package backend

import "encoding/json"

// Database represents a data storage for terraform stacks
type Database interface {
	Collection(name string) *Collection
	Insert(name string, col Collection) error
	Create(name string) error
	CreateRecord(action *CreateRecordAction) error
	ReadRecord(action *ReadRecordAction) error
	UpdateRecord(action *UpdateRecordAction) error
	DeleteRecord(action *DeleteRecordAction) error
	Close()
}

func NewCollection(name string, db Database) Collection {
	return Collection{
		name: name,
		db:   db,
		kv:   make(map[string][]byte),
	}
}

type Collection struct {
	name string
	db   Database
	kv   map[string][]byte
}

type CreateRecordAction struct {
	Collection Collection
	Id         string
	Payload    []byte
}

type ReadRecordAction struct {
	Collection Collection
	Id         string
	Result     *[]byte
}

type UpdateRecordAction struct {
	Collection Collection
	Id         string
	Payload    []byte
}
type DeleteRecordAction struct {
	Collection Collection
	Id         string
}

func (mc Collection) Name() (name string) {
	return mc.name
}

func (mc Collection) Count() (num int) {
	return len(mc.kv)
}

func (mc Collection) Create(id string, doc interface{}) (err error) {
	var b []byte
	if b, err = json.Marshal(doc); err == nil {
		if err = mc.db.CreateRecord(&CreateRecordAction{
			Collection: mc,
			Id:         id,
			Payload:    b,
		}); err == nil {
			mc.kv[id] = b
		}
	}
	return
}

func (mc Collection) Read(id string, result interface{}) (err error) {
	var b []byte
	var exists bool

	if b, exists = mc.kv[id]; !exists {
		err = mc.db.ReadRecord(&ReadRecordAction{
			Collection: mc,
			Id:         id,
			Result:     &b,
		});
		if err != nil {
			return
		}
	}
	return json.Unmarshal(b, result)
}

func (mc Collection) Update(id string, doc interface{}) (err error) {
	var b []byte
	if b, err = json.Marshal(doc); err == nil {
		if err = mc.db.UpdateRecord(&UpdateRecordAction{
			Collection: mc,
			Id:         id,
			Payload:    b,
		}); err != nil {
			return
		}
	}
	return mc.Read(id, doc)
}

func (mc Collection) Delete(id string) (err error) {
	if err = mc.db.DeleteRecord(&DeleteRecordAction{
		Collection: mc,
		Id:         id,
	}); err == nil {
		delete(mc.kv, id)
	}
	return
}

// Result set.
//type Result interface {
//	// Count returns the number of items that match the set conditions.
//	Count() (int64, error)
//	// One fetches the first result within the result set.
//	One(interface{}) error
//	// All fetches all results within the result set.
//	All(interface{}) error
//	// Limit defines the maximum number of results in this set.
//	Limit(int64) Result
//	// Skip ignores first *n* results.
//	Skip(int64) Result
//	// Sort results by given fields.
//	Sort(...string) Result
//	// Cursor executes query and returns cursor capable of going over all the results.
//	Cursor() (Cursor, error)
//}
//
//// Cursor API
//type Cursor interface {
//	// Close closes the cursor, preventing further enumeration.
//	Close() error
//	// Next reads the next result.
//	Next(result interface{}) (bool, error)
//}
