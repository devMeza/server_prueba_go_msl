package user_model

import base_models "main/pkg/base/models"

type Struct struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func init() {
	base_models.NewModel[Struct]("User", "users")
}
