package game

// Entity codes
const (
	Nobody = iota
	Player
	PawnNorth
	PawnEast
	PawnSouth
	PawnWest
	DamagedRook
	Rook
	DamagedBishop2
	DamagedBishop1
	Bishop
	DamagedKnight3
	DamagedKnight2
	DamagedKnight1
	Knight
	DamagedQueen4
	DamagedQueen3
	DamagedQueen2
	DamagedQueen1
	Queen
	King
	LastShadowPawnNorth
	LastShadowPawnEast
	LastShadowPawnSouth
	LastShadowPawnWest
	LastShadowRook
	LastShadowBishop
	LastShadowKnight
	LastShadowQueen
	ShadowPlayerNorth
	ShadowPlayerEast
	ShadowPlayerSouth
	ShadowPlayerWest
	ShadowPawn
	ShadowRookV
	ShadowRookH
	ShadowBishopL
	ShadowBishopR
	ShadowKnight
	ShadowQueenV
	ShadowQueenH
	ShadowQueenRD
	ShadowQueenLD
	ShadowKingV
	ShadowKingH
	ShadowKingRD
	ShadowKingLD
)

const Zero = 0

// Game field
const (
	FieldSide      = 9
	FieldSize      = FieldSide * FieldSide
	FieldPerimeter = 4 * (FieldSide - 1)
	TooSmall       = -1
	TooBig         = FieldSide
	Center         = FieldSide / 2
	FirstIndex     = Zero
	LastIndex      = FieldSide - 1
)

// Directions
const (
	NoDirection = iota
	North
	East
	South
	West
	NorthEast
	SouthEast
	SouthWest
	NorthWest
	DirectionsNumber     = 8
	DirectionsNumberHalf = DirectionsNumber / 2
)

const HPMax = 9
const AddEnemyEveryNSteps = 5

// Score
const (
	LastShadowScore      = 1
	LastShadowQueenScore = 5
	PawnScore            = 5
	CoinScore            = 10
	RookScore            = 10
	BishopScore          = 30
	KnightScore          = 50
	QueenScore           = 100
	KingScore            = 1000
)

const (
	WalkStep = 1
	RunStep  = 2
)

const (
	GraphicCellPixels     = 70
	GraphicHalfCellPixels = GraphicCellPixels / 2
	GraphicSideCells      = 10
	GraphicSidePixels     = GraphicCellPixels * GraphicSideCells
	GraphicSideHalfPixels = GraphicSidePixels / 2
)

// Events
const OneTurnMaxEvents = 50
const (
	EventEmpty = iota
	EventNewGame
	EventEndGame

	EventFromPawnToQueen

	EventPawnAttacks
	EventPawnAttacksLast
	EventRookAttacks
	EventDamagedRookAttacks
	EventRookAttacksLast
	EventBishopAttacks
	EventDamagedBishopAttacks
	EventBishopAttacksLast
	EventQueenAttacks
	EventDamagedQueenAttacks
	EventQueenAttacksLast
	EventKingAttacks
	EventKingAttacksLast

	EventPawnEaten
	EventRookEaten
	EventBishopEaten
	EventQueenEaten
	EventKingEaten

	EventVileStenchAttacks
	EventVileStenchAttacksLast

	EventBishopHeals
	EventRookHeals
	EventQueenHeals

	EventSkipTurn

	EventPawnInserted
	EventRookInserted
	EventBishopInserted
	EventKingInserted

	EventCoinFound
)
