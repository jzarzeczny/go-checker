package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jzarzeczny/go-checker/controllers"
	"github.com/jzarzeczny/go-checker/dotEnvParser"
	"github.com/jzarzeczny/go-checker/jsonParser"
)

func main() {
	envVariables, err := dotEnvParser.ParseEnvVariables()

	if err != nil {
		fmt.Println("Error during parsing of env variables", err)
	}

	urlList, err := jsonParser.Parser()

	if err != nil {
		fmt.Println("Failed to parse JSON ", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetWebsiteStatus(w, r, urlList, envVariables["TOKEN"])
	})

	port := envVariables["PORT"]

	if port == "" {
		port = ":8080"
	}

	fmt.Printf("Server listening on port %s", port)

	http.ListenAndServe(port, nil)
}
