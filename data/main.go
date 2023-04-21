package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

func main() {
	err := filepath.Walk("./enron_mail_20110402/maildir/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			mime, err := enmime.ReadEnvelope(bytes.NewReader(content))
			if err != nil {
				return err
			}
			emailData := EmailData{
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
			}
			jsonData, err := json.Marshal(emailData)
			if err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
			sendToZincSearch(jsonData)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func sendToZincSearch(jsonData []byte) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_doc", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "password")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
