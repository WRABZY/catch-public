package game

import (
	"image"
	"image/draw"
	"strconv"
)

var (
	base                                  *image.NRGBA
	baseWSOD                              *image.NRGBA
	startCoordinates                      map[int]int
	endCoordinates                        map[int]int
	hpStartCoordinates                    map[int]int
	hpEndCoordinates                      map[int]int
	endTopBorderCoordinatesY              map[int]int
	startBotBorderCoordinatesY            map[int]int
	startRightBorderCoordinatesX          map[int]int
	startLeftRightBorderCoordinatesY      map[int]int
	endLeftRightBorderCoordinatesY        map[int]int
	endLeftBorderCoordinatesX             map[int]int
	startScoreCoordinates                 map[int]map[int]int
	endScoreCoordinates                   map[int]map[int]int
	startAttackMarkerNorthWestCoordinates map[int]int
	endAttackMarkerNorthWestCoordinates   map[int]int
	startAttackMarkerNorthCoordinates     map[int]int
	endAttackMarkerNorthCoordinates       map[int]int
	startAttackMarkerNorthEastCoordinates map[int]int
	endAttackMarkerNorthEastCoordinates   map[int]int
)

func init() {
	base = image.NewNRGBA(image.Rect(0, 0, GraphicSidePixels, GraphicSidePixels))
	baseWSOD = image.NewNRGBA(image.Rect(0, 0, GraphicSidePixels, GraphicSidePixels))

	// Headers

	var xLeftTop int = 0
	var yLeftTop int = 0
	var xRightBot int = 0
	var yRightBot int = GraphicCellPixels
	for i := 0; i < GraphicSideCells; i++ {
		xLeftTop = i * GraphicCellPixels
		xRightBot = xLeftTop + GraphicCellPixels
		draw.Draw(base, image.Rect(xLeftTop, yLeftTop, xRightBot, yRightBot), assetBaseHeader, image.Point{0, 0}, draw.Src)
		draw.Draw(baseWSOD, image.Rect(xLeftTop, yLeftTop, xRightBot, yRightBot), assetBaseHeaderWSOD, image.Point{0, 0}, draw.Src)
	}

	// HP
	xLeftTop = 0
	xRightBot = GraphicCellPixels
	for i := 1; i < GraphicSideCells; i++ {
		yLeftTop = i * GraphicCellPixels
		yRightBot = yLeftTop + GraphicCellPixels
		draw.Draw(base, image.Rect(xLeftTop, yLeftTop, xRightBot, yRightBot), assetBaseNoHP, image.Point{0, 0}, draw.Src)
		draw.Draw(baseWSOD, image.Rect(xLeftTop, yLeftTop, xRightBot, yRightBot), assetBaseNoHP, image.Point{0, 0}, draw.Src)
	}

	// Field cells
	var imageToDraw *image.NRGBA
	for i := 1; i < GraphicSideCells; i++ {
		for j := 1; j < GraphicSideCells; j++ {
			xLeftTop = i * GraphicCellPixels
			yLeftTop = j * GraphicCellPixels
			xRightBot = xLeftTop + GraphicCellPixels
			yRightBot = yLeftTop + GraphicCellPixels

			if i%2 == j%2 {
				imageToDraw = assetBaseWhite
			} else {
				imageToDraw = assetBaseBlack
			}
			draw.Draw(base, image.Rect(xLeftTop, yLeftTop, xRightBot, yRightBot), imageToDraw, image.Point{0, 0}, draw.Src)
			draw.Draw(baseWSOD, image.Rect(xLeftTop, yLeftTop, xRightBot, yRightBot), assetBaseWSOD, image.Point{0, 0}, draw.Src)
		}
	}

	startCoordinates = make(map[int]int)
	endCoordinates = make(map[int]int)
	endTopBorderCoordinatesY = make(map[int]int)
	startBotBorderCoordinatesY = make(map[int]int)
	startRightBorderCoordinatesX = make(map[int]int)
	startLeftRightBorderCoordinatesY = make(map[int]int)
	endLeftRightBorderCoordinatesY = make(map[int]int)
	endLeftBorderCoordinatesX = make(map[int]int)
	startAttackMarkerNorthWestCoordinates = make(map[int]int)
	endAttackMarkerNorthWestCoordinates = make(map[int]int)
	startAttackMarkerNorthCoordinates = make(map[int]int)
	endAttackMarkerNorthCoordinates = make(map[int]int)
	startAttackMarkerNorthEastCoordinates = make(map[int]int)
	endAttackMarkerNorthEastCoordinates = make(map[int]int)
	for x := 0; x < FieldSide; x++ {
		startCoordinates[x] = GraphicCellPixels + x*GraphicCellPixels
		endCoordinates[x] = startCoordinates[x] + GraphicCellPixels
		endTopBorderCoordinatesY[x] = startCoordinates[x] + assetTargetTopH.Bounds().Dy()
		startBotBorderCoordinatesY[x] = endCoordinates[x] - assetTargetBotH.Bounds().Dy()
		startRightBorderCoordinatesX[x] = endCoordinates[x] - assetTargetRightV.Bounds().Dx()
		startLeftRightBorderCoordinatesY[x] = startCoordinates[x] + assetTargetTopH.Bounds().Dy()
		endLeftRightBorderCoordinatesY[x] = endCoordinates[x] - assetTargetBotH.Bounds().Dy()
		endLeftBorderCoordinatesX[x] = startCoordinates[x] + assetTargetLeftV.Bounds().Dx()
		startAttackMarkerNorthWestCoordinates[x] = startCoordinates[x] - assetAttackMarker.Bounds().Dx()/2
		endAttackMarkerNorthWestCoordinates[x] = startAttackMarkerNorthWestCoordinates[x] + assetAttackMarker.Bounds().Dx()
		startAttackMarkerNorthCoordinates[x] = startCoordinates[x] + (GraphicCellPixels-assetAttackMarker.Bounds().Dx())/2
		endAttackMarkerNorthCoordinates[x] = startAttackMarkerNorthCoordinates[x] + assetAttackMarker.Bounds().Dx()
		startAttackMarkerNorthEastCoordinates[x] = endCoordinates[x] - assetAttackMarker.Bounds().Dx()/2
		endAttackMarkerNorthEastCoordinates[x] = startAttackMarkerNorthEastCoordinates[x] + assetAttackMarker.Bounds().Dx()
	}

	hpStartCoordinates = make(map[int]int)
	hpEndCoordinates = make(map[int]int)
	for y := 1; y <= HPMax; y++ {
		hpStartCoordinates[y] = GraphicSidePixels - y*GraphicCellPixels
		hpEndCoordinates[y] = hpStartCoordinates[y] + GraphicCellPixels
	}

	startScoreCoordinates = make(map[int]map[int]int, GraphicSideCells*2)
	endScoreCoordinates = make(map[int]map[int]int, GraphicSideCells*2)
	// key of outer map is score len (number of digits)
	// key of inner map is number of digit
	for length := 1; length <= GraphicSideCells*2; length++ {
		startScoreCoordinates[length] = make(map[int]int, length)
		endScoreCoordinates[length] = make(map[int]int, length)
		for i := 0; i < length; i++ {
			startScoreCoordinates[length][i] = GraphicSideHalfPixels - length/2*GraphicHalfCellPixels + i*GraphicHalfCellPixels
			endScoreCoordinates[length][i] = startScoreCoordinates[length][i] + GraphicHalfCellPixels
		}
	}
}

