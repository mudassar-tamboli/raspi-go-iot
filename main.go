// Author : Mudassar Tamboli
// Date   : 6/10/2018

package main

import (
	"fmt"
	PUBSUB "github.com/mudassar-tamboli/raspi-go-iot/pubsub"
	GPIO  "github.com/mudassar-tamboli/raspi-go-iot/gpio"
)

func main() {

	fmt.Println("IoT Of LED started");
	
	PUBSUB.Init()
	GPIO.Start()

	fmt.Println("IoT Of LED exited");
	return
}


