package username

type Database struct {
	names map[string]int
}

func NewDatabase(existingNames []string) *Database {
	database := new(Database)
	database.names = make(map[string]int)

	for _, name := range existingNames {
		counter, ok := database.names[name]
		if (!ok) {
			counter = 0
		}
		database.names[name] = counter + 1
	}

	return database
}