func (g *Game) RefreshFrame() {
	if g.hp > 0 {
		draw.Draw(g.Frame, image.Rect(0, 0, GraphicSidePixels, GraphicSidePixels), base, image.Point{0, 0}, draw.Src)
		g.drawEnemies()
		g.drawCharacter()
		g.drawScore()
		if g.TurnOfPlayer {
			g.drawTargets()
		}
	} else {
		draw.Draw(g.Frame, image.Rect(0, 0, GraphicSidePixels, GraphicSidePixels), baseWSOD, image.Point{0, 0}, draw.Src)
		g.drawWSODEnemies()
		g.drawWSODCharacter()
		g.drawWSODScore()
	}
	g.drawAttackMarkers()
}

func (g *Game) drawCharacter() {
	// Draw HP
	for y := 1; y <= int(g.hp); y++ {
		draw.Draw(g.Frame, image.Rect(0, hpStartCoordinates[y], GraphicCellPixels, hpEndCoordinates[y]), assetBaseHP, image.Point{0, 0}, draw.Src)
	}

	var imageToDraw *image.NRGBA
	if g.damaged {
		if g.HasHair {
			imageToDraw = assetCharSadHair
		} else {
			imageToDraw = assetCharSad
		}
	} else if g.deltaScore > 0 {
		if g.HasHair {
			imageToDraw = assetCharHappyHair
		} else {
			imageToDraw = assetCharHappy
		}
	} else if g.direction == NoDirection {
		if g.HasHair {
			imageToDraw = assetCharSleepingHair
		} else {
			imageToDraw = assetCharSleeping
		}
	} else {
		if g.HasHair {
			imageToDraw = assetCharHair
		} else {
			imageToDraw = assetChar
		}
	}
	draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startCoordinates[int(g.y)], endCoordinates[int(g.x)], endCoordinates[int(g.y)]), imageToDraw, image.Point{0, 0}, draw.Over)
}

