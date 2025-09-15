package generator

import (
	"strings"
	"sync"

	"github.com/tylerbu/ccl-test-lib/types"
)

// Pool manages reusable objects to reduce memory allocations
type Pool struct {
	testSuites  sync.Pool
	stringSlice sync.Pool
	stringMap   sync.Pool
}

// NewPool creates a new object pool for the generator
func NewPool() *Pool {
	return &Pool{
		testSuites: sync.Pool{
			New: func() interface{} {
				return &types.TestSuite{}
			},
		},
		stringSlice: sync.Pool{
			New: func() interface{} {
				return make([]string, 0, 16) // Pre-allocate capacity for common case
			},
		},
		stringMap: sync.Pool{
			New: func() interface{} {
				return make(map[string]interface{}, 8) // Pre-allocate capacity
			},
		},
	}
}

// GetTestSuite returns a reusable TestSuite from the pool
func (p *Pool) GetTestSuite() *types.TestSuite {
	suite := p.testSuites.Get().(*types.TestSuite)
	// Reset the suite for reuse
	*suite = types.TestSuite{}
	return suite
}

// PutTestSuite returns a TestSuite to the pool
func (p *Pool) PutTestSuite(suite *types.TestSuite) {
	if suite != nil {
		p.testSuites.Put(suite)
	}
}

// GetStringSlice returns a reusable string slice from the pool
func (p *Pool) GetStringSlice() []string {
	slice := p.stringSlice.Get().([]string)
	return slice[:0] // Reset length but keep capacity
}

// PutStringSlice returns a string slice to the pool
func (p *Pool) PutStringSlice(slice []string) {
	if slice != nil && cap(slice) <= 256 { // Avoid holding onto huge slices
		p.stringSlice.Put(slice)
	}
}

// GetStringMap returns a reusable string map from the pool
func (p *Pool) GetStringMap() map[string]interface{} {
	m := p.stringMap.Get().(map[string]interface{})
	// Clear the map for reuse
	for k := range m {
		delete(m, k)
	}
	return m
}

// PutStringMap returns a string map to the pool
func (p *Pool) PutStringMap(m map[string]interface{}) {
	if m != nil && len(m) <= 64 { // Avoid holding onto huge maps
		p.stringMap.Put(m)
	}
}

// StringBuilderPool provides a pool of strings.Builder instances
var StringBuilderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

// GetStringBuilder returns a reusable strings.Builder from the pool
func GetStringBuilder() *strings.Builder {
	sb := StringBuilderPool.Get().(*strings.Builder)
	sb.Reset()
	return sb
}

// PutStringBuilder returns a strings.Builder to the pool
func PutStringBuilder(sb *strings.Builder) {
	if sb != nil && sb.Cap() <= 4096 { // Avoid holding onto huge builders
		StringBuilderPool.Put(sb)
	}
}
