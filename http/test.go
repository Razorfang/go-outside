package main

import (
        "fmt"
        "log"
        "net/http"
)

var counter int = 0

func generateResponseString() string {
	return fmt.Sprintf("The value of the counter on the server is '%d'\n", counter)
}

func handleGet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, generateResponseString())
}

func handleInc(writer http.ResponseWriter, request *http.Request) {
	counter += 1
}

func main() {
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/inc", handleInc)
        log.Fatal(http.ListenAndServe(":1234", nil))
}
