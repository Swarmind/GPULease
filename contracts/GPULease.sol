// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract GPULease is Ownable {
    using SafeERC20 for IERC20;
    IERC20 public credit; 
    address public treasury;

     struct Lease {
        address user;
        address provider;
        uint startTime;
        uint storagePricePerSecond; // Price per second for storage
        uint computePricePerSecond; // Price per second for computation
        bool active;
        bool completed;
        bool paused; // Lease can be paused during execution 
        uint pausedAt; // Time when lease was paused
        uint pausedDuration; // Cumulative duration of pauses in seconds
    }

    struct FrozenFundsInfo {
        uint leaseId;
        uint256 amount;
    }

    mapping(address => uint[]) public userActiveLeases;
    mapping(address => uint256) public balances;
    mapping(uint => uint256) public frozenFunds;
    mapping(uint => Lease) public leases;
    uint public leaseCount = 0;

    uint public platformFeePercentage = 5; // 5% platform fee

    event Deposit(address indexed user, uint256 amount);
    event Withdraw(address indexed user, uint256 amount);
    
    event LeaseStarted(uint leaseId, address user, address provider, uint duration, uint amount);
    event LeaseCompleted(uint leaseId);
    event LeasePaused(uint leaseId);
    event LeaseResumed(uint leaseId);

 

    constructor(address credit_, address treasury_) Ownable(msg.sender) {
        credit = IERC20(credit_);
        treasury = treasury_;
    }


    //admin stuff
    function setPlatformFee(uint _feePercentage) public onlyOwner {
        require(_feePercentage <= 100, "Fee too high");
        platformFeePercentage = _feePercentage;
    }
    function setTreasury(address newTreasury) external onlyOwner {
        require(newTreasury != address(0), "zero treasury");
        uint amount = userBalance(treasury);
        balances[newTreasury] += amount;
        balances[treasury] = 0;    
        treasury = newTreasury;
    }


   //wallet stuff
    function deposit(uint256 amount) external {
        credit.safeTransferFrom(msg.sender, address(this), amount);
        balances[msg.sender] += amount;
        emit Deposit(msg.sender, amount);
    }

    function withdraw(uint256 amount) external {
        require(balances[msg.sender] >= amount, "insufficient balance");
        credit.safeTransfer(msg.sender, amount);
        balances[msg.sender] -= amount;
        emit Withdraw(msg.sender, amount);
    }
   
    function userBalance(address user) public view returns (uint256) {
    return balances[user];
    }


    //lease stuff
    function startLease(
        uint _duration,
        uint _storagePricePerSecond,
        uint _computePricePerSecond,
        address _provider,
        address _user
    ) public onlyOwner returns (uint leaseId) {
        require(_duration > 0, "Duration must be > 0");
        require(_storagePricePerSecond > 0 || _computePricePerSecond > 0, "At least one price must be > 0");
        require(_user != address(0), "invalid user");
        require(_provider != address(0), "invalid provider");
        
        // Calculate total amounts for both storage and compute
        uint totalStorageAmount = _duration * _storagePricePerSecond;
        uint totalComputeAmount = _duration * _computePricePerSecond;
        uint totalAmount = totalStorageAmount + totalComputeAmount;
        uint platformFee = (totalAmount * platformFeePercentage) / 100;
        totalAmount += platformFee; 
       
        require(balances[_user] >= totalAmount, "Insufficient token balance");
                
        // Deduct funds from user balance and lock them in lockedFunds mapping by leaseId
        balances[_user] = balances[_user] - totalAmount;
        frozenFunds[leaseCount] = totalAmount;
        
        leaseId = leaseCount;
        leaseCount++;
        
        leases[leaseId] = Lease({
            user: _user,
            provider: _provider,
            startTime: block.timestamp - 5 minutes, //so we won't need any cancel function
            storagePricePerSecond: _storagePricePerSecond,
            computePricePerSecond: _computePricePerSecond,
            active: true,
            completed: false,
            paused: false,
            pausedAt: 0,
            pausedDuration: 0
        });
        userActiveLeases[_user].push(leaseId);
        emit LeaseStarted(leaseId, _user, _provider, _duration, totalAmount);
        return leaseId;
    }

     function pauseLease(uint _leaseId) public onlyOwner {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        require(!lease.paused, "Lease is already paused");
        
        // Set the pause time
        lease.pausedAt = block.timestamp;
        lease.paused = true;
        
        emit LeasePaused(_leaseId);
    }
    
    function resumeLease(uint _leaseId) external onlyOwner {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        require(lease.paused, "Lease is not paused");
        
        uint pauseDuration = block.timestamp - lease.pausedAt;
        lease.pausedDuration += pauseDuration;
        lease.pausedAt = 0; // Reset last paused time
        lease.paused = false;
        
        emit LeaseResumed(_leaseId);
    }

    function completeLease(uint _leaseId) external onlyOwner  {
        Lease storage lease = leases[_leaseId];
        require(lease.active, "Lease is not active");
        require(!lease.completed, "Lease already completed");
        
        uint actualStorageCost;
        uint actualComputeCost;
        (actualStorageCost, actualComputeCost) = calculateActualCost(_leaseId);
        
        // Total cost based on the effective duration
        uint actualTotalCost = actualStorageCost + actualComputeCost; 
        
        
        // Calculate platform fee from the total actual cost
        uint platformFee = (actualTotalCost * platformFeePercentage) / 100;
        balances[treasury] += platformFee;
        frozenFunds[_leaseId] -= platformFee;

        uint providerAmount = actualTotalCost - platformFee;
        balances[lease.provider] += providerAmount;
        frozenFunds[_leaseId] -= providerAmount;

        balances[lease.user] += frozenFunds[_leaseId];

        delete frozenFunds[_leaseId];
        
        lease.completed = true;
        lease.active = false;

            address user = leases[_leaseId].user;
    uint[] storage leasesList = userActiveLeases[user];
    for (uint i = 0; i < leasesList.length; i++) {
        if (leasesList[i] == _leaseId) {
            leasesList[i] = leasesList[leasesList.length - 1];
            leasesList.pop();
            break;
        }
    }
        emit LeaseCompleted(_leaseId);
    }


    function calculateActualCost(uint _leaseId) 
    internal 
    view 
    returns (uint actualStorageCost, uint actualComputeCost) 
{
    Lease storage lease = leases[_leaseId];

    require(lease.startTime > 0, "Lease not started");

    uint duration = block.timestamp - lease.startTime;

    uint totalPaused = lease.pausedDuration;
    if (lease.paused) {
        totalPaused += (block.timestamp - lease.pausedAt);
    }

    actualStorageCost = duration * lease.storagePricePerSecond;

    uint activeDuration = duration > totalPaused ? duration - totalPaused : 0;
    actualComputeCost = activeDuration * lease.computePricePerSecond;

    return (actualStorageCost, actualComputeCost);
}
    
   function getUserFrozenFunds(address user) 
    external
    view
    returns (FrozenFundsInfo[] memory result) 
{
    uint[] storage leasesList = userActiveLeases[user];
    result = new FrozenFundsInfo[](leasesList.length);

    for (uint i = 0; i < leasesList.length; i++) {
        uint leaseId = leasesList[i];

        result[i] = FrozenFundsInfo({
            leaseId: leaseId,
            amount: frozenFunds[leaseId]
        });
    }

    return result;
}

}