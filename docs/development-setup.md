# ZCalendar Development Setup

## Prerequisites

### System Requirements
- **Operating System**: macOS 12+ (for iOS), Windows 10+/macOS/Linux (for Android/Backend)
- **Memory**: 16GB RAM recommended
- **Storage**: 50GB free space

### Required Tools

#### Backend Development
- **Node.js**: Version 18.x or higher
- **npm**: Version 8.x or higher (comes with Node.js)
- **PostgreSQL**: Version 14.x or higher
- **Redis**: Version 6.x or higher (for caching)

#### iOS Development
- **macOS**: Required for iOS development
- **Xcode**: Version 14.0 or higher
- **iOS Simulator**: Included with Xcode
- **CocoaPods**: For dependency management

#### Android Development
- **Android Studio**: Latest stable version
- **Android SDK**: API level 24+ (Android 7.0+)
- **Java/Kotlin**: JDK 11 or higher

#### General Tools
- **Git**: Version control
- **Docker**: For containerized development (optional)
- **Postman**: API testing (optional)

## Environment Setup

### 1. Clone Repository
```bash
git clone https://github.com/ilmsadmin/zcalendar.git
cd zcalendar
```

### 2. Backend Setup

#### Install Dependencies
```bash
cd backend
npm install
```

#### Database Setup
```bash
# Install PostgreSQL (macOS with Homebrew)
brew install postgresql
brew services start postgresql

# Create database
createdb zcalendar_dev
createdb zcalendar_test

# Install Redis
brew install redis
brew services start redis
```

#### Environment Configuration
```bash
# Copy environment template
cp .env.example .env

# Edit .env file with your settings
nano .env
```

Example `.env` file:
```env
NODE_ENV=development
PORT=3000
DATABASE_URL=postgresql://username:password@localhost:5432/zcalendar_dev
TEST_DATABASE_URL=postgresql://username:password@localhost:5432/zcalendar_test
REDIS_URL=redis://localhost:6379
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRES_IN=7d
REFRESH_TOKEN_SECRET=your-refresh-token-secret
```

#### Run Database Migrations
```bash
npm run migrate:dev
npm run seed:dev  # Optional: seed with sample data
```

#### Start Backend Server
```bash
npm run dev
```

The backend API will be available at `http://localhost:3000`

### 3. iOS Development Setup

#### Prerequisites
```bash
# Install Xcode Command Line Tools
xcode-select --install

# Install CocoaPods
sudo gem install cocoapods
```

#### Project Setup
```bash
cd app_ios
pod install
open ZCalendar.xcworkspace  # Open in Xcode
```

#### Configuration
1. Open `ZCalendar.xcworkspace` in Xcode
2. Select your development team in project settings
3. Update bundle identifier if needed
4. Configure API base URL in `Config.swift`:
```swift
enum Config {
    static let apiBaseURL = "http://localhost:3000/api/v1"
}
```

#### Run iOS App
- Select target device/simulator in Xcode
- Press `Cmd + R` to build and run

### 4. Android Development Setup

#### Android Studio Setup
1. Download and install Android Studio
2. Open Android Studio and import the project:
   ```
   File > Open > Select /path/to/zcalendar/app_android
   ```
3. Wait for Gradle sync to complete
4. Install any missing SDK components when prompted

#### Configuration
Update `local.properties` with your Android SDK path:
```properties
sdk.dir=/Users/username/Library/Android/sdk
```

Update API configuration in `app/src/main/java/com/zcalendar/di/NetworkModule.kt`:
```kotlin
@Provides
@Singleton
fun provideBaseUrl(): String = "http://10.0.2.2:3000/api/v1/" // For emulator
// Use "http://localhost:3000/api/v1/" for physical device on same network
```

#### Run Android App
- Select target device/emulator
- Click the green play button or press `Shift + F10`

## Development Workflow

### Git Workflow
```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Make changes and commit
git add .
git commit -m "feat: add your feature description"

# Push to remote
git push origin feature/your-feature-name

# Create pull request via GitHub
```

