const fs = require("fs");

const movementData = fs.readFileSync("input.txt", "utf8").split("\n");

let aim = 0,
  x = 0,
  y = 0;

for (const movement of movementData) {
  [direction, amount] = movement.split(" ");
  amount = parseInt(amount);
  if (direction === "forward") {
    x += amount;
    y += aim * amount;
  }
  if (direction === "up") {
    aim -= amount;
  }
  if (direction === "down") {
    aim += amount;
  }
}

console.log(x * y);
