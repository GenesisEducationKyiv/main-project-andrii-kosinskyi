package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/model"
	"bitcoin_checker_api/internal/repository"
)

type StorageRepository struct {
	Records    []*model.User
	RecordsMap map[string]struct{}
	Cfg        *config.Storage
}

func NewStorageRepository(cfg *config.Storage) (repository.Repository, error) {
	sr := &StorageRepository{
		Cfg:        cfg,
		RecordsMap: make(map[string]struct{}),
		Records:    make([]*model.User, 0),
	}

	f, err := os.ReadFile(cfg.Path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if len(f) != 0 {
		if err = json.Unmarshal(f, &sr.Records); err != nil {
			log.Fatal(err)
			return nil, err
		}

		for _, record := range sr.Records {
			sr.RecordsMap[record.Email] = struct{}{}
		}
	}

	return sr, nil
}

func (that *StorageRepository) Write(email string) error {
	if _, ok := that.RecordsMap[email]; ok {
		return fmt.Errorf("e-mail вже є в базі даних")
	}

	that.RecordsMap[email] = struct{}{}
	that.Records = append(that.Records, model.NewUser(email))

	jsonRecords, err := json.Marshal(&that.Records)
	if err != nil {
		return err
	}

	err = os.WriteFile(that.Cfg.Path, jsonRecords, 0o600)
	if err != nil {
		return err
	}
	return nil
}

func (that *StorageRepository) ReadAll() []*model.User {
	return that.Records
}
