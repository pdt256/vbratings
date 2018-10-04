package uuid

import (
	"encoding/hex"

	"github.com/satori/go.uuid"
)

type Generator interface {
	NewV4() string
}

type service struct{}

var _ Generator = (*service)(nil)

func NewService() *service {
	return &service{}
}

func (service) NewV4() string {
	return toHex(uuid.NewV4())
}

func toHex(u uuid.UUID) string {
	buf := make([]byte, 32)

	hex.Encode(buf[:], u[:])

	return string(buf)
}

type staticService struct {
	ids          []string
	currentIndex int
}

var _ Generator = (*staticService)(nil)

func NewStaticService(ids []string) *staticService {
	return &staticService{
		ids:          ids,
		currentIndex: 0,
	}
}

func (s *staticService) NewV4() string {
	if s.currentIndex >= len(s.ids) {
		return toHex(uuid.NewV4())
	}

	next := s.ids[s.currentIndex]

	s.currentIndex++

	return next
}
