// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/ForwarderPool.sol";

contract ForwarderPoolScript is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address mumbaiAssetForwarder= address(0xbdfA400EC4CFd2a3c2Ea2670C647127657970cEa);
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);
        ForwarderPool forwarderPool = new ForwarderPool();
        console.log("ForwarderPool address: %s", address(forwarderPool));
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        // console.log("ForwarderPool assetForwarder: %s", forwarderPool.assetForwarder());
        forwarderPool.setWhitelistedFiller(deployer, true);
        // console.log("ForwarderPool whitelistedFillers: %s", forwarderPool.whitelistedFillers(deployer));
        vm.stopBroadcast();
    }
}
