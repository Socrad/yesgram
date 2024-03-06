package yesgram

import "fmt"

type game struct {
	rowHints    [][]int
	columnHints [][]int
	rowCases    [][][]int
	columnCases [][][]int
	gameBoard   board
}

func NewGame(rowHints, columnHints [][]int) game {
	newGame := game{}
	newGame.rowHints = rowHints
	newGame.columnHints = columnHints
	newGame.gameBoard = getNewBoard(len(columnHints), len(rowHints))
	newGame.rowCases = getCases(newGame.rowHints, newGame.gameBoard.xSize)
	newGame.columnCases = getCases(newGame.columnHints, newGame.gameBoard.ySize)
	return newGame
}

/*
	노노그램을 푼다

모든 행에 대해 현재 상태와 상충되는 케이스를 제거 후 모든 케이스에 공통으로 겹쳐서 칠해지거나 막힐 수 밖에 없는 타일을 칠하고 막는다.
모든 열에 대해서도 같은 작업을 수행한다.
열에 대해 수행하는 과정에서 행의 상태가 바뀌었다면 행에 대해 다시 수행하면 행의 상태가 바뀔 수 있다. 열을 다시 수행할 때도 마찬가지다.
더 이상 변경이 없을때까지 위 과정을 반복한다.
*/
func (g *game) Solve() {
	isComplete := true
	for isComplete {
		isComplete = g.conjunctionCheck()
	}
}

// 보드상태를 보여준다
func (g *game) ShowBoard() {
	for row := 0; row < g.gameBoard.ySize; row++ {
		rowLine := g.gameBoard.getRowLine(row)
		for _, state := range rowLine {
			tile := "□"
			if state == FILLED {
				tile = "■"
			} else if state == BLOCKED {
				tile = "X"
			}
			fmt.Print(tile)
		}
		fmt.Println("")
	}
}

// 확인용 기능
func (g *game) ShowGameStatus() {
	rowCasesCounter := 0
	columnCasesCounter := 0
	for _, cases := range g.rowCases {
		rowCasesCounter += len(cases)
	}

	for _, cases := range g.columnCases {
		columnCasesCounter += len(cases)
	}

	fmt.Println("남은 rowCases:", rowCasesCounter)
	fmt.Println("남은 colCases:", columnCasesCounter)

}
