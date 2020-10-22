package Utils

type Response struct {
	Code       int      `json:"code"`
	Data       interface{}        `json:"data"`
}

func ResponseOk(data interface{}) Response {
	return Response {
		Code:200,
		Data:data,
	}
}
