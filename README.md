# Fullstack App

This repository contains a **full-stack** application, built with **Go** for the **API layer** and **React** for **Interactive layer**. This makes deployment and maintenance straightforward.

## Table of Contents

-   [Project Overview](#project-overview)
-   [Installation](#installation)
-   [Running the Application](#running-the-application)
    -   [Development Mode](#development-mode)
    -   [Production Build](#production-build)
-   [Commands](#commands)
    -   [Install Dependencies](#install-dependencies)
    -   [Build the Application](#build-the-application)
    -   [Run the Application](#run-the-application)
    -   [Test the Application](#test-the-application)
    -   [Swagger Documentation](#swagger-documentation)
-   [Contributing](#contributing)
-   [License](#license)

## Project Overview

This project integrates a **Go** backend with an **embedded React client**. The React client is bundled and served directly by the Go application, which simplifies deployment. The backend provides an API, which is handled through Go’s Fiber package, while the React frontend handles the UI/UX.

-   **API**: Go (with Fiber)
-   **Client**: Go (with embedded React)

### Features

-   Serve both API and frontend from a single binary.
-   Integrated Swagger API documentation.
-   Development server with hot reload.
-   Cross-platform support (Linux, macOS, Windows).

## Installation

Ensure you have **Go** and **Node.js** installed on your machine. If you don't have them, you can follow the installation guides for Go and Node.js.

1. Clone the repository:

    ```bash
    git clone <repository-url>
    cd <project-directory>
    ```

2. Install **Go** & **Node** dependencies:

    ```bash
    make deps
    ```

## Running the Application

### Development Mode

In development mode, the backend API and the frontend React app runs concurrently. The React app will be served on a development server, and the Go backend will handle API requests and serve the built React app.

1. Start the development environment:

    ```bash
    make dev
    ```

    This will:

    - Build the Go API and generate Swagger documentation.
    - Keep watching the react app for changes and keep building it in realtime.
    - Start the Go backend with hot reloading using `air`.

### Production Build

In production, a single executable is built, containing both the backend and the bundled React client.

1. Build for all platforms:

    ```bash
    make build
    ```

2. Run the generated binary:

    ```bash
    make run
    ```

This will start the API and serve the React app.

## Commands

### Install Dependencies

Install Go and Node dependencies:

```bash
make deps
```

### Build the Application

Build the Go API and the React client:

```bash
make build
```

### Run the Application

Run the Go API and serve the React client:

```bash
make run
```

### Test the Application

Run tests for the Go API:

```bash
make test
```

### Swagger Documentation

Generate Swagger documentation:

```bash
make swagger
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
