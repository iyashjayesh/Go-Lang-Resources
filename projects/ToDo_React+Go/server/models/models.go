package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoList struct{
	ID	primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`	
	Task	string	`json:"task,omitempty"`
	Status	bool	`json:"status,omitempty"`
}

func main(){

	

}