package main

import (
	"net/http"
	"os"

	"github.com/pa3ng/test-slack-app/pkg/http/router"
)

var (
	port string = "8080"
)

func init() {
	if "" == os.Getenv("VERIFICATION_TOKEN") {
		panic("VERIFICATION_TOKEN is not set!")
	}

	if "" != os.Getenv("PORT") {
		port = os.Getenv("PORT")
	}
}

func main() {
	r := router.NewRouter()

	err := http.ListenAndServe(":"+port, http.Handler(r))
	if err != nil {
		panic(err)
	}
}
