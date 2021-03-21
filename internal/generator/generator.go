package generator

// Generator interface
type Generator interface {
	RenderHTMLFile(templateFile string, sel string) (string, error)
}
