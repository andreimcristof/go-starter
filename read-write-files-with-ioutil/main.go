package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	result := getData()

	for _, data := range result.Data {
		data.Content = getContent(data.Link)
	}

	mO, err := json.Marshal(result)
	handleError(err)

	writeFile(mO)
}

func writeFile(content []byte) {
	var writePermission = 0755
	wErr := ioutil.WriteFile(
		"./changed_data.json",
		content,
		os.FileMode(writePermission))

	handleError(wErr)
}

func getContent(path string) string {
	output, err := ioutil.ReadFile(path)
	handleError(err)
	return string(output)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getData() DataContainer {
	output, err := ioutil.ReadFile("./data/data.json")
	handleError(err)

	var container DataContainer
	uErr := json.Unmarshal(output, &container)
	handleError(uErr)

	return container
}
