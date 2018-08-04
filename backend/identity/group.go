package identity

func AdminGroup() *Group {
	return NewGroup("admin")
}

func NewGroup(name string) *Group {
	return &Group{Name: name}
}

// Group of users which can have a role attached to it
type Group struct {
	Name  string
	users []User
}

// Look at https://github.com/mikespook/gorbac for roles

func (g *Group) AddUser(user *User) {
	user.groups = append(user.groups, *g)
	g.users = append(g.users, *user)
}
