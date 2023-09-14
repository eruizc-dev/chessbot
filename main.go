package main

import "fmt"

func main() {
	game:= Game{ 0, DefaultBoard() }

	for i := 0; i < 1; i++ {
		fmt.Printf("Iteration %d\n", i)

		Print(game.board)

		moves := Moves(game)

		fmt.Print(moves)

		//move := moves[240 + rand.Intn(16)]

		//fmt.Printf("Moving %d,%d to %d,%d\n", move.from.x, move.from.y, move.to.x, move.to.y)
		//game.board[move.to.y][move.to.x] = game.board[move.from.y][move.from.x] 
		//game.board[move.from.y][move.from.x] = 0

		game.turn = 1 - game.turn
	}
}

type Piece uint8
type Board [8][8]Piece
type Game struct {
	turn uint8
	board Board
} 

type Position struct {
	x int
	y int
}

type Move struct {
	from Position
	to Position
}

const (
	White_Pawn 		Piece = 0b00000001
	White_Knight	Piece = 0b00000011
	White_Bishop	Piece = 0b01000011
	White_Rook		Piece = 0b00000101
	White_Queen		Piece = 0b01001001
	White_King		Piece = 0b00100000
	Black_Pawn 		Piece = 0b10000001
	Black_Knight	Piece = 0b10000011
	Black_Bishop	Piece = 0b11000011
	Black_Rook		Piece = 0b10000101
	Black_Queen		Piece = 0b11001001
	Black_King		Piece = 0b10100000
)

func Value(p Piece) uint8 {
	return (uint8) (p & 0b00111111)
}

func Color(p Piece) uint8 {
	return (uint8) (p) >> 7
}

func MovesDiagonally(p Piece) bool {
	return (p & 0b01000000) != 0
}

func MovesVertically(p Piece) bool {
	return (p & 0b00001100) != 0
}

func IsKnight(p Piece) bool {
	return p == White_Knight || p == Black_Knight
}

func IsKing(p Piece) bool {
	return (p & 0b00100000) == 0b00100000
}

func FromFemstring(fenstring string) Board {
	return Board { { } }
}

func Moves(game Game) []Move {
	var possible_moves []Move

	for y, row := range game.board {
		for x, piece := range row {
			from := Position{ x, y } 
			possible_moves = PossibleMoves(piece, from, game)
		}
	}

	return possible_moves
}

func PossibleMoves(piece Piece, from Position, game Game) []Move {
	var possible_moves []Move

	if MovesVertically(piece) {
		deltas := []Position { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 } }
		return Foo(from, 8, deltas, possible_moves, game)
	}

	return possible_moves
}

func PossibleMoves2(game Game) []Move {
	possible_moves := make([]Move, 256)

	for y, row := range game.board {
		for x, piece := range row {
			if Color(piece) != game.turn {
				continue
			}

			from := Position{ x, y } 

			if MovesVertically(piece) {
				deltas := []Position { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 } }
				possible_moves = Foo(from, 8, deltas, possible_moves, game)
			}

			if MovesDiagonally(piece) {
				deltas := []Position { { 1, 1 }, { 1, -1 }, { -1, -1 }, { -1, 1 } }
				possible_moves = Foo(from, 8, deltas, possible_moves, game)
			}

			if IsKnight(piece) {
				deltas := []Position {
					{ 1, 2 }, { 2, 1 },
					{ 1, -2 }, { 2, -1 },
					{ -1, -2 }, { -2, -1 },
					{ -1, 2 }, { -2, 1 },
				}
				possible_moves = Foo(from, 1, deltas, possible_moves, game)
			}

			if IsKing(piece) {
				deltas := []Position {
					{ 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 }, // Vertically
					{ 1, 1 }, { 1, -1 }, { -1, -1 }, { -1, 1 }, // Diagonally
				}
				possible_moves = Foo(from, 1, deltas, possible_moves, game)
			}
		}
	}

	return possible_moves
}

func Foo(piece Position, iterations int, deltas []Position, possible_moves []Move, game Game) []Move {
	for _, delta := range deltas {
		for scalar := 1; scalar < iterations; scalar++ {
			x_position := piece.x + (delta.x * scalar)
			y_position := piece.y + (delta.y * scalar)

			if OutOfBounds(x_position, y_position) || Color(game.board[x_position][y_position]) == game.turn {
				break
			}

			possible_moves = append(possible_moves, Move{piece, Position{x_position, y_position}})

			if game.board[x_position][y_position] != 0 {
				break
			}
		}
	}
	return possible_moves
}

func OutOfBounds(x int, y int) bool {
	return x < 0 || x > 7 || y < 0 || y > 7
}

func Print(board Board) {
	for i := len(board) - 1; i >= 0; i -- {
		for _, piece := range board[i] {
			fmt.Printf(" %2d ", Value(piece))
		}
		fmt.Print("\n")
	}
}

func DefaultBoard() Board {
	return Board {
		{ White_Rook, White_Knight, White_Bishop, White_Queen, White_King, White_Bishop, White_Knight, White_Rook },
		//{ White_Pawn, White_Pawn, White_Pawn, White_Pawn, White_Pawn, White_Pawn, White_Pawn, White_Pawn },
		{ 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0 },
		//{ Black_Pawn, Black_Pawn, Black_Pawn, Black_Pawn, Black_Pawn, Black_Pawn, Black_Pawn, Black_Pawn },
		{ Black_Rook, Black_Knight, Black_Bishop, Black_Queen, Black_King, Black_Bishop, Black_Knight, Black_Rook },
	}
}
