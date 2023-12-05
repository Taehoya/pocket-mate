package entities

type Bound int

const (
	SpringBound Bound = iota
	BasicBound
)

type Note struct {
	Bound      Bound
	NoteColor  string
	BoundColor string
}

var (
	defaultColour = "#22659A"
)

func NewNote(bound Bound, noteColourParam, bounColourParam string) *Note {
	if noteColourParam == "" {
		noteColourParam = defaultColour
	}

	if bounColourParam == "" {
		bounColourParam = defaultColour
	}

	return &Note{
		Bound:      bound,
		NoteColor:  noteColourParam,
		BoundColor: bounColourParam,
	}
}

func (b Bound) String() string {
	switch b {
	case SpringBound:
		return "SpringNote"
	case BasicBound:
		return "BasicNote"
	default:
		return "BasicNote"
	}
}
