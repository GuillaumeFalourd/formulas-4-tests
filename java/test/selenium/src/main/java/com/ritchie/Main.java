package com.ritchie;

import com.ritchie.formula.Formula;

public class Main {

  public static void main(String[] args) {

    String url = System.getenv("RIT_URL");
    String title = System.getenv("RIT_TITLE");

    Formula formula = new Formula(url, title);
    formula.Run();
  }
}
