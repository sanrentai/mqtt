package mqtt

// MQTT的配置信息格式
type Config struct {
	Broker   string `json:"Broker"`   // Broker地址，例如tcp://127.0.0.1:1883
	ClientID string `json:"ClientID"` // ClientID
}
