package converter

// nestedStruct is container for structs that found inside
// an array or another structs
type nestedStruct struct {
	List []string
	Map  map[string]struct{}
}

func (ns *nestedStruct) add(structDecl string) {
	if _, exist := ns.Map[structDecl]; exist {
		return
	}

	ns.Map[structDecl] = struct{}{}
	ns.List = append(ns.List, structDecl)
}
