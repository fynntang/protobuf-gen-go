package entity

import (
	"bitbucket.org/antalphadev/earn/pkg/dictionary"
	pkgRedis "bitbucket.org/antalphadev/earn/pkg/redis"
	"bitbucket.org/antalphadev/earn/pkg/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type BaseEntity struct {
	ID        ID        `gorm:"column:id;primary_key"`
	CreatedAt time.Time `gorm:"<-:create;column:created_at"`
	CreatedBy string    `gorm:"<-:create;column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type BaseTenantEntity struct {
	BaseEntity
	TenantID string `gorm:"column:tenant_id;index"`
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	if b.ID == 0 {
		b.ID = b.ID.FromInt64(utils.Snowflake().Generate().Int64())
	}
	b.CreatedAt = time.Now()
	ctx := tx.Statement.Context
	log.Printf("BeforeCreate ctx: %#v\n", ctx)
	if uid, ok := ctx.Value(dictionary.UserFieldUserID).(ID); ok {
		log.Printf("BeforeCreate uid: %#v\n", uid)
		b.CreatedBy = uid.String()
		b.UpdatedBy = b.CreatedBy
	}
	b.UpdatedAt = b.CreatedAt

	return nil
}

func (b *BaseEntity) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	ctx := tx.Statement.Context
	if uid, ok := ctx.Value(dictionary.UserFieldUserID).(ID); ok {
		tx.Statement.SetColumn("updated_by", uid.String())
	}

	return nil
}

func (b *BaseEntity) AfterFind(tx *gorm.DB) error {
	users := userNicknameCache(tx.Statement.Context)

	for id, nickname := range users {
		if b.CreatedBy == id {
			b.CreatedBy = nickname
		}
		if b.UpdatedBy == id {
			b.UpdatedBy = nickname
		}
	}

	return nil
}

func userNicknameCache(ctx context.Context) map[string]string {

	cacheKey := "dictionary:users"
	var userCache string
	var users = make(map[string]string)

	cache := pkgRedis.GetRedis().Get(ctx, cacheKey)
	if err := cache.Err(); err != nil && !errors.Is(err, redis.Nil) {
		return nil
	}
	if cache.Val() != "" && cache.Val() != "{}" {
		_ = json.Unmarshal([]byte(cache.Val()), &users)
		return users
	} else {
		var result []map[string]interface{}

		if db, ok := ctx.Value(dictionary.ContextDB).(*gorm.DB); ok {
			db.Table("users").Select([]string{"id", "nickname"}).Find(&result)
		}

		for _, item := range result {
			users[strconv.FormatInt(item["id"].(int64), 10)] = item["nickname"].(string)
		}
		usersMarshal, _ := json.Marshal(users)
		userCache = string(usersMarshal)

		pkgRedis.GetRedis().Set(ctx, cacheKey, userCache, time.Hour)
	}

	return users
}
