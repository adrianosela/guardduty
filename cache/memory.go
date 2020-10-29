package cache

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/adrianosela/guardduty/categorization"
	"github.com/adrianosela/guardduty/store"
)

// MemoryCache is an in-memory
// implementation of the Cache interface
type MemoryCache struct {
	// inherit read/write lock behavior
	sync.RWMutex

	categorization categorization.Categorization
	db             store.Store
	ttl            time.Duration
}

// NewMemoryCache is the MemoryCache constructor
func NewMemoryCache(db store.Store, ttl time.Duration) Cache {
	c := &MemoryCache{db: db, ttl: ttl}
	go c.worker()
	return c
}

// GetSeverity determines the severity of a given finding
func (m *MemoryCache) GetSeverity(finding string) (string, bool) {
	m.RLock()
	defer m.RUnlock()
	severity, ok := m.categorization.Mapping[finding]
	return severity, ok
}

// worker checks for data freshness every ttl millisecond
// and updates the data if a newer version is available
func (m *MemoryCache) worker() {
	ticker := time.NewTicker(m.ttl)
	m.refresh() // note: ignoring error

	// ending condition
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			m.refresh() // note: ignoring error
		case <-sig:
			return
		}
	}
}

// refresh refreshes the data in the cache
// if the data from upstream source is fresher
func (m *MemoryCache) refresh() error {
	fresh, err := m.db.GetCategorization()
	if err != nil {
		return err
	}

	m.RLock()
	stale := fresh.Version != m.categorization.Version
	m.RUnlock()

	if stale {
		m.Lock()
		m.categorization = *fresh
		m.Unlock()
	}

	return nil
}
