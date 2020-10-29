package board

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrieves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	x, err := at.Coord(0)
	y, err1 := at.Coord(1)

	if err != nil || err1 != nil {
		return nil
	}

	piece := c[x][y]

	if piece == nil {
		return nil
	}

	return piece
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	if c.PieceAt(from) == nil {
		return fmt.Errorf("No piece at %v", from.String())
	} else if c.PieceAt(to) != nil {
		return fmt.Errorf("There is a piece at %v", to.String())
	} else {
		x, _ := from.Coord(0)
		y, _ := from.Coord(1)
		place := c[x][y]
		return c.PlacePieceAt(place, to)
	}
}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	if c.PieceAt(at) != nil {
		return fmt.Errorf("There is a piece at %v", at.String())
	}

	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	c[x][y] = p

	return nil
}
