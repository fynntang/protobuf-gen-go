package entities

type User struct {
	entity.BaseEntity
}

func (u *User) TableName() string { return "users" }
