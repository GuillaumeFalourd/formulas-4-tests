package com.ritchie.formula

class Formula(
    private val numberOne: Int, 
    private val numberTwo: Int
    ) {
    fun Run() {
        val sum = numberOne + numberTwo
        println(String.format("Sum = %s", sum))
    }
}
