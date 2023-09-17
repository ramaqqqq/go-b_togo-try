package usecase

import "bookingoto-try/features/user"

type UserUseCaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUseCase(userRepo user.UserRepo) user.UserUseCase {
	return &UserUseCaseImpl{userRepo}
}

func (r *UserUseCaseImpl) ReadAllUser() ([]map[string]interface{}, error) {
	return r.userRepo.ReadAllUser()
}

func (r *UserUseCaseImpl) ReadSingleId(cstId string) (map[string]interface{}, error) {
	return r.userRepo.ReadSingleId(cstId)
}
