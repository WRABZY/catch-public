package game

import (
	"math/rand/v2"
)

func (g *Game) addEvent(event int8) {
	if g.nextTurnEvent < OneTurnMaxEvents {
		g.TurnEvents[g.nextTurnEvent] = event
		g.nextTurnEvent++
	}
}

func (g *Game) insertEnemy() {
	if g.hp > Zero {
		who := rand.IntN(100)
		if who == 0 { // 1 %
			g.insert(King)
			g.addEvent(EventKingInserted)
		} else if who < 60 { // 59 %
			g.insert(PawnNorth)
			g.addEvent(EventPawnInserted)
		} else if who < 90 { // 30 %
			g.insert(Rook)
			g.addEvent(EventRookInserted)
		} else { //	10 %
			g.insert(Bishop)
			g.addEvent(EventBishopInserted)
		}
	}
}

func (g *Game) insert(enemy int8) {
	var x int8
	var y int8
	var enemyToInsert = enemy

	for i := Zero; i < FieldPerimeter; i++ {
		switch g.nextEnemyDir {
		case North:
			x = g.nextEnemyPos
			y = Zero
			if enemy < Rook {
				enemyToInsert = PawnNorth // enemy < Rook is Pawn, there are 4 types of pawns depending on the direction of insertion
			}
			g.nextEnemyDir = East
		case East:
			y = g.nextEnemyPos
			x = LastIndex
			if enemy < Rook {
				enemyToInsert = PawnEast
			}
			g.nextEnemyDir = South
		case South:
			x = g.nextEnemyPos
			y = LastIndex
			if enemy < Rook {
				enemyToInsert = PawnSouth
			}
			g.nextEnemyDir = West
		case West:
			y = g.nextEnemyPos
			x = Zero
			if enemy < Rook {
				enemyToInsert = PawnWest
			}
			g.nextEnemyDir = North
		}

		// Position rotation
		if g.nextEnemyPos == LastIndex {
			g.nextEnemyPos = Zero
		} else {
			g.nextEnemyPos++
		}

		// Try to insert
		if g.cellIsFree(x, y) {
			g.field[y][x] = enemyToInsert
			g.enemiesXYNotMoved[x][y] = false
			return
		}
	}
}

func (g *Game) switchTurnParameters() {
	g.TurnOfPlayer = !g.TurnOfPlayer
	g.deltaScore = Zero
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
	g.nextTurnEvent = 0
	for i := range g.TurnEvents {
		g.TurnEvents[i] = EventEmpty
	}
}

func (g *Game) cleanShadows() {
	for y, arr := range g.field {
		for x, v := range arr {
			if v > King && (v < ShadowPlayerNorth || v > ShadowPlayerWest) { // Cleaning only enemy shadows
				g.field[y][x] = Nobody
			}
		}
	}
}

func (g *Game) cleanMoves() {
	for x, m := range g.enemiesXYNotMoved {
		for y, v := range m {
			if !v {
				g.enemiesXYNotMoved[x][y] = true
			}
		}
	}
}

func (g *Game) movePlayerToNoDirection() {
	g.dirSteps = 1
	g.direction = NoDirection
	g.yNorth = g.y - 1
	g.xEast = g.x + 1
	g.ySouth = g.y + 1
	g.xWest = g.x - 1
	g.addEvent(EventSkipTurn)
}

func (g *Game) movePlayerToNorth() {
	// Move
	g.field[g.y][g.x] = ShadowPlayerNorth
	g.y = g.yNorth
	g.eatEnemy()
	g.field[g.y][g.x] = Player

	// Calc next steps
	if g.direction != North {
		g.dirSteps = 1
		g.direction = North
	} else {
		g.dirSteps++
	}

	var newYNorth int8
	if g.dirSteps == 1 {
		newYNorth = g.y - WalkStep
	} else {
		newYNorth = g.y - RunStep
	}

	g.yNorth = newYNorth
	g.xEast = g.x + 1
	g.ySouth = g.y + 1
	g.xWest = g.x - 1
}

func (g *Game) movePlayerToEast() {
	// Move
	g.field[g.y][g.x] = ShadowPlayerEast
	g.x = g.xEast
	g.eatEnemy()
	g.field[g.y][g.x] = Player

	// Calc next steps
	if g.direction != East {
		g.dirSteps = 1
		g.direction = East
	} else {
		g.dirSteps++
	}

	var newXEast int8
	if g.dirSteps == 1 {
		newXEast = g.x + WalkStep
	} else {
		newXEast = g.x + RunStep
	}

	g.yNorth = g.y - 1
	g.xEast = newXEast
	g.ySouth = g.y + 1
	g.xWest = g.x - 1
}

