package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var imageCache = cache.New(24*time.Hour, 1*time.Hour)

func cacheImage(url string, content []byte) {
	imageCache.Set(url, content, cache.DefaultExpiration)
}

func getImageCache(url string) ([]byte, bool) {
	cacheContent, ok := imageCache.Get(url)
	if !ok {
		return nil, false
	}

	return cacheContent.([]byte), true
}
