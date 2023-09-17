package user

type UserRepo interface {
	ReadAllUser() ([]map[string]interface{}, error)
	ReadSingleId(cstId string) (map[string]interface{}, error)
}
