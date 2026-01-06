# go-iam

[![codecov](https://codecov.io/github/melvinodsa/go-iam/graph/badge.svg?token=TWJXNBHTQL)](https://codecov.io/github/melvinodsa/go-iam)
[![Go Report Card](https://goreportcard.com/badge/github.com/melvinodsa/go-iam)](https://goreportcard.com/report/github.com/melvinodsa/go-iam)
[![GoDoc](https://godoc.org/github.com/melvinodsa/go-iam?status.svg)](https://godoc.org/github.com/melvinodsa/go-iam)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

**go-iam** is a lightweight, multi-tenant Identity and Access Management (IAM) server built in **Golang**. It provides robust authentication and fine-grained authorization for modern applications. With support for custom roles, third-party auth providers, and multi-client setups, `go-iam` gives you full control over access management in a scalable and modular way.

> ✅ Admin UI: [go-iam-ui](https://github.com/melvinodsa/go-iam-ui)  
> 🐳 Docker Setup: [go-iam-docker](https://github.com/melvinodsa/go-iam-docker)  
> 🔐 Backend: [go-iam](https://github.com/melvinodsa/go-iam)  
> 📦 SDK: [go-iam-sdk](https://github.com/melvinodsa/go-iam-sdk)  
> 🚀 Examples: [go-iam-examples](https://github.com/melvinodsa/go-iam-examples)
> 💬 Reddit Community: [r/GoIAM](https://www.reddit.com/r/GoIAM/)

<img src=".github/go-iam.png" alt="go-iam overview" width="400"/>

---

## ✨ Features

### 🔀 Multi-Tenancy

- Create and manage **Projects**
- Strict **isolation** of data between tenants

### 🔐 Authentication Provider Integration

- Google, Microsoft, GitHub OAuth login support
- Easily extendable to add more providers
- **Shared credentials** support across multiple clients

### 🧩 Client Management

- Multiple apps (clients) per project
- Avoid duplicate OAuth credentials

### 🧱 Role-Based Access Control (RBAC)

- Define resources and group them into roles
- Create custom roles and assign to users
- Granular access control for different actions/resources

### 🛠️ Admin UI

- React-based Admin interface for managing:
  - Projects
  - Users
  - Roles
  - Resources
  - Clients

---

## 🧰 Tech Stack

| Component     | Tech                |
| ------------- | ------------------- |
| Backend       | Golang              |
| Database      | MongoDB             |
| Caching (opt) | Redis               |
| Frontend      | React + Vite (PNPM) |

---

## 🏗️ Architecture Overview

### Provider Level

The `Provider` struct acts as the central dependency injection container, wiring together all core services, database connections, cache, and middleware. It is initialized at application startup and made available throughout the request lifecycle.

```mermaid
classDiagram
    class Provider {
        +Service S
        +DB D
        +CacheService C
        +ProjectMiddlewares PM
        +AuthMiddlewares AM
        +SdkClient AuthClient
    }

    class Service
    class DB
    class CacheService
    class ProjectMiddlewares
    class AuthMiddlewares
    class SdkClient

    Provider o-- Service : S
    Provider o-- DB : D
    Provider o-- CacheService : C
    Provider o-- ProjectMiddlewares : PM
    Provider o-- AuthMiddlewares : AM
    Provider o-- SdkClient : AuthClient
```

### Service Level

The `Service` struct is a container for all business logic services, each responsible for a specific domain (projects, clients, users, roles, etc.). Each service is constructed with its required dependencies, such as stores, other services, or utility providers.

```mermaid
classDiagram
    class Service {
        +ProjectService Projects
        +ClientService Clients
        +AuthProviderService AuthProviders
        +AuthService Auth
        +UserService User
        +ResourceService Resources
        +RoleService Role
        +PolicyService Policy
        +AuthSyncService AuthSync
    }

    class ProjectService
    class ClientService
    class AuthProviderService
    class AuthService
    class UserService
    class ResourceService
    class RoleService
    class PolicyService
    class AuthSyncService

    class ProjectStore
    class ClientStore
    class AuthProviderStore
    class ResourceStore
    class RoleStore
    class UserStore
    class PolicyStore
    class CacheService
    class JWTService
    class EncryptService

    Service o-- ProjectService : Projects
    Service o-- ClientService : Clients
    Service o-- AuthProviderService : AuthProviders
    Service o-- AuthService : Auth
    Service o-- UserService : User
    Service o-- ResourceService : Resources
    Service o-- RoleService : Role
    Service o-- PolicyService : Policy
    Service o-- AuthSyncService : AuthSync

    ProjectService <.. ProjectStore : uses
    ClientService <.. ClientStore : uses
    AuthProviderService <.. AuthProviderStore : uses
    AuthService <.. AuthProviderService : uses
    AuthService <.. ClientService : uses
    AuthService <.. CacheService : uses
    AuthService <.. JWTService : uses
    AuthService <.. EncryptService : uses
    AuthService <.. UserService : uses
    UserService <.. UserStore : uses
    UserService <.. RoleService : uses
    ResourceService <.. ResourceStore : uses
    RoleService <.. RoleStore : uses
    PolicyService <.. PolicyStore : uses
    AuthSyncService <.. AuthService : uses
```

### Key Concepts

- **Provider**: Central container for all dependencies (services, DB, cache, middleware).
- **Service**: Domain-specific business logic, constructed with required dependencies.
- **Dependency Injection**: All components are initialized and injected at startup for testability and modularity.
- **Middleware**: Project and authentication middlewares are initialized and attached to the Fiber app.

### Usage

- Initialize the provider using your application config:
  ```go
  provider, err := providers.InjectDefaultProviders(config)
  if err != nil {
      log.Fatal(err)
  }
  ```
- Attach the provider to your Fiber app:
  ```go
  app.Use(providers.Handle(provider))
  ```
- Access services and dependencies in handlers via:
  ```go
  p := providers.GetProviders(c)
  userService := p.S.User
  ```

For more details, see the code in the `providers/` directory and the service implementations in `services/`.

---

## 🚀 Getting Started

### Option 1: 🔧 Manual Setup (Development)

#### Prerequisites

- Go 1.21+
- MongoDB
- Redis (optional, recommended)
- Google OAuth Credentials

#### Run the Backend

```bash
git clone https://github.com/melvinodsa/go-iam.git
cd go-iam
cp sample.env .env
go run main.go
```

### Option 2: 🐳 Docker-Based Local Setup (Recommended for Testing)

Use the official go-iam-docker repo to spin up everything with Docker Compose, including:

- MongoDB
- Redis
- go-iam (backend)
- go-iam-ui (admin frontend)

#### Steps

```bash
git clone https://github.com/melvinodsa/go-iam-docker.git
cd go-iam-docker
cp sample.env .env
docker compose up -d
```

#### Access

- Admin UI: [http://localhost:4173](http://localhost:4173)
- API: [http://localhost:3000](http://localhost:3000)
- API Docs: [http://localhost:3000/docs](http://localhost:3000/docs)

---

## 🧪 Testing

### Running Tests

Due to the extensive test suite in this project, it's important to use the correct testing command to avoid test caching issues that might cause local tests to pass while CI/CD fails.

**Use this command for reliable testing:**

```bash
go test -count=1 ./... -v --race -cover
```

**Flags explanation:**

- `-count=1`: Disables test result caching to ensure fresh test runs
- `./...`: Runs tests for all packages recursively
- `-v`: Verbose output showing individual test results
- `--race`: Enables race condition detection
- `-cover`: Shows test coverage information

**Why `-count=1` is important:**

- With many tests, Go may cache results and show false positives locally
- CI/CD environments don't use cached results, leading to inconsistencies
- This flag ensures your local testing matches CI/CD behavior

### Running Specific Test Suites

```bash
# Test specific package
go test -count=1 ./services/user -v --race -cover

# Test specific function
go test -count=1 ./services/user -v --race -cover -run TestCopyUserResources
```

---

## 📦 Environment Variables

Some important environment variables used in `.env`:

| Variable                                       | Description                                                           |
| ---------------------------------------------- | --------------------------------------------------------------------- |
| `LOGGER_LEVEL`                                 | Logger level `1 - Debug` (refer., `https://docs.gofiber.io/api/log/`) |
| `DB_HOST`                                      | MongoDB URI (e.g., `mongodb://user:pass@host/db`)                     |
| `JWT_SECRET`                                   | Secret key used for generating and verifying JWT tokens               |
| `REDIS_HOST`, `REDIS_PASSWORD`, `ENABLE_REDIS` | Redis host address and toggle to enable Redis caching                 |
| `ENCRYPTER_KEY`                                | Optional symmetric key for encrypting sensitive fields - change this  |
| `AUTH_PROVIDER_REFETCH_INTERVAL_IN_MINUTES`    | Interval in minutes to refetch and sync third-party auth providers    |
| `TOKEN_CACHE_TTL_IN_MINUTES`                   | Interval for which the authentication token should be valid           |

---

## License

- Community Edition: [Apache 2.0](./LICENSE) (Open Source, free to use)
