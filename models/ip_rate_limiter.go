package models

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// IPRateLimiter .
type IPRateLimiter struct {
	ips map[string]*rateLimitRef
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

type rateLimitRef struct {
	rateLimiter *rate.Limiter
	lastTime    time.Time
}

// NewIPRateLimiter .
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rateLimitRef),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	go func() {
		for ; true; <-ticker.C {
			fmt.Println("IP remover Tick at", time.Now())

			for key, lim := range i.ips {
				if time.Now().Sub(lim.lastTime).Minutes() > 5 {
					fmt.Println("key removed: " + key)
					delete(i.ips, key)
				}
			}
		}
	}()

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := &rateLimitRef{rate.NewLimiter(i.r, i.b), time.Now()}

	i.ips[ip] = limiter

	return limiter.rateLimiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	limiter.lastTime = time.Now()
	i.mu.Unlock()

	return limiter.rateLimiter
}

func (i *IPRateLimiter) GetLimiterIP(ip string) *rate.Limiter {
	ipAddress := strings.Split(ip, ":")[0]

	return i.GetLimiter(ipAddress)
}

func (i *IPRateLimiter) IsAllowed(ip string) bool {
	ipAddress := strings.Split(ip, ":")[0]

	rateLimiter := i.GetLimiter(ipAddress)

	if rateLimiter != nil && !rateLimiter.Allow() {
		return false
	}

	return true
}
