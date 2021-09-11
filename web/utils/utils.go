package utils

const (
	Ok                  = 200
	Redirect            = 301
	NotFound            = 404
	InternalServerError = 505
)

type JsonFormat struct {
	Code    int
	Message string
	Data    interface{}
}

func (json *JsonFormat) Ok(data interface{}) JsonFormat {
	return JsonFormat{
		Code:    Ok,
		Message: "OK",
		Data:    data,
	}
}
func (json *JsonFormat) Error(interface{}) JsonFormat {
	return JsonFormat{
		Code:    NotFound,
		Message: "NotFound",
		Data:    nil,
	}
}
