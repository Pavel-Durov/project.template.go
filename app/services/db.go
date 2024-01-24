package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"p3ld3v.dev/template/app/services/db/sqlc"
)

type DbStore interface {
	GetUser(id int64) (*sqlc.User, error)
	CreateUser(name string) (*sqlc.User, error)
	Connect() error
}

type SQLStore struct {
	query         *sqlc.Queries
	pool          *pgxpool.Pool
	connectionUrl string
	logger        Logger
}

func NewDbService(connectionUrl string, logger Logger) (DbStore, error) {
	return &SQLStore{
		connectionUrl: connectionUrl,
		logger:        logger,
		pool:          nil,
	}, nil
}

func (db *SQLStore) Connect() error {
	if db.pool != nil {
		return nil
	}
	config, err := pgxpool.ParseConfig(db.connectionUrl)
	if err != nil {
		return err
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		db.logger.Info("Connected to database")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}
	db.pool = pool
	db.query = sqlc.New(db.pool)
	return nil
}

func (db *SQLStore) GetUser(id int64) (*sqlc.User, error) {
	ctx := context.Background()
	user, err := db.query.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *SQLStore) CreateUser(name string) (*sqlc.User, error) {
	ctx := context.Background()
	user, err := db.query.AddUser(ctx, name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
