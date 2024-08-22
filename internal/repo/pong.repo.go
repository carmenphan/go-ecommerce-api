package repo

type PongRepo struct{}

func NewPongRepo() *PongRepo {
	return &PongRepo{}
}
func (pr *PongRepo) GetIoPong() string {
	return "Io Pong"
}
