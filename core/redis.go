package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
)

func (this *Common) Redisexec(redisHost, redisPassword, command string) error {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})
	var args []string
	if command == "" {
		if _, err := client.Ping().Result(); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		args = strings.Split(command, " ")
	}
	//result := client.Do("SMEMBERS", "letters")
	switch count := len(args); count {
	case 1:
		result := client.Do(args[0])
		resStr, err := result.Result()
		if err != nil {
			return err
		} else {
			fmt.Println(resStr)
			return nil
		}
	case 2:
		result := client.Do(args[0], args[1])
		resStr, err := result.Result()
		if err != nil {
			return err
		} else {
			fmt.Println(resStr)
			return nil
		}
	case 3:
		result := client.Do(args[0], args[1], args[2])
		resStr, err := result.Result()
		if err != nil {
			return err
		} else {
			fmt.Println(resStr)
			return nil
		}
	}
	return nil
}

//func InitRedisPool() (redis.Conn, error) {
//	conn, err := redis.Dial("tcp", "127.0.0.1:6380",
//		redis.DialConnectTimeout(time.Duration(2)*time.Second),
//		redis.DialPassword("de477b4a-ef25-7cf9-8098-ee9d1245dc7f"),
//		redis.DialDatabase(0),
//	)
//	if err != nil {
//		fmt.Println(err)
//		//log.Error(err)
//		fmt.Println(err)
//	}
//	return conn, err
//
//}
