const fs = require("fs");

const depthData = fs
  .readFileSync("input.txt", "utf8")
  .split("\n")
  .map((m) => parseInt(m));

const solution = depthData
  .map((currentDepth, idx) => (currentDepth > depthData[idx - 1] ? 1 : 0))
  .reduce((acc, val) => acc + val);

console.log(solution);
