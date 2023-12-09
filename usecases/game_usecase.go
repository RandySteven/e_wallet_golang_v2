package usecases

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/utils"
)

type gameUsecase struct {
	gameRepo        interfaces.GameRepository
	userRepo        interfaces.UserRepository
	transactionRepo interfaces.TransactionRepository
	walletRepo      interfaces.WalletRepository
	boxesRepo       interfaces.BoxRepository
}

// GetUserCurrentChance implements interfaces.GameUsecase.
func (usecase *gameUsecase) GetUserCurrentChance(ctx context.Context, userId uint) (*models.User, error) {
	return usecase.userRepo.GetById(ctx, userId)
}

// ChooseReward implements interfaces.GameUsecase.
func (usecase *gameUsecase) ChooseReward(ctx context.Context, chooseReward *req.ChooseReward) (*models.Game, error) {
	game, err := usecase.gameRepo.GetById(ctx, chooseReward.GameID)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, &apperror.ErrDataNotFound{Data: "game"}
	}

	if game.WinBoxID != 0 {
		return nil, &apperror.ErrInvalidRequest{Field: enums.BoxId}
	}

	boxId, err := utils.ValidateWinBox(game, chooseReward)
	if err != nil {
		return nil, err
	}

	game.WinBoxID = boxId
	game, err = usecase.gameRepo.CreateRewardTransaction(ctx, game)
	if err != nil {
		return nil, err
	}

	return game, nil
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
