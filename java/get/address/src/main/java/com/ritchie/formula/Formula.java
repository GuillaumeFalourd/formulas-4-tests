package com.ritchie.formula;

import com.github.tomaslanger.chalk.Chalk;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

public class Formula {

    private final String cep;

    public Formula(String cep) {
        this.cep = cep;
    }

    public void Run() {
        String cepformatted = cep.replaceAll("-", "");
        try {
            String viacepUrl = "http://viacep.com.br/ws/" + cepformatted + "/json/";
            URL url = new URL(viacepUrl);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");
            conn.setRequestProperty("Accept", "application/json");
            if (conn.getResponseCode() != 200) {
                throw new RuntimeException("Failed : HTTP Error code : "
                        + conn.getResponseCode());
            }
            InputStreamReader in = new InputStreamReader(conn.getInputStream());
            BufferedReader br = new BufferedReader(in);
            String output;
            while ((output = br.readLine()) != null) {
                System.out.println(output);
            }
            conn.disconnect();

        } catch (Exception e) {
            System.out.println("Exception: " + e);
        }
    }
}
