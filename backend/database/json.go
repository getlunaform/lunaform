package database

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type fileClient interface {
	Write([]byte) (n int, err error)
}

// JSONDatabase represents a store on disk encoded in JSON format.
// This database is better than MemoryDatabase (but honestly,
// pretty much everything is), but still not a good solution for
// live/production.
type JSONDatabase struct {
	localStore   MemoryDatabase
	file         fileClient
	flushPending chan bool
	flushPeriod  int
	flushMux     sync.Mutex
}

// NewJSONDatabase returns a driver to a JSON file on disk
func NewJSONDatabase(path string) (j JSONDatabase, err error) {
	var file fileClient
	if file, err = os.Open(path); os.IsNotExist(err) {
		if file, err = os.Create(path); err != nil {
			// If we can't create the file and it doesn't exist, bail.
			return
		}
	}

	j = JSONDatabase{
		file:         file,
		flushPending: make(chan bool),
		flushPeriod:  10,
		flushMux:     sync.Mutex{},
	}
	j.localStore, err = NewMemoryDatabase()
	go j.startFlusher()

	return
}

// Close will not do anything as we write directly to disk
// and keep a copy in memory
func (jd JSONDatabase) Close() (err error) {
	return
}

// Ping does nothing as the file is on a local disk.
func (jd JSONDatabase) Ping() error {
	return nil
}

// Create a record in the json file on disk
func (jd JSONDatabase) Create(recordType, key string, doc interface{}) (err error) {
	jd.localStore.Create(recordType, key, doc)
	go jd.flush()

	return
}

// Read a record from a json file on disk
func (jd JSONDatabase) Read(recordType, key string, i interface{}) (err error) {
	return
}

// Update a record to a json file on disk
func (jd JSONDatabase) Update(recordType, key string, doc interface{}) (err error) {
	return
}

// Delete a record from the json file on disk
func (jd JSONDatabase) Delete(recordType, key string) (err error) {
	return
}

func (jd JSONDatabase) flush() {
	jd.flushMux.Lock()
	jd.flushPending <- true
	jd.flushMux.Unlock()
}

func (jd JSONDatabase) startFlusher() {
	for {

		// Make the flushPending reset operation atomic.
		jd.flushMux.Lock()
		doFlush := <-jd.flushPending
		jd.flushPending <- false
		jd.flushMux.Unlock()

		if doFlush {
			payload, err := json.Marshal(jd.localStore.collections)
			if err != nil {
				panic(err)
			}
			jd.file.Write(payload)
		}
		time.Sleep(time.Duration(jd.flushPeriod * time.Millisecond))
	}
}
