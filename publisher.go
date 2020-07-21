package mqtt

// 通用发布消息接口
func (m Client) Publish(topic string, payload interface{}, qos byte, retained bool) (err error) {
	token := m.client.Publish(topic, qos, retained, payload)
	if token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return
}
