package yesgram

import "slices"

/*
행과 열을 차례로 읽으며 현재 상태와 상충되는 케이스를 제거하고, 케이스들을 conjunction하여 칠해지거나 막힌게 확실한 타일을 채운다.
변경한 사항이 없으면 false를 리턴한다
*/
func (g *game) conjunctionCheck() bool {
	isChanged := false
	for row, rowlineCases := range g.rowCases {
		originLine := g.gameBoard.getRowLine(row)
		conformedCases := getConformedCases(originLine, rowlineCases)

		lineState := getLineStateCasesConjunction(conformedCases)
		lineState = getLineStateDisjunction(originLine, lineState)

		if slices.Compare(originLine, lineState) != 0 {
			isChanged = true
			g.gameBoard.setRowLine(lineState, row)
		}
		g.rowCases[row] = conformedCases
	}

	for column, columnlineCases := range g.columnCases {
		originLine := g.gameBoard.getColumnLine(column)
		conformedCases := getConformedCases(originLine, columnlineCases)

		lineState := getLineStateCasesConjunction(conformedCases)
		lineState = getLineStateDisjunction(originLine, lineState)

		if slices.Compare(originLine, lineState) != 0 {
			isChanged = true
			g.gameBoard.setColumnLine(lineState, column)
		}
		g.columnCases[column] = conformedCases
	}

	return isChanged
}

// 현재 lineState와 상충되는 경우를 제거한 나머지 케이스를 얻는다
func getConformedCases(lineState []int, lineStateCases [][]int) [][]int {
	conformedCases := [][]int{}
	for _, lineStateCase := range lineStateCases {
		isMatched := true
		for index, state := range lineStateCase {
			if lineState[index]*state == -1 { // 어느 한쪽이 FILLED 인데 다른쪽이 BLOCKED 라면 곱하면 -1이 된다
				isMatched = false
				break
			}
		}
		if isMatched {
			conformedCases = append(conformedCases, lineStateCase)
		}
	}
	return conformedCases
}

// 모든 케이스에 공통으로 겹치는 타일상태의 결과를 얻는다.
func getLineStateCasesConjunction(lineStateCases [][]int) []int {
	conjunction := slices.Clone(lineStateCases[0])
	for _, lineStateCase := range lineStateCases {
		for index, state := range lineStateCase {
			if conjunction[index] != state {
				conjunction[index] = NEUTRAL
			}
		}
	}
	return conjunction
}

// 두 라인스테이트의 논리합적 결과를 얻는다.
func getLineStateDisjunction(a, b []int) []int {
	result := slices.Clone(a)
	for index, state := range result {
		if state == 0 {
			result[index] = b[index]
		}
	}

	return result
}
