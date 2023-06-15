package entities

type Auth struct {
	entity.BaseEntity
}

func (a *Auth) TableName() string { return "auths" }
