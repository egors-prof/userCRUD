package models

import "time"

type User struct {
	Id    int    `db:"id" json:"id"`
	FullName string `db:"full_name" json:"full_name"`
	Username string `db:"username" json:"username"`
	Password string `db:"hash_pass" json :"hash_pass"`
	CreatedAt time.Time `db:"created_at" json :"created_at"`
	UpdatedAt time.Time `db:"updated_at" json :"updated_at"`
}


type SignInRequest struct{
	UserName string `db:"username" json:"username"`
	Password string `db:"hash_pass" json :"hash_pass"`

}
type SignUpRequest struct{
	FullName string `db:"full_name" json:"full_name"`
	Username string `db:"username" json:"username"`
	Password string `db:"hash_pass" json :"hash_pass"`
}
type TokenPairResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	
}
