# PokèmonBattleGo

Pokèmon battles powered by **Go**!
*(Note: I know, Pokèmon is spelled with "é" but I'm too lazy to correct it, anyway...)*

## Architecture

All the main server logic is contained inside the `github.com/ar3s3ru/PokemonBattleGo/pbgServer` package.

To download the package, use the command:

```
    go get github.com/ar3s3ru/PokemonBattleGo/pbgServer
```

### Go? *Programming by interfaces!*

Embracing the *second* most important feature of Go language, the actual server logic is **decoupled** using *3 different interfaces*, that is:

```go
    IDataMechanism interface {
        // Handles data models accesses, creation and removal
    }

    ISessionMechanism interface {
        // Handles user session mechanisms
    }

    IAuthMechanism interface {
        // Implements login/logout, register/unregister functions
    }
```

The main reason is to abstract from data storage solutions (*in-memory database storage, like __Redis__*, databases like *__PostgreSQL__* or *__MongoDB__*), session creation and retrieval (like *direct connection to the server using HTTP(S) protocol*, or *using a SaaS (i.e. __Firebase__, __GCM__, ...)* to handle sensible information delivery), and different authorization handlers.

### Modeling with OOP? No thanks, *interfaces are just fine...*

Actually, even *__data models__ are interfaces*, here...

Main reason is: **WRITE YOUR OWN MODELS!**

Blindly trust user initializations, given a third-party data structure, is **foolish**.

Remember the golden rule:
> _never, **EVER**, trust the final user_.

So, interfaces **are just fine**.

From [Interfaces][go-interfaces-ref] on *Effective Go*:

> Interfaces in Go provide a way to specify the behavior of an object: 
> if something can do *this*, then it can be used *here*.


It's so much easier, and painless, thinking to *how data models behave* rather than *what their structure should be*.

#### TODO

    * Finish this README.md

> *Danilo Cianfrone*

[go-interfaces-ref]: https://golang.org/doc/effective_go.html#interfaces_and_types