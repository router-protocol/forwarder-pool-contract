# Documentation for ForwarderPool Contract and Test Cases
### Overview
The ForwarderPool contract is a Solidity smart contract designed for Ethereum-based blockchain networks. It facilitates the forwarding of both ERC20 tokens and native blockchain tokens (like Ethereum) to other addresses. This contract also integrates whitelisting functionality to restrict certain actions to authorized addresses only.

### Contract Features
ERC20 and Native Token Management: Handles deposits of ERC20 and native tokens, tracking the balances internally.
- Asset Forwarding: Uses an IAssetForwarder interface to forward assets.
- Whitelisting: Provides the ability to whitelist addresses, allowing only authorized addresses to perform certain actions.
- Owner Privileges: Includes functions that can only be executed by the contract owner, like withdrawing tokens and managing whitelisted addresses.
Key Functions
- setAssetForwarder(address _assetForwarder): Sets the address of the asset forwarder. Restricted to the contract owner.

- setWhitelistedFiller(address filler, bool whitelisted): Adds or removes an address from the whitelist. Restricted to the contract owner.

- depositERC20(address token, uint256 amount): Allows users to deposit ERC20 tokens into the contract.

- depositNativeToken(): Allows users to deposit native blockchain tokens (e.g., ETH) into the contract.

- iRelay(IAssetForwarder.RelayData memory relayData): Executes the relay of tokens using the asset forwarder, subject to whitelist and balance checks.

- iRelayMessage(IAssetForwarder.RelayDataMessage memory relayMessageData): Similar to iRelay but for relaying messages.

- withdrawNativeToken(address payable recipient, uint256 amount): Allows the owner to withdraw native tokens from the contract.

- withdrawERC20(address token, address recipient, uint256 amount): Allows the owner to withdraw ERC20 tokens from the contract.

### Test Cases Overview
The test cases use the Foundry framework, a smart contract development tool for Ethereum. They are designed to validate the functionality of the ForwarderPool contract in various scenarios.

Key Test Functions
- testSetAssetForwarder(): Tests setting the asset forwarder address.

- testSetWhitelistedFiller(): Tests adding an address to the whitelist.

- testDepositERC20(): Tests depositing ERC20 tokens into the forwarder pool.

- testDepositNativeToken(): Tests depositing native tokens into the forwarder pool.

- testIRelayWithErc20(): Tests the execution of a relay with ERC20 tokens.

- testIRelayWithNativeToken(): Tests the execution of a relay with native tokens.

- testIRelayMessageWithErc20(): Tests the execution of a relay message with ERC20 tokens.

- testIRelayMessageWithNativeToken(): Tests the execution of a relay message with native tokens.

### Test Setup
Uses Ethereum forks for testing (mumbaiFork).
Involves mock contracts and users (Interactor, Dummy, DamnValuableToken, etc.).
Verifies contract behavior under various network conditions and states.
Conclusion
The ForwarderPool contract and its tests provide a robust framework for managing and forwarding assets in a decentralized environment, with an emphasis on security through owner privileges and whitelisting. The test cases ensure that the contract functions as expected under different scenarios, contributing to its reliability and trustworthiness.

## Installation:

Install dependencies using 
```
forge install
```

## Compile Contracts:

Compile the smart contracts.
```
forge build
```
## Executing Tests:

Run tests using the command 
```
forge test
```
To test specific contracts or scenarios, use: 
```
forge test --match-contract <ContractName>
forge test --match-test <testName>
```