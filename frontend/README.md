# Sharingan Frontend

The Sharingan Frontend is a Next.js application that provides a user interface for the Sharingan application monitoring system.

## Getting Started

### Prerequisites

- Node.js 18+
- npm or yarn

### Development

```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Start production server
npm run start

# Run linting
npm run lint
```

## Project Structure

```
frontend/
├── public/           # Static files
├── src/              # Source code
│   ├── app/          # Next.js app directory
│   │   ├── api/      # API routes
│   │   │   └── auth/ # Authentication endpoints
│   │   └── components/ # React components
│   ├── pages/        # Page components
│   └── styles/       # CSS styles
├── next.config.ts    # Next.js configuration
└── tsconfig.json     # TypeScript configuration
```

## Key Features

- **Authentication**: Integration with Keycloak for secure user authentication
- **Application Dashboard**: Monitor and manage applications
- **Incident Management**: Track and respond to application incidents
- **User Subscriptions**: Subscribe to application alerts
- **Responsive UI**: Mobile-friendly interface

## Environment Configuration

Create a `.env.local` file in the frontend directory with the following variables:

```
NEXT_PUBLIC_API_URL=http://localhost:8000
NEXT_PUBLIC_KEYCLOAK_URL=http://localhost:8080
NEXT_PUBLIC_KEYCLOAK_REALM=sharingan
NEXT_PUBLIC_KEYCLOAK_CLIENT_ID=frontend
```

## Authentication Flow

The frontend uses Next.js with NextAuth.js for authentication, integrating with Keycloak as the identity provider. The authentication flow is illustrated in the main project README.