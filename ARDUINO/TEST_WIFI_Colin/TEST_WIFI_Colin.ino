#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <Adafruit_NeoPixel.h>
#include <ESP8266HTTPClient.h>



bool etat = 0,etat_old = 0;

int DataIn = 12;
int DataIn2 = 13;
int Bouton = 15;
int httpCode;
int i;


Adafruit_NeoPixel pixels(10, DataIn, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels2(10, DataIn2, NEO_GRB + NEO_KHZ800);


// Replace with your network credentials
const char* ssid = "XXXXX";
const char* password = "XXXX";

ESP8266WebServer server(80);   //instantiate server at port 80 (http port)

WiFiClient client;
HTTPClient http;

      
String page = "";
void setup(void) {


  Serial.begin(115200);
  WiFi.begin(ssid, password); //begin WiFi connection
  Serial.println("");

  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  WiFi.setAutoReconnect(true);
  WiFi.persistent(true);

  server.on("/", []() {
    server.send(200, "text/html", page);
  });

  server.on("/LEDVertON", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels.setPixelColor(i, pixels.Color(0, 150, 0));
    pixels.show();
    

  }
    delay(10);
  });

    server.on("/LEDOFF", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels.setPixelColor(i, pixels.Color(0, 0, 0));
    pixels.show();
    

  }
    delay(10);
  });

  server.on("/LEDBleuON", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels.setPixelColor(i, pixels.Color(0, 0, 150));
    pixels.show();
    

  }
    delay(10);
  });

  server.on("/LEDRougeON", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels.setPixelColor(i, pixels.Color(150, 0, 0));
    pixels.show();
    

  }
    delay(10);
  });

server.on("/LEDVertON2", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels2.setPixelColor(i, pixels.Color(0, 150, 0));
    pixels2.show();
    

  }
    delay(10);
  });

  server.on("/LEDBleuON2", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels2.setPixelColor(i, pixels.Color(0, 0, 150));
    pixels2.show();
    

  }
    delay(10);
  });

  server.on("/LEDRougeON2", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels2.setPixelColor(i, pixels.Color(150, 0, 0));
    pixels2.show();
    

  }
    delay(10);
  });

  server.on("/LEDRougeBlanc2", []() {
    server.send(200, "text/html", page);
    for (i = 0; i <= 9; i++) {
    pixels2.setPixelColor(i, pixels.Color(150, 150, 150));
    pixels2.show();
    

  }
    delay(10);
  });

  
  
  server.begin();
  Serial.println("Web server started!");

  pinMode (Bouton, INPUT);
  
  pixels.begin();
  pixels2.begin();

  http.begin(client, "http://192.168.249.26"); 
  

  
}

void loop(void) {

  

  server.handleClient();

  etat_old = etat;
  
  etat = digitalRead(Bouton);

   
  
  if(etat == HIGH and (etat != etat_old))
  {
      //Serial.println("BOUTON");

      for (i = 0; i <= 9; i++) {
      pixels.setPixelColor(i, pixels.Color(0, 0, 150));
      pixels.show();
      }
      
      //http.begin(client, "http://192.168.249.26");  //Specify request destination
      httpCode = http.GET();                  //Send the request
      //http.end();   //Close connection

      for (i = 0; i <= 9; i++) {
      pixels.setPixelColor(i, pixels.Color(0, 0, 0));
      pixels.show();
      //Serial.println("BOUTONFini");
      }
  }
  
}
