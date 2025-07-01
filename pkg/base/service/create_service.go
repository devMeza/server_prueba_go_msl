package base_service

import (
	helpers "main/pkg/base/helpers"
	base_models "main/pkg/base/models"
	models "main/pkg/base/models"
)

// Crear un nuevo modelo y guardarlo
// func NewController[T any](model models.Model[T]) (*Controller[T], bool) {
// 	controller := &Controller[T]{
// 		Model: model,
// 	}
// 	SaveController(controller)
// 	return controller, true
// }

// Crear un nuevo controlador con métodos personalizados
func NewServices[T any](model models.Model[T], methods Methods[T]) *Service[T] {
	var controller *Service[T]

	if methods != nil {
		controller = &Service[T]{Model: model, Methods: methods}
	} else {
		controller = &Service[T]{Model: model}
	}

	SaveService(controller)
	return controller
}

// Guardar un controlador en el mapa global
func SaveService[T any](service *Service[T]) {
	helpers.SaveStructure(service, &services)
}

// InitGeneric inicializa un controlador genérico con un modelo y métodos opcionales.
func Init[T any, M base_models.Model[T]](c Methods[T]) *Service[T] {
	if c == nil {
		m, _ := models.GetModel[T]()
		if m == nil {
			return nil
		}
		return NewServices(*m, nil)
	}
	return NewServices(c.GetModel(), c)
}
