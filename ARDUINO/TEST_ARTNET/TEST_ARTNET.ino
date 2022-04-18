#include <ESP8266WiFi.h>
#include <WiFiUdp.h>
#include <ArtnetWifi.h>
#include <FastLED.h>

//Wifi settings - be sure to replace these with the WiFi network that your computer is connected to
const char *ssid = "nom du routeur";
const char *password = "mot de passe du routeur";

// LED Strip
const int numLeds = 4; // Change if your setup has more or less LED's
const int numberOfChannels = numLeds * 3; // Total number of DMX channels you want to receive (1 led = 3 channels)

#define DATA_PIN1 2 //The data pin that the WS2812 strips are connected to.
#define DATA_PIN2 5 //The data pin that the WS2812 strips are connected to.
#define DATA_PIN3 4 //The data pin that the WS2812 strips are connected to.
#define DATA_PIN4 0 //The data pin that the WS2812 strips are connected to.  

CRGB leds1[30];
CRGB leds2[30];
CRGB leds3[30];
CRGB leds4[30];

// Artnet settings
ArtnetWifi artnet;
const int startUniverse = 0;

bool sendFrame = 1;
int previousDataLength = 0;

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
    if (i > 20){
      state = false;
      break;
    }
    i++;
  }
  if (state){
    Serial.println("");
    Serial.print("Connected to ");
    Serial.println(ssid);
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());
  } else {
    Serial.println("");
    Serial.println("Connection failed.");
  }

  return state;
}

void onDmxFrame(uint16_t universe, uint16_t length, uint8_t sequence, uint8_t* data)
{
  sendFrame = 1;
  // set brightness of the whole strip 
  if (universe == 15)
  {
    FastLED.setBrightness(data[0]);
  }
  // read universe and put into the right part of the display buffer
  
   for (int i = 0; i < length / 3; i++)
  {
    int led = i + (universe - startUniverse) * (previousDataLength / 3);
    if (led < numLeds)
    {
      for(int j = 0;j<30;j++)
      {
        switch (led)
        {
          case 0 :
            leds1[j] = CRGB(data[i * 3], data[i * 3 + 1], data[i * 3 + 2]);
            break;
          case 1 :
            leds2[j] = CRGB(data[i * 3], data[i * 3 + 1], data[i * 3 + 2]);
            break;
         case 2 :
            leds3[j] = CRGB(data[i * 3], data[i * 3 + 1], data[i * 3 + 2]);
            break;
         case 3 :
            leds4[j] = CRGB(data[i * 3], data[i * 3 + 1], data[i * 3 + 2]);
            break;
          
        }
        
      }
      
   
    }

    
  }
     
  previousDataLength = length; 
    FastLED.show(); 
 
}

void setup()
{
  Serial.begin(115200);
  ConnectWifi();
  artnet.begin();
  
  FastLED.addLeds<WS2812, DATA_PIN1, GRB>(leds1, 30);
  FastLED.addLeds<WS2812, DATA_PIN2, GRB>(leds2, 30);
  FastLED.addLeds<WS2812, DATA_PIN3, GRB>(leds3, 30);
  FastLED.addLeds<WS2812, DATA_PIN4, GRB>(leds4, 30);

  // onDmxFrame will execute every time a packet is received by the ESP32
  artnet.setArtDmxCallback(onDmxFrame);
}

void loop()
{
  // we call the read function inside the loop
  artnet.read();
}
