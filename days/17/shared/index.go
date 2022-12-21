package shared

type DisplayType uint8

const (
  Empty DisplayType = iota
  StandingRock
  FallingRock
)

type BoardDisplay [40][7]DisplayType
