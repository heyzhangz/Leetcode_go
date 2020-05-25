package question146

import "testing"

func contain(arr []string, key string) bool {

	for _, a := range arr {
		if a == key {
			return true
		}
	}

	return false
}

func TestQ146(t *testing.T) {

	//test mine
	print("run test mine\n")
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	println(cache.Get(1))
	cache.Put(3, 3)
	cache.Put(3, 5)
	cache.Put(3, 4)
	println(cache.Get(2))
	println(cache.Get(1))
	println(cache.Get(3))
	println(cache.Get(4))
	println(cache.Get(5))
}
