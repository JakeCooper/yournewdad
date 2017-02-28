package kaa

import (
	"reflect"
	"sort"
	"testing"
)

func Test_FilterKillZones(t *testing.T) {
	data, err := NewMoveRequest(gameString16)

	if err != nil {
		t.Errorf("%v", err)
	}

	directions, err := bestMoves(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, dir := range directions {
		if dir == RIGHT {
			t.Errorf("If you go right you kill yourself what is wrong with you", directions)
		}
	}

}

func Test_FilterPossibleMoves(t *testing.T) {
	data, err := NewMoveRequest(gameString3)
	if err != nil {
		t.Errorf("%v", err)
	}

	directions := FilterPossibleMoves(data, []string{UP, DOWN, LEFT, RIGHT})

	notRightOrDown := []string{LEFT, UP}

	// sort both of the strings so that deep equal will be able to see them
	sort.Strings(notRightOrDown)
	sort.Strings(directions)

	if !reflect.DeepEqual(directions, notRightOrDown) {
		t.Errorf("expected all directions except down, got %v", directions)
	}

}

func Test_FilterTail(t *testing.T) {
	data, err := NewMoveRequest(gameString15)

	if err != nil {
		t.Errorf("%v", err)
	}

	directions, err := bestMoves(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	right := []string{RIGHT}

	// sort both of the strings so that deep equal will be able to see them
	sort.Strings(right)
	sort.Strings(directions)

	if !reflect.DeepEqual(directions, right) {
		t.Errorf("expected right, got %v", directions)
	}

}

func Test_MovingIntoTightSpaces(t *testing.T) {
	data, err := NewMoveRequest(gameString7)
	if err != nil {
		t.Errorf("%v", err)
	}

	directions := []string{LEFT, RIGHT}
	foodDirections := FilterClosestFoodDirections(data, directions)
	expectedFoodDirections := []string{RIGHT}
	sort.Strings(foodDirections)
	sort.Strings(expectedFoodDirections)

	if !reflect.DeepEqual(foodDirections, expectedFoodDirections) {
		t.Errorf("expected %v directions got %v", expectedFoodDirections, foodDirections)
	}
}

func Test_ClosestFood(t *testing.T) {
	data, err := NewMoveRequest(gameString3)
	if err != nil {
		t.Errorf("%v", err)
	}

	directions := []string{LEFT, UP}
	foodDirections := FilterClosestFoodDirections(data, directions)
	expectedFoodDirections := []string{LEFT, UP}
	sort.Strings(foodDirections)
	sort.Strings(expectedFoodDirections)

	if !reflect.DeepEqual(foodDirections, expectedFoodDirections) {
		t.Errorf("expected %v directions got %v", expectedFoodDirections, foodDirections)
	}

}

func Test_ClosestFood3(t *testing.T) {
	data, err := NewMoveRequest(gameString13)
	if err != nil {
		t.Errorf("%v", err)
	}

	foodDirections, err := bestMoves(data)
	if err != nil {
		t.Errorf("%v", err)
	}
	expectedFoodDirections := []string{LEFT, DOWN}
	sort.Strings(foodDirections)
	sort.Strings(expectedFoodDirections)

	if !reflect.DeepEqual(foodDirections, expectedFoodDirections) {
		t.Errorf("expected %v directions got %v", expectedFoodDirections, foodDirections)
	}

}

func Test_MinimizationOfSpace(t *testing.T) {
	data, err := NewMoveRequest(gameString8)
	if err != nil {
		t.Errorf("%v", err)
	}

	directions := []string{LEFT, DOWN}
	filteredMoves := FilterMinimizeSpace(data, directions)

	expectedDirection := []string{LEFT}
	sort.Strings(filteredMoves)
	sort.Strings(expectedDirection)

	if !reflect.DeepEqual(expectedDirection, filteredMoves) {
		t.Errorf("expected %v directions got %v", expectedDirection, filteredMoves)
	}
}

func TestDontMoveOntoTheKeyArea(t *testing.T) {
	data, err := NewMoveRequest(gameString6)
	if err != nil {
		t.Errorf("%v", err)
	}

	directions := []string{RIGHT, UP}
	filteredMoves := FilterKeyArea(data, directions)

	expectedDirection := []string{RIGHT}
	sort.Strings(filteredMoves)
	sort.Strings(expectedDirection)

	if !reflect.DeepEqual(expectedDirection, filteredMoves) {
		t.Errorf("expected %v directions got %v", expectedDirection, filteredMoves)
	}

}

func Test_FilteringSpace(t *testing.T) {
	data, err := NewMoveRequest(gameString11)
	if err != nil {
		t.Errorf("%v", err)
	}

	filteredMoves, err := bestMoves(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedDirection := []string{LEFT}
	sort.Strings(filteredMoves)
	sort.Strings(expectedDirection)

	if !reflect.DeepEqual(expectedDirection, filteredMoves) {
		t.Errorf("expected %v directions got %v", expectedDirection, filteredMoves)
	}

}

func Test_CutOff(t *testing.T) {
	data, err := NewMoveRequest(gameString10)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	filteredMoves, err := bestMoves(data)
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedDirection := []string{DOWN}
	sort.Strings(filteredMoves)
	sort.Strings(expectedDirection)

	if !reflect.DeepEqual(expectedDirection, filteredMoves) {
		t.Errorf("expected %v directions got %v", expectedDirection, filteredMoves)
	}
}
