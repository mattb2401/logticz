package pkg

import (
	"fmt"
	"time"

	"github.com/logticz/logticz/lib"

	"github.com/go-kit/log"
	"github.com/logticz/logticz/internal/store/models"
	"gorm.io/gorm"
)

//AuthenticationSvc service interface
type AuthenticationSvc interface {
	HandleLogin(username, password string) (*models.User, error)
	HandleRegistration(names, username, password string) (*models.User, error)
}

//authentication defines the logger and db
type authentication struct {
	logger *log.Logger
	db     *gorm.DB
}

//NewAuthenticationSvc initiates a new AuthenticationSvc
func NewAuthenticationSvc(logger *log.Logger, db *gorm.DB) AuthenticationSvc {
	return &authentication{
		logger: logger,
		db:     db,
	}
}

// HandleLogin Route: POST /authenticate
//
// Only works with Basic Authentication (username and password).
//
// Responses:
// 200: userdata
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError
func (a *authentication) HandleLogin(username, password string) (*models.User, error) {
	if len(username) <= 0 || len(password) <= 0 {
		return nil, fmt.Errorf("wrong username or password")
	}
	user := models.User{}
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("wrong username or password")
	}
	if !lib.CheckArgonPassword(password, user.Password) {
		return nil, fmt.Errorf("wrong username or password")
	}
	return &user, nil
}

// HandleRegistration Route: POST /register
//
// Create new users with parameters, name, username, password
//
// Responses:
// 200: userdata
// 401: unauthorisedError
// 403: forbiddenError
// 500: internalServerError
func (a *authentication) HandleRegistration(names, username, password string) (*models.User, error) {
	if len(names) <= 0 || len(username) <= 0 || len(password) <= 0 {
		return nil, fmt.Errorf("missing required paramters")
	}
	hash, err := lib.HashArgonPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}
	user := models.User{
		Name:     names,
		Username: username,
		Password: hash,
		Added:    time.Now(),
	}
	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
