package astar

type Point struct {
	X int32
	Y int32
}

func createUint8Array(x, y int32) [][]uint8 {
	array := make([][]uint8, x)
	for i := range array {
		array[i] = make([]uint8, y)
	}
	return array
}

func createInt32Array(x, y int32) [][]int32 {
	array := make([][]int32, x)
	for i := range array {
		array[i] = make([]int32, y)
	}
	return array
}
