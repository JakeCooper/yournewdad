package main

func headOn(data *MoveRequest, sid int) bool {
	theirHead := data.Snakes.Data[sid].Head()
	myHead := data.Snakes.Data[data.MyIndex].Head()

	myHeadToTheirHead := myHead.Dist(theirHead)
	if myHeadToTheirHead.X+myHeadToTheirHead.Y != 4 || len(data.Snakes.Data[sid].Coords.Data) < data.MyLength {
		return false
	}
	// returns the first piece of a snakes body
	myFirstBody := &(data.Snakes.Data[data.MyIndex].Coords.Data[1])
	theirFirstBody := &(data.Snakes.Data[sid].Coords.Data[1])
	myHeadToTheirBody := myHead.Dist(theirFirstBody)
	theirHeadToMyBody := theirHead.Dist(myFirstBody)

	if (totalDist(myHeadToTheirBody) > totalDist(myHeadToTheirHead)) && (totalDist(theirHeadToMyBody) > totalDist(myHeadToTheirHead)) {
		return true
	}
	return false
}

func totalDist(p *Point) int {
	return p.X + p.Y
}