func (g *Game) movePlayerToSouth() {
	// Move
	g.field[g.y][g.x] = ShadowPlayerSouth
	g.y = g.ySouth
	g.eatEnemy()
	g.field[g.y][g.x] = Player

	// Calc next steps
	if g.direction != South {
		g.dirSteps = 1
		g.direction = South
	} else {
		g.dirSteps++
	}

	var newYSouth int8
	if g.dirSteps == 1 {
		newYSouth = g.y + WalkStep
	} else {
		newYSouth = g.y + RunStep
	}

	g.yNorth = g.y - 1
	g.xEast = g.x + 1
	g.ySouth = newYSouth
	g.xWest = g.x - 1
}

func (g *Game) movePlayerToWest() {
	// Move
	g.field[g.y][g.x] = ShadowPlayerWest
	g.x = g.xWest
	g.eatEnemy()
	g.field[g.y][g.x] = Player

	// Calc next steps
	if g.direction != West {
		g.dirSteps = 1
		g.direction = West
	} else {
		g.dirSteps++
	}

	var newXWest int8
	if g.dirSteps == 1 {
		newXWest = g.x - WalkStep
	} else {
		newXWest = g.x - RunStep
	}

	g.yNorth = g.y - 1
	g.xEast = g.x + 1
	g.ySouth = g.y + 1
	g.xWest = newXWest
}

func (g *Game) eatEnemy() {
	switch g.field[g.y][g.x] {
	case King:
		g.deltaScore = KingScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventKingEaten)
	case Queen, DamagedQueen1, DamagedQueen2, DamagedQueen3, DamagedQueen4:
		if rand.IntN(100) < 10 {
			g.insert(King)
			g.addEvent(EventKingInserted)
		}
		g.deltaScore = QueenScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventQueenEaten)
	case Knight, DamagedKnight1, DamagedKnight2, DamagedKnight3:
		g.deltaScore = KnightScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
	case Bishop:
		g.deltaHP = 3
		g.deltaScore = BishopScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventBishopEaten)
		g.addEvent(EventBishopHeals)
	case DamagedBishop1:
		g.deltaHP = 2
		g.deltaScore = BishopScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventBishopEaten)
		g.addEvent(EventBishopHeals)
	case DamagedBishop2:
		g.deltaHP = 1
		g.deltaScore = BishopScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventBishopEaten)
		g.addEvent(EventBishopHeals)
	case Rook, DamagedRook:
		g.deltaScore = RookScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventRookEaten)
	case PawnNorth, PawnEast, PawnSouth, PawnWest:
		g.deltaScore = PawnScore
		delete(g.enemiesXYNotMoved[g.x], g.y)
		g.addEvent(EventPawnEaten)
	case ShadowPawn:
		g.deltaHP = -1
		g.damaged = true
		if g.hp > 1 {
			g.addEvent(EventVileStenchAttacks)
		} else {
			g.addEvent(EventVileStenchAttacksLast)
		}
	case LastShadowRook:
		g.deltaHP = 1
		g.deltaScore = LastShadowScore
		g.addEvent(EventRookHeals)
	// TODO
	/*case LastShadowKnight:
	g.deltaHP = 1
	g.deltaScore = LastShadowScore*/
	case LastShadowQueen:
		g.deltaHP = 2
		g.deltaScore = LastShadowQueenScore
		g.addEvent(EventQueenHeals)
	case ShadowKingRD, ShadowKingLD:
		g.deltaScore = CoinScore
		g.addEvent(EventCoinFound)
	}

	g.score += g.deltaScore
	g.hp += g.deltaHP

	if g.hp > HPMax {
		g.hp = HPMax
	} else if g.hp < Zero {
		g.hp = Zero
	}
}

// North pawn moves to south
func (g *Game) movePawnNorthFrom(x, y int8) {
	nextY := y + 1

	// Attack, if you can
	if g.damagePlayerFromTo(x, y, x-1, nextY) || g.damagePlayerFromTo(x, y, x+1, nextY) {
		if g.hp > 0 {
			g.field[y][x] = LastShadowPawnNorth // Pawn dies after attack
			delete(g.enemiesXYNotMoved[x], y)
			g.addEvent(EventPawnAttacks)
		} else {
			g.addEvent(EventPawnAttacksLast)
		}
		return
	}

	// Move, if you can
	if g.cellIsFree(x, nextY) {
		if nextY == LastIndex { // Pawn becomes Queen
			g.field[nextY][x] = Queen
			g.addEvent(EventFromPawnToQueen)
		} else {
			g.field[nextY][x] = PawnNorth
		}
		g.field[y][x] = ShadowPawn
		delete(g.enemiesXYNotMoved[x], y)
		g.enemiesXYNotMoved[x][nextY] = false // already moved
		return
	}
}

