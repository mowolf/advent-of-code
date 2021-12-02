const fs = require("fs");

const movementData = fs.readFileSync("input.txt", "utf8").split("\n");

const forward = movementData
  .filter((v) => v.includes("forward"))
  .map((v) => parseInt(v.split(" ")[1]))
  .reduce((a, b) => a + b);
const up = movementData
  .filter((v) => v.includes("up"))
  .map((v) => parseInt(v.split(" ")[1]))
  .reduce((a, b) => a + b);
const down = movementData
  .filter((v) => v.includes("down"))
  .map((v) => parseInt(v.split(" ")[1]))
  .reduce((a, b) => a + b);

y = parseInt(down - up);

console.log("x: " + forward, "y: " + y, forward * y);
