package model

type City struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (*City) TableName() string {
	return "cities"
}
