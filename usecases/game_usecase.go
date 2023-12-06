package usecases

import (
	"assignment_4/apperror"
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
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

// ChooseReward implements interfaces.GameUsecase.
func (usecase *gameUsecase) ChooseReward(ctx context.Context, chooseReward *req.ChooseReward) (*models.Game, error) {
	game, err := usecase.gameRepo.GetById(ctx, chooseReward.GameID)
	if err != nil {
		return nil, err
	}
	//check if box is exists in game
	var boxId uint = 0
	switch chooseReward.BoxID {
	case game.BoxID1:
		boxId = game.BoxID1
	case game.BoxID2:
		boxId = game.BoxID2
	case game.BoxID3:
		boxId = game.BoxID3
	case game.BoxID4:
		boxId = game.BoxID4
	case game.BoxID5:
		boxId = game.BoxID5
	case game.BoxID6:
		boxId = game.BoxID6
	case game.BoxID7:
		boxId = game.BoxID7
	case game.BoxID8:
		boxId = game.BoxID8
	case game.BoxID9:
		boxId = game.BoxID9
	default:
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
