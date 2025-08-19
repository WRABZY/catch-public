package main

import (
	"bytes"
	"catch/game"
	"fmt"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pm "catch/metrics_main/metrics"

	"google.golang.org/grpc"
)

const (
	methodSendMessage            = "/sendMessage"
	methodDeleteMessage          = "/deleteMessage"
	methodSendPhoto              = "/sendPhoto"
	methodEditMessageMedia       = "/editMessageMedia"
	methodEditMessageReplyMarkup = "/editMessageReplyMarkup"
	methodAnswerCallbackQuery    = "/answerCallbackQuery"
	methodSendInvoice            = "/sendInvoice"
	methodAnswerPreCheckoutQuery = "/answerPreCheckoutQuery"

	typeJSON = "Content-Type: application/json"

	inlineKeyboardClose       = `{"inline_keyboard":[[{"text":"❎","callback_data":"inline0inline"}]]}`
	inlineKeyboardPayTemplate = `{"inline_keyboard":[[{"text":"XTR %d","pay":"True"},{"text":"❎","callback_data":"inline0inline"}]]}`

	inlineKeyboardWait = `{"inline_keyboard":[
	   [{"text":"⏳","callback_data":"inline1inline"},
		{"text":"⏳","callback_data":"inline2inline"},
		{"text":"⏳","callback_data":"inline3inline"}],
	   [{"text":"⏳","callback_data":"inline4inline"},
		{"text":"⏳","callback_data":"inline5inline"},
		{"text":"⏳","callback_data":"inline6inline"}],
	   [{"text":"⏳","callback_data":"inline7inline"},
		{"text":"⏳","callback_data":"inline8inline"},
		{"text":"⏳","callback_data":"inline9inline"}]]}`

	inlineKeyboardShopAliveNewGame = `{"inline_keyboard":[
	   [{"text":"🛍","callback_data":"inline1inline"},
		{"text":"🟥","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🟨","callback_data":"inline4inline"},
		{"text":"💤","callback_data":"inline5inline"},
		{"text":"🟩","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🟦","callback_data":"inline8inline"},
		{"text":"🔒","callback_data":"inline9inline"}]]}`

	inlineKeyboardShopAliveBack = `{"inline_keyboard":[
	   [{"text":"🛍","callback_data":"inline1inline"},
		{"text":"🟥","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🟨","callback_data":"inline4inline"},
		{"text":"💤","callback_data":"inline5inline"},
		{"text":"🟩","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🟦","callback_data":"inline8inline"},
		{"text":"⬅️","callback_data":"inline9inline"}]]}`

	inlineKeyboardShopAliveForward = `{"inline_keyboard":[
	   [{"text":"🛍","callback_data":"inline1inline"},
		{"text":"🔒","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🔒","callback_data":"inline4inline"},
		{"text":"🔒","callback_data":"inline5inline"},
		{"text":"🔒","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🔒","callback_data":"inline8inline"},
		{"text":"➡️","callback_data":"inline9inline"}]]}`

	inlineKeyboardHairAliveNewGame = `{"inline_keyboard":[
	   [{"text":"💇","callback_data":"inline1inline"},
		{"text":"🟥","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🟨","callback_data":"inline4inline"},
		{"text":"💤","callback_data":"inline5inline"},
		{"text":"🟩","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🟦","callback_data":"inline8inline"},
		{"text":"🔒","callback_data":"inline9inline"}]]}`

	inlineKeyboardHairAliveBack = `{"inline_keyboard":[
	   [{"text":"💇","callback_data":"inline1inline"},
		{"text":"🟥","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🟨","callback_data":"inline4inline"},
		{"text":"💤","callback_data":"inline5inline"},
		{"text":"🟩","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🟦","callback_data":"inline8inline"},
		{"text":"⬅️","callback_data":"inline9inline"}]]}`

	inlineKeyboardHairAliveForward = `{"inline_keyboard":[
	   [{"text":"🔒","callback_data":"inline1inline"},
		{"text":"🔒","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🔒","callback_data":"inline4inline"},
		{"text":"🔒","callback_data":"inline5inline"},
		{"text":"🔒","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🔒","callback_data":"inline8inline"},
		{"text":"➡️","callback_data":"inline9inline"}]]}`

	inlineKeyboardShopDeadBack = `{"inline_keyboard":[
	   [{"text":"🛍","callback_data":"inline1inline"},
		{"text":"🔒","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🔒","callback_data":"inline4inline"},
		{"text":"🔒","callback_data":"inline5inline"},
		{"text":"🔒","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🔒","callback_data":"inline8inline"},
		{"text":"⬅️","callback_data":"inline9inline"}]]}`

	inlineKeyboardShopDeadForward = `{"inline_keyboard":[
	   [{"text":"🛍","callback_data":"inline1inline"},
		{"text":"🔒","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🔒","callback_data":"inline4inline"},
		{"text":"🔒","callback_data":"inline5inline"},
		{"text":"🔒","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🔒","callback_data":"inline8inline"},
		{"text":"➡️","callback_data":"inline9inline"}]]}`

	inlineKeyboardHairDeadBack = `{"inline_keyboard":[
	   [{"text":"💇","callback_data":"inline1inline"},
		{"text":"🔒","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🔒","callback_data":"inline4inline"},
		{"text":"🔒","callback_data":"inline5inline"},
		{"text":"🔒","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🔒","callback_data":"inline8inline"},
		{"text":"⬅️","callback_data":"inline9inline"}]]}`

	inlineKeyboardHairDeadForward = `{"inline_keyboard":[
	   [{"text":"🔒","callback_data":"inline1inline"},
		{"text":"🔒","callback_data":"inline2inline"},
		{"text":"🏆","callback_data":"inline3inline"}],
	   [{"text":"🔒","callback_data":"inline4inline"},
		{"text":"🔒","callback_data":"inline5inline"},
		{"text":"🔒","callback_data":"inline6inline"}],
	   [{"text":"❓","callback_data":"inline7inline"},
		{"text":"🔒","callback_data":"inline8inline"},
		{"text":"➡️","callback_data":"inline9inline"}]]}`

	leaderboardMessageTemplate0 = "🥇 %s\n%d\n\n"
	leaderboardMessageTemplate1 = "🥈 %s\n%d\n\n"
	leaderboardMessageTemplate2 = "🥉 %s\n%d\n\n"
	leaderboardMessageTemplate  = "#%2d %s\n%d\n\n"

	callback0 = uint8('0')
	callback1 = uint8('1')
	callback2 = uint8('2')
	callback3 = uint8('3')
	callback4 = uint8('4')
	callback5 = uint8('5')
	callback6 = uint8('6')
	callback7 = uint8('7')
	callback8 = uint8('8')
	callback9 = uint8('9')
	callbackP = uint8('P')

	answerCallbackIsBusyTemplate = `{"callback_query_id":"%d","text":"⏳\nПожалуйста, дождитесь завершения обработки предыдущего запроса","show_alert":true}`
	answerCallbackOkTemplate     = `{"callback_query_id":"%d"}`
	answerCallbackAlertTemplate  = `{"callback_query_id":"%d","text":"%s","show_alert":true}`

	cantMove      = "😿\nCan't move"
	cantMoveUp    = cantMove + " up"
	cantMoveDown  = cantMove + " down"
	cantMoveLeft  = cantMove + " left"
	cantMoveRight = cantMove + " right"
	cantSleep     = "😿\nThere's no time for sleep now"
	cantDoIt      = "😿\nCan't do it now"
)

