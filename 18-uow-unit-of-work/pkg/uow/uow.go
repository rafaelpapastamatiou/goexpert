package uow

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNoTransactionStarted      = errors.New("no transaction started")
	ErrTransactionAlreadyStarted = errors.New("transaction already started")
	ErrRepositoryNotFound        = errors.New("repository not found")
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UnitOfWorkInterface interface {
	Register(name string, fc RepositoryFactory)
	Unregister(name string)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func() error) error
	Commit() error
	Rollback() error
}

type UnitOfWork struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
	committed    bool
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (uow *UnitOfWork) Register(name string, fc RepositoryFactory) {
	uow.Repositories[name] = fc
}

func (uow *UnitOfWork) Unregister(name string) {
	delete(uow.Repositories, name)
}

func (uow *UnitOfWork) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if uow.Tx == nil {
		tx, err := uow.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		uow.Tx = tx
	}

	repo, ok := uow.Repositories[name]
	if !ok {
		return nil, ErrRepositoryNotFound
	}

	return repo(uow.Tx), nil
}

func (uow *UnitOfWork) Do(ctx context.Context, fn func() error) error {
	if uow.Tx != nil {
		return ErrTransactionAlreadyStarted
	}

	tx, err := uow.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	uow.Tx = tx
	defer uow.Tx.Rollback()

	err = fn()
	if err != nil {
		return err
	}

	return uow.Tx.Commit()
}

func (uow *UnitOfWork) Rollback() error {
	if uow.Tx == nil && uow.committed {
		uow.committed = false
		return nil
	}

	if uow.Tx == nil {
		return ErrNoTransactionStarted
	}

	err := uow.Tx.Rollback()
	if err != nil {
		return err
	}

	uow.Tx = nil
	return nil
}

func (uow *UnitOfWork) Commit() error {
	if uow.Tx == nil {
		return ErrNoTransactionStarted
	}

	err := uow.Tx.Commit()
	if err != nil {
		return err
	}

	uow.Tx = nil
	uow.committed = true
	return nil
}

var _ UnitOfWorkInterface = &UnitOfWork{}
