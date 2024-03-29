package fetcher

import (
	"testing"
)

func TestGetWebsiteStatus(t *testing.T) {
	url := "https://www.google.com"
	ch := make(chan bool)

	go GetWebsiteStatus(url)

	result := <-ch

	if !result {
		t.Errorf("GetWebsiteStatus(%s) = %t; want true", url, result)
	}
}