type GameServer struct {
	databaseMutex         sync.Mutex
	userCacheMutex        sync.RWMutex
	userCache             map[int64]*User
	apiUrl                string
	gameBoofer            *GameBoofer
	lastCleaningDaemonRun time.Time
	leaderboardMutex      sync.RWMutex
	leaderboardNames      [10]string
	leaderboardScores     [10]uint64
	pm.UnimplementedCatchMetricsServer
	price int
}

type User struct {
	id           int64
	name         string
	game         *game.Game
	lastActivity time.Time
	chatId       int64
	messageId    int64
	token        chan struct{}
	gotHair      bool
	thisFrameId  string
	prevFrameId  string
	watchingBack bool
}

type GameBoofer struct {
	m     sync.Mutex
	games []*game.Game
	ptr   int
}

func newGameBoofer() *GameBoofer {
	var initialSize = 100
	var gb = GameBoofer{
		games: make([]*game.Game, initialSize),
	}
	for gb.ptr < initialSize {
		gb.games[gb.ptr] = game.NewGame()
		gb.ptr++
	}
	return &gb
}

func (gb *GameBoofer) getGame() *game.Game {
	gb.m.Lock()
	defer gb.m.Unlock()
	if gb.ptr > 0 {
		gb.ptr--
		return gb.games[gb.ptr]
	}
	return game.NewGame()
}

