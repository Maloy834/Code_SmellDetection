package main

type sumation struct {
	total int
}
func (s* sumation)sum(x  int,y int,z int) {
   // s := sumation{}
	s.total =x+y+z
   println(s.total)
}
func main() {
	s := sumation{}
	col:=complexityVisitor{}
	col.Complexity=5
	col.Count=3
	s.sum(1, 2, 3)

}