func (g *Game) drawEnemies() {
	var imageToDraw *image.NRGBA
	for y, arr := range *g.field {
		for x, val := range arr {
			switch val {
			case PawnNorth:
				imageToDraw = assetPawnNorth
			case PawnEast:
				imageToDraw = assetPawnEast
			case PawnSouth:
				imageToDraw = assetPawnSouth
			case PawnWest:
				imageToDraw = assetPawnWest
			case DamagedRook:
				imageToDraw = assetDamagedRook
			case Rook:
				imageToDraw = assetRook
			case DamagedBishop2:
				imageToDraw = assetDamagedBishop2
			case DamagedBishop1:
				imageToDraw = assetDamagedBishop1
			case Bishop:
				imageToDraw = assetBishop
			case DamagedQueen4:
				imageToDraw = assetDamagedQueen4
			case DamagedQueen3:
				imageToDraw = assetDamagedQueen3
			case DamagedQueen2:
				imageToDraw = assetDamagedQueen2
			case DamagedQueen1:
				imageToDraw = assetDamagedQueen1
			case Queen:
				imageToDraw = assetQueen
			case King:
				imageToDraw = assetKing
			case ShadowPlayerNorth:
				if y%2 == x%2 {
					imageToDraw = assetShadowPlayerNorthWhite
				} else {
					imageToDraw = assetShadowPlayerNorthBlack
				}
			case ShadowPlayerEast:
				if y%2 == x%2 {
					imageToDraw = assetShadowPlayerEastWhite
				} else {
					imageToDraw = assetShadowPlayerEastBlack
				}
			case ShadowPlayerSouth:
				if y%2 == x%2 {
					imageToDraw = assetShadowPlayerSouthWhite
				} else {
					imageToDraw = assetShadowPlayerSouthBlack
				}
			case ShadowPlayerWest:
				if y%2 == x%2 {
					imageToDraw = assetShadowPlayerWestWhite
				} else {
					imageToDraw = assetShadowPlayerWestBlack
				}
			case ShadowPawn:
				if y%2 == x%2 {
					imageToDraw = assetShadowPawnWhite
				} else {
					imageToDraw = assetShadowPawnBlack
				}
			case LastShadowPawnNorth:
				imageToDraw = assetLastShadowPawnNorth
			case LastShadowPawnEast:
				imageToDraw = assetLastShadowPawnEast
			case LastShadowPawnSouth:
				imageToDraw = assetLastShadowPawnSouth
			case LastShadowPawnWest:
				imageToDraw = assetLastShadowPawnWest
			case ShadowRookV:
				if y%2 == x%2 {
					imageToDraw = assetShadowRookVWhite
				} else {
					imageToDraw = assetShadowRookVBlack
				}
			case ShadowRookH:
				if y%2 == x%2 {
					imageToDraw = assetShadowRookHWhite
				} else {
					imageToDraw = assetShadowRookHBlack
				}
			case LastShadowRook:
				imageToDraw = assetLastShadowRook
			case ShadowBishopL:
				if y%2 == x%2 {
					imageToDraw = assetShadowBishopLWhite
				} else {
					imageToDraw = assetShadowBishopLBlack
				}
			case ShadowBishopR:
				if y%2 == x%2 {
					imageToDraw = assetShadowBishopRWhite
				} else {
					imageToDraw = assetShadowBishopRBlack
				}
			case LastShadowBishop:
				imageToDraw = assetLastShadowBishop
			case ShadowQueenV:
				if y%2 == x%2 {
					imageToDraw = assetShadowQueenVWhite
				} else {
					imageToDraw = assetShadowQueenVBlack
				}
			case ShadowQueenH:
				if y%2 == x%2 {
					imageToDraw = assetShadowQueenHWhite
				} else {
					imageToDraw = assetShadowQueenHBlack
				}
			case ShadowQueenRD:
				if y%2 == x%2 {
					imageToDraw = assetShadowQueenRDWhite
				} else {
					imageToDraw = assetShadowQueenRDBlack
				}
			case ShadowQueenLD:
				if y%2 == x%2 {
					imageToDraw = assetShadowQueenLDWhite
				} else {
					imageToDraw = assetShadowQueenLDBlack
				}
			case LastShadowQueen:
				imageToDraw = assetLastShadowQueen
			case ShadowKingV:
				if y%2 == x%2 {
					imageToDraw = assetShadowKingVWhite
				} else {
					imageToDraw = assetShadowKingVBlack
				}
			case ShadowKingH:
				if y%2 == x%2 {
					imageToDraw = assetShadowKingHWhite
				} else {
					imageToDraw = assetShadowKingHBlack
				}
			case ShadowKingRD:
				if y%2 == x%2 {
					imageToDraw = assetShadowKingRDWhite
				} else {
					imageToDraw = assetShadowKingRDBlack
				}
			case ShadowKingLD:
				if y%2 == x%2 {
					imageToDraw = assetShadowKingLDWhite
				} else {
					imageToDraw = assetShadowKingLDBlack
				}
			default:
				continue
			}
			draw.Draw(g.Frame, image.Rect(startCoordinates[x], startCoordinates[y], endCoordinates[x], endCoordinates[y]), imageToDraw, image.Point{0, 0}, draw.Src)
		}
	}
}

