const fs = require("fs");

const movementData = fs.readFileSync("input.txt", "utf8").split("\n");

const getInput = (name) => {
  return movementData
    .filter((v) => v.includes(name))
    .map((v) => parseInt(v.split(" ")[1]))
    .reduce((a, b) => a + b);
};

const forward = getInput("forward");
const up = getInput("up");
const down = getInput("down");

const y = down - up;

console.log("x: " + forward, "y: " + y, "product: " + forward * y);
