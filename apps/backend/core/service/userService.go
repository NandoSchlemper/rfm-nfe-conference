package service

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*model.User, error) {
    user := model.NewUser(name, email)
    return user, s.repo.Save(user)
}


