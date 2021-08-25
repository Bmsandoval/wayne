package hello

import "github.com/bmsandoval/wayne/db/models"

func (h *Helper) Get() ([]models.Greetings, error) {
	sdb := h.AppCtx.DB.SDB

	greetingModel := h.AppCtx.DB.GreetingModel

	query := greetingModel.Select("*")

	var result []models.Greetings
	err := sdb.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}
