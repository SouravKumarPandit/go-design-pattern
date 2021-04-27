package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var colors = []string{"Red", "Green", "Blue", "White", "Black"}

func main() {

	fmt.Println(colors)
	factory := NewShapeFactory()
	for i := 0; i < 20; i++ {

		circle := factory.getCircle(getRandomColor(colors...))
		circle.x = getRandomInt(50, 100)
		circle.y = getRandomInt(30, 100)
		circle.radius = getRandomInt(5, 100)
		circle.draw()
	}
}

func getRandomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)

}
func getRandomColor(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
