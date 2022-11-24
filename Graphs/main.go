package main

import (
	"fmt"
)

// 1st - Implement the structure of the graph and vertex
// Graph represents and adjacency list graph
// Will needs a list that holds pointers of vertices in a slice

type Graph struct {
	vertices []*Vertex
}

// Vertex represents a vertex graph
// Holds a key
// bc it's an adjacency list, each vertex in the graph will have a list of vertices that is holding the vertex of the neighbors

type Vertex struct {
	key      int
	adjacent []*Vertex
}

// 2nd - Implement the addVertex op and addEdge op
// addVertex adds a vertex to the graph
// It's a method to the Graph, so, it's a pointer receiver, and it takes a key
// I will create a vertex inside the body that has k as the key, then it will be appended to the vertex list in the graph
// Test it here. Better safe, than sorry =P
// Disallow duplicate keys and test it again

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("\n Sorry, vertex %v already exists. You cannot add it again. =( \n", k)
		fmt.Print(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// addEdge

// 3rd - Implement a func that gets a vertex
// getVertex

// 4th - Implement Contains
// Contains checks if the keys are already in the vertices of the graph
// Takes a list and keys and returns true if that list has that vertex and key value
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacency list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v: ", v.key) //prints the key of each vertex
		//loop for printing out other loops inside the adjacency list
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func main() {

	test := &Graph{}

	//adding nodes to it

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.Print()

	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)

}