

# identity
`import "github.com/zeebox/terraform-server/backend/identity"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type ApiKey](#ApiKey)
* [type Group](#Group)
* [type MemoryIdentityProvider](#MemoryIdentityProvider)
  * [func NewMemoryIdentityProvider() MemoryIdentityProvider](#NewMemoryIdentityProvider)
  * [func (mip MemoryIdentityProvider) ChangePassword(user User, password string) (err error)](#MemoryIdentityProvider.ChangePassword)
  * [func (mip MemoryIdentityProvider) ConsumeEndpoint(payload []byte) (err error)](#MemoryIdentityProvider.ConsumeEndpoint)
  * [func (mip MemoryIdentityProvider) CreateUser(username string, password string) (user User, err error)](#MemoryIdentityProvider.CreateUser)
  * [func (mip MemoryIdentityProvider) IsEditable() (editable bool)](#MemoryIdentityProvider.IsEditable)
  * [func (mip MemoryIdentityProvider) IsFederated() (federated bool)](#MemoryIdentityProvider.IsFederated)
  * [func (mip MemoryIdentityProvider) LoginUser(user User, password string) (loggedin bool)](#MemoryIdentityProvider.LoginUser)
  * [func (mip MemoryIdentityProvider) ReadUser(username string) (user User, err error)](#MemoryIdentityProvider.ReadUser)
  * [func (mip MemoryIdentityProvider) UpdateUser(user User) (err error)](#MemoryIdentityProvider.UpdateUser)
* [type Provider](#Provider)
  * [func NewDatabaseIdentityProvider(db database.Database) (idp Provider, err error)](#NewDatabaseIdentityProvider)
* [type SSHKey](#SSHKey)
* [type User](#User)
  * [func (u *User) ChangePassword(password string) (err error)](#User.ChangePassword)
  * [func (u *User) LoggedIn() bool](#User.LoggedIn)
  * [func (u *User) Login(password string) bool](#User.Login)
  * [func (u *User) Logout()](#User.Logout)


#### <a name="pkg-files">Package files</a>
[auth_group.go](/src/github.com/zeebox/terraform-server/backend/identity/auth_group.go) [auth_user.go](/src/github.com/zeebox/terraform-server/backend/identity/auth_user.go) [database.go](/src/github.com/zeebox/terraform-server/backend/identity/database.go) [memory.go](/src/github.com/zeebox/terraform-server/backend/identity/memory.go) [provider.go](/src/github.com/zeebox/terraform-server/backend/identity/provider.go) 






## <a name="ApiKey">type</a> [ApiKey](/src/target/auth_user.go?s=808:986#L48)
``` go
type ApiKey struct {
    Value                string
    DateCreated          time.Time
    DateExpired          time.Time
    ValidationPeriod     time.Duration
    AutomaticallyExpired bool
}
```









## <a name="Group">type</a> [Group](/src/target/auth_group.go?s=18:37#L3)
``` go
type Group struct{}
```









## <a name="MemoryIdentityProvider">type</a> [MemoryIdentityProvider](/src/target/memory.go?s=393:454#L18)
``` go
type MemoryIdentityProvider struct {
    // contains filtered or unexported fields
}
```
Memory IdentityProvider will store user details in RAM. Once this
struct is released, all data is lost. This is really only used for
development and will probably be deprecated in time.







### <a name="NewMemoryIdentityProvider">func</a> [NewMemoryIdentityProvider](/src/target/memory.go?s=70:125#L9)
``` go
func NewMemoryIdentityProvider() MemoryIdentityProvider
```




### <a name="MemoryIdentityProvider.ChangePassword">func</a> (MemoryIdentityProvider) [ChangePassword](/src/target/memory.go?s=1762:1850#L78)
``` go
func (mip MemoryIdentityProvider) ChangePassword(user User, password string) (err error)
```



### <a name="MemoryIdentityProvider.ConsumeEndpoint">func</a> (MemoryIdentityProvider) [ConsumeEndpoint](/src/target/memory.go?s=621:698#L30)
``` go
func (mip MemoryIdentityProvider) ConsumeEndpoint(payload []byte) (err error)
```



### <a name="MemoryIdentityProvider.CreateUser">func</a> (MemoryIdentityProvider) [CreateUser](/src/target/memory.go?s=712:813#L34)
``` go
func (mip MemoryIdentityProvider) CreateUser(username string, password string) (user User, err error)
```



### <a name="MemoryIdentityProvider.IsEditable">func</a> (MemoryIdentityProvider) [IsEditable](/src/target/memory.go?s=456:518#L22)
``` go
func (mip MemoryIdentityProvider) IsEditable() (editable bool)
```



### <a name="MemoryIdentityProvider.IsFederated">func</a> (MemoryIdentityProvider) [IsFederated](/src/target/memory.go?s=537:601#L26)
``` go
func (mip MemoryIdentityProvider) IsFederated() (federated bool)
```



### <a name="MemoryIdentityProvider.LoginUser">func</a> (MemoryIdentityProvider) [LoginUser](/src/target/memory.go?s=1583:1670#L74)
``` go
func (mip MemoryIdentityProvider) LoginUser(user User, password string) (loggedin bool)
```



### <a name="MemoryIdentityProvider.ReadUser">func</a> (MemoryIdentityProvider) [ReadUser](/src/target/memory.go?s=1072:1154#L49)
``` go
func (mip MemoryIdentityProvider) ReadUser(username string) (user User, err error)
```



### <a name="MemoryIdentityProvider.UpdateUser">func</a> (MemoryIdentityProvider) [UpdateUser](/src/target/memory.go?s=1502:1569#L70)
``` go
func (mip MemoryIdentityProvider) UpdateUser(user User) (err error)
```



## <a name="Provider">type</a> [Provider](/src/target/provider.go?s=18:350#L3)
``` go
type Provider interface {
    IsEditable() bool
    IsFederated() bool

    CreateUser(username string, password string) (User, error)
    ReadUser(username string) (User, error)
    UpdateUser(user User) error

    LoginUser(user User, password string) bool
    ChangePassword(user User, password string) error

    ConsumeEndpoint(payload []byte) error
}
```






### <a name="NewDatabaseIdentityProvider">func</a> [NewDatabaseIdentityProvider](/src/target/database.go?s=95:175#L8)
``` go
func NewDatabaseIdentityProvider(db database.Database) (idp Provider, err error)
```




## <a name="SSHKey">type</a> [SSHKey](/src/target/auth_user.go?s=666:806#L40)
``` go
type SSHKey struct {
    Public          []byte
    Private         []byte
    PublicPath      string
    PrivatePath     string
    ServerGenerated bool
}
```









## <a name="User">type</a> [User](/src/target/auth_user.go?s=45:222#L8)
``` go
type User struct {
    IsEditable bool
    Username   string
    Password   string
    ApiKeys    []*ApiKey
    SSHKeys    []*SSHKey
    Idp        Provider
    // contains filtered or unexported fields
}
```









### <a name="User.ChangePassword">func</a> (\*User) [ChangePassword](/src/target/auth_user.go?s=441:499#L32)
``` go
func (u *User) ChangePassword(password string) (err error)
```



### <a name="User.LoggedIn">func</a> (\*User) [LoggedIn](/src/target/auth_user.go?s=386:416#L28)
``` go
func (u *User) LoggedIn() bool
```



### <a name="User.Login">func</a> (\*User) [Login](/src/target/auth_user.go?s=224:266#L19)
``` go
func (u *User) Login(password string) bool
```



### <a name="User.Logout">func</a> (\*User) [Logout](/src/target/auth_user.go?s=337:360#L24)
``` go
func (u *User) Logout()
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
