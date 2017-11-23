package myredis

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func Print(){
	beego.Debug("go,redis")
}

func init() {
	beego.Debug("初始化redis")
	redisHost := beego.AppConfig.String("redis.ip") + ":6379"
	Pool = newPool(redisHost)
	close()
}
func newPool(server string) *redis.Pool {
	database, _ := beego.AppConfig.Int("redis.database")
	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server, redis.DialDatabase(database), redis.DialPassword(beego.AppConfig.String("redis.password")))
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) (string, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data string
	data, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}