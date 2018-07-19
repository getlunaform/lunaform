package database

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"fmt"
	"time"
)

// jsonDatabase stores data on disk in json files.
// This database is better than MemoryDatabase (but honestly,
// pretty much everything is), but still not a good solution for
// live/production.
type jsonDatabase struct {
	file fileClient
	db   BufferedDriver
}

type fileClient interface {
	//io.Closer
	io.Reader
	//io.ReaderAt
	io.Seeker
	io.WriterAt
	io.Writer
	Stat() (os.FileInfo, error)
	Truncate(size int64) error
}

// NewJSONDBDriver returns a json database object
func NewJSONDBDriver(dbFile fileClient) (Driver, error) {
	jdb := jsonDatabase{
		file: dbFile,
	}

	b := new(bytes.Buffer)
	b.ReadFrom(jdb.file)

	c := make(map[string]string)
	err := json.Unmarshal(b.Bytes(), &c)
	mdb, err := NewMemoryDBDriverWithCollection(c)
	if err != nil {
		return nil, err
	}
	jdb.db = mdb

	go jdb.doEvery(2*time.Second, jdb.Flush)

	return jdb, err
}

// Close the file pointer
func (jdb jsonDatabase) Close() (err error) {
	return jdb.Flush(nil)
}

func (jdb jsonDatabase) Flush(t *time.Time) (err error) {
	fmt.Sprint("Closing db connection and flushing to file")
	var b []byte
	if err = jdb.file.Truncate(0); err != nil {
		return
	}
	if b, err = jdb.db.Bytes(); err != nil {
		return err
	}
	_, err = jdb.file.WriteAt(b, 0)
	return
}

// Create a record in the JSON file on disk
func (jdb jsonDatabase) Create(recordType, key string, doc interface{}) error {
	return jdb.db.Create(recordType, key, doc)
}

// Delete a record in the JSON file on disk
func (jdb jsonDatabase) Delete(recordType, key string) error {
	return jdb.db.Delete(recordType, key)
}

// Ping mock. Implementing a no-op for the json file db
func (jdb jsonDatabase) Ping() error {
	return jdb.db.Ping()
}

// Read a record from the JSON file on disk
func (jdb jsonDatabase) Read(recordType, key string, i interface{}) (err error) {
	return jdb.db.Read(recordType, key, i)
}

// List all records of a given type from the JSON file on disk
func (jdb jsonDatabase) List(recordType string, i interface{}) (err error) {
	return jdb.db.List(recordType, i)
}

// Update a record in the JSON file on disk
func (jdb jsonDatabase) Update(recordType, key string, doc interface{}) (err error) {
	return jdb.db.Update(recordType, key, doc)
}

func (jdb jsonDatabase) doEvery(d time.Duration, f func(*time.Time) error) {
	for x := range time.Tick(d) {
		if err := f(&x); err != nil {
			panic(err)
		}
	}
}
