package v1

import (
	"fmt"
	"net/url"

	sdkerrors "cosmossdk.io/errors"
	base "github.com/dun-io/imhub/types"
	v1base "github.com/dun-io/imhub/types/v1"
)

func (m *Node) GetAddr() base.NodeAddress {
	if m.Address == "" {
		return nil
	}

	addr, err := base.NodeAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return addr
}
func (m *Node) Validate() error {
	if m.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.Address)
	}

	if m.RemoteURL == "" {
		return fmt.Errorf("remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length cannot be greater than %d chars", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return sdkerrors.Wrapf(err, "invalid remote_url %s", m.RemoteURL)
	}
	if remoteURL.Scheme != "https" {
		return fmt.Errorf("remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return fmt.Errorf("remote_url port cannot be empty")
	}

	if m.InactiveAt.IsZero() {
		if !m.Status.Equal(v1base.StatusInactive) {
			return fmt.Errorf("invalid inactive_at %s; expected positive", m.InactiveAt)
		}
	}
	if !m.InactiveAt.IsZero() {
		if !m.Status.Equal(v1base.StatusActive) {
			return fmt.Errorf("invalid inactive_at %s; expected zero", m.InactiveAt)
		}
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return fmt.Errorf("status must be one of [active, inactive]")
	}
	if m.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}
