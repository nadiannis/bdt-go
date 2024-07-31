package main

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type User struct {
	ID       int64
	Username string
	Age      int
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() ([]*User, error) {
	query := "SELECT * FROM users"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Username, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) Insert(user *User) error {
	query := "INSERT INTO users (username, age) VALUES ($1, $2) RETURNING id"

	args := []any{user.Username, user.Age}
	
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return r.db.QueryRow(query, args...).Scan(&user.ID)
}

func (r *UserRepository) GetByID(id int64) (*User, error) {
	query := "SELECT id, username, age FROM users WHERE id = $1"

	var user User

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Age)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r *UserRepository) Update(user *User) error {
	query := `
						UPDATE users
						SET username = $1, age = $2
						WHERE id = $3`
	
	args := []any{user.Username, user.Age, user.ID}

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteByID(id int64) error {
	query := "DELETE FROM users WHERE id = $1"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
