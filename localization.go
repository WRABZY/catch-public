package main

import (
	"catch/game"
	"fmt"
)

const defaultLang = "en"

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
		game.EventPawnEaten:             fmt.Sprintf("💰 Пешка съедена: +%d очков\n", game.PawnScore),
		game.EventRookEaten:             fmt.Sprintf("💰 Ладья съедена: +%d очков\n", game.RookScore),
		game.EventBishopEaten:           fmt.Sprintf("💰 Слон съеден: +%d очков\n", game.BishopScore),
		game.EventQueenEaten:            fmt.Sprintf("💰 Королева съедена: +%d очков\n", game.QueenScore),
		game.EventKingEaten:             fmt.Sprintf("💰 Король съеден: +%d очков\n", game.KingScore),
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
		game.EventCoinFound:             fmt.Sprintf("💰 Найдена ржавая золотая монетка: +%d очков\n", game.CoinScore),
	},
	"en": map[int8]string{
		game.EventNewGame:               "😺 New game\n",
		game.EventEndGame:               "😿 Game over\n",
		game.EventFromPawnToQueen:       "👸 The Pawn becomes a Queen\n",
		game.EventPawnAttacks:           "🪖 The Pawn attacks and dies from retaliatory damage\n",
		game.EventPawnAttacksLast:       "🪖 The Pawn deals the fatal blow\n",
		game.EventRookAttacks:           "💀 The Rook attacks and takes retaliatory damage\n",
		game.EventDamagedRookAttacks:    "💀 The Rook attacks and dies from retaliatory damage\n",
		game.EventRookAttacksLast:       "💀 The Rook deals the fatal blow\n",
		game.EventBishopAttacks:         "🐘 The Bishop attacks and takes retaliatory damage\n",
		game.EventDamagedBishopAttacks:  "🐘 The Bishop attacks and dies from retaliatory damage\n",
		game.EventBishopAttacksLast:     "🐘 The Bishop deals the fatal blow\n",
		game.EventQueenAttacks:          "👸 The Queen attacks and takes retaliatory damage\n",
		game.EventDamagedQueenAttacks:   "👸 The Queen attacks and dies from retaliatory damage\n",
		game.EventQueenAttacksLast:      "👸 The Queen deals the fatal blow\n",
		game.EventKingAttacks:           "🤴 The King attacks\n",
		game.EventKingAttacksLast:       "🤴 The King deals the fatal blow\n",
		game.EventPawnEaten:             fmt.Sprintf("💰 The Pawn is eaten: +%d points\n", game.PawnScore),
		game.EventRookEaten:             fmt.Sprintf("💰 The Rook is eaten: +%d points\n", game.RookScore),
		game.EventBishopEaten:           fmt.Sprintf("💰 The Bishop is eaten: +%d points\n", game.BishopScore),
		game.EventQueenEaten:            fmt.Sprintf("💰 The Queen is eaten: +%d points\n", game.QueenScore),
		game.EventKingEaten:             fmt.Sprintf("💰 The King is eaten: +%d points\n", game.KingScore),
		game.EventVileStenchAttacks:     "😱 Echo of War deals damage\n",
		game.EventVileStenchAttacksLast: "😱 Echo of War deals the fatal damage\n",
		game.EventBishopHeals:           "💚 The energy of the Bishop heals\n",
		game.EventRookHeals:             "💚 The energy of the fallen Rook heals\n",
		game.EventQueenHeals:            "💚 The energy of the fallen Queen heals\n",
		game.EventSkipTurn:              "💤 Skipping a turn...\n",
		game.EventPawnInserted:          "🪖 The Pawn enters the battle\n",
		game.EventRookInserted:          "💀 The Rook enters the battle\n",
		game.EventBishopInserted:        "🐘 The Bishop enters the battle\n",
		game.EventKingInserted:          "🤴 The King enters the battle\n",
		game.EventCoinFound:             fmt.Sprintf("💰 Rusty gold coin found: +%d points\n", game.CoinScore),
	},
}

var answerCallbackIsBusyText = map[string]string{
	"ru": "⏳\nПожалуйста, дождитесь ответа на предыдущее действие",
	"en": "⏳\nPlease wait for a response to the previous action",
}

