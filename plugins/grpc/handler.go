package grpc

import (
	"context"

	"github.com/anthonydevelops/tailored/plugins/grpc/model"
)

// CreateProject creates a new user project
func (*Service) CreateProject(ctx context.Context, request *model.Project) (*model.Empty, error) {
	return nil, nil
}

// DeleteProject deletes a previously stored user project
func (*Service) DeleteProject(ctx context.Context, request *model.ID) (*model.Empty, error) {
	return nil, nil
}

// GetProject returns a user project
func (*Service) GetProject(ctx context.Context, request *model.ID) (*model.Project, error) {
	return nil, nil
}
