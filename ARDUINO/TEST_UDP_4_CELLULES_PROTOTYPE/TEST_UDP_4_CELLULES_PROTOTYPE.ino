#include <ESP8266WiFi.h>
#include <WiFiUdp.h>
#include <ArduinoJson.h>
#include <Adafruit_NeoPixel.h>


int i;
Adafruit_NeoPixel pixels1(36, 2, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels2(36, 5, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels3(36, 4, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels4(36, 0, NEO_GRB + NEO_KHZ800);

WiFiUDP udp;

char packetBuffer[1023];
unsigned int localPort = 1053;
const int ID = 0;

const char *ssid = "TP-Link_BFDC";
const char *password = "56640113";

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);

  pixels1.begin();
  pixels2.begin();
  pixels3.begin();
  pixels4.begin();

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

  double couleur, bleu, rouge, vert;

  bleu = 0;
  rouge = 0;
  vert = 10;

  couleur = rouge * 256 * 256 + vert * 256 + bleu;

  pixels1.fill(couleur, 0, 36);
  pixels1.show();
  pixels2.fill(couleur, 0, 36);
  pixels2.show();
  pixels3.fill(couleur, 0, 36);
  pixels3.show();
  pixels4.fill(couleur, 0, 36);
  pixels4.show();
  

  
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

double couleur, bleu, rouge, vert;

  bleu = (int) obj["RgbColor"][2];
  rouge = (int) obj["RgbColor"][0];
  vert = (int) obj["RgbColor"][1];

  couleur = rouge * 256 * 256 + vert * 256 + bleu;

  if(number == 0)
  {
    pixels1.fill(couleur, 0, 36);
    pixels1.show();
  }
  if(number == 1)
  {
    pixels2.fill(couleur, 0, 36);
    pixels2.show();
  }
  if(number == 2)
  {
    pixels3.fill(couleur, 0, 36);
    pixels3.show();
  }
  if(number == 3)
  {
    pixels4.fill(couleur, 0, 36);
    pixels4.show();
  }


  delay(10);

  
}
