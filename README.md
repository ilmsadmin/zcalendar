# ZCalendar

A comprehensive cross-platform calendar application with native mobile apps and a powerful backend API.

## Overview

ZCalendar provides users with a seamless calendar experience across iOS and Android devices, featuring real-time synchronization, smart reminders, and intuitive event management.

## Project Structure

```
zcalendar/
├── docs/                 # Project documentation
│   ├── README.md         # Documentation overview
│   ├── project-overview.md
│   ├── architecture.md
│   ├── api.md
│   ├── database.md
│   ├── ios.md
│   ├── android.md
│   ├── development-setup.md
│   ├── deployment.md
│   └── user-manual.md
├── backend/              # Node.js API server
│   └── README.md
├── app_ios/              # iOS application (Swift)
│   └── README.md
├── app_android/          # Android application (Kotlin)
│   └── README.md
└── README.md            # This file
```

## Key Features

- **Cross-Platform**: Native iOS and Android applications
- **Real-time Sync**: Seamless synchronization across devices
- **Multiple Calendar Views**: Day, week, month, and list views
- **Event Management**: Create, edit, delete, and organize events
- **Smart Reminders**: Customizable notifications and alerts
- **Calendar Sharing**: Share calendars with family and colleagues
- **Offline Support**: Core functionality available without internet
- **Recurring Events**: Flexible recurring event patterns

## Technology Stack

### Backend
- **Runtime**: Node.js 18+
- **Framework**: Express.js
- **Database**: PostgreSQL with Prisma ORM
- **Cache**: Redis
- **Authentication**: JWT tokens

### iOS App
- **Language**: Swift 5.7+
- **UI Framework**: SwiftUI
- **Architecture**: MVVM
- **Local Storage**: Core Data
- **Minimum iOS**: 14.0+

### Android App
- **Language**: Kotlin 1.8+
- **UI Framework**: Jetpack Compose
- **Architecture**: MVVM with Clean Architecture
- **Local Storage**: Room
- **Minimum Android**: 7.0 (API 24)

## Quick Start

### Prerequisites
- Node.js 18+ (for backend)
- Xcode 14+ (for iOS development)
- Android Studio (for Android development)
- PostgreSQL 14+
- Redis 6+

### Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/ilmsadmin/zcalendar.git
   cd zcalendar
   ```

2. **Backend Setup**
   ```bash
   cd backend
   npm install
   cp .env.example .env
   # Configure your .env file
   npm run migrate:dev
   npm run dev
   ```

3. **iOS Setup**
   ```bash
   cd app_ios
   pod install
   open ZCalendar.xcworkspace
   # Build and run in Xcode
   ```

4. **Android Setup**
   ```bash
   cd app_android
   # Open in Android Studio
   # Sync Gradle and run
   ```

For detailed setup instructions, see [Development Setup Guide](./docs/development-setup.md).

## Documentation

Comprehensive documentation is available in the `docs/` directory:

- **[Project Overview](./docs/project-overview.md)** - High-level project description and goals
- **[Architecture](./docs/architecture.md)** - System architecture and design patterns
- **[API Documentation](./docs/api.md)** - Backend API specifications
- **[Database Design](./docs/database.md)** - Database schema and design
- **[iOS Documentation](./docs/ios.md)** - iOS app development details
- **[Android Documentation](./docs/android.md)** - Android app development details
- **[Development Setup](./docs/development-setup.md)** - Environment setup guide
- **[Deployment Guide](./docs/deployment.md)** - Production deployment instructions
- **[User Manual](./docs/user-manual.md)** - End-user documentation

## Contributing

We welcome contributions to ZCalendar! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Write or update tests as needed
5. Follow the existing code style and conventions
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Development Guidelines

- Follow the established architecture patterns
- Write tests for new functionality
- Update documentation as needed
- Use meaningful commit messages
- Ensure code passes all linting and tests

## Testing

Each component has its own testing strategy:

- **Backend**: Jest for unit and integration tests
- **iOS**: XCTest for unit and UI tests
- **Android**: JUnit and Espresso for testing

Run tests in each directory:
```bash
# Backend
cd backend && npm test

# iOS (in Xcode)
Product → Test (Cmd+U)

# Android
cd app_android && ./gradlew test
```

## Deployment

ZCalendar supports multiple deployment strategies:

- **Backend**: Docker containers, Kubernetes, or traditional server deployment
- **iOS**: App Store distribution via Xcode and App Store Connect
- **Android**: Google Play Store via Android Studio and Play Console

See the [Deployment Guide](./docs/deployment.md) for detailed instructions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: Check the [docs](./docs/) directory
- **Issues**: Create a GitHub issue for bugs or feature requests
- **Email**: support@zcalendar.com (if available)

## Roadmap

### Version 1.0
- [x] Project setup and documentation
- [ ] Backend API implementation
- [ ] iOS app development
- [ ] Android app development
- [ ] Basic calendar functionality
- [ ] User authentication
- [ ] Event CRUD operations

### Version 1.1
- [ ] Real-time synchronization
- [ ] Push notifications
- [ ] Calendar sharing
- [ ] Recurring events

### Version 2.0
- [ ] Advanced features (widgets, integrations)
- [ ] Performance optimizations
- [ ] Enhanced UI/UX
- [ ] Analytics and insights

## Acknowledgments

- Material Design team for Android design guidelines
- Apple Human Interface Guidelines for iOS design
- Open source community for various tools and libraries used in this project