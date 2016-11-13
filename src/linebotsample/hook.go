package linebot

import (
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func hookHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	cf, err := LoadConfig(c)
	if err != nil {
		msg := err.Error()
		log.Infof(c, msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	client := urlfetch.Client(c)
	cl, err := linebot.New(cf.Secret, cf.Token, linebot.WithHTTPClient(client))
	if err != nil {
		msg := err.Error()
		log.Infof(c, msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	ev, err := cl.ParseRequest(r)
	if err != nil {
		msg := err.Error()
		log.Infof(c, msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	for _, e := range ev {
		log.Infof(c, "Event: %#v", e)
		if e.Type != linebot.EventTypeMessage {
			continue
		}

		switch m := e.Message.(type) {
		case *linebot.TextMessage:
			log.Infof(c, "TextMessage %#v", m)
			if strings.Contains(m.Text, "Go") {
				_, err := cl.ReplyMessage(e.ReplyToken, linebot.NewTextMessage("ʕ ◔ϖ◔ʔ Go~")).Do()
				if err != nil {
					log.Errorf(c, "send error: %v", err)
				}
			}
		}
	}
}
