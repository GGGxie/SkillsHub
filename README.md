# Skills Hub

A community-driven AI skills repository for discovering, creating, learning, and sharing. Built with Go (backend) and React (frontend), featuring bilingual support (English/Chinese) and Google OAuth login.

## Tech Stack

### Backend
- **Go** with [Gin](https://github.com/gin-gonic/gin) web framework
- **SQLite** database (zero-config, file-based)
- **JWT** authentication
- **Google OAuth2** login
- RESTful API design

### Frontend
- **React 18** + **TypeScript** + **Vite**
- **TailwindCSS** for styling
- **react-i18next** for internationalization (EN/ZH)
- **@react-oauth/google** for Google Sign-In
- **react-router-dom** for routing
- **lucide-react** for icons

## Features

- Skill discovery with search, sort (hottest/latest), and pagination
- Featured skills carousel
- Skill detail pages with Markdown rendering
- Comment system
- Like/vote system
- Article/learning section
- Help & FAQ page
- Submit new skills
- Google OAuth login
- Bilingual UI (English/Chinese) with language switcher
- Responsive design (mobile-friendly)
- Modern, clean UI inspired by skills-hub.cc

## Project Structure

```
SkillsHub/
├── backend/
│   ├── cmd/
│   │   └── main.go              # Entry point
│   ├── internal/
│   │   ├── config/config.go     # Configuration
│   │   ├── handler/
│   │   │   ├── auth.go          # Auth handlers (Google OAuth)
│   │   │   ├── skill.go         # Skill CRUD handlers
│   │   │   └── article.go       # Article handlers
│   │   ├── middleware/auth.go   # JWT middleware
│   │   ├── model/model.go       # Data models
│   │   └── service/db.go        # Database init & seed
│   ├── go.mod
│   └── .env.example
├── frontend/
│   ├── src/
│   │   ├── api/index.ts         # API client
│   │   ├── components/
│   │   │   ├── Layout.tsx       # App layout
│   │   │   ├── Header.tsx       # Navigation header
│   │   │   ├── Footer.tsx       # Footer
│   │   │   └── SkillCard.tsx    # Skill card component
│   │   ├── hooks/useAuth.tsx    # Auth context & hook
│   │   ├── i18n/
│   │   │   ├── index.ts         # i18n setup
│   │   │   ├── en.json          # English translations
│   │   │   └── zh.json          # Chinese translations
│   │   ├── pages/
│   │   │   ├── HomePage.tsx     # Main discover page
│   │   │   ├── SkillDetailPage.tsx
│   │   │   ├── AuthPage.tsx     # Login/signup
│   │   │   ├── AuthCallback.tsx
│   │   │   ├── LearnPage.tsx    # Articles
│   │   │   ├── HelpPage.tsx     # Help & FAQ
│   │   │   ├── SubmitPage.tsx   # Submit skill
│   │   │   └── NotFoundPage.tsx
│   │   ├── App.tsx
│   │   ├── main.tsx
│   │   └── index.css
│   ├── package.json
│   ├── vite.config.ts
│   ├── tailwind.config.js
│   └── .env.example
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- Google OAuth credentials (from [Google Cloud Console](https://console.cloud.google.com/apis/credentials))

### 1. Setup Google OAuth

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project (or select existing)
3. Navigate to **APIs & Services > Credentials**
4. Click **Create Credentials > OAuth Client ID**
5. Set application type to **Web application**
6. Add authorized redirect URIs:
   - `http://localhost:8080/api/auth/google/callback`
   - `http://localhost:5173` (for frontend)
7. Copy the **Client ID** and **Client Secret**

### 2. Backend Setup

```bash
cd backend

# Copy env file and fill in your values
cp .env.example .env

# Edit .env with your Google OAuth credentials
# GOOGLE_CLIENT_ID=your-client-id
# GOOGLE_CLIENT_SECRET=your-client-secret

# Download dependencies
go mod tidy

# Run the server
go run cmd/main.go
```

The API server starts at `http://localhost:8080`.

### 3. Frontend Setup

```bash
cd frontend

# Copy env file and fill in your values
cp .env.example .env

# Edit .env with your Google Client ID
# VITE_GOOGLE_CLIENT_ID=your-client-id

# Install dependencies
npm install

# Start dev server
npm run dev
```

The frontend starts at `http://localhost:5173`.

## API Endpoints

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/api/skills` | List skills (paginated) | No |
| GET | `/api/skills/featured` | Get featured skills | No |
| GET | `/api/skills/categories` | Get categories | No |
| GET | `/api/skills/:id` | Get skill detail | No |
| POST | `/api/skills` | Create skill | Yes |
| POST | `/api/skills/:id/like` | Toggle like | Yes |
| GET | `/api/skills/:id/comments` | Get comments | No |
| POST | `/api/skills/:id/comments` | Add comment | Yes |
| GET | `/api/articles` | List articles | No |
| GET | `/api/articles/:id` | Get article | No |
| GET | `/api/auth/google` | Google OAuth URL | No |
| GET | `/api/auth/google/callback` | OAuth callback | No |
| POST | `/api/auth/google/token` | Frontend token login | No |
| GET | `/api/auth/me` | Get current user | Yes |

## License

MIT
