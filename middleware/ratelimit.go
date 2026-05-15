package middleware

import (
	"app/pkg/resp"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*bucket
	rate    int
	window  time.Duration
}

type bucket struct {
	tokens   int
	lastSeen time.Time
}

func RateLimitMiddleware(rate int, window time.Duration) gin.HandlerFunc {
	rl := &rateLimiter{
		buckets: make(map[string]*bucket),
		rate:    rate,
		window:  window,
	}

	go func() {
		for {
			time.Sleep(window)
			rl.mu.Lock()
			for ip, b := range rl.buckets {
				if time.Since(b.lastSeen) > window {
					delete(rl.buckets, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		rl.mu.Lock()
		b, exists := rl.buckets[ip]
		if !exists {
			b = &bucket{tokens: rate - 1, lastSeen: time.Now()}
			rl.buckets[ip] = b
			rl.mu.Unlock()
			c.Next()
			return
		}

		now := time.Now()
		elapsed := now.Sub(b.lastSeen)
		b.tokens += int(elapsed.Seconds() * float64(rate) / window.Seconds())
		if b.tokens > rate {
			b.tokens = rate
		}

		if b.tokens <= 0 {
			rl.mu.Unlock()
			c.JSON(http.StatusTooManyRequests, resp.Output(resp.RESP_FAIL, nil, "Too many requests"))
			c.Abort()
			return
		}

		b.tokens--
		b.lastSeen = now
		rl.mu.Unlock()
		c.Next()
	}
}