func (gb *GameBoofer) returnGame(g *game.Game) {
	gb.m.Lock()
	defer gb.m.Unlock()

	if gb.ptr < len(gb.games) {
		gb.games[gb.ptr] = g
	} else {
		gb.games = append(gb.games, g)
	}
	gb.ptr++
}

func (gs *GameServer) startCleaningDaemon() {
	go func() {
		for {
			gs.userCacheMutex.Lock()
			for id, user := range gs.userCache {
				if time.Since(user.lastActivity).Minutes() > 59 {
					select {
					case user.token <- struct{}{}:
					default:
						continue
					}
					// TODO SAVE TO DATABASE
					user.game.ResetGame()
					gs.gameBoofer.returnGame(user.game)
					delete(gs.userCache, id)
				}
			}
			gs.lastCleaningDaemonRun = time.Now()
			gs.userCacheMutex.Unlock()
			time.Sleep(60 * time.Minute)
		}
	}()
}

func (gs *GameServer) startGRPCMetricsServerDaemon(port int) {
	go func(p int) {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", p))
		if err != nil {
			log.Printf("gRPC metrics server daemon failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pm.RegisterCatchMetricsServer(grpcServer, gs)
		grpcServer.Serve(lis)
	}(port)
}

func main() {
	localConfig, err := os.ReadFile("config.local")
	if err != nil {
		log.Fatal(err)
	}

	lineSep := "\r\n"
	lineIndex1 := bytes.Index(localConfig, []byte(lineSep))
	if lineIndex1 == -1 {
		lineSep = "\n"
		lineIndex1 = bytes.Index(localConfig, []byte(lineSep))
	}
	lineSeparatorLen := len(lineSep)
	lineIndex2 := lineIndex1 + lineSeparatorLen + bytes.Index(localConfig[lineIndex1+lineSeparatorLen:], []byte(lineSep))
	lineIndex3 := lineIndex2 + lineSeparatorLen + bytes.Index(localConfig[lineIndex2+lineSeparatorLen:], []byte(lineSep))
	lineIndex4 := lineIndex3 + lineSeparatorLen + bytes.Index(localConfig[lineIndex3+lineSeparatorLen:], []byte(lineSep))

	addr := string(localConfig[len("addr="):lineIndex1])
	path := string(localConfig[lineIndex1+lineSeparatorLen+len("path=") : lineIndex2])
	sysEP := string(localConfig[lineIndex2+lineSeparatorLen+len("sys_endpoint=") : lineIndex3])
	sysToken := string(localConfig[lineIndex3+lineSeparatorLen+len("sys_token=") : lineIndex4])
	price, err := strconv.Atoi(string(localConfig[lineIndex4+lineSeparatorLen+len("price="):]))
	if err != nil {
		price = 100 // default
	}

	server := &GameServer{
		userCache:  make(map[int64]*User),
		apiUrl:     addr + path,
		gameBoofer: newGameBoofer(),
		price:      price,
	}
	server.startCleaningDaemon()
	server.startGRPCMetricsServerDaemon(9890)
	server.initDatabase()
	server.leaderboardNames, server.leaderboardScores = server.loadLeaderboard()

	http.Handle("/", server)
	http.Handle("/"+sysEP+"/", newSysController(sysToken, server))
	log.Fatal(http.ListenAndServe(`:443`, nil))
}

func (gs *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Request body reading error:", err)
		return
	}

	go gs.handleRequest(body)
	w.WriteHeader(http.StatusOK)
}

func (gs *GameServer) handleRequest(body []byte) {
	log.Println("request:", string(body)) // TODO DELETE

	userId, err := getUserId(body)
	if err != nil {
		return
	}

	gs.userCacheMutex.RLock()
	user, ok := gs.userCache[userId]
	gs.userCacheMutex.RUnlock()
	if !ok {
		chId, err := getChatId(body)
		if err != nil {
			// Maybe payment
			preCheckoutQueryId := getPreCheckoutQueryId(body)
			if preCheckoutQueryId == "" {
				return
			}
			gs.sendErrorPreCheckoutQueryAnswer(preCheckoutQueryId, `Для совершения покупок необходимо начать игру.`)
			return
		}
		if bytes.Contains(body, startText) { // == command "New game"
			user = new(User)
			user.token = make(chan struct{}, 1)
			user.token <- struct{}{}

			gs.userCacheMutex.Lock()
			gs.userCache[userId] = user
			gs.userCacheMutex.Unlock()

			user.id = userId
			user.chatId = chId
			user.messageId = -1
			user.name = getName(body)
			user.game = gs.gameBoofer.getGame()
			user.game.HasHair, user.gotHair = gs.loadUserUsingHair(userId)
			user.lastActivity = time.Now()

			user.game.RefreshFrame()
			gs.sendNewGameMessage(user)
			<-user.token
		} else if bytes.Contains(body, wakeUpText) {
			// TODO DB
			return
		}
		incomingMessageId, err := getMessageId(body)
		if err != nil {
			return
		}
		gs.sendDeleteMessage(chId, incomingMessageId)
		return
	}

	callbackId, err := getCallbackId(body)
	select {
	case user.token <- struct{}{}:
	default:
		if err != nil {
			// Maybe payment
			preCheckoutQueryId := getPreCheckoutQueryId(body)
			if preCheckoutQueryId == "" {
				return
			}
			gs.sendErrorPreCheckoutQueryAnswer(preCheckoutQueryId, `Пожалуйста, дождитесь завершения обработки предыдущего запроса`)
			return
		}
		gs.answerCallbackIsBusy(callbackId)
		return
	}

	if err != nil {
		// Maybe payment
		preCheckoutQueryId := getPreCheckoutQueryId(body)
		if preCheckoutQueryId == "" {
			// This is not callback and not payment
			if bytes.Contains(body, startText) {
				thisMessageId, err := getMessageId(body)
				if err == nil {
					gs.sendDeleteMessage(user.chatId, thisMessageId)
				}
				if user.messageId != -1 {
					if user.game.PlayerIsAlive() {
						gs.sendDeleteMessage(user.chatId, user.messageId)
					} else {
						gs.sendDeleteReplyMarkup(user)
					}
				}
				user.game.ResetGame()
				user.game.RefreshFrame()
				gs.sendNewGameMessage(user)
			}
		} else {
			if user.gotHair {
				gs.sendErrorPreCheckoutQueryAnswer(preCheckoutQueryId, `Покупка уже выполнена`)
			} else {
				if ok := gs.updatePaymentsDatabase(user.id, true); ok {
					user.gotHair = true
					user.game.HasHair = true
					user.game.RefreshFrame()
					gs.sendActualFrame(user, true)
					gs.answerOkPreCheckoutQuery(preCheckoutQueryId)
				} else {
					gs.sendErrorPreCheckoutQueryAnswer(preCheckoutQueryId, `Покупка сейчас недоступна, пожалуйста, повторите попытку позднее`)
				}
			}
		}
	} else {
		// This is callback
		gs.handleCallback(user, body, callbackId)
	}

	<-user.token
}

func (gs *GameServer) answerCallbackIsBusy(callbackId int64) {
	resp, _ := http.Post(
		gs.apiUrl+methodAnswerCallbackQuery,
		typeJSON,
		strings.NewReader(fmt.Sprintf(answerCallbackIsBusyTemplate, callbackId)),
	)
	if resp != nil {
		resp.Body.Close()
	}
}

func (gs *GameServer) handleCallback(user *User, body []byte, callbackId int64) {
	callbackData, err := getCallbackData(body)
	if err != nil {
		return
	}
	var alertMessage string
	switch callbackData {
	case callback0: // Close message
		gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
		messageId, err := getMessageId(body)
		if err == nil {
			gs.sendDeleteMessage(user.chatId, messageId)
		}
		return
	case callback1:
		if user.gotHair {
			if user.watchingBack {
				alertMessage = cantDoIt
			} else {
				gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
				user.game.HasHair = !user.game.HasHair
				gs.updatePaymentsDatabase(user.id, user.game.HasHair)
				user.game.RefreshFrame()
				gs.sendActualFrame(user, true)
				return
			}
		} else {
			gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
			gs.sendShop(user)
			return
		}
	case callback2:
		if user.watchingBack || !user.game.MovePlayerTo(game.North) {
			alertMessage = cantMoveUp
		}
	case callback3:
		if gs.leaderboardScores[0] == 0 {
			alertMessage = "No results yet"
		} else {
			gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
			gs.sendLeaderboard(user)
			return
		}
	case callback4:
		if user.watchingBack || !user.game.MovePlayerTo(game.West) {
			alertMessage = cantMoveLeft
		}
	case callback5:
		if user.watchingBack || !user.game.MovePlayerTo(game.NoDirection) {
			alertMessage = cantSleep
		}
	case callback6:
		if user.watchingBack || !user.game.MovePlayerTo(game.East) {
			alertMessage = cantMoveRight
		}
	case callback7:
		gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
		gs.sendHelp(user)
		return
	case callback8:
		if user.watchingBack || !user.game.MovePlayerTo(game.South) {
			alertMessage = cantMoveDown
		}
	case callback9:
		if user.prevFrameId == "" {
			alertMessage = cantDoIt
		} else {
			gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
			user.watchingBack = !user.watchingBack
			if user.watchingBack {
				gs.sendPrevFrameById(user)
			} else {
				gs.sendThisFrameById(user)
			}
			return
		}
	default:
		return
	}

	if alertMessage == "" {
		gs.answerCallback(fmt.Sprintf(answerCallbackOkTemplate, callbackId))
		user.game.RefreshFrame()
		gs.sendActualFrame(user, false)
		if !user.game.TurnOfPlayer {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				user.game.MoveEnemies()
				user.game.RefreshFrame()
			}()
			go func() {
				defer wg.Done()
				time.Sleep(2 * time.Second)
			}()
			wg.Wait()
			gs.sendActualFrame(user, false)
		}
		if !user.game.PlayerIsAlive() {
			gs.insertNewScore(user)
		}
	} else {
		gs.answerCallback(fmt.Sprintf(answerCallbackAlertTemplate, callbackId, alertMessage))
	}
}

