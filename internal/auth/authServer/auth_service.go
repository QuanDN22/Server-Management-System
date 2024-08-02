package gRPCServer

import (
	"context"
	"fmt"

	"github.com/QuanDN22/Server-Management-System/internal/auth/domain"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/QuanDN22/Server-Management-System/pkg/utils"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ping
func (a *AuthServer) Ping(ctx context.Context, _ *emptypb.Empty) (*auth.PingMessage, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &auth.PingMessage{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	return &auth.PingMessage{
		Message: fmt.Sprintf("Pong auth server, %v", roles),
	}, nil
}

// login
func (a *AuthServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
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
	}, nil
}

// signup
func (a *AuthServer) Signup(ctx context.Context, in *auth.SignupRequest) (*emptypb.Empty, error) {
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
		return &emptypb.Empty{}, nil
	}

	// fmt.Println("failed to create user")

	return &emptypb.Empty{}, status.Error(codes.Internal, "failed to create user")
}

// // logout
// func (a *AuthServer) Logout(ctx context.Context, in *auth.LogoutRequest) (*emptypb.Empty, error) {

// }

// // change password
func (a *AuthServer) ChangePassword(ctx context.Context, in *auth.ChangePasswordRequest) (*emptypb.Empty, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	// dig the roles from the claims
	username := token.Claims.(jwt.MapClaims)["username"].(string)

	var userInToken domain.User
	res := a.db.First(&userInToken, "username = ?", username)
	if res.Error != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to get user information")
	}
	if res.RowsAffected == 0 {
		return &emptypb.Empty{}, status.Error(codes.Internal, "permission denied, user deleted")
	}
	
	if in.GetUserId() < 1 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "invalid user id")
	}

	// user can only change their own password
	if userInToken.ID != uint(in.GetUserId()) {
		return &emptypb.Empty{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	if res.Error != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to get user information")
	}

	if res.RowsAffected == 0 {
		return &emptypb.Empty{}, status.Error(codes.NotFound, "user not found")
	}

	oldPassword := in.GetOldPassword()
	newPassword := in.GetNewPassword()

	if oldPassword == "" || newPassword == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "missing old password or new password")
	}

	if oldPassword == newPassword {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "old password and new password are the same")
	}

	if !utils.CompareHashPassword(oldPassword, userInToken.Password) {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "invalid old password")
	}

	userInToken.Password, err = utils.GenerateHashPassword(newPassword)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to hash password")
	}

	res = a.db.Save(&userInToken)
	if res.RowsAffected == 1 {
		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, status.Error(codes.Internal, "failed to change password")
}

// get user information by ID
// admin can get all users and can not get information of other admin
// user can only get their own information
func (a *AuthServer) GetUserByID(ctx context.Context, id *auth.UserID) (*auth.User, error) {
	if id.GetUserId() < 1 {
		return &auth.User{}, status.Error(codes.InvalidArgument, "invalid user id")
	}

	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &auth.User{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	// dig the roles from the claims
	username := token.Claims.(jwt.MapClaims)["username"].(string)
	roles := token.Claims.(jwt.MapClaims)["roles"]

	var userInToken domain.User
	res := a.db.First(&userInToken, "username = ?", username)
	if res.RowsAffected == 0 {
		return &auth.User{}, status.Error(codes.Internal, "permission denied, user deleted")
	}

	if roles == "user" {
		// user can not get information of other user
		if userInToken.ID != uint(id.GetUserId()) {
			return &auth.User{}, status.Error(codes.PermissionDenied, "permission denied")
		}

		// user can only get their own information
		return &auth.User{
			UserId:    id.GetUserId(),
			Username:  userInToken.UserName,
			Email:     userInToken.Email,
			CreatedAt: userInToken.CreatedAt.String(),
			UpdatedAt: userInToken.UpdatedAt.String(),
		}, nil
	}

	var user domain.User
	res = a.db.First(&user, "id = ?", id.GetUserId())

	if res.RowsAffected == 0 {
		return &auth.User{}, status.Error(codes.NotFound, "user not found")
	}

	fmt.Println("user", user)

	// admin can not get information of other admin
	if user.Role == "admin" && user.UserName != userInToken.UserName {
		return &auth.User{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	return &auth.User{
		UserId:    id.GetUserId(),
		Username:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

// api for admin
// get all users
func (a *AuthServer) GetAllUsers(ctx context.Context, _ *emptypb.Empty) (*auth.GetAllUsersResponse, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &auth.GetAllUsersResponse{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	fmt.Println("token ", token)

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	fmt.Println(roles)

	if roles != "admin" {
		return &auth.GetAllUsersResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	var users []*auth.User

	res := a.db.Model(&domain.User{}).Select("id", "username", "email", "created_at", "updated_at").Where("role = ?", "user").Find(&users)

	if res.Error != nil {
		return &auth.GetAllUsersResponse{}, status.Error(codes.Internal, "failed to get all users")
	}

	fmt.Println("users", users)

	return &auth.GetAllUsersResponse{
		Users: users,
	}, nil
}

// api for admin
// admin delete a user by ID
func (a *AuthServer) DeleteUserByID(ctx context.Context, in *auth.UserID) (*emptypb.Empty, error) {
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

	var user domain.User
	res := a.db.First(&user, "id = ?", userID)

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

	return &emptypb.Empty{}, nil
}

// api for admin
// get email of admin
func (a *AuthServer) GetAdminEmail(ctx context.Context, _ *emptypb.Empty) (*auth.GetAdminEmailResponse, error) {
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		return &auth.GetAdminEmailResponse{}, status.Error(codes.Unauthenticated, "no auth provided")
	}

	fmt.Println("token ", token)

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	fmt.Println(roles)

	if roles != "admin" {
		return &auth.GetAdminEmailResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	var emails []string
	res := a.db.Model(&domain.User{}).Where("role = ?", "admin").Select("email").Find(&emails)

	if res.Error != nil {
		return &auth.GetAdminEmailResponse{}, status.Error(codes.Internal, "failed to get admin email")
	}

	fmt.Println("emails", emails)

	return &auth.GetAdminEmailResponse{
		Email: emails,
	}, nil
}
