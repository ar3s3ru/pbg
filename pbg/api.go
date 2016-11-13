package pbg

import (
    "encoding/json"
    "github.com/valyala/fasthttp"
)

type (
    // APIResponser è un'interfaccia che si occupa di tradurre le risposte dai
    // RequestHandler di fasthttp (ovvero, le funzioni che rendono il server un middleware)
    // nel formato scelto per la comunicazione Client<->Server (nel rispetto dei principi REST)
    APIResponser interface {
        // Prende in input lo status code della risposta HTTP, eventuali dati in caso di successo e
        // l'errore in caso di fallimento
        //
        // Restituisce la risposta tradotta pronta da essere mandata al client
        Writer(statusCode int, data interface{}, err error) []byte

        // Ritorna l'header HTTP Content-Type che definisce il tipo di risposta del server
        ContentType() []byte
    }

    // APIResponser che si occupa di tradurre le risposte del server in formato JSON
    // e le invia usando le strutture JSONSuccess e JSONError
    JSONResponser func(statusCode int, data interface{}, err error) []byte

    // Rappresenta la struttura JSON di un messaggio di successo
    JSONSuccess struct {
        Data interface{} `json:"data"`
    }

    // Rappresenta la struttura JSON di un messaggio d'errore
    JSONError struct {
        // Presentazione testuale dell'errore HTTP
        Error string `json:"error"`
        // Chiarisce ulteriormente le cause dell'errore
        Message string `json:"message"`
    }
)

var (
    // HTTP header utilizzato per denotare i contenuti JSON
    // Verrà inserito nel contesto della risposta HTTP
    jsonContentType = []byte("application/json; charset=utf-8")

    // In caso di errori nell'unmarshalling JSON, ritorna questo errore generico
    typeAssertError = []byte(`{"error":"Internal Server Error","message":"Some error occurred"}`)
)

// Adapter non esportato per il server
// Serve a tradurre i dati della risposta HTTP secondo lo standard di comunicazione
// scelto per parlare col client, secondo i principi REST
func (srv *server) apiWriter(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        // Esegui prima la richiesta
        // Al termine della richiesta, SI SPERA che l'utente abbia reso disponibile
        // i risultati tramite contesto in APIErrorKey (errore) o APIDataKey (successo)
        handler(ctx)

        var err error
        if apiErr := ctx.UserValue(APIErrorKey); apiErr != nil {
            // C'è stato un errore, a quanto pare...
            var ok bool
            if err, ok = apiErr.(error); !ok {
                // ...ma l'errore non è del tipo error, come dovrebbe essere
                ctx.SetStatusCode(fasthttp.StatusInternalServerError)
                ctx.SetBody(typeAssertError)
                return
            }
        }

        // Nessun errore, traduci i dati e scrivi nel contesto della risposta
        ctx.Response.SetBody(srv.apiResponser.Writer(
            ctx.Response.StatusCode(),
            ctx.UserValue(APIDataKey),
            err,
        ))

        // Segnala al client il tipo della risposta
        ctx.SetContentTypeBytes(srv.apiResponser.ContentType())
    }
}

// Costruisce un nuovo oggetto di tipo JSONResponser da passare
// al factory method del Server
func NewJSONResponser() JSONResponser {
    // In verità, questo APIResponser è la stessa funzione Writer che vorremmo
    // avere (vedere l'implementazione di JSONResponser.Writer() per capire)
    return func(statusCode int, data interface{}, err error) []byte {
        var (
            response []byte // Body della risposta
            errInt error    // Errore interno, può capitare durante il marshalling in JSON
        )

        if data != nil {
            // La richiesta ha avuto successo
            dataJson := JSONSuccess{data}
            response, errInt = json.Marshal(dataJson)
        } else if err != nil {
            // La richiesta non ha avuto successo
            errJson := JSONError{fasthttp.StatusMessage(statusCode), err.Error()}
            response, errInt = json.Marshal(errJson)
        } else {
            // TODO: insert error here (se i due casi precedenti non sono validi)
        }

        if errInt != nil {
            // Abbiamo avuto un errore interno, in particolare durante il mashalling in JSON
            // Quindi utilizziamo una risposta JSON "hardcoded" per descrivere l'errore
            return []byte(`{"error":"` + fasthttp.StatusMessage(statusCode) + `","message":"` + errInt.Error() + `"}`)
        }

        return response
    }
}


func (jr JSONResponser) Writer(statusCode int, data interface{}, err error) []byte {
    // Chiama la funzione di JSONResponser sui dati passati al Writer
    return jr(statusCode, data, err)
}

func (jr JSONResponser) ContentType() []byte {
    return jsonContentType
}

// Aggiungi l'entità 'ent' al contesto della risposta 'ctx' con la chiave key,
// inoltre setta lo statusCode passato come argomento
func writeAPI(ctx *fasthttp.RequestCtx, key string, ent interface{}, stat int) {
    ctx.SetUserValue(key, ent)
    ctx.SetStatusCode(stat)
}

// Aggiungi i dati al contesto della risposta HTTP
func WriteAPISuccess(ctx *fasthttp.RequestCtx, data interface{}, statusCode int) {
    writeAPI(ctx, APIDataKey, data, statusCode)
}

// Aggiungi l'errore dell'handler al contesto della risposta HTTP
func WriteAPIError(ctx *fasthttp.RequestCtx, err error, statusCode int) {
    writeAPI(ctx, APIErrorKey, err, statusCode)
}
