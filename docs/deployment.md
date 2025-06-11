# ZCalendar Deployment Guide

## Overview
This guide covers deployment strategies for the ZCalendar application across different environments and platforms.

## Infrastructure Requirements

### Backend Infrastructure
- **Server**: Cloud instance (AWS EC2, DigitalOcean, etc.)
- **Database**: PostgreSQL (managed service recommended)
- **Cache**: Redis instance
- **Load Balancer**: For production scaling
- **CDN**: For static asset delivery
- **Monitoring**: Application and infrastructure monitoring

### Mobile App Distribution
- **iOS**: Apple App Store
- **Android**: Google Play Store
- **Enterprise**: Internal distribution (optional)

## Backend Deployment

### Environment Setup

#### Production Environment Variables
```env
NODE_ENV=production
PORT=3000
DATABASE_URL=postgresql://user:pass@prod-db:5432/zcalendar
REDIS_URL=redis://prod-redis:6379
JWT_SECRET=production-secret-key
API_RATE_LIMIT=1000
CORS_ORIGIN=https://zcalendar.com
```

#### Docker Configuration
```dockerfile
# Dockerfile
FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci --only=production

# Copy source code
COPY . .

# Build application
RUN npm run build

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD node healthcheck.js

# Start application
CMD ["npm", "start"]
```

#### Docker Compose for Production
```yaml
# docker-compose.prod.yml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=production
      - DATABASE_URL=postgresql://user:pass@db:5432/zcalendar
      - REDIS_URL=redis://redis:6379
    depends_on:
      - db
      - redis
    restart: unless-stopped

  db:
    image: postgres:14
    environment:
      POSTGRES_DB: zcalendar
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

  redis:
    image: redis:6-alpine
    restart: unless-stopped

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - app
    restart: unless-stopped

volumes:
  postgres_data:
```

### Deployment Strategies

#### 1. Manual Deployment
```bash
# Build and deploy to server
npm run build
rsync -avz --exclude node_modules . user@server:/path/to/app
ssh user@server "cd /path/to/app && npm ci --production && pm2 restart app"
```

#### 2. CI/CD with GitHub Actions
```yaml
# .github/workflows/deploy.yml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Setup Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18'
        cache: 'npm'
    
    - name: Install dependencies
      run: npm ci
    
    - name: Run tests
      run: npm test
    
    - name: Build application
      run: npm run build
    
    - name: Deploy to server
      uses: appleboy/ssh-action@v0.1.5
      with:
        host: ${{ secrets.HOST }}
        username: ${{ secrets.USERNAME }}
        key: ${{ secrets.PRIVATE_KEY }}
        script: |
          cd /path/to/app
          git pull origin main
          npm ci --production
          npm run migrate:up
          pm2 restart zcalendar
```

#### 3. Container Orchestration (Kubernetes)
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: zcalendar-backend
spec:
  replicas: 3
  selector:
    matchLabels:
      app: zcalendar-backend
  template:
    metadata:
      labels:
        app: zcalendar-backend
    spec:
      containers:
      - name: backend
        image: zcalendar/backend:latest
        ports:
        - containerPort: 3000
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: zcalendar-secrets
              key: database-url
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
---
apiVersion: v1
kind: Service
metadata:
  name: zcalendar-backend-service
spec:
  selector:
    app: zcalendar-backend
  ports:
  - port: 80
    targetPort: 3000
  type: LoadBalancer
```

### Database Migration Strategy
```bash
# Production migration process
npm run migrate:status          # Check migration status
npm run migrate:up              # Apply pending migrations
npm run migrate:rollback        # Rollback if needed (emergency)

# Blue-green deployment with migrations
# 1. Deploy new version to green environment
# 2. Run migrations on green database
# 3. Switch traffic to green
# 4. Keep blue as fallback
```

## Mobile App Deployment

### iOS App Store Deployment

#### Prerequisites
- Apple Developer Account ($99/year)
- App Store Connect access
- Code signing certificates
- Provisioning profiles

#### Build Process
```bash
# Archive for distribution
xcodebuild -workspace ZCalendar.xcworkspace \
           -scheme ZCalendar \
           -configuration Release \
           -archivePath build/ZCalendar.xcarchive \
           archive

# Export for App Store
xcodebuild -exportArchive \
           -archivePath build/ZCalendar.xcarchive \
           -exportPath build \
           -exportOptionsPlist ExportOptions.plist
```

#### App Store Connect Process
1. **Create App Record**
   - App name, bundle ID, SKU
   - App categories and metadata
   - Age rating questionnaire

2. **Upload Build**
   - Use Xcode or Transporter app
   - Wait for processing (15-60 minutes)
   - Resolve any issues reported

3. **Submission for Review**
   - Complete app information
   - Add screenshots and descriptions
   - Set pricing and availability
   - Submit for review (1-3 days)

#### TestFlight Beta Testing
```bash
# Upload beta build
xcodebuild -exportArchive \
           -archivePath build/ZCalendar.xcarchive \
           -exportPath build \
           -exportOptionsPlist ExportOptionsTestFlight.plist

