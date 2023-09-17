package user

type UserUseCase interface {
	ReadAllUser() ([]map[string]interface{}, error)
	ReadSingleId(cstId string) (map[string]interface{}, error)
}
