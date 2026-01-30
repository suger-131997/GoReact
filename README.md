# Go + html/template + React

This is an experimental project to verify how to seamlessly integrate Go's standard `net/http` package, `html/template`, and **React**.
Designed as a proof-of-concept for a technical blog post, it explores methods to achieve a modern frontend development experience within a simple Go server architecture.

## Getting Started

### Prerequisites

- Go 1.25+
- Node.js (npm)
- [air](https://github.com/air-verse/air) (for live reloading)

### Installation

1. **Install Dependencies**
   ```bash
   go mod download
   npm install
   ```

2. **Start Development Server**
   ```bash
   make dev
   ```
   The server will be available at `http://localhost:3000`.

3. **Build and Run for Production**
   ```bash
   make run
   ```
