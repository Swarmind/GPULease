import  ethers  from "hardhat";

async function main() {
  // Получаем контракт
  const GPULease = await ethers.getContractFactory("GPULease");

  // Получаем токен (например, USDC)
  const USDC = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"; // USDC на Sepolia

  // Развертываем контракт
  const gpuLease = await GPULease.deploy(USDC);
  await gpuLease.waitForDeployment();

  console.log("GPULease deployed to:", gpuLease.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
