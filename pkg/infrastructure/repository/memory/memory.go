package memory

import (
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

func (s *Storage) GetAll() ([]domain.Port, error) {
	s.dataMu.RLock()
	defer s.dataMu.RUnlock()

	result := make([]domain.Port, 0, len(s.data))
	for _, v := range s.data {
		result = append(result, v)
	}

	return result, nil
}

func (s *Storage) Insert(p domain.Port) (string, error) {
	id := p.ID
	if id == "" {
		id = uuid.NewString()
	}

	return s.upsert(id, p)
}

func (s *Storage) Update(id string, p domain.Port) error {
	_, err := s.upsert(id, p)
	return err
}

func (s *Storage) upsert(id string, p domain.Port) (string, error) {
	s.dataMu.Lock()
	defer s.dataMu.Unlock()

	s.data[id] = p

	return id, nil
}
