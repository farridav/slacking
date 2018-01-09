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

type Message struct {
	Text     string
	Username string
	Emoji    string
}

func (Message) DoSend(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp

}

func SendMessage(webhook, msg, username, emoji string) (string, string) {
	message := &Message{Text: msg, Username: username, Emoji: emoji}

	payload, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp := message.DoSend(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return resp.Status, string(body)
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

		status, body := SendMessage(*webhook, msg, *username, *emoji)

		fmt.Println("response Status:", status)
		fmt.Println("response Body:", body)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
}
