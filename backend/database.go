package backend

import "encoding/json"

// Record is an untyped, schemaless record type which all
// record types embed and implement
type Record map[string]interface{}

// Key returns a record's Key
func (r Record) Key() string {
	return r["Key"].(string)
}

// Type returns a record's type
func (r Record) Type() string {
	return r["Type"].(string)
}

// Driver represents a low level storage serialiser/ deserialiser
// This is wrapped in the Database
type Driver interface {
	Create(recordType, key string, doc interface{}) error
	Read(recordType, key string) (interface{}, error)
	Update(recordType, key string, doc interface{}) error
	Delete(recordType, key string) error

	Ping() error
	Close() error
}

// Database stores data for terraform server
type Database struct {
	driver Driver
}
