package helper

type ResponseMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithToken struct {
	Status      bool   `json:"status"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}
