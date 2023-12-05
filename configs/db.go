package configs

import (
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"assignment_4/repositories"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	UserRepository         interfaces.UserRepository
	WalletRepository       interfaces.WalletRepository
	TransactionRepository  interfaces.TransactionRepository
	SourceOfFundRepository interfaces.SourceOfFundRepository
	db                     *gorm.DB
}

func NewRepository(config *models.Config) (*Repository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName,
	)
	log.Println(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("DB error : db.go : ", err)
		return nil, err
	}
	return &Repository{
		UserRepository:         repositories.NewUserRepository(db),
		WalletRepository:       repositories.NewWalletRepository(db),
		TransactionRepository:  repositories.NewTransactionRepository(db),
		SourceOfFundRepository: repositories.NewSourceOfFundRepo(db),
		db:                     db,
	}, nil
}

func (r *Repository) Automigrate() error {

	r.db.Exec("CREATE SEQUENCE my_sequence")

	return r.db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Game{},
		&models.Wallet{},
		&models.SourceOfFund{},
		&models.Transaction{},
	)
}

func (r *Repository) TrxBegin() *gorm.DB {
	return r.db.Begin()
}
