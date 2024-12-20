package domain

type Pokemon struct {
	Name  string   `njson:"name"`
	Types []string `njson:"types.#.type.name"`
}
