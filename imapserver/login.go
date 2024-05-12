package imapserver

import (
	"github.com/fpawel/go-imap"
	"github.com/fpawel/go-imap/internal/imapwire"
)

func (c *Conn) handleLogin(tag string, dec *imapwire.Decoder) error {
	var username, password string
	if !dec.ExpectSP() || !dec.ExpectAString(&username) || !dec.ExpectSP() || !dec.ExpectAString(&password) || !dec.ExpectCRLF() {
		return dec.Err()
	}
	if err := c.checkState(imap.ConnStateNotAuthenticated); err != nil {
		return err
	}
	if !c.canAuth() {
		return &imap.Error{
			Type: imap.StatusResponseTypeNo,
			Code: imap.ResponseCodePrivacyRequired,
			Text: "TLS is required to authenticate",
		}
	}
	if err := c.session.Login(username, password); err != nil {
		return err
	}
	c.state = imap.ConnStateAuthenticated
	return c.writeCapabilityStatus(tag, imap.StatusResponseTypeOK, "Logged in")
}
