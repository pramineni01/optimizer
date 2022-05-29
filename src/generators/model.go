package generators


interface Model {
	// func Print(format: string) ([]byte, error)
	// func ToBytes(out: *[]bytes) error
	// func Save(format: string, loc: string): (bool, error)
}

interface ModelGenerator {
	func generate() model
}