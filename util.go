package main

import (
	"bytes"
	"errors"
	"fmt"
)

func swap(arr []*Snake, a, b int) {
	arr[b], arr[a] = arr[a], arr[b]
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MinMax(data *MoveRequest, direc string) {
	// generated the hazards without the hazards around the other snakes

	data.GenHazards(data, false)
	myHead := data.Snakes.Data[data.MyIndex].Head()
	if direc != "" {
		myHeadtmp, err := GetPointInDirection(myHead, direc, data)
		if err != nil {
			return
		}
		myHead = myHeadtmp
		if myHead == nil {
			return
		}
		if myHead != nil {
			data.Hazards[myHead.String()] = true
		}
	}

	ret := quickStats2(data, direc)

	if direc == "" && len(ret.sortedFood) > 0 {
		data.DistToFood = ret.sortedFood[0].moves
	}
	if direc != "" {
		data.Direcs[direc].ClosestFood = ret.ClosestFood
		data.Direcs[direc].Food = ret.Food
		data.Direcs[direc].Moves = ret.Moves
		data.Direcs[direc].SeeTail = ret.SeeTail
		data.Direcs[direc].KeySnakeData = ret.KeySnakeData
		data.Direcs[direc].FoodHash = ret.FoodHash
		data.Direcs[direc].sortedFood = ret.sortedFood
		data.Direcs[direc].MoveHash = ret.MoveHash
	}

	if direc != "" {
		data.Direcs[direc].minMaxArr = ret.minMaxArr
	} else if direc == "" {
		data.minMaxArr = ret.minMaxArr
		data.KSD = ret.KeySnakeData
	}

}

func GenMinMaxStats(arr MMArray) MinMaxMetaData {
	ret := MinMaxMetaData{}
	ret.movesHash = make(map[string]int)
	ret.tiesHash = make(map[string][]int)
	ret.snakes = make(map[int]MinMaxSnakeMD)
	for i := range arr {
		for j := range arr[i] {
			p := &Point{X: i, Y: j}
			ids := arr[i][j].snakeIds

			if arr[i][j].tie {
				ret.tiesHash[p.String()] = ids
			}

			for _, id := range ids {
				s, ok := ret.snakes[id]
				if !ok {
					ret.snakes[id] = MinMaxSnakeMD{}
				}
				if arr[i][j].tie {
					s.ties++
				} else {
					s.moves++
					ret.movesHash[p.String()] = id
				}
				ret.snakes[id] = s
			}
		}
	}
	return ret
}

func stringAllMinMAX(data *MoveRequest) string {
	var buffer bytes.Buffer
	buffer.WriteString("\n board\n ")
	buffer.WriteString(data.minMaxArr.String())
	for _, direc := range []string{UP, RIGHT, DOWN, LEFT} {
		if data.Direcs[direc].minMaxArr != nil {
			buffer.WriteString(fmt.Sprintf("%v\n", direc))
			buffer.WriteString(data.Direcs[direc].minMaxArr.String())
		}
	}
	return buffer.String()
}

func findGuaranteedClosestFood(data *MoveRequest, direc string) *FoodData {
	for _, food := range data.Direcs[direc].sortedFood {
		for _, id := range data.minMaxArr[food.pnt.Y][food.pnt.X].snakeIds {
			if id == data.MyIndex {
				return food
			}
		}
	}
	return nil
}

func getTail(ind int, data *MoveRequest) (*Point, error) {
	if (ind < 0) || (ind >= len(data.Snakes.Data)) {
		return nil, errors.New("Index out of bounds")
	}
	snake := data.Snakes.Data[ind]
	fmt.Println(snake)
	return &(snake.Coords.Data[len(snake.Coords.Data)-1]), nil

}

func IsSnakeHead(p *Point, data *MoveRequest) bool {
	if p != nil && data.SnakeHeads[p.String()] {
		return true
	}
	return false
}

func getTaunt(turn int) string {
	if turn < 30 {
		return "Distributed Consensus"
	} else if turn < 60 {
		return "Trustless, just like my ex-wife"
	} else if turn < 90 {
		return "Craig Wright is Satoshi Nakamoto"
	} else if turn < 120 {
		return "BLOCKCHAINNNN"
	} else if turn < 150 {
		return "Job Posting: Blockchain 3.0 Developers ONLY"
	} else if turn < 180 {
		return "Future of the internet"
	} else if turn < 180 {
		return "BLOCK-MF-CHAIN"
	} else if turn < 200 {
		return "Reverse Mortgaged my house to buy bitcoin"
	} else if turn < 250 {
		return "Homeless: Accepting ETH"
	} else if turn < 300 {
		return "AI is out, blockchains are in"
	} else if turn < 350 {
		return "If this server had a blockchain maybe it would crash less"
	} else if turn < 400 {
		return "POS: Proof of stake or piece of sh*t; you decide"
	} else if turn < 450 {
		return "ETH > BTC"
	} else if turn < 475 {
		return "Peep my ICO live after this event. http://www.definitelynotascam.net/"
	} else if turn < 500 {
		return "Dogs on the blockchain"
	} else if turn < 550 {
		return "Ripple is garbage"
	} else if turn < 600 {
		return "BITCONNNNNNNNNECTTTTTTT"
	} else if turn < 650 {
		return "Elixir isn't webscale. Long live Golang."
	} else if turn < 675 {
		return "Even Bitcoin has more transactions than this server."
	} else if turn < 700 {
		return "Send ETH: 0xBAe33EC3765498AA53B386420A2DCAe93E343C6a"
	}
	return "Off The Blockchain"
}

// get the position of all neighbouring snake tiles and
// return the snake data corresponding to the last piece
// of snake that you see
// if there are no snakes around you return nil
func FindMinSnakePointInSurroundingArea(p *Point, data *MoveRequest, KeySnakeData map[int]*SnakeData) {
	pts := []*Point{
		p.UpHazard(data),
		p.DownHazard(data),
		p.LeftHazard(data),
		p.RightHazard(data)}

	for _, pt := range pts {
		if pt != nil {
			sd := data.SnakeHash[pt.String()]
			if sd != nil {
				if KeySnakeData[sd.id] == nil || sd.lengthLeft < KeySnakeData[sd.id].lengthLeft {
					KeySnakeData[sd.id] = sd
				}
			}
		}
	}
}

// returns the number of valid neighbours to a point p
func GetNumNeighbours(data *MoveRequest, p *Point) (int, error) {
	if p == nil {
		return 0, nil
	}
	neighbours := 0
	for _, d := range []string{UP, DOWN, LEFT, RIGHT} {
		neighbour, err := GetPointInDirection(p, d, data)
		if err != nil {
			return 0, err
		}
		//fmt.Printf("In Loop neighbour %v, %v\n", p, d)
		if neighbour != nil {
			neighbours += 1
		}
	}
	//fmt.Printf("getting neighbours %v, %v\n", direc, neighbours)
	return neighbours, nil
}

// returns a point representing traven in the direction direc
// i.e. if you pass in direc "up" it will give you the point
// that is above p
// will only return points that are valid moves
func GetPointInDirection(p *Point, direc string, data *MoveRequest) (*Point, error) {
	if p == nil {
		return nil, nil
	}
	switch direc {
	case UP:
		return p.Up(data), nil
	case DOWN:
		return p.Down(data), nil
	case LEFT:
		return p.Left(data), nil
	case RIGHT:
		return p.Right(data), nil
	}
	return nil, errors.New(fmt.Sprintf("could not find direction %v", direc))
}

func GetPointInDirectionHazards(p *Point, direc string, data *MoveRequest) (*Point, error) {
	if p == nil {
		return nil, nil
	}
	switch direc {
	case UP:
		return p.UpHazard(data), nil
	case DOWN:
		return p.DownHazard(data), nil
	case LEFT:
		return p.LeftHazard(data), nil
	case RIGHT:
		return p.RightHazard(data), nil
	}
	return nil, errors.New(fmt.Sprintf("could not find direction %v", direc))
}

func toStringPointer(str string) *string {
	return &str
}

func getMyHead(data *MoveRequest) (*Point, error) {
	for _, snake := range data.Snakes.Data {
		if snake.Id == data.You.Id && len(data.You.Body.Data) > 0 {
			return snake.Head(), nil
		}
	}
	return &Point{}, errors.New("Could not get head")
}

func getMyTail(data *MoveRequest) (*Point, error) {
	return getTail(data.MyIndex, data)
}