### Backend Development

#### Running Tests
```bash
cd backend
npm test                    # Run all tests
npm run test:watch         # Run tests in watch mode
npm run test:coverage      # Run tests with coverage
```

#### Code Quality
```bash
npm run lint              # Check code style
npm run lint:fix          # Fix auto-fixable issues
npm run format           # Format code with Prettier
```

#### Database Operations
```bash
npm run migrate:create migration_name    # Create new migration
npm run migrate:up                      # Apply pending migrations
npm run migrate:down                    # Rollback last migration
npm run seed:create seed_name           # Create new seed
```

### iOS Development

#### Running Tests
```bash
# In Xcode
Product > Test (Cmd + U)

# Command line
xcodebuild test -workspace ZCalendar.xcworkspace -scheme ZCalendar -destination 'platform=iOS Simulator,name=iPhone 14'
```

#### Code Quality
- Enable SwiftLint in Xcode for code style checking
- Use Xcode's built-in formatter (Ctrl + I)

### Android Development

#### Running Tests
```bash
cd app_android
./gradlew test                    # Unit tests
./gradlew connectedAndroidTest    # Instrumented tests
```

#### Code Quality
```bash
./gradlew ktlintCheck            # Kotlin linting
./gradlew ktlintFormat           # Format Kotlin code
./gradlew detekt                 # Static analysis
```

## API Testing

### Using Postman
1. Import the Postman collection from `docs/postman_collection.json`
2. Set up environment variables:
   - `base_url`: `http://localhost:3000/api/v1`
   - `auth_token`: JWT token from login response

### Using curl
```bash
# Login
curl -X POST http://localhost:3000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'

# Get events
curl -X GET "http://localhost:3000/api/v1/events?start=2024-01-01&end=2024-01-31" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Debugging

### Backend Debugging
```bash
# Debug mode with inspector
npm run debug

# Then attach debugger in VS Code or Chrome DevTools
```

### iOS Debugging
- Use Xcode's built-in debugger
- Set breakpoints in source code
- Use `po` command in debugger console
- View memory graph for memory leaks

### Android Debugging
- Use Android Studio's debugger
- Set breakpoints in Kotlin/Java code
- Use `adb logcat` for log viewing
- Android Studio's Layout Inspector for UI debugging

## Common Issues and Solutions

### Backend Issues

#### Database Connection Error
```bash
# Check PostgreSQL status
brew services list | grep postgresql

# Restart PostgreSQL
brew services restart postgresql
```

#### Port Already in Use
```bash
# Find process using port 3000
lsof -ti:3000

# Kill the process
kill -9 $(lsof -ti:3000)
```

### iOS Issues

#### Pod Install Fails
```bash
cd app_ios
pod deintegrate
pod install --repo-update
```

#### Simulator Issues
```bash
# Reset simulator
Device > Erase All Content and Settings...

# Or restart simulator
xcrun simctl shutdown all
xcrun simctl boot "iPhone 14"
```

### Android Issues

#### Gradle Sync Fails
1. File > Invalidate Caches and Restart
2. Clean and rebuild project
3. Check internet connection for dependency downloads

#### Emulator Issues
```bash
# List available AVDs
emulator -list-avds

# Start specific AVD
emulator -avd Pixel_4_API_30
```

## IDE Configuration

### VS Code (Backend)
Recommended extensions:
- ES6 String HTML
- Prettier
- ESLint
- GitLens
- Thunder Client (API testing)

### Xcode (iOS)
Recommended settings:
- Enable SwiftLint build phase
- Set up code snippets for common patterns
- Configure source control integration

### Android Studio
Recommended plugins:
- Kotlin
- Android Parcelable code generator
- ADB Idea
- GitToolBox

## Performance Monitoring

### Backend
- Use `npm run monitor` for performance monitoring
- Set up logging with appropriate levels
- Monitor database query performance

### Mobile Apps
- Use Xcode Instruments for iOS performance profiling
- Use Android Studio Profiler for Android performance analysis
- Monitor memory usage and battery consumption