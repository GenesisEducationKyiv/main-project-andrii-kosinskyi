package repository

import (
	"encoding/json"
	"log"
	"os"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/model"
)

type LocalRepository struct {
	Records    []*model.User
	RecordsMap map[string]struct{}
	Cfg        *config.Storage
}

func NewLocalRepository(cfg *config.Storage) (*LocalRepository, error) {
	sr := &LocalRepository{
		Cfg:        cfg,
		RecordsMap: make(map[string]struct{}),
		Records:    make([]*model.User, 0),
	}

	f, err := os.ReadFile(cfg.Path)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
			return nil, err
		}

		if os.IsNotExist(err) {
			_, err = os.Create(cfg.Path)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
		}
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

func (that *LocalRepository) Write(email string) error {
	if ok := that.ExistsByEmail(email); ok {
		return ErrRecordExists
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

func (that *LocalRepository) ReadAll() []*model.User {
	return that.Records
}

func (that *LocalRepository) ExistsByEmail(e string) bool {
	_, ok := that.RecordsMap[e]
	return ok
}
