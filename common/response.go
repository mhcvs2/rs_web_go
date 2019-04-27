package common

type ResponseCode int

const (
	SUCCESS  ResponseCode = 200
	ERROR                 = 500
	NO_LOGIN              = 401
)

type Response struct {
	Status  ResponseCode `json:"status"`
	Message string       `json:"message"`
	Result  interface{}  `json:"result"`
}

func NewResponse(status ResponseCode, msg string, result interface{}) *Response {
	if result == nil {
		result = make([]string, 0)
	}
	return &Response{Status: status, Message: msg, Result: result}
}

func NewSuccessResponse(result interface{}) *Response {
	return NewResponse(SUCCESS, "success", result)
}

func NewErrorResponse(result interface{}) *Response {
	return NewResponse(ERROR, "error", result)
}

func NewNoLoginResponse(result interface{}) *Response {
	return NewResponse(NO_LOGIN, "not login", result)
}

func NewPageResponse(data interface{}, dataName string, page *PageParam) *Response {
	result := make(map[string]interface{})
	result[dataName] = data
	result["pagination"] = page
	return NewResponse(SUCCESS, "success", result)
}