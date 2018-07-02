package database

import (
	"encoding/json"
	"fmt"
)

// memoryDatabase represents an in memory store for our server.
//   This driver is an ephemeral database stored in RAM, and
//   primarily used for development. When the server shuts down
//   all the state in it is lost. You probably shouldn't use it.
type memoryDatabase struct {
	collections map[string]string
}

// NewMemoryDBDriver returns a memory database object
func NewMemoryDBDriver() (Driver, error) {
	return memoryDatabase{
		collections: make(map[string]string),
	}, nil
}

// NewMemoryDBDriverWithCollection seeds the memory database with some initialisaiton data
func NewMemoryDBDriverWithCollection(collection map[string]string) (BufferedDriver, error) {
	return memoryDatabase{
		collections: collection,
	}, nil
}

// Close doesn't do anything as there is no connection to sever.
func (md memoryDatabase) Close() error {
	return nil
}

// Ping doesn't do anything as the database is inside the app memory space.
func (md memoryDatabase) Ping() error {
	return nil
}

// Create a record in memory
func (md memoryDatabase) Create(recordType, key string, doc interface{}) (err error) {
	if md.exists(recordType, key) {
		return fmt.Errorf("%q %q already exists", recordType, key)
	}

	md.collections[md.key(recordType, key)], err = md.serialize(doc)

	return
}

// Read a record from memory
func (md memoryDatabase) Read(recordType, key string, i interface{}) error {
	if !md.exists(recordType, key) {
		return fmt.Errorf("%q %q does not exist", recordType, key)
	}

	md.deserialize(md.collections[md.key(recordType, key)], i)

	return nil
}

// List records of a given type from memory
func (md memoryDatabase) List(recordType string) (err error) {
	return nil
}

// Update a record in memory
func (md memoryDatabase) Update(recordType, key string, doc interface{}) (err error) {
	if !md.exists(recordType, key) {
		return fmt.Errorf("%q %q does not exist", recordType, key)
	}

	md.collections[md.key(recordType, key)], err = md.serialize(doc)

	return
}

// Delete a record from memory
func (md memoryDatabase) Delete(recordType, key string) error {
	delete(md.collections, md.key(recordType, key))

	return nil
}

func (md memoryDatabase) key(recordType, key string) string {
	return fmt.Sprintf("%s %s", recordType, key)
}

func (md memoryDatabase) exists(recordType, key string) (ok bool) {
	_, ok = md.collections[md.key(recordType, key)]

	return
}

// The following exist because we expect to be able to serialise to/ from
// a pointer to a struct. The original implementation, for example, of
// Read() looked like:
//
// func (md memoryDatabase) Read(recordType, key string, i interface{}) error {
//     if !md.exists(recordType, key) {
//         return fmt.Errorf("%q %q doesn't exist", recordType, key)
//     }
//
//     i = md.collections[md.key(recordType, key)]
//
//     return nil
// }
//
// this meant we were never really updating the i that was passed in (scoping, I assume)
// and so we weren't seeing the data we expected.
// Instead, then, we're going to do json

func (md memoryDatabase) serialize(i interface{}) (string, error) {
	v, err := json.Marshal(i)

	return string(v), err
}

func (md memoryDatabase) deserialize(s string, i interface{}) error {
	return json.Unmarshal([]byte(s), i)
}

// Bytes slice representation of the database
func (md memoryDatabase) Bytes() (b []byte, err error) {
	s, err := md.serialize(md.collections)
	return []byte(s), err
}
