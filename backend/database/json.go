package database

import (
	"encoding/json"
	"io"
	"os"
	"time"
	"io/ioutil"
	"sync"
)

// jsonDatabase stores data on disk in json files.
// This database is better than MemoryDatabase (but honestly,
// pretty much everything is), but still not a good solution for
// live/production.
type jsonDatabase struct {
	file  string
	db    BufferedDriver
	mutex *sync.Mutex
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
func NewJSONDBDriver(dbFile string) (d Driver, err error) {
	jdb := jsonDatabase{
		file:  dbFile,
		mutex: &sync.Mutex{},
	}

	var (
		dat []byte
		mdb BufferedDriver
	)
	if dat, err = ioutil.ReadFile(dbFile); err != nil {
		return nil, err
	}

	c := make([]map[string]string, 0)
	if len(dat) > 0 {
		if err = json.Unmarshal(dat, &c); err != nil {
			return
		}
	}

	if mdb, err = NewMemoryDBDriverWithCollection(c); err != nil {
		return nil, err
	}
	jdb.db = mdb

	go jdb.doEvery(2*time.Second, jdb.Flush)

	return jdb, err
}

// Close the file pointer
func (jdb jsonDatabase) Close() (err error) {
	jdb.lock()
	defer jdb.unlock()

	return jdb.Flush(nil)
}

func (jdb jsonDatabase) Flush(t *time.Time) (err error) {
	jdb.lock()
	defer jdb.unlock()

	var b []byte
	if b, err = jdb.db.Bytes(); err != nil {
		return
	}
	if err = ioutil.WriteFile(jdb.file, b, 0644); err != nil {
		return
	}

	return
}

// Create a record in the JSON file on disk
func (jdb jsonDatabase) Create(recordType DBTableRecordType, key string, doc interface{}) error {
	jdb.lock()
	defer jdb.unlock()

	return jdb.db.Create(recordType, key, doc)
}

// Delete a record in the JSON file on disk
func (jdb jsonDatabase) Delete(recordType DBTableRecordType, key string) error {
	jdb.lock()
	defer jdb.unlock()

	return jdb.db.Delete(recordType, key)
}

// Ping mock. Implementing a no-op for the json file db
func (jdb jsonDatabase) Ping() error {
	jdb.lock()
	defer jdb.unlock()

	return jdb.db.Ping()
}

// Read a record from the JSON file on disk
func (jdb jsonDatabase) Read(recordType DBTableRecordType, key string, i interface{}) (err error) {
	jdb.lock()
	defer jdb.unlock()

	return jdb.db.Read(recordType, key, i)
}

// List all records of a given type from the JSON file on disk
func (jdb jsonDatabase) List(recordType DBTableRecordType, i interface{}) (err error) {
	jdb.lock()
	defer jdb.unlock()

	return jdb.db.List(recordType, i)
}

// Update a record in the JSON file on disk
func (jdb jsonDatabase) Update(recordType DBTableRecordType, key string, doc interface{}) (err error) {
	jdb.lock()
	defer jdb.unlock()

	return jdb.db.Update(recordType, key, doc)
}

func (jdb jsonDatabase) doEvery(d time.Duration, f func(*time.Time) error) {
	for x := range time.Tick(d) {
		if err := f(&x); err != nil {
			panic(err)
		}
	}
}

func (jdb jsonDatabase) lock() {
	jdb.mutex.Lock()
}

func (jdb jsonDatabase) unlock() {
	jdb.mutex.Unlock()
}
