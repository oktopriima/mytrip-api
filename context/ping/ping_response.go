package ping

type pingResponse struct {
	Data string `json:"message"`
}

func (this pingResponse) GetData() string {
	return this.Data
}
