// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/ForwarderPool.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../src/interfaces/IGateway.sol";
import "../src/interfaces/IAssetForwarder.sol";


contract ForwarderPoolScript is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address assetForwarder= address(0xbdfA400EC4CFd2a3c2Ea2670C647127657970cEa);
        address gateway= address(0x423853a63DBAC343caE1E7aEE6b1c4488888dd13);
        string memory alphaAssetForwarderMiddleware = "router12fykm2xhg5ces2vmf4q2aem8c958exv3v0wmvrspa8zucrdwjedsjnnxzt";


        string memory refundClaimer = "router19y0l9k97hmjrqtcqvenattetk7uxwhjfyk56lt";
        string memory feePayer = "router19y0l9k97hmjrqtcqvenattetk7uxwhjfyk56lt";

        bytes memory zeroBytes = new bytes(0);

        uint64 start = 1708470322;
        uint64 expiryTime = start + 1000;

        ERC20 usdt = ERC20(0x22bAA8b6cdd31a0C5D1035d6e72043f4Ce6aF054);
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);
        ForwarderPool forwarderPool = new ForwarderPool(
            IAssetForwarder(assetForwarder),
            IGateway(gateway),
            alphaAssetForwarderMiddleware
        );
        console.log("ForwarderPool address: %s", address(forwarderPool));
        // setting forwarderPool fee payer
        forwarderPool.setDappMetadata(feePayer);
        // setting forwarderPool refund claimer
        forwarderPool.setForwarderRefundClaimer(uint64(expiryTime), refundClaimer, zeroBytes);
        // setting forwarderPool whitelisted forwarders
        forwarderPool.setWhitelistedForwarder(deployer, true);
        forwarderPool.setWhitelistedForwarder(address(0xcbA625Dc657CB92D85BcD2c6573605e3ea6E61a9), true); // actual forwarer in alpha-devnet
        console.log("ForwarderPool whitelistedForwarders: %s", forwarderPool.whitelistedForwarders(deployer));
        usdt.approve(address(forwarderPool),10_000_000_000_000);
        forwarderPool.depositERC20(address(usdt),10_000_000_000_000);
        forwarderPool.approveAssetForwarder(address(usdt),10_000_000_000_000);
        vm.stopBroadcast();
    }
}

// https://mumbai.polygonscan.com/address/0x6d277d651ca4b4684bda699bba74b685e11a48d1