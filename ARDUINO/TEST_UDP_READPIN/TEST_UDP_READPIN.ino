#include <ESP8266WiFi.h>
#include <WiFiUdp.h>

WiFiUDP udp;

char packetBuffer[255];
unsigned int localPort = 1053;

const char *ssid = "LATITUDE";
const char *password = "12341234";

void setup() {

  pinMode(5, INPUT_PULLUP);
  
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
  

}

int y = 0;
int etatPrecedent = 0;
void loop() {

  int etatPresent = digitalRead(5);

  if(etatPresent != etatPrecedent) 
  {
    etatPrecedent = etatPresent;

    Serial.println(etatPresent);

    if(!etatPresent) 
    {
      udp.beginPacket("10.42.0.1", 1053);
      udp.write("ON");
      udp.endPacket();
    }
    else
    {
      udp.beginPacket("10.42.0.1", 1053);
      udp.write("OFF");
      udp.endPacket();
    }
    
  }
}
