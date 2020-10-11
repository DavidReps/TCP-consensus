import "math/rand"

type node struct{
	state float32
	ID int
}


func newNode(id int) *node{
	n := node{state: rand.Float32(), ID: id}
	return &n
}

func nodeMaker(n int) []*node{
	var nodes []*node
	for i := 0; i < n; i++{
		nodes = append(nodes, newNode(i+1))
	}
	return nodes
}
