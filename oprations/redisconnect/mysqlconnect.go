package redisconnect

import "github.com/go-redis/redis/v8"

// 声明一个全局的redisDb变量
// var RedisDB *redis.Client

// 根据redis配置初始化一个客户端
func RedisConnect() (RedisDB *redis.Client, err error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	// 但是语法报错
	// _, err = RedisDB.Ping().Result()
	// if err != nil {
	// 	return nil, err
	// }
	return RedisDB, nil
}

