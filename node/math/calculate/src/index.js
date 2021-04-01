const run = require("./formula/formula")

const NUMBER_ONE = process.env.NUMBER_ONE
const NUMBER_TWO = process.env.NUMBER_TWO
const OPERATION = process.env.OPERATION

run(NUMBER_ONE, NUMBER_TWO, OPERATION)
