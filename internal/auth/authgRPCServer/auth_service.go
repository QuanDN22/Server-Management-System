package authgRPCServer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/QuanDN22/Server-Management-System/internal/auth/domain"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/QuanDN22/Server-Management-System/pkg/utils"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// login
func (a *AuthGrpcServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	// fmt.Printf("Login request %s, %s", in.GetUsername(), in.GetPassword())
	// fmt.Println("5. server auth /login")

	un := in.GetUsername()
	pw := in.GetPassword()

	// fmt.Printf("un: %s %s\n", un, pw)

	// fmt.Println("5. here")

	if un == "" || pw == "" {
		return &auth.LoginResponse{}, status.Error(codes.InvalidArgument, "missing username or password")
	}

	var existUser domain.User
	res := a.db.Where(&domain.User{UserName: un}).First(&existUser)

	// fmt.Println(existUser)

	if res.RowsAffected == 0 {
		return &auth.LoginResponse{}, status.Error(codes.NotFound, "user not found")
	}

	if !utils.CompareHashPassword(pw, existUser.Password) {
		return &auth.LoginResponse{}, status.Error(codes.InvalidArgument, "invalid password")
	}

	tokenString, err := a.issuer.IssueToken(existUser.UserName, existUser.Role)
	if err != nil {
		return &auth.LoginResponse{}, status.Error(codes.Internal, "failed to issue token")
	}

	return &auth.LoginResponse{
		AccessToken: tokenString,
	}, status.Error(codes.OK, "successful login")
}

// ping
func (a *AuthGrpcServer) Ping(ctx context.Context, _ *emptypb.Empty) (*auth.PingMessage, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &auth.PingMessage{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	return &auth.PingMessage{
		Message: fmt.Sprintf("Pong auth server, %v", roles),
	}, status.Error(codes.OK, "ping successful")
}

// signup
func (a *AuthGrpcServer) Signup(ctx context.Context, in *auth.SignupRequest) (*emptypb.Empty, error) {
	var user domain.User

	user.UserName = in.GetUsername()
	user.Password = in.GetPassword()
	user.Email = in.GetEmail()
	user.Role = "user"

	// fmt.Println("user: ", user)

	if user.UserName == "" || user.Password == "" || user.Email == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "missing username or password or email")
	}

	var existUser domain.User
	res := a.db.First(&existUser, "username = ?", user.UserName)

	if res.RowsAffected == 1 {
		return &emptypb.Empty{}, status.Error(codes.AlreadyExists, "username already exists")
	}

	// fmt.Println("existUser: ", existUser)

	res = a.db.First(&existUser, "email = ?", user.Email)
	if res.RowsAffected == 1 {
		return &emptypb.Empty{}, status.Error(codes.AlreadyExists, "email already exists")
	}

	// fmt.Println("existUser: ", existUser)

	if user.Email == existUser.Email {
		return &emptypb.Empty{}, status.Error(codes.AlreadyExists, "email already exists")
	}

	var err error
	user.Password, err = utils.GenerateHashPassword(user.Password)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to hash password")
	}

	// fmt.Println("hashed password: ", user.Password)

	res = a.db.Create(&user)
	if res.RowsAffected == 1 {
		fmt.Println("successful")
		return &emptypb.Empty{}, status.Error(codes.OK, "successful")
	}

	// fmt.Println("failed to create user")

	return &emptypb.Empty{}, status.Error(codes.Internal, "failed to create user")
}

// // logout
// func (a *AuthGrpcServer) Logout(ctx context.Context, in *auth.LogoutRequest) (*emptypb.Empty, error) {

// }

// // change password
// func (a *AuthGrpcServer) ChangePassword(ctx context.Context, in *auth.ChangePasswordRequest) (*emptypb.Empty, error) {

// }

// admin delete a user by ID
func (a *AuthGrpcServer) DeleteUserByID(ctx context.Context, in *auth.DeleteUserRequest) (*emptypb.Empty, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	fmt.Println("token ", token)

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	fmt.Println(roles)

	if roles != "admin" {
		return &emptypb.Empty{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	userID := in.GetUserId()
	id, _ := strconv.Atoi(userID)

	var user domain.User
	res := a.db.First(&user, "id = ?", id)

	if res.RowsAffected == 0 {
		return &emptypb.Empty{}, status.Error(codes.NotFound, "user not found")
	}

	fmt.Println("user", user)

	if user.Role != "user" {
		return &emptypb.Empty{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	res = a.db.Delete(&user)
	if res.RowsAffected != 1 || res.Error != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to delete user")
	}

	return &emptypb.Empty{}, status.Error(codes.OK, "successful delete user by id given")
}
