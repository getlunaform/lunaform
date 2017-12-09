

# database
`import "github.com/zeebox/terraform-server/backend/database"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Database](#Database)
  * [func NewDatabaseWithDriver(driver Driver) Database](#NewDatabaseWithDriver)
* [type Driver](#Driver)
* [type JSONDatabase](#JSONDatabase)
  * [func NewJSONDatabase(dbFile fileClient) (jdb JSONDatabase, err error)](#NewJSONDatabase)
  * [func (jdb JSONDatabase) Close() (err error)](#JSONDatabase.Close)
  * [func (jdb JSONDatabase) Create(recordType, key string, doc interface{}) error](#JSONDatabase.Create)
  * [func (jdb JSONDatabase) Delete(recordType, key string) error](#JSONDatabase.Delete)
  * [func (jdb JSONDatabase) Ping() error](#JSONDatabase.Ping)
  * [func (jdb JSONDatabase) Read(recordType, key string, i interface{}) (err error)](#JSONDatabase.Read)
  * [func (jdb JSONDatabase) Update(recordType, key string, doc interface{}) (err error)](#JSONDatabase.Update)
* [type MemoryDatabase](#MemoryDatabase)
  * [func NewMemoryDatabase() (MemoryDatabase, error)](#NewMemoryDatabase)
  * [func (md MemoryDatabase) Bytes() []byte](#MemoryDatabase.Bytes)
  * [func (md MemoryDatabase) Close() error](#MemoryDatabase.Close)
  * [func (md MemoryDatabase) Create(recordType, key string, doc interface{}) error](#MemoryDatabase.Create)
  * [func (md MemoryDatabase) Delete(recordType, key string) error](#MemoryDatabase.Delete)
  * [func (md MemoryDatabase) Ping() error](#MemoryDatabase.Ping)
  * [func (md MemoryDatabase) Read(recordType, key string, i interface{}) error](#MemoryDatabase.Read)
  * [func (md MemoryDatabase) Update(recordType, key string, doc interface{}) error](#MemoryDatabase.Update)
* [type Record](#Record)
  * [func (r Record) Key() string](#Record.Key)
  * [func (r Record) Type() string](#Record.Type)
* [type RedisDatabase](#RedisDatabase)
  * [func NewRedisDatabase(namespace, address, password string, database int) (r RedisDatabase, err error)](#NewRedisDatabase)
  * [func (r RedisDatabase) Close() error](#RedisDatabase.Close)
  * [func (r RedisDatabase) Create(recordType, key string, doc interface{}) error](#RedisDatabase.Create)
  * [func (r RedisDatabase) Delete(recordType, key string) error](#RedisDatabase.Delete)
  * [func (r RedisDatabase) Ping() error](#RedisDatabase.Ping)
  * [func (r RedisDatabase) Read(recordType, key string, i interface{}) (err error)](#RedisDatabase.Read)
  * [func (r RedisDatabase) Update(recordType, key string, doc interface{}) (err error)](#RedisDatabase.Update)


#### <a name="pkg-files">Package files</a>
[database.go](/src/github.com/zeebox/terraform-server/backend/database/database.go) [dynamodb.go](/src/github.com/zeebox/terraform-server/backend/database/dynamodb.go) [json.go](/src/github.com/zeebox/terraform-server/backend/database/json.go) [memory.go](/src/github.com/zeebox/terraform-server/backend/database/memory.go) [redis.go](/src/github.com/zeebox/terraform-server/backend/database/redis.go) 






## <a name="Database">type</a> [Database](/src/target/database.go?s=734:773#L30)
``` go
type Database struct {
    // contains filtered or unexported fields
}
```
Database stores data for terraform server







### <a name="NewDatabaseWithDriver">func</a> [NewDatabaseWithDriver](/src/target/database.go?s=835:885#L35)
``` go
func NewDatabaseWithDriver(driver Driver) Database
```
NewDatabaseWithDriver creates a new Database struct with





## <a name="Driver">type</a> [Driver](/src/target/database.go?s=433:687#L19)
``` go
type Driver interface {
    Create(recordType, key string, doc interface{}) error
    Read(recordType, key string, i interface{}) error
    Update(recordType, key string, doc interface{}) error
    Delete(recordType, key string) error

    Ping() error
    Close() error
}
```
Driver represents a low level storage serialiser/ deserialiser
This is wrapped in the Database










## <a name="JSONDatabase">type</a> [JSONDatabase](/src/target/json.go?s=266:332#L14)
``` go
type JSONDatabase struct {
    // contains filtered or unexported fields
}
```
JSONDatabase stores data on disk in json files.
This database is better than MemoryDatabase (but honestly,
pretty much everything is), but still not a good solution for
live/production.







### <a name="NewJSONDatabase">func</a> [NewJSONDatabase](/src/target/json.go?s=546:615#L31)
``` go
func NewJSONDatabase(dbFile fileClient) (jdb JSONDatabase, err error)
```
NewJSONDatabase returns a json database object





### <a name="JSONDatabase.Close">func</a> (JSONDatabase) [Close](/src/target/json.go?s=836:879#L44)
``` go
func (jdb JSONDatabase) Close() (err error)
```
Close the file pointer




### <a name="JSONDatabase.Create">func</a> (JSONDatabase) [Create](/src/target/json.go?s=1005:1082#L51)
``` go
func (jdb JSONDatabase) Create(recordType, key string, doc interface{}) error
```
Create a record in the JSON file on disk




### <a name="JSONDatabase.Delete">func</a> (JSONDatabase) [Delete](/src/target/json.go?s=1176:1236#L56)
``` go
func (jdb JSONDatabase) Delete(recordType, key string) error
```
Delete a record in the JSON file on disk




### <a name="JSONDatabase.Ping">func</a> (JSONDatabase) [Ping](/src/target/json.go?s=1337:1373#L61)
``` go
func (jdb JSONDatabase) Ping() error
```
Ping mock. Implementing a no-op for the json file db




### <a name="JSONDatabase.Read">func</a> (JSONDatabase) [Read](/src/target/json.go?s=1445:1524#L66)
``` go
func (jdb JSONDatabase) Read(recordType, key string, i interface{}) (err error)
```
Read a record from the JSON file on disk




### <a name="JSONDatabase.Update">func</a> (JSONDatabase) [Update](/src/target/json.go?s=1615:1698#L71)
``` go
func (jdb JSONDatabase) Update(recordType, key string, doc interface{}) (err error)
```
Updated a record in the JSON file on disk




## <a name="MemoryDatabase">type</a> [MemoryDatabase](/src/target/memory.go?s=302:363#L12)
``` go
type MemoryDatabase struct {
    // contains filtered or unexported fields
}
```
MemoryDatabase represents an in memory store for our server.
This driver is an ephemeral database stored in RAM, and
primarily used for development. When the server shuts down
all the state in it is lost. You probably shouldn't use it.







### <a name="NewMemoryDatabase">func</a> [NewMemoryDatabase](/src/target/memory.go?s=419:467#L17)
``` go
func NewMemoryDatabase() (MemoryDatabase, error)
```
NewMemoryDatabase returns a memory database object





### <a name="MemoryDatabase.Bytes">func</a> (MemoryDatabase) [Bytes](/src/target/memory.go?s=2942:2981#L112)
``` go
func (md MemoryDatabase) Bytes() []byte
```
Bytes slice representation of the database




### <a name="MemoryDatabase.Close">func</a> (MemoryDatabase) [Close](/src/target/memory.go?s=610:648#L24)
``` go
func (md MemoryDatabase) Close() error
```
Close doesn't do anything as there is no connection to sever.




### <a name="MemoryDatabase.Create">func</a> (MemoryDatabase) [Create](/src/target/memory.go?s=826:904#L34)
``` go
func (md MemoryDatabase) Create(recordType, key string, doc interface{}) error
```
Create a record in memory




### <a name="MemoryDatabase.Delete">func</a> (MemoryDatabase) [Delete](/src/target/memory.go?s=1678:1739#L67)
``` go
func (md MemoryDatabase) Delete(recordType, key string) error
```
Delete a record from memory




### <a name="MemoryDatabase.Ping">func</a> (MemoryDatabase) [Ping](/src/target/memory.go?s=742:779#L29)
``` go
func (md MemoryDatabase) Ping() error
```
Ping doesn't do anything as the database is inside the app memory space.




### <a name="MemoryDatabase.Read">func</a> (MemoryDatabase) [Read](/src/target/memory.go?s=1111:1185#L45)
``` go
func (md MemoryDatabase) Read(recordType, key string, i interface{}) error
```
Read a record from memory




### <a name="MemoryDatabase.Update">func</a> (MemoryDatabase) [Update](/src/target/memory.go?s=1391:1469#L56)
``` go
func (md MemoryDatabase) Update(recordType, key string, doc interface{}) error
```
Update a record in memory




## <a name="Record">type</a> [Record](/src/target/database.go?s=112:146#L5)
``` go
type Record map[string]interface{}
```
Record is an untyped, schemaless record type which all
record types embed and implement










### <a name="Record.Key">func</a> (Record) [Key](/src/target/database.go?s=178:206#L8)
``` go
func (r Record) Key() string
```
Key returns a record's Key




### <a name="Record.Type">func</a> (Record) [Type](/src/target/database.go?s=270:299#L13)
``` go
func (r Record) Type() string
```
Type returns a record's type




## <a name="RedisDatabase">type</a> [RedisDatabase](/src/target/redis.go?s=251:321#L18)
``` go
type RedisDatabase struct {
    // contains filtered or unexported fields
}
```
RedisDatabase represents a redis store for our server
It holds state and docs and all kinds of awesome sauce







### <a name="NewRedisDatabase">func</a> [NewRedisDatabase](/src/target/redis.go?s=601:702#L33)
``` go
func NewRedisDatabase(namespace, address, password string, database int) (r RedisDatabase, err error)
```
NewRedisDatabase returns a redis database object





### <a name="RedisDatabase.Close">func</a> (RedisDatabase) [Close](/src/target/redis.go?s=1160:1196#L55)
``` go
func (r RedisDatabase) Close() error
```
Close will close a connection to the redis database




### <a name="RedisDatabase.Create">func</a> (RedisDatabase) [Create](/src/target/redis.go?s=1259:1335#L60)
``` go
func (r RedisDatabase) Create(recordType, key string, doc interface{}) error
```
Create a record within redis




### <a name="RedisDatabase.Delete">func</a> (RedisDatabase) [Delete](/src/target/redis.go?s=2133:2192#L104)
``` go
func (r RedisDatabase) Delete(recordType, key string) error
```
Delete a record from redis




### <a name="RedisDatabase.Ping">func</a> (RedisDatabase) [Ping](/src/target/redis.go?s=1015:1050#L48)
``` go
func (r RedisDatabase) Ping() error
```
Ping checks whether redis is up and usable
it returns an error to determine this; a nil value means the redis database is usable




### <a name="RedisDatabase.Read">func</a> (RedisDatabase) [Read](/src/target/redis.go?s=1513:1591#L71)
``` go
func (r RedisDatabase) Read(recordType, key string, i interface{}) (err error)
```
Read a record from redis




### <a name="RedisDatabase.Update">func</a> (RedisDatabase) [Update](/src/target/redis.go?s=1879:1961#L93)
``` go
func (r RedisDatabase) Update(recordType, key string, doc interface{}) (err error)
```
Update a redis record








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
