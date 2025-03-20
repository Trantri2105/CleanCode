package main

import (
	"fmt"
)

type Repository interface {
	GetData(key string) string
}

type OriginalRepository struct{}

func (s *OriginalRepository) GetData(key string) string {
	fmt.Printf("OriginalRepository: Fetching data for key %s from database\n", key)
	// Fetching data from database
	return fmt.Sprintf("Data for key %s", key)
}

// CacheProxy implements the same interface but adds caching
type CacheProxy struct {
	repo  *OriginalRepository
	cache map[string]string
}

func NewCacheProxy(repo *OriginalRepository) *CacheProxy {
	return &CacheProxy{
		repo:  repo,
		cache: make(map[string]string),
	}
}

// GetData checks if the data is already in cache, if not fetches from original repository
func (p *CacheProxy) GetData(key string) string {
	if data, found := p.cache[key]; found {
		fmt.Printf("CacheProxy: Returning cached data for key %s\n", key)
		return data
	}

	data := p.repo.GetData(key)
	p.cache[key] = data
	return data
}

func main() {
	// Create the original repository
	service := &OriginalRepository{}

	// Create the proxy
	proxy := NewCacheProxy(service)

	fmt.Println("\n--- First request ---")
	fmt.Println(proxy.GetData("ABC"))

	fmt.Println("\n--- Second request (same key) ---")
	fmt.Println(proxy.GetData("ABC"))

	fmt.Println("\n--- Third request (different key) ---")
	fmt.Println(proxy.GetData("XYZ"))
}
