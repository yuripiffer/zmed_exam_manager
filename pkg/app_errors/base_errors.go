package app_errors

type ErrorStruct struct {
	AppError
	Title    string
	Detail   string
	Err      string
	ErrorKey string
	Data     interface{}
}

type AppError interface {
	error
	GetTitle() string
	GetDetail() string
	GetErr() string
	GetKey() string
	GetData() interface{}
	SetData(data interface{})
}

func (e *ErrorStruct) Error() string {
	return e.GetDetail()
}

func (e *ErrorStruct) GetTitle() string {
	return e.Title
}

func (e *ErrorStruct) GetDetail() string {
	return e.Detail
}

func (e *ErrorStruct) GetErr() string {
	return e.Err
}

func (e *ErrorStruct) GetKey() string {
	return e.ErrorKey
}

func (e *ErrorStruct) SetData(data interface{}) {
	e.Data = data
}

func (e *ErrorStruct) GetData() interface{} {
	return e.Data
}
