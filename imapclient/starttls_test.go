package imapclient_test

import (
	"context"
	"crypto/tls"
	"testing"

	"github.com/fpawel/go-imap/imapclient"
)

func TestStartTLS(t *testing.T) {
	conn, server := newMemClientServerPair(t)
	defer conn.Close()
	defer server.Close()

	options := imapclient.Options{
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client, err := imapclient.NewStartTLS(context.Background(), conn, &options)
	if err != nil {
		t.Fatalf("NewStartTLS() = %v", err)
	}
	defer client.Close()

	if err := client.Noop().Wait(context.Background()); err != nil {
		t.Fatalf("Noop().Wait() = %v", err)
	}
}
