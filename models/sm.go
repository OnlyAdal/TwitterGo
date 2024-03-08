package models

type Secret struct {
	Host     string `json:"host"` //alt + 96 comillas alrrevez
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSign  string `json:"jwtsign"`
	Database string `json:"database"`
}
