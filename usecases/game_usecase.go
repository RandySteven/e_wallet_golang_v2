package usecases

import (
	"assignment_4/apperror"
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"context"
)

type gameUsecase struct {
	gameRepo        interfaces.GameRepository
	userRepo        interfaces.UserRepository
	transactionRepo interfaces.TransactionRepository
	walletRepo      interfaces.WalletRepository
	boxesRepo       interfaces.BoxRepository
}

// PlayGame implements interfaces.GameUsecase.
func (usecase *gameUsecase) PlayGame(ctx context.Context, game *models.Game) (*models.Game, error) {
	user, err := usecase.userRepo.GetById(ctx, game.UserID)
	if err != nil {
		return nil, err
	}

	if user.Chance == 0 {
		return nil, &apperror.ErrZeroChance{}
	}

	boxes, err := usecase.boxesRepo.GetNineRandomBoxes(ctx)
	if err != nil {
		return nil, err
	}
	game.Box1 = &boxes[0]
	game.Box2 = &boxes[1]
	game.Box3 = &boxes[2]
	game.Box4 = &boxes[3]
	game.Box5 = &boxes[4]
	game.Box6 = &boxes[5]
	game.Box7 = &boxes[6]
	game.Box8 = &boxes[7]
	game.Box9 = &boxes[8]

	game, err = usecase.gameRepo.Save(ctx, game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func NewGameUsecase(
	gameRepo interfaces.GameRepository,
	userRepo interfaces.UserRepository,
	transactionRepo interfaces.TransactionRepository,
	walletRepo interfaces.WalletRepository,
	boxesRepo interfaces.BoxRepository,
) *gameUsecase {
	return &gameUsecase{
		gameRepo:        gameRepo,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
		boxesRepo:       boxesRepo,
	}
}

var _ interfaces.GameUsecase = &gameUsecase{}
