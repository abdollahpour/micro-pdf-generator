package main

import (
	"fmt"
	"github.com/abdollahpour/micro-pdf-generator/internal/config"
	"log"
	"net/http"

	. "github.com/abdollahpour/micro-pdf-generator/internal/handler"
)

func main() {
	http.HandleFunc("/", PdfHandler)
	log.Println(fmt.Sprintf("Listen on port %d", config.Config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
