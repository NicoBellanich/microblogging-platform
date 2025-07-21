package domain

type User struct {
	Name         string
	Following    []*User
	Publications MessageList
}

func NewUser(userName string) *User {
	return &User{
		Name: userName,
	}
}

func (user *User) AddPublication(msg Message) {
	user.Publications.AddMessage(&msg)
}

func (user *User) AddFollowing(newFollowing *User) {
	if newFollowing != nil {
		user.Following = append(user.Following, newFollowing)
	}
}

func (user *User) GetAllPublications() MessageList {
	return user.Publications
}
