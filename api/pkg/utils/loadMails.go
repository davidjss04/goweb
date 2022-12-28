package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	// "os"
	// "path"
	"strings"

	"github.com/davidjss04/api/pkg/models"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CreateJson(url string, nameOput string) *[]byte {

	currentRunDir, err := os.Getwd()
	check(err)

	newDirectory := models.NewDirectory(url)
	index := models.NewIndex("newmaisv6")

	models.DirsExplorer(&newDirectory, false, &index)

	var b []byte
	b, err = json.Marshal(index)
	check(err)

	outputDir := path.Join(currentRunDir, nameOput)
	f, err := os.Create(outputDir)
	check(err)
	defer f.Close()
	_, err = f.Write(b)
	check(err)

	fmt.Println("currentRunDir", currentRunDir)
	fmt.Println("Output saved to", outputDir)
	fmt.Println("Mails loaded")

	return &b
}

func LoadMails(url string, nameOput string) {
	b := CreateJson(url, nameOput)

	// CreateIndex()

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

func CreateIndex() {

	query := `{
		"name": "newmaisv6",
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
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

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
