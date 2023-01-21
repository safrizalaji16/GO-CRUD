package domain

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel
	ID       int64  `bun:"id,pk,autoincrement" json:"id"`
	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}
