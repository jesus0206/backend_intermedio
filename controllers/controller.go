package controllers

import "yofio/intermedio/store"

// A Controller response
type Controller struct {
	repo store.IRepository
}

// NewController function
func NewController() *Controller {
	return &Controller{
		repo: store.NewRepository(),
	}
}
