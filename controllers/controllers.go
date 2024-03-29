package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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

	resultMap := make(map[string]bool)
	var wg sync.WaitGroup

	wg.Add((len(urlList)))

	for _, data := range urlList {
		go func(data interfaces.URLData) {
			defer wg.Done()

			success := fetcher.GetWebsiteStatus(data.URL)
			fmt.Printf("\n Trying to get value of %s and it's %v \n", data.URL, success)
			resultCh <- interfaces.Result{Name: data.Name, Status: success}

		}(data)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		resultMap[result.Name] = result.Status
	}

	json, err := json.Marshal(resultMap)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshalling JSON: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(json)

}
