package main

import (
	"context"
	pb "github.com/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo Repository
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
	userEncode, err := h.repo.GetByEmailAndPassword(user)
	if err != nil {
		return err
	}
	compareErr:=bcrypt.CompareHashAndPassword([]byte(userEncode.Password),[]byte(user.Password))
	if compareErr!=nil{
		return compareErr
	}
	token.Token = "`x_2nam"
	return nil
}

func (h *Handler) ValidateToken(c context.Context, token *pb.Token, token2 *pb.Token) error {
     return nil
}

