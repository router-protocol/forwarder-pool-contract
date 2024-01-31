// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/ForwarderPool.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";


contract ForwarderPoolScript is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address mumbaiAssetForwarder= address(0xbdfA400EC4CFd2a3c2Ea2670C647127657970cEa);
        ERC20 usdt = ERC20(0x22bAA8b6cdd31a0C5D1035d6e72043f4Ce6aF054);
        address deployer = vm.addr(deployerPrivateKey);
        console.log("Deployer address: %s", deployer);
        ForwarderPool forwarderPool = new ForwarderPool();
        console.log("ForwarderPool address: %s", address(forwarderPool));
        forwarderPool.setAssetForwarder(mumbaiAssetForwarder);
        // console.log("ForwarderPool assetForwarder: %s", forwarderPool.assetForwarder());
        forwarderPool.setWhitelistedFiller(deployer, true);
        forwarderPool.setWhitelistedFiller(address(0xcbA625Dc657CB92D85BcD2c6573605e3ea6E61a9), true); // actual forwarer in alpha-devnet
        // console.log("ForwarderPool whitelistedFillers: %s", forwarderPool.whitelistedFillewrs(deployer));
        usdt.approve(address(forwarderPool),10_000_000_000_000);
        forwarderPool.depositERC20(address(usdt),10_000_000_000_000);
        forwarderPool.approveAssetForwarder(address(usdt),10_000_000_000_000);
        vm.stopBroadcast();
    }
}
