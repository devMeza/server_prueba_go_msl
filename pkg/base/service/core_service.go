package base_service

import (
	models "main/pkg/base/models"
	"main/pkg/hooks"
	"sync"
)

// Controller agrupa un modelo con sus métodos
type Service[T any] struct {
	models.Model[T]
	Methods[T]
	hooks.Hookable
	hooks.Cleaners
}

// Methods define los métodos CRUD
type Methods[T any] interface {
	Read(filter map[string]any, config map[string]int) ([]map[string]any, error)
	Insert(data T) error
	Update(filter map[string]any, data T) error
	Delete(filter map[string]any) error
	GetModel() models.Model[T]
	NewService() *Service[T]
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var services sync.Map // key: string, value: *Model[any]
