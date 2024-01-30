// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;
import "../../src/interfaces/IMessageHandler.sol";

contract Interactor {
    function handleMessage(
        address tokenSent,
        uint256 amount,
        bytes memory message
    ) public payable {
        (address stakingContract, bytes memory data) = abi.decode(
            message,
            (address, bytes)
        );

        bytes memory payload = abi.encodeWithSelector(0x2d1e0c02, data);
        (bool success, bytes memory returnData) = address(stakingContract).call(
            payload
        );
    }
}