package ggraph

// GraphNodeIndexType node map index type
type GraphNodeIndexType = uint32

// GraphLinkLayerType link layer index type
type GraphLinkLayerType = uint32

// GraphLinkIndexType link map index type
type GraphLinkIndexType = uint64

// GraphLinkWeightType default weight type
type GraphLinkWeightType = float32

// IGraph interface for graph
type IGraph interface {
	// Build a new graph
	New() *Graph
	// Node returns a new node, while appending to the graph
	Node(index GraphNodeIndexType) Node
	// Layer returns a new layer within the name
	Layer(name string) Layer
	// Link returns a new link, while appending to the graph
	Link(layer string, index GraphLinkIndexType) Link

	// NextNodeIndex returns next node index when appending a node
	NextNodeIndex() GraphNodeIndexType
	// NextLinkIndex returns next link index when appending a link
	NextLinkIndex(layer string) GraphLinkIndexType

	// Nodes return node collection as a map
	// key index template-able
	Nodes() map[GraphNodeIndexType]*Node

	// Layers return layer collection as a map
	Layers() map[string]*Layer

	// Links return link collection as a map
	// key index template-able
	Links(layer string) map[string]map[GraphLinkIndexType]*Link
}

// Graph struct for a graph
type Graph struct {
	nodes map[GraphNodeIndexType]*Node
	links map[GraphLinkIndexType]*Link
}

// GraphComponent graph component
type GraphComponent struct {
	graph *Graph
}

// New impl from IGraph
func (g *Graph) New() *Graph {
	graph := new(Graph)
	graph.nodes = *new(map[GraphNodeIndexType]*Node)
	graph.links = *new(map[GraphLinkIndexType]*Link)
	return graph
}

// NextNodeIndex impl from IGraph
func (g Graph) NextNodeIndex() GraphNodeIndexType {
	cursor := GraphNodeIndexType(len(g.nodes) + 1)
	_, exists := g.nodes[cursor]
	for ; exists; _, exists = g.nodes[cursor] {
		cursor++
	}
	return cursor
}

// NextLinkIndex impl from IGraph
func (g Graph) NextLinkIndex(layer string) GraphLinkIndexType {
	cursor := GraphLinkIndexType(len(g.links) + 1)
	_, exists := g.links[cursor]
	for ; exists; _, exists = g.links[cursor] {
		cursor++
	}
	return cursor

}

// Node impl from IGraph
func (g Graph) Node() Node {
	ni := g.NextNodeIndex()
	node := new(Node)
	node.graph = &g
	node.id = ni
	g.nodes[ni] = node
	return *node
}

// Link impl from IGraph
func (g Graph) Link(layer string) Link {
	li := g.NextLinkIndex(layer)
	link := new(Link)
	link.graph = &g
	link.id = li
	g.links[li] = link
	return *link
}
