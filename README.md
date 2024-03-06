Package to solve the nonograms (= picross = griddlers)

Example main.go


	package main
	
	import "github.com/Socrad/yesgram"
	
	func main() {

		rowHints := [][]int{{2, 2}, {2, 3}, {3}, {3, 5}, {1, 9, 1}, {1, 13}, {7, 5}, {3, 3}, {2, 1, 3}, {2, 3}, {2, 2}, {1, 2, 2}, {1, 2, 2, 2}, {2, 11}, {5, 5}}
		columnHints := [][]int{{2, 3}, {2, 1, 2}, {3, 2}, {1, 5, 1}, {7, 2}, {4, 4}, {3, 3}, {3, 1}, {3, 1, 3}, {3, 5}, {7, 2, 2}, {9, 2}, {3, 6, 2}, {2, 6}, {3, 5}}
	 
		game := yesgram.NewGame(rowHints, columnHints)
		game.Solve()
		game.ShowBoard()
	}
