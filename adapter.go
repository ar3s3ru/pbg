package pbg

import "github.com/valyala/fasthttp"

type (
	// Adapter è un decoratore, ovvero una funzione che prende in input
	// una funzione e ne restituisce un'altra appartenente alla stessa famiglia;
	// i decoratori servono ad estendere le funzionalità di routine pre-esistenti.
	//
	// In questo caso, Adapter è un decoratore per i RequestHandler del package fasthttp.
	Adapter func(fasthttp.RequestHandler) fasthttp.RequestHandler
)

// Applica un insieme di adapter ad un RequestHandler.
// Gli adapter vengonono applicati in ordine.
func Adapt(rh fasthttp.RequestHandler, adapters ...Adapter) fasthttp.RequestHandler {
	for _, adapter := range adapters {
		// rh diventa un RequestHandler con funzionalità estesa da adapter
		rh = adapter(rh)
	}

	return rh
}