func (gs *GameServer) answerCallback(message string) {
	callbackAnswerResponse, _ := http.Post(
		gs.apiUrl+methodAnswerCallbackQuery,
		typeJSON,
		strings.NewReader(message),
	)
	if callbackAnswerResponse != nil {
		callbackAnswerResponse.Body.Close()
	}
}

func (gs *GameServer) sendActualFrame(user *User, hairOnly bool) {
	var frameBuffer bytes.Buffer

	png.Encode(&frameBuffer, user.game.Frame)

	var caption strings.Builder
	for _, event := range user.game.TurnEvents {
		if event != game.EventEmpty {
			caption.WriteString(localizationEvents["ru"][event])
		}
	}
	dataToSend := map[string]io.Reader{
		"media":        strings.NewReader(`{"type":"photo","media":"attach://gameframe","caption":"` + caption.String() + `","show_caption_above_media":true}`),
		"chat_id":      strings.NewReader(strconv.FormatInt(user.chatId, 10)),
		"message_id":   strings.NewReader(strconv.FormatInt(user.messageId, 10)),
		"gameframe":    &frameBuffer,
		"reply_markup": strings.NewReader(gs.getKeyboard(user)),
	}

	var bufferedData bytes.Buffer
	var err error
	wmultipart := multipart.NewWriter(&bufferedData)
	for key, r := range dataToSend {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if key == "photo" || key == "gameframe" {
			if fw, err = wmultipart.CreateFormFile(key, "gameframe"); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = wmultipart.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return
		}

	}
	typeMultipart := "Content-Type: " + wmultipart.FormDataContentType()
	wmultipart.Close()

	resp, _ := http.Post(gs.apiUrl+methodEditMessageMedia, typeMultipart, &bufferedData)
	if resp != nil {
		body, _ := io.ReadAll(resp.Body)
		if !hairOnly {
			user.prevFrameId = user.thisFrameId
		}
		user.thisFrameId = getFileId(body)
		resp.Body.Close()
	}
}

