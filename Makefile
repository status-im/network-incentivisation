generate:
	mkdir -p golang/nodes
	abigen --sol contracts/Nodes.sol --pkg nodes --out golang/nodes/nodes.go
	abigen --sol contracts/NodesV2.sol --pkg nodes --out golang/nodes/nodesv2.go

