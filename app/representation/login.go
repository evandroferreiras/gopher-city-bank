package representation

// LoginBody struct to illustrate post body
type LoginBody struct {
	Cpf    string `bson:"cpf" json:"cpf" validate:"required"`
	Secret string `bson:"secret" json:"secret" validate:"required"`
}

// LoginResponse struct to illustrate login response
type LoginResponse struct {
	Token string `bson:"token" json:"token"`
}
