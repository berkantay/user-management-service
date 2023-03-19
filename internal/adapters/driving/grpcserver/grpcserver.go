package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/berkantay/user-management-service/internal/adapters/driving/proto"
	"github.com/berkantay/user-management-service/internal/model"

	"google.golang.org/grpc"
)

type UserService interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	RemoveUser(userId string) error
	QueryUsers(query *model.UserQuery) ([]model.User, error)
	DatabaseHealthCheck(ctx context.Context) error
	Echo(ctx context.Context) error
	GracefullShutdown() error
}

type Server struct {
	service UserService
	pb.UnimplementedUserApiServer
}

func NewServer(service UserService) *Server {

	return &Server{
		service: service,
	}
}

func (s *Server) Run() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}

	userManagementService := s

	grpcServer := grpc.NewServer()

	pb.RegisterUserApiServer(grpcServer, userManagementService)

	grpcServer.Serve(listen)
	defer grpcServer.Stop()

}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	wrappedMessage := createUserRequestToUser(req)

	err := s.service.CreateUser(wrappedMessage)

	if err != nil {
		return &pb.CreateUserResponse{
			Status: &pb.Status{
				Code:    "INTERNAL",
				Message: "Could not create user.",
			},
		}, err
	}

	return &pb.CreateUserResponse{
		Status: &pb.Status{
			Code:    "OK",
			Message: "User created.",
		}, //TODO Fill user info from db
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	fmt.Println("Deleting user", req.Id)

	err := s.service.RemoveUser(req.Id)

	if err != nil {
		return &pb.DeleteUserResponse{
			Status: &pb.Status{
				Code:    "INTERNAL",
				Message: "Could not delete user.",
			},
		}, err
	}

	return &pb.DeleteUserResponse{
		Status: &pb.Status{
			Code:    "OK",
			Message: "User deleted.",
		}, //TODO Fill user info from db
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	fmt.Println("Updating user", req.Id)

	err := s.service.UpdateUser(updateUserRequestToUser(req))

	if err != nil {
		return &pb.UpdateUserResponse{
			Status: &pb.Status{
				Code:    "INTERNAL",
				Message: "Could not update user.",
			},
		}, err
	}

	return &pb.UpdateUserResponse{
		Status: &pb.Status{
			Code:    "OK",
			Message: "User updated.",
		}, //TODO Fill user info from db
	}, nil

}

func (s *Server) QueryUsers(ctx context.Context, req *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {

	userQuery := toUserQuery(req)

	user, err := s.service.QueryUsers(userQuery)

	if err != nil {
		return &pb.QueryUsersResponse{
			Status: &pb.Status{
				Code:    "INTERNAL",
				Message: "Internal error occured",
			},
		}, err
	}

	if user == nil {
		return &pb.QueryUsersResponse{
			Status: &pb.Status{
				Code:    "NOT_FOUND",
				Message: "Could not found any user.",
			},
		}, err
	}

	return toPbQueryResponse(user), nil

}

func createUserRequestToUser(req *pb.CreateUserRequest) *model.User { //TODO:move this wrapping layer from server logic

	return &model.User{

		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
		Password:  req.Password,
		Email:     req.Email,
		Country:   req.Country,
	}
}

func updateUserRequestToUser(req *pb.UpdateUserRequest) *model.User { //TODO:move this wrapping layer from server logic

	return &model.User{
		ID:        req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
		Password:  req.Password,
		Email:     req.Email,
		Country:   req.Country,
	}
}

func toUserQuery(req *pb.QueryUsersRequest) *model.UserQuery {

	return &model.UserQuery{
		ID:        req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		NickName:  req.NickName,
		Email:     req.Email,
		Country:   req.Country,
	}

}

func toPbQueryResponse(users []model.User) *pb.QueryUsersResponse {

	payload := make([]*pb.UserPayload, 0)

	for _, u := range users {

		payload = append(payload, &pb.UserPayload{

			Id:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			NickName:  u.NickName,
			Password:  u.Password,
			Email:     u.Email,
			Country:   u.Country,
		})

	}

	return &pb.QueryUsersResponse{
		Status: &pb.Status{
			Code:    "OK",
			Message: "Users queried",
		},
		Payload: payload,
	}

}
