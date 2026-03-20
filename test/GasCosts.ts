import { network } from "hardhat";
const { ethers } = await network.connect();
import { expect } from "chai";

describe("GPULease Gas Tests (raw wei)", function () {
  let lease: any;
  let token: any;
  let owner: any;
  let user: any;
  let provider: any;
  let treasury: any;

  async function logGas(tx: any, label: string) {
    const receipt = await tx.wait();

    const gasUsed = receipt.gasUsed;
    const gasPrice = receipt.gasPrice ?? receipt.effectiveGasPrice;
    const totalWei = gasUsed * gasPrice;

    console.log(`\n=== ${label} ===`);
    console.log("gasUsed:", gasUsed.toString());
    console.log("gasPrice:", gasPrice.toString());
    console.log("totalWei:", totalWei.toString());
  }

  beforeEach(async () => {
    [owner, user, provider, treasury] = await ethers.getSigners();

    // Deploy token
    const Token = await ethers.getContractFactory("MockERC20");
    token = await Token.deploy("Mock", "MCK");
    await token.waitForDeployment();

    // Deploy lease contract
    const Lease = await ethers.getContractFactory("GPULease");
    lease = await Lease.deploy(token.target, treasury.address);
    await lease.waitForDeployment();

    // Setup balances
    await token.mint(user.address, ethers.parseEther("10000"));
    await token.connect(user).approve(lease.target, ethers.parseEther("10000"));

    await lease.connect(user).deposit(ethers.parseEther("1000"));
  });

  it("startLease gas", async () => {
    const tx = await lease.startLease(
      3600,
      ethers.parseEther("0.0001"),
      ethers.parseEther("0.0002"),
      provider.address,
      user.address
    );

    await logGas(tx, "startLease");
  });

  it("pause + resume gas", async () => {
    await lease.startLease(
      3600,
      ethers.parseEther("0.0001"),
      ethers.parseEther("0.0002"),
      provider.address,
      user.address
    );

    await logGas(await lease.pauseLease(0), "pauseLease");
    await logGas(await lease.resumeLease(0), "resumeLease");
  });

  it("completeLease gas", async () => {
    await lease.startLease(
      3600,
      ethers.parseEther("0.0001"),
      ethers.parseEther("0.0002"),
      provider.address,
      user.address
    );

    // имитируем время
    await ethers.provider.send("evm_increaseTime", [1800]);
    await ethers.provider.send("evm_mine", []);

    await logGas(await lease.completeLease(0), "completeLease");
  });

  it("deposit + withdraw gas", async () => {
    await logGas(
      await lease.connect(user).deposit(ethers.parseEther("100")),
      "deposit"
    );

    await logGas(
      await lease.connect(user).withdraw(ethers.parseEther("50")),
      "withdraw"
    );
  });

  it("multiple leases scaling gas", async () => {
    for (let i = 0; i < 5; i++) {
      await lease.startLease(
        3600,
        ethers.parseEther("0.0001"),
        ethers.parseEther("0.0002"),
        provider.address,
        user.address
      );
    }

    const gas = await lease.getUserFrozenFunds.estimateGas(user.address);

    console.log("\n=== getUserFrozenFunds (5 leases) ===");
    console.log("estimated gas:", gas.toString());
  });
});