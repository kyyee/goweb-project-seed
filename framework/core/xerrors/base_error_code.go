package xerrors

import (
	"fmt"

	"github.com/pkg/errors"
)

const ERROR_CODE_PRIFIX = "00201"
const SUCCESS_CODE = "0000000000"
const (
	SYS_INTERNAL_ERROR                = "00001"
	SYS_INIT_ERROR                    = "00002"
	UNKNOWN_ERROR                     = "00003"
	MEMORY_ALLOCATION_FAILED          = "00004"
	MEMORY_OVERFLOW_ERROR             = "00005"
	REQUEST_TIMEOUT_ERROR             = "00006"
	INVALID_HANDLE                    = "00007"
	ENVIRONMENT_CHANGE_FAILED         = "00008"
	SYS_BUSY                          = "00009"
	DATA_SOURCE_ERROR                 = "00010"
	FREQUENT_OPERATION                = "00011"
	UNKNOWN_SOURCE_ERROR              = "00012"
	REQUEST_LIMIT_ERROR               = "00013"
	GET_CONFIGURATION_ERROR           = "00014"
	SET_CONFIGURATION_ERROR           = "00015"
	PROGRAM_RESTART_REQUIRED          = "00016"
	SYS_RESTART_REQUIRED              = "00017"
	FEATURE_NOT_SUPPORTED             = "00018"
	VALIDATION_FAILED                 = "00019"
	CALL_FAILED                       = "00020"
	MESSAGE_COMPONENT_EXCEPTION       = "00021"
	INSUFFICIENT_STORAGE_CAPACITY     = "00022"
	STORAGE_SERVICE_ERROR             = "00023"
	NETWORK_ERROR                     = "00024"
	ILLEGAL_IP_ADDRESS                = "00025"
	IP_FORBIDDEN_ERROR                = "00026"
	DOMAIN_NAME_RESOLUTION_FAILED     = "00027"
	PORT_ERROR                        = "00028"
	DOMAIN_ERROR                      = "00029"
	IP_LIMIT_ERROR                    = "00030"
	IP_EMPTY_ERROR                    = "00031"
	CONNECTION_FAILED                 = "00032"
	DATA_SENDING_FAILED               = "00033"
	DATA_RECEPTION_FAILED             = "00034"
	HTTP_CLIENT_INITIALIZATION_FAILED = "00035"
	HTTP_REQUEST_FAILED               = "00036"
	HTTP_REQUEST_TIMEOUT              = "00037"
	HTTP_ERROR_400                    = "00038"
	HTTP_ERROR_404                    = "00039"
	HTTP_ERROR_500                    = "00040"
	KAFKA_TOPIC_INVALID               = "00041"
	KAFKA_MESSAGE_SENDING_FAILED      = "00042"
	DATABASE_ERROR                    = "00043"
	DATABASE_OPERATION_FAILED         = "00044"
	RESULT_EMPTY_ERROR                = "00045"
	ID_EMPTY_ERROR                    = "00046"
	DATA_DUPLICATION                  = "00047"
	DATA_UPDATING                     = "00048"
	DATA_NOT_UNIQUE_ERROR             = "00049"
	DICTIONARY_ACQUISITION_FAILED     = "00050"
	API_NOT_EXIST_ERROR               = "00051"
	INVALID_PARAM_ERROR               = "00052"
	INVALID_INPUT_PARAMETER           = "00053"
	MISSING_REQUIRED_PARAMETERS       = "00054"
	TYPE_NOT_SUPPORTED                = "00055"
	DUPLICATE_NAME                    = "00056"
	PARAMETER_EMPTY                   = "00057"
	LANGUAGE_NOT_SUPPORTED            = "00058"
	API_OFF_ERROR                     = "00059"
	INTERFACE_MAINTENANCE             = "00060"
	VERSION_ERROR                     = "00061"
	VERIFICATION_CODE_ERROR           = "00062"
	NOT_EMPTY                         = "00063"
	MOBILE_NUMBER_ERROR               = "00064"
	ID_NUMBER_ERROR                   = "00065"
	EMAIL_ERROR                       = "00066"
	LONGITUDE_AND_LATITUDE_ERROR      = "00067"
	REQUEST_PATH_ERROR                = "00068"
	DATA_FORMAT_ERROR                 = "00069"
	TIME_FORMAT_ERROR                 = "00070"
	INVALID_TIME                      = "00071"
	PLATE_NUMBER_ERROR                = "00072"
	CAMERA_NUMBER_ERROR               = "00073"
	CALLBACK_URL_ERROR                = "00074"
	TIME_FORMAT_LIMIT_ERROR           = "00075"
	PAGE_NUMBER_LIMIT_ERROR           = "00076"
	PAGE_SIZE_LIMIT_ERROR             = "00077"
	THIRD_PARTY_INTERFACE_FAILED      = "00078"
	JSON_PARSING_ERROR                = "00079"
	JSON_FIELD_TYPE_ERROR             = "00080"
	JSON_FIELD_VALUE_ERROR            = "00081"
	JSON_FIELD_TYPE_MISSING           = "00082"
	JSON_OBJECT_CREATE_FAILED         = "00083"
	JSON_ARRAY_CREATE_FAILED          = "00084"
	JSON_FORMAT_ERROR                 = "00085"
	FILE_NAME_ERROR                   = "00086"
	FILE_ADDRESS_NOT_EXIST            = "00087"
	FILE_ID_NOT_EXIST                 = "00088"
	FILE_ID_TOO_LONG                  = "00089"
	FILE_TYPE_ERROR                   = "00090"
	FILE_CREATE_ERROR                 = "00091"
	FILE_OPEN_ERROR                   = "00092"
	DOWNLOAD_ERROR                    = "00093"
	FILE_READ_ERROR                   = "00094"
	FILE_WRITE_ERROR                  = "00095"
	FILE_FORMAT_ERROR                 = "00096"
	FILE_IS_ENCRYPTED                 = "00097"
	FILE_TRANSFER_ERROR               = "00098"
	EXCEL_BUILD_FAILED                = "00099"
	EXCEL_IMPORT_FAILED               = "00100"
	ENCODING_CONVERSION_FAILED        = "00101"
	FILE_UNRECOGNIZED                 = "00102"
	FILE_SIZE_ERROR                   = "00103"
	FILE_UPLOAD_ERROR                 = "00104"
	FILE_IS_BEING_UPLOADED            = "00105"
	USER_NOT_EXIST                    = "00106"
	USER_NOT_LOGGED_IN                = "00107"
	PERMISSION_DENY_ERROR             = "00108"
	PASSWORD_ERROR                    = "00109"
	USERNAME_ALREADY_EXISTS           = "00110"
	USERNAME_LENGTH_ERROR             = "00111"
	PASSWORD_LENGTH_ERROR             = "00112"
	USERNAME_OR_PASSWORD_FORMAT_ERROR = "00113"
	USERNAME_OR_PASSWORD_EMPTY        = "00114"
	LOGIN_EXPIRED                     = "00115"
	SECURITY_AUTHENTICATION_FAILED    = "00116"
	ACCOUNT_LOCKED                    = "00117"
	USER_ADD_ERROR                    = "00118"
	USER_DELETE_ERROR                 = "00119"
	USER_UPDATE_ERROR                 = "00120"
	USER_QUERY_ERROR                  = "00121"
	PASSWORD_UPDATE_ERROR             = "00122"
	NEW_PASSWORD_ERROR                = "00123"
	NEW_PASSWORD_LENGTH_ERROR         = "00124"
	OLD_PASSWORD_LENGTH_ERROR         = "00125"
	OLD_PASSWORD_ERROR                = "00126"
	PASSWORD_SAME_ERROR               = "00127"
	USERNAME_OR_PASSWORD_ERROR        = "00128"
	DEVICE_ACQUISITION_FAILED         = "00129"
	CHANNEL_ACQUISITION_FAILED        = "00130"
	DEVICE_OFFLINE                    = "00131"
	RESOURCE_TREE_ACQUISITION_FAILED  = "00132"
	RESOURCE_EXIST_ERROR              = "00133"
	DEVICE_NOT_EXIST                  = "00134"
	DEVICE_ALREADY_EXISTS             = "00135"
	DEVICE_LIMIT_ERROR                = "00136"
	CHANNEL_LIMIT_ERROR               = "00137"
	CAMERA_ID_ERROR                   = "00138"
	CAMERA_NAME_ERROR                 = "00139"
	IPC_CODE_TOO_LONG                 = "00140"
	IPC_NAME_TOO_LONG                 = "00141"
	IPC_STYLE_ACQUISITION_FAILED      = "00142"

	INSERT_FAILED               = "01001"
	UPDATE_FAILED               = "01002"
	DELETE_FAILED               = "01003"
	PRIMARY_KEY_NOT_EMPTY_ERROR = "01004"
	SQL_EXCEPTION               = "01005"
	PARAM_SWITCH_ERROR          = "01006"
)

