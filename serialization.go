package main

import (
	"bytes"
	"catch/game"
	"encoding/gob"
)

type SUser struct {
	Id           int64
	Name         string
	Game         []byte
	ChatId       int64
	MessageId    int64
	GotHair      bool
	ThisFrameId  string
	PrevFrameId  string
	WatchingBack bool
	Lang         string
	Action       string
	BotAction    string
}

func (u *User) ToSerial() []byte {
	suser := SUser{
		Id:           u.id,
		Name:         u.name,
		Game:         u.game.ToSerial(),
		ChatId:       u.chatId,
		MessageId:    u.messageId,
		GotHair:      u.gotHair,
		ThisFrameId:  u.thisFrameId,
		PrevFrameId:  u.prevFrameId,
		WatchingBack: u.watchingBack,
		Lang:         u.lang,
		Action:       u.action,
		BotAction:    u.botAction,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(suser)
	return buf.Bytes()
}

func (u *User) RestoreFrom(b []byte, game *game.Game) {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	var su = new(SUser)
	err := dec.Decode(su)
	if err != nil {
		return
	}
	u.id = su.Id
	u.name = su.Name
	u.game = game
	u.chatId = su.ChatId
	u.messageId = su.MessageId
	u.gotHair = su.GotHair
	u.thisFrameId = su.ThisFrameId
	u.prevFrameId = su.PrevFrameId
	u.watchingBack = su.WatchingBack
	u.lang = su.Lang
	u.action = su.Action
	u.botAction = su.BotAction
	game.RestoreFrom(su.Game)
}
