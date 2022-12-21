package shared

type Board [][7]bool

type StateCollection struct {
  Board Board
}

type DisplayType uint8

const (
  Empty DisplayType = iota
  StandingRock
  FallingRock
)

type BoardDisplay [40][7]DisplayType
