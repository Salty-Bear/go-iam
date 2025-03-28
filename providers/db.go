package providers

import (
	"fmt"

	"github.com/melvinodsa/go-iam/config"
	"github.com/melvinodsa/go-iam/db"
)

func NewDBConnection(cnf config.AppConfig) (db.DB, error) {
	conn, err := db.NewMongoConnection(cnf.DB.Host())
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}
	return conn, nil
}
