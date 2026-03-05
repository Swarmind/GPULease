// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// Интерфейс ERC-20
interface IERC20 {
    function transfer(address to, uint256 amount) external returns (bool);
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
}

contract GPULease {
    struct Lease {
        address user;
        address provider;
        uint startTime;
        uint duration;
        uint pricePerSecond;
        uint totalAmount;
        bool active;
        bool completed;
    }

    IERC20 public token;
    mapping(uint => Lease) public leases;
    uint public leaseCount = 0;
    
    // Platform fee parameters
    address public deployer;
    uint public platformFeePercentage = 5; // 5% platform fee
    
    event LeaseStarted(uint leaseId, address user, address provider, uint duration, uint amount);
    event LeaseCompleted(uint leaseId, uint refund, uint providerReward);
    event LeaseCancelled(uint leaseId, uint refund);
    event PaymentReceived(uint leaseId, uint amount);
    event PlatformFeeCollected(uint leaseId, uint feeAmount);
    
    modifier onlyOwner(uint _leaseId) {
        require(leases[_leaseId].user == msg.sender || leases[_leaseId].provider == msg.sender, "Only user or provider can call this");
        _;
    }
    
    modifier onlyDeployer() {
        require(msg.sender == deployer, "Only deployer can call this");
        _;
    }
    
    constructor(address _token) {
        token = IERC20(_token);
        deployer = msg.sender; // Set deployer as the contract owner
    }
    
    function setPlatformFee(uint _feePercentage) external onlyDeployer {
        platformFeePercentage = _feePercentage;
    }
    
    function startLease(
        uint _duration,
        uint _pricePerSecond,
        address _provider
    ) external returns (uint leaseId) {
        require(_duration > 0, "Duration must be > 0");
        require(_pricePerSecond > 0, "Price per second must be > 0");
        
        uint totalAmount = _duration * _pricePerSecond;
        require(token.balanceOf(msg.sender) >= totalAmount, "Insufficient token balance");
        
        // Calculate platform fee
        uint platformFee = (totalAmount * platformFeePercentage) / 100;
        
        // Transfer tokens to contract (including platform fee)
        require(token.transferFrom(msg.sender, address(this), totalAmount), "Transfer failed");
        
        // Transfer platform fee to deployer
        require(token.transfer(deployer, platformFee), "Platform fee transfer failed");
        
        leaseId = leaseCount;
        leaseCount++;
        
        leases[leaseId] = Lease({
            user: msg.sender,
            provider: _provider,
            startTime: block.timestamp,
            duration: _duration,
            pricePerSecond: _pricePerSecond,
            totalAmount: totalAmount,
            active: true,
            completed: false
        });
        
        emit LeaseStarted(leaseId, msg.sender, _provider, _duration, totalAmount);
        emit PaymentReceived(leaseId, totalAmount);
        emit PlatformFeeCollected(leaseId, platformFee);
        
        return leaseId;
    }
    
    function completeLease(uint _leaseId) external onlyOwner(_leaseId) {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        uint actualDuration = block.timestamp - lease.startTime;
        uint actualCost = actualDuration * lease.pricePerSecond;
        uint refund = lease.totalAmount - actualCost;
        
        if (refund > 0) {
            require(token.transfer(lease.user, refund), "Refund transfer failed");
        }
        require(token.transfer(lease.provider, actualCost), "Provider transfer failed");
        
        lease.completed = true;
        lease.active = false;
        
        emit LeaseCompleted(_leaseId, refund, actualCost);
    }
    
    function cancelLease(uint _leaseId) external onlyOwner(_leaseId) {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        require(block.timestamp < lease.startTime + 300, "Cannot cancel after 5 minutes");
        
        uint refund = lease.totalAmount;
        require(token.transfer(lease.user, refund), "Refund transfer failed");
        
        lease.active = false;
        lease.completed = true;
        
        emit LeaseCancelled(_leaseId, refund);
    }
    
    function getLeaseStatus(uint _leaseId) external view returns (bool active, bool completed, uint startTime, uint duration) {
        Lease storage lease = leases[_leaseId];
        return (lease.active, lease.completed, lease.startTime, lease.duration);
    }
    
    function getContractBalance() external view returns (uint) {
        return token.balanceOf(address(this));
    }
}