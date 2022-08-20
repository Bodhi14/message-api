/* package main

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
)

func main() {

    // Set variable values
	gatewayURL := "http://web.springedge.com"
	apikey := "AUTH_API_KEY"
	sender := "SEDEMO"
	urlStr := gatewayURL + "/api/web/send/"
    
    //Params
    v := url.Values{}
    v.Set("to","9073423666")
    v.Set("apikey",apikey)
    v.Set("sender",sender)
    v.Set("message","Hello, this is my first message sent using Golang!!!")

    rb := *strings.NewReader(v.Encode())

    client := &http.Client()

    req, _ := http.NewRequest("POST", urlStr, &rb)

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    // Making Request
    resp, _ := client.Do(req)

    // Print Response
    fmt.Println(resp.Status)
}
*/


/*
package main
import (
  "fmt"
  textmagic "github.com/textmagic/textmagic-rest-go"
)
func main() {
  client := textmagic.NewClient("<USERNAME>", "<APIV2_KEY>")
  params := map[string]string{
    "phones": "99900000",
    "text"  : "Hello from TextMagic Go",
  }
  message, err := client.CreateMessage(params)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(message.Id)
  }
}
*/

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

type VonageResponse struct {
    Messages []struct {
        Status    string `json:"status"`
        ErrorText string `json:"error-text"`
    } `json:"messages"`
}

func main() {
    apiKey := "b8a4d95d"
    apiSecret := "FuRIhuR2ZjLQVJsv"
    apiPath := "https://rest.nexmo.com/sms/json"

    from := "Vonage APIs"
    to := "+919830014541"

    message := "Real Madrid >>> FC Barcelona"

    // Request body
    body := url.Values{
        "from": {from},
        "to": {to},
        "text": {message},
        "api_key": {apiKey},
        "api_secret": {apiSecret},
    }

    // Create a HTTP post request
    r, err := http.NewRequest("POST", apiPath, strings.NewReader(body.Encode()))
    if err != nil {
        fmt.Println(err)
        return
    }

    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    res, err := client.Do(r)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    resp := &VonageResponse{}
    derr := json.NewDecoder(res.Body).Decode(resp)
    if derr != nil {
        fmt.Println(err)
        return
    }

    if len(resp.Messages) <= 0 {
        fmt.Println("Vonage error: Internal Error")
        return
    }

    // A status of zero indicates success; a non-zero value means something went wrong.
    if resp.Messages[0].Status != "0" {
        fmt.Errorf("Vonage error: %v (status: %v)", resp.Messages[0].ErrorText, resp.Messages[0].Status)
        return
    }

    fmt.Println("SMS sent successfully.")
}








