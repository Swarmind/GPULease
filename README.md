# GPULease - Ethereum GPU Leasing Platform

This project implements a decentralized GPU leasing platform on Ethereum using Solidity smart contracts. The platform allows users to lease GPU resources from providers in a trustless manner.

## Project Overview

The core of the project is the `GPULease.sol` contract, which enables:

- **Lease Creation**: Users can start GPU leases by specifying duration and price per second
- **Secure Payments**: Funds are held in escrow and only released to providers upon lease completion
- **Automatic Refunds**: Users receive refunds for unused time after lease completion
- **Cancellation Policy**: Leases can be cancelled within 5 minutes with full refund
- **Real-time Status Tracking**: Users can check lease status and contract balance

The project uses Hardhat 3 Beta with TypeScript, mocha for testing, and ethers.js for Ethereum interactions.

## Smart Contract Features

The `GPULease` contract includes:

- **Lease Management**: Start, complete, and cancel leases
- **Payment Processing**: Handle ERC-20 token transfers for payments
- **Access Control**: Only lease owners or providers can perform actions
- **Event Logging**: Emit events for lease creation, completion, and cancellation
- **Balance Tracking**: Monitor contract balance and lease status

## Usage

### Running Tests

To run all tests in the project:

```shell
npx hardhat test
```

You can also selectively run Solidity or TypeScript tests:

```shell
npx hardhat test solidity
npx hardhat test mocha
```

### Deployment to Sepolia

This project includes an Ignition module to deploy the contract. You can deploy to a local chain or to Sepolia.

To run the deployment to a local chain:

```shell
npx hardhat ignition deploy ignition/modules/Counter.ts
```

To run the deployment to Sepolia, you need an account with funds. The provided Hardhat configuration includes a `SEPOLIA_PRIVATE_KEY` variable.

Set the `SEPOLIA_PRIVATE_KEY` variable using the `hardhat-keystore` plugin:

```shell
npx hardhat keystore set SEPOLIA_PRIVATE_KEY
```

After setting the variable, run the deployment:

```shell
npx hardhat ignition deploy --network sepolia ignition/modules/Counter.ts
```

## Smart Contract Details

The `GPULease` contract manages GPU leasing operations with the following structure:

- **Lease Structure**: Tracks user, provider, duration, price, and status
- **Token Integration**: Uses ERC-20 tokens for payments
- **Time-Based Calculations**: Calculates actual costs based on lease duration
- **Refund System**: Provides automatic refunds for unused time
- **Security Measures**: Includes access controls and time-based cancellation limits

## Contract Function Usage

Here's how to use the functions in the GPULease contract:

### 1. Starting a Lease

To start a lease, call the `startLease` function with the following parameters:

```solidity
function startLease(
    uint _duration,
    uint _pricePerSecond,
    address _provider
) external returns (uint leaseId)
```

**Parameters:**
- `_duration`: The duration of the lease in seconds
- `_pricePerSecond`: The price per second in wei
- `_provider`: The address of the GPU provider

**Returns:**
- `leaseId`: The ID of the newly created lease

**Example:**
```javascript
// Assuming you have already deployed the contract and have the contract instance
const leaseId = await contract.startLease(3600, 1000000000, providerAddress);
```

### 2. Completing a Lease

To complete a lease, call the `completeLease` function with the lease ID:

```solidity
function completeLease(uint _leaseId) external onlyOwner(_leaseId)
```

**Parameters:**
- `_leaseId`: The ID of the lease to complete

**Requirements:**
- The caller must be the lease owner or provider
- The lease must be active
- The lease must not be already completed

**Example:**
```javascript
// Complete a lease with ID 0
await contract.completeLease(0);
```

### 3. Canceling a Lease

To cancel a lease, call the `cancelLease` function with the lease ID:

```solidity
function cancelLease(uint _leaseId) external onlyOwner(_leaseId)
```

**Parameters:**
- `_leaseId`: The ID of the lease to cancel

**Requirements:**
- The caller must be the lease owner or provider
- The lease must be active
- The lease must not be already completed
- The cancellation must happen within 5 minutes of lease start

**Example:**
```javascript
// Cancel a lease with ID 0
await contract.cancelLease(0);
```

### 4. Getting Lease Status

To get the status of a lease, call the `getLeaseStatus` function with the lease ID:

```solidity
function getLeaseStatus(uint _leaseId) external view returns (bool active, bool completed, uint startTime, uint duration)
```

**Parameters:**
- `_leaseId`: The ID of the lease to check

**Returns:**
- `active`: Whether the lease is active
- `completed`: Whether the lease is completed
- `startTime`: The start time of the lease
- `duration`: The duration of the lease

**Example:**
```javascript
const [active, completed, startTime, duration] = await contract.getLeaseStatus(0);
```

### 5. Getting Contract Balance

To get the current balance of the contract, call the `getContractBalance` function:

```solidity
function getContractBalance() external view returns (uint)
```

**Returns:**
- The current balance of the contract in wei

**Example:**
```javascript
const balance = await contract.getContractBalance();
```

This project demonstrates a practical implementation of decentralized GPU resource allocation on Ethereum.