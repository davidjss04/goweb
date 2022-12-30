package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/davidjss04/api/pkg/models"
	"github.com/go-chi/chi/v5"
)

type strucResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore int `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			ID        string    `json:"_id"`
			Score     int       `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    struct {
				Content   string `json:"content"`
				From      string `json:"from"`
				MessageID string `json:"message-id"`
				To        string `json:"to"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func GetMailById(w http.ResponseWriter, r *http.Request) {
	mailIdParam := chi.URLParam(r, "mailId")
	bodyReq := getMailOfZincSearch(mailIdParam)

	var struc strucResponse
	var mail models.Mail

	json.Unmarshal([]byte(*bodyReq), &struc)

	mail.MessageId = struc.Hits.Hits[0].Source.MessageID
	mail.From = struc.Hits.Hits[0].Source.From
	mail.To = struc.Hits.Hits[0].Source.To
	mail.Content = struc.Hits.Hits[0].Source.Content

	fmt.Println(struc.Hits.Hits[0].Source.From)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(mail)
}

func SearchMail(w http.ResponseWriter, r *http.Request) {
	phrase := chi.URLParam(r, "phrase")
	bodyReq := searchMailOfZincSearch(phrase)

	var struc strucResponse
	var mails []models.Mail

	json.Unmarshal([]byte(*bodyReq), &struc)

	for _, hit := range struc.Hits.Hits {
		mail := models.Mail{}
		mail.MessageId = hit.Source.MessageID
		mail.From = hit.Source.From
		mail.To = hit.Source.To
		mail.Content = hit.Source.Content
		mails = append(mails, mail)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(mails)
}

func searchMailOfZincSearch(phrase string) *[]byte {

	query := `{
		"search_type": "matchphrase",
		"query": {
				"term": "` + phrase + `"
		},
		"sort_fields": ["-@timestamp"],
		"from": 0,
		"max_results": 1000,
		"_source": [
			"message-id",
			"from",
			"to",
			"content"
		]
	}`

	req, err := http.NewRequest("POST", "http://localhost:4080/api/newmaisv6/_search", strings.NewReader(query))
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

	return &body
}

func getMailOfZincSearch(mailId string) *[]byte {

	query := `{
		"search_type": "alldocuments",
		"query": {
			"message-id": "` + mailId + `"
		},
		"sort_fields": ["-@timestamp"],
		"from": 0,
		"max_results": 1,
		"_source": [
			"message-id",
			"from",
			"to",
			"content"
		]
	}`

	req, err := http.NewRequest("POST", "http://localhost:4080/api/newmaisv6/_search", strings.NewReader(query))
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

	return &body
}
