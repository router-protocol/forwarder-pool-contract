# IAssetForwarder
[Git Source](https://github.com/router-protocol/forwarder-pool-contract/blob/0450db7d2136e0474953942b8882399767544265/src/interfaces/IAssetForwarder.sol)

**Author:**
Router Protocol.


## Functions
### iDepositUSDC


```solidity
function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount)
    external
    payable;
```

### iDeposit


```solidity
function iDeposit(DepositData memory depositData, bytes memory destToken, bytes memory recipient) external payable;
```

### iDepositInfoUpdate


```solidity
function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal)
    external
    payable;
```

### iDepositMessage


```solidity
function iDepositMessage(
    DepositData memory depositData,
    bytes memory destToken,
    bytes memory recipient,
    bytes memory message
) external payable;
```

### iRelay


```solidity
function iRelay(RelayData memory relayData) external payable;
```

### iRelayMessage


```solidity
function iRelayMessage(RelayDataMessage memory relayData) external payable;
```

## Events
### FundsDeposited

```solidity
event FundsDeposited(
    uint256 partnerId,
    uint256 amount,
    bytes32 destChainIdBytes,
    uint256 destAmount,
    uint256 depositId,
    address srcToken,
    address depositor,
    bytes recipient,
    bytes destToken
);
```

### iUSDCDeposited

```solidity
event iUSDCDeposited(
    uint256 partnerId,
    uint256 amount,
    bytes32 destChainIdBytes,
    uint256 usdcNonce,
    address srcToken,
    bytes32 recipient,
    address depositor
);
```

### FundsDepositedWithMessage

```solidity
event FundsDepositedWithMessage(
    uint256 partnerId,
    uint256 amount,
    bytes32 destChainIdBytes,
    uint256 destAmount,
    uint256 depositId,
    address srcToken,
    bytes recipient,
    address depositor,
    bytes destToken,
    bytes message
);
```

### FundsPaid

```solidity
event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce);
```

### DepositInfoUpdate

```solidity
event DepositInfoUpdate(
    address srcToken,
    uint256 feeAmount,
    uint256 depositId,
    uint256 eventNonce,
    bool initiatewithdrawal,
    address depositor
);
```

### FundsPaidWithMessage

```solidity
event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData);
```

## Structs
### DestDetails

```solidity
struct DestDetails {
    uint32 domainId;
    uint256 fee;
    bool isSet;
}
```

### RelayData

```solidity
struct RelayData {
    uint256 amount;
    bytes32 srcChainId;
    uint256 depositId;
    address destToken;
    address recipient;
}
```

### RelayDataMessage

```solidity
struct RelayDataMessage {
    uint256 amount;
    bytes32 srcChainId;
    uint256 depositId;
    address destToken;
    address recipient;
    bytes message;
}
```

### DepositData

```solidity
struct DepositData {
    uint256 partnerId;
    uint256 amount;
    uint256 destAmount;
    address srcToken;
    address refundRecipient;
    bytes32 destChainIdBytes;
}
```

