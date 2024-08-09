# API with WebSockets using Go and Chi

This project implements a RESTful API in Go with support for WebSockets for real-time notifications. It uses the Chi library for routing and Gorilla WebSocket for WebSocket communication.

## Project Structure

- `apiHandler`: Structure that manages the API handlers, WebSocket connections, and client subscriptions.
- `NewHandler`: Function that initializes Chi routing and middleware, including routes for room and message management.
- `handleSubscribe`: Handler that manages WebSocket connections for a specific room.
- `handleCreateRoom`: Handler that creates a new room.
- `handleCreateRoomMessage`: Handler that creates a new message in a room and notifies connected clients.

## Dependencies

- [Chi](https://github.com/go-chi/chi) for HTTP routing.
- [Gorilla WebSocket](https://github.com/gorilla/websocket) for WebSocket communication.
- [PGX](https://github.com/jackc/pgx) for PostgreSQL interaction.
- [slog](https://pkg.go.dev/log/slog) for structured logging.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Andreoew/semana-tech-go-react-server
   cd semana-tech-go-react-server
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Configure your PostgreSQL database and adjust the settings as needed.

## Code Structure

### apiHandler

```go
type apiHandler struct {
    q           *pgstore.Queries
    r           *chi.Mux
    upgrader    websocket.Upgrader
    subscribers map[string]map[*websocket.Conn]context.CancelFunc
    mu          *sync.Mutex
}
```
