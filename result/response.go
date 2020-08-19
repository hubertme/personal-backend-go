package result

type Response struct {
	ErrorCode int `json:"error_code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		ErrorCode: 0,
		Message:   "",
		Data:      data,
	}
}

func DevError(data interface{}) Response {
	return Response{
		ErrorCode: DEV_REPORT_ERROR,
		Message:   "",
		Data:      data,
	}
}

func Error(errorCode int, message string, data interface{}) Response {
	return Response{
		ErrorCode: errorCode,
		Message:   message,
		Data:      data,
	}
}