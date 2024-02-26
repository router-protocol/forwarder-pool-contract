# MultiCaller
[Git Source](https://github.com/router-protocol/forwarder-pool-contract/blob/0450db7d2136e0474953942b8882399767544265/src/Multicaller.sol)

Logic is 100% copied from "https://github.com/Uniswap/v3-periphery/blob/main/contracts/base/Multicall.sol" just version upgraded
please check if this can create conflict with handling msg.value


## Functions
### multicall


```solidity
function multicall(bytes[] calldata data) external returns (bytes[] memory results);
```

