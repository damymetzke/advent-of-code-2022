package shared

type Position struct {
	X int
	Y int
}

type Board [][7]bool

type Piece []Position

type StateCollection struct {
	Board         Board
	PieceType     int
	PiecePosition Position
}

type DisplayType uint8

const (
	Empty DisplayType = iota
	StandingRock
	FallingRock
)

type BoardDisplay [40][7]DisplayType

func GetPiece(id int) Piece {
	switch id {
	case 0:
		return Piece{
			Position{
				X: 0,
				Y: 0,
			},
			Position{
				X: 1,
				Y: 0,
			},
			Position{
				X: 2,
				Y: 0,
			},
			Position{
				X: 3,
				Y: 0,
			},
		}
	case 1:
		return Piece{
			Position{
				X: 1,
				Y: 0,
			},
			Position{
				X: 0,
				Y: 1,
			},
			Position{
				X: 1,
				Y: 1,
			},
			Position{
				X: 2,
				Y: 1,
			},
			Position{
				X: 1,
				Y: 2,
			},
		}
	case 2:
		return Piece{
			Position{
				X: 0,
				Y: 0,
			},
			Position{
				X: 1,
				Y: 0,
			},
			Position{
				X: 2,
				Y: 0,
			},
			Position{
				X: 2,
				Y: 1,
			},
			Position{
				X: 2,
				Y: 2,
			},
		}
	case 3:
		return Piece{
			Position{
				X: 0,
				Y: 0,
			},
			Position{
				X: 0,
				Y: 1,
			},
			Position{
				X: 0,
				Y: 2,
			},
			Position{
				X: 0,
				Y: 3,
			},
		}
	case 4:
		return Piece{
			Position{
				X: 0,
				Y: 0,
			},
			Position{
				X: 1,
				Y: 0,
			},
			Position{
				X: 0,
				Y: 1,
			},
			Position{
				X: 1,
				Y: 1,
			},
		}
	}
	return Piece{}
}
