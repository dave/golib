package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"unsafe"
)

func asNode(n *types.Node) *Node      { return (*Node)(unsafe.Pointer(n)) }
func asTypesNode(n *Node) *types.Node { return (*types.Node)(unsafe.Pointer(n)) }