func (gs *GameServer) sendThisFrameById(user *User) {
	dataToSend := `{
		"chat_id":` + fmt.Sprint(user.chatId) + `,
		"message_id":` + fmt.Sprint(user.messageId) + `,
		"media":{"type":"photo","media":"` + user.thisFrameId + `","caption":"` + "Просмотр актуального кадра состояния игры." + `","show_caption_above_media":true},
		"reply_markup":` + gs.getKeyboard(user) + `
	}`

	resp, _ := http.Post(gs.apiUrl+methodEditMessageMedia, typeJSON, strings.NewReader(dataToSend))
	if resp != nil {
		resp.Body.Close()
	}
}

func (gs *GameServer) sendPrevFrameById(user *User) {
	var caption string

	if user.game.PlayerIsAlive() {
		caption = "Просмотр кадра состояния игры до вражеского хода.\nДля продолжения вернитесь к актуальному кадру."
	} else {
		caption = "Просмотр кадра состояния игры до вражеского хода."
	}

	dataToSend := `{
		"chat_id":` + fmt.Sprint(user.chatId) + `,
		"message_id":` + fmt.Sprint(user.messageId) + `,
		"media":{"type":"photo","media":"` + user.prevFrameId + `","caption":"` + caption + `","show_caption_above_media":true},
		"reply_markup":` + gs.getKeyboard(user) + `
	}`

	resp, _ := http.Post(gs.apiUrl+methodEditMessageMedia, typeJSON, strings.NewReader(dataToSend))
	if resp != nil {
		resp.Body.Close()
	}
}

