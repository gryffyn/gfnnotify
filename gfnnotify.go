package gfnnotify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Notifier struct {
	apptoken string
	endpoint string
}

type Msg struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Priority string `json:"priority"`
	Extras   struct {
		ClientDisplay struct {
			ContentType string `json:"contentType"`
		} `json:"client::display"`
	} `json:"extras"`
}

func (m *Notifier) SendMarkdownMsg(msg Msg) (int, error) {
	msg.Extras.ClientDisplay.ContentType = "text/markdown"
	js, err := json.Marshal(msg)
	form, err := http.Post(m.endpoint+"/message?token="+m.apptoken, "application/json", bytes.NewBuffer(js))
	defer form.Body.Close()
	return form.StatusCode, err
}

func (m *Notifier) SendPlainMsg(msg Msg) (int, error) {
	data := url.Values{
		"title":    {msg.Title},
		"message":  {msg.Message},
		"priority": {msg.Priority},
	}
	form, err := http.PostForm(m.endpoint+"/message?token="+m.apptoken, data)
	defer form.Body.Close()
	return form.StatusCode, err
}
