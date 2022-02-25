package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pborman/uuid"
	"rsc.io/quote"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}


func main() {
    uuidWithHphen := uuid.NewRandom()
    uuid := strings.Replace(uuidWithHphen.String(), "-", "", -1)
    fmt.Println(uuid)

	fmt.Println(quote.Go())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

