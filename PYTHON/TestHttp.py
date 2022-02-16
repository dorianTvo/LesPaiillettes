import requests
import time


for i in range(200):
	r = requests.get('http://192.168.249.127/LEDVertON')
	time.sleep(0.1)
	r = requests.get('http://192.168.249.127/LEDBleuON2')
	time.sleep(0.1)
	r = requests.get('http://192.168.249.127/LEDRougeON')
	time.sleep(0.1)
	r = requests.get('http://192.168.249.127/LEDVertON2')
	time.sleep(0.1)
	r = requests.get('http://192.168.249.127/LEDBleuON')
	time.sleep(0.1)
	r = requests.get('http://192.168.249.127/LEDRougeON2')
	time.sleep(0.1)
	
	
