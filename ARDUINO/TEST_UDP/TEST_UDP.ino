#include <ESP8266WiFi.h>
#include <WiFiUdp.h>

WiFiUDP udp;

char packetBuffer[255];
unsigned int localPort = 1053;

const char *ssid = "LATITUDE";
const char *password = "12341234";

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);

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

  udp.beginPacket("10.42.0.1", 1053);
  udp.write("Hello from client");
  udp.endPacket();

  delay(500);
  
  int packetSize = udp.parsePacket();

  if (packetSize) {
      int len = udp.read(packetBuffer, 255);
      
      if (len > 0) 
        packetBuffer[len - 1] = 0;
        
      Serial.println(packetBuffer);
  
      udp.beginPacket("10.42.0.1", 1053);
      udp.write("Hello from client");
      udp.endPacket();
  } 
}

int y = 0;
  
void loop() {
}
