package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/davidjss04/indexer/pkg/models"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func LoadMails(url string, nameIndex string, wg *sync.WaitGroup) {
	defer wg.Done()
	b := createJson(url, nameIndex)

	createIndex(nameIndex)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", strings.NewReader(string(*b)))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func createJson(url string, nameIndex string) *[]byte {

	newDirectory := models.NewDirectory(url)
	index := models.NewIndex(nameIndex)

	models.DirecotoryExplorer(&newDirectory, &index)

	var b []byte
	b, err := json.Marshal(index)
	check(err)

	return &b

}

func createIndex(name string) {

	query := `{
		"name": "` + name + `",
		"storage_type": "disk",
		"shard_num": 1,
		"mappings": {
			"properties": {
				"message-id": {
					"type": "text",
					"index": true,
					"store": true,
					"highlightable": true
				},
				"from": {
					"type": "text",
					"index": true,
					"store": true,
					"highlightable": true
				},
				"to": {
					"type": "text",
					"index": true,
					"sortable": true,
					"aggregatable": true
				},
				"content": {
					"type": "text",
					"index": true,
					"sortable": true,
					"aggregatable": true
				}
			}
		}
	}
	`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/index", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}