func (g *Game) drawScore() {
	if g.score > 0 {
		stringToPrint := strconv.FormatUint(g.score, 10)
		stringLen := len(stringToPrint)
		var imageToDraw *image.NRGBA
		for i := 0; i < stringLen; i++ {
			switch stringToPrint[i] {
			case '0':
				imageToDraw = assetScore0
			case '1':
				imageToDraw = assetScore1
			case '2':
				imageToDraw = assetScore2
			case '3':
				imageToDraw = assetScore3
			case '4':
				imageToDraw = assetScore4
			case '5':
				imageToDraw = assetScore5
			case '6':
				imageToDraw = assetScore6
			case '7':
				imageToDraw = assetScore7
			case '8':
				imageToDraw = assetScore8
			case '9':
				imageToDraw = assetScore9
			}
			draw.Draw(g.Frame, image.Rect(startScoreCoordinates[stringLen][i], Zero, endScoreCoordinates[stringLen][i], GraphicCellPixels), imageToDraw, image.Point{0, 0}, draw.Src)
		}
	}
}

func (g *Game) drawTargets() {
	// North
	if g.yNorth != g.y {
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startCoordinates[int(g.yNorth)], endCoordinates[int(g.x)], endTopBorderCoordinatesY[int(g.yNorth)]), assetTargetTopH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startBotBorderCoordinatesY[int(g.yNorth)], endCoordinates[int(g.x)], endCoordinates[int(g.yNorth)]), assetTargetTopH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startLeftRightBorderCoordinatesY[int(g.yNorth)], endLeftBorderCoordinatesX[int(g.x)], endLeftRightBorderCoordinatesY[int(g.yNorth)]), assetTargetTopV, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startRightBorderCoordinatesX[int(g.x)], startLeftRightBorderCoordinatesY[int(g.yNorth)], endCoordinates[int(g.x)], endLeftRightBorderCoordinatesY[int(g.yNorth)]), assetTargetTopV, image.Point{0, 0}, draw.Src)
	}
	// East
	if g.xEast != g.x {
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.xEast)], startCoordinates[int(g.y)], endCoordinates[int(g.xEast)], endTopBorderCoordinatesY[int(g.y)]), assetTargetRightH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.xEast)], startBotBorderCoordinatesY[int(g.y)], endCoordinates[int(g.xEast)], endCoordinates[int(g.y)]), assetTargetRightH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.xEast)], startLeftRightBorderCoordinatesY[int(g.y)], endLeftBorderCoordinatesX[int(g.xEast)], endLeftRightBorderCoordinatesY[int(g.y)]), assetTargetRightV, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startRightBorderCoordinatesX[int(g.xEast)], startLeftRightBorderCoordinatesY[int(g.y)], endCoordinates[int(g.xEast)], endLeftRightBorderCoordinatesY[int(g.y)]), assetTargetRightV, image.Point{0, 0}, draw.Src)
	}
	// South
	if g.ySouth != g.y {
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startCoordinates[int(g.ySouth)], endCoordinates[int(g.x)], endTopBorderCoordinatesY[int(g.ySouth)]), assetTargetBotH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startBotBorderCoordinatesY[int(g.ySouth)], endCoordinates[int(g.x)], endCoordinates[int(g.ySouth)]), assetTargetBotH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startLeftRightBorderCoordinatesY[int(g.ySouth)], endLeftBorderCoordinatesX[int(g.x)], endLeftRightBorderCoordinatesY[int(g.ySouth)]), assetTargetBotV, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startRightBorderCoordinatesX[int(g.x)], startLeftRightBorderCoordinatesY[int(g.ySouth)], endCoordinates[int(g.x)], endLeftRightBorderCoordinatesY[int(g.ySouth)]), assetTargetBotV, image.Point{0, 0}, draw.Src)
	}
	// West
	if g.xWest != g.x {
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.xWest)], startCoordinates[int(g.y)], endCoordinates[int(g.xWest)], endTopBorderCoordinatesY[int(g.y)]), assetTargetLeftH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.xWest)], startBotBorderCoordinatesY[int(g.y)], endCoordinates[int(g.xWest)], endCoordinates[int(g.y)]), assetTargetLeftH, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.xWest)], startLeftRightBorderCoordinatesY[int(g.y)], endLeftBorderCoordinatesX[int(g.xWest)], endLeftRightBorderCoordinatesY[int(g.y)]), assetTargetLeftV, image.Point{0, 0}, draw.Src)
		draw.Draw(g.Frame, image.Rect(startRightBorderCoordinatesX[int(g.xWest)], startLeftRightBorderCoordinatesY[int(g.y)], endCoordinates[int(g.xWest)], endLeftRightBorderCoordinatesY[int(g.y)]), assetTargetLeftV, image.Point{0, 0}, draw.Src)
	}
}

