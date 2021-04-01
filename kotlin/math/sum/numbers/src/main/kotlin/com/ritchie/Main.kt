package com.ritchie

import com.ritchie.formula.Formula

object Main {
    @JvmStatic
    fun main(args: Array<String?>?) {
        val numberOne: Int = (System.getenv("RIT_NUMBER_ONE").toInt())
        val numberTwo: Int = (System.getenv("RIT_NUMBER_TWO").toInt())
        val formula = Formula(numberOne, numberTwo)
        formula.Run()
    }
}
