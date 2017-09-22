package datatypes

const (
	SORTED_SET_HEIGHT int = 32
)

type SkipList struct {
	levels [SORTED_SET_HEIGHT]ListNode
}

func (sl *SkipList) insert(v int)  {

}