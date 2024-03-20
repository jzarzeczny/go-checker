package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jzarzeczny/go-checker/data"
	"github.com/jzarzeczny/go-checker/fetcher"
	"github.com/jzarzeczny/go-checker/interfaces"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		resultCh := make(chan interfaces.Result)

		for _, data := range data.UrlDataList {
			go func(data interfaces.URLData) {
				ch := make(chan bool)
				defer close(ch)
				go fetcher.GetWebsiteStatus(data.URL, ch)
				success := <-ch
				resultCh <- interfaces.Result{Name: data.Name, Status: success}

			}(data)
		}

		var results []interfaces.Result
		for range data.UrlDataList {
			result := <-resultCh
			results = append(results, result)
		}

		json, err := json.Marshal(results)

		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshalling JSON: %v", err), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		w.Write(json)
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
