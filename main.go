package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jzarzeczny/go-checker/controllers"
	"github.com/jzarzeczny/go-checker/jsonParser"
)

func main() {

	urlList, err := jsonParser.Parser()

	if err != nil {
		fmt.Println("Failed to parse JSON ", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { controllers.GetWebsiteStatus(w, r, urlList) })

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
