package mem

import (
    "io/ioutil"
    "encoding/json"
)

type sourceFile struct {
    Generation int       `json:"generation"`
    PNumbers   int       `json:"pokemon_count"`
    MNumbers   int       `json:"move_count"`
    PList      []pokèmon `json:"pokemons"`
    MList      []move    `json:"moves"`
}

func marshalSourceFile(file string) (*sourceFile, error) {
    if file == "" {
        return nil, ErrInvalidDataSourceFile
    } else if file, err := ioutil.ReadFile(file); err != nil {
        return nil, err
    } else {
        source := sourceFile{}
        if err := json.Unmarshal(file, &source); err != nil {
            return nil, err
        }

        return &source, nil
    }
}
