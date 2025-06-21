# 🏥 VCM Medical Platform

**Advanced Medical Treatment Platform with 95% Efficacy**

A comprehensive multi-stakeholder medical platform connecting patients with specialized doctors and managing a complete ecosystem of agents, sales channels, distributors, and medical professionals.

## 🚀 Railway Deployment Guide

### Step 1: Push to GitHub
```bash
git add .
git commit -m "VCM Medical Platform - Ready for deployment"
git push origin main
```

### Step 2: Deploy on Railway
1. **In Railway Dashboard:**
   - Click "Deploy from GitHub repo"
   - Select your repository
   - Railway auto-detects Go application

2. **Add PostgreSQL:**
   - Click "Add Service" → "PostgreSQL"
   - Copy DATABASE_URL from PostgreSQL service

### Step 3: Set Environment Variables
```bash
DATABASE_URL=your_postgresql_connection_string_from_railway
JWT_SECRET=your-super-secret-production-key
PORT=8080
ENVIRONMENT=production
```

### Step 4: Setup Database
1. Connect to Railway PostgreSQL
2. Run SQL from `database/schema.sql`

### Step 5: Access Your App
- **Live URL:** `https://your-app.railway.app`
- **API Health:** `https://your-app.railway.app/health`

## ✨ Features

### 🔐 Authentication System
- Multi-user registration (8 user types)
- JWT-based authentication
- OTP verification (demo: 123456)
- Role-based access control

### 👥 User Types Supported
- **Patient (0)** - Medical treatment seekers
- **Doctor (5)** - Medical professionals  
- **Agent (1)** - Sales representatives
- **Distributor (4)** - Wholesale partners
- **Admin (11)** - Platform administrators

### 📱 Pages & Features
- **Home** - Landing page with platform overview
- **Register** - Account creation with user type selection
- **Login** - Secure authentication
- **Verify OTP** - Account verification (demo OTP: 123456)
- **Dashboard** - User-specific dashboard with stats and actions

## 🛠️ Technology Stack

### Backend (Go)
- **Fiber** - Web framework
- **GORM** - Database ORM
- **JWT** - Authentication
- **PostgreSQL** - Database
- **bcrypt** - Password hashing

### Frontend (React)
- **React 18** + TypeScript
- **Tailwind CSS** - Styling
- **React Router** - Navigation
- **React Hook Form** - Form handling
- **Zustand** - State management
- **Axios** - API communication

## 💻 Local Development

### Backend
```bash
cd backend
go mod tidy
go run cmd/main.go
```

### Frontend  
```bash
cd frontend
npm install
npm run dev
```

### Build Frontend
```bash
cd frontend
npm run build
```

## 📊 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login  
- `POST /api/v1/auth/verify-otp` - OTP verification

### Protected Routes
- `GET /api/v1/dashboard` - Dashboard data
- `GET /api/v1/profile` - User profile
- `PUT /api/v1/profile` - Update profile

### Patient Routes
- `GET /api/v1/patient/assessments` - Get assessments
- `POST /api/v1/patient/assessments` - Create assessment
- `GET /api/v1/patient/appointments` - Get appointments
- `POST /api/v1/patient/appointments` - Book appointment

## 🏗️ Project Structure

```
vcm-medical-platform/
├── backend/                    # Go backend
│   ├── cmd/main.go            # Application entry
│   ├── internal/              # Internal packages
│   │   ├── api/               # API routes
│   │   ├── config/            # Configuration
│   │   ├── database/          # Database connection
│   │   ├── handlers/          # Request handlers
│   │   ├── models/            # Data models
│   │   ├── middleware/        # Middleware
│   │   └── utils/             # Utilities
│   └── go.mod                 # Go dependencies
├── frontend/                   # React frontend
│   ├── src/                   # Source code
│   │   ├── components/        # UI components
│   │   ├── pages/             # Page components
│   │   ├── store/             # State management
│   │   └── utils/             # Utilities
│   ├── package.json           # Node dependencies
│   └── tailwind.config.js     # Tailwind config
├── database/                   # Database files
│   └── schema.sql             # Database schema
├── railway.json               # Railway config
└── README.md                  # Documentation
```

## 🔐 Demo Credentials

**Test Account:**
- Email: demo@vcm.com  
- Password: password123
- OTP: 123456 (any 6 digits work in demo)

## 🌍 Medical Specialties

1. **Cancer Immunotherapy** - CAR-T, BiTE, Oncolytic therapies
2. **Autoimmune Disorders** - Psoriasis, Rheumatoid Arthritis, Lupus
3. **Ophthalmology** - Optic nerve, Glaucoma, Macular degeneration
4. **Neurological Sciences** - Stroke, Alzheimer's, Autism
5. **Respiratory Medicine** - Tuberculosis, Pneumonia, Bronchitis
6. **Infectious Diseases** - HPV, Antibiotic-resistant infections

## 📈 Scalability Features

- **Microservices Ready** - Modular architecture
- **Database Optimization** - Indexed queries
- **JWT Stateless** - Horizontal scaling ready
- **API Versioning** - Future-proof endpoints
- **Error Handling** - Comprehensive error responses

## 🚀 Deployment Status

✅ **Backend API** - Go server with Fiber framework  
✅ **Frontend SPA** - React application with Tailwind CSS  
✅ **Database** - PostgreSQL with complete schema  
✅ **Authentication** - JWT with multi-user support  
✅ **Railway Config** - Ready for one-click deployment  
✅ **Health Check** - `/health` endpoint for monitoring  

## 📞 Support

For deployment or technical issues:
1. Check `/health` endpoint
2. Verify environment variables
3. Review Railway deployment logs
4. Ensure PostgreSQL connection

## 🎯 Next Steps

1. **Extend Features** - Add assessment forms, chat system
2. **Mobile App** - React Native companion
3. **Analytics** - Comprehensive reporting dashboard
4. **Payments** - WeChat Pay and Stripe integration
5. **AI Features** - Treatment recommendations

---

**🎉 Your VCM Medical Platform is ready for global medical innovation!**

© 2024 VAMOS BIOTECH (Shanghai) Co., Ltd.
