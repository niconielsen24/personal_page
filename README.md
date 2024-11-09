
# TicTacToe Server

A simple TicTacToe server written in Go, implementing a game backend that handles game initialization and moves via API endpoints. The server is designed to be used with a React frontend, allowing users to play TicTacToe games in real-time.

## Features

- **Game Initialization**: Create a new TicTacToe game.
- **Make a Move**: Make a move in an ongoing game by providing a `game_id` and position.
- **CORS Support**: The server supports Cross-Origin Resource Sharing (CORS) to allow frontend communication.

## API Endpoints

### `POST /initGame`

Initializes a new TicTacToe game.

**Response**: 
- `200 OK`: Returns the initial game state.
- `500 Internal Server Error`: If there was an issue initializing the game.

### `PUT /makeMove`

Makes a move in the game identified by `game_id`.

**Request Body**:
```json
{
  "game_id": "uuid",
  "position": {
    "x": 0,
    "y": 0
  }
}
```

**Response**:
- `200 OK`: Returns the updated game state.
- `400 Bad Request`: If the move is invalid.
- `404 Not Found`: If the game with the specified `game_id` doesn't exist.
- `500 Internal Server Error`: If there was an issue processing the move.

## Frontend

The project uses **React** for the frontend. The frontend communicates with the Go server via the exposed API endpoints, allowing users to play the TicTacToe game interactively.

## Getting Started

### Prerequisites

- Go 1.18+ for the backend server
- Node.js and npm/yarn for the frontend (React)

### Backend Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/niconielsen24/personal_page.git
   cd personal_page
   ```

2. Install dependencies and run the server:
   ```bash
   go run main.go
   ```

3. The server will start at `http://localhost:8080`.

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Run the frontend:
   ```bash
   npm start
   ```

4. The frontend will be accessible at `http://localhost:3000`.

## License

This project is licensed under the MIT License - see the [LICENSE](https://www.mit.edu/~amini/LICENSE.md) file for details.
