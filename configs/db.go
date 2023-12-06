package configs

import (
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository           interfaces.UserRepository
	WalletRepository         interfaces.WalletRepository
	TransactionRepository    interfaces.TransactionRepository
	SourceOfFundRepository   interfaces.SourceOfFundRepository
	GameRepository           interfaces.GameRepository
	BoxRepository            interfaces.BoxRepository
	ForgotPasswordRepository interfaces.ForgotPasswordRepository
	db                       *gorm.DB
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
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Println("DB error : db.go : ", err)
		return nil, err
	}
	return &Repository{
		UserRepository:           repositories.NewUserRepository(db),
		WalletRepository:         repositories.NewWalletRepository(db),
		TransactionRepository:    repositories.NewTransactionRepository(db),
		SourceOfFundRepository:   repositories.NewSourceOfFundRepo(db),
		GameRepository:           repositories.NewGameRepository(db),
		ForgotPasswordRepository: repositories.NewForgotPasswordRepository(db),
		BoxRepository:            repositories.NewBoxRepository(db),
		db:                       db,
	}, nil
}

func (r *Repository) Automigrate() error {

	r.db.Exec("CREATE SEQUENCE IF NOT EXISTS my_sequence")

	return r.db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.ForgotPasswordToken{},
		&models.Game{},
		&models.Wallet{},
		&models.SourceOfFund{},
		&models.Transaction{},
	)
}

func (r *Repository) TrxBegin() *gorm.DB {
	return r.db.Begin()
}
