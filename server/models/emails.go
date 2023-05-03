package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type EmailSearchResult struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index     string  `json:"_index"`
			Type      string  `json:"_type"`
			ID        string  `json:"_id"`
			Score     float64 `json:"_score"`
			Timestamp string  `json:"@timestamp"`
			Source    Email   `json:"_source"`
			Highlight struct {
				Content []string `json:"content"`
			} `json:"highlight"`
		} `json:"hits"`
	} `json:"hits"`
}

type Email struct {
	Id        string   `json:"id"`
	From      string   `json:"from"`
	To        string   `json:"to"`
	Content   string   `json:"content"`
	Subject   string   `json:"subject"`
	Date      string   `json:"date"`
	Highlight []string `json:"highlight"`
}

type ApiResponse struct {
	Took   int     `json:"took"`
	Emails []Email `json:"emails"`
}

var (
	apiEndpoint = "http://ec2-18-231-3-43.sa-east-1.compute.amazonaws.com:4080/api/emails/_search"
	httpClient  = &http.Client{}
)

func GetEmails(s string) (ApiResponse, error) {
	requestBody := map[string]interface{}{
		"search_type": "matchphrase",
		"query": map[string]interface{}{
			"term":  s,
			"field": "content",
		},
		"from":        0,
		"max_results": 200,
		"_source":     []string{},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<strong>"},
			"post_tags": []string{"</strong>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{
					"pre_tags":  []string{},
					"post_tags": []string{},
				},
				"content": map[string]interface{}{
					"pre_tags":  []string{},
					"post_tags": []string{},
				},
			},
		},
	}

	var apiResponse ApiResponse

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Println(err)
		return apiResponse, err
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return apiResponse, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "password")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return apiResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return apiResponse, err
	}
	resp.Body.Close()

	response := EmailSearchResult{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		return apiResponse, err
	}

	apiResponse = ApiResponse{
		Took:   response.Hits.Total.Value,
		Emails: convertToEmails(response),
	}

	return apiResponse, nil
}

func convertToEmails(response EmailSearchResult) []Email {
	emails := []Email{}

	for _, hit := range response.Hits.Hits {
		email := Email{
			Id:        hit.ID,
			From:      hit.Source.From,
			To:        hit.Source.To,
			Subject:   hit.Source.Subject,
			Content:   hit.Source.Content,
			Date:      hit.Source.Date,
			Highlight: hit.Highlight.Content,
		}
		emails = append(emails, email)
	}

	return emails
}
