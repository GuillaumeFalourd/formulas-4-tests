package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {

  public static void main(String[] args) {
    String input1 = System.getenv("NUMBER_ONE");
    String input2 = System.getenv("NUMBER_TWO");
    Formula formula = new Formula(Integer.valueOf(input1), Integer.valueOf(input2));
    System.out.println(formula.Run());
  }
}
