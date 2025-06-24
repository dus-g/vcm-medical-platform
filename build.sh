#!/bin/bash
set -e

echo "🏗️  Building VCM Medical Platform..."

# Build frontend first
echo "📦 Building React frontend..."
cd frontend
npm ci
npm run build
cd ..

# Verify frontend build exists
if [ ! -d "frontend/dist" ]; then
    echo "❌ Frontend build failed - no dist directory found"
    exit 1
fi

echo "✅ Frontend build complete"

# Build Go binary
echo "🚀 Building Go backend..."
go build -o main .

echo "✅ Build complete!"
