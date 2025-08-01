# Ujian Full-Stack Web App

## Overview
A simple authentication app using Go (Core.go) for backend and JSCroot for frontend. MongoDB stores user data. Frontend is deployable to GitHub Pages, backend to Render/Railway.

## Backend
- **Framework:** Core.go
- **Database:** MongoDB (`backend-ujian`), collection `User`
- **Endpoints:**
  - `POST /signup` — JSON `{username, password}`
  - `POST /login` — JSON `{username, password}`
- **CORS:** Enabled for all routes

## Frontend
- **JSCroot** via CDN
- **Pages:**
  - `index.html` — Login form
  - `signup.html` — Sign-up form
  - `dashboard.html` — Post-login
- **Assets:**
  - `assets/app.js` — Handles form logic, fetch API
  - `assets/style.css` — Simple styling

## Deployment
### Backend (Render)
1. Create a Render service for Go web server.
2. Set MongoDB connection string in `controllers/auth.go`.
3. Add `RENDER_API_KEY` and `YOUR_RENDER_SERVICE_ID` to GitHub repo secrets.
4. On push to `main`, GitHub Actions triggers deployment.

### Frontend (GitHub Pages)
1. Push `index.html`, `signup.html`, `dashboard.html`, and `assets/` to `main` branch.
2. Enable GitHub Pages in repo settings (root or `/docs`).
3. Update fetch URLs in `assets/app.js` to point to backend's public URL after deployment.

## Local Development
- **Backend:**
  - Run `go run main.go` (requires MongoDB running locally)
- **Frontend:**
  - Open `index.html` in browser

## Notes
- For production, use HTTPS and hash passwords.
- Update backend URL in frontend JS after deployment.
- See `.github/workflows/deploy-backend.yml` for CI/CD config.

---

**Key files:**
- `main.go`, `controllers/auth.go`, `go.mod`
- `index.html`, `signup.html`, `dashboard.html`, `assets/app.js`, `assets/style.css`
- `.github/workflows/deploy-backend.yml`