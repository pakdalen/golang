package students

import (
	"context"
	"endProject/pkg/api"
)

type GRPCServer struct{}

func (a *GRPCServer) GetStudent(ctx context.Context, req *api.GetStudentListRequest) (*api.GetStudentsListResponse, error) {
	return &api.GetStudentsListResponse{
		Gpa: req.GetMarks() / req.GetSubjects(),
	}, nil
}
