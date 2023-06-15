package entities

type Role struct {
	entity.BaseEntity
}

func (r *Role) TableName() string { return "roles" }
