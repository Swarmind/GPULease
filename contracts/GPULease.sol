// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract GPULease is AccessControl {
    
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

    // User balances mapping
    mapping(address => uint256) public userBalances;
    
    // Locked funds for leases mapping  
    mapping(uint => uint256) public lockedFunds;
    
    IERC20 public token;
    mapping(uint => Lease) public leases;
    uint public leaseCount = 0;
    
    // Platform fee parameters
    uint public platformFeePercentage = 5; // 5% platform fee
    
    event LeaseStarted(uint leaseId, address user, address provider, uint duration, uint amount);
    event LeaseCompleted(uint leaseId, uint refund, uint providerReward);
    event LeaseCancelled(uint leaseId, uint refund);
    event PaymentReceived(uint leaseId, uint amount);
    event PlatformFeeCollected(uint leaseId, uint feeAmount);
    event UserDeposited(address user, uint amount);
    event UserWithdrawn(address user, uint amount);
    
    address deployer;


    // operation with leaseID shoud be available only by deal participant or the admin of the contract (in case of delegated call)
    modifier onlyLeaseOwner(uint _leaseId) {
        require(leases[_leaseId].user == msg.sender || leases[_leaseId].provider == msg.sender || msg.sender == deployer, "Only user or provider can call this");
        _;
    }
    
    constructor(address _token) {
        token = IERC20(_token);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender); // Grant default admin role to deployer
        deployer = msg.sender;
    }
    
    function setPlatformFee(uint _feePercentage) external onlyRole(DEFAULT_ADMIN_ROLE) {
        platformFeePercentage = _feePercentage;
    }
    
    /**
     * @dev Deposit tokens to user balance
     */
    function deposit(uint256 amount) external {
        require(amount > 0, "Deposit amount must be > 0");
        token.transferFrom(msg.sender, address(this), amount);
        userBalances[msg.sender] = userBalances[msg.sender] + amount;
        emit UserDeposited(msg.sender, amount);
    }
    
    /**
     * @dev Withdraw tokens from user balance
     */
    function withdraw(uint256 amount) external {
        require(userBalances[msg.sender] >= amount, "Insufficient balance");
        userBalances[msg.sender] = userBalances[msg.sender] - amount;
        token.transfer(msg.sender, amount);
        emit UserWithdrawn(msg.sender, amount);
    }
    
    function startLeaseWithUser(
        uint _duration,
        uint _pricePerSecond,
        address _provider,
        address _user
    ) external returns (uint leaseId) {
        require(_duration > 0, "Duration must be > 0");
        require(_pricePerSecond > 0, "Price per second must be > 0");
        
        uint totalAmount = _duration * _pricePerSecond;
        require(userBalances[_user] >= totalAmount, "Insufficient token balance");
        
        // Calculate platform fee
        uint platformFee = (totalAmount * platformFeePercentage) / 100;
        
        // Deduct funds from user balance and lock them in lockedFunds mapping by leaseId
        userBalances[_user] = userBalances[_user] - totalAmount;
        lockedFunds[leaseCount] = totalAmount;
        
        leaseId = leaseCount;
        leaseCount++;
        
        leases[leaseId] = Lease({
            user: _user,
            provider: _provider,
            startTime: block.timestamp,
            duration: _duration,
            pricePerSecond: _pricePerSecond,
            totalAmount: totalAmount,
            active: true,
            completed: false
        });
        
        emit LeaseStarted(leaseId, _user, _provider, _duration, totalAmount);
        emit PaymentReceived(leaseId, totalAmount);
        emit PlatformFeeCollected(leaseId, platformFee);
        
        return leaseId;
    }
    
    function startLease(
        uint _duration,
        uint _pricePerSecond,
        address _provider
    ) external returns (uint leaseId) {
        uint lid = this.startLeaseWithUser(_duration, _pricePerSecond, _provider, msg.sender);
        return lid;
    }
    
    function completeLease(uint _leaseId) external onlyLeaseOwner(_leaseId) {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        uint actualDuration = block.timestamp - lease.startTime;
        uint actualCost = actualDuration * lease.pricePerSecond;
        uint refund = lease.totalAmount - actualCost;
        
        // Refund unused amount back to user's balance
        if (refund > 0) {
            userBalances[lease.user] = userBalances[lease.user] + refund;
        }
        // Transfer actual cost to provider
        require(token.transfer(lease.provider, actualCost), "Provider transfer failed");
        
        // Unlock the funds from lockedFunds mapping 
        delete lockedFunds[_leaseId];
        
        lease.completed = true;
        lease.active = false;
        
        emit LeaseCompleted(_leaseId, refund, actualCost);
    }
    
    function cancelLease(uint _leaseId) external onlyLeaseOwner(_leaseId) {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        require(block.timestamp < lease.startTime + 300, "Cannot cancel after 5 minutes");
        
        uint refund = lease.totalAmount;
        
        // Return funds back to user's balance
        userBalances[lease.user] = userBalances[lease.user] + refund;
        
        // Unlock the funds from lockedFunds mapping 
        delete lockedFunds[_leaseId];
        
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