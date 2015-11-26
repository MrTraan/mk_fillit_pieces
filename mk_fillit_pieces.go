package main

import (
	"bytes"
	"math/rand"
	"os"
	"time"
)

type maxOff struct {
	x, y int
}

var pMaxOff = map[int]maxOff{
	0: {3, 0},
	1: {1, 2},
	2: {1, 2},
	3: {2, 2},
	4: {2, 1},
	5: {2, 1},
	6: {1, 2},
}

var pRaw = map[int][4][]byte{
	0: {{'#', '.', '.', '.'}, {'#', '.', '.', '.'}, {'#', '.', '.', '.'}, {'#', '.', '.', '.'}},
	1: {{'.', '#', '#', '.'}, {'#', '#', '.', '.'}, {'.', '.', '.', '.'}, {'.', '.', '.', '.'}},
	2: {{'#', '#', '.', '.'}, {'.', '#', '#', '.'}, {'.', '.', '.', '.'}, {'.', '.', '.', '.'}},
	3: {{'#', '#', '.', '.'}, {'#', '#', '.', '.'}, {'.', '.', '.', '.'}, {'.', '.', '.', '.'}},
	4: {{'#', '.', '.', '.'}, {'#', '.', '.', '.'}, {'#', '#', '.', '.'}, {'.', '.', '.', '.'}},
	5: {{'.', '#', '.', '.'}, {'.', '#', '.', '.'}, {'#', '#', '.', '.'}, {'.', '.', '.', '.'}},
	6: {{'.', '#', '.', '.'}, {'#', '#', '#', '.'}, {'.', '.', '.', '.'}, {'.', '.', '.', '.'}},
}

func AddPiece(b *bytes.Buffer) {
	pNum := 6 //rand.Intn(len(pRaw))
	offX := rand.Intn(pMaxOff[pNum].x + 1)
	offY := rand.Intn(pMaxOff[pNum].y + 1)
	for i := 0; i < offY; i++ {
		b.Write([]byte{'.', '.', '.', '.', '\n'})
	}
	for i := 0; i < 4-offY; i++ {
		b.Write([]byte{'.', '.', '.', '.'}[:offX])
		b.Write(pRaw[pNum][i][:4-offX])
		b.WriteByte('\n')
	}
}

func main() {
	var b bytes.Buffer

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(27)
	for i := 0; i <= n; i++ {
		AddPiece(&b)
		if i != n {
			b.WriteByte('\n')
		}
	}
	b.WriteTo(os.Stdout)
}
