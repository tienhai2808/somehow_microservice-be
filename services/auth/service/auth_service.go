package service

import (
	"context"
	"time"

	"github.com/SomeHowMicroservice/shm-be/services/auth/protobuf"
)

type AuthService interface {
	SignUp(ctx context.Context, req *protobuf.SignUpRequest) (string, error)

	VerifySignUp(ctx context.Context, req *protobuf.VerifySignUpRequest) (*protobuf.AuthResponse, string, time.Duration, string, time.Duration, error)

	SignIn(ctx context.Context, req *protobuf.SignInRequest) (*protobuf.AuthResponse, string, time.Duration, string, time.Duration, error)

	RefreshToken(ctx context.Context, req *protobuf.RefreshTokenRequest) (string, time.Duration, string, time.Duration, error)

	ChangePassword(ctx context.Context, req *protobuf.ChangePasswordRequest) (string, time.Duration, string, time.Duration, error)

	ForgotPassword(ctx context.Context, req *protobuf.ForgotPasswordRequest) (string, error)

	VerifyForgotPassword(ctx context.Context, req *protobuf.VerifyForgotPasswordRequest) (string, error)

	ResetPassword(ctx context.Context, req *protobuf.ResetPasswordRequest) error
}
