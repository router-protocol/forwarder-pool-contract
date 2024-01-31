// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./interfaces/IAssetForwarder.sol";
// import "./Multicaller.sol";

contract ForwarderPool is Ownable {

    constructor() Ownable(msg.sender) {

    }

    mapping(address => uint256) private erc20Deposits;
    mapping(address => uint256) private nativeTokenDeposits;
    mapping(address => bool) public whitelistedFillers;
    IAssetForwarder public assetForwarder;

    address private constant NATIVE_ADDRESS =
        0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    function setAssetForwarder(address _assetForwarder) external onlyOwner {
        assetForwarder = IAssetForwarder(_assetForwarder);
    }

    function setWhitelistedFiller(address filler, bool whitelisted) external onlyOwner {
        whitelistedFillers[filler] = whitelisted;
    }

    modifier onlyWhitelistedFiller() {
        require(whitelistedFillers[msg.sender], "ForwarderPool: caller is not whitelisted filler");
        _;
    }

    function depositERC20(address token, uint256 amount) external {
        IERC20(token).transferFrom(msg.sender, address(this), amount);
        erc20Deposits[token] += amount;
    }

    function depositNativeToken() external payable {
        nativeTokenDeposits[msg.sender] += msg.value;
    }

    function iRelay(IAssetForwarder.RelayData memory relayData) external onlyWhitelistedFiller {
        require(assetForwarder != IAssetForwarder(address(0)), "ForwarderPool: assetForwarder not set");
        if (isNative(relayData.destToken)) {
            require(
                address(this).balance >= relayData.amount,
                "ForwarderPool: insufficient native token balance"
            );
            assetForwarder.iRelay{value: relayData.amount}(relayData);
        } else {
            require(erc20Deposits[relayData.destToken] >= relayData.amount, "ForwarderPool: insufficient ERC20 deposit");
            erc20Deposits[relayData.destToken] -= relayData.amount;
            // approve erc20 token to assetForwarder
            // IERC20(relayData.destToken).approve(address(assetForwarder), relayData.amount);
            assetForwarder.iRelay(relayData);
        }
    }

    function approveAssetForwarder(address token, uint256 amount) external onlyWhitelistedFiller {
        IERC20(token).approve(address(assetForwarder), amount);
    }

    function iRelayMessage(IAssetForwarder.RelayDataMessage memory relayMessageData) external onlyWhitelistedFiller {
        require(assetForwarder != IAssetForwarder(address(0)), "ForwarderPool: assetForwarder not set");
        if (isNative(relayMessageData.destToken)) {
            require(
                address(this).balance >= relayMessageData.amount,
                "ForwarderPool: insufficient native token balance"
            );
        } else {
            require(erc20Deposits[relayMessageData.destToken] >= relayMessageData.amount, "ForwarderPool: insufficient ERC20 deposit");
            erc20Deposits[relayMessageData.destToken] -= relayMessageData.amount;
            // approve erc20 token to assetForwarder
            // IERC20(relayMessageData.destToken).approve(address(assetForwarder), relayMessageData.amount);
        }
        assetForwarder.iRelayMessage(relayMessageData);
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
