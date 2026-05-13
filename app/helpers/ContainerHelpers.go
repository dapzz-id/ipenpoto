package helpers

// Helpers aggregates all helper packages
type Helpers struct {
	String     StringHelpers
	Regex      RegexHelpers
	Conversion ConversionHelpers
}

// New creates a new Helpers instance with all sub-helpers initialized
func New() Helpers {
	return Helpers{
		String:     NewStringHelpers(),
		Regex:      NewRegexHelpers(),
		Conversion: NewConversionHelpers(),
	}
}
