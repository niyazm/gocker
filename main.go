package main
import (
	"fmt"
	"net/http"
	"html"
	"log"
	"context"
	"github.com/go-redis/redis/v8"
	)

var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}

func main() {
	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world, from badonkadocker %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/redisTest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "doing a redis test %q", html.EscapeString(r.URL.Path))
		ExampleClient()
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you might be looking for /heartbeat %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
