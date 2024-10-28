package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/configs"
	"github.com/iypetrov/gopizza/internal/database"
	"github.com/iypetrov/gopizza/internal/dtos"
	"github.com/iypetrov/gopizza/internal/toasts"
	"github.com/lib/pq"
)

type Auth struct {
	db            *sql.DB
	queries       *database.Queries
	cognitoClient *cip.Client
}

func NewAuth(db *sql.DB, queries *database.Queries, cognitoClient *cip.Client) Auth {
	return Auth{
		db:            db,
		queries:       queries,
		cognitoClient: cognitoClient,
	}
}

func (srv *Auth) CreateUser(ctx context.Context, email, password, address string) (uuid.UUID, error) {
	result, err := srv.cognitoClient.SignUp(ctx, &cip.SignUpInput{
		ClientId: aws.String(configs.Get().AWS.CognitoClientID),
		Username: aws.String(email),
		Password: aws.String(password),
	})
	if err != nil {
		return uuid.Nil, toasts.GetAWSError(err)
	}

	userID, err := uuid.Parse(*result.UserSub)
	if err != nil {
		return uuid.Nil, toasts.ErrUserCreation
	}
	p := database.CreateUserParams{
		ID:        userID,
		Email:     email,
		Address:   address,
		CreatedAt: time.Now().UTC(),
	}
	_, err = srv.queries.CreateUser(ctx, p)
	if err != nil {
		var pgErr *pq.Error

		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "23505" {
				return uuid.Nil, toasts.ErrUserAlreadyExists
			}
		}

		return uuid.Nil, toasts.ErrUserCreation
	}

	return userID, nil
}

func (srv *Auth) VerifyUserCode(ctx context.Context, id uuid.UUID, email, code string) error {
	_, err := srv.cognitoClient.ConfirmSignUp(ctx, &cip.ConfirmSignUpInput{
		ClientId:         aws.String(configs.Get().AWS.CognitoClientID),
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
	})
	if err != nil {
		return toasts.GetAWSError(err)
	}

	p := database.ConfirmUserParams{
		ID: id,
		ConfirmedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}
	_, err = srv.queries.ConfirmUser(ctx, p)
	if err != nil {
		return toasts.ErrUserCreation
	}

	return nil
}

func (srv *Auth) VerifyUser(ctx context.Context, email, password string) (dtos.UserCookie, error) {
	result, err := srv.cognitoClient.InitiateAuth(ctx, &cip.InitiateAuthInput{
		ClientId:       aws.String(configs.Get().AWS.CognitoClientID),
		AuthFlow:       "USER_PASSWORD_AUTH",
		AuthParameters: map[string]string{"USERNAME": email, "PASSWORD": password},
	})
	if err != nil {
		return dtos.UserCookie{}, toasts.GetAWSError(err)
	}

	user, err := srv.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return dtos.UserCookie{}, toasts.ErrUserNotFound
	}

	cookie := dtos.UserCookie{
		ID:           user.ID.String(),
		Email:        email,
		AccessToken:  *result.AuthenticationResult.AccessToken,
		RefreshToken: *result.AuthenticationResult.RefreshToken,
	}
	return cookie, nil
}
