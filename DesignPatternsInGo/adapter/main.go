package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	UserID    int    `json:"userId" xml:"userId"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*ToDo, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*ToDo, error) {
	return rs.Remote.GetData()
}

type JSONBackend struct{}

func (jb *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*ToDo, error) {
	xmlFile := `
<?xml version="1.0" encoding="UTF-8" ?>
<root>
	<userId>1</userId>
	<id>1</id>
	<title>delectus aut autem</title>
	<completed>false</completed>
</root>
`
	var todo ToDo
	err := xml.Unmarshal([]byte(xmlFile), &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func main() {
	// No adapter
	todo := getRemoteData()
	fmt.Println("TODO without adapter:\t\t", todo.ID, todo.Title)

	// With adapter, using JSON
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tdFromJSON, err := jsonAdapter.CallRemoteService()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("TODO with adapter (JSON):\t", tdFromJSON.ID, tdFromJSON.Title)

	// With adapter, using XML
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tdFromXML, err := xmlAdapter.CallRemoteService()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("TODO with adapter (XML):\t", tdFromXML.ID, tdFromXML.Title)
}

func getRemoteData() *ToDo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal(err)
	}

	return &todo
}
