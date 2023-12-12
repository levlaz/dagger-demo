package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/launchdarkly/go-sdk-common/v3/ldcontext"
	ld "github.com/launchdarkly/go-server-sdk/v7"
	"rsc.io/quote/v4"
)

var client, _ = ld.MakeClient(os.Getenv("LD_SDK_KEY"), 5*time.Second)

func GetGreeting() string {
	var greeting string
	flagKey := "greeting-flipper"
	context := ldcontext.NewBuilder(uuid.New().String()).Anonymous(true).Build()

	value, detail, err := client.BoolVariationDetail(flagKey, context, false)
	if err != nil {
		panic(err)
	}
	if value {
		greeting = quote.Glass()
	} else {
		greeting = quote.Go()
	}

	fmt.Println(detail)
	fmt.Println(context)

	return greeting
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(GetGreeting()))
	})
	http.ListenAndServe(":3000", r)
}
