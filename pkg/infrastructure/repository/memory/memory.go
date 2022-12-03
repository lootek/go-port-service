package memory

import (
	"sort"
	"sync"

	"github.com/google/uuid"
	"github.com/lootek/go-port-service/pkg/core/domain"
)

type Storage struct {
	dataMu sync.RWMutex
	data   map[string]domain.Port
}

func NewStorage() *Storage {
	return &Storage{
		data: map[string]domain.Port{},
	}
}

func (s *Storage) GetAll() []domain.Port {
	s.dataMu.RLock()
	defer s.dataMu.RUnlock()

	result := make([]domain.Port, 0, len(s.data))
	for _, v := range s.data {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func (s *Storage) Upsert(ports []domain.Port) error {
	for _, p := range ports {
		id := p.ID
		if id == "" {
			id = uuid.NewString()
		}

		// Make the synchronization block as narrow as possible
		// to allow future parallelization for the better performance
		s.dataMu.Lock()
		s.data[id] = p
		s.dataMu.Unlock()
	}

	return nil
}
