package cache

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var cDb *cache.Cache

func InitCache() {
	cDb = cache.New(10*time.Minute, 60*time.Minute)
	log.Println("初始化缓存系统成功!")
}

func NewCache() *cache.Cache {
	return cDb
}
