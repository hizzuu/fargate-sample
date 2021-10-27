package interactor

import (
	"context"

	"github.com/hizzuu/app/internal/domain"
	"github.com/hizzuu/app/internal/usecase/repository"
)

type userInteractor struct {
	userRepo repository.UserRepository
}

type UserInteractor interface {
	Get(ctx context.Context, id int64) (*domain.User, error)
	List()
	Create()
	Update()
	Delete()
}

func NewUserInteractor(userRepo repository.UserRepository) UserInteractor {
	return &userInteractor{
		userRepo,
	}
}

func (i *userInteractor) Get(ctx context.Context, id int64) (*domain.User, error) {
	return i.userRepo.Get(ctx, id)
}

func (i *userInteractor) List() {
	i.userRepo.List()
}

func (i *userInteractor) Create() {
	i.userRepo.Create()
}

func (i *userInteractor) Update() {
	i.userRepo.Update()
}

func (i *userInteractor) Delete() {
	i.userRepo.Delete()
}
