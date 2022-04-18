package errno

import "fmt"

//在实际开发中，一个错误类型通常包含两部分：Code部分，用来唯一标识一个错误；Message部分，用来展示错误信息，
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

//errno.go源码文件中有两个核心函数New()和DecodeErr(),一个用来新建定制的错误，一个用来解析定制的错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}
func DecodeErr(err error) (int, string) {
	//如果err=nil,则errno.DecodeErr(err)会返回成功的code:0和message:OK。
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

//同时也提供了Add()和Addf()函数，如果想对外展示更多的信息可以调用比函数，使用方法
func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}
