package waMsgTransport

import (
	"github.com/KuntilBogel/waSocket/proto/armadilloutil"
	"github.com/KuntilBogel/waSocket/proto/waMsgApplication"
)

const (
	MessageApplicationVersion = 2
)

func (msg *MessageTransport_Payload) Decode() (*waMsgApplication.MessageApplication, error) {
	return armadilloutil.Unmarshal(&waMsgApplication.MessageApplication{}, msg.GetApplicationPayload(), MessageApplicationVersion)
}

func (msg *MessageTransport_Payload) Set(payload *waMsgApplication.MessageApplication) (err error) {
	msg.ApplicationPayload, err = armadilloutil.Marshal(payload, MessageApplicationVersion)
	return
}
