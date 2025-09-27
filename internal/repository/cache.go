package repository

import (
	"context"
	"encoding/json"
	"time"
	"github.com/redis/go-redis/v9"
)
type Cache struct{
	rdb * redis.Client
}

func NewCache(rdb*redis.Client)*Cache{
	return &Cache{rdb:rdb}
}

func (c *Cache)Set(ctx context.Context,key string ,value interface{},dur time.Duration)error{
	rawB,err:=json.Marshal(value)
	if err!=nil{
		return err
	}
	
	err=c.rdb.Set(ctx,key,rawB,dur).Err()
	if err!=nil{
		return err
	}
	return nil

}

func (c*Cache)Get(ctx context.Context,key string)(string ,error){
	result,err:=c.rdb.Get(ctx,key).Result()
	if err!=nil{
		return "",err 
	}
	return result,nil
}