// Author : Mudassar Tamboli
// Date   : 6/10/2018

package gpio

import (
         "encoding/json"
         "fmt"
         "math/rand"
         "time"
    MQTT "github.com/eclipse/paho.mqtt.golang"
         "github.com/stianeikeland/go-rpio"
)

const (
	
    GreenLedPin1 uint8 = 21
    GreenLedPin2 uint8 = 20
    GreenLedPin3 uint8 = 16

    BlueLedPin1  uint8 = 7
    BlueLedPin2  uint8 = 8
    BlueLedPin3  uint8 = 25

    RedLedPin1   uint8 = 24
    RedLedPin2   uint8 = 23
    RedLedPin3   uint8 = 18    

)

type CallBackPubGreenLedStatus func(interface{})
type CallBackPubBlueLedStatus func(interface{})
type CallBackPubRedLedStatus func(interface{})

var PubGreenLedStatus CallBackPubGreenLedStatus
var PubBlueLedStatus CallBackPubBlueLedStatus
var PubRedLedStatus CallBackPubBlueLedStatus

func blockForever() {
    c := make(chan struct{})
    <-c
}

func Start() {

    rand.Seed(time.Now().UTC().UnixNano())

    fmt.Println("gpio : LED blink started")
    rpio.Open()
    defer rpio.Close()

    go goDancingGreenLed() // publisher of status of green led
    go goDancingBlueLed()  // publisher of status of blue led
    go goStillRedLed()     // Both publisher and subscriber for red led
	
	blockForever()

}

func SubRedLedAction(client MQTT.Client, msg MQTT.Message) {

    var rledPin1, rledPin2, rledPin3 rpio.Pin

    rledPin1 = rpio.Pin(RedLedPin1)
    rledPin2 = rpio.Pin(RedLedPin2)
    rledPin3 = rpio.Pin(RedLedPin3)

    b := msg.Payload()
    var m map[string]string
    //    choke <- [2]string{msg.Topic(), string(msg.Payload())}

    err := json.Unmarshal(b, &m)

    if err != nil {
        fmt.Println("gpio : SubRedLedAction : ", msg.Topic(), string(msg.Payload()))
    } else {

        fmt.Println("gpio : SubRedLedAction : ", msg.Topic(), m["LED1"], m["LED2"], m["LED3"])

        rledPin1.Output()
        rledPin2.Output()
        rledPin3.Output()

        if m["LED1"] == "ON" {
            rledPin1.Write(rpio.High)
        } else {
            rledPin1.Write(rpio.Low)
        }

        if m["LED2"] == "ON" {
            rledPin2.Write(rpio.High)
        } else {
            rledPin2.Write(rpio.Low)
        }

        if m["LED3"] == "ON" {
            rledPin3.Write(rpio.High)
        } else {
            rledPin3.Write(rpio.Low)
        }
    }
}

func goDancingGreenLed() {

    fmt.Println("gpio : goDancingGreenLed (go routine)")
    var gledPin1, gledPin2, gledPin3 rpio.Pin
    var gledState1, gledState2, gledState3 rpio.State

    gledPin1 = rpio.Pin(GreenLedPin1)
    gledPin2 = rpio.Pin(GreenLedPin2)
    gledPin3 = rpio.Pin(GreenLedPin3)

    gledPin1.Output()
    gledPin2.Output()
    gledPin3.Output()

    for {

        gledMap := make(map[string]string)

        gledState1 = GetRandomOnOff()
        gledState2 = GetRandomOnOff()
        gledState3 = GetRandomOnOff()

        rpio.WritePin(gledPin1, gledState1)
        rpio.WritePin(gledPin2, gledState2)
        rpio.WritePin(gledPin3, gledState3)

        if gledPin1.Read() == rpio.High {
            gledMap["LED1"] = "ON"
        } else {
            gledMap["LED1"] = "OFF"
        }

        if gledPin2.Read() == rpio.High {
            gledMap["LED2"] = "ON"
        } else {
            gledMap["LED2"] = "OFF"
        }

        if gledPin3.Read() == rpio.High {
            gledMap["LED3"] = "ON"
        } else {
            gledMap["LED3"] = "OFF"
        }

        gledJSON, _ := json.Marshal(gledMap)

        fmt.Println("gpio : Calling PubGreenLedStatus ==> ", gledMap)
        PubGreenLedStatus(gledJSON)

        time.Sleep(time.Second * 7)
    }
}

func goDancingBlueLed() {

    fmt.Println("gpio : goDancingBlueLed (go routine)")
    var bledPin1, bledPin2, bledPin3 rpio.Pin
    var bledState1, bledState2, bledState3 rpio.State

    bledPin1 = rpio.Pin(BlueLedPin1)
    bledPin2 = rpio.Pin(BlueLedPin2)
    bledPin3 = rpio.Pin(BlueLedPin3)

    bledPin1.Output()
    bledPin2.Output()
    bledPin3.Output()

    for {

        bledMap := make(map[string]string)

        bledState1 = GetRandomOnOff()
        bledState2 = GetRandomOnOff()
        bledState3 = GetRandomOnOff()

        rpio.WritePin(bledPin1, bledState1)
        rpio.WritePin(bledPin2, bledState2)
        rpio.WritePin(bledPin3, bledState3)

        if bledPin1.Read() == rpio.High {
            bledMap["LED1"] = "ON"
        } else {
            bledMap["LED1"] = "OFF"
        }

        if bledPin2.Read() == rpio.High {
            bledMap["LED2"] = "ON"
        } else {
            bledMap["LED2"] = "OFF"
        }

        if bledPin3.Read() == rpio.High {
            bledMap["LED3"] = "ON"
        } else {
            bledMap["LED3"] = "OFF"
        }

        bledJSON, _ := json.Marshal(bledMap)

        fmt.Println("gpio : Calling PubBlueLedStatus  ==> ", bledMap)
        PubBlueLedStatus(bledJSON)

        time.Sleep(time.Second * 4)
    }
}

func goStillRedLed() {

    fmt.Println("gpio : goStillRedLed (go routine)")
    var rledPin1, rledPin2, rledPin3 rpio.Pin

    rledPin1 = rpio.Pin(RedLedPin1)
    rledPin2 = rpio.Pin(RedLedPin2)
    rledPin3 = rpio.Pin(RedLedPin3)

    for {

        rledMap := make(map[string]string)

        if rledPin1.Read() == rpio.High {
            rledMap["LED1"] = "ON"
        } else {
            rledMap["LED1"] = "OFF"
        }

        if rledPin2.Read() == rpio.High {
            rledMap["LED2"] = "ON"
        } else {
            rledMap["LED2"] = "OFF"
        }

        if rledPin3.Read() == rpio.High {
            rledMap["LED3"] = "ON"
        } else {
            rledMap["LED3"] = "OFF"
        }

        rledJSON, _ := json.Marshal(rledMap)

        fmt.Println("gpio : Calling PubRedLedStatus  ==> ", rledMap)
        PubRedLedStatus(rledJSON)

        time.Sleep(time.Second * 5)
    }
}

func GetRandomOnOff() (ledState rpio.State) {

    var val_1_0 int = (rand.Intn((1-0)+1) + 0)
    ledState = (rpio.State)(val_1_0)
    return
}

