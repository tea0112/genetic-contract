const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const ControllerModule = buildModule("ControllerModule", (m) => {
  const controller = m.contract("Controller");

  return { controller };
});

module.exports = ControllerModule