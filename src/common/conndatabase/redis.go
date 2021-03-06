package conndatabase
import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)
//设置缓存key
var (
	LocationGeo = "LocationGeo"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := beego.AppConfig.String("redis.host")
	Pool = newPool(redisHost)
}

/*

MaxIdle：最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。

MaxActive：最大的激活连接数，表示同时最多有N个连接

IdleTimeout：最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭

Dial：建立连接

// 从池里获取连接
rc := Pool.Get()
// 用完后将连接放回连接池
defer rc.Close()

*/

type RedisServer struct {

}

func newPool(server string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 45),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 50),
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
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

func poolClose() {
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

func DealConn(comm string,key string) ([]byte,error) {
		conn := Pool.Get()
		defer conn.Close()
		var data []byte
		data, err := redis.Bytes(conn.Do(comm, key))
		if err != nil {
			return nil, err
		}
	   return data,err
}


