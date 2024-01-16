package types

const (
	// ModuleName defines the module name
	ModuleName = "example"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_example"
)

var (
	ParamsKey = []byte("p_example")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
