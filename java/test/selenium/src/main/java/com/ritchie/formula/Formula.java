package com.ritchie.formula;

import org.openqa.selenium.WebDriver;
//import org.openqa.selenium.firefox.FirefoxDriver;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.chrome.ChromeOptions;

public class Formula {

    private final String url;
    private final String title;

    public Formula(String url, String title) {
        this.url = url;
        this.title = title;
    }

    public void Run() {
        // declaration and instantiation of objects/variables

        // System.setProperty("webdriver.gecko.driver","C:\\geckodriver.exe");
		// WebDriver driver = new FirefoxDriver();
		// System.setProperty("webdriver.chrome.driver","/usr/bin/local/chromedriver");
        ChromeOptions options = new ChromeOptions();
        options.addArguments("disable-infobars");
        options.addArguments("--start-maximized");
        options.addArguments("--headless");
        WebDriver driver = new ChromeDriver(options);

        String baseUrl = url;
        String expectedTitle = title;
        String actualTitle = "";

        // Launch Navigator and direct it to the Base URL
        driver.get(baseUrl);

        // Get Actual Title
        actualTitle = driver.getTitle();

        /*
         * Compare the actual title of the page with the expected one and print
         * The result as "Passed" or "Failed"
         */
        if (actualTitle.contentEquals(expectedTitle)){
            System.out.println("\u001B[32m" + "Test Passed!" + "\u001B[0m");
        } else {
            System.out.println("\u001B[31m" + "Test Failed" + "\u001B[0m");
            System.out.println("\u001B[31m" + "Expected: " + expectedTitle + "\u001B[0m");
            System.out.println("\u001B[31m" + "Got: " + actualTitle + "\u001B[0m");
        }

        //Close Navigator
        driver.close();

        // WebDriver driver  = new ChromeDriver();
		// driver.navigate().to("http://google.com");
		// String appTitle = driver.getTitle();
		// System.out.println("Application title is :: "+appTitle);
		// driver.quit();
    }
}
