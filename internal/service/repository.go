package service

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) IsExistEmail(email string) (bool, error) {
	var count int
	query, args, err := squirrel.Select("COUNT(*)").From("subscribers").Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		return false, err
	}
	err = r.db.Get(&count, query, args...)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) AddSubscriber(email string) error {
	query, args, err := squirrel.Insert("subscribers").Columns("email").Values(email).ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (r *Repository) GetAllSubscribers() ([]string, error) {
	var emails []string
	query, args, err := squirrel.Select("email").From("subscribers").ToSql()
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&emails, query, args...)
	if err != nil {
		return nil, err
	}
	return emails, nil
}
