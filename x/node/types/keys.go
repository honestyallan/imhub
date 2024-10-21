package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"
	base "github.com/dun-io/imhub/types"
	"time"
)

const (
	// ModuleName defines the module name
	ModuleName = "node"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
)

var (
	ParamsKey                  = []byte("p_node")
	NodeKeyPrefix              = []byte{0x10}
	ActiveNodeKeyPrefix        = append(NodeKeyPrefix, 0x01)
	InactiveNodeKeyPrefix      = append(NodeKeyPrefix, 0x02)
	NodeForPlanKeyPrefix       = []byte{0x11}
	NodeForInactiveAtKeyPrefix = []byte{0x12}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func ActiveNodeKey(addr base.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}
func InactiveNodeKey(addr base.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}
func NodeForInactiveAtKey(at time.Time, addr base.NodeAddress) []byte {
	return append(GetNodeForInactiveAtKeyPrefix(at), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}
func GetNodeForInactiveAtKeyPrefix(at time.Time) []byte {
	return append(NodeForInactiveAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}
