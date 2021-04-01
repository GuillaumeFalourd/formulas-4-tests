const clc = require("cli-color")

function Run(numberOne, numberTwo) {
    var multiplication = parseInt(numberOne) * parseInt(numberTwo)
    console.log(clc.green("Multiplication = "+ multiplication))
}

const formula = Run
module.exports = formula
