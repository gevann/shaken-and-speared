# shaken-and-speared

This repository contains a scaffold for a weekly word game. Each player receives the same set of letters for the week. Players may join a group to compete against one another. The game provides a *word to beat* and each player submits a word built from their remaining letters to score more points than the others.

## Structure

 - `backend/` – Go HTTP server exposing placeholder API endpoints.

## Getting Started

1. Build the server (requires Go):

   ```bash
   cd backend
   go build -o wordgame
   ```

2. Start the server:

   ```bash
   ./wordgame
   ```

   The backend uses SQLite for persistent storage. A database file `wordgame.db`
   will be created automatically in the working directory when the server starts.
   Ensure the process has write access to this directory.

The server runs on port `3000` by default and exposes the following placeholder endpoints:

- `GET /api/status` – health check.
- `GET /api/game/week` – retrieve the week's letters.
- `POST /api/game/word` – submit a word.

These endpoints currently return mock data and should be expanded with real game logic.
