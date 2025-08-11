package game

import (
	"image"
	"math/rand/v2"
)

type Game struct {
	Frame               *image.NRGBA
	HasHair             bool
	TurnOfPlayer        bool
	TurnEvents          []int8
	nextTurnEvent       int
	field               *[FieldSide][FieldSide]int8
	enemiesXYNotMoved   map[int8]map[int8]bool
	hp                  int8
	deltaHP             int8
	damaged             bool
	attackFromNorth     bool
	attackFromNorthEast bool
	attackFromEast      bool
	attackFromSouthEast bool
	attackFromSouth     bool
	attackFromSouthWest bool
	attackFromWest      bool
	attackFromNorthWest bool
	steps               uint64
	score               uint64
	deltaScore          uint64
	direction           int8
	dirSteps            int8
	x                   int8
	y                   int8
	yNorth              int8
	xEast               int8
	ySouth              int8
	xWest               int8
	nextEnemyDir        int8
	nextEnemyPos        int8
}

func NewGame() *Game {
	var game = &Game{
		Frame:             image.NewNRGBA(image.Rect(0, 0, GraphicSidePixels, GraphicSidePixels)),
		TurnOfPlayer:      true,
		field:             &[FieldSide][FieldSide]int8{},
		enemiesXYNotMoved: make(map[int8]map[int8]bool),
		hp:                HPMax,
		direction:         NoDirection,
		x:                 Center,
		y:                 Center,
		yNorth:            Center - 1,
		xEast:             Center + 1,
		ySouth:            Center + 1,
		xWest:             Center - 1,
		TurnEvents:        make([]int8, OneTurnMaxEvents),
	}

	var i int8
	for i = Zero; i < FieldSide; i++ {
		game.field[i] = [FieldSide]int8{}
		game.enemiesXYNotMoved[i] = make(map[int8]bool)
	}
	game.field[game.y][game.x] = Player

	game.nextEnemyDir = int8(rand.IntN(DirectionsNumberHalf) + North)
	game.nextEnemyPos = int8(rand.IntN(FieldSide))

	game.addEvent(EventNewGame)

	game.insert(PawnNorth)
	game.addEvent(EventPawnInserted)

	return game
}

func (g *Game) ResetGame() {
	g.TurnOfPlayer = true
	for i := range g.TurnEvents {
		g.TurnEvents[i] = EventEmpty
	}
	g.nextTurnEvent = 0
	var i int8
	for i = Zero; i < FieldSide; i++ {
		for j := range g.field[i] {
			g.field[i][j] = Nobody
		}
		g.enemiesXYNotMoved[i] = make(map[int8]bool)
	}
	g.field[Center][Center] = Player
	g.hp = HPMax
	g.deltaHP = Zero
	g.damaged = false
	g.attackFromNorth = false
	g.attackFromNorthEast = false
	g.attackFromEast = false
	g.attackFromSouthEast = false
	g.attackFromSouth = false
	g.attackFromSouthWest = false
	g.attackFromWest = false
	g.attackFromNorthWest = false
	g.steps = Zero
	g.score = Zero
	g.deltaScore = Zero
	g.direction = NoDirection
	g.dirSteps = Zero
	g.x = Center
	g.y = Center
	g.yNorth = Center - 1
	g.xEast = Center + 1
	g.ySouth = Center + 1
	g.xWest = Center - 1
	g.nextEnemyDir = int8(rand.IntN(DirectionsNumberHalf) + North)
	g.nextEnemyPos = int8(rand.IntN(FieldSide))
	g.addEvent(EventNewGame)
	g.insert(PawnNorth)
	g.addEvent(EventPawnInserted)
}

func (g *Game) PlayerIsAlive() bool {
	return g.hp > 0
}

func (g *Game) GetScore() uint64 {
	return g.score
}

func (g *Game) MoveEnemies() {
	if !g.TurnOfPlayer && g.hp > Zero {
		g.switchTurnParameters()
		g.cleanMoves()
		g.cleanShadows()
		for x, m := range g.enemiesXYNotMoved {
			for y, notMoved := range m {
				if notMoved {
					switch g.field[y][x] {
					case PawnNorth:
						g.movePawnNorthFrom(x, y)
					case PawnEast:
						g.movePawnEastFrom(x, y)
					case PawnSouth:
						g.movePawnSouthFrom(x, y)
					case PawnWest:
						g.movePawnWestFrom(x, y)
					case Rook, DamagedRook:
						g.moveRook(x, y)
					case Bishop, DamagedBishop1, DamagedBishop2:
						g.moveBishop(x, y)
					case Queen, DamagedQueen1, DamagedQueen2, DamagedQueen3, DamagedQueen4:
						g.moveQueen(x, y)
					case King:
						g.moveKing(x, y)
					}
					if g.hp == 0 {
						g.addEvent(EventEndGame)
						return
					}
				}
			}
		}

		if g.steps%AddEnemyEveryNSteps == 0 {
			g.insertEnemy()
		}

		// If game field is not empty - return
		for _, row := range g.enemiesXYNotMoved {
			if len(row) > 0 {
				return
			}
		}

		g.insertEnemy()
	}
}

func (g *Game) MovePlayerTo(direction int8) bool {
	if g.TurnOfPlayer && g.hp > Zero {
		g.switchTurnParameters()
		switch direction {
		case North:
			if g.y > Zero {
				g.movePlayerToNorth()
			} else {
				g.switchTurnParameters()
				return false
			}
		case East:
			if g.x < LastIndex {
				g.movePlayerToEast()
			} else {
				g.switchTurnParameters()
				return false
			}
		case South:
			if g.y < LastIndex {
				g.movePlayerToSouth()
			} else {
				g.switchTurnParameters()
				return false
			}
		case West:
			if g.x > Zero {
				g.movePlayerToWest()
			} else {
				g.switchTurnParameters()
				return false
			}
		case NoDirection:
			g.movePlayerToNoDirection()
		}

		if g.yNorth < Zero {
			g.yNorth = Zero
		}
		if g.xEast > LastIndex {
			g.xEast = LastIndex
		}
		if g.ySouth > LastIndex {
			g.ySouth = LastIndex
		}
		if g.xWest < Zero {
			g.xWest = Zero
		}

		g.steps++
		return true
	}
	return false
}