# Distribute to internal testers (immediate)
# Distribute to external testers (requires beta review)
```

### Android Play Store Deployment

#### Prerequisites
- Google Play Developer Account ($25 one-time)
- Play Console access
- Signing key for app releases

#### Build Process
```bash
cd app_android

# Build release APK
./gradlew assembleRelease

# Build App Bundle (recommended)
./gradlew bundleRelease

# Sign the release
jarsigner -verbose -sigalg SHA256withRSA -digestalg SHA-256 \
          -keystore release-key.keystore \
          app-release-unsigned.apk \
          alias_name

# Optimize APK
zipalign -v 4 app-release-unsigned.apk app-release.apk
```

#### Play Console Process
1. **Create App**
   - App name and default language
   - App category and tags
   - Content rating questionnaire

2. **Upload Release**
   - Upload AAB file to production track
   - Add release notes
   - Review and rollout percentage

3. **Store Listing**
   - App description and screenshots
   - Feature graphic and icon
   - Pricing and distribution

#### Internal Testing Track
```bash
# Upload to internal testing
./gradlew publishReleaseBundle --track internal

# Promote to alpha/beta/production when ready
```

## Monitoring and Maintenance

### Backend Monitoring

#### Application Monitoring
```javascript
// Setup monitoring with PM2
module.exports = {
  apps: [{
    name: 'zcalendar',
    script: './dist/server.js',
    instances: 'max',
    exec_mode: 'cluster',
    env: {
      NODE_ENV: 'production'
    },
    error_file: './logs/err.log',
    out_file: './logs/out.log',
    log_file: './logs/combined.log',
    time: true
  }]
}
```

#### Health Checks
```javascript
// Health check endpoint
app.get('/health', async (req, res) => {
  try {
    // Check database connection
    await db.raw('SELECT 1');
    
    // Check Redis connection
    await redis.ping();
    
    res.status(200).json({
      status: 'healthy',
      timestamp: new Date().toISOString(),
      version: process.env.npm_package_version
    });
  } catch (error) {
    res.status(503).json({
      status: 'unhealthy',
      error: error.message
    });
  }
});
```

#### Logging Configuration
```javascript
// Winston logger setup
const winston = require('winston');

const logger = winston.createLogger({
  level: 'info',
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.errors({ stack: true }),
    winston.format.json()
  ),
  transports: [
    new winston.transports.File({ filename: 'error.log', level: 'error' }),
    new winston.transports.File({ filename: 'combined.log' }),
    new winston.transports.Console({
      format: winston.format.simple()
    })
  ]
});
```

### Mobile App Analytics

#### iOS Analytics (App Store Connect)
- App downloads and revenue
- Crash reports and diagnostics
- User engagement metrics
- Performance metrics

#### Android Analytics (Play Console)
- Install statistics
- User acquisition reports
- Crash and ANR reports
- Performance data

#### Custom Analytics
```swift
// iOS - Firebase Analytics
Analytics.logEvent("event_created", parameters: [
  "event_type": "meeting",
  "duration": 60
])
```

```kotlin
// Android - Firebase Analytics
firebaseAnalytics.logEvent("event_created") {
    param("event_type", "meeting")
    param("duration", 60)
}
```

## Security Considerations

### Backend Security
- SSL/TLS encryption for all communications
- Security headers (HSTS, CSP, etc.)
- Rate limiting and DDoS protection
- Regular security updates
- Vulnerability scanning

### Mobile App Security
- Code obfuscation for release builds
- Certificate pinning for API communication
- Secure storage for sensitive data
- Regular security reviews

## Backup and Recovery

### Database Backups
```bash
# Automated daily backups
pg_dump -h $DB_HOST -U $DB_USER -d $DB_NAME | gzip > backup_$(date +%Y%m%d).sql.gz

# Restore from backup
gunzip -c backup_20240115.sql.gz | psql -h $DB_HOST -U $DB_USER -d $DB_NAME
```

### Application Backups
- Configuration files
- SSL certificates
- Application logs
- User uploaded files

## Performance Optimization

### Backend Optimization
- Enable gzip compression
- Implement caching strategies
- Database query optimization
- CDN for static assets

### Mobile App Optimization
- Image optimization and caching
- Lazy loading of data
- Offline capabilities
- Background sync optimization

## Rollback Procedures

### Backend Rollback
```bash
# Quick rollback with PM2
pm2 stop zcalendar
git checkout previous-stable-tag
npm ci --production
npm run migrate:down  # If needed
pm2 start zcalendar
```

### Mobile App Rollback
- **iOS**: Use App Store Connect to revert to previous version
- **Android**: Use Play Console staged rollouts to control percentage
- **Emergency**: Remove app from stores temporarily if critical issue