package appcompatcache

import "time"

type CacheEntry struct {
	Name  string    `json:"name"`
	Epoch uint64    `json:"epoch"`
	Time  time.Time `json:"time"`
}
