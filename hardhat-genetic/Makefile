compile:
	solc --abi ./contracts/Controller.sol -o ./contracts --overwrite
	abigen --abi ./contracts/Controller.abi --pkg contracts --type Controller --out ./contracts/Controller.go
node:
	npx hardhat node
deploy:
	npx hardhat ignition deploy ./ignition/modules/Controller.js --network localhost