func (gs *GameServer) sendNewGameMessage(user *User) {
	var frameBuffer bytes.Buffer
	png.Encode(&frameBuffer, user.game.Frame)

	var caption strings.Builder
	for _, event := range user.game.TurnEvents {
		if event != game.EventEmpty {
			caption.WriteString(localizationEvents["ru"][event])
		}
	}
	dataToSend := map[string]io.Reader{
		"photo":                    &frameBuffer,
		"chat_id":                  strings.NewReader(strconv.FormatInt(user.chatId, 10)),
		"reply_markup":             strings.NewReader(gs.getKeyboard(user)),
		"caption":                  strings.NewReader(caption.String()),
		"show_caption_above_media": strings.NewReader("true"),
	}

	var bufferedData bytes.Buffer
	var err error
	wmultipart := multipart.NewWriter(&bufferedData)
	for key, r := range dataToSend {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if key == "photo" || key == "gameframe" {
			if fw, err = wmultipart.CreateFormFile(key, "gameframe"); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = wmultipart.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return
		}
	}
	typeMultipart := "Content-Type: " + wmultipart.FormDataContentType()
	wmultipart.Close()

	resp, err := http.Post(gs.apiUrl+methodSendPhoto, typeMultipart, &bufferedData)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp != nil {
		resp.Body.Close()
	}

	user.prevFrameId = ""
	user.thisFrameId = getFileId(body)

	resultMessageId, err := getResultMessageId(body)
	if err != nil {
		return
	}
	user.messageId = resultMessageId
}

