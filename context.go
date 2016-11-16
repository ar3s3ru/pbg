package pbg

const (
	// Context key per i dati delle risposte negli handler API
	APIDataKey = "apiData"
	// Context key per gli errori delle risposte negli handler API
	APIErrorKey = "apiError"

	// Context key per accede al logger negli handler
	// Disponibile con l'adapter Server.WithLogger()
	LoggerKey = "logger"
	// Context key per accedere alla sessione autenticata negli handler
	// Usare un decoratore per controllare il token d'autorizzazione nelle richieste HTTP
	SessionKey = "session"

	// Context key per accedere all'interfaccia del DB delle mosse
	// Disponibile col decoratore Server.WithMoveDBAccess()
	MoveDBInterfaceKey = "moveDBInterface"
	// Context key per accedere all'interfaccia del DB dei pokèmon
	// Disponibile col decoratore Server.WithPokèmonDBAccess()
	PokèmonDBInterfaceKey = "pokèmonDBInterface"
	// Context key per accedere all'interfaccia del DB degli allenatori
	// Disponibile col decoratore Server.WithTrainerDBAccess()
	TrainerDBInterfaceKey = "trainerDBInterface"
	// Context key per accedere all'interfaccia col DB delle sessioni del server
	// Disponibile col decoratore Server.WithSessionDBAccess()
	SessionDBInterfaceKey = "sessionDBInterface"
)
