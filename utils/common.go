package utils

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"

	"github.com/go-playground/validator/v10"
)

func Validate(obj any, err error) []string {
	var errBadRequests []string

	if err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("%s field is %s", fieldErr.Field(), fieldErr.ActualTag())
			errBadRequests = append(errBadRequests, errMsg)
		}
	}

	return errBadRequests
}

func ValidateWinBox(game *models.Game, chooseReward *req.ChooseReward) (uint, error) {
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
		return boxId, &apperror.ErrInvalidRequest{Field: enums.BoxId}
	}
	return boxId, nil
}
