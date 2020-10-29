package store

import "github.com/adrianosela/guardduty/categorization"

// Store represents the operations our
// storage solution must support
type Store interface {
	GetCategorization() (*categorization.Categorization, error)
}
