package main

import (
	"advanced-go-platzi/cache"
	"advanced-go-platzi/sync"
)

func main() {
	sync.DoSync()
	cache.DoCache()
}
