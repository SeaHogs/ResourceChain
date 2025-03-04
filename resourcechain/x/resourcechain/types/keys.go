package types

const (
	// ModuleName defines the module name
	ModuleName = "resourcechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_resourcechain"
)

var (
	ParamsKey = []byte("p_resourcechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
