package types

const (
	// ModuleName defines the module name
	ModuleName = "node"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_node"
)

var (
	ParamsKey = []byte("p_node")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
