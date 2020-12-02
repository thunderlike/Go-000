package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

var dafaultName = "notexits"

// ServiceInterface interface
type ServiceInterface interface {
	GetUserNameByID(userID int) (name string, err error)
}

// ServerClient init server
type ServerClient struct {
	dbInstall *Connect
}

// NewServerClient server client
func NewServerClient(dbInstall *Connect) ServiceInterface {
	return &ServerClient{dbInstall}
}

// GetUserNameByID service get user info by userID
func (c *ServerClient) GetUserNameByID(userID int) (name string, err error) {
	stmt, err := c.dbInstall.db.Prepare("select name from user_test where id = ?")
	if err != nil {
		return dafaultName, errors.Wrap(err, "GetUserNameByID function Prepare Error")
	}
	defer stmt.Close()
	err = stmt.QueryRow(userID).Scan(&name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return dafaultName, nil
	} else {
		return dafaultName, errors.Wrap(err, "GetUserNameByID function QueryRow Error")
	}
	return name, nil
}
