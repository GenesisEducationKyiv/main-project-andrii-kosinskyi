package repositories

import "bitcoin_checker_api/internal/models"

type Repository interface {
	Write(email string) error
	ReadAll() []*models.User
}
