package telegram

type User struct {
	Id string
}

func (u *User) Recipient() string {
	return u.Id
}
