#include <WiFi.h>
#include <WiFiUdp.h>
#include <Adafruit_NeoPixel.h>


#define ID 0

//Wifi settings - be sure to replace these with the WiFi network that your computer is connected to

const char *ssid = "SFR-6fd8";
const char *password = "D96ZQHYEGWHQ";

// LED Strip
const int numLeds = 4; // Change if your setup has more or less LED's
const int numberOfChannels = numLeds * 3; // Total number of DMX channels you want to receive (1 led = 3 channels)


Adafruit_NeoPixel pixels1(30, 18, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels2(30, 19, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels3(30, 21, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels4(30, 22, NEO_GRB + NEO_KHZ800);


WiFiUDP udp;
uint8_t packetBuffer[1023];
unsigned int localPort = 1053;

// connect to wifi – returns true if successful or false if not
boolean ConnectWifi(void)
{
  boolean state = true;
  int i = 0;

  WiFi.begin(ssid, password);
  Serial.println("");
  Serial.println("Connecting to WiFi");

  // Wait for connection
  Serial.print("Connecting");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
    if (i > 20) {
      state = false;
      break;
    }
    i++;
  }
  if (state) {
    Serial.println("");
    Serial.print("Connected to ");
    Serial.println(ssid);
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());

    double couleur = 0 * 256 * 256 + 100 * 256;

    pixels1.fill(couleur, 0, 30);
    pixels1.show();
    pixels2.fill(couleur, 0, 30);
    pixels2.show();
    pixels3.fill(couleur, 0, 30);
    pixels3.show();
    pixels4.fill(couleur, 0, 30);
    pixels4.show();
          

    
  } else {
    Serial.println("");
    Serial.println("Connection failed.");

    double couleur = 0;

    pixels1.fill(couleur, 0, 30);
    pixels1.show();
    pixels2.fill(couleur, 0, 30);
    pixels2.show();
    pixels3.fill(couleur, 0, 30);
    pixels3.show();
    pixels4.fill(couleur, 0, 30);
    pixels4.show();
    
  }

  return state;
}

void onDmxFrame(uint16_t universe, uint16_t length, uint8_t sequence, uint8_t* data)
{
  double couleur;
  

  if ((ID < 32) && (universe == 1))
  {
    for (int i = 0; i < 4; i++)
  {
          
      couleur = data[(ID * 12 +1) + (i * 3)] * 256 * 256 + data[(ID * 12+1) + (i * 3) + 1] * 256 + data[(ID * 12+1) + (i * 3) + 2];
           
      
      switch (i)
      {
        case 0 :
          pixels1.fill(couleur, 0, 30);
          pixels1.show();
          break;
        case 1 :
          pixels2.fill(couleur, 0, 30);
          pixels2.show();
          break;
        case 2 :
          pixels3.fill(couleur, 0, 30);
          pixels3.show();
          break;
        case 3 :
          pixels4.fill(couleur, 0, 30);
          pixels4.show();
          break;

      }
    
  }     
  }
  
  if ((ID >= 32) && (universe == 2))
  {
    for (int i = 0; i < 4; i++)
  {
          
      couleur = data[((ID-32) * 12+1) + (i * 3)] * 256 * 256 + data[((ID-32) * 12+1) + (i * 3) + 1] * 256 + data[((ID-32) * 12+1) + (i * 3) + 2];
           
      
      switch (i)
      {
        case 0 :
          pixels1.fill(couleur, 0, 30);
          pixels1.show();
          break;
        case 1 :
          pixels2.fill(couleur, 0, 30);
          pixels2.show();
          break;
        case 2 :
          pixels3.fill(couleur, 0, 30);
          pixels3.show();
          break;
        case 3 :
          pixels4.fill(couleur, 0, 30);
          pixels4.show();
          break;

      }
    
  }     
  }
  
}



void setup()
{
  Serial.begin(115200);
  
  pixels1.begin();
  pixels2.begin();
  pixels3.begin();
  pixels4.begin();

  double couleur = 100 * 256 * 256;

    pixels1.fill(couleur, 0, 30);
    pixels1.show();
    pixels2.fill(couleur, 0, 30);
    pixels2.show();
    pixels3.fill(couleur, 0, 30);
    pixels3.show();
    pixels4.fill(couleur, 0, 30);
    pixels4.show();

    
    ConnectWifi();
  udp.begin(localPort);

}

void loop()
{
  int packetSize = udp.parsePacket();

  if (packetSize)
  {
    int len = udp.read(packetBuffer, 1023);

    Serial.println("id");
    for (int i = 0; i < 20; i++)
  {
    Serial.println(packetBuffer[i]);
  }
/*
    Serial.println("value1");
    Serial.println(packetBuffer[4]);

    Serial.println("value2");
    Serial.println(packetBuffer[5]);

    Serial.println("value3");
    Serial.println(packetBuffer[6]);
    */
  
    onDmxFrame(packetBuffer[0], 0, 0, packetBuffer);

  
  }

}
