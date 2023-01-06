package main

type sand point

func (s sand) canGoDown(obstacles map[point]bool) bool {
	temp := s.goDown()
	return !obstacles[point(temp)]
}

func (s sand) goDown() sand {
	s.y++
	return s
}

func (s sand) canGoDownLeft(obstacles map[point]bool) bool {
	temp := s.goDownLeft()
	return !obstacles[point(temp)]
}

func (s sand) goDownLeft() sand {
	temp := s.goDown()
	temp.x--
	return temp
}

func (s sand) canGoDownRight(obstacles map[point]bool) bool {
	temp := s.goDownRight()
	return !obstacles[point(temp)]
}

func (s sand) goDownRight() sand {
	temp := s.goDown()
	temp.x++
	return temp
}

func createSandStartPoint() sand {
	return sand{
		x: 500,
		y: 0,
	}
}
