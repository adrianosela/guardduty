package cache

// Cache represents the operations our
// caching functionality must support
type Cache interface {
	GetSeverity(finding string) (string, bool)
}
