package com.ritchie.formula;

import com.google.gson.*;

public class Formula {

  private String username;
  private String token;
  private String email;

  public String Run() {
    Gson gson = new Gson();
    return gson.toJson(this, Formula.class);
  }

  public Formula(String username, String token, String email) {
    this.username = username;
    this.token = token;
    this.email = email;
  }

  public String getUsername() {
    return username;
  }

  public void setUsername(String username) {
    this.username = username;
  }

  public String getToken() {
    return token;
  }

  public void setToken(String token) {
    this.token = token;
  }

  public String getEmail() {
    return email;
  }

  public void setEmail(String email) {
    this.email = email;
  }
}