func (gs *GameServer) sendDeleteMessage(chatId, messageId int64) {
	resp, err := http.Post(fmt.Sprintf("%s%s?chat_id=%d&message_id=%d", gs.apiUrl, methodDeleteMessage, chatId, messageId), typeJSON, nil)
	if err != nil {
		return
	}
	resp.Body.Close()
}

func (gs *GameServer) sendLeaderboard(user *User) {
	var builder strings.Builder
	builder.WriteString("Лучшие игроки в Catch!\n\n")
	gs.leaderboardMutex.RLock()
	builder.WriteString(fmt.Sprintf(leaderboardMessageTemplate0, gs.leaderboardNames[0], gs.leaderboardScores[0]))
	if gs.leaderboardScores[1] != 0 {
		builder.WriteString(fmt.Sprintf(leaderboardMessageTemplate1, gs.leaderboardNames[1], gs.leaderboardScores[1]))
		if gs.leaderboardScores[2] != 0 {
			builder.WriteString(fmt.Sprintf(leaderboardMessageTemplate2, gs.leaderboardNames[2], gs.leaderboardScores[2]))
			for i := 3; i < len(gs.leaderboardScores); i++ {
				if gs.leaderboardScores[i] == 0 {
					break
				} else {
					builder.WriteString(fmt.Sprintf(leaderboardMessageTemplate, i+1, gs.leaderboardNames[i], gs.leaderboardScores[i]))
				}
			}
		}
	}
	gs.leaderboardMutex.RUnlock()

	dataToSend := `{
		"chat_id":` + fmt.Sprint(user.chatId) + `,
		"text":"` + builder.String() + `",
		"reply_markup":` + inlineKeyboardClose + `
	}`

	resp, err := http.Post(gs.apiUrl+methodSendMessage, typeJSON, strings.NewReader(dataToSend))
	if err != nil {
		return
	}
	resp.Body.Close()
}

func (gs *GameServer) sendHelp(user *User) {
	var text = `*Об игре*
Правила игры: [telegra\\.ph/Catch](https://telegra.ph/Catch-Pravila-igry-08-18)
Автор: [WRABZY](https://t.me/WRABZY)`

	dataToSend := `{
		"chat_id":` + fmt.Sprint(user.chatId) + `,
		"text":"` + text + `",
		"parse_mode":"MarkdownV2",
		"reply_markup":` + inlineKeyboardClose + `,
		"link_preview_options": {"url": "https://telegra.ph/Catch-Pravila-igry-08-18"}
	}`

	resp, err := http.Post(gs.apiUrl+methodSendMessage, typeJSON, strings.NewReader(dataToSend))
	if err != nil {
		return
	}
	bbb, _ := io.ReadAll(resp.Body)
	log.Println(string(bbb))
	resp.Body.Close()
}

