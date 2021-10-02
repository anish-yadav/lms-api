package handler

import "github.com/anish-yadav/lms-api/internal/pkg/class"

type ClassRequestResponse struct {
	ID string `json:"id"`
}

type ClassResponse struct {
	classes []*class.Class
}
