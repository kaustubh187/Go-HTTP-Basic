package main

import (
    "fmt"
    "net/http"
)



func hello(w http.ResponseWriter, req *http.Request){
    
    fmt.Println("Request Received")
    fmt.Fprint(w,"Hello!!")


}
func main(){

    http.HandleFunc("/",hello)
    fmt.Println("Server Running")
    http.ListenAndServe(":8080",nil)
}