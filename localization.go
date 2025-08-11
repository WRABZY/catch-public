package main

import (
	"catch/game"
	"fmt"
)

var ruEventPawnEatenMessage = fmt.Sprintf("💰 Пешка съедена: +%d очков\n", game.PawnScore)
var ruEventRookEatenMessage = fmt.Sprintf("💰 Ладья съедена: +%d очков\n", game.RookScore)
var ruEventBishopEatenMessage = fmt.Sprintf("💰 Слон съеден: +%d очков\n", game.BishopScore)
var ruEventQueenEatenMessage = fmt.Sprintf("💰 Королева съедена: +%d очков\n", game.QueenScore)
var ruEventKingEatenMessage = fmt.Sprintf("💰 Король съеден: +%d очков\n", game.KingScore)
var ruEventCoinFound = fmt.Sprintf("💰 Найдена ржавая золотая монетка: +%d очков\n", game.CoinScore)

var localizationEvents = map[string]map[int8]string{
	"ru": map[int8]string{
		game.EventNewGame:               "😺 Новая игра\n",
		game.EventEndGame:               "😿 Игра окончена\n",
		game.EventFromPawnToQueen:       "👸 Пешка становится Королевой\n",
		game.EventPawnAttacks:           "🪖 Пешка атакует и погибает от ответного урона\n",
		game.EventPawnAttacksLast:       "🪖 Пешка наносит фатальный удар\n",
		game.EventRookAttacks:           "💀 Ладья атакует и получает урон в ответ\n",
		game.EventDamagedRookAttacks:    "💀 Ладья атакует и погибает от ответного урона\n",
		game.EventRookAttacksLast:       "💀 Ладья наносит фатальный удар\n",
		game.EventBishopAttacks:         "🐘 Слон атакует и получает урон в ответ\n",
		game.EventDamagedBishopAttacks:  "🐘 Слон атакует и погибает от ответного урона\n",
		game.EventBishopAttacksLast:     "🐘 Слон наносит фатальный удар\n",
		game.EventQueenAttacks:          "👸 Королева атакует и получает урон в ответ\n",
		game.EventDamagedQueenAttacks:   "👸 Королева атакует и погибает от ответного урона\n",
		game.EventQueenAttacksLast:      "👸 Королева наносит фатальный удар\n",
		game.EventKingAttacks:           "🤴 Король атакует\n",
		game.EventKingAttacksLast:       "🤴 Король наносит фатальный удар\n",
		game.EventPawnEaten:             ruEventPawnEatenMessage,
		game.EventRookEaten:             ruEventRookEatenMessage,
		game.EventBishopEaten:           ruEventBishopEatenMessage,
		game.EventQueenEaten:            ruEventQueenEatenMessage,
		game.EventKingEaten:             ruEventKingEatenMessage,
		game.EventVileStenchAttacks:     "😱 Эхо войны наносит урон\n",
		game.EventVileStenchAttacksLast: "😱 Эхо войны наносит фатальный урон\n",
		game.EventBishopHeals:           "💚 Энергия Слона исцеляет\n",
		game.EventRookHeals:             "💚 Энергия павшей Ладьи исцеляет\n",
		game.EventQueenHeals:            "💚 Энергия павшей Королевы исцеляет\n",
		game.EventSkipTurn:              "💤 Пропуск хода...\n",
		game.EventPawnInserted:          "🪖 Пешка вступает в бой\n",
		game.EventRookInserted:          "💀 Ладья вступает в бой\n",
		game.EventBishopInserted:        "🐘 Слон вступает в бой\n",
		game.EventKingInserted:          "🤴 Король вступает в бой\n",
		game.EventCoinFound:             ruEventCoinFound,
	},
}
