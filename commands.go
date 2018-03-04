package main

import (
	"fmt"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
}

func handleStart(res http.ResponseWriter, req *http.Request) {
	respond(res, GameStartResponse{
		Taunt:    toStringPointer("Trustless, just like my ex wife."),
		Color:    "gold",
		Name:     "Snakes on a Blockchain",
		HeadType: "shades",
		TailType: "fat-rattle",
		HeadUrl:  toStringPointer("https://i.imgur.com/a/Xg0ki"),
	})
}

func handleMove(res http.ResponseWriter, req *http.Request) {
	//ctx := appengine.NewContext(req)
	str := getMoveRequestString(req)

	// log the json blob that comes in if requested
	logging := os.Getenv("YND_LOG")
	if len(logging) > 0 {
		log.Printf(str)
	}
	fmt.Println(str)
	data, err := NewMoveRequest(str)
	if err != nil {
		respond(res, MoveResponse{
			Move:  "up",
			Taunt: "can't parse this!",
		})
		return
	}

	// log move request
	//log.Infof(ctx, "%v", data)
	//if appengine.IsDevAppServer() {
	//	if imAgressive(data) {
	//		log.Infof(ctx, stringAllMinMAX(data))
	//	}
	//}

	move, err := getMove(data, req)

	if err != nil {
		respond(res, MoveResponse{
			Move:  "up",
			Taunt: "Couldn't parse",
		})
		//log.Errorf(ctx, "Could not find a move for this data")
		return
	}
	taunt := getTaunt(data.Turn)
	respond(res, MoveResponse{
		Move:  move,
		Taunt: taunt,
	})
}

func getMove(data *MoveRequest, req *http.Request) (string, error) {
	//ctx := appengine.NewContext(req)

	moves, err := bestMoves(data)

	if err != nil {
		//log.Errorf(ctx, "generating MetaData: %v", err)
		return "", err
	}

	//log.Printf("%v\n", moves)
	if len(moves) < 1 {
		return "", err
	}

	rand.Seed(time.Now().Unix())

	return moves[rand.Intn(len(moves))], nil
}
