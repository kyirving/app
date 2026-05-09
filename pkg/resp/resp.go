package resp

const (
	RESP_SUCC                       = 200
	RESP_CREATE                     = 201
	RESP_UPDATE                     = 202
	RESP_DELETE                     = 204
	RESP_RESET                      = 205
	RESP_PARTIAL                    = 206
	RESP_BAD_REQUEST                = 400
	RESP_UNAUTHORIZED               = 401
	RESP_FORBIDDEN                  = 403
	RESP_NOT_FOUND                  = 404
	RESP_METHOD_NOT_ALLOWED         = 405
	RESP_NOT_ACCEPTABLE             = 406
	RESP_GONE                       = 410
	RESP_FAIL                       = 500
	RESP_BAD_GATEWAY                = 502
	RESP_GATEWAY_TIMEOUT            = 504
	RESP_HTTP_VERSION_NOT_SUPPORTED = 505
	RESP_INSUFFICIENT_STORAGE       = 507
	RESP_LOOP_DETECTED              = 508
	RESP_NOT_EXTENDABLE             = 510
)

var HttpCodeMap = map[int]string{
	RESP_SUCC:                       "success",
	RESP_CREATE:                     "create",
	RESP_UPDATE:                     "update",
	RESP_DELETE:                     "delete",
	RESP_RESET:                      "reset",
	RESP_PARTIAL:                    "partial",
	RESP_BAD_REQUEST:                "bad request",
	RESP_UNAUTHORIZED:               "unauthorized",
	RESP_FORBIDDEN:                  "forbidden",
	RESP_NOT_FOUND:                  "not found",
	RESP_METHOD_NOT_ALLOWED:         "method not allowed",
	RESP_NOT_ACCEPTABLE:             "not acceptable",
	RESP_GONE:                       "gone",
	RESP_BAD_GATEWAY:                "bad gateway",
	RESP_GATEWAY_TIMEOUT:            "gateway timeout",
	RESP_HTTP_VERSION_NOT_SUPPORTED: "http version not supported",
	RESP_INSUFFICIENT_STORAGE:       "insufficient storage",
	RESP_LOOP_DETECTED:              "loop detected",
	RESP_NOT_EXTENDABLE:             "not extendable",
	RESP_FAIL:                       "fail",
}

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Output(code int, data interface{}, message ...string) *Resp {

	res := &Resp{
		Code: code,
		Data: data,
	}

	if message != nil {
		res.Message = message[0]
		return res
	}

	if msg, ok := HttpCodeMap[code]; ok {
		res.Message = msg
	}
	return res
}
