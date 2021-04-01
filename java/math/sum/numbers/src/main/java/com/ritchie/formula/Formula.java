package com.ritchie.formula;

public class Formula {

  private Integer input1;
  private Integer input2;

  public String Run() {
    Integer sum = input1 + input2;
    return String.format("The sum is %s", sum);
  }

  public Formula(Integer input1, Integer input2) {
    this.input1 = input1;
    this.input2 = input2;
  }

  public Integer getInput1() {
    return input1;
  }

  public void setInput1(Integer input1) {
    this.input1 = input1;
  }

  public Integer getInput2() {
    return input2;
  }

  public void setInput2(Integer input2) {
    this.input2 = input2;
  }
}
