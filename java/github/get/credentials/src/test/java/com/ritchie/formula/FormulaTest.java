package com.ritchie.formula;

import static org.junit.Assert.*;
import com.google.gson.*;

import org.junit.Test;

public class FormulaTest {

  @Test
  public void run() {
    Formula formula = new Formula("Username", "Token", "Email");
    Gson gson = new Gson();
    String expected = gson.toJson(formula, Formula.class);
    assertEquals(expected, formula.Run());
  }
}
