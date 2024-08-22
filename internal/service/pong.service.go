package service

import "github.com/carmenphan/go-ecommerce-backend-api/internal/repo"

type PongService struct {
	pongRepo *repo.PongRepo
}

func NewPongService() *PongService {
	return &PongService{
		pongRepo: repo.NewPongRepo(),
	}
}
func (ps *PongService) GetPong() string {
	return ps.pongRepo.GetIoPong()
}
