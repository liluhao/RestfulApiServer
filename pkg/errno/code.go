package errno

//code.go文件统一存自定义错误码
/*在实际开发中引入错误码有如下好处：
可以非常方便地定位问题和定位代码行（看到错误码知道什么意思，grep错误码可以定位到错误码所在行）;
如果API对外开放，有个错误码会更专业些;
错误码包含一定的信息，通过错误码可以判断出错误级别、错误模块和具体错误信息;
在实际业务开发中，一个条错误信息需要包含两部分内容：直接展示给用户的message和用于开发人员debug的error。
message可能会直接展示给用户error:是用于debug的错误信息，可能包含敏感内部信息，不宜对外展示
业务开发过程中，可能需要判断错误是哪种类型以便做相应的逻辑处理，通过定制的错误码很容易做到这点
*/
//Go中的HTTP服务器开发都是引用net/http包，该包中只有60个错误码，基本都是跟HTTP请求相关的。在大型系统中，这些错误码完全不够用，而且跟业务没有任何关联，满足不了业务需求。
//以下全是自定义的错误，部分模仿新浪,即我们定制的
/*错误代码通常由5位数组成：
从左往右第1位为系统级参数：1为系统级错误；2为普通错误，通常是由用户非法操作引起的
从左往右第2位+第3位为服务模块代码：服务模块为两位数，一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
从左往右第4位+第5位为具体错误代码：错误码为两位数：防止一个模块定制过多的错误码，后期不好维护*/
/*code=0说明是正确返回，code>0说明是错误返回
错误通常包话系统级错误码和服务级错误码
错误码均为>=0的数
在apiserver中HTTP Code固定为http.StatusOK,错误码通过code来表示。
*/
/*
如果AP川是对外的，错误信息数量有限，则制定错误码非常容易，强烈建议使用错误码。如果是内部系统，特
别是庞大的系统，内部错误会非常多，这时候没必要为每一个错误制定错误码，而只需为常见的错误制定错误
码，对于普通的错误，系统在处理时会统一作为InternalServerError处理。
*/
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."} //即不传如任何参数

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
)
