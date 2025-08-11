package main

import (
	"errors"
	"strconv"
)

const (
	userIdErrorMessage          = "error while extracting user_id"
	chatIdErrorMessage          = "error while extracting chat_id"
	callbackIdErrorMessage      = "error while extracting callback_id"
	messageIdErrorMessage       = "error while extracting message_id"
	resultMessageIdErrorMessage = "error while extracting message_id of sended message"
	callbackDataErrorMessage    = "error while extracting callback data"
)

var (
	// not public
)

func getUserId(body []byte) (int64, error) {
	// not public
}

func getName(body []byte) string {
	// not public
}

func getChatId(body []byte) (int64, error) {
	// not public
}

func getCallbackId(body []byte) (int64, error) {
	// not public
}

func getMessageId(body []byte) (int64, error) {
	// not public
}

func getResultMessageId(body []byte) (int64, error) {
	// not public
}

func getCallbackData(body []byte) (byte, error) {
	// not public
}

func getFileId(body []byte) string {
	// not public
}
