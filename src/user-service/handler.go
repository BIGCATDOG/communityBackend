package main

import (
	"context"
	"fmt"
	pb "github.com/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo Repository
	tokenService *TokenService
}

func (h *Handler) Create(c context.Context, user *pb.User, response *pb.Response) error {
	 encryPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	 if err!=nil{
	 	return err
	 }
	 user.Password = string(encryPassword)
     repoErr:=h.repo.Create(user)
     if repoErr!=nil{
     	return err
	 }
	 response.User=user
	 return nil
}

func (h *Handler) Get(c context.Context, user *pb.User, response *pb.Response) error {
      u,err:=h.repo.Get(user.Id)
      if err!=nil{
      	return err
	  }
	  response.User=u
	  return nil
}

func (h *Handler) GetAll(c context.Context, request *pb.Request, response *pb.Response) error {
      users,err := h.repo.GetAll()
      if err!=nil{
      	return err
	  }
	  response.Users = users
	  return nil
}

func (h *Handler) Auth(c context.Context, user *pb.User, token *pb.Token) error {
	userCopy := *user
	userEncode, err := h.repo.GetByEmailAndPassword(user)
	if err != nil {
		return err
	}
	fmt.Println(userEncode)
	fmt.Println(user)
	compareErr:=bcrypt.CompareHashAndPassword([]byte(userEncode.Password),[]byte(userCopy.Password))
	if compareErr!=nil{
		return compareErr
	}
	fmt.Println("verify password")
	tokenRes,tokenErr:=h.tokenService.Encode(userEncode)
	if tokenErr!=nil{
		return tokenErr
	}
	token.Token = tokenRes
	return nil
}

func (h *Handler) ValidateToken(c context.Context, token *pb.Token, token2 *pb.Token) error {
   claims,err:=  h.tokenService.Decode(token.Token)
   if err !=nil || claims.Valid()!=nil {
   	 token2.Valid = false
   	 return err
   }
   token2.Valid=true
   return nil
}

