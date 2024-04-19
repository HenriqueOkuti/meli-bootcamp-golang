package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) JSON(code int, data interface{}) {
	r.Code = code
	r.Data = data
}

func (r *Response) JSONError(code int, err error) {
	r.Code = code
	r.Error = err.Error()
}
