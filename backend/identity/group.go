package identity

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
