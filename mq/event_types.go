package mq

type BasicEvent struct {
	EventId    string `json:"event_id"`
	OccurredAt string `json:"occurred_at"`
}

type K8sResourceEvent struct {
	BasicEvent    `json:",inline"`
	UserId        string `json:"user_id"`
	ResourceType  string `json:"resource_type"`
	Action        string `json:"action"`
	K8sObjectKind string `json:"k8s_obj_kind,omitempty"`
}

type StrategyEvent struct {
	BasicEvent           `json:",inline"`
	UserId               string                 `json:"user_id"`
	ResourceType         string                 `json:"resource_type"`
	Action               string                 `json:"action"`
	StrategyName         string                 `json:"strategy_name,omitempty"`
	StrategyImage        string                 `json:"image_name,omitempty"`
	StrategyType         string                 `json:"strategy_type,omitempty"`
	StrategyDeployEnvs   map[string]interface{} `json:"deploy_envs,omitempty"`
	StrategyUpdateDetail map[string]interface{} `json:"update_detail,omitempty"`
}

type ProductEvent struct {
	BasicEvent    `json:",inline"`
	UserId        string `json:"user_id"`
	ResourceType  string `json:"resource_type"`
	Action        string `json:"action"`
	ProductId     string `json:"pid,omitempty"`
	ProductName   string `json:"product_name,omitempty"`
	ProductImage  string `json:"image_name,omitempty"`
	BotType       string `json:"bot_type,omitempty"`
	ProductDigest string `json:"image_digest,omitempty"`
}

type PaymentEvent struct {
	BasicEvent        `json:",inline"`
	UserId            string `json:"user_id"`
	Action            string `json:"action"`
	TransactionStatus string `json:"tx_status,omitempty"`
	ProductId         string `json:"pid,omitempty"`
	ProductName       string `json:"product_name,omitempty"`
	ProductImage      string `json:"image_name,omitempty"`
	ProductDigest     string `json:"image_digest,omitempty"`
}
