// Package memory is an output adapter to store entities in memory
package memory

// memoryRepository implements the trip.IRepository
type memoryRepository struct {
	*memoryRepositoryBase // DO NOT REMOVE IT
}

// NewMemoryRepository in memory repository constructor function
func NewMemoryRepository(cryptoKey []byte) *memoryRepository {
	return &memoryRepository{memoryRepositoryBase: newMemoryRepositoryBase(cryptoKey)}
}
