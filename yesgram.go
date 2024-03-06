package yesgram

import "fmt"

type game struct {
	rowHints    [][]int
	columnHints [][]int
	rowCases    [][][]int
	columnCases [][][]int
	gameBoard   board
}

/*
	노노그램

노노그램을 생성할 때, 힌트가 칠해질 수 있는 모든 케이스를 생성해서 함께 저장
*/
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

모든 행에 대해 현재 상태와 상충되는 케이스를 제거 후, 남은 케이스들에서 공통으로 칠해지거나 막히는 타일을 칠하고 막는다.
모든 열에 대해서도 같은 작업을 수행한다.
열에 대해 작업을 수행하는 과정에서 행의 상태가 바뀌었다면 남은 케이스 중 행의 상태와 상충되는 케이스가 있을 수 있다.
그러므로 행에 대해 다시 같은 작업을 수행하면 행의 상태가 바뀔 수 있다.
열에 대해서도 마찬가지다.
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
