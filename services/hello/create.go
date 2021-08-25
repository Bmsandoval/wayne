package hello

import (
	"github.com/bmsandoval/wayne/db/models"
	"time"
)

func (h *Helper) Create(greetingModel models.Greetings) (*models.Greetings, error) {
	sdb := h.AppCtx.DB.SDB

	//Set "Defaults"
	now := time.Now().UTC()
	greetingModel.CreatedAt = now
	greetingModel.UpdatedAt = &now

	//Create Greeting
	if err := sdb.Insert(&greetingModel); err != nil {
		return nil, err
	}
	return &greetingModel, nil
}
