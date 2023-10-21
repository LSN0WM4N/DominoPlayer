package dominoplayer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Game struct {
	Name     string  `json:"name"`
	Position int     `json:"position"`
	Pieces   []Piece `json:"pieces"`
	History  [][]int `json:"history"`
	Heads    []int32 `json:"heads"`
}

var mainGame Game = Game{}

func Start(w http.ResponseWriter, r *http.Request) {
	log.Println(w)
	var answer start = start{St: true}
	output, _ := json.Marshal(answer)
	w.Write(output)
}

func Reset(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "application/json")
	resetData := postReset(r)

	mainGame.Name = "SN0WM4N"
	mainGame.Position = resetData.Position
	mainGame.Pieces = parsePieces(resetData.Pieces)
	mainGame.History = nil
	mainGame.Heads = []int32{-1, -1}
}

func Step(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "application/json")
	stepData := getStep(r)

	mainGame.Heads = stepData.Heads
	shouldPass := true

	if mainGame.Heads[0] == -1 || mainGame.Heads[1] == -1 {
		shouldPass = false
	} else {
		moves := ValidMoves(CreatePiece(mainGame.Heads))

		if len(moves) != 0 {
			shouldPass = false
		}
	}

	if shouldPass {
		var answer = []byte(`{"piece": null, "head": null}`)
		w.Write(answer)
	} else {
		piece := Chose(CreatePiece(mainGame.Heads))
		var answer stepAnswer = stepAnswer{[]int32{piece.x, piece.y}, selectHead(piece, mainGame.Heads)}
		w.Write(postStep(answer))
		fmt.Println(mainGame.Pieces)
		fmt.Println(answer)

		Remove(piece)
	}
}
