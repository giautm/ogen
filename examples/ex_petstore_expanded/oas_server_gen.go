// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddPet implements addPet operation.
	//
	// Creates a new pet in the store. Duplicates are allowed.
	//
	// POST /pets
	AddPet(ctx context.Context, req NewPet) (AddPetRes, error)
	// DeletePet implements deletePet operation.
	//
	// Deletes a single pet based on the ID supplied.
	//
	// DELETE /pets/{id}
	DeletePet(ctx context.Context, params DeletePetParams) (DeletePetRes, error)
	// FindPetByID implements find pet by id operation.
	//
	// Returns a user based on a single ID, if the user does not have access to the pet.
	//
	// GET /pets/{id}
	FindPetByID(ctx context.Context, params FindPetByIDParams) (FindPetByIDRes, error)
	// FindPets implements findPets operation.
	//
	// Returns all pets from the system that the user has access to
	// Nam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet.
	// Aenean at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum nec, cursus
	// id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea
	// dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra tempor ante molestie
	// imperdiet. Vivamus id aliquam diam. Cras quis velit non tortor eleifend sagittis. Praesent at enim
	// pharetra urna volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim
	// enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor, consectetur non augue vel,
	//  vehicula interdum libero. Morbi euismod sagittis libero sed lacinia.
	// Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu condimentum ligula luctus
	// nec. Phasellus semper velit eget aliquet faucibus. In a mattis elit. Phasellus vel urna viverra,
	// condimentum lorem id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio. Morbi
	// rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec scelerisque
	// tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis felis eu dolor tristique, sit
	// amet auctor felis gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi. Fusce
	// sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium,
	// pulvinar elit eu, euismod sapien.
	//
	// GET /pets
	FindPets(ctx context.Context, params FindPetsParams) (FindPetsRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s, err := newConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
