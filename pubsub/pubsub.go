// Author : Mudassar Tamboli
// Date   : 6/10/2018

package pubsub

import (
         "fmt"
         "os"
    MQTT "github.com/eclipse/paho.mqtt.golang"
         "github.com/mudassar-tamboli/raspi-go-iot/gpio"
)


func Init() {

    fmt.Println("pubsub : Raspberry Pi Pub/Sub initializing...")
    optsPub := MQTT.NewClientOptions()
    //optsPub.AddBroker("tcp://iot.eclipse.org:1883")
    optsPub.AddBroker("tcp://m14.cloudmqtt.com:14205")
    optsPub.SetUsername("htmbxcyz")     // TODO
    optsPub.SetPassword("rH2_IZj43nDy") // TODO
    optsPub.SetClientID("rasp-pi-go")
    optsPub.SetCleanSession(false)
    optsPub.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
        fmt.Println("SetDefaultPublishHandler : ", msg.Topic(), string(msg.Payload()))
    })

	fmt.Println("pubsub : Raspberry Pi MQTT broker configured")

    clientPub := MQTT.NewClient(optsPub)
    if tokenPub := clientPub.Connect(); tokenPub.Wait() && tokenPub.Error() != nil {
        panic(tokenPub.Error())
    }

    gpio.PubGreenLedStatus = func(ledMapJSON interface{}) {
        tokenPub := clientPub.Publish("plain/led/status/green", 0, false, ledMapJSON)
        tokenPub.Wait()
    }

    gpio.PubBlueLedStatus = func(ledMapJSON interface{}) {
        tokenPub := clientPub.Publish("secure/led/status/blue", 0, false, ledMapJSON)
        tokenPub.Wait()
    }

    gpio.PubRedLedStatus = func(ledMapJSON interface{}) {
        tokenPub := clientPub.Publish("secure/led/status/red", 0, false, ledMapJSON)
        tokenPub.Wait()
    }
    
    if tokenPub := clientPub.Subscribe("secure/led/action/red", 0, gpio.SubRedLedAction); tokenPub.Wait() && tokenPub.Error() != nil {
        fmt.Println(tokenPub.Error())
        os.Exit(1)
    }
	
	fmt.Println("pubsub : Raspberry Pi Pub/Sub callbacks registered")

}

