package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jhillyerd/enmime"
)

type EmailData struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Subject         string `json:"subject"`
	Content         string `json:"content"`
	MessageID       string `json:"message_id"`
	Date            string `json:"date"`
	ContentType     string `json:"content_type"`
	MimeVersion     string `json:"mime_version"`
	ContentEncoding string `json:"content_transfer_encoding"`
	XFrom           string `json:"x_from"`
	XTo             string `json:"x_to"`
	Xcc             string `json:"x_cc"`
	Xbcc            string `json:"x_bcc"`
	XFolder         string `json:"x_folder"`
	XOrigin         string `json:"x_origin"`
	XFilename       string `json:"x_filename"`
}

type BulkData struct {
	Index   string      `json:"index"`
	Records []EmailData `json:"records"`
}

type PropertyDetail struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}

type Mapping struct {
	Properties map[string]PropertyDetail `json:"properties"`
}

type IndexerData struct {
	Name         string  `json:"name"`
	StorageType  string  `json:"storage_type"`
	ShardNum     int     `json:"shard_num"`
	MappingField Mapping `json:"mappings"`
}

func main() {
	log.Println("Starting indexer...")
	indexerData, err := createIndexerFromJsonFile("./index.json")
	if err != nil {
		log.Fatal(err)
	}

	sent := createIndexOnZincSearch(indexerData)
	if sent != "ok" {
		log.Fatal("Failed to send indexer to ZincSearch")
	}

	// log.Println("Start indexing...")
	// startTime := time.Now()
	//
	// var records []EmailData
	// var m sync.Mutex
	// var wg sync.WaitGroup
	//
	// err := filepath.Walk("./maildir/", func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !info.IsDir() {
	// 		wg.Add(1)
	// 		go func(p string) {
	// 			defer wg.Done()
	// 			emailData, err := processFile(p)
	// 			if err != nil {
	// 				log.Println(err)
	// 				return
	// 			}
	// 			m.Lock()
	// 			records = append(records, emailData)
	// 			m.Unlock()
	// 		}(path)
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// wg.Wait()
	//
	// sendBulkToZincSearch(records)
	//
	// duration := time.Since(startTime)
	// log.Printf("Finished indexing. Time taken: %.2f seconds", duration.Seconds())
}

func processFile(path string) (EmailData, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return EmailData{}, err
	}
	mime, err := enmime.ReadEnvelope(bytes.NewReader(content))
	if err != nil {
		return EmailData{}, err
	}
	return EmailData{
		From:            mime.GetHeader("From"),
		To:              mime.GetHeader("To"),
		Subject:         mime.GetHeader("Subject"),
		Content:         mime.Text,
		MessageID:       mime.GetHeader("Message-ID"),
		Date:            mime.GetHeader("Date"),
		ContentType:     mime.GetHeader("Content-Type"),
		MimeVersion:     mime.GetHeader("Mime-Version"),
		ContentEncoding: mime.GetHeader("Content-Transfer-Encoding"),
		XFrom:           mime.GetHeader("X-From"),
		XTo:             mime.GetHeader("X-To"),
		Xcc:             mime.GetHeader("X-cc"),
		Xbcc:            mime.GetHeader("X-bcc"),
		XFolder:         mime.GetHeader("X-Folder"),
		XOrigin:         mime.GetHeader("X-Origin"),
		XFilename:       mime.GetHeader("X-Filename"),
	}, nil
}

func sendBulkToZincSearch(records []EmailData) {
	bulkData := BulkData{
		Index:   "emails",
		Records: records,
	}

	jsonData, err := json.Marshal(bulkData)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", bytes.NewReader(jsonData))
	if err != nil {
		log.Println(err)
		return
	}
	req.SetBasicAuth("admin", "password")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func createIndexerFromJsonFile(filepath string) (IndexerData, error) {
	var indexerData IndexerData

	file, err := os.Open(filepath)
	if err != nil {
		return indexerData, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&indexerData)
	if err != nil {
		return indexerData, err
	}

	return indexerData, nil
}

func createIndexOnZincSearch(indexerData IndexerData) string {
	jsonData, err := json.Marshal(indexerData)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://yourhost.com/create_indexer", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("failed to create indexer, status code: %d", resp.StatusCode)
	}
	log.Println("Indexer created successfully")

	return "ok"
}
