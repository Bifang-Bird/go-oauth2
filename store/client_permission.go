package store

import (
	"context"
	"errors"
	"sync"

	oauth2 "github.com/Bifang-Bird/goOauth2"
)

// NewClientStore create client store
func NewClientPermissionStore() *ClientPermissionStore {
	return &ClientPermissionStore{
		data: make(map[string][]*oauth2.ClientPermissionInfo),
	}
}

// ClientStore client information store
type ClientPermissionStore struct {
	sync.RWMutex
	data map[string][]*oauth2.ClientPermissionInfo
}

// GetByID according to the ID for the client information
func (cs *ClientPermissionStore) GetByID(ctx context.Context, id string) ([]*oauth2.ClientPermissionInfo, error) {
	cs.RLock()
	defer cs.RUnlock()

	if c, ok := cs.data[id]; ok {
		return c, nil
	}
	return nil, errors.New("not found")
}

// Set set client information
func (cs *ClientPermissionStore) Set(id string, cli []*oauth2.ClientPermissionInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()

	cs.data[id] = cli
	return
}
