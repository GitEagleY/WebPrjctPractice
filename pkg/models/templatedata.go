package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSRFtoken string //cross site request forgery token
	Flash     string
	Warnaing  string
	Error     error
}
