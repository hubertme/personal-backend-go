package result

type Response struct {
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
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
		Message:   "Please contact developer to report this error",
		Data:      data,
	}
}

func FormatError(data interface{}) Response {
	return Response{
		ErrorCode: BAD_FORMATTED_REQUEST,
		Message:   "Bad formatted request sent",
		Data:      data,
	}
}

func EmptyError(data interface{}) Response {
	return Response{
		ErrorCode: EMPTY_FIELD_REQUEST,
		Message:   "Empty request field detected",
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
