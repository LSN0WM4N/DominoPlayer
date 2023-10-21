package dominoplayer

import (
	"encoding/json"
	"net/http"
)

type start struct {
	St bool `json:"start"`
}

type reset struct {
	Position   int     `json:"position"`
	Pieces     [][]int `json:"pieces"`
	Max_number int     `json:"max_number"`
	Timeout    int     `json:"timeout"`
	Score      int     `json:"score"`
}

type step struct {
	Heads []int32 `json:"heads"`
}

type stepAnswer struct {
	Piece []int32 `json:"piece"`
	Head  int32   `json:"head"`
}

type Piece struct {
	x int32
	y int32
}

var pieces []Piece

func (this *Piece) used() {
	this.x = -1
	this.y = -1
}

func (this *Piece) isAble() bool {
	return (this.x != -1 && this.y != -1)
}

func (this *Piece) match(piece Piece) bool {
	if !this.isAble() {
		return false
	}

	return (this.x == piece.x || this.x == piece.y || this.y == piece.x || this.y == piece.y)
}

func CreatePiece(arr []int32) Piece {
	if arr[0] > arr[1] {
		arr[0], arr[1] = arr[1], arr[0]
	}

	return Piece{x: arr[0], y: arr[1]}
}

func isEqual(a, b Piece) bool {
	return (a.x == b.x && a.y == b.y) || (a.y == b.x && a.x == b.y)
}

func ValidMoves(piece Piece) []Piece {
	var answer []Piece
	if !piece.isAble() {
		return mainGame.Pieces
	}

	for _, i := range mainGame.Pieces {
		if i.isAble() && i.match(piece) {
			answer = append(answer, i)
		}
	}

	return answer
}

func Chose(piece Piece) Piece {
	moves := ValidMoves(piece)

	var sum int32 = -1
	var ans Piece = Piece{x: -1, y: -1}

	for _, i := range moves {
		if i.x+i.y >= sum {
			ans = i
			sum = i.x + i.y
		}
	}

	return ans
}

func Remove(piece Piece) {
	for index, i := range mainGame.Pieces {
		if isEqual(i, piece) {
			mainGame.Pieces[index].used()
		}
	}
}

func postReset(r *http.Request) reset {
	var answer reset
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&answer)
	return answer
}

func getStep(r *http.Request) step {
	var answer step
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&answer)
	return answer
}

func postStep(myStep stepAnswer) []byte {
	answer, _ := json.Marshal(myStep)
	return answer
}

func parsePieces(pieces [][]int) []Piece {
	var answer []Piece

	for _, i := range pieces {
		piece := Piece{x: int32(i[0]), y: int32(i[1])}
		answer = append(answer, piece)
	}

	return answer
}

func selectHead(piece Piece, heads []int32) int32 {
	if piece.x == heads[0] || piece.y == heads[0] {
		return 0
	} else {
		return 1
	}
}