var Msg = map[string]string{
	SYS_INTERNAL_ERROR:                "系统内部异常",
	SYS_INIT_ERROR:                    "系统初始化异常",
	UNKNOWN_ERROR:                     "未知错误",
	MEMORY_ALLOCATION_FAILED:          "内存分配失败",
	MEMORY_OVERFLOW_ERROR:             "内存溢出",
	REQUEST_TIMEOUT_ERROR:             "请求超时",
	INVALID_HANDLE:                    "非法句柄",
	ENVIRONMENT_CHANGE_FAILED:         "环境变更获取失败",
	SYS_BUSY:                          "系统繁忙",
	DATA_SOURCE_ERROR:                 "数据源异常",
	FREQUENT_OPERATION:                "当前操作频繁",
	UNKNOWN_SOURCE_ERROR:              "未知的请求源",
	REQUEST_LIMIT_ERROR:               "请求超过次数限制",
	GET_CONFIGURATION_ERROR:           "获取配置错误",
	SET_CONFIGURATION_ERROR:           "设置配置错误",
	PROGRAM_RESTART_REQUIRED:          "需要重启程序",
	SYS_RESTART_REQUIRED:              "需要重启系统",
	FEATURE_NOT_SUPPORTED:             "特性不支持",
	VALIDATION_FAILED:                 "验证失败",
	MESSAGE_COMPONENT_EXCEPTION:       "消息组件异常",
	INSUFFICIENT_STORAGE_CAPACITY:     "存储容量不足",
	STORAGE_SERVICE_ERROR:             "存储服务错误",
	NETWORK_ERROR:                     "网络错误",
	ILLEGAL_IP_ADDRESS:                "非法IP地址",
	IP_FORBIDDEN_ERROR:                "被禁止的IP",
	DOMAIN_NAME_RESOLUTION_FAILED:     "解析域名获取IP失败",
	PORT_ERROR:                        "端口错误",
	DOMAIN_ERROR:                      "域名错误",
	IP_LIMIT_ERROR:                    "当前IP请求超过限制",
	IP_EMPTY_ERROR:                    "IP地址不能为空",
	CONNECTION_FAILED:                 "连接失败",
	DATA_SENDING_FAILED:               "数据发送失败",
	DATA_RECEPTION_FAILED:             "数据接收失败",
	HTTP_CLIENT_INITIALIZATION_FAILED: "HTTP客户端初始化失败",
	HTTP_REQUEST_FAILED:               "HTTP请求失败",
	HTTP_REQUEST_TIMEOUT:              "HTTP请求超时",
	HTTP_ERROR_400:                    "HTTP错误码400",
	HTTP_ERROR_404:                    "HTTP错误码404",
	HTTP_ERROR_500:                    "HTTP错误码500",
	KAFKA_TOPIC_INVALID:               "kafka的topic无效",
	KAFKA_MESSAGE_SENDING_FAILED:      "kafka消息发送失败",
	DATABASE_ERROR:                    "数据库错误",
	DATABASE_OPERATION_FAILED:         "数据库操作失败",
	RESULT_EMPTY_ERROR:                "查询无结果",
	ID_EMPTY_ERROR:                    "ID不能为空",
	DATA_DUPLICATION:                  "数据重复",
	DATA_UPDATING:                     "数据更新中",
	DATA_NOT_UNIQUE_ERROR:             "查询结果不唯一",
	DICTIONARY_ACQUISITION_FAILED:     "字典获取失败",
	API_NOT_EXIST_ERROR:               "接口不存在",
	INVALID_PARAM_ERROR:               "参数校验错误",
	INVALID_INPUT_PARAMETER:           "无效入参",
	MISSING_REQUIRED_PARAMETERS:       "缺少必要参数",
	TYPE_NOT_SUPPORTED:                "类型不支持",
	DUPLICATE_NAME:                    "名称重复",
	PARAMETER_EMPTY:                   "参数为空",
	LANGUAGE_NOT_SUPPORTED:            "不支持该语言",
	API_OFF_ERROR:                     "接口停用",
	INTERFACE_MAINTENANCE:             "接口维护",
	VERSION_ERROR:                     "版本号错误",
	VERIFICATION_CODE_ERROR:           "验证码错误",
	NOT_EMPTY:                         "%s不能为空",
	MOBILE_NUMBER_ERROR:               "手机号码错误",
	ID_NUMBER_ERROR:                   "身份证号码错误",
	EMAIL_ERROR:                       "邮箱错误",
	LONGITUDE_AND_LATITUDE_ERROR:      "经纬度错误",
	REQUEST_PATH_ERROR:                "请求路径错误",
	DATA_FORMAT_ERROR:                 "数据格式错误",
	TIME_FORMAT_ERROR:                 "时间参数错误",
	INVALID_TIME:                      "无效时间",
	PLATE_NUMBER_ERROR:                "车牌号错误",
	CAMERA_NUMBER_ERROR:               "摄像头编号错误",
	CALLBACK_URL_ERROR:                "回调地址URL错误",
	TIME_FORMAT_LIMIT_ERROR:           "时间格式必须是ISO-8601",
	PAGE_NUMBER_LIMIT_ERROR:           "页码超出限制",
	PAGE_SIZE_LIMIT_ERROR:             "分页大小超出限制",
	THIRD_PARTY_INTERFACE_FAILED:      "调用第三方接口失败",
	JSON_PARSING_ERROR:                "json解析错误",
	JSON_FIELD_TYPE_ERROR:             "json消息字段类型错误",
	JSON_FIELD_VALUE_ERROR:            "json消息字段取值错误",
	JSON_FIELD_TYPE_MISSING:           "json消息字段缺失",
	JSON_OBJECT_CREATE_FAILED:         "创建json对象失败",
	JSON_ARRAY_CREATE_FAILED:          "创建json数组失败",
	JSON_FORMAT_ERROR:                 "json格式或内容错误",
	FILE_NAME_ERROR:                   "文件名错误",
	FILE_ADDRESS_NOT_EXIST:            "文件地址不存在",
	FILE_ID_NOT_EXIST:                 "文件ID不存在",
	FILE_ID_TOO_LONG:                  "文件ID超长",
	FILE_TYPE_ERROR:                   "文件类型错误",
	FILE_CREATE_ERROR:                 "创建文件失败",
	FILE_OPEN_ERROR:                   "打开文件失败",
	DOWNLOAD_ERROR:                    "下载文件失败",
	FILE_READ_ERROR:                   "读取文件失败",
	FILE_WRITE_ERROR:                  "写入文件失败",
	FILE_FORMAT_ERROR:                 "文件格式错误",
	FILE_IS_ENCRYPTED:                 "文件被加密",
	FILE_TRANSFER_ERROR:               "文件转换失败",
	EXCEL_BUILD_FAILED:                "Excel生成失败",
	EXCEL_IMPORT_FAILED:               "Excel导入失败",
	ENCODING_CONVERSION_FAILED:        "文本编码转换失败",
	FILE_UNRECOGNIZED:                 "未识别文件",
	FILE_SIZE_ERROR:                   "文件大小超过限制",
	FILE_UPLOAD_ERROR:                 "文件上传失败",
	FILE_IS_BEING_UPLOADED:            "其它用户正在上传该文件",
	USER_NOT_EXIST:                    "用户不存在",
	USER_NOT_LOGGED_IN:                "用户未登录",
	PERMISSION_DENY_ERROR:             "用户没有权限",
	PASSWORD_ERROR:                    "密码错误",
	USERNAME_ALREADY_EXISTS:           "用户名已存在",
	USERNAME_LENGTH_ERROR:             "用户名长度错误",
	PASSWORD_LENGTH_ERROR:             "密码长度错误",
	USERNAME_OR_PASSWORD_FORMAT_ERROR: "用户名或密码格式错误",
	USERNAME_OR_PASSWORD_EMPTY:        "用户名或密码为空",
	LOGIN_EXPIRED:                     "用户登录已过期",
	SECURITY_AUTHENTICATION_FAILED:    "安全认证失败",
	ACCOUNT_LOCKED:                    "账号被锁定",
	USER_ADD_ERROR:                    "添加用户失败",
	USER_DELETE_ERROR:                 "删除用户失败",
	USER_UPDATE_ERROR:                 "修改用户失败",
	USER_QUERY_ERROR:                  "查询用户失败",
	PASSWORD_UPDATE_ERROR:             "修改密码失败",
	NEW_PASSWORD_ERROR:                "新密码错误",
	NEW_PASSWORD_LENGTH_ERROR:         "新密码长度错误",
	OLD_PASSWORD_LENGTH_ERROR:         "原密码长度错误",
	OLD_PASSWORD_ERROR:                "原密码错误",
	PASSWORD_SAME_ERROR:               "新密码与原密码相同",
	USERNAME_OR_PASSWORD_ERROR:        "用户名或密码错误",
	DEVICE_ACQUISITION_FAILED:         "设备获取失败",
	CHANNEL_ACQUISITION_FAILED:        "通道获取失败",
	DEVICE_OFFLINE:                    "设备离线",
	RESOURCE_TREE_ACQUISITION_FAILED:  "资源树获取失败",
	RESOURCE_EXIST_ERROR:              "资源已经存在",
	DEVICE_NOT_EXIST:                  "设备不存在",
	DEVICE_ALREADY_EXISTS:             "设备已经存在",
	DEVICE_LIMIT_ERROR:                "设备数达到上限",
	CHANNEL_LIMIT_ERROR:               "通道数达到上限",
	CAMERA_ID_ERROR:                   "相机ID错误",
	CAMERA_NAME_ERROR:                 "相机名错误",
	IPC_CODE_TOO_LONG:                 "IPC设备编码超长",
	IPC_NAME_TOO_LONG:                 "IPC设备名称超长",
	IPC_STYLE_ACQUISITION_FAILED:      "获取IPC款型失败",

	INSERT_FAILED:               "创建失败",
	UPDATE_FAILED:               "更新失败",
	DELETE_FAILED:               "删除失败",
	PRIMARY_KEY_NOT_EMPTY_ERROR: "主键不能为空",
	SQL_EXCEPTION:               "SQL执行异常",
	PARAM_SWITCH_ERROR:          "参数转换异常",
}

type Error struct {
	Code  string
	Msg   string
	Cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code = %s, message = %s, cause = %+v", e.Code, e.Msg, e.Cause)
}
func (e *Error) Unwrap() error {
	return e.Cause
}

func NewAssignCodeError(code string) error {
	var msg string
	var ok bool
	if msg, ok = Msg[code]; !ok {
		msg = "未知错误"
	}
	return errors.WithStack(&Error{
		Code: code,
		Msg:  msg,
	})
}

func NewAssignCodeAndCauseError(code string, err error) error {
	var msg string
	var ok bool
	if msg, ok = Msg[code]; !ok {
		msg = "未知错误"
	}
	return errors.WithStack(&Error{
		Code:  code,
		Msg:   msg,
		Cause: err,
	})
}

func NewCustomError(code string, format string, a ...any) error {
	return errors.WithStack(&Error{
		Code: code,
		Msg:  fmt.Sprintf(format, a...),
	})
}

func NewCustomAndCauseError(code string, err error, format string, a ...any) error {
	return errors.WithStack(&Error{
		Code:  code,
		Msg:   fmt.Sprintf(format, a...),
		Cause: err,
	})
}
