# Hacker News-like GraphQL API

This is a Hacker News-like GraphQL API built with **Golang**, **GORM**, **gqlgen**, and **Chi**. It provides a simple yet powerful API for creating and managing users and links.

## Features

- **User Creation**: Users can register an account.
- **User Login**: Users can log in to their accounts.
- **Link Creation**: Authenticated users can create links.
- **Built-in Authentication**: The API includes a JWT-based authentication system.
- **Query Links**: Users can query links either by user or retrieve all links.

## Technologies Used

- **Golang**: The programming language used for the backend.
- **GORM**: An ORM library for Golang that handles database operations.
- **gqlgen**: A library for building GraphQL servers in Go.
- **Chi**: A lightweight router for Go.

## Getting Started

To get a local copy up and running, follow these steps.

### Prerequisites

- Go (version 1.16 or later)
- PostgreSQL (or another supported database)


## Environment Variables

To run the server, you need to configure environment variables. Copy the `.env.example` file from the root of the project and rename it to `.env`. Update the values as needed:

```bash
cp .env.example .env
```

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Anuolu-2020/hackernews-api-clone.git
   cd your-repo
   ```

2. Install the required dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your database and update the connection string in the configuration.

4. Run the application:

   ```bash
   go run server.go
   ```

## API Endpoints

### User

- **Register User**: `POST /query`
- **Login User**: `POST /query`

### Links

- **Create Link**: `POST /query`
- **Query Links**: `POST /query`

## Example Queries

### Register User

```graphql
mutation create {
	createUser(input: { username: "exampleUser", password: "123" })
}
```

### Login User

```graphql
mutation create {
	login(input: { username: "exampleUser", password: "123" })
}

```

### Create Link

```graphql
mutation create {
	createLink(input: { title: "real link!", address: "www.graphql.org" }) {
		user {
			name
		}
	}
}

```

### Query Links

```graphql
query {
	links {
		title
		address
		user {
			name
		}
	}
}

```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Golang](https://golang.org/)
- [GORM](https://gorm.io/)
- [gqlgen](https://gqlgen.com/)
- [Chi](https://github.com/go-chi/chi)


