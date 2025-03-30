package service

import "main.go/internal/repo"

type ItemService struct {
	ItemRepo *repo.ItemRepo
}

func Create() *ItemService {
	return &ItemService{
		ItemRepo: repo.Create(),
	}
}

func (i *ItemService) CreateItem() int {
	x := i.ItemRepo.CreateItem()
	return x
}
