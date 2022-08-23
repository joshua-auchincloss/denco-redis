package goserver

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/joshua-auchincloss/denco-redis/goserver/utils"
	"github.com/go-redis/redis/v9"
	"github.com/naoina/denco"
)

func Index(w http.ResponseWriter, r *http.Request, params denco.Params) {
	fmt.Fprintf(w, "Redis + Denco!\n")
}

// gets current site count and returns the value, setting if it does not exist and incrementing upon exit
func Count(client *redis.Client, ctx context.Context, ttl_s int, w http.ResponseWriter, r *http.Request, params denco.Params) {
	val, err := client.Get(ctx, "count").Int()

	switch {

	// unset keys return redis.Nil err, not a value
	case err == redis.Nil:
		fmt.Fprintf(w, "Does not exist... creating.\nset count = 1\n%vs TTL", ttl_s)
		// 1e+9 = s in ns
		client.Set(ctx, "count", 1, time.Duration(ttl_s*1e+9))

	// catch actual errors
	case err != nil:
		fmt.Fprintf(w, "Oh no ...\n%v", err)

	// compare to same type
	case val != 0:
		client.Incr(ctx, "count")
		fmt.Fprintf(w, "%v", val+1)
	}
}

// checks health of redis cluster
func checkRedisHealth(client *redis.Client, ctx context.Context) string {
	// test internal connection health
	log.Println("Client: PING")
	pong, err := client.Do(ctx, "PING").Result()
	if err != nil {
		panic(err)
	}
	log.Println("Redis:", pong)
	return fmt.Sprintf("%v", pong)
}


func main() {
	var ctx = context.Background()
	// specified in docker-compose
	redis_host := utils.getenv("REDIS_HOST", "localhost")
	redis_port := utils.getenv("REDIS_PORT", "6379")

	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")

	url := fmt.Sprintf("%v:%v", redis_host, redis_port)

	client := redis.NewClient(
		&redis.Options{
			Addr:     url,
			Password: "",
			DB:       0,
		})

	checkRedisHealth(client, ctx)

	mux := denco.NewMux()

	handler, err := mux.Build([]denco.Handler{
		mux.GET("/", Index),
		mux.GET("/ping", func(w http.ResponseWriter, r *http.Request, params denco.Params) {
			// 30 sec TTL, pass other vars from denco (writer, request, params)
			Count(client, ctx, 30, w, r, params)
		}),
		mux.GET("/shouldhide", func(w http.ResponseWriter, r *http.Request, params denco.Params){
			user, pass, ok := r.BasicAuth()
			fmt.Sprintf("COOKIES: %v, USER: %v, PASS: %v, OK: %v",r.Cookies(), user, pass, ok)
			w.Header().Set("WWW-Authenticate", "Invalid Credentials")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Sprintf("RESPONSE: %v", w.Header())

		  }),
	})

	if err != nil {
		panic(err)
	}
	// specified in docker-compose
	port := utils.getenv("DENCO_PORT", "8080")
	serve := fmt.Sprintf(":%v", port)

	srv := &http.Server{
		Addr:    serve,
		Handler: handler,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	log.Fatal(srv.ListenAndServeTLS(*certFile, *keyFile))
}
