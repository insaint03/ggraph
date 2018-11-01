package ggraph

// Link interfaced
type Link struct {
	GraphComponent
	id GraphLinkIndexType

	Source Node
	Target Node
	Layer  string
	Weight GraphLinkWeightType
}
