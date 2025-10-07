package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)
func GenerateHashPassword(password string)(string,error){
	hasher:=sha256.New()
	_,err:=hasher.Write([]byte(password))
	if err!=nil{
		return "",err
	}
	hashed:=hasher.Sum(nil)
	log.Println(hashed)
	str:=hex.EncodeToString(hashed)
	return str,err

}