var (
	answerCantMoveUp = map[string]string{
		"ru": "😿\nДвижение вверх недоступно",
		"en": "😿\nUpward movement is not available",
	}
	answerСantMoveDown = map[string]string{
		"ru": "😿\nДвижение вниз недоступно",
		"en": "😿\nDownward movement is not available",
	}
	answerСantMoveLeft = map[string]string{
		"ru": "😿\nДвижение влево недоступно",
		"en": "😿\nLeft movement is not available",
	}
	answerCantMoveRight = map[string]string{
		"ru": "😿\nДвижение вправо недоступно",
		"en": "😿\nRight movement is not available",
	}
	answerCantSleep = map[string]string{
		"ru": "😿\nСейчас не время спать",
		"en": "😿\nNow is not the time to sleep",
	}
	answerCantDoIt = map[string]string{
		"ru": "😿\nЭта функция пока недоступна",
		"en": "😿\nThis feature is not available yet",
	}
	answerNoResults = map[string]string{
		"ru": "Таблица лучших игроков пуста",
		"en": "The top players table is empty",
	}
	answerBoughtAlready = map[string]string{
		"ru": "Покупка уже выполнена",
		"en": "The purchase has already been completed",
	}
	answerBoughtIsNotAvailable = map[string]string{
		"ru": "Покупка сейчас недоступна. Пожалуйста, повторите попытку позднее",
		"en": "Purchase is currently unavailable. Please try again later",
	}
	answerStartForBuy = map[string]string{
		"ru": "Для совершения покупок необходимо начать игру",
		"en": "To make purchases you need to start the game",
	}
)

var (
	actionPathRed = map[string]string{
		"ru": "Действие: 🟥\n",
		"en": "Action: 🟥\n",
	}
	actionPathYellow = map[string]string{
		"ru": "Действие: 🟨\n",
		"en": "Action: 🟨\n",
	}
	actionPathGreen = map[string]string{
		"ru": "Действие: 🟩\n",
		"en": "Action: 🟩\n",
	}
	actionPathBlue = map[string]string{
		"ru": "Действие: 🟦\n",
		"en": "Action: 🟦\n",
	}
	actionSleep = map[string]string{
		"ru": "Действие: 💤\n",
		"en": "Action: 💤\n",
	}
	actionWatchBackDead = map[string]string{
		"ru": "👁 Взгляд в прошлое\n",
		"en": "👁 A look into the past\n",
	}
	actionWatchBackAlive = map[string]string{
		"ru": "👁 Взгляд в прошлое.\nДля продолжения игры вернитесь в настоящее нажатием на ➡️\n",
		"en": "👁 A look into the past.\nTo continue the game, return to the present by pressing the ➡️\n",
	}
	actionHaircut = map[string]string{
		"ru": "💅 Изменён стиль\n",
		"en": "💅 Style changed\n",
	}
	actionBot = map[string]string{
		"ru": "🤖 Бот завершил ход\n",
		"en": "🤖 The bot has completed its turn\n",
	}
)

var haircutName = map[string]string{
	"ru": "Лавандовое каре",
	"en": "Lavender bob",
}

var shopDescription = map[string]string{
	"ru": "Вы можете поддержать разработчика, купив для главной фигуры игры альтернативный образ",
	"en": "You can support the developer by purchasing an alternative image for the game's main figure",
}

var leaderboardHead = map[string]string{
	"ru": "Лучшие игроки в Catch!\n\n",
	"en": "Best Catch! players\n\n",
}

var helpHead = map[string]string{
	"ru": "*Об игре*\n",
	"en": "*About*\n",
}

var helpLink = map[string]string{
	"ru": "Правила игры: [telegra\\\\.ph/Catch](https://telegra.ph/Catch-Pravila-igry-08-18)\n",
	"en": "Game rules: [telegra\\\\.ph/Catch](https://telegra.ph/Catch-Pravila-igry-08-18)\n",
}

var authorLink = map[string]string{
	"ru": "Канал автора: [WRABZY](https://t.me/WRABZY)\n",
	"en": "Author's channel: [WRABZY](https://t.me/WRABZY)\n",
}
