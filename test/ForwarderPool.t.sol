// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/ForwarderPool.sol";
import "../src/interfaces/IAssetForwarder.sol";
import "./utils/Utilities.sol";
import "./utils/DamnValuableToken.sol";
import "./utils/Interactor.sol";
import "./utils/Dummy.sol";

contract ForwarderTest is Test {
    ForwarderPool public forwarderPool;
    // the identifiers of the forks
    uint256 mumbaiFork;
    uint256 fujiFork;
    address private mumbaiAssetForwarder;
    address private fujiAssetForwarder;
    
    Utilities internal utils;
    DamnValuableToken internal dvt;
    Interactor internal interactor;
    Dummy internal dummy;
    address internal owner;
    address internal filler1;
    address internal filler2;   


    string MUMBAI_RPC_URL = "https://rpc.ankr.com/polygon_mumbai";
    // string FUJI_RPC_URL = FUJI_RPC_URL;

    function setUp() public {
        // setting up the mock users
        utils = new Utilities();
        address payable[] memory users = utils.createUsers(3);
        owner = users[0];
        filler1 = users[1];
        filler2 = users[2];

        // setting up the fork chains
        mumbaiFork = vm.createFork(MUMBAI_RPC_URL);
        // vm.rollFork(45_331_671);
        // fujiFork = vm.createFork(FUJI_RPC_URL);
        mumbaiAssetForwarder= address(0xbdfA400EC4CFd2a3c2Ea2670C647127657970cEa);
        // fujiAssetForwarder= address(0x8d1eec6aa6cd3b661ea177dc750c63874adae63c);
        vm.selectFork(mumbaiFork);

        interactor = new Interactor();
        dummy = new Dummy();
        // setting up some erc20 tokens
        dvt = new DamnValuableToken();
        vm.label(address(dvt), "DVT");
        dvt.transfer(filler1, 100 ether);
        dvt.transfer(filler2, 100 ether);
        dvt.transfer(owner, 100 ether);


        // setting up the forwarder pool
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        forwarderPool = new ForwarderPool();
        vm.stopPrank();
    }

    function testSetAssetForwarder() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        assertEq(address(forwarderPool.assetForwarder()), mumbaiAssetForwarder);
        vm.stopPrank();
        // vm.selectFork(optimismFork);
        // assertEq(vm.activeFork(), optimismFork);
        // forwarderPool.setAssetForwarder(fujiAssetForwarder);
        // assertEq(address(forwarderPool.assetForwarder()), fujiAssetForwarder);
    }

    function testSetWhitelistedFiller() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        forwarderPool.setWhitelistedFiller(filler1, true);
        assertEq(forwarderPool.whitelistedFillers(filler1),true);
        vm.stopPrank();
    }

    function testDepositERC20() public {
        vm.selectFork(mumbaiFork);
        dvt.approve(address(forwarderPool), 50);
        forwarderPool.depositERC20(address(dvt), 50);
        assertEq(dvt.balanceOf(address(forwarderPool)), 50);
    }

    function testDepositNativeToken() public {
        vm.selectFork(mumbaiFork);
        uint256 balanceBefore = address(forwarderPool).balance;
        forwarderPool.depositNativeToken{value: 1 ether}();
        uint256 balanceAfter = address(forwarderPool).balance;
        assertEq(balanceAfter - balanceBefore, 1 ether);
    }

    function testApproveAssetForwarder() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        forwarderPool.setWhitelistedFiller(filler1, true);
        vm.stopPrank();
        
        vm.startPrank(filler1);
        forwarderPool.approveAssetForwarder(address(dvt), 50 ether);
        assertEq(dvt.allowance(address(forwarderPool), address(mumbaiAssetForwarder)), 50 ether);
        vm.stopPrank();
    }

    function testExecuteIRelayWithErc20() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.startPrank(owner);
        // set contract state
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        // token deposit
        dvt.approve(address(forwarderPool), 50 ether);
        forwarderPool.depositERC20(address(dvt), 50 ether);
        // whitelist filler
        forwarderPool.setWhitelistedFiller(filler1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(filler1);
        uint256 balanceBefore = dvt.balanceOf(address(0xa1));
        forwarderPool.approveAssetForwarder(address(dvt),1000000);
        forwarderPool.iRelay(
            IAssetForwarder.RelayData(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(dvt),
                address(0xa1)
            )
        );
        bytes32 messageHash = keccak256(
            abi.encode(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(dvt),
                address(0xa1),
                address(mumbaiAssetForwarder)
            )
        );
        (bool success, bytes memory returnData)=mumbaiAssetForwarder.call(abi.encodeWithSignature("executeRecord(bytes32)", messageHash));
        (bool result) = abi.decode(returnData, (bool));
        assertEq(result, true);
        uint256 balanceAfter = dvt.balanceOf(address(0xa1));
        assertEq(balanceAfter - balanceBefore, 1000000);
        vm.stopPrank();
    }

    function testExecuteIRelayWithNativeToken() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.deal(address(owner), 1 ether);
        vm.startPrank(owner);
        // set contract state
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        // token deposit
        forwarderPool.depositNativeToken{value: 1 ether}();
        // whitelist filler
        forwarderPool.setWhitelistedFiller(filler1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(filler1);
        uint256 balanceBefore = address(0xa1).balance;

        forwarderPool.iRelay(
            IAssetForwarder.RelayData(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE),
                address(0xa1)
            )
        );
        bytes32 messageHash = keccak256(
            abi.encode(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE),
                address(0xa1),
                address(mumbaiAssetForwarder)
            )
        );
        (bool success, bytes memory returnData)=mumbaiAssetForwarder.call(abi.encodeWithSignature("executeRecord(bytes32)", messageHash));
        (bool result) = abi.decode(returnData, (bool));
        assertEq(result, true);
        uint256 balanceAfter = address(0xa1).balance;
        assertEq(balanceAfter - balanceBefore, 1000000);
        vm.stopPrank();            
    }    

    function testExecuteIRelayMessageWithErc20() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.startPrank(owner);
        // set contract state
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        // token deposit
        dvt.approve(address(forwarderPool), 50 ether);
        forwarderPool.depositERC20(address(dvt), 50 ether);
        // whitelist filler
        forwarderPool.setWhitelistedFiller(filler1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(filler1);
        uint256 balanceBefore = dvt.balanceOf(address(interactor));
        //message payload
        bytes memory payload = abi.encode(address(dummy), bytes("message"));
        forwarderPool.approveAssetForwarder(address(dvt),1000000);
        forwarderPool.iRelayMessage(
            IAssetForwarder.RelayDataMessage(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(dvt),
                address(interactor),
                payload
            )
        );
        bytes32 messageHash = keccak256(
            abi.encode(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(dvt),
                address(interactor),
                address(mumbaiAssetForwarder),
                payload
            )
        );
        (bool success, bytes memory returnData)=mumbaiAssetForwarder.call(abi.encodeWithSignature("executeRecord(bytes32)", messageHash));
        (bool result) = abi.decode(returnData, (bool));
        assertEq(result, true);
        uint256 balanceAfter = dvt.balanceOf(address(interactor));
        assertEq(balanceAfter - balanceBefore, 1000000);
        //expect count of dummy contract to be incremented
        uint count = dummy.count();
        assertEq(count, 1);
        vm.stopPrank();
    }

    function testExecuteIRelayMessageWithNativeToken() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.deal(address(owner), 1 ether);
        vm.startPrank(owner);
        // set contract state
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        // token deposit
        forwarderPool.depositNativeToken{value: 1 ether}();
        // whitelist filler
        forwarderPool.setWhitelistedFiller(filler1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(filler1);
        uint256 balanceBefore = address(interactor).balance;
        //message payload
        bytes memory payload = abi.encode(address(dummy), bytes("message"));

        forwarderPool.iRelayMessage(
            IAssetForwarder.RelayDataMessage(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE),
                address(interactor),
                payload
            )
        );
        bytes32 messageHash = keccak256(
            abi.encode(
                1000000,
                0x6f736d6f2d746573742d35000000000000000000000000000000000000000000,
                186,
                address(0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE),
                address(interactor),
                address(mumbaiAssetForwarder),
                payload
            )
        );
        (bool success, bytes memory returnData)=mumbaiAssetForwarder.call(abi.encodeWithSignature("executeRecord(bytes32)", messageHash));
        (bool result) = abi.decode(returnData, (bool));
        assertEq(result, true);
        uint256 balanceAfter = address(interactor).balance;
        assertEq(balanceAfter - balanceBefore, 1000000);
        //expect count of dummy contract to be incremented
        uint count = dummy.count();
        assertEq(count, 1);
        vm.stopPrank();
    }
}