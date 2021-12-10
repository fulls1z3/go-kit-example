package main

type Service interface {
	add(name string) error
	remove(id int) error
	getAll() ([]model, error)
}

type svc struct{}

func NewService() Service {
	return &svc{}
}

func (s *svc) add(name string) error {
	return nil
}

func (s *svc) remove(id int) error {
	return nil
}

func (s *svc) getAll() ([]model, error) {
	return []model{}, nil
}

type model struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
