package literals

type Digital struct {
	value string
}

func (d *Digital) compare(other Digital) int {
	if len(d.value) < len(other.value) {
		return -1;
	} else if len(d.value) > len(other.value) {
		return 1;
	} else {
		for i := 0; i < len(d.value); i++ {
			if uint(d.value[i]) < uint(other.value[i]) {
				return -1
			} else if uint(d.value[i]) > uint(other.value[i]) {
				return 1
			}
		}
	}
	return 0;
}
