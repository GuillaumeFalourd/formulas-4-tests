package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {

  public static void main(String[] args) {
    String input1 = System.getenv("GIT_USER");
    String input2 = System.getenv("GIT_TOKEN");
    String input3 = System.getenv("GIT_EMAIL");
    Formula formula = new Formula(input1, input2, input3);
    System.out.println(formula.Run());
  }
}
