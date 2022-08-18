package astar

import (
	"errors"
	"math"
)

type Astar struct {
	width  int32
	height int32
	points [][]uint8
	path   []Point
}

//NewAstar 新建寻路
func NewAstar(w, h int32) *Astar {
	return &Astar{
		width:  w,
		height: h,
		points: createUint8Array(w, h),
		path:   make([]Point, 0),
	}
}

//SetData 设置数据
func (a *Astar) SetData(data []byte) (err error) {
	if len(data) != int(a.width)*int(a.height) {
		err = errors.New("长度错误。")
		return
	}
	for x := 0; x < int(a.width); x++ {
		h := data[x*int(a.height) : x*int(a.height)+int(a.height)]
		for y := 0; y < int(a.height); y++ {
			a.points[x][y] = uint8(h[y])
		}
	}
	a.path = make([]Point, 0)
	return
}

//Find 寻找路径
func (a *Astar) Find(x1, y1, x2, y2 int32) (err error) {
	a.path = make([]Point, 0)
	if x1 < 0 || x1 > a.width-1 ||
		y1 < 0 || y1 > a.height-1 ||
		x2 < 0 || x2 > a.width-1 ||
		y2 < 0 || y2 > a.height-1 {
		err = errors.New("无效的点。")
		return
	}

	openList := make([]Point, 0)
	openListVal := createInt32Array(a.width, a.height)
	parentX := createInt32Array(a.width, a.height)
	parentY := createInt32Array(a.width, a.height)
	openListG := createInt32Array(a.width, a.height)
	openListH := createInt32Array(a.width, a.height)
	closeListVal := createInt32Array(a.width, a.height)

	historyF := int32(-1)
	tempPoint := Point{X: x1, Y: y1}
	openList = append(openList, tempPoint)

	var currentF, tempVal, currentX, currentY, tempX, tempY int32
	for openListG[x2][y2] == 0 {
		if len(openList) == 0 {
			err = errors.New("寻路错误。")
			return
		}

		for n := 0; n < len(openList); n++ {
			if n == 0 {
				historyF = openListH[openList[n].X][openList[n].Y] + openListG[openList[n].X][openList[n].Y]
				tempVal = int32(n)
				currentX = openList[n].X
				currentY = openList[n].Y
			} else {
				currentF = openListH[openList[n].X][openList[n].Y] + openListG[openList[n].X][openList[n].Y]
				if historyF >= currentF {
					historyF = currentF
					tempVal = int32(n)
					currentX = openList[n].X
					currentY = openList[n].Y
				}
			}
		}
		openList = append(openList[:tempVal], openList[tempVal+1:]...)
		openListVal[currentX][currentY] = 0
		tempPoint.X = currentX
		tempPoint.Y = currentY
		closeListVal[currentX][currentY] = 1
		xd := []int32{-1, 1, -1, 1, 0, -1, 1, 0}
		yd := []int32{-1, -1, 1, 1, -1, 0, 0, 1}
		for n := 0; n < 8; n++ {
			tempX = currentX + xd[n : n+1][0]
			tempY = currentY + yd[n : n+1][0]
			if tempX < 1 || tempY < 1 {
				//...
			} else if a.points[tempX][tempY] == 0 {
				//...
			} else if closeListVal[tempX][tempY] == 1 {
				//...
			} else {
				if openListVal[tempX][tempY] == 1 {
					tempVal = openListG[currentX][currentY]
					if n > 4 {
						tempVal += 10
					} else {
						tempVal += 14
					}
					if openListG[tempX][tempY] > tempVal {
						openListG[tempX][tempY] = openListG[currentX][currentY]
						if n > 4 {
							openListG[tempX][tempY] += 10
						} else {
							openListG[tempX][tempY] += 14
						}
						parentX[tempX][tempY] = currentX
						parentY[tempX][tempY] = currentY
					}
				} else {
					openListVal[tempX][tempY] = 1
					tempPoint.X = tempX
					tempPoint.Y = tempY
					openList = append(openList, tempPoint)

					parentX[tempX][tempY] = currentX
					parentY[tempX][tempY] = currentY
					openListG[tempX][tempY] = openListG[currentX][currentY]
					if n > 4 {
						openListG[tempX][tempY] += 10
					} else {
						openListG[tempX][tempY] += 14
					}
					openListH[tempX][tempY] = int32((math.Abs(float64(x2)-float64(tempX)) + math.Abs(float64(y2)-float64(tempY))) * 10)
				}
			}
		}
	}
	tempPoint.X = x2
	tempPoint.Y = y2
	a.path = append(a.path, tempPoint)

	tempPoint.X = parentX[x2][y2]
	tempPoint.Y = parentY[x2][y2]

	for !(tempPoint.X == 0 && tempPoint.Y == 0) {
		a.path = append(a.path, tempPoint)
		if tempPoint.X > 0 && tempPoint.Y > 0 {
			tempX = parentX[tempPoint.X][tempPoint.Y]
			tempY = parentY[tempPoint.X][tempPoint.Y]
			tempPoint.X = tempX
			tempPoint.Y = tempY
		}
	}
	return
}

//GetPath 获取路径
func (a *Astar) GetPath() []Point {
	return a.path
}

//CheckPoint 同行状态
func (a *Astar) CheckPoint(x, y int32) bool {
	if x < 0 || x > a.width-1 || y < 0 || y > a.height-1 {
		return false
	}
	return a.points[x][y] == 1
}
