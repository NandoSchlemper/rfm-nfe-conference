package model

type User struct {
    ID string
    Name string
    Email string
}

func NewUser(name string, email string) *User {
    return &User{
        ID: generateUUID(),
        Name: name,
        Email: email,
    }
}
