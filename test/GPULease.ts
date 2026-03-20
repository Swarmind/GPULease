import { expect } from "chai";
import { network } from "hardhat";
const { ethers } = await network.connect();
describe("GPULease", function () {
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;

  let token: any;
  let lease: any;

  beforeEach(async () => {
    [owner, user, provider, treasury] = await ethers.getSigners();

    // Deploy MockERC20
    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MOCK");

    // Mint tokens to user
    await token.mint(user.address, ethers.parseEther("1000"));

    // Deploy GPULease
    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(token.target, treasury.address);

    // Approve + deposit
    await token.connect(user).approve(lease.target, ethers.parseEther("1000"));
    await lease.connect(user).deposit(ethers.parseEther("1000"));
  });

  it("should deposit correctly", async () => {
    const balance = await lease.userBalance(user.address);
    expect(balance).to.equal(ethers.parseEther("1000"));
  });

  it("should start lease and freeze funds", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    const frozen = await lease.frozenFunds(0);
    expect(frozen).to.be.gt(0);

    const userBalance = await lease.userBalance(user.address);
    expect(userBalance).to.be.lt(ethers.parseEther("1000"));
  });

  it("should pause and resume lease", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    await lease.pauseLease(0);

    let leaseData = await lease.leases(0);
    expect(leaseData.paused).to.equal(true);

    await lease.resumeLease(0);

    leaseData = await lease.leases(0);
    expect(leaseData.paused).to.equal(false);
  });

  it("should complete lease and distribute funds", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    // немного прокрутим время
    await ethers.provider.send("evm_increaseTime", [500]);
    await ethers.provider.send("evm_mine", []);

    await lease.completeLease(0);

    const leaseData = await lease.leases(0);
    expect(leaseData.completed).to.equal(true);

    const providerBalance = await lease.userBalance(provider.address);
    expect(providerBalance).to.be.gt(0);

    const treasuryBalance = await lease.userBalance(treasury.address);
    expect(treasuryBalance).to.be.gt(0);
  });

  it("should return frozen funds list", async () => {
    const duration = 1000;
    const price = ethers.parseEther("0.001");

    await lease.startLease(
      duration,
      price,
      price,
      provider.address,
      user.address
    );

    const result = await lease.getUserFrozenFunds(user.address);

    expect(result.length).to.equal(1);
    expect(result[0].leaseId).to.equal(0);
    expect(result[0].amount).to.be.gt(0);
  });

  it("should withdraw correctly", async () => {
    await lease.connect(user).withdraw(ethers.parseEther("100"));

    const balance = await lease.userBalance(user.address);
    expect(balance).to.equal(ethers.parseEther("900"));
  });

  it("should not allow withdraw more than balance", async () => {
    await expect(
      lease.connect(user).withdraw(ethers.parseEther("2000"))
    ).to.be.revertedWith("insufficient balance");
  });

  it("should allow owner to change platform fee", async () => {
    await lease.setPlatformFee(10);
    const fee = await lease.platformFeePercentage();
    expect(fee).to.equal(10);
  });

  it("should move treasury balance on change", async () => {
    // artificially give treasury balance
    await lease.setPlatformFee(10);

    await lease.setTreasury(provider.address);

    const newTreasury = await lease.treasury();
    expect(newTreasury).to.equal(provider.address);
  });
});