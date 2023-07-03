package repository

import (
	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/model"
)

type MockRepository struct {
	Records    []*model.User
	RecordsMap map[string]struct{}
	Cfg        *config.Storage
}

func NewMockRepository(cfg *config.Storage) (*MockRepository, error) {
	sr := &MockRepository{
		Cfg:        cfg,
		RecordsMap: make(map[string]struct{}),
		Records:    make([]*model.User, 0),
	}

	return sr, nil
}

func (that *MockRepository) Write(email string) error {
	if ok := that.ExistsByEmail(email); ok {
		return ErrRecordExists
	}

	that.RecordsMap[email] = struct{}{}
	that.Records = append(that.Records, model.NewUser(email))

	return nil
}

func (that *MockRepository) ReadAll() []*model.User {
	return that.Records
}

func (that *MockRepository) ExistsByEmail(e string) bool {
	_, ok := that.RecordsMap[e]
	return ok
}