func (g *Game) drawAttackMarkers() {
	if g.attackFromNorth {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthCoordinates[int(g.x)], startAttackMarkerNorthWestCoordinates[int(g.y)], endAttackMarkerNorthCoordinates[int(g.x)], endAttackMarkerNorthWestCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromNorthEast {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthEastCoordinates[int(g.x)], startAttackMarkerNorthWestCoordinates[int(g.y)], endAttackMarkerNorthEastCoordinates[int(g.x)], endAttackMarkerNorthWestCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromEast {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthEastCoordinates[int(g.x)], startAttackMarkerNorthCoordinates[int(g.y)], endAttackMarkerNorthEastCoordinates[int(g.x)], endAttackMarkerNorthCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromSouthEast {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthEastCoordinates[int(g.x)], startAttackMarkerNorthEastCoordinates[int(g.y)], endAttackMarkerNorthEastCoordinates[int(g.x)], endAttackMarkerNorthEastCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromSouth {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthCoordinates[int(g.x)], startAttackMarkerNorthEastCoordinates[int(g.y)], endAttackMarkerNorthCoordinates[int(g.x)], endAttackMarkerNorthEastCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromSouthWest {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthWestCoordinates[int(g.x)], startAttackMarkerNorthEastCoordinates[int(g.y)], endAttackMarkerNorthWestCoordinates[int(g.x)], endAttackMarkerNorthEastCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromWest {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthWestCoordinates[int(g.x)], startAttackMarkerNorthCoordinates[int(g.y)], endAttackMarkerNorthWestCoordinates[int(g.x)], endAttackMarkerNorthCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
	if g.attackFromNorthWest {
		draw.Draw(g.Frame, image.Rect(startAttackMarkerNorthWestCoordinates[int(g.x)], startAttackMarkerNorthWestCoordinates[int(g.y)], endAttackMarkerNorthWestCoordinates[int(g.x)], endAttackMarkerNorthWestCoordinates[int(g.y)]), assetAttackMarker, image.Point{0, 0}, draw.Over)
	}
}

func (g *Game) drawWSODEnemies() {
	var imageToDraw *image.NRGBA
	for y := g.y - 1; y < g.y+2; y++ {
		for x := g.x - 1; x < g.x+2; x++ {
			if g.cellExists(x, y) {
				switch g.field[y][x] {
				case PawnNorth:
					imageToDraw = assetPawnNorth
				case PawnEast:
					imageToDraw = assetPawnEast
				case PawnSouth:
					imageToDraw = assetPawnSouth
				case PawnWest:
					imageToDraw = assetPawnWest
				case DamagedRook:
					imageToDraw = assetDamagedRook
				case Rook:
					imageToDraw = assetRook
				case DamagedBishop2:
					imageToDraw = assetDamagedBishop2
				case DamagedBishop1:
					imageToDraw = assetDamagedBishop1
				case Bishop:
					imageToDraw = assetBishop
				case DamagedQueen4:
					imageToDraw = assetDamagedQueen4
				case DamagedQueen3:
					imageToDraw = assetDamagedQueen3
				case DamagedQueen2:
					imageToDraw = assetDamagedQueen2
				case DamagedQueen1:
					imageToDraw = assetDamagedQueen1
				case Queen:
					imageToDraw = assetQueen
				case King:
					imageToDraw = assetKing
				case LastShadowPawnNorth:
					imageToDraw = assetLastShadowPawnNorth
				case LastShadowPawnEast:
					imageToDraw = assetLastShadowPawnEast
				case LastShadowPawnSouth:
					imageToDraw = assetLastShadowPawnSouth
				case LastShadowPawnWest:
					imageToDraw = assetLastShadowPawnWest
				case LastShadowRook:
					imageToDraw = assetLastShadowRook
				case LastShadowBishop:
					imageToDraw = assetLastShadowBishop
				case LastShadowQueen:
					imageToDraw = assetLastShadowQueen
				default:
					if x%2 == y%2 {
						imageToDraw = assetBaseWhite
					} else {
						imageToDraw = assetBaseBlack
					}
				}
				draw.Draw(g.Frame, image.Rect(startCoordinates[int(x)], startCoordinates[int(y)], endCoordinates[int(x)], endCoordinates[int(y)]), imageToDraw, image.Point{0, 0}, draw.Src)
			}
		}
	}
}

func (g *Game) drawWSODCharacter() {
	var imageToDraw *image.NRGBA
	if g.HasHair {
		imageToDraw = assetCharDeadHair
	} else {
		imageToDraw = assetCharDead
	}
	draw.Draw(g.Frame, image.Rect(startCoordinates[int(g.x)], startCoordinates[int(g.y)], endCoordinates[int(g.x)], endCoordinates[int(g.y)]), imageToDraw, image.Point{0, 0}, draw.Over)
}

func (g *Game) drawWSODScore() {
	if g.score > 0 {
		stringToPrint := strconv.FormatUint(g.score, 10)
		stringLen := len(stringToPrint)
		var imageToDraw *image.NRGBA
		for i := 0; i < stringLen; i++ {
			switch stringToPrint[i] {
			case '0':
				imageToDraw = assetScoreWSOD0
			case '1':
				imageToDraw = assetScoreWSOD1
			case '2':
				imageToDraw = assetScoreWSOD2
			case '3':
				imageToDraw = assetScoreWSOD3
			case '4':
				imageToDraw = assetScoreWSOD4
			case '5':
				imageToDraw = assetScoreWSOD5
			case '6':
				imageToDraw = assetScoreWSOD6
			case '7':
				imageToDraw = assetScoreWSOD7
			case '8':
				imageToDraw = assetScoreWSOD8
			case '9':
				imageToDraw = assetScoreWSOD9
			}
			draw.Draw(g.Frame, image.Rect(startScoreCoordinates[stringLen][i], Zero, endScoreCoordinates[stringLen][i], GraphicCellPixels), imageToDraw, image.Point{0, 0}, draw.Src)
		}
	}
}
