package interfaces

type URLData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Result represents the result structure containing name and status information
type Result struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
