package anchor

// AnchorMap - contains map of anchors
type AnchorMap struct {
	// anchorMap - for each anchor key in the patterns it will maintain information if the key exists in the resource
	// if anchor key of the pattern exists in the resource then (key)=true else (key)=false
	anchorMap map[string]bool
	// AnchorError - used in validate to break execution of the recursion when if condition fails
	AnchorError validateAnchorError
}

// NewAnchorMap -initialize anchorMap
func NewAnchorMap() *AnchorMap {
	return &AnchorMap{anchorMap: map[string]bool{}}
}

// KeysAreMissing - if any of the anchor key doesn't exists in the resource then it will return true
// if any of (key)=false then return KeysAreMissing() as true
// if all the keys exists in the pattern exists in resource then return KeysAreMissing() as false
func (ac *AnchorMap) KeysAreMissing() bool {
	for _, v := range ac.anchorMap {
		if !v {
			return true
		}
	}
	return false
}

// CheckAnchorInResource checks if condition anchor key has values
func (ac *AnchorMap) CheckAnchorInResource(pattern map[string]interface{}, resource interface{}) {
	for key := range pattern {
		if a := Parse(key); IsCondition(a) || IsExistence(a) || IsNegation(a) {
			val, ok := ac.anchorMap[key]
			if !ok {
				ac.anchorMap[key] = false
			} else if !val {
				if resourceHasValueForKey(resource, a.Key()) {
					ac.anchorMap[key] = true
				}
			}
		}
	}
}
