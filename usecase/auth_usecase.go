package usecase

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/model"
	"savegen-api/repository"
	"savegen-api/util"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(req dto.RegisterRequest) (entity.User, string, error)
	Login(req dto.LoginRequest) (entity.User, string, error)
	Guest(req dto.GuestRequest) (entity.User, string, error)
	GoogleLogin(req dto.GoogleLoginRequest) (entity.User, string, error)
	GetCurrentUser(userID int) (entity.User, error)
}

type authUsecase struct {
	userRepository repository.UserRepository
}

type AuthUsecaseConfig struct {
	UserRepository repository.UserRepository
}

func NewAuthUsecase(cfg *AuthUsecaseConfig) AuthUsecase {
	return &authUsecase{userRepository: cfg.UserRepository}
}

var (
	ErrEmailAlreadyUsed    = errors.New("email is already registered")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInvalidGoogleToken  = errors.New("invalid Google ID token")
	ErrGoogleEmailMismatch = errors.New("Google account has no verified email")
)

func (u *authUsecase) Register(req dto.RegisterRequest) (entity.User, string, error) {
	_, found, err := u.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return entity.User{}, "", err
	}
	if found {
		return entity.User{}, "", model.ErrInvalidInput{Field: "email", Reason: ErrEmailAlreadyUsed.Error()}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, "", err
	}
	hashStr := string(hash)

	user := entity.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: &hashStr,
		IsGuest:      false,
	}
	created, err := u.userRepository.CreateUser(user)
	if err != nil {
		return entity.User{}, "", err
	}

	token, err := util.SignToken(created.ID)
	if err != nil {
		return entity.User{}, "", err
	}
	return created, token, nil
}

func (u *authUsecase) Login(req dto.LoginRequest) (entity.User, string, error) {
	user, found, err := u.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return entity.User{}, "", err
	}
	if !found || user.PasswordHash == nil {
		return entity.User{}, "", model.ErrInvalidInput{Field: "credentials", Reason: ErrInvalidCredentials.Error()}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(req.Password)); err != nil {
		return entity.User{}, "", model.ErrInvalidInput{Field: "credentials", Reason: ErrInvalidCredentials.Error()}
	}

	token, err := util.SignToken(user.ID)
	if err != nil {
		return entity.User{}, "", err
	}
	return user, token, nil
}

func (u *authUsecase) Guest(req dto.GuestRequest) (entity.User, string, error) {
	_, found, err := u.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return entity.User{}, "", err
	}
	if found {
		return entity.User{}, "", model.ErrInvalidInput{Field: "email", Reason: ErrEmailAlreadyUsed.Error()}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, "", err
	}
	hashStr := string(hash)

	username := req.Username
	if username == "" {
		username = "Guest"
	}

	user := entity.User{
		Username:     username,
		Email:        req.Email,
		PasswordHash: &hashStr,
		IsGuest:      true,
	}
	created, err := u.userRepository.CreateUser(user)
	if err != nil {
		return entity.User{}, "", err
	}

	token, err := util.SignToken(created.ID)
	if err != nil {
		return entity.User{}, "", err
	}
	return created, token, nil
}

// googleTokenInfo represents the fields we care about from Google's tokeninfo endpoint.
type googleTokenInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Aud           string `json:"aud"`
}

// verifyGoogleIDToken talks to Google's tokeninfo endpoint to validate the ID token
// and pull the identity claims. For production you should switch to local JWKS-based
// verification, but this is the simplest correct implementation.
func verifyGoogleIDToken(idToken string) (*googleTokenInfo, error) {
	endpoint := "https://oauth2.googleapis.com/tokeninfo?" + url.Values{"id_token": {idToken}}.Encode()
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, ErrInvalidGoogleToken
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var info googleTokenInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, ErrInvalidGoogleToken
	}
	if info.Sub == "" || info.Email == "" {
		return nil, ErrInvalidGoogleToken
	}
	if info.EmailVerified != "true" {
		return nil, ErrGoogleEmailMismatch
	}
	return &info, nil
}

func (u *authUsecase) GoogleLogin(req dto.GoogleLoginRequest) (entity.User, string, error) {
	info, err := verifyGoogleIDToken(req.IDToken)
	if err != nil {
		return entity.User{}, "", model.ErrInvalidInput{Field: "id_token", Reason: err.Error()}
	}

	// 1. Try to find by google_sub (returning Google user).
	if user, found, err := u.userRepository.FindUserByGoogleSub(info.Sub); err != nil {
		return entity.User{}, "", err
	} else if found {
		token, err := util.SignToken(user.ID)
		if err != nil {
			return entity.User{}, "", err
		}
		return user, token, nil
	}

	// 2. Try to find by email. If found, auto-link Google to existing account.
	if user, found, err := u.userRepository.FindUserByEmail(info.Email); err != nil {
		return entity.User{}, "", err
	} else if found {
		if err := u.userRepository.UpdateGoogleSub(user.ID, info.Sub); err != nil {
			return entity.User{}, "", err
		}
		sub := info.Sub
		user.GoogleSub = &sub
		token, err := util.SignToken(user.ID)
		if err != nil {
			return entity.User{}, "", err
		}
		return user, token, nil
	}

	// 3. Brand new user: create one with no password, linked to Google.
	sub := info.Sub
	username := info.Name
	if username == "" {
		username = "Google User"
	}
	user := entity.User{
		Username:  username,
		Email:     info.Email,
		GoogleSub: &sub,
		IsGuest:   false,
	}
	created, err := u.userRepository.CreateUser(user)
	if err != nil {
		return entity.User{}, "", err
	}
	token, err := util.SignToken(created.ID)
	if err != nil {
		return entity.User{}, "", err
	}
	return created, token, nil
}

func (u *authUsecase) GetCurrentUser(userID int) (entity.User, error) {
	return u.userRepository.GetUserById(userID)
}
