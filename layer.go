package ggraph

// Layer link layer
type Layer struct {
	GraphComponent
	id    GraphLinkLayerType
	Links map[GraphLinkIndexType]*Link
}
