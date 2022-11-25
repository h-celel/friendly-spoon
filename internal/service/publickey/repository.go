package publickey

import (
	"sync"
)

type Repository struct {
	sync.RWMutex
	publicKey []byte
}

func (r *Repository) Get() []byte {
	r.RLock()
	defer r.RUnlock()
	return r.publicKey
}

func (r *Repository) Put(publicKey []byte) {
	r.Lock()
	defer r.Unlock()
	r.publicKey = publicKey
}
