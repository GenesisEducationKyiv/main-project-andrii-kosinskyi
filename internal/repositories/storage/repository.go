package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/models"
	"bitcoin_checker_api/internal/repositories"
)

type InternalStorageRepository struct {
	records    []*models.User
	recordsMap map[*models.User]struct{}
	cfg        *config.Config
}

func NewInternalStorageRepository(cfg *config.Config) (repositories.Repository, error) {
	isr := &InternalStorageRepository{}
	f, err := os.ReadFile(cfg.InternalStorage.Path)
	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
		return nil, err
	}

	if err = toml.Unmarshal(f, isr); err != nil {
		// failed to encode
		log.Fatal(err)
		return nil, err
	}

	for _, record := range isr.records {
		isr.recordsMap[record] = struct{}{}
	}

	isr.cfg = cfg

	return isr, nil
}

func (that *InternalStorageRepository) Write(email string) error {
	newUser := models.NewUser(email)
	if _, ok := that.recordsMap[newUser]; !ok {
		that.records = append(that.records, newUser)
		newRecords, err := toml.Marshal(that.records)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(that.cfg.InternalStorage.Path, newRecords, 0600)
		if err != nil {
			panic(err)
		}
		return nil
	}
	return fmt.Errorf("e-mail вже є в базі даних")
}

func (that *InternalStorageRepository) ReadAll() []*models.User {
	return that.records
}
