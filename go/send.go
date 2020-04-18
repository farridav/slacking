package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Our swappable Http Handler type
type HttpHandler func(req *http.Request) map[string]string

// Our production Http Handler
func MakeRequest(req *http.Request) map[string]string {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return map[string]string{
		"status": resp.Status,
		"body":   string(body),
	}
}

// Our Message sending utlity
func SendMessage(sender HttpHandler, webhook string, msg string, username string, emoji string, channel string) map[string]string {
	message := map[string]string{
		"text":       msg,
		"username":   username,
		"icon_emoji": emoji,
	}

	if channel != "" {
		message["channel"] = channel
	}

	payload, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	return sender(req)
}

func main() {
	/*
	   A Script for reading a text file of messages, and sending them to a Slack webhook

	   # Setup

	   Go to https://<subdomain>.slack.com/services/B8NRSBB0V and setup an integration (grab the webhook_url)

	       go run send.go --webhook '<your url from above>'

	*/
	webhook := flag.String("webhook", "", "Your slack webhook url, see http://bit.ly/2EapumJ")
	input := flag.String("input", "../messages.txt", "Path to your message file")
	username := flag.String("username", "Mr Shipit", "The username to use")
	emoji := flag.String("emoji", ":shipit:", "Your icon emoji")
	channel := flag.String("channel", "", "Override channel with this channel")

	flag.Parse()

	if *webhook == "" {
		panic("--webhook needed")
	}

	file, err := os.Open(*input)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var msg string
	for {
		msg, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		response := SendMessage(MakeRequest, *webhook, msg, *username, *emoji, *channel)

		fmt.Println("response Status:", response["status"])
		fmt.Println("response Body:", response["body"])
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
}
