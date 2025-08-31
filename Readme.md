Here’s a polished **README.md** based on your steps, formatted and cleaned up for clarity:

---

# GraphQL Service

This repository provides a GraphQL service built with **gqlgen**. It follows a domain-driven folder structure where resolvers, handlers, DTOs, and use cases are separated by responsibility.

---

## 🚀 How to add a new domain

### 1. Define schema
Create a new file inside `graph/schema` with your domain name, e.g.:

```bash
graph/schema/todo.graphqls
```

Inside it, define the request/response types and extend `Query` or `Mutation`:

```graphql
type Todo {
  id: ID!
  text: String!
  done: Boolean!
}

extend type Query {
  todos: [Todo!]!
}

extend type Mutation {
  createTodo(text: String!): Todo!
}
```

---

### 2. Generate code
Run:

```bash
make gen
```

This will:
- Update `internal/dto` with generated types
- Create or update `<domain>.resolvers.go` inside `graph/resolver/`

---

### 3. Add a handler
Create a new handler in the `handler/` folder for your domain. Example:

```
handler/
  rest.go
```

---

### 4. Expose handler methods
In your domain file (e.g., `todo.go`), implement the required methods and expose them.

---

### 5. Wire up the resolver
- Check `graph/resolver/resolver.go` → ensure the new handler is added to the `Resolver` struct.
- Update `cmd/app/main.go` to initialize the handler and inject it into the resolver.

Example in `graph/resolver/resolver.go`:

```go
type Resolver struct {
    TodoHandler *handler.TodoHandler
    UserHandler *handler.UserHandler
}
```

And in `main.go`:

```go
srv := handler.NewDefaultServer(
    generated.NewExecutableSchema(
        generated.Config{
            Resolvers: &graph.Resolver{
                TodoHandler: handler.NewTodoHandler(),
            },
        },
    ),
)
```

---

### 6. Implement resolver logic
Go to your generated resolver file, e.g.:

```
graph/resolver/todo.resolvers.go
```

And invoke the corresponding handler methods inside the resolver functions.

---

## 📂 Project Structure

```
graph/
    generated/ # Auto-generated files
    schema/                 # GraphQL schema files (.graphqls)
        domain.graphqls     # Domain-specific schema files
    resolver/               # Auto-generated resolver stubs
        resolver.go         # Root Resolver struct
        domain.resolver.go  # Domain-specific resolvers

internal/
    dto/       # Auto-generated GraphQL DTOs
    domain/
        handler/            # Handlers per domain
        usecase/            # Business logic layer
        repository/         # Data access / external APIs
        domain.go           # collection of interfaces
    models/                 # Domain models
    config/                 # Application configuration
    utils/                  # Helper functions
    pkg/                    # Dependency and 3rd party invocation packages
cmd/app/main.go     # Application entry point
```

---

## 🛠️ Development

- **Add new domain:** create `<domain>.graphqls` under `graph/schema`
- **Generate code:** run `make gen`
- **Implement logic:** in `handler/` + resolver files
- **Wire dependencies:** in `resolver.go` and `main.go`

---

✅ That’s it! The rest (business logic and data access) should follow clean separation of **usecase** (domain logic) and **repository** (DB/API integration).
