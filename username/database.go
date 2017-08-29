package username

type Database struct {
	names map[string]int
}

func NewDatabase(existingNames []string) *Database {
	names := make(map[string]int)

	for _, name := range existingNames {
		counter, ok := names[name]
		if !ok {
			counter = 0
		}
		names[name] = counter + 1
	}

	return &Database{names}
}