// East pawn moves to west
func (g *Game) movePawnEastFrom(x, y int8) {
	nextX := x - 1

	// Attack, if you can
	if g.damagePlayerFromTo(x, y, nextX, y-1) || g.damagePlayerFromTo(x, y, nextX, y+1) {
		if g.hp > 0 {
			g.field[y][x] = LastShadowPawnEast // Pawn dies after attack
			delete(g.enemiesXYNotMoved[x], y)
			g.addEvent(EventPawnAttacks)
		} else {
			g.addEvent(EventPawnAttacksLast)
		}
		return
	}

	// Move, if you can
	if g.cellIsFree(nextX, y) {
		if nextX == Zero { // Pawn becomes Queen
			g.field[y][nextX] = Queen
			g.addEvent(EventFromPawnToQueen)
		} else {
			g.field[y][nextX] = PawnEast
		}
		g.field[y][x] = ShadowPawn
		delete(g.enemiesXYNotMoved[x], y)
		g.enemiesXYNotMoved[nextX][y] = false // already moved
		return
	}
}

// South pawn moves to north
func (g *Game) movePawnSouthFrom(x, y int8) {
	nextY := y - 1

	// Attack, if you can
	if g.damagePlayerFromTo(x, y, x-1, nextY) || g.damagePlayerFromTo(x, y, x+1, nextY) {
		if g.hp > 0 {
			g.field[y][x] = LastShadowPawnSouth // Pawn dies after attack
			delete(g.enemiesXYNotMoved[x], y)
			g.addEvent(EventPawnAttacks)
		} else {
			g.addEvent(EventPawnAttacksLast)
		}
		return
	}

	// Move, if you can
	if g.cellIsFree(x, nextY) {
		if nextY == Zero { // Pawn becomes Queen
			g.field[nextY][x] = Queen
			g.addEvent(EventFromPawnToQueen)
		} else {
			g.field[nextY][x] = PawnSouth
		}
		g.field[y][x] = ShadowPawn
		delete(g.enemiesXYNotMoved[x], y)
		g.enemiesXYNotMoved[x][nextY] = false // already moved
		return
	}
}

// West pawn moves to east
func (g *Game) movePawnWestFrom(x, y int8) {
	nextX := x + 1

	// Attack, if you can
	if g.damagePlayerFromTo(x, y, nextX, y-1) || g.damagePlayerFromTo(x, y, nextX, y+1) {
		if g.hp > 0 {
			g.field[y][x] = LastShadowPawnWest // Pawn dies after attack
			delete(g.enemiesXYNotMoved[x], y)
			g.addEvent(EventPawnAttacks)
		} else {
			g.addEvent(EventPawnAttacksLast)
		}
		return
	}

	// Move, if you can
	if g.cellIsFree(nextX, y) {
		if nextX == LastIndex { // Pawn becomes Queen
			g.field[y][nextX] = Queen
			g.addEvent(EventFromPawnToQueen)
		} else {
			g.field[y][nextX] = PawnWest
		}
		g.field[y][x] = ShadowPawn
		delete(g.enemiesXYNotMoved[x], y)
		g.enemiesXYNotMoved[nextX][y] = false // already moved
		return
	}
}

func (g *Game) moveRook(x, y int8) {
// not public
}

func (g *Game) moveBishop(x, y int8) {
// not public
}

func (g *Game) moveQueen(x, y int8) {
// not public
}

func (g *Game) moveKing(x, y int8) {
// not public
}

func (g *Game) rookAttackTransform(xFrom, yFrom, xTo, yTo int8) {
	if g.hp > 0 {
		delete(g.enemiesXYNotMoved[xFrom], yFrom)
		switch g.field[yFrom][xFrom] {
		case Rook:

			g.field[yTo][xTo] = DamagedRook       // Rook becomes DamagedRook after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventRookAttacks)
		case DamagedRook:
			g.field[yTo][xTo] = LastShadowRook // DamagedRook dies after attack
			g.addEvent(EventDamagedRookAttacks)
		}
	} else {
		g.field[yTo][xTo] = g.field[yFrom][xFrom]
		g.addEvent(EventRookAttacksLast)
	}
}

