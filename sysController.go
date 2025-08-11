package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type SysContoller struct {
	token  string
	server *GameServer
}

func newSysController(token string, server *GameServer) *SysContoller {
	return &SysContoller{token: token, server: server}
}

func (sc *SysContoller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Header["Authorization"][0][7:] == sc.token && r.Method == "GET" { // [7:] > "Bearer ..."
		switch {
		case strings.HasSuffix(r.URL.Path, "/status"):
			answer := `{
				"number_of_active_users":` + fmt.Sprint(len(sc.server.userCache)) + `,
				"game_boofer_games_available":` + fmt.Sprint(sc.server.gameBoofer.ptr) + `,
				"game_boofer_size":` + fmt.Sprint(len(sc.server.gameBoofer.games)) + `,
				"last_cleaning_daemon_run_minutes_ago":` + fmt.Sprint(int(time.Since(sc.server.lastCleaningDaemonRun).Minutes())) + `
			}`
			w.Write([]byte(answer))
		}
	}
}
