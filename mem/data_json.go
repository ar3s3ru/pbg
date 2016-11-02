package mem

type sourceFile struct {
    Generation int       `json:"generation"`
    PNumbers   int       `json:"pokemon_count"`
    MNumbers   int       `json:"move_count"`
    PList      []pokèmon `json:"pokemons"`
    MList      []move    `json:"moves"`
}
