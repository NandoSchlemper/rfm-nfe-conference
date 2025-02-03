package repository

type UserRepository interface {
    Save(user *model.User) error
    FindById(id string) (*model.User, error)
}
