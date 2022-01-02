package user

//Interface is interface to describe how to handle an user
type Interface interface {
	FetchUserByName(name string) (User, error)
	FetchUserById(id string) (User, error)
	UpdateUser(user User) error
	CreateUser(user User) error
}

// User should be representation of model in database
type User struct {
	uid         string
	name        string
	email       string
	phoneNumber string
}
