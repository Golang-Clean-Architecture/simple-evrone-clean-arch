package entity

type Todo struct {
	Name   string `json:"name" bson:"name"`
	Status string `json:"status" bson:"status"`
}
