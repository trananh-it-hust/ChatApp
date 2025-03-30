package repo

type ItemRepo struct {
}

func Create() *ItemRepo {
	return &ItemRepo{}
}

func (i *ItemRepo) CreateItem() int {
	return 1
}
