package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
	"strings"
	"net/http"
)

var (
	messageAddPath = "/message/add"
	messageUpdatePath = "/message/update"
	messageDeletePath = "/message/delete"
	messageStatusPath = "/message/status"
	messageListPath = "/message/list"
	getMessageByNamePath = "/message/getMessageByName"
	getConsumersByNamePath = "/message/getConsumersByName"
	messageReload = "/message/reload"
)

func NewMessageByNode(node map[string]string) *Message {
	return NewMessage(node["manager_uri"], node["token_header_name"], node["token"], node["publish_uri"])
}

func NewMessage(managerUri string, tokenHeader string, token string, publishUri string) *Message {
	return &Message{
		ManagerUri: managerUri,
		PublishUri: publishUri,
		TokenHeaderName: tokenHeader,
		Token: token,
	}
}

type Message struct {
	ManagerUri string
	PublishUri string
	TokenHeaderName string
	Token string
}

func (m *Message) GetMessages() (messages []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageListPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpGet(url, nil, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return messages, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return messages, errors.New(fmt.Sprintf(v["message"].(string)))
	}
	for _, items := range v["data"].([]interface{}) {
		messages = append(messages, items.(map[string]interface{}))
	}
	return messages, nil
}

func (m *Message) AddMessage(message map[string]string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageAddPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpPost(url, message, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return nil
}

func (m *Message) UpdateMessage(message map[string]string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageUpdatePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpPost(url, message, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return nil
}

func (m *Message) DeleteMessage(name string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageDeletePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return nil
}

func (m *Message) GetMessageByName(name string) (message map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, getMessageByNamePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return message, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return message, errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return v["data"].(map[string]interface{}), nil
}

func (m *Message) GetConsumersByName(name string) (consumers []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, getConsumersByNamePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return consumers, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return consumers, errors.New(fmt.Sprintf(v["message"].(string)))
	}
	for _, items := range v["data"].([]interface{}) {
		consumers = append(consumers, items.(map[string]interface{}))
	}

	return
}

func (m *Message) ReloadMessage(name string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageReload)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return nil
}

func (m *Message) GetConsumersStatus(name string) (consumerStatus []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageStatusPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}
	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return consumerStatus, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return consumerStatus, errors.New(fmt.Sprintf(v["message"].(string)))
	}
	for _, items := range v["data"].([]interface{}) {
		if items == nil {
			continue
		}
		consumerStatus = append(consumerStatus, items.(map[string]interface{}))
	}
	return
}

func (m *Message) Publish(name, method, data, tokenHeader, routeKeyHeader, routeKey string) (err error) {

	message, err := m.GetMessageByName(name)
	if err != nil {
		return
	}
	headerValues := map[string]string{}
	if message["is_need_token"].(bool) {
		headerValues[tokenHeader] = message["token"].(string)
	}
	if routeKey != "" {
		headerValues[routeKeyHeader] = routeKey
	}
	url := fmt.Sprintf("%s/publish/%s", m.PublishUri, name)

	var req *http.Request
	if method == "get" {
		if !strings.Contains(url, "?") {
			url += "?"
		}
		req, err = http.NewRequest("GET", url+data, nil)
	}else {
		req, err = http.NewRequest("POST", url, strings.NewReader(data))
	}
	if err != nil {
		return
	}
	if len(headerValues) > 0 {
		for key, value := range headerValues {
			req.Header.Set(key, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	code := resp.StatusCode
	defer resp.Body.Close()
	if code != 200 {
		return errors.New(fmt.Sprintf("request status: %d", code))
	}

	return nil
}