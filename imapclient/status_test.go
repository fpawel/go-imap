package imapclient_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/fpawel/go-imap"
)

func TestStatus(t *testing.T) {
	client, server := newClientServerPair(t, imap.ConnStateAuthenticated)
	defer client.Close()
	defer server.Close()

	options := imap.StatusOptions{
		NumMessages: true,
		NumUnseen:   true,
	}
	data, err := client.Status("INBOX", &options).Wait(context.Background())
	if err != nil {
		t.Fatalf("Status() = %v", err)
	}

	wantNumMessages := uint32(1)
	wantNumUnseen := uint32(1)
	want := &imap.StatusData{
		Mailbox:     "INBOX",
		NumMessages: &wantNumMessages,
		NumUnseen:   &wantNumUnseen,
	}
	if !reflect.DeepEqual(data, want) {
		t.Errorf("Status() = %#v but want %#v", data, want)
	}
}
