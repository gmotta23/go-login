package models

func GetAllModels() []interface{} {
	var models []interface{}

	models = append(models, &User{})

	return models
}
