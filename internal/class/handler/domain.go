package handler

import "github.com/anish-yadav/lms-api/internal/pkg/class"

type CreateClassRequest struct {
	Section  string `json:"section" validate:"required"`
	Semester string `json:"semester" validate:"required"`
	Stream   string `json:"stream" validate:"required"`
}



func (c *CreateClassRequest) toClass() *class.Class {
	return class.NewClass(c.Section,c.Semester,c.Stream)
}
