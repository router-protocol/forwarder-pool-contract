# IGateway
[Git Source](https://github.com/router-protocol/forwarder-pool-contract/blob/0450db7d2136e0474953942b8882399767544265/src/interfaces/IGateway.sol)

from

*Interface of the Gateway Self External Calls.*


## Functions
### iSend


```solidity
function iSend(
    uint256 version,
    uint256 routeAmount,
    string calldata routeRecipient,
    string calldata destChainId,
    bytes calldata requestMetadata,
    bytes calldata requestPacket
) external payable returns (uint256);
```

### setDappMetadata


```solidity
function setDappMetadata(string memory feePayerAddress) external payable returns (uint256);
```

### currentVersion


```solidity
function currentVersion() external view returns (uint256);
```

### eventNonce


```solidity
function eventNonce() external view returns (uint256);
```

