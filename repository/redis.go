package repository

import (
	"github.com/go-redis/redis"
	"shared"
	"time"
)

//Редис Репозиторий
type Redis struct{
	Client *redis.Client
	Config *shared.RedisCfg
}

func (r *Redis) Open() (error) {
	r.Client = redis.NewClient(&redis.Options{
		Addr: r.Config.HostAndPort(),
		DB:   r.Config.Db,
	})
	_, err := r.Client.Ping().Result()
	return err
}

func (r *Redis) Get() (shared.State){
	val, err := r.Client.Get(r.Config.Key).Result()
	if err != nil {
		panic(err)
	}
	v := shared.FromJSON(val)
	return v
}

func (r *Redis) Put(state shared.State) (error){
	err := r.Client.Set(r.Config.Key, state.ToJSON(), 0).Err()
	time.Sleep(100 * time.Millisecond)
	return err
}

func (r *Redis) Close(){
	r.Client.Close()
}

func NewRedis(config *shared.RedisCfg) (*Redis){
	r := Redis{Config: config}
	return &r;
}