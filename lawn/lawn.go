package lawn

import (
	"plants-vs-zombies/characters"
	"strconv"
)

type Lawn interface {
	GetPosition(xAxis int, yAxis int) Position
	GetShooter(position Position) characters.Shooter
}

type lawn struct {
	positions []Position
	shooters  map[Position]characters.Shooter
}

func New(boardSetup []string) Lawn {
	positions := initLawnPositions(boardSetup)

	row := boardSetup[0]
	shooters := initRowShooters(row)

	return lawn{positions: positions, shooters: shooters}
}

func (l lawn) GetShooter(position Position) characters.Shooter {
	return l.shooters[position]
}

func (l lawn) GetPosition(xAxis int, yAxis int) Position {
	position := Position{Column: xAxis, Row: yAxis}
	return position
}

func initRowShooters(boardRow string) map[Position]characters.Shooter {
	shooters := make(map[Position]characters.Shooter)
	shooters = fillRowShooters(boardRow, shooters)
	return shooters
}

func fillRowShooters(boardRow string, shooters map[Position]characters.Shooter) map[Position]characters.Shooter {
	for i := 0; i < len(boardRow); i++ {
		position := Position{Column: i, Row: 0}
		shooter := parseShooterInBoardRow(boardRow, i)
		shooters[position] = shooter
	}
	return shooters
}

func parseShooterInBoardRow(boardRow string, shooterIndex int) characters.Shooter {
	powerStr := string(boardRow[shooterIndex])
	power, err := strconv.Atoi(powerStr)
	if err != nil {
		panic("invalid attack power")
	}
	return characters.Shooter{AttackPower: power}
}

func initLawnPositions(boardSetup []string) []Position {
	positions := make([]Position, len(boardSetup))
	return positions
}