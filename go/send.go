package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"flag"
)


func main() {
    /*
    A Script for reading a text file of messages, and sending them to a Slack webhook

    # Setup
    Go to https://<subdomain>.slack.com/services/B8NRSBB0V and setup an integration (grab the webhook_url)

        go run send.go --webhook '<your url from above>'

    */
    webhook := flag.String("webhook", "", "Your slack webhook url, see http://bit.ly/2EapumJ")
    input := flag.String("channel", "../messages.txt", "Path to your message file")
    username := flag.String("username", "Mr Shipit", "The username to use")
    emoji := flag.String("emoji", ":shipit:", "Your icon emoji")

    flag.Parse()

    if *webhook == "" {
        panic("-webhook needed")
    }

	file, err := os.Open(*input)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		var payload = []byte(`{"text": "` + line + `", "username": "` + *username + `", "icon_emoji": "` + *emoji + `"}`)

		req, err := http.NewRequest("POST", *webhook, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
		fmt.Println()
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}
}
