package main

import "fmt"

// Target
type chessPiece interface {
  move() string
}

// Client
type Player struct {}

func (p *Player) PlayMove(piece chessPiece) {
  fmt.Println(piece.move())
}

// Third Party Adaptees
type thirdPartyKnight struct {}

func (k *thirdPartyKnight) jump() string {
  return "Knight jumps in an L shape"
}

type thirdPartyBishop struct {}

func (b *thirdPartyBishop) slide() string {
  return "Bishop slides diagonally"
}

// Adapters 
type knightAdapter struct {
  knight *thirdPartyKnight
}

func (ka *knightAdapter) move() string {
  return ka.knight.jump()
}

type bishopAdapter struct {
  bishop *thirdPartyBishop
}

func (ba *bishopAdapter) move() string {
  return ba.bishop.slide()
}
