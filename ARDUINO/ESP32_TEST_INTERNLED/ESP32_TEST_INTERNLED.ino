#include <ESP8266WiFi.h>
#include <WiFiUdp.h>
#include <ArtnetWifi.h>
//#include <FastLED.h>
#include <Adafruit_NeoPixel.h>


#define ID 63

//Wifi settings - be sure to replace these with the WiFi network that your computer is connected to

//const char *ssid = "nom du routeur";
//const char *password = "mot de passe du routeur";

const char *ssid = "nom du routeur";
const char *password = "mot de passe du routeur";

// LED Strip
const int numLeds = 4; // Change if your setup has more or less LED's
const int numberOfChannels = numLeds * 3; // Total number of DMX channels you want to receive (1 led = 3 channels)




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

    digitalWrite(LED_BUILTIN, HIGH);

          

    
  } else {
    Serial.println("");
    Serial.println("Connection failed.");



    
  }

  return state;
}

void onDmxFrame(uint16_t universe, uint16_t length, uint8_t sequence, uint8_t* data)
{
  double couleur;
  sendFrame = 1;

  if ((ID < 32) && (universe == 0))
  {
    if (data[(ID * 12)] > 0)
    {
      digitalWrite(LED_BUILTIN, LOW);
    }
    else
    {
      digitalWrite(LED_BUILTIN, HIGH);
    }
  
  }
  
  if ((ID >= 32) && (universe == 1))
  {
    if (data[((ID-32) * 12)] > 0)
    {
      digitalWrite(LED_BUILTIN, LOW);
    }
    else
    {
      digitalWrite(LED_BUILTIN, HIGH);
    }
  }
  
}



void setup()
{
  Serial.begin(115200);

pinMode(LED_BUILTIN, OUTPUT);
digitalWrite(LED_BUILTIN, LOW);
    
    ConnectWifi();
  artnet.begin();


  // onDmxFrame will execute every time a packet is received by the ESP32
  artnet.setArtDmxCallback(onDmxFrame);
}

void loop()
{
  // we call the read function inside the loop
  artnet.read();
}
