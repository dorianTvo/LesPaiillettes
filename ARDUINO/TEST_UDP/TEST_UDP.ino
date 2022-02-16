#include <ESP8266WiFi.h>
#include <WiFiUdp.h>

WiFiUDP udp;

char packetBuffer[255];
unsigned int localPort = 8080;

const char *ssid = "LATITUDE";
const char *password = "vzam011101";

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);

  udp.begin(localPort);

  Serial.print(F("UDP Client : ")); 
  Serial.println(WiFi.localIP());
}

void loop() {
  int packetSize = udp.parsePacket();

  if (packetSize) {
     Serial.println("EN"); 
  }
  
  Serial.println("FR");
  delay(500);

}