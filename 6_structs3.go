package main
import "fmt"
type rect struct {
	width int
	height int
}
func (r rect) area() int {
	return r.width * r.height
}

func main(){
	r := rect{width: 12, height: 9}
	fmt.Println(r.area())
}