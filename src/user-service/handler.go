package main

import pb "user-service/proto/user"

type Handler struct {
	repo Repository
}

func (h *Handler) Create(c interface{}, user *pb.User, response *pb.Response) error {
     err:=h.repo.Create(user)
     if err!=nil{
     	return err
	 }
	 response.User=user
	 return nil
}

func (h *Handler) Get(c interface{}, user *pb.User, response *pb.Response) error {
      u,err:=h.repo.Get(user.Id)
      if err!=nil{
      	return err
	  }
	  response.User=u
	  return nil
}

func (h *Handler) GetAll(c interface{}, request *pb.Request, response *pb.Response) error {
      users,err := h.repo.GetAll()
      if err!=nil{
      	return err
	  }
	  response.Users = users
	  return nil
}

func (h *Handler) Auth(c interface{}, user *pb.User, token *pb.Token) error {
	_, err := h.repo.GetByEmailAndPassword(user)
	if err != nil {
		return err
	}
	token.Token = "`x_2nam"
	return nil
}

func (h *Handler) ValidateToken(c interface{}, token *pb.Token, token2 *pb.Token) error {
     return nil
}

