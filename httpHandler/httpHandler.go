package main

import (
	"net/http"
	"fmt"
	"encoding/json")

func main() {
    http.HandleFunc("/", display)
    http.ListenAndServe(":8080", nil)
}

func display(w http.ResponseWriter, r *http.Request) {
 var request map[string]interface{}
 decoder := json.NewDecoder(r.Body)
 err := decoder.Decode(&request)
 if err != nil {
    panic(err)
 }
 serviceRequest:=request["serviceRequest"].(map[string]interface{})
//userInfo:=request["userInfo"].(map[string]interface{})
 command:=serviceRequest["COMMAND"].(map[string]interface{})

 for key,value:= range command{
 fmt.Println(key,value.(string))
 }
}