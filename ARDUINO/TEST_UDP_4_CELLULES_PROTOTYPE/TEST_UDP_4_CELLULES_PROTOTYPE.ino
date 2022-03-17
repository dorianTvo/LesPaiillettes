#include <ESP8266WiFi.h>
#include <WiFiUdp.h>
#include <ArduinoJson.h>
#include <Adafruit_NeoPixel.h>

int DataIn = 0;
Adafruit_NeoPixel pixels(9, DataIn, NEO_GRB + NEO_KHZ800);

//const int capacity = JSON_OBJECT_SIZE(4);
//DynamicJsonBuffer jb(capacity);

WiFiUDP udp;

char packetBuffer[1023];
unsigned int localPort = 1053;
const int ID = 0;

const char *ssid = "TP-Link_BFDC";
const char *password = "56640113";

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);

  pixels.begin();

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
  
  udp.begin(localPort);

  pixels.setPixelColor(8, pixels.Color(255,0,0));
  pixels.show();
  delay(100);
  pixels.setPixelColor(8, pixels.Color(0,255,0));
  pixels.show();

  
}
  
void loop() 
{
  int packetSize = udp.parsePacket();

  if (packetSize) 
  {
      int len = udp.read(packetBuffer, 1023);
      
      //if (len > 0) 
        //packetBuffer[len - 1] = 0;

      //Serial.print(packetBuffer);
        
     //JsonObject& obj = jb.parseObject(packetBuffer);

     DynamicJsonDocument obj(32768);
     deserializeJson(obj, packetBuffer);

    lightLed(0, obj[ID]["Cellules"][0]);
    lightLed(1, obj[ID]["Cellules"][1]);
    lightLed(2, obj[ID]["Cellules"][2]);
    lightLed(3, obj[ID]["Cellules"][3]); 
  } 
}

void lightLed(int number, JsonObject obj) 
{
  pixels.setPixelColor(number, pixels.Color((int) obj["RgbColor"][0], (int) obj["RgbColor"][1], (int) obj["RgbColor"][2]));
  pixels.show();


  //Serial.println("affiche");
  //Serial.println(number);
  //Serial.println((int) obj["RgbColor"][0]);
  //Serial.println((int) obj["RgbColor"][1]);
  //Serial.println((int) obj["RgbColor"][2]);
  
}
