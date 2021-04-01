const clc = require("cli-color")

function Run(numberOne, numberTwo) {
    var sum = parseInt(numberOne) + parseInt(numberTwo)
    console.log(clc.green("Sum = "+ sum))
}

const formula = Run
module.exports = formula
