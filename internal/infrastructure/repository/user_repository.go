package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
)

var (
	ErrUserAlreadyExists = errors.New("user alredy exists")
	ErrUserNotFound      = errors.New("user not found")
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *entity.User) (*int64, error) {
	if r.ExistsByEmail(user.Email) {
		return nil, ErrUserAlreadyExists
	}

	stmt := `INSERT INTO users (name, email, password, modified_at) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(stmt, user.Name, user.Email, user.Password, user.ModifiedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *UserRepository) FindAll(filters *Filters) ([]entity.User, error) {
	stmt := fmt.Sprintf("SELECT * FROM users ORDER BY id %s", filters.Order)

	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.ModifiedAt, &user.Deleted, &user.LastLogin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindById(id int64) (*entity.User, error) {
	var user entity.User
	stmt := `SELECT * FROM users WHERE id = ? AND deleted = ?`
	err := r.db.QueryRow(stmt, id, 0).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.ModifiedAt, &user.Deleted, &user.LastLogin)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	if r.ExistsByEmail(user.Email) {
		return ErrUserAlreadyExists
	}

	user.ModifiedAt = time.Now()
	stmt := `UPDATE users SET (name = ?, email = ?, password = ?, modified_at = ? WHERE id = ?)`
	_, err := r.db.Exec(stmt, user.Name, user.Email, user.Password, user.ModifiedAt, user.ID)
	return err
}

func (r *UserRepository) Delete(id int64) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	stmt := `UPDATE users SET deleted = 1 WHERE id = ?`
	_, err = r.db.Exec(stmt, id)
	return err
}

func (r *UserRepository) ExistsByEmail(email string) bool {
	var id *int64
	stmt := `SELECT id FROM users WHERE email = ?`
	r.db.QueryRow(stmt, email).Scan(&id)
	return id != nil
}
