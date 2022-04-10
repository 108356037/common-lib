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
	K8sObjectKind string `json:"k8s_obj_kind"`
}

type StrategyEvent struct {
	BasicEvent `json:",inline"`
	//EventType          string                 `json:"event_type"`
	UserId               string                 `json:"user_id"`
	ResourceType         string                 `json:"resource_type"`
	Action               string                 `json:"action"`
	StrategyName         string                 `json:"strategy_name"`
	StrategyUpdateDetail map[string]interface{} `json:"update_detail,omitempty"`
}

type ProductEvent struct {
	BasicEvent    `json:",inline"`
	UserId        string `json:"user_id"`
	ResourceType  string `json:"resource_type"`
	Action        string `json:"action"`
	ProductName   string `json:"product_name"`
	ProductImage  string `json:"image_name"`
	ProductDigest string `json:"image_digest,omitempty"`
}
