package mysql

import (
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
	"testTask/cmd/internal/db/models"
)

type DB struct {
	db *gorm.DB
}

var modelsList = []interface{}{
	new(models.Auth),
	new(models.User),
	new(models.UserProfile),
	new(models.UserData),
}

func NewDatabase(dsn string, needRecreateDB bool) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	if needRecreateDB {
		if err = db.Migrator().DropTable(modelsList...); err != nil {
			return nil, err
		}
		if err = autoMigrate(db); err != nil {
			return nil, err
		}

		b, err := os.ReadFile("../cmd/internal/db/mysql/init.sql")
		if err != nil {
			return nil, err
		}

		tx := db.Begin()
		defer tx.Rollback()

		for _, q := range strings.Split(string(b), ";") {
			q = strings.TrimSpace(q)
			if q == "" {
				continue
			}
			if err = tx.Exec(q).Error; err != nil {
				return nil, err
			}
		}
		if err = tx.Commit().Error; err != nil {
			return nil, err
		}
	} else {
		if err = autoMigrate(db); err != nil {
			return nil, err
		}
	}
	return &DB{db}, nil
}

func (d *DB) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (d *DB) prepareSelectUserProfileWithAllData(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx).Model(new(models.User)).Joins("UserProfile").Joins("UserData")
}

func (d *DB) GetAllProfiles(ctx context.Context) ([]models.ProfileAPI, error) {
	var r []models.ProfileAPI
	if err := d.prepareSelectUserProfileWithAllData(ctx).Scan(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (d *DB) GetProfileByUsername(ctx context.Context, username string) (models.ProfileAPI, error) {
	var r models.ProfileAPI
	if err := d.
		prepareSelectUserProfileWithAllData(ctx).
		Where("username = ?", username).
		Scan(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, nil
		}
		return r, err
	}
	return r, nil
}

func (d *DB) IsExistAuthKey(ctx context.Context, key string) (bool, error) {
	var auth models.Auth
	if err := d.db.WithContext(ctx).Where("api_key = ?", key).Take(&auth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Auth{}, &models.User{}, &models.UserProfile{}, &models.UserData{}); err != nil {
		return err
	}
	return nil
}
