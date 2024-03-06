package yesgram

var BLOCKED = -1
var FILLED = 1
var NEUTRAL = 0

type board struct {
	xSize int
	ySize int
	state []int
}

func getNewBoard(xSize, ySize int) board {
	newBoard := board{}
	newBoard.xSize = xSize
	newBoard.ySize = ySize
	newBoard.state = make([]int, xSize*ySize)
	return newBoard
}

func (b *board) getRowLine(row int) []int {
	return b.state[row*b.xSize : (row+1)*b.xSize]
}

func (b *board) setRowLine(lineState []int, row int) {
	for index, state := range lineState {
		b.state[row*b.xSize+index] = state
	}
}

func (b *board) getColumnLine(column int) []int {
	lineState := make([]int, b.ySize)
	for index := 0; index < b.ySize; index++ {
		lineState[index] = b.state[index*b.xSize+column]
	}
	return lineState
}

func (b *board) setColumnLine(lineState []int, column int) {
	for index := 0; index < b.ySize; index++ {
		b.state[index*b.xSize+column] = lineState[index]
	}
}
