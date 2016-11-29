# PokèmonBattleGo

Pokèmon battles powered by **Go**!

*(Note: I know, Pokèmon is spelled with "é" but I'm too lazy to correct it, anyway...)*

## Architecture

All the main server logic is contained inside the `github.com/ar3s3ru/pbg` package.

To download the package, use the command:

```
    go get github.com/ar3s3ru/pbg
```

### Component-based architecture

Since we'd like to abstract from concrete implementation of data storage (hence, usage of one DB instead of another, or just using in-heap memory
objects), I choose to structure the whole system as **independent components**, following the principles behind *CBSE*, 
or *Component-based software engineering*.

After identifying key roles inside the system (at analysis stage), which are:
* **Server HTTP** (for middlewares registration)
* **PokèmonDB**
* **MoveDB**
* **TrainerDB**
* **SessionDB**

I proceed with writing interfaces that actually *decouples* the server logic, such as:

```go
    // Middleware registration
    ServerHTTP interface { ... }

    // Data storage
    PokèmonDBComponent interface { ... }
    MoveDBComponent    interface { ... }
    TrainerDBComponent interface { ... }
    SessionDBComponent interface { ... }
```

Every DB software-component allows **CRUD operations** through a **component-required interface**:

```go
    PokèmonDBInterface interface { ... }
    MoveDBInterface    interface { ... }
    TrainerDBInterface interface { ... }
    SessionDBInterface interface { ... } 
```

Such interfaces are provided by their own component (following *ObjectPool semantic*) and can be requested inside an HTTP handler through
a Server-level decorator (see *ServerAdapter* interfaces for reference); the Server will make such interfaces accessible through the
*fasthttp.RequestCtx* object of the HTTP handler.

### REST Architecture

The executable server should follow the REST Architecture.
As such, models representations at module level and models representations at middleware level **should be different**.

Meaning:

```go
    // At framework level we could have this interface
    Model interface {
        Name() string
        Dog()  Dog     // Dog is another model interface
    }

    // At module level we could have this model
    mem.Model struct {
        name string
        dog  Dog
    }

    func (m *mem.Model) Name() string { return m.name }
    func (m *mem.Model) Dog() Dog { return m.dog }

    // At middleware level, we would like to implement HATEOAS principle from REST architecture, and thus...
    RESTDog struct {
        Dog             // Embedding
    } 

    RESTModel struct {
        Model
    }

    func (rd *RESTDog) MarshalJSON() ([]byte, error) {
        /**
            Do marshalling here with this structure:

            { 
                "name": <dog name>, 
                "href": "http(s)://server:port/dog/102"  // HATEOAS, state of this dog is available at that link
            }
         */
    }

    func (rm *RESTModel) MarshalJSON() ([]byte, error) {
        // Create temporary RESTDog value that holds the original Dog interface value
        dog := &RESTDog{rm.Dog()}

        // Do marshalling here...
    }
```

Moreover, we can use a *Negotiator-like* handler that will output API data with requested encoding/format.
For this purpouse, we have an **APIResponser** interface (look at the documentation for further references).

## TODO

Use cases implementated as far are:

* Retrieve Pokèmon and Moves registered and available (permanently) on the server
* Create and setup new Pokèmon trainer profiles

Must still implement:

* Matchmaking mechanisms
* Database migration (considering MongoDB as of today)

## Contributing

As a **GDG Pisa community project**, everyone can contribute as they like.
All the code here is open source (*duh!*), without license (so free to use and modify).

License could be used in further commits, though.

Anyway, feel free to do *significant* pull request when you like :)


* - Danilo Cianfrone*