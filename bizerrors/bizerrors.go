package bizerrors

import (
	"fmt"
	"net/http"
)

var SILENT = 0
var WARN_MESSAGE = 1
var ERROR_MESSAGE = 2
var NOTIFICATION = 3
var REDIRECT = 9

// BizError struct
type BizError struct {
	HttpStatus int `json:"-"`
	//code       string
	//message    string
	Success      bool   `json:"success"`      // always false
	ErrorCode    string `json:"errorCode"`    // code for errorType
	ErrorMessage string `json:"errorMessage"` // message display to user
	ShowType     int    `json:"showType"`     // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
	TraceId      string `json:"traceId"`      // Convenient for back-end Troubleshooting: unique request ID
	Host         string `json:"host"`         // onvenient for backend Troubleshooting: host of current access server
}

// Code returns code.
func (e *BizError) Code() string {
	return e.ErrorCode
}

// Error returns message.
func (e *BizError) Error() string {
	return e.ErrorMessage
}

var (
	NotificationWarning = func(msg string) *BizError {
		return &BizError{HttpStatus: http.StatusBadRequest, Success: false, ErrorCode: "NOTIFICATION_WARNING", ErrorMessage: msg,
			ShowType: NOTIFICATION,
		}
	}
	UnknownError = func(msg string) *BizError {
		return &BizError{HttpStatus: http.StatusInternalServerError, Success: false, ErrorCode: "SYS_ERROR", ErrorMessage: msg,
			ShowType: 2,
		}
	}
	BadRequest = func(msg string, a ...interface{}) *BizError {
		return &BizError{HttpStatus: http.StatusBadRequest, Success: false, ErrorCode: "BAD_REQUEST", ErrorMessage: fmt.Sprintf(msg, a...),
			ShowType: 2,
		}
	}
	MissingCtxValueTenantID = &BizError{HttpStatus: http.StatusBadRequest, Success: false, ErrorCode: "MISSING_CTX_VALUE_TENANT_ID", ErrorMessage: "Context中缺失tenant_id",
		ShowType: 2,
	}
	MissingAccessToken = &BizError{HttpStatus: http.StatusUnauthorized, Success: false, ErrorCode: "UNAUTHORIZED", ErrorMessage: "ACCESS TOKEN缺失",
		ShowType: 2,
	}
	Unauthorized = &BizError{HttpStatus: http.StatusUnauthorized, Success: false, ErrorCode: "UNAUTHORIZED", ErrorMessage: "ACCESS TOKEN失效",
		ShowType: 2,
	}
	ErrNotFound = func(msg string) *BizError {
		return &BizError{HttpStatus: http.StatusNotFound, Success: false, ErrorCode: "NOT_FOUND", ErrorMessage: msg,
			ShowType: 1,
		}
	}
	ErrorContractDeadlineInvalidFormat  = &BizError{ErrorCode: "001", HttpStatus: http.StatusBadRequest, ErrorMessage: "deadline Invalid format", ShowType: 1}
	ErrorContractActionInvalidParameter = &BizError{ErrorCode: "002", HttpStatus: http.StatusBadRequest, ErrorMessage: "contract action invalid", ShowType: 1}
	//ErrorCarActionInvalid               = &BizError{ErrorCode: "004", HttpStatus: http.StatusBadRequest, ErrorMessage: "car action invalid"}
	ErrorGetTransportByIdNotFind  = &BizError{ErrorCode: "005", HttpStatus: http.StatusBadRequest, ErrorMessage: "transport Not Find", ShowType: 1}
	ErrorTransportIsPush          = &BizError{ErrorCode: "006", HttpStatus: http.StatusBadRequest, ErrorMessage: "派车单已经推送，无法操作", ShowType: 1}
	ErrorCarStatusStatusDelivered = &BizError{ErrorCode: "007", HttpStatus: http.StatusBadRequest, ErrorMessage: "派车单已经完成，无法操作", ShowType: 1}
	//ErrorTransportDelete                = func(message string) *BizError {
	//  return &BizError{ErrorCode: "010", HttpStatus: http.StatusBadRequest,
	//    ErrorMessage:             message}
	//}
	//ErrorDoActionForTransportCar             = &BizError{ErrorCode: "013", HttpStatus: http.StatusBadRequest, ErrorMessage: "不能执行该操作"}
	ErrorGetCompanyByIdNotFind  = &BizError{ErrorCode: "015", HttpStatus: http.StatusBadRequest, ErrorMessage: "Company Not Find", ShowType: 1}
	ErrorStatementActionInvalid = &BizError{ErrorCode: "017", HttpStatus: http.StatusBadRequest, ErrorMessage: "statement action invalid", ShowType: 1}
	//NoDataResultError = &BizError{http.StatusOK, "20101", "没有有效数据返回"}
	//VerifyTokenError = &BizError{http.StatusOK, "20101", "用户未登陆"}
	//VerifyTokenFailure = &BizError{http.StatusOK, "10016", "你已在其他设备登陆"}
	//VerifyTokenFailure2 = &BizError{http.StatusOK, "10017", "你已在其他设备登陆"}
	//VerifyTokenServerError = &BizError{http.StatusNotFound, "20101", "用户未登陆"}
	//VerifySignError = &BizError{http.StatusOK, "10012", "签名错误"}
	//VerifySignServerError = &BizError{http.StatusNotFound, "10012", "签名错误"}
	//VerifyCodeError = &BizError{http.StatusOK, "10013", "验证码有误"}
	//PreventFrequentError = &BizError{http.StatusOK, "10212", "您提交过快，请稍后提交"}
	//ProuctEmptyError = &BizError{http.StatusOK, "50000", "商品不存在"}
	//ProuctIdEmptyError = &BizError{http.StatusOK, "50001", "商品ID不能为空"}
	//UserNotExistError = &BizError{http.StatusOK, "132", "用户不存在"}
	//UserPasswordError = &BizError{http.StatusOK, "133", "用户或密码不正确"}
	//UserPermissionError = &BizError{http.StatusBadRequest, "1", "用户无发布权限"}
	//QuickLoginNotExistError = &BizError{http.StatusOK, "15601", "第三方登录用户未注册"}
	//QuickLoginExistError = &BizError{http.StatusOK, "15602", "第三方登录已注册"}
	//QuickUidNullError = &BizError{http.StatusOK, "15603", "第三方登录用户凭证不能为空"}
	//QuickUidBindPhoneError = &BizError{http.StatusOK, "15604", "手机号已被绑定"}
	//AntiRubbishRequestError = &BizError{http.StatusOK, "10301", "网易反垃圾请求错误"}
	//AntiRubbishADError = &BizError{http.StatusOK, "10302", "检查评论内容有广告或关键字嫌疑"}
	//AntiRubbishError = &BizError{http.StatusOK, "10303", "检查评论内容不通过"}
	//AntiRubbishUserNameError = &BizError{http.StatusOK, "10304", "用户名检查不通过"}
	//AntiRubbishUserNameAdError = &BizError{http.StatusOK, "10305", "检查用户名有广告嫌疑"}
	//AntiRubbishPhoneExistdError = &BizError{http.StatusOK, "10306", "手机号已存在"}
	//AntiRubbishUserBlacklistError = &BizError{http.StatusOK, "10307", "用户已经黑名单"}
	//AntiRubbishUserNameExistError = &BizError{http.StatusOK, "10308", "用户名已存在"}
	//AntiRubbishPhoneMathError = &BizError{http.StatusOK, "10309", "手机号不匹配"}
	//AntiRubbishPhoneNullError = &BizError{http.StatusOK, "10310", "手机号不能为空"}
	//AntiRubbishUpdateCodeError = &BizError{http.StatusOK, "10311", "修改验证码状态失败"}
	//ArticleNotExistdError = &BizError{http.StatusOK, "10312", "文章不存在"}
	//ReportArticleNotExistdError = &BizError{http.StatusOK, "10314", "举报的评论信息不存在"}
	//FindArticleNotExistdError = &BizError{http.StatusOK, "10315", "发表评论时没有找到对应的文章信息"}
	//UserLimitCommentsError = &BizError{http.StatusOK, "10316", "您已被限制评论"}
	//RepeatCommitCommentsError = &BizError{http.StatusOK, "10317", "短时间内重复评论"}
	//CommentsLimitGt500Error = &BizError{http.StatusOK, "10318", "评论限制500字, 您已经超出"}
	//CommentsNotExistdError = &BizError{http.StatusOK, "10319", "评论信息不存在"}
	//ForbiddenCommentError = &BizError{http.StatusOK, "10320", "该文章禁止评论"}
	//PhoneNotExistdError = &BizError{http.StatusOK, "10321", "手机号不存在"}
	//AntiRubbishUserAvatarError = &BizError{http.StatusOK, "10322", "系统暂不支持头像修改"}
)
