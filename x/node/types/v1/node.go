package v1

import (
	sdkerrors "cosmossdk.io/errors"
	"fmt"
	base "github.com/dun-io/imhub/types"
	v1base "github.com/dun-io/imhub/types/v1"
)

func (m *Node) GetAddress() base.NodeAddress {
	if m.AccAddr == "" {
		return nil
	}

	addr, err := base.NodeAddressFromBech32(m.AccAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Node) Validate() error {
	if m.AccAddr == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.AccAddr); err != nil {
		return sdkerrors.Wrapf(err, "invalid address %s", m.AccAddr)
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
