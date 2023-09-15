package main

//import "testing"

//func TestValue(t *testing.T) {
//	expected_values := map[Piece]uint8{
//		White_Pawn 	: 1,
//		White_Knight: 3,
//		White_Bishop: 3,
//		White_Rook	: 5,
//		White_Queen	: 9,
//		White_King	: 32,
//		Black_Pawn 	: 1,
//		Black_Knight: 3,
//		Black_Bishop: 3,
//		Black_Rook	: 5,
//		Black_Queen	: 9,
//		Black_King	: 32,
//	}
//
//	for piece, expected_value := range expected_values {
//		value := Value(piece)
//		if value != expected_value {
//			t.Fatalf("Piece %d expected to be worth %d but was %d", piece, expected_value, value);
//		}
//	}
//}
//
//func TestMovesVertically(t *testing.T) {
//	expected_values := map[Piece]bool{
//		White_Pawn 	: false,
//		White_Knight: false,
//		White_Bishop: false,
//		White_Rook	: true,
//		White_Queen	: true,
//		White_King	: false,
//		Black_Pawn 	: false,
//		Black_Knight: false,
//		Black_Bishop: false,
//		Black_Rook	: true,
//		Black_Queen	: true,
//		Black_King	: false,
//	}
//
//	for piece, expected_value := range expected_values {
//		value := MovesVertically(piece)
//		if value != expected_value {
//			if expected_value == true {
//				t.Fatalf("Piece %d should be allowed to move vertically", piece);
//			} else {
//				t.Fatalf("Piece %d should not be allowed to move vertically", piece);
//			}
//		}
//	}
//}
//
//func TestMovesDiagonally(t *testing.T) {
//	expected_values := map[Piece]bool{
//		White_Pawn 	: false,
//		White_Knight: false,
//		White_Bishop: true,
//		White_Rook	: false,
//		White_Queen	: true,
//		White_King	: false,
//		Black_Pawn 	: false,
//		Black_Knight: false,
//		Black_Bishop: true,
//		Black_Rook	: false,
//		Black_Queen	: true,
//		Black_King	: false,
//	}
//
//	for piece, expected_value := range expected_values {
//		value := MovesDiagonally(piece)
//		if value != expected_value {
//			if expected_value == true {
//				t.Fatalf("Piece %d should be allowed to move diagonally", piece);
//			} else {
//				t.Fatalf("Piece %d should not be allowed to move diagonally", piece);
//			}
//		}
//	}
//}
