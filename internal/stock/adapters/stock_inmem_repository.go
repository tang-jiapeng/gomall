package adapters

import (
	"context"
	"github.com/0X10CC/gorder/common/genproto/orderpb"
	domain "github.com/0X10CC/gorder/stock/domain/stock"
	"sync"
)

type MemoryStoreRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

var stub = map[string]*orderpb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub_item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
}

func NewMemoryStoreRepository() *MemoryStoreRepository {
	return &MemoryStoreRepository{
		lock:  &sync.RWMutex{},
		store: make(map[string]*orderpb.Item),
	}
}

func (m MemoryStoreRepository) GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var (
		res     []*orderpb.Item
		missing []string
	)
	for _, id := range ids {
		if item, exist := m.store[id]; exist {
			res = append(res, item)
		} else {
			missing = append(missing, id)
		}
	}
	if len(res) == len(ids) {
		return res, nil
	}
	return res, domain.NotFoundError{Missing: missing}
}
