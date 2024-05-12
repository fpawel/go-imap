package imapclient

import (
	"context"
	"fmt"

	"github.com/fpawel/go-imap"
	"github.com/fpawel/go-imap/internal/imapwire"
)

// Capability sends a CAPABILITY command.
func (c *Client) Capability() *CapabilityCommand {
	cmd := &CapabilityCommand{}
	c.beginCommand("CAPABILITY", cmd).end()
	return cmd
}

func (c *Client) handleCapability(ctx context.Context) error {
	caps, err := readCapabilities(c.dec)
	if err != nil {
		return err
	}
	c.setCaps(ctx, caps)
	if cmd := findPendingCmdByType[*CapabilityCommand](c); cmd != nil {
		cmd.caps = caps
	}
	return nil
}

// CapabilityCommand is a CAPABILITY command.
type CapabilityCommand struct {
	cmd
	caps imap.CapSet
}

func (cmd *CapabilityCommand) Wait(ctx context.Context) (imap.CapSet, error) {
	err := cmd.cmd.Wait(ctx)
	return cmd.caps, err
}

func readCapabilities(dec *imapwire.Decoder) (imap.CapSet, error) {
	caps := make(imap.CapSet)
	for dec.SP() {
		var name string
		if !dec.ExpectAtom(&name) {
			return caps, fmt.Errorf("in capability-data: %v", dec.Err())
		}
		caps[imap.Cap(name)] = struct{}{}
	}
	return caps, nil
}
