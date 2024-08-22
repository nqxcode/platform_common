package pagination

const DefaultLimit = uint64(30)

// Limit represent offset and limit
type Limit struct {
	Offset uint64
	Limit  uint64
}
