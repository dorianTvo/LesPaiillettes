package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("RECEPTION REQUETE")
		go dorian_esp_LED()
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":80"))
}

func main2() {
	for i := 0; i < 20; i++ {
		go colin_esp_LEDOn()
		go dorian_esp_LED()

		/*for y := 0; y < 64; y++ {
			fausse_requete()
		}*/

		time.Sleep(1 * time.Second)

		go colin_esp_LEDOff()
		go dorian_esp_LEDOff()

		time.Sleep(1 * time.Second)
	}

	time.Sleep(1 * time.Second)
}

func fausse_requete() {
	_, err := http.Get("http://192.168.249.204/")
	if err != nil {
		log.Fatalln(err)
	}
}

func colin_esp_LEDOn() {
	_, err := http.Get("http://192.168.249.127/LEDVertON")
	if err != nil {
		log.Fatalln(err)
	}
}

func colin_esp_LEDOff() {
	_, err := http.Get("http://192.168.249.127/LEDOFF")
	if err != nil {
		log.Fatalln(err)
	}
}

func dorian_esp_LED() {
	_, err := http.Get("http://192.168.249.114/LED")
	if err != nil {
		log.Fatalln(err)
	}
}

func dorian_esp_LEDOff() {
	_, err := http.Get("http://192.168.249.114/LEDOff")
	if err != nil {
		log.Fatalln(err)
	}
}
