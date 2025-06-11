# ZCalendar Backend

This directory contains the backend API server for the ZCalendar application.

## Technology Stack

- **Runtime**: Node.js 18+
- **Framework**: Express.js
- **Database**: PostgreSQL with Prisma ORM
- **Authentication**: JWT tokens
- **Caching**: Redis
- **Testing**: Jest
- **Code Quality**: ESLint + Prettier

## Getting Started

1. Install dependencies:
   ```bash
   npm install
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your database and other settings
   ```

3. Run database migrations:
   ```bash
   npm run migrate:dev
   ```

4. Start development server:
   ```bash
   npm run dev
   ```

## Project Structure

```
backend/
├── src/
│   ├── controllers/     # Request handlers
│   ├── services/        # Business logic
│   ├── models/          # Database models
│   ├── routes/          # API routes
│   ├── middleware/      # Custom middleware
│   ├── utils/           # Utility functions
│   └── config/          # Configuration files
├── tests/               # Test files
├── migrations/          # Database migrations
├── seeds/               # Database seed files
└── docs/                # API documentation
```

## Available Scripts

- `npm run dev` - Start development server with hot reload
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run test` - Run test suite
- `npm run test:watch` - Run tests in watch mode
- `npm run lint` - Check code style
- `npm run migrate:dev` - Run database migrations

## API Endpoints

The API documentation is available at `/api/docs` when running the server.

Key endpoints:
- `POST /api/v1/auth/login` - User authentication
- `GET /api/v1/events` - Get user events
- `POST /api/v1/events` - Create new event
- `PUT /api/v1/events/:id` - Update event
- `DELETE /api/v1/events/:id` - Delete event

## Environment Variables

Required environment variables:
- `DATABASE_URL` - PostgreSQL connection string
- `REDIS_URL` - Redis connection string
- `JWT_SECRET` - Secret for JWT token signing
- `PORT` - Server port (default: 3000)

See `.env.example` for a complete list.

## Contributing

1. Follow the existing code style
2. Write tests for new features
3. Update documentation as needed
4. Run linting before committing

For detailed development guidelines, see the main project documentation.