package main

import "fmt"

/*

	Beispiel:

	Diese Implementierung definiert zwei Structs: Piece und PieceTable. Jedes Piece
	repräsentiert einen zusammenhängenden Textbereich im Dokument und PieceTable ist
	die Hauptdatenstruktur, die die Stücke verwaltet und Einfügungen unterstützt.

	Die Funktion NewPieceTable erstellt eine neue PieceTable mit einem Originalstück,
	das das gesamte Dokument darstellt. Die Insert-Methode nimmt einen Offset und die
	einzufügenden Daten an und teilt die betroffenen Teile in drei Teile auf: den Anfang,
	die Mitte und das Ende.

*/

type Piece struct {
	start  int
	end    int
	buffer []byte
}

type PieceTable struct {
	original *Piece
	add      *Piece
	pieces   []*Piece
}

func NewPieceTable(original []byte) *PieceTable {
	orig := &Piece{0, len(original), original}
	return &PieceTable{
		original: orig,
		add:      &Piece{0, 0, []byte{}},
		pieces:   []*Piece{orig},
	}
}

func (p *PieceTable) Insert(offset int, data []byte) {
	add := p.add
	last := p.pieces[len(p.pieces)-1]
	if add.end != 0 {
		p.pieces = append(p.pieces, add)
		last = add
		p.add = &Piece{0, 0, []byte{}}
	}
	if offset == last.end {
		last.buffer = append(last.buffer, data...)
		last.end += len(data)
		return
	}
	var (
		startPiece *Piece
		endPiece   *Piece
		startOff   int
		endOff     int
	)
	for i := range p.pieces {
		piece := p.pieces[i]
		if offset >= piece.end {
			continue
		}
		if offset == piece.start {
			startPiece = piece
			startOff = 0
		} else {
			startPiece = &Piece{piece.start, offset, piece.buffer}
			p.pieces = append(p.pieces, startPiece)
			startOff = offset - piece.start
		}
		if offset+len(data) == piece.end {
			endPiece = piece
			endOff = piece.end - offset
		} else if offset+len(data) < piece.end {
			endPiece = &Piece{offset + len(data), piece.end, piece.buffer}
			p.pieces = append(p.pieces, endPiece)
			endOff = piece.end - (offset + len(data))
		} else {
			endPiece = piece
			endOff = len(piece.buffer) - startOff
		}
		break
	}
	newPiece := &Piece{0, 0, make([]byte, len(data))}
	copy(newPiece.buffer, data)
	newPiece.start = startPiece.end
	newPiece.end = newPiece.start + len(data)
	p.pieces = append(p.pieces, newPiece)
	if endPiece != nil {
		p.pieces = append(p.pieces, &Piece{endPiece.start + endOff, endPiece.end, endPiece.buffer})
	}
}

func (p *PieceTable) GetText() []byte {
	var buf []byte
	for i := range p.pieces {
		piece := p.pieces[i]
		if piece == p.original {
			continue
		}
		if piece == p.add {
			piece = p.original
		}
		buf = append(buf, piece.buffer[piece.start:piece.end]...)
	}
	return buf
}

func main() {

	// Create a new PieceTable with some initial text
	pt := NewPieceTable([]byte("Hello, world!"))

	// Insert some text at offset 7
	pt.Insert(7, []byte("beautiful "))

	// Get the full text from the PieceTable
	text := pt.GetText()
	fmt.Println(string(text))
	// Output: Hello, beautiful world!

}
