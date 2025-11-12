package node

type Choice struct {
	ID       ChoiceID
	Text     string
	Effects  []Effect
	ToNodeID NodeID
}
