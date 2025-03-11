package chat

type Req struct {
	ReqIdentifier int32  `json:"reqIdentifier" validate:"required"` //消息类型
	Token         string `json:"token" `
	SendID        string `json:"sendID" validate:"required"`
	OperationID   string `json:"operationID" validate:"required"`
	MsgIncr       string `json:"msgIncr" validate:"required"`
	Data          []byte `json:"data"`
}

type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	MsgIncr       string `json:"msgIncr"`
	OperationID   string `json:"operationID"`
	ErrCode       int32  `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	Data          []byte `json:"data"`
}
