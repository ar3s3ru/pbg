package mem

import (
    "bytes"
    "encoding/base64"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

var (
    authorizationPrefix = []byte("Basic ")
    authorizationComma  = []byte(":")
    authorizationPassw  = []byte("x")
)

// Controlla la validità dell'header Authorization nelle richieste HTTP
// Deve essere del tipo "(token):x", dove token è in formato UUID
// e l'header in Base64.
//
// Ritorna il valore del token in []byte in caso di successo, o un errore
// nel caso, appunto, d'errore
func BasicAuthorization(header []byte) ([]byte, error) {
    if bytes.HasPrefix(header, authorizationPrefix) {
        // Decode from base64 to UTF-8 string
        payload, err := base64.StdEncoding.DecodeString(string(header[len(authorizationPrefix):]))
        if err == nil {
            // Splitting the payload string into "token:password"
            pair := bytes.SplitN(payload, authorizationComma, 2)
            // Checking password validation
            if len(pair) == 2 && bytes.Equal(pair[1], authorizationPassw) {
                // Returns Session token
                return pair[0], nil
            }
        }
    }
    // Some error occurred, that means we had an invalid Authorization header
    return nil, pbg.ErrInvalidAuthorizationHeader
}

// Converte un file JSON in due dataset, rispettivamente, per i Pokèmon e per le mosse
// Ritorna un errore nel caso di problemi con la conversione; i primi campi saranno settati su nil
func WithDatasetFile(file string) ([]pbg.Pokèmon, []pbg.Move, error) {
    sf, err := marshalSourceFile(file)
    if err != nil {
        return nil, nil,err
    }

    pokèmons := convertLtoPL(sf.PList)
    moves    := convertLtoML(sf.MList)

    return pokèmons, moves, nil
}

// Converte uno slice di tipo mem.move (non esportato) in uno slice
// di tipo pbg.Move (interfaccia)
//
// Operazione d'ordine theta(len(moves))
func convertLtoML(moves []move) []pbg.Move {
    if moves == nil {
        panic("Must use a valid move list!")
    }

    list := make([]pbg.Move, len(moves), len(moves))
    for i := range moves {
        list[i] = pbg.Move(&moves[i])
    }

    return list
}

// Converte uno slice di tipo mem.pokèmon (non esportato) in uno slice
// di tipo pbg.Pokèmon (interfaccia)
//
// Operazione d'ordine theta(len(pokèmons))
func convertLtoPL(pokèmons []pokèmon) []pbg.Pokèmon {
    if pokèmons == nil {
        panic("Must use a valid pokèmon list!")
    }

    list := make([]pbg.Pokèmon, len(pokèmons), len(pokèmons))
    for i := range pokèmons {
        list[i] = pbg.Pokèmon(&pokèmons[i])
    }

    return list
}

// Controlla che l'indice i sia nell'intervallo 0 < i < (upperBoud + 1)
func inRange(i int, upperBound int) bool {
    return i >= 1 && i <= upperBound
}

// Controlla la validità delle IVs passate per argomento
// Ricordiamo che le IVs vanno da 0 a 31, per ogni IVs
func checkIvs(ivs [6]int) bool {
    sum := 0
    for _, v := range ivs {
        if v < 0 || v > 31 {
            return false
        }

        sum += v
    }

    if sum < 0 || sum > (6 * 31) {
        return false
    }

    return true
}

// Controlla la validità delle EVs passate come argomento
// Ricordiamo che le EVs hanno una somma totale di 510 e che ogni EV
// ha un valore massimo di 255
func checkEvs(evs [6]int) bool {
    sum := 0
    for _, v := range evs {
        if v < 0 || v > 255 {
            return false
        }

        sum += v
    }

    if sum < 0 || sum > 510 {
        return false
    }

    return true
}

// Controlla la validità del livello passato come argomento
// Ricordiamo che il livello di un pokemon si trova nell'intervallo [1, 100]
func checkLevel(level int) bool {
    return level >= 1 && level <= 100
}
