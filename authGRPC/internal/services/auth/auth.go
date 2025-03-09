package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorthDice/AuthGRPC/internal/domain/models"
	"github.com/NorthDice/AuthGRPC/internal/lib/jwt"
	"github.com/NorthDice/AuthGRPC/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passwordHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (user models.User, err error)
	IsAdmin(ctx context.Context, userID int64) (isAdmin bool, err error)
}

type AppProvider interface {
	App(ctx context.Context, appID int32) (app models.App, err error)
}

var (
	invalidCredentials = errors.New("invalid credentials")
)

// New returns a new instance of the Auth service.
func New(
	log *slog.Logger,
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver:    userSaver,
		usrProvider: userProvider,
		log:         log,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

// Login checks if user with given credentials exists in the system.
func (auth *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	appID int32,
) (string, error) {
	const op = "auth.Login"

	log := auth.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("Attempting to log in")

	user, err := auth.usrProvider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			auth.log.Warn("User not found")

			return "", fmt.Errorf("%s: %w", op, invalidCredentials)
		}

		auth.log.Warn("Failed to get user")

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		auth.log.Warn("Invalid credentials")

		return "", fmt.Errorf("%s: %w", op, invalidCredentials)
	}

	app, err := auth.appProvider.App(ctx, appID)
	if err != nil {
		auth.log.Warn("Failed to get app")
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Successfully logged in")

	token, err := jwt.NewToken(user, app, auth.tokenTTL)
	if err != nil {
		auth.log.Warn("Failed to create token")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

// RegisterNewUser registers a new user with the given email and password.
func (auth *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "auth.RegisterNewUser"

	log := auth.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("Registering new user")

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Failed to generate password", err)
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := auth.usrSaver.SaveUser(ctx, email, passwordHash)
	if err != nil {
		log.Error("Failed to save user", err)

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("User registered")

	return id, nil
}

// IsAdmin checks if the user is admin
func (auth *Auth) IsAdmin(
	ctx context.Context,
	userID int64,
) (bool, error) {
	const op = "auth.IsAdmin"

	log := auth.log.With(
		slog.String("op", op),
	)

	log.Info("Checking if user is admin")

	isAdmin, err := auth.usrProvider.IsAdmin(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Checked if user is admin", slog.Bool("isAdmin", isAdmin))

	return isAdmin, nil
}
