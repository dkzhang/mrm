package handlers

import "mrm/ent"

type Handler struct {
	DbClient *ent.Client
}

func NewHandler(dbClient *ent.Client) *Handler {
	return &Handler{
		DbClient: dbClient,
	}
}