func (g *Game) bishopAttackTransform(xFrom, yFrom, xTo, yTo int8) {
	if g.hp > 0 {
		delete(g.enemiesXYNotMoved[xFrom], yFrom)
		switch g.field[yFrom][xFrom] {
		case Bishop:
			g.field[yTo][xTo] = DamagedBishop1    // Bishop becomes DamagedBishop1 after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventBishopAttacks)
		case DamagedBishop1:
			g.field[yTo][xTo] = DamagedBishop2    // DamagedBishop1 becomes DamagedBishop2 after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventBishopAttacks)
		case DamagedBishop2:
			g.field[yTo][xTo] = LastShadowBishop // DamagedBishop2 dies after attack
			g.addEvent(EventDamagedBishopAttacks)
		}
	} else {
		g.field[yTo][xTo] = g.field[yFrom][xFrom]
		g.addEvent(EventBishopAttacksLast)
	}
}

func (g *Game) queenAttackTransform(xFrom, yFrom, xTo, yTo int8) {
	if g.hp > 0 {
		delete(g.enemiesXYNotMoved[xFrom], yFrom)
		switch g.field[yFrom][xFrom] {
		case Queen:
			g.field[yTo][xTo] = DamagedQueen1     // Queen becomes DamagedQueen1 after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventQueenAttacks)
		case DamagedQueen1:
			g.field[yTo][xTo] = DamagedQueen2     // DamagedQueen1 becomes DamagedQueen2 after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventQueenAttacks)
		case DamagedQueen2:
			g.field[yTo][xTo] = DamagedQueen3     // DamagedQueen2 becomes DamagedQueen3 after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventQueenAttacks)
		case DamagedQueen3:
			g.field[yTo][xTo] = DamagedQueen4     // DamagedQueen3 becomes DamagedQueen4 after attack
			g.enemiesXYNotMoved[xTo][yTo] = false // already moved
			g.addEvent(EventQueenAttacks)
		case DamagedQueen4:
			g.field[yTo][xTo] = LastShadowQueen // DamagedQueen4 dies after attack
			g.addEvent(EventDamagedQueenAttacks)
		}
	} else {
		g.field[yTo][xTo] = g.field[yFrom][xFrom]
		g.addEvent(EventQueenAttacksLast)
	}
}

func (g *Game) moveKingFromTo(xFrom, yFrom, xTo, yTo int8) bool {
	if !g.cellIsFree(xTo, yTo) {
		return false
	}
	delete(g.enemiesXYNotMoved[xFrom], yFrom)
	g.enemiesXYNotMoved[xTo][yTo] = false // already moved
	g.field[yTo][xTo] = King
	if xFrom == xTo {
		g.field[yFrom][xFrom] = ShadowKingV
	} else if yFrom == yTo {
		g.field[yFrom][xFrom] = ShadowKingH
	} else if (xFrom < xTo) == (yFrom < yTo) {
		g.field[yFrom][xFrom] = ShadowKingRD
	} else {
		g.field[yFrom][xFrom] = ShadowKingLD
	}

	return true
}

func (g *Game) damagePlayerFromTo(xFrom, yFrom, xTo, yTo int8) bool {
	if !g.cellIsPlayer(xTo, yTo) {
		return false
	}
	g.deltaHP--
	g.hp--
	if g.hp < Zero {
		g.hp = Zero
	}
	g.damaged = true
	if xFrom == xTo {
		if yFrom < yTo {
			g.attackFromNorth = true
		} else {
			g.attackFromSouth = true
		}
	} else if yFrom == yTo {
		if xFrom < xTo {
			g.attackFromWest = true
		} else {
			g.attackFromEast = true
		}
	} else if xFrom < xTo && yFrom < yTo {
		g.attackFromNorthWest = true
	} else if xFrom < xTo && yFrom > yTo {
		g.attackFromSouthWest = true
	} else if xFrom > xTo && yFrom < yTo {
		g.attackFromNorthEast = true
	} else {
		g.attackFromSouthEast = true
	}
	return true
}

func (g *Game) cellIsDangerous(x, y int8) bool {
	return ((x == g.xEast || x == g.xWest) && y == g.y) || (x == g.x && (y == g.yNorth || y == g.ySouth))
}

func (g *Game) cellIsVeryDangerousForPlayer(x, y int8) bool {
	if !g.cellExists(x, y) {
		return false
	}
	dif := (g.x - x) * (g.y - y)
	return dif == 1 || dif == -1
}

func (g *Game) cellIsFree(x, y int8) bool {
	return g.cellExists(x, y) && (g.field[y][x] == Nobody || g.field[y][x] > LastShadowQueen)
}

func (g *Game) cellIsPlayer(x, y int8) bool {
	return g.cellExists(x, y) && (g.field[y][x] == Player)
}

func (g *Game) cellExists(x, y int8) bool {
	return x > TooSmall && x < TooBig && y > TooSmall && y < TooBig
}
