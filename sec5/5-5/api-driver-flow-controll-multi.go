package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter{
	// 可変長のRateLimiterを受け取って、その機能を集約したmultiLimiterを返す
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit) // リミットが短い順にソート
	// sortの第一引数が対象オブジェクトのリスト、第二引数は関数で、各要素をとって順序関係を返す
	return &multiLimiter{limiters: limiters}
	// limiters（スライス）を要素にもつmultiLimiterオブジェクトのポインタを返す
}

type multiLimiter struct {
	limiters []RateLimiter
}

// メソッドの定義
func (l *multiLimiter) Wait(ctx context.Context) error {
	// ぜんぶ通ればブロックされない
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
	// 一番厳しいlimitを返す
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration/time.Duration(eventCount))
}

func Open() *APIConnection {
	secondLimit := rate.NewLimiter(Per(2, time.Second), 1)
	minuteLimit := rate.NewLimiter(Per(2, time.Minute), 10)
	return &APIConnection{
		rateLimiter: MultiLimiter(secondLimit, minuteLimit),
	}
}

type APIConnection struct {
	rateLimiter RateLimiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { 
		return err // ctxがキャンセルされたらエラーが返る
	}
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { 
		return err // ctxがキャンセルされたらエラーが返る
	}
	return nil
}

func main() {
	defer fmt.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}