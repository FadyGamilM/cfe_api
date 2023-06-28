package services

import (
	"database/sql"
	"time"

	"github.com/FadyGamilM/cfe_api/core/domain"
)

// limit all db operations for 3 seconds only
const dbTimeOut = 3 * time.Second

// define the connection pool to apply all db operations using it
var db *sql.DB

type Models struct {
	coffee domain.Coffee
}
