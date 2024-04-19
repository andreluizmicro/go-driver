package repository

import (
	"database/sql"
	"github.com/andreluizmicro/go-driver-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-driver-api/internal/domain/entity"
	"github.com/andreluizmicro/go-driver-api/internal/domain/exception"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/Presenter"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository/filter"
	"time"
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
		return nil, exception.ErrUserAlreadyExists
	}

	stmt := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	result, err := r.db.Exec(stmt, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *UserRepository) FindAll(filters *filter.Filters) (contracts.PaginateInterface, error) {
	stmt := "SELECT * FROM users WHERE deleted = 0 ORDER BY name LIMIT ?, ?"

	rows, err := r.db.Query(stmt, filters.PerPage, (filters.Page-1)*filters.PerPage)
	if err != nil {
		return &Presenter.PaginatePresenter{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var users []interface{}

	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.ModifiedAt,
			&user.Deleted,
			&user.LastLogin,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return &Presenter.PaginatePresenter{
		Data: users,
	}, nil
}

func (r *UserRepository) FindById(id int64) (*entity.User, error) {
	var user entity.User
	stmt := `SELECT * FROM users WHERE id = ? AND deleted = ?`
	err := r.db.QueryRow(stmt, id, 0).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.ModifiedAt,
		&user.Deleted,
		&user.LastLogin,
	)
	if err != nil {
		return nil, exception.ErrUserNotFound
	}
	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	if r.ExistsByEmail(user.Email) {
		return exception.ErrUserAlreadyExists
	}

	stmt := `UPDATE users SET name = ?, email = ?, modified_at = ? WHERE id = ?`
	_, err := r.db.Exec(stmt, user.Name, user.Email, time.Now(), user.ID)
	return err
}

func (r *UserRepository) Delete(id int64) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	if user == nil {
		return exception.ErrUserNotFound
	}

	stmt := `UPDATE users SET deleted = 1 WHERE id = ?`
	_, err = r.db.Exec(stmt, id)
	return err
}

func (r *UserRepository) ExistsByEmail(email string) bool {
	var id *int64
	stmt := `SELECT id FROM users WHERE email = ?`
	err := r.db.QueryRow(stmt, email).Scan(&id)
	if err != nil {
		return false
	}
	return id != nil
}
