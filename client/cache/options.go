package cache

// KeyComparator key comparator function
type KeyComparator func(a, b string) bool

// ScanOptions scan options
type ScanOptions struct {
	KeyComparator *KeyComparator
	ScanCount     int
}

// ScanOption scan option
type ScanOption func(*ScanOptions)

// WithKeyComparator key comparator function
func WithKeyComparator(keyComparator KeyComparator) ScanOption {
	return func(options *ScanOptions) {
		options.KeyComparator = &keyComparator
	}
}

// WithScanCount scan limit
func WithScanCount(count int) ScanOption {
	return func(options *ScanOptions) {
		options.ScanCount = count
	}
}
