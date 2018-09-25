package datatypes

const (
	SKIPLIST_HEIGHT = 32
)

type SkipList struct {
	levels [SKIPLIST_HEIGHT]SkipListEntry
}

func (sl *SkipList) insert(v int)  {
	/*for l := range sl.levels {

	}*/
}