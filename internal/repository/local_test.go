package repository_test

import (
	"errors"
	"fmt"
	"testing"

	"bitcoin_checker_api/internal/repository"

	"bitcoin_checker_api/config"
)

func TestLocalRepository_Write(t *testing.T) {
	repo, err := repository.NewLocalRepository(&config.Storage{Path: "./storage.json"})
	if err != nil {
		t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
	}
	defer repo.RemoveAll()

	if err = repo.Write("taras@schevchenko.com"); err != nil || len(repo.Records) == 0 {
		t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
	}
}

func TestLocalRepository_DoNotWriteDuplicateRecord(t *testing.T) {
	repo, err := repository.NewLocalRepository(&config.Storage{Path: "./storage.json"})
	if err != nil {
		t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
	}
	defer repo.RemoveAll()

	err = repo.Write("taras@schevchenko.com")
	if err != nil {
		t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
	}
	err = repo.Write("taras@schevchenko.com")
	if !errors.Is(err, repository.ErrRecordExists) || len(repo.Records) != 1 {
		t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
	}
}

func TestLocalRepository_ReadAll(t *testing.T) {
	numRecords := 7
	repo, err := repository.NewLocalRepository(&config.Storage{Path: "./storage.json"})
	if err != nil {
		t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
	}
	defer repo.RemoveAll()

	for i := 0; i < numRecords; i++ {
		err = repo.Write(fmt.Sprintf("taras@schevchenko%d.com", i))
		if err != nil {
			t.Errorf("TestRepository_Write() err = %v record len = %d", err, len(repo.Records))
		}
	}

	records := repo.ReadAll()
	if len(records) != numRecords {
		t.Errorf("TestRepository_Write() expected record len = %d but actual record len = %d", numRecords, len(repo.Records))
	}
}
