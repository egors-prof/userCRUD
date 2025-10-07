package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)
type Cache struct{
	rdb * redis.Client
	cacheLog zerolog.Logger
}

func NewCache(rdb*redis.Client,logger zerolog.Logger)*Cache{
	return &Cache{
		rdb:rdb,
		cacheLog:logger,
	}
}

func (c *Cache)Set(ctx context.Context,key string ,value interface{},dur time.Duration)error{
	rawB,err:=json.Marshal(value)
	if err!=nil{
		c.cacheLog.Error().Err(err).Msg("error while setting value in cache")
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
	if errors.Is(err,redis.Nil){
		c.cacheLog.Error().Err(err).Msg("no passed id in cache")
		return "",err 
	}
	return result,nil
}