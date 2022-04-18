#include <WiFi.h>
#include <WiFiUdp.h>
#include <ArtnetWifi.h>
//#include <FastLED.h>
#include <Adafruit_NeoPixel.h>


#define ID 1

//Wifi settings - be sure to replace these with the WiFi network that your computer is connected to

const char *ssid = "nom du routeur";
const char *password = "mot de passe du routeur";

// LED Strip
const int numLeds = 4; // Change if your setup has more or less LED's
const int numberOfChannels = numLeds * 3; // Total number of DMX channels you want to receive (1 led = 3 channels)


Adafruit_NeoPixel pixels1(30, 23, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels2(30, 22, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels3(30, 1, NEO_GRB + NEO_KHZ800);
Adafruit_NeoPixel pixels4(30, 3, NEO_GRB + NEO_KHZ800);


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
  sendFrame = 1;
  // set brightness of the whole strip
  if (universe == 15)
  {
    //FastLED.setBrightness(data[0]);
  }
  // read universe and put into the right part of the display buffer

  for (int i = 0; i < length / 3; i++)
  {
    couleur = data[ID * i * 3] * 256 * 256 + data[ID * i * 3 + 1] * 256 + data[ID * i * 3 + 2];

    int led = i + (universe - startUniverse) * (previousDataLength / 3);
    if (led < numLeds)
    {

      switch (led)
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

  //previousDataLength = length;

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
  artnet.begin();


  // onDmxFrame will execute every time a packet is received by the ESP32
  artnet.setArtDmxCallback(onDmxFrame);
}

void loop()
{
  // we call the read function inside the loop
  artnet.read();
}
