package ping

type PingUsecase interface {
	Test() (error, PingResponse)
}

type PingResponse interface {
	GetData() string
}

type pingUsecase struct {
}

func NewPingUsecase() PingUsecase {
	return &pingUsecase{}
}

func (this pingUsecase) Test() (error, PingResponse) {
	return nil, pingResponse{Data: "tes"}
}
