package main

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
	visited  bool
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		key:      key,
		adjacent: []*Vertex{},
	}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) AddEdge(v1, v2 *Vertex) {
	v1.adjacent = append(v1.adjacent, v2)
	v2.adjacent = append(v2.adjacent, v1)
}

func (g *Graph) DFS(v *Vertex) {
	v.visited = true
	fmt.Printf("Visited vertex: %d\n", v.key)

	for _, adjacent := range v.adjacent {
		if !adjacent.visited {
			g.DFS(adjacent)
		}
	}
}

func main() {
	g := &Graph{}

	vertices := []*Vertex{}
	for i := 0; i < 8; i++ {
		vertex := NewVertex(i)
		vertices = append(vertices, vertex)
		g.AddVertex(vertex)
	}

	g.AddEdge(vertices[0], vertices[1])
	g.AddEdge(vertices[0], vertices[2])
	g.AddEdge(vertices[1], vertices[3])
	g.AddEdge(vertices[1], vertices[4])
	g.AddEdge(vertices[2], vertices[5])
	g.AddEdge(vertices[2], vertices[6])
	g.AddEdge(vertices[3], vertices[7])

	g.DFS(vertices[0])
}
