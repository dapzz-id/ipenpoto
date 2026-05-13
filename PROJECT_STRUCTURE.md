# 📁 IPENPOTO API - Project Structure

```
ipenpoto/
├── app/
│   ├── controllers/          # HTTP request handlers
│   │   └── AuthController.go
│   │
│   ├── middleware/           # Request/Response middleware
│   │   ├── JWTMiddleware.go
│   │   ├── RoleMiddleware.go
│   │   ├── RateLimitMiddleware.go
│   │   └── SecurityLogger.go
│   │
│   ├── services/             # Business logic layer
│   │   ├── AuthService.go
│   │   ├── JWTService.go
│   │   └── TokenBlacklistService.go
│   │
│   ├── repositories/         # Data access layer
│   │   └── UserRepository.go
│   │
│   ├── enums/                # Enumeration types
│   │   ├── ApprovalStatusEnum.go
│   │   ├── CorporateMemberEnum.go
│   │   ├── StatusUserEnum.go
│   │   └── UserEnum.go
│   │
│   ├── requests/             # Request DTOs
│   │   └── AuthRequest.go
│   │
│   ├── responses/            # Response DTOs
│   │   └── AuthResponse.go
│   │
│   ├── utils/                # Utility functions
│   │   ├── Response.go
│   │   └── Validation.go
│   │
│   ├── helpers/              # ✨ Helper functions (NEW)
│   │   ├── ContainerHelpers.go
│   │   ├── StringHelpers.go
│   │   ├── RegexHelpers.go
│   │   └── ConversionHelpers.go
│   │
│   └── tests/                # ✨ Test files (NEW)
│       ├── unit/             # Unit tests
│       │   ├── AuthServiceTest.go
│       │   └── JWTServiceTest.go
│       │
│       └── helpers/          # Test helpers
│           └── TestHelpers.go
│
├── config/                   # Configuration
│   ├── database.go
│   └── redis.go
│
├── database/
│   ├── models/               # Data models
│   │   ├── User.go
│   │   ├── CorporateMember.go
│   │   ├── CorporateProfile.go
│   │   ├── PhotographerProfile.go
│   │   ├── RegistrationRequest.go
│   │   └── UserRoleCorporateRequest.go
│   │
│   └── seeders/              # Database seeders
│
├── routes/
│   └── api.go                # API route definitions
│
├── nginx/                    # Nginx configuration
│   ├── nginx.conf
│   └── sites/
│       └── default.conf
│
├── main.go                   # Application entry point
├── go.mod                    # Go dependencies
├── go.sum                    # Go dependency checksums
├── docker-compose.yml        # Docker Compose configuration
├── Dockerfile.dev            # Development Dockerfile
│
└── docs/                     # Documentation
    ├── SECURITY_AUDIT.md     # Security improvements
    └── SECURITY_CHECKLIST.md # Comprehensive checklist

```

---

## 📂 Folder Purposes

### `app/controllers/`
HTTP request handlers that receive client requests and return responses.

### `app/middleware/`
Cross-cutting concerns:
- `JWTMiddleware.go` - Token validation
- `RoleMiddleware.go` - Role-based access control
- `RateLimitMiddleware.go` - Rate limiting
- `SecurityLogger.go` - Security event logging

### `app/services/`
Business logic layer containing domain logic and services.

### `app/repositories/`
Data access layer for database operations.

### `app/enums/`
Type-safe enumeration definitions.

### `app/requests/`
Request Data Transfer Objects (DTOs).

### `app/responses/`
Response Data Transfer Objects (DTOs).

### `app/utils/`
Utility functions for common operations.

### `app/helpers/` ✨ NEW
Helper functions organized by type:
- **String helpers** - String operations
- **Regex helpers** - Regular expression utilities
- **Conversion helpers** - Type conversions

**Usage:**
```go
import "ipenpoto/app/helpers"

h := helpers.New()
h.String.TrimAndLower("  HELLO  ") // "hello"
h.Regex.IsValidEmail("test@example.com")
h.Conversion.StringToInt("123")
```

### `app/tests/` ✨ NEW
Test files organized by type:
- **`unit/`** - Unit tests
- **`helpers/`** - Test helper functions

**Run tests:**
```bash
go test ./app/tests/unit -v
go test ./... -v
```

### `config/`
Configuration and initialization:
- Database connection
- Redis connection
- Environment variables

### `database/models/`
GORM model definitions for database tables.

