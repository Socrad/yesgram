package yesgram

import (
	"github.com/Socrad/mathlogic"
)

/*
	한줄의 타일의 상태에 대한 모든 케이스를 얻는다.

칠해진 타일 사이에는 1개의 막힌 블록이 있고 양 옆에는 0개의 막힌 블록이 있는 프로토 타입을 생성한 후
막힌 블록의 길이의 합이 가용여백의 크기인 모든 경우를 생성하여 각 경우를 프로토 타입의 막힌 블록에 더해주면
모든 케이스를 생성할 수 있다.
*/
func getLineStateCases(hint []int, lineSize int) [][]int {
	lineStateCases := [][]int{}
	lineCasePrototype, freeSpace := getLineCasePrototype(hint, lineSize)

	n := len(lineCasePrototype.blockedTiles)
	total := freeSpace
	nNumbersTotalCases := mathlogic.GetNnumbersTotalCases(n, total) // n개 숫자의 합이 total이 되는 모든 경우를 얻는다.
	for _, aCase := range nNumbersTotalCases {
		aLineCase := lineCase{}
		aLineCase.filledTiles = lineCasePrototype.filledTiles

		for index, value := range aCase {
			aLineCase.blockedTiles = append(aLineCase.blockedTiles, lineCasePrototype.blockedTiles[index]+value)
		}
		lineStateCases = append(lineStateCases, getPaintedState(aLineCase))
	}
	return lineStateCases
}

/*
매 경우의 블록타일의 길이와 채워진 타일의 길이를 저장하기 위한 구조체
*/
type lineCase struct {
	blockedTiles []int
	filledTiles  []int
}

// 경우 생성에 사용될 프로토타입과 가변여백의 크기를 얻는다.
func getLineCasePrototype(hint []int, lineSize int) (lineCase, int) {
	totalLength := 0
	lineCasePrototype := lineCase{}

	lineCasePrototype.blockedTiles = append(lineCasePrototype.blockedTiles, 0)
	for _, length := range hint {
		lineCasePrototype.filledTiles = append(lineCasePrototype.filledTiles, length)
		totalLength += length
		lineCasePrototype.blockedTiles = append(lineCasePrototype.blockedTiles, 1)
		totalLength++
	}
	lineCasePrototype.blockedTiles[len(hint)] = 0
	totalLength--

	freeSpace := lineSize - totalLength

	return lineCasePrototype, freeSpace
}

// lineCase대로 타일을 칠하고 막아 한 경우의 줄 상태를 얻는다.
func getPaintedState(aCase lineCase) []int {
	lineState := []int{}
	for index, length := range aCase.filledTiles {
		for i := 0; i < aCase.blockedTiles[index]; i++ {
			lineState = append(lineState, BLOCKED)
		}
		for i := 0; i < length; i++ {
			lineState = append(lineState, FILLED)
		}
	}
	for i := 0; i < aCase.blockedTiles[len(aCase.filledTiles)]; i++ { // blockedTiles 는 filledTiles 보다 1개 많다.
		lineState = append(lineState, BLOCKED)
	}
	return lineState
}
