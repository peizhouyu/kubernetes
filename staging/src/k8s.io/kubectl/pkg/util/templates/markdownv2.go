package templates

const linebreak2 = "\n"

// ASCIIRenderer implements blackfriday.Renderer
//var _ blackfriday.Renderer = &ASCIIRendererV2{}

// ASCIIRenderer is a blackfriday.Renderer intended for rendering markdown
// documents as plain text, well suited for human reading on terminals.
type ASCIIRendererV2 struct {
	Indentation string

	listItemCount uint
	listLevel     uint
}
