package database

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// memoryDatabase represents an in memory store for our server.
//   This driver is an ephemeral database stored in RAM, and
//   primarily used for development. When the server shuts down
//   all the state in it is lost. You probably shouldn't use it.
type memoryDatabase struct {
	collections []*Record
}

// NewMemoryDBDriver returns a memory database object
func NewMemoryDBDriver() (Driver, error) {
	return &memoryDatabase{
		collections: []*Record{},
	}, nil
}

// NewMemoryDBDriverWithCollection seeds the memory database with some initialisaiton data
func NewMemoryDBDriverWithCollection(collection []map[string]string) (BufferedDriver, error) {
	rs := make([]*Record, len(collection))

	for i, value := range collection {
		rs[i] = &Record{
			Type:  DBTableRecordType(value["Type"]),
			Key:   value["Key"],
			Value: value["Value"],
		}
	}

	return &memoryDatabase{
		collections: rs,
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
func (md *memoryDatabase) Create(recordType DBTableRecordType, key string, doc interface{}) (err error) {
	if md.exists(recordType, key) {
		return RecordExistsError(fmt.Errorf("%q %q already exists", recordType, key).(RecordExistsError))
	}

	if data, err := md.serialize(doc); err == nil {
		md.collections = append(
			md.collections,
			&Record{
				Key:   key,
				Type:  recordType,
				Value: data,
			},
		)
	}

	return
}

// Read a record from memory
func (md memoryDatabase) Read(recordType DBTableRecordType, key string, i interface{}) (err error) {
	if !md.exists(recordType, key) {
		return RecordDoesNotExistError(fmt.Errorf("%q %q does not exist", recordType, key))
	}

	for _, record := range md.collections {
		if record.Key == key && record.Type == recordType {
			err = md.deserialize(record.Value, i)
			break
		}
	}

	return
}

// List records of a given type from memory
func (md memoryDatabase) List(recordType DBTableRecordType, i interface{}) (err error) {
	elemType := getElemType(i)

	slice := reflect.ValueOf(i).Elem()

	for _, r := range md.collections {
		if r.Type == recordType {
			recType := reflect.New(elemType)
			json.Unmarshal([]byte(r.Value), recType.Interface())
			slice.Set(reflect.Append(slice, recType))
		}
	}

	return
}

// Update a record in memory
func (md *memoryDatabase) Update(recordType DBTableRecordType, key string, doc interface{}) (err error) {
	if !md.exists(recordType, key) {
		return RecordDoesNotExistError(fmt.Errorf("%q %q does not exist", recordType, key))
	}

	var data string
	for i, r := range md.collections {
		if r.Type == recordType && r.Key == key {
			if data, err = md.serialize(doc); err == nil {
				md.collections[i].Value = data
			}
			break
		}
	}

	return
}

// Delete a record from memory
func (md *memoryDatabase) Delete(recordType DBTableRecordType, key string) (err error) {

	for i, r := range md.collections {
		if r.Type == recordType && r.Key == key {
			md.collections = append(
				md.collections[:i],
				md.collections[i+1:]...,
			)
			break
		}
	}

	return
}

func (md memoryDatabase) key(recordType DBTableRecordType, key string) string {
	return fmt.Sprintf("%s %s", recordType, key)
}

func (md memoryDatabase) exists(recordType DBTableRecordType, key string) (ok bool) {
	for _, r := range md.collections {
		if r.Type == recordType && r.Key == key {
			return true
		}
	}
	return false
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
