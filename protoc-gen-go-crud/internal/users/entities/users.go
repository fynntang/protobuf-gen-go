package entities

type Users struct {
	entity.BaseEntity
}

func (u *Users) TableName() string { return "userss" }
