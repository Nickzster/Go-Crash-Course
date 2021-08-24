package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string
}

func index(w http.ResponseWriter, r *http.Request){
	resp := Response {"Hello, world!"}
	fmt.Println(resp);

	js, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "About!")
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server listening on port 3000!")
	http.ListenAndServe(":3000", nil)

}