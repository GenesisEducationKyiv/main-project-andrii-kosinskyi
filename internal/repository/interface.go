package repository

import "bitcoin_checker_api/internal/model"

type Repository interface {
	Write(email string) error
	ReadAll() []*model.User
	ExistsByEmail(e string) bool
}
