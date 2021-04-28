package middleware

import (
	"fmt"
	"github.com/go-redis/redis"
	"lingye-gin/src/config"
	"math/rand"
	"time"
)

var RedisClient *redis.Client

type RedisPool struct{}

func (p RedisPool) Init() {
	config.Logger.Info("RedisPool Init")
	// 连接服务器
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppProps.Redis.Addr,
		Password: config.AppProps.Redis.Passwd,
		DB:       config.AppProps.Redis.Database,
	})

	// 心跳
	pong, err := RedisClient.Ping().Result()

	// Output: PONG <nil>
	config.Logger.Info("RedisPool Ping And ", pong, err)
	config.Logger.Info("RedisPool Ok")
}

func (p RedisPool) Test() {
	ExampleClient_String()
	ExampleClient_List()
	ExampleClient_Hash()
	ExampleClient_Set()
	ExampleClient_SortSet()
	ExampleClient_HyperLogLog()
	ExampleClient_CMD()
	ExampleClient_Scan()
	ExampleClient_Tx()
	ExampleClient_Script()
	ExampleClient_PubSub()
}

func ExampleClient_String() {
	config.Logger.Println("ExampleClient_String")
	defer config.Logger.Println("ExampleClient_String")

	//kv读写
	err := RedisClient.Set("key", "value", 1*time.Second).Err()
	config.Logger.Println(err)

	//获取过期时间
	tm, err := RedisClient.TTL("key").Result()
	config.Logger.Println(tm)

	val, err := RedisClient.Get("key").Result()
	config.Logger.Println(val, err)

	val2, err := RedisClient.Get("missing_key").Result()
	if err == redis.Nil {
		config.Logger.Println("missing_key does not exist")
	} else if err != nil {
		config.Logger.Println("missing_key", val2, err)
	}

	//不存在才设置 过期时间 nx ex
	value, err := RedisClient.SetNX("counter", 0, 1*time.Second).Result()
	config.Logger.Println("setnx", value, err)

	//Incr
	result, err := RedisClient.Incr("counter").Result()
	config.Logger.Println("Incr", result, err)
}

func ExampleClient_List() {
	config.Logger.Println("ExampleClient_List")
	defer config.Logger.Println("ExampleClient_List")

	//添加
	config.Logger.Println(RedisClient.RPush("list_test", "message1").Err())
	config.Logger.Println(RedisClient.RPush("list_test", "message2").Err())

	//设置
	config.Logger.Println(RedisClient.LSet("list_test", 2, "message set").Err())

	//remove
	ret, err := RedisClient.LRem("list_test", 3, "message1").Result()
	config.Logger.Println(ret, err)

	rLen, err := RedisClient.LLen("list_test").Result()
	config.Logger.Println(rLen, err)

	//遍历
	lists, err := RedisClient.LRange("list_test", 0, rLen-1).Result()
	config.Logger.Println("LRange", lists, err)

	//pop没有时阻塞
	result, err := RedisClient.BLPop(1*time.Second, "list_test").Result()
	config.Logger.Println("result:", result, err, len(result))
}

func ExampleClient_Hash() {
	config.Logger.Println("ExampleClient_Hash")
	defer config.Logger.Println("ExampleClient_Hash")

	datas := map[string]interface{}{
		"name": "LI LEI",
		"sex":  1,
		"age":  28,
		"tel":  123445578,
	}

	//添加
	if err := RedisClient.HMSet("hash_test", datas).Err(); err != nil {
		config.Logger.Fatal(err)
	}

	//获取
	rets, err := RedisClient.HMGet("hash_test", "name", "sex").Result()
	config.Logger.Println("rets:", rets, err)

	//成员
	retAll, err := RedisClient.HGetAll("hash_test").Result()
	config.Logger.Println("retAll", retAll, err)

	//存在
	bExist, err := RedisClient.HExists("hash_test", "tel").Result()
	config.Logger.Println(bExist, err)

	bRet, err := RedisClient.HSetNX("hash_test", "id", 100).Result()
	config.Logger.Println(bRet, err)

	//删除
	config.Logger.Println(RedisClient.HDel("hash_test", "age").Result())
}

func ExampleClient_Set() {
	config.Logger.Println("ExampleClient_Set")
	defer config.Logger.Println("ExampleClient_Set")

	//添加
	ret, err := RedisClient.SAdd("set_test", "11", "22", "33", "44").Result()
	config.Logger.Println(ret, err)

	//数量
	count, err := RedisClient.SCard("set_test").Result()
	config.Logger.Println(count, err)

	//删除
	ret, err = RedisClient.SRem("set_test", "11", "22").Result()
	config.Logger.Println(ret, err)

	//成员
	members, err := RedisClient.SMembers("set_test").Result()
	config.Logger.Println(members, err)

	bret, err := RedisClient.SIsMember("set_test", "33").Result()
	config.Logger.Println(bret, err)

	RedisClient.SAdd("set_a", "11", "22", "33", "44")
	RedisClient.SAdd("set_b", "11", "22", "33", "55", "66", "77")
	//差集
	diff, err := RedisClient.SDiff("set_a", "set_b").Result()
	config.Logger.Println(diff, err)

	//交集
	inter, err := RedisClient.SInter("set_a", "set_b").Result()
	config.Logger.Println(inter, err)

	//并集
	union, err := RedisClient.SUnion("set_a", "set_b").Result()
	config.Logger.Println(union, err)

	ret, err = RedisClient.SDiffStore("set_diff", "set_a", "set_b").Result()
	config.Logger.Println(ret, err)

	rets, err := RedisClient.SMembers("set_diff").Result()
	config.Logger.Println(rets, err)
}

