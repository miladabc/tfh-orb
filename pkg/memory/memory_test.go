package memory

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemoryStore(t *testing.T) {
	t.Parallel()

	m := New[string, int]()

	key, value := "store", 123

	m.Store(key, value)
	v, exists := m.Get(key)
	require.True(t, exists)
	require.Equal(t, value, v)
}

func TestMemoryGet(t *testing.T) {
	t.Parallel()

	m := New[string, int]()

	key, value := "get", 123

	_, exists := m.Get(key)
	require.False(t, exists)

	m.Store(key, value)
	v, exists := m.Get(key)
	require.True(t, exists)
	require.Equal(t, value, v)
}

func TestMemoryStoreOverwrite(t *testing.T) {
	t.Parallel()

	m := New[string, int]()

	key, value := "overwrite", 123
	m.Store(key, value)

	newValue := 456
	m.Store(key, newValue)

	v, exists := m.Get(key)
	require.True(t, exists)
	require.Equal(t, newValue, v)
}

func TestConcurrency(t *testing.T) {
	t.Parallel()

	m := New[int, int]()

	const goroutines = 1000

	var wg sync.WaitGroup

	wg.Add(goroutines)

	for i := range goroutines {
		go func() {
			defer wg.Done()

			m.Store(i, i+1)
		}()
	}

	wg.Wait()

	require.Equal(t, goroutines, m.Len())
}
