package game

import (
	"bytes"
	"encoding/gob"
)

type SGame struct {
	HasHair             bool
	TurnOfPlayer        bool
	TurnEvents          []int8
	NextTurnEvent       int
	Field               [FieldSide][FieldSide]int8
	EnemiesXYNotMoved   map[int8]map[int8]bool
	HP                  int8
	DeltaHP             int8
	Damaged             bool
	AttackFromNorth     bool
	AttackFromNorthEast bool
	AttackFromEast      bool
	AttackFromSouthEast bool
	AttackFromSouth     bool
	AttackFromSouthWest bool
	AttackFromWest      bool
	AttackFromNorthWest bool
	Steps               uint64
	Score               uint64
	DeltaScore          uint64
	Direction           int8
	DirSteps            int8
	X                   int8
	Y                   int8
	YNorth              int8
	XEast               int8
	YSouth              int8
	XWest               int8
	NextEnemyDir        int8
	NextEnemyPos        int8
}

func (g *Game) ToSerial() []byte {
	sgame := SGame{
		HasHair:             g.HasHair,
		TurnOfPlayer:        g.TurnOfPlayer,
		TurnEvents:          g.TurnEvents,
		NextTurnEvent:       g.nextTurnEvent,
		Field:               *g.field,
		EnemiesXYNotMoved:   g.enemiesXYNotMoved,
		HP:                  g.hp,
		DeltaHP:             g.deltaHP,
		Damaged:             g.damaged,
		AttackFromNorth:     g.attackFromNorth,
		AttackFromNorthEast: g.attackFromNorthEast,
		AttackFromEast:      g.attackFromEast,
		AttackFromSouthEast: g.attackFromSouthEast,
		AttackFromSouth:     g.attackFromSouth,
		AttackFromSouthWest: g.attackFromSouthWest,
		AttackFromWest:      g.attackFromWest,
		AttackFromNorthWest: g.attackFromNorthWest,
		Steps:               g.steps,
		Score:               g.score,
		DeltaScore:          g.deltaScore,
		Direction:           g.direction,
		DirSteps:            g.dirSteps,
		X:                   g.x,
		Y:                   g.y,
		YNorth:              g.yNorth,
		XEast:               g.xEast,
		YSouth:              g.ySouth,
		XWest:               g.xWest,
		NextEnemyDir:        g.nextEnemyDir,
		NextEnemyPos:        g.nextEnemyPos,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(sgame)
	return buf.Bytes()
}

func (g *Game) RestoreFrom(b []byte) {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	var sg = new(SGame)
	err := dec.Decode(sg)
	if err != nil {
		return
	}

	g.HasHair = sg.HasHair
	g.TurnOfPlayer = sg.TurnOfPlayer
	g.TurnEvents = sg.TurnEvents
	g.nextTurnEvent = sg.NextTurnEvent
	g.field = &sg.Field
	g.enemiesXYNotMoved = sg.EnemiesXYNotMoved
	g.hp = sg.HP
	g.deltaHP = sg.DeltaHP
	g.damaged = sg.Damaged
	g.attackFromNorth = sg.AttackFromNorth
	g.attackFromNorthEast = sg.AttackFromNorthEast
	g.attackFromEast = sg.AttackFromEast
	g.attackFromSouthEast = sg.AttackFromSouthEast
	g.attackFromSouth = sg.AttackFromSouth
	g.attackFromSouthWest = sg.AttackFromSouthWest
	g.attackFromWest = sg.AttackFromWest
	g.attackFromNorthWest = sg.AttackFromNorthWest
	g.steps = sg.Steps
	g.score = sg.Score
	g.deltaScore = sg.DeltaScore
	g.direction = sg.Direction
	g.dirSteps = sg.DirSteps
	g.x = sg.X
	g.y = sg.Y
	g.yNorth = sg.YNorth
	g.xEast = sg.XEast
	g.ySouth = sg.YSouth
	g.xWest = sg.XWest
	g.nextEnemyDir = sg.NextEnemyDir
	g.nextEnemyPos = sg.NextEnemyPos
}
