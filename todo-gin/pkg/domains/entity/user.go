package entity

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MockDataUser() []User {
	return []User{
		{
			ID:        1,
			Username:  "admin",
			Name:      "Admin",
			Email:     "admin@contag.id",
			Password:  "admin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Username:  "user",
			Name:      "User",
			Email:     "user@contag.id",
			Password:  "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
