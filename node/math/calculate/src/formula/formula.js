const clc = require("cli-color")

function execute(command) {
    const { exec } = require('child_process');

    exec(command, (error, stdout, stderr) => {
        if (error) {
          console.error(`exec error: ${error}`);
          return;
        }
        console.log(`stdout: ${stdout}`);
        console.error(`stderr: ${stderr}`);
    })
}

function Run(numberOne, numberTwo, operation) {
    console.log("Number operations")

    if(operation === 'sum'){
        execute(`echo '{ "number_one": "'${numberOne}'", "number_two": "'${numberTwo}'", "operation": "'${operation}'" }' | rit math sum numbers --stdin`)
    } else if (operation === 'multiply'){
        execute(`echo '{ "number_one": "'${numberOne}'", "number_two": "'${numberTwo}'", "operation": "'${operation}'" }' | rit math multiply numbers --stdin`)
    } else {
        console.log(clc.red("(Invalid Input) Use sum or multiply!"))
    }
}

const formula = Run
module.exports = formula
