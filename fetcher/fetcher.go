package fetcher

import (
	"fmt"
	"net/http"
)

func GetWebsiteStatus(url string, ch chan <- bool){
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		ch <- false
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("URL %s returned status code %d \n", url, res.StatusCode)
		ch <- false
		return
	}

	ch <- true

}