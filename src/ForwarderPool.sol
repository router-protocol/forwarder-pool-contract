// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IAssetForwarder.sol";
import "./interfaces/IGateway.sol";
// import "./Multicaller.sol";

/**
 * @title ForwarderPool
 * @author Router Protocol
 * @notice Pooling funds to be used by set of whitelisted forwarders 
 * @custom:version v0.5.0
 */
contract ForwarderPool is Ownable {

    /** 
     * Initializes a new instance of the contract setting up key components necessary for its operation.
     * This includes setting up the asset forwarder, gateway, and asset forwarder middleware.
     * The constructor also establishes the contract's owner upon deployment.
     * 
     * @param _assetForwarder The asset forwarder interface that handles the forwarding of assets.
     * @param _gateway The gateway interface to interact with different blockchain protocols.
     * @param _assetForwarderMiddleware A string representing middleware used in asset forwarding.
     */
    constructor(
        IAssetForwarder _assetForwarder,
        IGateway _gateway,
        string memory _assetForwarderMiddleware
    ) Ownable(msg.sender) {
        assetForwarder = _assetForwarder;
        gateway = _gateway;
        assetForwarderMiddleware = _assetForwarderMiddleware;
    }

    mapping(address => uint256) private erc20Deposits;
    mapping(address => uint256) private nativeTokenDeposits;
    mapping(address => bool) public whitelistedForwarders;
    IAssetForwarder public assetForwarder;
    IGateway public gateway;
    string public assetForwarderMiddleware;

    string private constant destChainId = "router_9625-1"; // this will be depended on router enviroment

    address private constant NATIVE_ADDRESS =
        0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /**
     * Updates the asset forwarder contract address. Only callable by the contract owner.
     *
     * @param _assetForwarder The new address of the IAssetForwarder contract.
     */
    function setAssetForwarder(address _assetForwarder) external onlyOwner {
        assetForwarder = IAssetForwarder(_assetForwarder);
    }

    /**
     * Updates the gateway contract address. Only callable by the contract owner.
     *
     * @param _gateway The new address of the IGateway contract.
     */
    function setGateway(address _gateway) external onlyOwner {
        gateway = IGateway(_gateway);
    }

    /**
     * Updates the asset forwarder middleware configuration. Only callable by the contract owner.
     *
     * @param _assetForwarderMiddleware The new middleware configuration for the asset forwarder.
     */
    function setAssetForwarderMiddleware(string memory _assetForwarderMiddleware) external onlyOwner {
        assetForwarderMiddleware = _assetForwarderMiddleware;
    }

    /**
     * Updates the whitelist status of a forwarder address. Only callable by the contract owner.
     *
     * @param forwarder The address of the forwarder to update.
     * @param whitelisted Boolean indicating whether the forwarder is whitelisted.
     */
    function setWhitelistedForwarder(address forwarder, bool whitelisted) external onlyOwner {
        whitelistedForwarders[forwarder] = whitelisted;
    }

    /**
     * Modifier to restrict function access to only whitelisted forwarders.
     */
    modifier onlyWhitelistedForwarder() {
        require(whitelistedForwarders[msg.sender], "ForwarderPool: caller is not whitelisted forwarder");
        _;
    }


    /**
     * Allows a anyone to deposit ERC20.
     *
     * @param token The address of the ERC20 token to deposit.
     * @param amount The amount of tokens to deposit.
     */
    function depositERC20(address token, uint256 amount) external {
        IERC20(token).transferFrom(msg.sender, address(this), amount);
        erc20Deposits[token] += amount;
    }

    /**
     * Allows a anyone to deposit native tokens (e.g., ETH).
     */
    function depositNativeToken() external payable {
        nativeTokenDeposits[msg.sender] += msg.value;
    }

    /**
     * Sets the refund claimer for the forwarder.
     * 
     * @param expiryTime The expiry time for the refund claimer in seconds.
     * @param refundClaimer The router chain address of the refund claimer.
     * @param requestMetadata Additional metadata for the request. Like router gas.
     */
    function setForwarderRefundClaimer(
        uint64 expiryTime,
        string calldata refundClaimer,
        bytes calldata requestMetadata
    ) public onlyOwner payable {
        // Encode the expiry time and refund claimer into a byte array
        bytes memory packet = abi.encode(expiryTime, refundClaimer);

        // Encode the asset forwarder middleware and the previously created packet into a byte array
        bytes memory requestPacket = abi.encode(assetForwarderMiddleware, packet);

        // Send the request packet using the gateway contract
        gateway.iSend{ value: msg.value }(
          1,
          0,
          string(""),
          destChainId,
          requestMetadata,
          requestPacket
        );
    }

    /**
     * Set's the router chain fee payer 
     * 
     * @param feePayerAddress The router chain address of the fee payer.
     */
    function setDappMetadata(string memory feePayerAddress) external onlyOwner {
      // Set the Dapp metadata in the gateway contract
      gateway.setDappMetadata(feePayerAddress);
    }

    /**
     * Approves the asset forwarder to spend a certain amount of tokens on behalf of the contract.
     * 
     * @param token The address of the token to approve.
     * @param amount The amount of tokens to approve.
     */
    function approveAssetForwarder(address token, uint256 amount) external onlyWhitelistedForwarder {
        // Approve the asset forwarder to spend the specified amount of tokens
        IERC20(token).approve(address(assetForwarder), amount);
    }

    /**
     * This function relays assets using the provided relay data.
     * 
     * @param relayData The data required to perform the relay operation.
     */
    function iRelay(IAssetForwarder.RelayData memory relayData) external onlyWhitelistedForwarder {
        // Ensure that the asset forwarder is set
        require(assetForwarder != IAssetForwarder(address(0)), "ForwarderPool: assetForwarder not set");

        // If the destination token is native (like ETH), call the iRelay function on the assetForwarder
        // contract with the relay amount as value.
        if (isNative(relayData.destToken)) {
            assetForwarder.iRelay{value: relayData.amount}(relayData);
        } else {
            // If the destination token is not native, call the iRelay function on the assetForwarder
            // contract without sending any value.
            assetForwarder.iRelay(relayData);
        }
    }

    /**
     * This function relays a message using the provided relay message data.
     * 
     * @param relayMessageData The data required to perform the relay operation.
     */
    function iRelayMessage(IAssetForwarder.RelayDataMessage memory relayMessageData) external onlyWhitelistedForwarder {
        // Ensure that the asset forwarder is set
        require(assetForwarder != IAssetForwarder(address(0)), "ForwarderPool: assetForwarder not set");

        // If the destination token is native (like ETH), check if the contract has enough balance
        // and then call the iRelayMessage function on the assetForwarder contract with the relay amount as value.
        if (isNative(relayMessageData.destToken)) {
            require(
                address(this).balance >= relayMessageData.amount,
                "ForwarderPool: insufficient native token balance"
            );
            assetForwarder.iRelayMessage{value: relayMessageData.amount}(relayMessageData);
        } else {
            // If the destination token is not native, check if the contract has enough ERC20 deposits
            // and then call the iRelayMessage function on the assetForwarder contract without sending any value.
            require(erc20Deposits[relayMessageData.destToken] >= relayMessageData.amount, "ForwarderPool: insufficient ERC20 deposit");
            erc20Deposits[relayMessageData.destToken] -= relayMessageData.amount;
            // The following line is commented out, but it would approve the asset forwarder to spend the specified amount of ERC20 tokens
            // IERC20(relayMessageData.destToken).approve(address(assetForwarder), relayMessageData.amount);
            assetForwarder.iRelayMessage(relayMessageData);
        }
    }

    function isNative(address token) internal pure returns (bool) {
        return token == NATIVE_ADDRESS;
    }

    function withdrawNativeToken(address payable recipient, uint256 amount) external onlyOwner {
        recipient.transfer(amount);
    }

    function withdrawERC20(address token, address recipient, uint256 amount) external onlyOwner {
        IERC20(token).transfer(recipient, amount);
    }
}
