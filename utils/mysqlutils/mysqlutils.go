package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

const (
	ErrorNoRows         = "no rows in result set"
	ErrorRecordNotFound = "record not found"
)

func ParseError(err error, obj interface{}) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError) //type assertion
	if !ok {
		if strings.Contains(err.Error(), ErrorRecordNotFound) {
			return errors.NewNotFoundError(fmt.Sprintf("id '%v' not found", obj))
		}

	} else if sqlErr.Number == 1062 {
		return errors.NewConflictError(fmt.Sprintf("'%v' is already existed", obj))
	}

	return errors.NewInternalServerError(fmt.Sprintf(err.Error()))
}
