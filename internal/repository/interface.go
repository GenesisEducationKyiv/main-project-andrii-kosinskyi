package repository

import "bitcoin_checker_api/internal/model"

type Repository interface {
	Write(email string) error
	ReadAll() []*model.User
	RemoveAll() error
	ExistsByEmail(e string) bool
}
