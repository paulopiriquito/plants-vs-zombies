package game_test

import (
	"plants-vs-zombies/game"
	"strconv"
	"testing"
)

func TestGameSetup(t *testing.T) {
	t.Run("Given an empty lawn map, when the game is setup, then the game should be playable", TestGameWithEmptyLawn)
	t.Run("Given a lawn map with a basic shooter, when the game is setup, then should have a basic shooter", TestGameWithSingleShooter)
	t.Run("Given a lawn map with multiple shooters, when I retrieve one shooter, then should must match the one in setup", TestGameWithMultipleShooters)
}

func TestGameOverStates(t *testing.T) {

}

func TestGameWithMultipleShooters(t *testing.T) {
	testShooterPower := 2
	testLawnMap := []string{"1" + strconv.Itoa(testShooterPower)}
	testGame := game.NewGame(testLawnMap)
	lawn := testGame.Lawn()

	shooterPosition := lawn.GetPosition(1, 0)
	testShooter := lawn.GetShooter(shooterPosition)

	if testShooter.AttackPower != testShooterPower {
		t.Error("testShooter must have the setup attack power")
	}
}

func TestGameWithSingleShooter(t *testing.T) {
	testShooterPower := 3
	lawnMap := []string{strconv.Itoa(testShooterPower)}
	testGame := game.NewGame(lawnMap)
	lawn := testGame.Lawn()

	position := lawn.GetPosition(0, 0)
	testShooter := lawn.GetShooter(position)

	if testShooter.AttackPower != testShooterPower {
		t.Error("testShooter must have the setup attack power")
	}
}

func TestGameWithEmptyLawn(t *testing.T) {
	lawnMap := []string{""}
	testGame := game.NewGame(lawnMap)

	err := testGame.Start()
	if err != nil {
		t.Error("testGame must be playable")
	}
}