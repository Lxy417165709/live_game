package lg

import (
	"fmt"
	"math/rand"
	"time"
)
const (
	FlagOfDeath = iota
	FlagOfLiving
)

type LiveGameBoard struct {
	board          [][]int
	nextStateBoard [][]int
	rows, cols     int
}

func NewLiveGameBoard(rows, cols int, sowingPercent int) *LiveGameBoard {
	const fullPercent = 100
	liveGameBoard := &LiveGameBoard{}
	liveGameBoard.board = get2DSlice(rows, cols)
	liveGameBoard.nextStateBoard = get2DSlice(rows, cols)
	liveGameBoard.rows, liveGameBoard.cols = rows, cols
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rows; i++ {
		for t := 0; t < cols; t++ {
			if rand.Intn(fullPercent) < sowingPercent {
				liveGameBoard.board[i][t] = FlagOfLiving
			} else {
				liveGameBoard.board[i][t] = FlagOfDeath
			}
		}
	}
	return liveGameBoard
}

func (g *LiveGameBoard) NextState() {
	newBoard := get2DSlice(g.rows, g.cols)
	for i := 0; i < g.rows; i++ {
		for t := 0; t < g.cols; t++ {
			countOfAroundLivingCeil := g.getCountOfAroundLivingCell(i, t)
			switch countOfAroundLivingCeil {
			case 2:
				newBoard[i][t] = g.board[i][t]
			case 3:
				newBoard[i][t] = FlagOfLiving
			default:
				newBoard[i][t] = FlagOfDeath
			}
		}
	}
	g.board = newBoard
}

func (g *LiveGameBoard) Show() {
	for i := 0; i < g.rows; i++ {
		for t := 0; t < g.cols; t++ {
			if g.board[i][t] == FlagOfLiving {
				fmt.Print("*")
			}
			if g.board[i][t] == FlagOfDeath {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (g *LiveGameBoard) getCountOfAroundLivingCell(x, y int) int {
	const offset = 1
	countOfAroundLivingCell := 0
	for i := -offset; i <= offset; i++ {
		for t := -offset; t <= offset; t++ {
			refX, refY := (x+i+g.rows)%g.rows, (y+t+g.cols)%g.cols
			if refX == x && refY == y {
				continue
			}
			if g.board[refX][refY] == FlagOfLiving {
				countOfAroundLivingCell++
			}
		}
	}
	return countOfAroundLivingCell
}

func get2DSlice(rows, cols int) [][]int {
	slice := make([][]int, rows)
	for i := 0; i < rows; i++ {
		slice[i] = make([]int, cols)
	}
	return slice
}
