// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IAssetForwarder.sol";
import "./Multicaller.sol";

contract ForwarderPool is Ownable, Multicaller {
    mapping(address => uint256) private erc20Deposits;
    mapping(address => uint256) private nativeTokenDeposits;
    mapping(address => bool) private whitelistedFillers;
    IAssetForwarder private assetForwarder;

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

    function executeIRelay(RelayData memory relayData) external onlyWhitelistedFiller {
        require(assetForwarder != IAssetForwarder(address(0)), "ForwarderPool: assetForwarder not set");
        if (isNative(relayData.destToken)) {
            require(
                nativeTokenDeposits[address(this)] >= relayData.amount,
                "ForwarderPool: insufficient native token balance"
            );
        } else {
            require(erc20Deposits[relayData.srcToken] >= relayData.amount, "ForwarderPool: insufficient ERC20 deposit");
            erc20Deposits[relayData.srcToken] -= relayData.amount;
            // approve erc20 token to assetForwarder
            IERC20(relayData.srcToken).approve(address(assetForwarder), relayData.amount);
        }
        assetForwarder.execute(relayData);
    }

    function executeIRelayMessage(RelayMessageData memory relayMessageData) external onlyWhitelistedFiller {
        require(assetForwarder != IAssetForwarder(address(0)), "ForwarderPool: assetForwarder not set");
        if (isNative(relayMessageData.destToken)) {
            require(
                nativeTokenDeposits[address(this)] >= relayMessageData.amount,
                "ForwarderPool: insufficient native token balance"
            );
        } else {
            require(erc20Deposits[relayMessageData.srcToken] >= relayMessageData.amount, "ForwarderPool: insufficient ERC20 deposit");
            erc20Deposits[relayMessageData.srcToken] -= relayMessageData.amount;
            // approve erc20 token to assetForwarder
            IERC20(relayMessageData.srcToken).approve(address(assetForwarder), relayMessageData.amount);
        }
        assetForwarder.executeMessage(relayMessageData);
    }

    function isNative(address token) internal pure returns (bool) {
        return token == NATIVE_ADDRESS;
    }
}
