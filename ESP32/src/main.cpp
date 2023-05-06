#include <Arduino.h>
#include <string>
#include <WiFi.h>
using namespace std;

const char ssid[10] = "S22_plus", pass[10] = "eqms6209";
const int monitor_speed = 921600, activated = 4095, hz = 512;
const int R = 23, G = 22, B = 21, sensorPin = 32, laserPin = 25, buzzerPin = 26;
int timer = 0, state = 0;

void connectToWiFi(){

  do{

    WiFi.begin(ssid, pass);

    Serial.print("Connecting...");


    int maxAttempts = 100;

    int numAttempts = 0;

    while(WiFi.status() != WL_CONNECTED){

      Serial.print(".");

      delay(100);

      if(++numAttempts >= maxAttempts){ // If no connection after maxAttempts * 100 milliseconds (100 max = 10 seconds)

        Serial.print("Could not connect - waiting 60 seconds to try again");

        delay(60000); //Wait for 1 minute

        break;

      }

    }

  } while(WiFi.status() != WL_CONNECTED);

  Serial.println("");

  Serial.println("Connected!");

}

void setup() {
  Serial.begin(monitor_speed);
  // rgb pins
  pinMode(R, OUTPUT);
  pinMode(G, OUTPUT);
  pinMode(B, OUTPUT);

  // buzzer pin
  pinMode(buzzerPin, OUTPUT);

  // laser pin
  pinMode(laserPin, OUTPUT);
  Serial.println("----");
  connectToWiFi();
}

void loop() {

  // start laser HIGH = ON / LOW = OFF
  digitalWrite(laserPin, HIGH);
  int laserRead = analogRead(sensorPin);
  delay(1000);

  //check if laser hit sensor, and change color correspondingly to it
  if (laserRead != activated) { 
    timer++;
    Serial.print("timing: ");
    Serial.println(timer);
    if (timer == 5) //turns light red if timer exceeds 5s
    { // checks if laser is blocked for longer period
      analogWrite(R, 255);
      analogWrite(G, 0);
      analogWrite(B, 0);
      tone(buzzerPin, hz, 1000);
      // next step, send API-call to /request
      state = 1;
      Serial.println("Trash is full, request is sent");
    }
  } else { //turns light green 
    analogWrite(R, 0);
    analogWrite(G, 255);
    analogWrite(B, 0);
    timer = 0;
    if(state == 1){
      // next step, send API-call to /invoice
      state = 0;
      Serial.println("Trash has been emptied, invoice is sent");
    }
  }
}