package cache_test

import (
	"alexioannides/go-playpen/learn-go-with-pocket-sized-projects/06-caching-with-generics/cache"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testTTL = time.Millisecond * 100

/*
	TestCache is an integration test.

- Create a cache and check that it is empty?
- Upsert a new key <5, fünf> in the cache
- Read the value for this entry
- Upsert for the same entry with the new value
- Read the new value
- Upsert another key <3, drei> and check that is doesn't override
- Delete 5 and check that 3 still exists
*/
func TestCache(t *testing.T) {
	c := cache.New[int, string](testTTL)

	c.Upsert(5, "fünf")

	v, found := c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "fünf", v)

	c.Upsert(5, "pum")

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Upsert(3, "drei")

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)

	v, found = c.Read(5)
	assert.True(t, found)
	assert.Equal(t, "pum", v)

	c.Delete(5)

	v, found = c.Read(5)
	assert.False(t, found)
	assert.Equal(t, "", v)

	v, found = c.Read(3)
	assert.True(t, found)
	assert.Equal(t, "drei", v)
}

// TestCache_Parallel_goroutines simulates a number of parallel tasks each operating on the cache.
// It passes even when we use "go test -race .".
func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string](testTTL)

	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			// Perform one operation that alters the content of the cache in each goroutine.
			// The mutex prevents any race condition from happening.
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()
}

// TestCache_Parallel runs two goroutines that have concurrent access to write to the cache.
func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string](testTTL)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}

func TestCache_TTL(t *testing.T) {
	c := cache.New[string, string](testTTL)
	c.Upsert("Norwegian", "Blue")

	// can we find the value we just cached?
	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "Blue", got)

	// wait until TTL expires and try again
	time.Sleep(testTTL * 2)
	got, found = c.Read("Norwegian")
	assert.False(t, found)
	assert.Equal(t, "", got)

}
