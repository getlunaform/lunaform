package database

type DBTableRecordType string

func (dbtype DBTableRecordType) String() string {
	return string(dbtype)
}

// Record is an untyped, schemaless record type which all
// record types embed and implement
type Record struct {
	Key   string
	Type  DBTableRecordType
	Value string
}

// Driver represents a low level storage serialiser/ deserialiser
// This is wrapped in the Database
type Driver interface {
	Create(recordType DBTableRecordType, key string, doc interface{}) error
	Read(recordType DBTableRecordType, key string, i interface{}) error
	List(recordType DBTableRecordType, i interface{}) error
	Update(recordType DBTableRecordType, key string, doc interface{}) error
	Delete(recordType DBTableRecordType, key string) error

	Ping() error
	Close() error
}

// BufferedDriver represents a low level storage which exposes a Bytes method to allow
// serialization to disk
type BufferedDriver interface {
	Driver
	Bytes() ([]byte, error)
}

// Database stores data for terraform server
type Database struct {
	driver Driver
}

// NewDatabaseWithDriver creates a new Database struct with
func NewDatabaseWithDriver(driver Driver) Database {
	return Database{
		driver: driver,
	}
}

func (db *Database) Create(recordType DBTableRecordType, key string, doc interface{}) error {
	return db.driver.Create(recordType, key, doc)
}

func (db *Database) Read(recordType DBTableRecordType, key string, i interface{}) error {
	return db.driver.Read(recordType, key, i)
}

func (db *Database) List(recordType DBTableRecordType, i interface{}) (err error) {
	return db.driver.List(recordType, i)
}

func (db *Database) Update(recordType DBTableRecordType, key string, doc interface{}) error {
	return db.driver.Update(recordType, key, doc)
}

func (db *Database) Delete(recordType DBTableRecordType, key string) error {
	return db.driver.Delete(recordType, key)
}

func (db *Database) Ping() error {
	return db.driver.Ping()
}

func (db *Database) Close() error {
	return db.driver.Close()
}

type RecordExistsError error
type RecordDoesNotExistError error
