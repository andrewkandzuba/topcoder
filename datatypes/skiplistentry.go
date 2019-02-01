package datatypes

type SkipListEntry struct {
	value uint
	next  *SkipListEntry
	down  *SkipListEntry
}

func NewSkipListEntry(value uint, next *SkipListEntry, down *SkipListEntry) *SkipListEntry {
	return &SkipListEntry{value: value, next: next, down: down}
}

func NewSkipListEntryDefault() *SkipListEntry {
	return &SkipListEntry{value: 0, next: nil, down: nil}
}

func (ls SkipListEntry) Value() uint {
	return ls.value
}

func (ls SkipListEntry) Next() *SkipListEntry {
	return ls.next
}

func (ls SkipListEntry) Down() *SkipListEntry {
	return ls.down
}
