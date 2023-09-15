package main


type Color byte
const (
	None	Color = 0b11000000
	White	Color = 0b00000000
	Black	Color = 0b01000000
)

type Type byte
const (
	Pawn		Type = 0b00001
	Bishop	Type = 0b00011
	Knight	Type = 0b10011
	Rook		Type = 0b00101
	Queen		Type = 0b00111
	King		Type = 0b01111
)

type Piece byte
const (
	Empty				Piece = 0b11000000
	WhitePawn		Piece = 0b00000001
	WhiteBishop	Piece = 0b00000011
	WhiteKnight	Piece = 0b00010011
	WhiteRook		Piece = 0b00000101
	WhiteQueen	Piece = 0b00000111
	WhiteKing		Piece = 0b00001111
	BlackPawn		Piece = 0b01000001
	BlackBishop	Piece = 0b01000011
	BlackKnight	Piece = 0b01010011
	BlackRook		Piece = 0b01000101
	BlackQueen	Piece = 0b01000111
	BlackKing		Piece = 0b01001111
)

func (p Piece) Type() Type {
	return Type(p & 0b00011111)
}

func (p Piece) Color() Color {
	return Color(p & 0b11000000)
}

func (p Piece) Value() uint8 {
	return uint8(p & 0b00001111)
}

func (p Piece) MovesOrthogonally() bool {
	return p & 0b00000100 != 0
}

func (p Piece) MovesDiagonally() bool {
	return p & 0b00000010 != 0
}

type Delta struct {
	x	int
	y int
}

func (p Piece) Deltas() (uint8, []Delta) {
	switch p.Type() {
		case Pawn:
			// TODO: implement pawn's deltas
			return 0, []Delta{}
		case Bishop:
			return 7, []Delta{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
		case Knight:
			return 1, []Delta{{1, 2}, {2, 1}, {-1, 2}, {-2, 1}, {1, -2}, {2, -1}, {-1, -2}, {-2, -1}}
		case Rook:
			return 7, []Delta{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		case Queen:
			return 7, []Delta{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		case King:
			return 1, []Delta{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		default:
			return 0, nil
	}
}

func (p Piece) Icon() rune {
	switch p {
	case WhitePawn:
		return '♙'
	case BlackPawn:
		return '♟'
	case WhiteBishop:
		return '♗'
	case BlackBishop:
		return '♝'
	case WhiteKnight:
		return '♘'
	case BlackKnight:
		return '♞'
	case WhiteRook:
		return '♖'
	case BlackRook:
		return '♜'
	case WhiteQueen:
		return '♕'
	case BlackQueen:
		return '♛'
	case WhiteKing:
		return '♔'
	case BlackKing:
		return '♚'
	default:
		return '·'
	}
}
