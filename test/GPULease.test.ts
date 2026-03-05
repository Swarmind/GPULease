import { expect } from "chai";
import  ethers  from "hardhat";
import { Contract } from "ethers";

const  parseEther  = ethers;

const deployContract = async (): Promise<Contract> => {
  const [owner, user, provider] = await ethers.getSigners();
  
  // Deploy a mock ERC-20 token
  const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
  const mockToken = await ERC20Mock.deploy("Mock Token", "MTK", parseEther("10000"));
  await mockToken.waitForDeployment();
  
  // Deploy GPULease contract
  const GPULease = await ethers.getContractFactory("GPULease");
  const gpulease = await GPULease.deploy(mockToken.target);
  await gpulease.waitForDeployment();
  
  return gpulease;
};

describe("GPULease Contract", function () {
  let gpulease: Contract;
  let mockToken: Contract;
  let owner: any;
  let user: any;
  let provider: any;
  
  beforeEach(async () => {
    [owner, user, provider] = await ethers.getSigners();
    gpulease = await deployContract();
    mockToken = await ethers.getContractAt("ERC20Mock", await gpulease.token());
  });
  
  describe("Lease Creation", function () {
    it("should allow users to start a lease", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      const totalAmount = duration * pricePerSecond;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      const tx = await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Check events
      await expect(tx).to.emit(gpulease, "LeaseStarted").withArgs(
        0, // leaseId
        user.address,
        provider.address,
        duration,
        parseEther(totalAmount.toString())
      );
      
      // Check lease status
      const lease = await gpulease.leases(0);
      expect(lease.user).to.equal(user.address);
      expect(lease.provider).to.equal(provider.address);
      expect(lease.startTime).to.be.gt(0);
      expect(lease.duration).to.equal(duration);
      expect(lease.pricePerSecond).to.equal(pricePerSecond);
      expect(lease.totalAmount).to.equal(parseEther(totalAmount.toString()));
      expect(lease.active).to.be.true;
      expect(lease.completed).to.be.false;
    });
    
    it("should revert if duration is zero", async function () {
      await expect(
        gpulease.connect(user).startLease(0, 1, provider.address)
      ).to.be.revertedWith("Duration must be > 0");
    });
    
    it("should revert if price per second is zero", async function () {
      await expect(
        gpulease.connect(user).startLease(100, 0, provider.address)
      ).to.be.revertedWith("Price per second must be > 0");
    });
    
    it("should revert if user has insufficient balance", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Try to start lease without sufficient balance
      await expect(
        gpulease.connect(user).startLease(duration, pricePerSecond, provider.address)
      ).to.be.revertedWith("Insufficient token balance");
    });
  });
  
  describe("Lease Completion", function () {
    let leaseId: number;
    
    beforeEach(async function () {
      // Setup initial conditions
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      const tx = await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      leaseId = 0;
    });
    
    it("should allow lease owner to complete a lease", async function () {
      // Advance time to simulate lease completion
      await ethers.provider.send("evm_increaseTime", [100]);
      
      // Complete lease
      const tx = await gpulease.connect(user).completeLease(0);
      
      // Check events
      await expect(tx).to.emit(gpulease, "LeaseCompleted").withArgs(
        0, // leaseId
        parseEther("0"), // refund
        parseEther("100") // providerReward
      );
      
      // Check lease status
      const lease = await gpulease.leases(0);
      expect(lease.active).to.be.false;
      expect(lease.completed).to.be.true;
      
      // Check contract balance
      expect(await mockToken.balanceOf(gpulease.target)).to.equal(0);
    });
    
    it("should allow provider to complete a lease", async function () {
      // Advance time to simulate lease completion
      await ethers.provider.send("evm_increaseTime", [100]);
      
      // Complete lease
      const tx = await gpulease.connect(provider).completeLease(0);
      
      // Check events
      await expect(tx).to.emit(gpulease, "LeaseCompleted").withArgs(
        0, // leaseId
        parseEther("0"), // refund
        parseEther("100") // providerReward
      );
      
      // Check lease status
      const lease = await gpulease.leases(0);
      expect(lease.active).to.be.false;
      expect(lease.completed).to.be.true;
      
      // Check contract balance
      expect(await mockToken.balanceOf(gpulease.target)).to.equal(0);
    });
    
    it("should revert if lease is not active", async function () {
      // Try to complete a lease that doesn't exist
      await expect(
        gpulease.connect(user).completeLease(1)
      ).to.be.revertedWith("Lease is not active");
    });
    
    it("should revert if lease is already completed", async function () {
      // Advance time to simulate lease completion
      await ethers.provider.send("evm_increaseTime", [100]);
      
      // Complete lease
      await gpulease.connect(user).completeLease(0);
      
      // Try to complete again
      await expect(
        gpulease.connect(user).completeLease(0)
      ).to.be.revertedWith("Lease already completed");
    });
  });
  
  describe("Lease Cancellation", function () {
    let leaseId: number;
    
    beforeEach(async function () {
      // Setup initial conditions
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      const tx = await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      leaseId = 0;
    });
    
    it("should allow lease owner to cancel a lease within 5 minutes", async function () {
      // Advance time to simulate lease start
      await ethers.provider.send("evm_increaseTime", [50]); // 50 seconds
      
      // Cancel lease
      const tx = await gpulease.connect(user).cancelLease(0);
      
      // Check events
      await expect(tx).to.emit(gpulease, "LeaseCancelled").withArgs(
        0, // leaseId
        parseEther("100") // refund
      );
      
      // Check lease status
      const lease = await gpulease.leases(0);
      expect(lease.active).to.be.false;
      expect(lease.completed).to.be.true;
      
      // Check contract balance
      expect(await mockToken.balanceOf(gpulease.target)).to.equal(0);
    });
    
    it("should allow provider to cancel a lease within 5 minutes", async function () {
      // Advance time to simulate lease start
      await ethers.provider.send("evm_increaseTime", [50]); // 50 seconds
      
      // Cancel lease
      const tx = await gpulease.connect(provider).cancelLease(0);
      
      // Check events
      await expect(tx).to.emit(gpulease, "LeaseCancelled").withArgs(
        0, // leaseId
        parseEther("100") // refund
      );
      
      // Check lease status
      const lease = await gpulease.leases(0);
      expect(lease.active).to.be.false;
      expect(lease.completed).to.be.true;
      
      // Check contract balance
      expect(await mockToken.balanceOf(gpulease.target)).to.equal(0);
    });
    
    it("should revert if lease is not active", async function () {
      // Try to cancel a lease that doesn't exist
      await expect(
        gpulease.connect(user).cancelLease(1)
      ).to.be.revertedWith("Lease is not active");
    });
    
    it("should revert if lease is already completed", async function () {
      // Advance time to simulate lease completion
      await ethers.provider.send("evm_increaseTime", [100]);
      
      // Complete lease
      await gpulease.connect(user).completeLease(0);
      
      // Try to cancel
      await expect(
        gpulease.connect(user).cancelLease(0)
      ).to.be.revertedWith("Lease already completed");
    });
    
    it("should revert if trying to cancel after 5 minutes", async function () {
      // Advance time to simulate lease start
      await ethers.provider.send("evm_increaseTime", [300]); // 5 minutes
      
      // Try to cancel
      await expect(
        gpulease.connect(user).cancelLease(0)
      ).to.be.revertedWith("Cannot cancel after 5 minutes");
    });
  });
  
  describe("Status and Balance", function () {
    it("should return correct lease status", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Check lease status
      const [active, completed, startTime, duration] = await gpulease.getLeaseStatus(0);
      expect(active).to.be.true;
      expect(completed).to.be.false;
      expect(startTime).to.be.gt(0);
      expect(duration).to.equal(duration);
    });
    
    it("should return correct contract balance", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Check contract balance
      expect(await gpulease.getContractBalance()).to.equal(parseEther("100"));
      
      // Complete lease
      await ethers.provider.send("evm_increaseTime", [100]);
      await gpulease.connect(user).completeLease(0);
      
      // Check contract balance
      expect(await gpulease.getContractBalance()).to.equal(0);
    });
  });
  
  describe("Access Control", function () {
    it("should revert if non-owner tries to complete lease", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Try to complete lease as a non-owner
      await expect(
        gpulease.connect(provider).completeLease(0)
      ).to.be.revertedWith("Only user or provider can call this");
    });
    
    it("should revert if non-owner tries to cancel lease", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Try to cancel lease as a non-owner
      await expect(
        gpulease.connect(provider).cancelLease(0)
      ).to.be.revertedWith("Only user or provider can call this");
    });
  });
  
  describe("Event Emission", function () {
    it("should emit LeaseStarted event", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      const tx = await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Check LeaseStarted event
      await expect(tx).to.emit(gpulease, "LeaseStarted").withArgs(
        0, // leaseId
        user.address,
        provider.address,
        duration,
        parseEther("100")
      );
    });
    
    it("should emit PaymentReceived event", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      const tx = await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Check PaymentReceived event
      await expect(tx).to.emit(gpulease, "PaymentReceived").withArgs(
        0, // leaseId
        parseEther("100")
      );
    });
    
    it("should emit LeaseCompleted event", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Advance time to simulate lease completion
      await ethers.provider.send("evm_increaseTime", [100]);
      
      // Complete lease
      const tx = await gpulease.connect(user).completeLease(0);
      
      // Check LeaseCompleted event
      await expect(tx).to.emit(gpulease, "LeaseCompleted").withArgs(
        0, // leaseId
        parseEther("0"), // refund
        parseEther("100") // providerReward
      );
    });
    
    it("should emit LeaseCancelled event", async function () {
      const duration = 100;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("100"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("100"));
      
      // Start lease
      await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Advance time to simulate lease start
      await ethers.provider.send("evm_increaseTime", [50]);
      
      // Cancel lease
      const tx = await gpulease.connect(user).cancelLease(0);
      
      // Check LeaseCancelled event
      await expect(tx).to.emit(gpulease, "LeaseCancelled").withArgs(
        0, // leaseId
        parseEther("100") // refund
      );
    });
  });
  
  describe("Edge Cases", function () {
    it("should handle zero duration and price", async function () {
      // Try to start lease with zero duration and price
      await expect(
        gpulease.connect(user).startLease(0, 0, provider.address)
      ).to.be.revertedWith("Duration must be > 0");
    });
    
    it("should handle very long leases", async function () {
      const duration = 1000000;
      const pricePerSecond = 1;
      
      // Transfer tokens to user
      await mockToken.connect(owner).transfer(user.address, parseEther("1000000"));
      
      // Approve GPULease contract to spend tokens
      await mockToken.connect(user).approve(gpulease.target, parseEther("1000000"));
      
      // Start lease
      const tx = await gpulease.connect(user).startLease(duration, pricePerSecond, provider.address);
      
      // Check lease status
      const lease = await gpulease.leases(0);
      expect(lease.duration).to.equal(duration);
      expect(lease.totalAmount).to.equal(parseEther("1000000"));
      
      // Advance time to simulate lease completion
      await ethers.provider.send("evm_increaseTime", [duration]);
      
      // Complete lease
      await gpulease.connect(user).completeLease(0);
      
      // Check contract balance
      expect(await mockToken.balanceOf(gpulease.target)).to.equal(0);
    });
  });
})