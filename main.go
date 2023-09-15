package main

import "fmt"

type Game struct {
	turn		Color
	board		[8][8]Piece
}

func main() {
	game := NewGame("")
	game.Print()
}

func NewGame(fen string) Game {
	return Game{
		White,
		[8][8]Piece{
			{ WhiteRook,	WhiteKnight,		WhiteBishop,		WhiteQueen,		WhiteKing,	WhiteBishop,	WhiteKnight,	WhiteRook	},
			{ WhitePawn,	WhitePawn,			WhitePawn,			WhitePawn,		WhitePawn,	WhitePawn,		WhitePawn,		WhitePawn	},
			//{ Empty,			Empty,					Empty,					Empty,				Empty,			Empty,				Empty,				Empty	},
			{ Empty,			Empty,					Empty,					Empty,				Empty,			Empty,				Empty,				Empty	},
			{ Empty,			Empty,					Empty,					Empty,				Empty,			Empty,				Empty,				Empty	},
			{ Empty,			Empty,					Empty,					Empty,				Empty,			Empty,				Empty,				Empty	},
			{ Empty,			Empty,					Empty,					Empty,				Empty,			Empty,				Empty,				Empty	},
			//{ Empty,			Empty,					Empty,					Empty,				Empty,			Empty,				Empty,				Empty	},
			{ BlackPawn,	BlackPawn,			BlackPawn,			BlackPawn,		BlackPawn,	BlackPawn,		BlackPawn,		BlackPawn	},
			{ BlackRook,	BlackKnight,		BlackBishop,		BlackQueen,		BlackKing,	BlackBishop,	BlackKnight,	BlackRook	},
		},
	}
}

func (g Game) Print() {
	for _, row := range g.board {
		for _, piece := range row {
			fmt.Printf("%c ", piece.Icon())
		}
		fmt.Println()
	}

	if g.turn == White {
		fmt.Println("Player to play: White")
	} else {
		fmt.Println("Player to play: Black")
	}

	fmt.Println("Current evaluation: ", g.Eval())

	fmt.Println("Available moves")
	for square, moves := range g.AvailableMoves() {
		if square.piece.Color() == g.turn {
			fmt.Printf("%c: %d\n", square.piece.Icon(), moves)
		}
	}
}

type Square struct {
	piece Piece
	x	int
	y int
}

func (g Game) Eval() int {
	eval := 0
	for _, row := range g.board {
		for _, piece := range row {
			switch piece.Color() {
			case White:
				eval += int(piece.Value())
			case Black:
				eval -= int(piece.Value())
			default:
				break
			}
		}
	}
	return eval
}

func (g Game) AvailableMoves() map[Square][]Square {
	// initialize map
	moves := make(map[Square][]Square)
	for x, row := range g.board {
		for y, piece := range row {
			if piece.Color() == g.turn {
				moves[g.Square(x, y)] = g.AvailableMovesFrom(g.Square(x, y))
			}
		}
	}
	return moves
}

func (g Game) Square(x, y int) Square {
	return Square{ g.board[x][y], x, y }
}

func (g Game) AvailableMovesFrom(from Square) []Square {
	piece := g.board[from.x][from.y]

	var moves []Square
	max_distance, deltas := piece.Deltas()
	for _, direction := range deltas {
		for distance := 1; distance <= int(max_distance); distance++ {
			destination := Square{ piece, from.x + (direction.x * distance), from.y + (direction.y * distance) }

			if destination.x < 0 || destination.x > 7 || destination.y < 0 || destination.y > 7 {
				break
			}

			if g.board[destination.x][destination.y].Color() == piece.Color() {
				break
			}

			moves = append(moves, destination)
		}
	}
	return moves
}

func (g Game) Move(from, to Square) {
	g.board[to.x][to.y] = g.board[from.x][from.y]
}
