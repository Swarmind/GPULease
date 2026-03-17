# GPULease - Ethereum GPU Leasing Platform

A smart contract for leasing GPU computing resources with flexible pricing and pause/resume functionality.

## Overview

GPULease is a decentralized smart contract that enables users to lease GPU compute and storage resources from providers. The system handles payment processing, platform fee collection, and provides mechanisms for pausing and resuming leases.

## Features

- **Flexible Pricing**: Separate pricing for storage and computation per second
- **Platform Fee System**: 5% platform fee (configurable)
- **Pause/Resume Functionality**: Leases can be paused during execution with accurate cost calculation
- **Lease Cancellation**: Within 5 minutes of creation
- **Refund Mechanism**: Automatic refunds for unused time
- **Access Control**: Role-based permissions using OpenZeppelin AccessControl

## Contract Structure

### Core Data Structures

**Lease Struct**
```solidity
struct Lease {
    address user;                    // The lease requester
    address provider;                // The GPU provider
    uint startTime;                  // When the lease started
    uint duration;                   // Total lease duration in seconds
    uint storagePricePerSecond;      // Price per second for storage
    uint computePricePerSecond;      // Price per second for computation
    uint totalAmount;                // Total amount to be paid
    bool active;                     // Whether the lease is currently active
    bool completed;                  // Whether the lease has been completed
    bool paused;                     // Whether the lease is currently paused
    uint lastPausedTime;             // Time when lease was last paused
    uint pausedDuration;             // Cumulative duration of pauses in seconds
}
```

### Key Functions

**Deposit & Withdraw**
- `deposit(amount)`: Add tokens to your account balance
- `withdraw(amount)`: Remove tokens from your account balance

**Lease Management**
- `startLease(duration, storagePricePerSecond, computePricePerSecond, provider)`: Start a lease for yourself
- `startLeaseWithUser(duration, storagePricePerSecond, computePricePerSecond, provider, user)`: Admin can start a lease on behalf of another user
- `pauseLease(leaseId)`: Pause an active lease (only the user or provider can call)
- `resumeLease(leaseId)`: Resume a paused lease (only the user or provider can call)
- `completeLease(leaseId)`: Complete an active lease and settle payments
- `cancelLease(leaseId)`: Cancel a lease within 5 minutes

**Utility Functions**
- `getContractBalance()`: Check contract's token balance
- `getLeaseStatus(leaseId)`: Get detailed information about a specific lease

## Security Features

- **ReentrancyGuard**: Prevents reentrant calls during critical operations
- **Access Control**: Role-based permissions for administrative functions
- **Platform Fee Management**: Admin can update platform fee percentage

## Events

The contract emits the following events:
- `LeaseStarted`: When a lease is created
- `LeaseCompleted`: When a lease is completed successfully 
- `LeaseCancelled`: When a lease is cancelled
- `PaymentReceived`: When payment is received for a lease
- `PlatformFeeCollected`: When platform fee is collected
- `UserDeposited`: When a user deposits tokens
- `UserWithdrawn`: When a user withdraws tokens
- `LeasePaused`: When a lease is paused
- `LeaseResumed`: When a lease is resumed

## Usage Flow

1. **Deposit Tokens**: Users must deposit tokens to their account balance before starting leases
2. **Start Lease**: Choose duration and pricing, then start the lease  
3. **Pause/Resume** (Optional): During execution, pause or resume the lease
4. **Complete/Cancellation**: Either complete the lease or cancel it within 5 minutes

## Contract Parameters

- `platformFeePercentage`: Default 5% platform fee (modifiable by admin)
- All prices are expressed in smallest token units (wei)

## Access Control

The contract uses OpenZeppelin's AccessControl with:
- `DEFAULT_ADMIN_ROLE`: Can set platform fees and perform administrative functions
- Lease participants: Users and providers can manage their own leases