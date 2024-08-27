package interfaces

type Saver interface {
	Save() error
	// SaveToFile(fileName string, content string) error
}

type Displayer interface {
	ToString()
}

type OutputAndSave interface {
	Saver

	//ToString and Displayer - any method of creation can be used
	// Displayer
	ToString()
}