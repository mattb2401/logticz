package pkg

import (
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/log"
	"github.com/logticz/logticz/internal/store/models"
	"github.com/logticz/logticz/lib"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"

	"os"
	"testing"

	"gorm.io/gorm"
)

//TestAuthentication_HandleLogin tests HandleLogin with mock
func TestAuthentication_HandleLogin(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		assert.Error(t, err)
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sqldb,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		assert.Error(t, err)
	}
	pass, err := lib.HashArgonPassword("testpassword")
	if err != nil {
		assert.Error(t, err)
	}
	user := models.User{
		ID:       1,
		Name:     "Matt",
		Username: "testuser",
		Password: pass,
		Added:    time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "added"}).AddRow(user.ID, user.Name, user.Username, user.Password, user.Added)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1 ORDER BY "users"."id" LIMIT 1`)).WithArgs("testuser").WillReturnRows(rows)
	authenticationSvc := NewAuthenticationSvc(&logger, db)
	response, err := authenticationSvc.HandleLogin("testuser", "testpassword")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, response.Username, user.Username)
	assert.Equal(t, response.Name, user.Name)
}

func TestAuthentication_HandleRegistration(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		assert.Error(t, err)
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sqldb,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		assert.Error(t, err)
	}
	pass, err := lib.HashArgonPassword("testpassword")
	if err != nil {
		assert.Error(t, err)
	}
	user := models.User{
		ID:       1,
		Name:     "Matt",
		Username: "testuser",
		Password: pass,
		Added:    time.Now(),
	}
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("name", "username", "password", "added") VALUES ($1, $2, $3, $4) RETURNING "id"`)).WithArgs(user.Name, user.Username, user.Password, user.Added)
	authenticationSvc := NewAuthenticationSvc(&logger, db)
	response, err := authenticationSvc.HandleRegistration(user.Name, user.Username, user.Password)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, response.Username, user.Username)
	assert.Equal(t, response.Name, user.Name)
}
