// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ForwarderPool is Ownable {
    mapping(address => uint256) private erc20Deposits;
    mapping(address => uint256) private nativeTokenDeposits;

    function depositERC20(address token, uint256 amount) external {
        IERC20(token).transferFrom(msg.sender, address(this), amount);
        erc20Deposits[token] += amount;
    }

    function depositNativeToken() external payable {
        nativeTokenDeposits[msg.sender] += msg.value;
    }

    function multicall(address contractAddress, uint256 amount) external onlyOwner {
        IERC20(contractAddress).approve(contractAddress, amount);
        (bool success, ) = contractAddress.call(abi.encodeWithSignature("execute()"));
        require(success, "Multicall failed");
    }
}
