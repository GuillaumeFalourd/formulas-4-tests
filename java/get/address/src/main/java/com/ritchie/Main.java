package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {

  public static void main(String[] args) {

    String cep = System.getenv("CEP");

    Formula formula = new Formula(cep);
    formula.Run();
  }
}
