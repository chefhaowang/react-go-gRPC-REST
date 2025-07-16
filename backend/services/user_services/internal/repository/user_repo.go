//CRUD of user respo
package repository

import (
	"database/sql"
	"log"
	"backend/services/user_services/internal/db"
)

// UserRepository handles all DB operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByEmail retrieves a user by their email from the database
func (r *UserRepository) GetUserByEmail(email string) (*db.User, error) {
	query := `SELECT id, email, password_hash, role, created_at FROM users WHERE email = ?`
	row := r.db.QueryRow(query, email)
	user := &db.User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		log.Printf("DB error: %v", err)
		return nil, err
	}
	return user, nil
}


// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id string) (*db.User, error) {
	query := `SELECT id, email, password_hash, role, created_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)
	user := &db.User{}
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		log.Printf("DB error: %v", err)
		return nil, err
	}
	return user, nil
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *db.User) error {
	query := `INSERT INTO users (id, email, password_hash, role) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.ID, user.Email, user.PasswordHash, user.Role)
	return err
}


// UpdateUser updates a user's email and role
func (r *UserRepository) UpdateUser(user *db.User) error {
	query := `UPDATE users SET email = ?, role = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Email, user.Role, user.ID)
	return err
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
