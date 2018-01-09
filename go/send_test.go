package main

import "testing"

func TestSendMessage(t *testing.T) {
	status, body := SendMessage("test", "test", "test", "test")

	if status != "200 OK" {
		t.Error("failure, got " + status)
	}

	if body != "OK" {
		t.Error("failure")
	}

}
