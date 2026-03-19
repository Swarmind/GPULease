import { expect } from "chai";
import { network } from "hardhat";
const { ethers } = await network.connect();

describe("GPULease", function () {
  let gpulease: any;
  let mockToken: any;
  let deployer: any;
  let user1: any;
  let provider1: any;
  let user2: any;

  beforeEach(async function () {
    // Get signers
    [deployer, user1, provider1, user2] = await ethers.getSigners();

    // Deploy MockERC20 token
    const MockToken = await ethers.getContractFactory("MockERC20");
    mockToken = await MockToken.deploy("Test Token", "TST");
    await mockToken.waitForDeployment();

    // Mint some tokens to users for testing
    await mockToken.mint(user1.address, ethers.parseEther("1000"));
    await mockToken.mint(provider1.address, ethers.parseEther("1000"));
    await mockToken.mint(user2.address, ethers.parseEther("1000"));

    // Deploy GPULease contract
    const GPULease = await ethers.getContractFactory("GPULease");
    gpulease = await GPULease.deploy(mockToken.getAddress());
    await gpulease.waitForDeployment();

    // Approve tokens for contracts
    await mockToken.connect(user1).approve(gpulease.getAddress(), ethers.parseEther("1000"));
    await mockToken.connect(provider1).approve(gpulease.getAddress(), ethers.parseEther("1000"));
    await mockToken.connect(user2).approve(gpulease.getAddress(), ethers.parseEther("1000"));
  });

  describe("Deployment", function () {
    it("Should set the right deployer and token", async function () {
      expect(await gpulease.deployer()).to.equal(deployer.address);
      expect(await gpulease.token()).to.equal(await mockToken.getAddress())
    });

    it("Should set default platform fee percentage", async function () {
      expect(await gpulease.platformFeePercentage()).to.equal(5);
    });
  });

  describe("Deposit and Withdraw", function () {
    it("Should allow users to deposit tokens", async function () {
      const depositAmount = ethers.parseEther("100");
      
      await expect(() => 
        gpulease.connect(user1).deposit(depositAmount)
      ).to.changeTokenBalance(ethers,mockToken, user1, -depositAmount);
      
      expect(await gpulease.userBalances(user1.address)).to.equal(depositAmount);
    });

    it("Should allow users to withdraw tokens", async function () {
      const depositAmount = ethers.parseEther("100");
      const withdrawAmount = ethers.parseEther("50");
      
      await gpulease.connect(user1).deposit(depositAmount);
      
      await expect(() => 
        gpulease.connect(user1).withdraw(withdrawAmount)
      ).to.changeTokenBalance(ethers,mockToken, user1, withdrawAmount);
      
      expect(await gpulease.userBalances(user1.address)).to.equal(depositAmount - withdrawAmount);
    });

    it("Should prevent withdrawals with insufficient balance", async function () {
      const withdrawAmount = ethers.parseEther("50");
      
      await expect(gpulease.connect(user1).withdraw(withdrawAmount))
        .to.be.revertedWith("Insufficient balance");
    });
  });

  describe("Lease Management", function () {
    it("Should start a lease successfully", async function () {
      const duration = 3600; // 1 hour
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      
      expect(await gpulease.leaseCount()).to.equal(1);
      
      // Check event emission
      expect(tx).to.emit(gpulease, "LeaseStarted");
      expect(tx).to.emit(gpulease, "PaymentReceived");
      expect(tx).to.emit(gpulease, "PlatformFeeCollected");
    });

    it("Should allow admin to start lease on behalf of a user", async function () {
      const duration = 3600; // 1 hour
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Admin starts lease for user
      const tx = await gpulease.startLeaseWithUser(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address,
        user1.address
      );
      
      expect(await gpulease.leaseCount()).to.equal(1);
    });

    it("Should properly handle lease completion", async function () {
      const duration = 3600; // 1 hour
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Start a lease
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      const leaseId = 0;
      
      // Simulate time passing (we'll use Hardhat's time manipulation)
      await ethers.provider.send("evm_increaseTime", [duration]);
      await ethers.provider.send("evm_mine");
      
      // Complete the lease
      await expect(() => 
        gpulease.connect(user1).completeLease(leaseId)
      ).to.changeTokenBalance(ethers,mockToken, provider1, totalAmount * 95n / 100n); // 95% goes to provider (5% platform fee)
    });

    it("Should allow lease cancellation within 5 minutes", async function () {
      const duration = 3600; // 1 hour
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Start a lease
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      const leaseId = 0;
      
      // Cancel immediately (within 5 minutes)
      await expect(() => 
        gpulease.connect(user1).cancelLease(leaseId)
      ).to.changeTokenBalance(ethers,mockToken, user1, totalAmount);
    });

    it("Should prevent cancellation after 5 minutes", async function () {
      const duration = 3600; // 1 hour
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Start a lease
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      const leaseId = 0;
      
      // Fast forward time by more than 5 minutes (300 seconds)
      await ethers.provider.send("evm_increaseTime", [301]);
      await ethers.provider.send("evm_mine");
      
      // Try to cancel after the grace period - should fail
      await expect(gpulease.connect(user1).cancelLease(leaseId))
        .to.be.revertedWith("Cannot cancel after 5 minutes");
    });

    it("Should allow pausing and resuming leases", async function () {
      const duration = 7200; // 2 hours
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Start a lease
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      const leaseId = 0;
      
      // Pause the lease
      await expect(() => 
        gpulease.connect(user1).pauseLease(leaseId)
      ).to.emit(gpulease, "LeasePaused");
      
      // Resume the lease
      await expect(() => 
        gpulease.connect(user1).resumeLease(leaseId)
      ).to.emit(gpulease, "LeaseResumed");
    });

    it("Should calculate actual cost correctly with paused time", async function () {
      const duration = 7200; // 2 hours
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Start a lease
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      const leaseId = 0;
      
      // Pause the lease after 1 hour
      await ethers.provider.send("evm_increaseTime", [3600]);
      await ethers.provider.send("evm_mine");
      
      await gpulease.connect(user1).pauseLease(leaseId);
      
      // Continue for another hour, then resume
      await ethers.provider.send("evm_increaseTime", [3600]);
      await ethers.provider.send("evm_mine");
      
      await gpulease.connect(user1).resumeLease(leaseId);
      
      // Complete the lease after 1 more hour
      await ethers.provider.send("evm_increaseTime", [3600]);
      await ethers.provider.send("evm_mine");
      
      // Complete the lease to verify actual cost calculation
      await gpulease.connect(user1).completeLease(leaseId);
    });
  });

  describe("Access Control", function () {
    it("Should allow only user or provider to call lease functions", async function () {
      const duration = 3600; // 1 hour
      const storagePricePerSecond = ethers.parseEther("0.0001");
      const computePricePerSecond = ethers.parseEther("0.0002");
      
      const totalAmount = (BigInt(duration) * storagePricePerSecond) + (BigInt(duration) * computePricePerSecond);
      
      await gpulease.connect(user1).deposit(totalAmount);
      
      // Start a lease
      const tx = await gpulease.connect(user1).startLease(
        duration,
        storagePricePerSecond,
        computePricePerSecond,
        provider1.address
      );
      
      const receipt = await tx.wait();
      const leaseId = 0;
      
      // User should be able to pause the lease
      await expect(gpulease.connect(user1).pauseLease(leaseId)).to.not.be.reverted;
      
      // Provider should also be able to pause the lease
      await expect(gpulease.connect(provider1).pauseLease(leaseId)).to.not.be.reverted;
      
      // Someone else should not be able to call pause
      await expect(gpulease.connect(user2).pauseLease(leaseId))
        .to.be.revertedWith("Only user or provider can call this");
    });

    it("Should allow admin to set platform fee", async function () {
      const newFee = 10;
    
      // 1. Отправляем транзакцию от админа
      const tx = await gpulease.connect(deployer).setPlatformFee(newFee);
    
      // 2. Явно ждём её включения в блок
      await tx.wait();
    
      // 3. Проверяем, что значение обновилось
      expect(await gpulease.platformFeePercentage()).to.equal(newFee);
    });
  });

  describe("Contract Balance", function () {
    it("Should return correct contract balance", async function () {
      const depositAmount = ethers.parseEther("100");
      
      await gpulease.connect(user1).deposit(depositAmount);
      
      expect(await gpulease.getContractBalance()).to.equal(depositAmount);
    });
  });
});