func ExampleClient_SortSet() {
	config.Logger.Println("ExampleClient_SortSet")
	defer config.Logger.Println("ExampleClient_SortSet")

	addArgs := make([]redis.Z, 100)
	for i := 1; i < 100; i++ {
		addArgs = append(addArgs, redis.Z{Score: float64(i), Member: fmt.Sprintf("a_%d", i)})
	}
	//Logger.Println(addArgs)

	Shuffle := func(slice []redis.Z) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for len(slice) > 0 {
			n := len(slice)
			randIndex := r.Intn(n)
			slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
			slice = slice[:n-1]
		}
	}

	//随机打乱
	Shuffle(addArgs)

	//添加
	ret, err := RedisClient.ZAddNX("sortset_test", addArgs...).Result()
	config.Logger.Println(ret, err)

	//获取指定成员score
	score, err := RedisClient.ZScore("sortset_test", "a_10").Result()
	config.Logger.Println(score, err)

	//获取制定成员的索引
	index, err := RedisClient.ZRank("sortset_test", "a_50").Result()
	config.Logger.Println(index, err)

	count, err := RedisClient.SCard("sortset_test").Result()
	config.Logger.Println(count, err)

	//返回有序集合指定区间内的成员
	rets, err := RedisClient.ZRange("sortset_test", 10, 20).Result()
	config.Logger.Println(rets, err)

	//返回有序集合指定区间内的成员分数从高到低
	rets, err = RedisClient.ZRevRange("sortset_test", 10, 20).Result()
	config.Logger.Println(rets, err)

	//指定分数区间的成员列表
	rets, err = RedisClient.ZRangeByScore("sortset_test", redis.ZRangeBy{Min: "(30", Max: "(50", Offset: 1, Count: 10}).Result()
	config.Logger.Println(rets, err)
}

//用来做基数统计的算法，HyperLogLog 的优点是，在输入元素的数量或者体积非常非常大时，计算基数所需的空间总是固定 的、并且是很小的。
//每个 HyperLogLog 键只需要花费 12 KB 内存，就可以计算接近 2^64 个不同元素的基 数
func ExampleClient_HyperLogLog() {
	config.Logger.Println("ExampleClient_HyperLogLog")
	defer config.Logger.Println("ExampleClient_HyperLogLog")

	for i := 0; i < 10000; i++ {
		RedisClient.PFAdd("pf_test_1", fmt.Sprintf("pfkey%d", i))
	}
	ret, err := RedisClient.PFCount("pf_test_1").Result()
	config.Logger.Println(ret, err)

	for i := 0; i < 10000; i++ {
		RedisClient.PFAdd("pf_test_2", fmt.Sprintf("pfkey%d", i))
	}
	ret, err = RedisClient.PFCount("pf_test_2").Result()
	config.Logger.Println(ret, err)

	RedisClient.PFMerge("pf_test", "pf_test_2", "pf_test_1")
	ret, err = RedisClient.PFCount("pf_test").Result()
	config.Logger.Println(ret, err)
}

func ExampleClient_PubSub() {
	config.Logger.Println("ExampleClient_PubSub")
	defer config.Logger.Println("ExampleClient_PubSub")
	//发布订阅
	pubsub := RedisClient.Subscribe("subkey")
	_, err := pubsub.Receive()
	if err != nil {
		config.Logger.Fatal("pubsub.Receive")
	}
	ch := pubsub.Channel()
	time.AfterFunc(1*time.Second, func() {
		config.Logger.Println("Publish")

		err = RedisClient.Publish("subkey", "test publish 1").Err()
		if err != nil {
			config.Logger.Fatal("RedisClient.Publish", err)
		}

		RedisClient.Publish("subkey", "test publish 2")
	})
	for msg := range ch {
		config.Logger.Println("recv channel:", msg.Channel, msg.Pattern, msg.Payload)
	}
}

func ExampleClient_CMD() {
	config.Logger.Println("ExampleClient_CMD")
	defer config.Logger.Println("ExampleClient_CMD")

	//执行自定义redis命令
	Get := func(rdb *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		RedisClient.Process(cmd)
		return cmd
	}

	v, err := Get(RedisClient, "NewStringCmd").Result()
	config.Logger.Println("NewStringCmd", v, err)

	v, err = RedisClient.Do("get", "RedisClient.do").String()
	config.Logger.Println("RedisClient.Do", v, err)
}

func ExampleClient_Scan() {
	config.Logger.Println("ExampleClient_Scan")
	defer config.Logger.Println("ExampleClient_Scan")

	//scan
	for i := 1; i < 1000; i++ {
		RedisClient.Set(fmt.Sprintf("skey_%d", i), i, 0)
	}

	cusor := uint64(0)
	for {
		keys, retCusor, err := RedisClient.Scan(cusor, "skey_*", int64(100)).Result()
		config.Logger.Println(keys, cusor, err)
		cusor = retCusor
		if cusor == 0 {
			break
		}
	}
}

func ExampleClient_Tx() {
	pipe := RedisClient.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	// Execute
	//
	//     MULTI
	//     INCR pipeline_counter
	//     EXPIRE pipeline_counts 3600
	//     EXEC
	//
	// using one rdb-server roundtrip.
	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

func ExampleClient_Script() {
	IncrByXX := redis.NewScript(`
        if redis.call("GET", KEYS[1]) ~= false then
            return redis.call("INCRBY", KEYS[1], ARGV[1])
        end
        return false
    `)

	n, err := IncrByXX.Run(RedisClient, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = RedisClient.Set("xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByXX.Run(RedisClient, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)
}
