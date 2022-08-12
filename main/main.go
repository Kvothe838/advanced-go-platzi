package main

import (
	"advanced-go-platzi/cache"
	"advanced-go-platzi/design_patterns"
	"advanced-go-platzi/sync"
)

func main() {
	sync.DoSync()
	cache.DoCache()
	cache.DoExpensiveFibo()
	design_patterns.DoFactory()
}
