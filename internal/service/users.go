package service

import (

	"CSR/internal/errs"
	"CSR/internal/models"
	"CSR/utils"
	"errors"
	"log"
	"os"

	"github.com/rs/zerolog"
)
func(s*Service)CreateNewUser(userRequest models.SignUpRequest)error{
	hashed,err:=utils.GenerateHashPassword(userRequest.Password)
	log.Println(hashed)
	if err!=nil{
		if errors.Is(err,errs.ErrHashing){
			return errs.ErrHashing
		}
	}
	userRequest.Password=hashed
	err=s.repository.CreateNewUser(userRequest)
	if err!=nil{
		return err
	}
	return nil 
}


func(s*Service)Authenticate(userRequest models.SignInRequest)(int ,error){
	logger:=zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msg("authentication . . . ")
	//1.username exists 
	user,err:=s.repository.GetUserByUsername(userRequest.UserName) 
	// log.Println("user",user)
	// log.Println("err",err)
	if err!=nil{
		log.Println("error in block 1")
		return 0,errs.ErrIncorrectUsernameOrPassword
	}

	//2.check password hashes 
	userReqHashPass,err:=utils.GenerateHashPassword(userRequest.Password)
	if err!=nil{
		log.Println("error in block 2")
		return 0,err
	}
	log.Println("hash_pass",string(userReqHashPass),"\n",user.Password)
	if !(userReqHashPass==user.Password){
		log.Println("password are not equal")
		return 0,errs.ErrIncorrectUsernameOrPassword
	}
	
	log.Println(user.Id,err)
	return user.Id,nil
}