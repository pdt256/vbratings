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
	buf := make([]byte, 32)
	u := uuid.NewV4()

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
	next := s.ids[s.currentIndex]

	s.currentIndex++

	return next
}
