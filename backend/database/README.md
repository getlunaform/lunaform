

# database
`import "github.com/zeebox/terraform-server/backend/database"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type JSONCollection](#JSONCollection)
* [type MemoryDatabase](#MemoryDatabase)
  * [func NewMemoryDatabase() (MemoryDatabase, error)](#NewMemoryDatabase)
  * [func (md MemoryDatabase) Close() error](#MemoryDatabase.Close)
  * [func (md *MemoryDatabase) Create(recordType, key string, doc interface{}) error](#MemoryDatabase.Create)
  * [func (md MemoryDatabase) Delete(recordType, key string) error](#MemoryDatabase.Delete)
  * [func (md MemoryDatabase) Ping() error](#MemoryDatabase.Ping)
  * [func (md MemoryDatabase) Read(recordType, key string, i interface{}) error](#MemoryDatabase.Read)
  * [func (md MemoryDatabase) Update(recordType, key string, doc interface{}) error](#MemoryDatabase.Update)
* [type RedisDatabase](#RedisDatabase)
  * [func NewRedisDatabase(namespace, address, password string, database int) (r RedisDatabase, err error)](#NewRedisDatabase)
  * [func (r RedisDatabase) Close() error](#RedisDatabase.Close)
  * [func (r RedisDatabase) Create(recordType, key string, doc interface{}) error](#RedisDatabase.Create)
  * [func (r RedisDatabase) Delete(recordType, key string) error](#RedisDatabase.Delete)
  * [func (r RedisDatabase) Ping() error](#RedisDatabase.Ping)
  * [func (r RedisDatabase) Read(recordType, key string, i interface{}) (err error)](#RedisDatabase.Read)
  * [func (r RedisDatabase) Update(recordType, key string, doc interface{}) (err error)](#RedisDatabase.Update)


#### <a name="pkg-files">Package files</a>
[dynamodb.go](/src/github.com/zeebox/terraform-server/backend/database/dynamodb.go) [json.go](/src/github.com/zeebox/terraform-server/backend/database/json.go) [memory.go](/src/github.com/zeebox/terraform-server/backend/database/memory.go) [redis.go](/src/github.com/zeebox/terraform-server/backend/database/redis.go) 






## <a name="JSONCollection">type</a> [JSONCollection](/src/target/json.go?s=67:97#L4)
``` go
type JSONCollection struct {
}
```
JSONFileDB stores data on disk in json files.










## <a name="MemoryDatabase">type</a> [MemoryDatabase](/src/target/memory.go?s=180:241#L14)
``` go
type MemoryDatabase struct {
    // contains filtered or unexported fields
}
```






### <a name="NewMemoryDatabase">func</a> [NewMemoryDatabase](/src/target/memory.go?s=54:102#L8)
``` go
func NewMemoryDatabase() (MemoryDatabase, error)
```




### <a name="MemoryDatabase.Close">func</a> (MemoryDatabase) [Close](/src/target/memory.go?s=243:281#L18)
``` go
func (md MemoryDatabase) Close() error
```



### <a name="MemoryDatabase.Create">func</a> (\*MemoryDatabase) [Create](/src/target/memory.go?s=354:433#L26)
``` go
func (md *MemoryDatabase) Create(recordType, key string, doc interface{}) error
```



### <a name="MemoryDatabase.Delete">func</a> (MemoryDatabase) [Delete](/src/target/memory.go?s=1118:1179#L56)
``` go
func (md MemoryDatabase) Delete(recordType, key string) error
```



### <a name="MemoryDatabase.Ping">func</a> (MemoryDatabase) [Ping](/src/target/memory.go?s=299:336#L22)
``` go
func (md MemoryDatabase) Ping() error
```



### <a name="MemoryDatabase.Read">func</a> (MemoryDatabase) [Read](/src/target/memory.go?s=611:685#L36)
``` go
func (md MemoryDatabase) Read(recordType, key string, i interface{}) error
```



### <a name="MemoryDatabase.Update">func</a> (MemoryDatabase) [Update](/src/target/memory.go?s=862:940#L46)
``` go
func (md MemoryDatabase) Update(recordType, key string, doc interface{}) error
```



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
