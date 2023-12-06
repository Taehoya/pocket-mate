package entities

type NoteType string

type Note struct {
	NoteType   NoteType
	NoteColor  string
	BoundColor string
}

var (
	defaultColour = "#22659A"
)

func NewNote(noteType NoteType, noteColourParam, bounColourParam string) *Note {
	if noteColourParam == "" {
		noteColourParam = defaultColour
	}

	if bounColourParam == "" {
		bounColourParam = defaultColour
	}

	return &Note{
		NoteType:   noteType,
		NoteColor:  noteColourParam,
		BoundColor: bounColourParam,
	}
}
