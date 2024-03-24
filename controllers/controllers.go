package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jzarzeczny/go-checker/fetcher"
	"github.com/jzarzeczny/go-checker/interfaces"
	"github.com/jzarzeczny/go-checker/validator"
)

func GetWebsiteStatus(w http.ResponseWriter, r *http.Request, urlList []interfaces.URLData, token string) {

	err := validator.ValidateToken(r, token)

	if err != nil {
		http.Error(w, "Token is invalid", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	resultCh := make(chan interfaces.Result)

	for _, data := range urlList {
		go func(data interfaces.URLData) {
			ch := make(chan bool)
			defer close(ch)
			go fetcher.GetWebsiteStatus(data.URL, ch)
			success := <-ch
			resultCh <- interfaces.Result{Name: data.Name, Status: success}

		}(data)
	}

	var results []interfaces.Result
	for range urlList {
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

}
