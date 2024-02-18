package postgres

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/spf13/viper/remote"
)

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {
	envVars := map[string]string{
		"PGHOST":     "postgres",
		"PGPORT":     "5432",
		"PGDATABASE": "postgres",
		"PGUSER":     "postgres",
		"PGPASSWORD": "password",
		"PGSSLMODE":  "disable",
	}

	for key, value := range envVars {
		if len(os.Getenv(key)) == 0 {
			if err := os.Setenv(key, value); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

type Store struct {
	Pool *pgxpool.Pool
}

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}

func (s Settings) toDSN() string {
	var args []string

	addArg := func(key, value string) {
		if len(value) > 0 {
			args = append(args, fmt.Sprintf("%s=%s", key, value))
		}
	}

	addArg("host", s.Host)
	addArg("port", fmt.Sprintf("%d", s.Port))
	addArg("dbname", s.Database)
	addArg("user", s.User)
	addArg("password", s.Password)
	addArg("sslmode", s.SSLMode)

	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {
	config, err := pgxpool.ParseConfig(settings.toDSN())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{Pool: conn}, nil
}
