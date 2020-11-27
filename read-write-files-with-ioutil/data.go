package main

// Data struct
type Data struct {
	Title       string
	Link        string
	Description string
	Content     string
}

// DataContainer struct
type DataContainer struct {
	Name string
	Data []*Data
}
