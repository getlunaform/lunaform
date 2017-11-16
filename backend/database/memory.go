package database

import (
	"github.com/zeebox/terraform-server/backend"
)

func NewMemoryDatabase() MemoryDatabase {
	return MemoryDatabase{
		collections: make(map[string]*backend.Collection),
	}
}

type MemoryDatabase struct {
	collections map[string]*backend.Collection
}

func (md MemoryDatabase) Close() {
	md.collections = make(map[string]*backend.Collection)
}

func (md MemoryDatabase) Create(name string) error {
	return md.Insert(name, backend.NewCollection(name, md))
}

func (md MemoryDatabase) Collection(name string) (c *backend.Collection) {
	var ok bool
	if c, ok = md.collections[name]; ok {
		return c
	} else {
		return nil
	}
}

func (md MemoryDatabase) Insert(name string, col backend.Collection) (err error) {
	md.collections[name] = &col
	return
}

/**
 * Record access callbacks
 * The memory database depends on the collection objects to preserve the data.
 */
func (md MemoryDatabase) CreateRecord(action *backend.CreateRecordAction) error {
	return nil
}

func (md MemoryDatabase) ReadRecord(action *backend.ReadRecordAction) error {
	return nil
}

func (md MemoryDatabase) UpdateRecord(action *backend.UpdateRecordAction) error {
	return nil
}

func (md MemoryDatabase) DeleteRecord(action *backend.DeleteRecordAction) error {
	return nil
}
