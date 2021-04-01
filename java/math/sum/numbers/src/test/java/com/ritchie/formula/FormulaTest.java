package com.ritchie.formula;

import static org.junit.Assert.*;

import org.junit.Test;

public class FormulaTest {

  @Test
  public void run() {
    Formula formula = new Formula(1, 2);
    String expected = "The sum is 3";
    assertEquals(expected, formula.Run());
  }
}
