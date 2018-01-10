package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func MockDoSend(req *http.Request) map[string]string {
	body, _ := ioutil.ReadAll(req.Body)
	return map[string]string{
		"status":       "200 OK",
		"body":         "OK",
		"host":         req.Host,
		"method":       req.Method,
		"content_type": req.Header.Get("Content-Type"),
		"data":         string(body),
	}
}

func TestSendMessage(t *testing.T) {
	//
	// Assert that given the input below, we are sending the correct payload
	//
	host := "someplace.com"
	webhook := "http://" + host
	msg := map[string]string{
		"channel":    "#test_channel",
		"icon_emoji": ":test_emoji:",
		"text":       "test message",
		"username":   "test username",
	}

	resp := SendMessage(MockDoSend, webhook, msg["text"], msg["username"], msg["icon_emoji"], msg["channel"])

	expected_bin, err := json.Marshal(msg)
	expected := string(expected_bin)

	if err != nil {
		panic(err)
	}

	// We use the webhook we passed in
	if resp["host"] != host {
		t.Error("failure, expected " + webhook + ", got " + resp["host"])
	}

	// We sent the correct payload to our request handler
	if resp["data"] != expected {
		t.Error("failure, expected " + expected + ", got " + resp["data"])
	}

	// Check correct content type
	if resp["content_type"] != "application/json" {
		t.Error("failure, expected application/json, got" + resp["content_type"])
	}

	// check correct HTTP verb
	if resp["method"] != "POST" {
		t.Error("failure, expected POST, got" + resp["method"])
	}

}
