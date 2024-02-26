// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/ForwarderPool.sol";
import "../src/interfaces/IGateway.sol";
import "../src/interfaces/IAssetForwarder.sol";
import "./utils/Utilities.sol";
import "./utils/DamnValuableToken.sol";
import "./utils/Interactor.sol";
import "./utils/Dummy.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";


contract ForwarderTest is Test {
    error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed);
    event ISendEvent(
        uint256 version,
        uint256 routeAmount,
        uint256 indexed eventNonce,
        address requestSender,
        string srcChainId,
        string destChainId,
        string routeRecipient,
        bytes requestMetadata,
        bytes requestPacket
    );
    event SetDappMetadataEvent(
        uint256 indexed eventNonce, 
        address dappAddress, 
        string chainId, 
        string feePayerAddress
        );



    ForwarderPool public forwarderPool;
    // the identifiers of the forks
    uint256 mumbaiFork;
    uint256 fujiFork;
    address private mumbaiAssetForwarder;
    address private fujiAssetForwarder;
    address private mumbaiGateway;
    string private alphaAssetForwarderMiddleware;
    
    Utilities internal utils;
    DamnValuableToken internal dvt;
    Interactor internal interactor;
    Dummy internal dummy;
    address internal owner;
    address internal forwarder1;
    address internal forwarder2;   

    string internal refundClaimer = "router19y0l9k97hmjrqtcqvenattetk7uxwhjfyk56lt";
    string internal feePayer = "router19y0l9k97hmjrqtcqvenattetk7uxwhjfyk56lt";


    string MUMBAI_RPC_URL = "https://rpc.ankr.com/polygon_mumbai";
    // string FUJI_RPC_URL = FUJI_RPC_URL;

    function setUp() public {
        // setting up the mock users
        utils = new Utilities();
        address payable[] memory users = utils.createUsers(3);
        owner = users[0];
        forwarder1 = users[1];
        forwarder2 = users[2];

        // setting up the fork chains
        mumbaiFork = vm.createFork(MUMBAI_RPC_URL);
        // vm.rollFork(45_331_671);
        // fujiFork = vm.createFork(FUJI_RPC_URL);
        mumbaiAssetForwarder= address(0xbdfA400EC4CFd2a3c2Ea2670C647127657970cEa);
        mumbaiGateway= address(0x423853a63DBAC343caE1E7aEE6b1c4488888dd13);
        alphaAssetForwarderMiddleware = "router12fykm2xhg5ces2vmf4q2aem8c958exv3v0wmvrspa8zucrdwjedsjnnxzt";

        // fujiAssetForwarder= address(0x8d1eec6aa6cd3b661ea177dc750c63874adae63c);
        vm.selectFork(mumbaiFork);

        interactor = new Interactor();
        dummy = new Dummy();
        // setting up some erc20 tokens
        dvt = new DamnValuableToken();
        vm.label(address(dvt), "DVT");
        dvt.transfer(forwarder1, 100 ether);
        dvt.transfer(forwarder2, 100 ether);
        dvt.transfer(owner, 100 ether);


        // setting up the forwarder pool
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        forwarderPool = new ForwarderPool(
            IAssetForwarder(mumbaiAssetForwarder),
            IGateway(mumbaiGateway),
            alphaAssetForwarderMiddleware
        );
        vm.stopPrank();
    }

    function testSetAssetForwarder() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        assertEq(address(forwarderPool.assetForwarder()), mumbaiAssetForwarder);
        vm.stopPrank();
        // vm.selectFork(optimismFork);
        // assertEq(vm.activeFork(), optimismFork);
        // forwarderPool.setAssetForwarder(fujiAssetForwarder);
        // assertEq(address(forwarderPool.assetForwarder()), fujiAssetForwarder);
    }

    function testSetGateway() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        assertEq(address(forwarderPool.gateway()), mumbaiGateway);
        vm.stopPrank();
        // vm.selectFork(optimismFork);
        // assertEq(vm.activeFork(), optimismFork);
        // forwarderPool.setGateway(fujiGateway);
        // assertEq(address(forwarderPool.gateway()), fujiGateway);
    }

    function testAlphaAssetForwarderMiddleware() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        assertEq(forwarderPool.assetForwarderMiddleware(), alphaAssetForwarderMiddleware);
        vm.stopPrank();
    }

    function testSetDappMetadata() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);

        vm.expectEmit(mumbaiGateway);
        uint256 requestNonce = IGateway(mumbaiGateway).eventNonce();

        emit SetDappMetadataEvent(
            requestNonce+1, // eventNonce
            address(forwarderPool),
            "80001", // mumbai chain id
            feePayer
        );

        forwarderPool.setDappMetadata(feePayer);
        // assertEq(forwarderPool.dappMetadata(), "metadata");
        vm.stopPrank();
    }

    function testSetWhitelistedForwarder() public {
        vm.selectFork(mumbaiFork);
        assertEq(vm.activeFork(), mumbaiFork);
        vm.startPrank(owner);
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        assertEq(forwarderPool.whitelistedForwarders(forwarder1),true);
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
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        vm.stopPrank();

        vm.startPrank(forwarder1);
        forwarderPool.approveAssetForwarder(address(dvt), 50 ether);
        assertEq(dvt.allowance(address(forwarderPool), address(mumbaiAssetForwarder)), 50 ether);
        vm.stopPrank();
    }

    function testSetForwarderRefundClaimer() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.startPrank(owner);


        uint64 start = 1708429322;
        uint64 expiryTime = start + 1000;

        bytes memory packet = abi.encode(uint64(expiryTime), refundClaimer);
        bytes memory requestPacket = abi.encode(alphaAssetForwarderMiddleware, packet);


        vm.expectEmit(mumbaiGateway);
        uint256 requestNonce = IGateway(mumbaiGateway).eventNonce();

        bytes memory zeroBytes = new bytes(0);


        emit ISendEvent(
            1,
            0,
            requestNonce+1, // eventNonce
            address(forwarderPool),
            "80001", // mumbai chain id
            "router_9625-1", // destChainId
            "", // routeRecipient
            zeroBytes, // requestMetadata
            requestPacket
        );

        forwarderPool.setForwarderRefundClaimer(uint64(expiryTime), refundClaimer, zeroBytes);
    }

    function testExecuteIRelayWithErc20() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.startPrank(owner);
        // token deposit
        dvt.approve(address(forwarderPool), 50 ether);
        forwarderPool.depositERC20(address(dvt), 50 ether);
        // whitelist forwarder
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(forwarder1);
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
        assert(success);
        (bool result) = abi.decode(returnData, (bool));
        assertEq(result, true);
        uint256 balanceAfter = dvt.balanceOf(address(0xa1));
        assertEq(balanceAfter - balanceBefore, 1000000);
        vm.stopPrank();
    }


    function testExceptRevertERC20InsufficientBalanceWhileExecuteIRelayWithErc20() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.startPrank(owner);
        // set contract state
        // token deposit
        dvt.approve(address(forwarderPool), 50);
        forwarderPool.depositERC20(address(dvt), 50);
        // whitelist forwarder
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(forwarder1);
        uint256 balanceBefore = dvt.balanceOf(address(0xa1));
        forwarderPool.approveAssetForwarder(address(dvt),1000000);
        vm.expectRevert(abi.encodeWithSelector(ERC20InsufficientBalance.selector,address(forwarderPool), 50, 1000000));

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
        assert(success);
        (bool result) = abi.decode(returnData, (bool));
        assertEq(result, false);
        uint256 balanceAfter = dvt.balanceOf(address(0xa1));
        assertEq(balanceAfter , balanceBefore);
        vm.stopPrank();
    }
    function testExecuteIRelayWithNativeToken() public {
        // set-up
        vm.selectFork(mumbaiFork);
        vm.deal(address(owner), 1 ether);
        vm.startPrank(owner);
        // set contract state
        // token deposit
        forwarderPool.depositNativeToken{value: 1 ether}();
        // whitelist forwarder
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(forwarder1);
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
        assert(success);
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
        // token deposit
        dvt.approve(address(forwarderPool), 50 ether);
        forwarderPool.depositERC20(address(dvt), 50 ether);
        // whitelist forwarder
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(forwarder1);
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
        assert(success);
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
        // token deposit
        forwarderPool.depositNativeToken{value: 1 ether}();
        // whitelist forwarder
        forwarderPool.setWhitelistedForwarder(forwarder1, true);
        vm.stopPrank();

        // execute the relay
        vm.startPrank(forwarder1);
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
        assert(success);
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