package main

func CreateSpiral(n int) [][]int {
	// assert n is in range
	if n < 1 {
		return [][]int{}
	}

	// define result set
	result := make([][]int, n)
	for i:=0; i<n;i++{
		result[i]=make([]int, n)
	}

	// keep current value
	val := 1

	// begin with moving right
	xMod, yMod := 1, 0
	// from top left corner
	x, y := 0, 0
	// set initial value
	result[y][x]=val
	val++

	// add first line
	for i:=0; i<n-1;i++ {
		x+=xMod
		y+=yMod
		result[y][x] = val
		val++
	}
	if n == 1 {
		return result
	}

	// proceed vertically
	// length will decrement every other time
	for l:=n-1;l>0; l-- {
		// ensure decrementing every OTHER time
		for c:=0;c<2;c++ {
			if xMod==1 { // currently moving right
				// start moving down
				xMod,yMod=0,1
			} else if yMod==1 { // currently moving down
				// start moving left
				xMod, yMod = -1, 0
			} else if xMod==-1 { // currently moving left
				// start moving up
				xMod,yMod=0,-1
			} else if yMod==-1 { // currently moving up
				// start moving right
				xMod,yMod=1,0
			}

			// actually add each line
			for i:=0; i<l;i++ {
				x+=xMod
				y+=yMod
				result[y][x] = val
				val++
			}
		}
	}

	return result
}
