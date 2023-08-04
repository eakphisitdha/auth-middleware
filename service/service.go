package service

import (
	"app/model"
	"app/repository"
	"app/transaction"
	"log"
)

type IService interface {
	Get(req model.GetRequest) ([]model.DataResponse, error)
	Add(req model.AddRequest) error
	Update(req model.UpdateRequest) error
	Delete(req model.DeleteRequest) error
}

type Service struct {
	r repository.IRepository
	t transaction.ITransaction
}

func NewService(r repository.IRepository, t transaction.ITransaction) IService {
	return &Service{r: r, t: t}
}

func (s *Service) Get(req model.GetRequest) ([]model.DataResponse, error) {
	data, err := s.r.Get(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	//
	// put your business logic
	//
	response := []model.DataResponse{}
	for _, i := range data {
		var res model.DataResponse
		//
		// put your business logic
		//
		res.ID = i.ID
		res.Name = i.Name
		if i.Field.Valid {
			res.Field = i.Field.String
		}
		//
		response = append(response, res)
	}

	return response, nil
}

func (s *Service) Add(req model.AddRequest) error {
	//
	// put your business logic
	//
	err := s.r.Add(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.t.Log("field", "value", "ADD")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) Update(req model.UpdateRequest) error {

	err := s.r.Update(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = s.t.Log("field", "value", "UPDATE")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) Delete(req model.DeleteRequest) error {

	err := s.r.Delete(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.t.Log("field", "value", "DELETE")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
