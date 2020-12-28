package main

type LRUCache struct {
	list map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	key int
	val int
}

func main() {

}
