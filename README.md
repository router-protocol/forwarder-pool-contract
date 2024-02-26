# Documentation for ForwarderPool Contract and Test Cases
### Overview
This is ForwarderPool contract facilitates the forwarding of both ERC20 tokens and native blockchain tokens for Router Nitro.

### Contract Features
- **ERC20 and Native Token Management:** Handles deposits of ERC20 and native tokens, tracking the balances internally.
- **Asset Forwarding:** Uses an IAssetForwarder interface to forward assets.
- **Whitelisting:** Provides the ability to whitelist addresses, allowing only authorized addresses to forwarder tokens.
- **Owner Privileges:** Includes functions that can only be executed by the contract owner, like withdrawing tokens and managing whitelisted addresses.


## Installation:

Install dependencies using 

```
forge install
```

## Compile Contracts:

Compile the smart contracts using the following command:

```
forge build
```

## Executing Tests:

Run tests using the following command:
 
```
forge test
```

To test specific contracts or scenarios, use: 

```
forge test --match-contract <ContractName>
forge test --match-test <testName>
```
## Deploying Contracts:

Simulate deployment with the following command:

```
source .env
forge script script/ForwarderPool.s.sol:ForwarderPoolScript --rpc-url $MUMBAI_RPC -vvv
```

Deploying the contract and verifying the deployment

```
source .env
forge script script/ForwarderPool.s.sol:ForwarderPoolScript --rpc-url $MUMBAI_RPC -vvv --broadcast --verify
```

If you encounter an error related to EIP1559 fees, add the `--legacy` flag:

```
Failed to estimate EIP1559 fees. This chain might not support EIP1559, try adding --legacy to your command.
```

Something like this 👆

```
source .env
forge script script/ForwarderPool.s.sol:ForwarderPoolScript --rpc-url $MUMBAI_RPC -vvv --broadcast --verify --legacy
```

# IMPORTANT: Setting up Router Chain Fee payer and Refund Claimer

Fee Payer: Pays your fees on router chain 
Refund Claimer: Address who can claim your refund.

### Setting Up Fee payer

The fee payer is responsible for paying your fees on the router chain. Send a request for fee payer by calling  `setDappMetadata(feePayer)` and passing the fee payer's router chain address:

```
        setDappMetadata(feePayer);
```

The feePayer will have to approve this request from router chain. [Router Chain Explorer](https://routerscan.io/feePayer)

### Setting Up Refund Claimer

The refund claimer is the address that can claim your refund. Set up the refund claimer by calling `setForwarderRefundClaimer(uint64(expiryTime), refundClaimer, gasBytes)` and passing the refund claimer's router chain address, a expiry and gas in bytes:

```
        setForwarderRefundClaimer(uint64(expiryTime), refundClaimer, zeroBytes);
```


