package main

import twilio "github.com/twilio/twilio-go"
import openapi "github.com/twilio/twilio-go/rest/api/v2010"
import "os"
import "fmt"
import "time"

func main() {
    client := twilio.NewRestClient()
    duration, err := time.ParseDuration("61m")
    sendWhen := time.Now().Add(duration)

    params := &openapi.CreateMessageParams{}
    params.SetFrom(os.Getenv("TWILIO_MESSAGING_SERVICE_SID"))
    params.SetTo("+919073423666")
    params.SetBody("Friendly reminder that you have an appointment with us next week.")
    params.SetScheduleType("fixed")
    params.SetSendAt(sendWhen)

    result, err := client.ApiV2010.CreateMessage(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(*result.Sid)
    }
}