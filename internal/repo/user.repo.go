package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}
func (ur *UserRepo) GetIoUser() string {
	return "Io User"
}
