package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "modernc.org/sqlite"
)

const (
	sqlDriverName              = "sqlite"
	dsnURI                     = "server_data"
	leaderboardTableName       = "leaderboard"
	leaderboardIdColumnName    = "id"
	leaderboardNameColumnName  = "name"
	leaderboardScoreColumnName = "score"

	paymentsTableName           = "payments"
	paymentsIdColumnName        = "id"
	paymentsUsingHairColumnName = "using_hair"
)

var (
	createLeaderboardTableQuery = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s INTEGER PRIMARY KEY NOT NULL, %s TEXT NOT NULL, %s TEXT NOT NULL);", leaderboardTableName, leaderboardIdColumnName, leaderboardNameColumnName, leaderboardScoreColumnName)
	updateLeaderboardRowQuery   = fmt.Sprintf("INSERT OR REPLACE INTO %s (%s, %s, %s) VALUES (?, ?, ?);", leaderboardTableName, leaderboardIdColumnName, leaderboardNameColumnName, leaderboardScoreColumnName)
	selectAllLeaderboardQuery   = fmt.Sprintf("SELECT * FROM %s", leaderboardTableName)

	createPaymentsTableQuery    = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s INTEGER PRIMARY KEY NOT NULL, %s TEXT NOT NULL);", paymentsTableName, paymentsIdColumnName, paymentsUsingHairColumnName)
	updatePaymentsRowQuery      = fmt.Sprintf("INSERT OR REPLACE INTO %s (%s, %s) VALUES (?, ?);", paymentsTableName, paymentsIdColumnName, paymentsUsingHairColumnName)
	selectInfoFromPaymentsQuery = fmt.Sprintf("SELECT %s FROM %s WHERE %s=?;", paymentsUsingHairColumnName, paymentsTableName, paymentsIdColumnName)
)

func (gs *GameServer) initDatabase() {
	db, err := sql.Open(sqlDriverName, dsnURI)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return
	}

	db.Exec(createLeaderboardTableQuery)
	db.Exec(createPaymentsTableQuery)
}

func (gs *GameServer) updateLeaderboardDatabase(retry bool) {
	gs.databaseMutex.Lock()
	defer gs.databaseMutex.Unlock()
	db, err := sql.Open(sqlDriverName, dsnURI)
	if err != nil {
		if retry {
			gs.updateLeaderboardDatabase(false)
		}
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		if retry {
			gs.updateLeaderboardDatabase(false)
		}
		return
	}

	for i, name := range gs.leaderboardNames {
		db.Exec(updateLeaderboardRowQuery, i, name, strconv.FormatUint(gs.leaderboardScores[i], 10))
	}
}

func (gs *GameServer) loadLeaderboard() (names [10]string, scores [10]uint64) {
	gs.databaseMutex.Lock()
	defer gs.databaseMutex.Unlock()
	db, err := sql.Open(sqlDriverName, dsnURI)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return
	}

	rows, err := db.Query(selectAllLeaderboardQuery)
	if err != nil {
		return
	}
	defer rows.Close()

	i := new(int64)
	n := new(string)
	s := new(string)
	for rows.Next() {
		err = rows.Scan(i, n, s)
		if err != nil || *i < 0 || int(*i) > len(names) || int(*i) > len(scores) {
			return
		}

		sc, err := strconv.ParseUint(*s, 10, 64)
		if err != nil {
			return
		}

		names[*i] = *n
		scores[*i] = sc
	}
	return
}

func (gs *GameServer) updatePaymentsDatabase(id int64, usingHair bool) bool {
	gs.databaseMutex.Lock()
	defer gs.databaseMutex.Unlock()
	db, err := sql.Open(sqlDriverName, dsnURI)
	if err != nil {
		return false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false
	}

	_, err = db.Exec(updatePaymentsRowQuery, id, usingHair)
	return err == nil
}

func (gs *GameServer) loadUserUsingHair(id int64) (usingHair, ok bool) {
	gs.databaseMutex.Lock()
	defer gs.databaseMutex.Unlock()
	db, err := sql.Open(sqlDriverName, dsnURI)
	if err != nil {
		return false, false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false, false
	}

	using := new(string)
	row := db.QueryRow(selectInfoFromPaymentsQuery, id)
	err = row.Scan(using)
	if err != nil {
		return false, false
	}

	return *using == "1", true
}