### `database/seeders/`
Database seeding scripts for initial data.

### `routes/`
API route definitions and middleware chain setup.

### `nginx/`
Nginx reverse proxy configuration.

---

## 🏗️ Architecture Pattern

```
┌─────────────────┐
│   HTTP Client   │
└────────┬────────┘
         │
    ┌────▼──────────┐
    │   Middleware  │
    │ (Security Log)│
    └────┬──────────┘
         │
    ┌────▼──────────┐
    │  Controllers  │ ◄─── Receive requests, validate input
    └────┬──────────┘
         │
    ┌────▼──────────┐
    │   Services    │ ◄─── Business logic, rules
    └────┬──────────┘
         │
    ┌────▼──────────┐
    │ Repositories  │ ◄─── Database queries
    └────┬──────────┘
         │
    ┌────▼──────────┐
    │   Database    │
    │  (PostgreSQL) │
    └────┬──────────┘
         │
    ┌────▼──────────┐
    │   Utilities   │ ◄─── Response, Validation
    │   Helpers     │ ◄─── String, Regex, Conversion
    └───────────────┘
```

---

## 📦 Dependencies

### Runtime Dependencies
- `gorm.io/gorm` - ORM
- `github.com/gofiber/fiber/v2` - Web framework
- `github.com/golang-jwt/jwt/v5` - JWT
- `github.com/redis/go-redis/v9` - Redis client
- `golang.org/x/crypto` - Cryptography
- `github.com/google/uuid` - UUID generation
- `github.com/go-playground/validator/v10` - Validation

### Development Dependencies
- Go 1.26.2+
- Docker & Docker Compose
- PostgreSQL 18.3
- Redis 8.6

---

## 🚀 Quick Start

### Setup
```bash
# 1. Clone and install dependencies
go mod download

# 2. Setup environment
cp .env.example .env
# Edit .env with your values

# 3. Run migrations
go run main.go --migrate

# 4. Start development server
go run main.go
```

### Testing
```bash
# Run all tests
go test ./... -v

# Run unit tests only
go test ./app/tests/unit -v

# With coverage
go test ./... -cover
```

### Docker
```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f app-api

# Stop services
docker-compose down
```

---

## 📝 Naming Conventions

### Files
- `snake_case` for filenames
- `_test.go` suffix for test files
- Descriptive names (e.g., `auth_service.go`, not `service.go`)

### Packages
- `lowercase` package names
- Match folder names

### Types & Functions
- `PascalCase` for exported identifiers
- `camelCase` for unexported identifiers

### Variables
- `camelCase` for all variables
- Meaningful names (avoid `a`, `b`, `x`)

---

## 🔐 Security Structure

All security-related components:
- `middleware/JWTMiddleware.go` - Token validation
- `middleware/RoleMiddleware.go` - Authorization
- `middleware/RateLimitMiddleware.go` - Rate limiting
- `middleware/SecurityLogger.go` - Logging
- `services/TokenBlacklistService.go` - Token blacklist
- `utils/Validation.go` - Input validation

See `SECURITY_AUDIT.md` for detailed information.

---

## 📊 Lines of Code

```
app/
├── controllers/        ~100 LOC
├── middleware/         ~300 LOC ✨
├── services/           ~200 LOC ✨
├── repositories/       ~50 LOC
├── enums/             ~100 LOC
├── requests/          ~50 LOC
├── responses/         ~50 LOC
├── utils/             ~150 LOC
├── helpers/           ~200 LOC ✨
└── tests/             ~300 LOC ✨
    ├── unit/          ~200 LOC
    └── helpers/       ~100 LOC
├── config/            ~100 LOC
├── database/          ~150 LOC
└── routes/            ~50 LOC

TOTAL: ~1,800+ LOC with comments & tests
```

---

## ✨ Recent Improvements

- ✅ Reorganized tests into `app/tests/unit/`
- ✅ Created `app/helpers/` for helper functions
- ✅ Added test helpers in `app/tests/helpers/`
- ✅ Removed unnecessary files
- ✅ Enterprise-grade security implementation

---

## 🎯 Next Steps

1. **Use Helpers** - Import and use helpers in controllers/services
2. **Write Tests** - Add more unit tests to `app/tests/unit/`
3. **Integration Tests** - Add integration tests (DB, Redis)
4. **API Documentation** - Generate Swagger/OpenAPI docs
5. **Monitoring** - Add logging and metrics

