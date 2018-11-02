package main
import ("fmt";"net/http";"bytes";"io/ioutil")

func main() {
    url:="http://52.36.2.105:8081/TxnWebapp/JsonSelector?LOGIN=Ussd_Bearer1&PASSWORD=MPtc1ToayCkCMZZeHUu0snA3aUaPbSFQ9UzIkNGbVRU=&REQUEST_GATEWAY_CODE=USSD&REQUEST_GATEWAY_TYPE=USSD&requestText="
	//url := "http://www.httpbin.org/post"
	var str = []byte(`{
  "serviceRequest": {
    "COMMAND":{  
         "MSISDN":"7786000014",
         "PROVIDER":"101",
         "MPIN":"1357",
         "TYPE":"CVMPINREQ",
         "LANGUAGE1":"1",
      "PAYID": "12"
    }
  }
  }`)
req, err := http.NewRequest("POST", url, bytes.NewBuffer(str))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

 body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}