func (gs *GameServer) sendShop(user *User) {
	var title = `Лавандовое каре`
	var description = `Вы можете поддержать разработчика, купив для главной фигуры игры альтернативный образ`
	var inlineKeyboardPay = fmt.Sprintf(inlineKeyboardPayTemplate, gs.price)
	var prices = fmt.Sprintf(`[{"label":"stars","amount":%d}]`, gs.price)

	dataToSend := `{
		"chat_id":` + fmt.Sprint(user.chatId) + `,
		"title":"` + title + `",
		"description":"` + description + `",
		"photo_url": "https://raw.githubusercontent.com/WRABZY/catch-public/refs/heads/main/promo/promo.png",
		"photo_width":300,
		"photo_height":300,
		"payload":"inlinePinline",
		"currency":"XTR",
		"prices":` + prices + `,
		"reply_markup":` + inlineKeyboardPay + `
	}`

	resp, err := http.Post(gs.apiUrl+methodSendInvoice, typeJSON, strings.NewReader(dataToSend))
	if err != nil {
		return
	}
	resp.Body.Close()
}

func (gs *GameServer) sendErrorPreCheckoutQueryAnswer(preCheckoutQueryId, reason string) {
	dataToSend := `{
		"pre_checkout_query_id":"` + preCheckoutQueryId + `",
		"ok":false,
		"error_message":"` + reason + `"
	}`

	resp, err := http.Post(gs.apiUrl+methodAnswerPreCheckoutQuery, typeJSON, strings.NewReader(dataToSend))
	if err != nil {
		return
	}
	resp.Body.Close()
}

func (gs *GameServer) answerOkPreCheckoutQuery(preCheckoutQueryId string) {
	dataToSend := `{
		"pre_checkout_query_id":"` + preCheckoutQueryId + `",
		"ok":true
	}`

	resp, err := http.Post(gs.apiUrl+methodAnswerPreCheckoutQuery, typeJSON, strings.NewReader(dataToSend))
	if err != nil {
		return
	}
	resp.Body.Close()
}

func (gs *GameServer) sendDeleteReplyMarkup(user *User) {
	dataToSend := `{
		"chat_id":` + fmt.Sprint(user.chatId) + `,
		"message_id":` + fmt.Sprint(user.messageId) + `
	}`

	resp, err := http.Post(gs.apiUrl+methodEditMessageReplyMarkup, typeJSON, strings.NewReader(dataToSend))
	if err != nil {
		return
	}
	resp.Body.Close()
}

func (gs *GameServer) insertNewScore(user *User) {
	gs.leaderboardMutex.Lock()
	defer gs.leaderboardMutex.Unlock()
	i := len(gs.leaderboardScores) - 1
	newScore := user.game.GetScore()
	if gs.leaderboardScores[i] < newScore {
		for ; i > 0 && gs.leaderboardScores[i-1] < newScore; i-- {
			gs.leaderboardScores[i] = gs.leaderboardScores[i-1]
			gs.leaderboardNames[i] = gs.leaderboardNames[i-1]
		}
		gs.leaderboardScores[i] = newScore
		gs.leaderboardNames[i] = user.name
		gs.updateLeaderboardDatabase(true)
	}
}

func (gs *GameServer) getKeyboard(user *User) string {
	if user.game.TurnOfPlayer {
		if user.prevFrameId == "" {
			if user.gotHair {
				return inlineKeyboardHairAliveNewGame
			} else {
				return inlineKeyboardShopAliveNewGame
			}
		} else {
			if user.game.PlayerIsAlive() {
				if user.gotHair {
					if user.watchingBack {
						return inlineKeyboardHairAliveForward
					} else {
						return inlineKeyboardHairAliveBack
					}
				} else {
					if user.watchingBack {
						return inlineKeyboardShopAliveForward
					} else {
						return inlineKeyboardShopAliveBack
					}
				}
			} else {
				if user.gotHair {
					if user.watchingBack {
						return inlineKeyboardHairDeadForward
					} else {
						return inlineKeyboardHairDeadBack
					}
				} else {
					if user.watchingBack {
						return inlineKeyboardShopDeadForward
					} else {
						return inlineKeyboardShopDeadBack
					}
				}
			}
		}
	}
	return inlineKeyboardWait
}
