package sync

import "sync"

// WithLock runs while holding lk.
func WithLock(lk sync.Locker, fn func()) {
	lk.Lock()
	defer lk.Unlock() // in case fn panics
	fn()
}
