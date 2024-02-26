# ForwarderPool
[Git Source](https://github.com/router-protocol/forwarder-pool-contract/blob/0450db7d2136e0474953942b8882399767544265/src/ForwarderPool.sol)

**Inherits:**
Ownable

**Author:**
Router Protocol

Pooling funds to be used by set of whitelisted forwarders


## State Variables
### erc20Deposits

```solidity
mapping(address => uint256) private erc20Deposits;
```


### nativeTokenDeposits

```solidity
mapping(address => uint256) private nativeTokenDeposits;
```


### whitelistedForwarders

```solidity
mapping(address => bool) public whitelistedForwarders;
```


### assetForwarder

```solidity
IAssetForwarder public assetForwarder;
```


### gateway

```solidity
IGateway public gateway;
```


### assetForwarderMiddleware

```solidity
string public assetForwarderMiddleware;
```


### destChainId

```solidity
string private constant destChainId = "router_9625-1";
```


### NATIVE_ADDRESS

```solidity
address private constant NATIVE_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
```


## Functions
### constructor

Initializes a new instance of the contract setting up key components necessary for its operation.
This includes setting up the asset forwarder, gateway, and asset forwarder middleware.
The constructor also establishes the contract's owner upon deployment.


```solidity
constructor(IAssetForwarder _assetForwarder, IGateway _gateway, string memory _assetForwarderMiddleware)
    Ownable(msg.sender);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_assetForwarder`|`IAssetForwarder`|The asset forwarder interface that handles the forwarding of assets.|
|`_gateway`|`IGateway`|The gateway interface to interact with different blockchain protocols.|
|`_assetForwarderMiddleware`|`string`|A string representing middleware used in asset forwarding.|


### setAssetForwarder

Updates the asset forwarder contract address. Only callable by the contract owner.


```solidity
function setAssetForwarder(address _assetForwarder) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_assetForwarder`|`address`|The new address of the IAssetForwarder contract.|


### setGateway

Updates the gateway contract address. Only callable by the contract owner.


```solidity
function setGateway(address _gateway) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_gateway`|`address`|The new address of the IGateway contract.|


### setAssetForwarderMiddleware

Updates the asset forwarder middleware configuration. Only callable by the contract owner.


```solidity
function setAssetForwarderMiddleware(string memory _assetForwarderMiddleware) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_assetForwarderMiddleware`|`string`|The new middleware configuration for the asset forwarder.|


### setWhitelistedForwarder

Updates the whitelist status of a forwarder address. Only callable by the contract owner.


```solidity
function setWhitelistedForwarder(address forwarder, bool whitelisted) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`forwarder`|`address`|The address of the forwarder to update.|
|`whitelisted`|`bool`|Boolean indicating whether the forwarder is whitelisted.|


### onlyWhitelistedForwarder

Modifier to restrict function access to only whitelisted forwarders.


```solidity
modifier onlyWhitelistedForwarder();
```

### depositERC20

Allows a anyone to deposit ERC20.


```solidity
function depositERC20(address token, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the ERC20 token to deposit.|
|`amount`|`uint256`|The amount of tokens to deposit.|


### depositNativeToken

Allows a anyone to deposit native tokens (e.g., ETH).


```solidity
function depositNativeToken() external payable;
```

### setForwarderRefundClaimer

Sets the refund claimer for the forwarder.


```solidity
function setForwarderRefundClaimer(uint64 expiryTime, string calldata refundClaimer, bytes calldata requestMetadata)
    public
    payable
    onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`expiryTime`|`uint64`|The expiry time for the refund claimer in seconds.|
|`refundClaimer`|`string`|The router chain address of the refund claimer.|
|`requestMetadata`|`bytes`|Additional metadata for the request. Like router gas.|


### setDappMetadata

Set's the router chain fee payer


```solidity
function setDappMetadata(string memory feePayerAddress) external onlyOwner;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`feePayerAddress`|`string`|The router chain address of the fee payer.|


### approveAssetForwarder

Approves the asset forwarder to spend a certain amount of tokens on behalf of the contract.


```solidity
function approveAssetForwarder(address token, uint256 amount) external onlyWhitelistedForwarder;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token to approve.|
|`amount`|`uint256`|The amount of tokens to approve.|


### iRelay

This function relays assets using the provided relay data.


```solidity
function iRelay(IAssetForwarder.RelayData memory relayData) external onlyWhitelistedForwarder;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`relayData`|`IAssetForwarder.RelayData`|The data required to perform the relay operation.|


### iRelayMessage

This function relays a message using the provided relay message data.


```solidity
function iRelayMessage(IAssetForwarder.RelayDataMessage memory relayMessageData) external onlyWhitelistedForwarder;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`relayMessageData`|`IAssetForwarder.RelayDataMessage`|The data required to perform the relay operation.|


### isNative


```solidity
function isNative(address token) internal pure returns (bool);
```

### withdrawNativeToken


```solidity
function withdrawNativeToken(address payable recipient, uint256 amount) external onlyOwner;
```

### withdrawERC20


```solidity
function withdrawERC20(address token, address recipient, uint256 amount) external onlyOwner;
```

