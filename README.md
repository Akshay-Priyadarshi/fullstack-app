# Fullstack App

This repository contains a **full-stack** application, built with [**Go**](https://go.dev/) for the **API layer** and [**React**](https://react.dev/) for **Client layer**.

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

This repository contains a **full-stack** application, built with [**Go**](https://go.dev/) for the **API layer** and [**React**](https://react.dev/) for **Client layer**. This makes deployment and maintenance straightforward. This repository also follows the **best practices** for Go and React applications. Apart from that this follows the [**Golang standards for project layout**](https://github.com/golang-standards/project-layout). Treat this as a starting point and further extend it to suit your needs.

### Features & Highlights

-   **Go API**:

    -   [**Fiber**](https://gofiber.io/) - Web framework.
    -   [**Air**](https://github.com/air-verse/air) - Live reload for Go apps.
    -   [**Swagger**](https://swagger.io/specification/) - API documentation.
    -   [**Make**](https://www.gnu.org/software/make/manual/make.html) - Build automation.

-   **React Client**:

    -   [**Vite**](https://vite.dev/) - Frontend build tool.
    -   [**React Router**](https://reactrouter.com/) - Routing.
    -   [**Tailwind CSS**](https://tailwindcss.com/) - Utility-first CSS framework.
    -   [**Axios**](https://axios-http.com/) - HTTP client.
    -   [**Tanstack Query**](https://tanstack.com/query/v3/) - Data fetching library.
    -   [**Storybook**](https://storybook.js.org/) - For showcasing components.
    -   [**EsBuild**](https://esbuild.github.io/) - For production builds.
    -   [**ShadCN UI**](https://ui.shadcn.com/) -Component library.
    -   [**Vitest**](https://vitest.dev/) - Testing framework.
    -   [**Cypress**](https://www.cypress.io/) - End-to-end testing.
    -   [**pnpm**](https://pnpm.io/) - Package manager.

-   **Fullstack**

    -   **Concurrent Development** - Run both the API and the React app concurrently.
    -   **Production Build** - Build a single executable for the API and the React app.

## Installation

Ensure you have **Go** and [**Node.js**](https://nodejs.org/en) installed on your machine. If you don't have them, you can follow the installation guides for [Go Installation](https://go.dev/doc/install) and [Node.js Installation](https://nodejs.org/en/download/package-manager).

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

Run tests for the Go API and the React client:

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
