# ZCalendar iOS App

This directory contains the native iOS application for ZCalendar.

## Requirements

- **iOS**: 14.0+
- **Xcode**: 14.0+
- **Swift**: 5.7+
- **CocoaPods**: For dependency management

## Technology Stack

- **UI Framework**: SwiftUI with UIKit integration
- **Architecture**: MVVM pattern
- **Data Persistence**: Core Data
- **Networking**: URLSession with async/await
- **Reactive Programming**: Combine framework
- **Testing**: XCTest

## Getting Started

1. Install CocoaPods dependencies:
   ```bash
   cd app_ios
   pod install
   ```

2. Open workspace in Xcode:
   ```bash
   open ZCalendar.xcworkspace
   ```

3. Configure your development team in project settings

4. Build and run on simulator or device

## Project Structure

```
app_ios/
├── ZCalendar/
│   ├── App/                 # App delegate and main app file
│   ├── Views/               # SwiftUI views
│   │   ├── Calendar/        # Calendar-related views
│   │   ├── Events/          # Event management views
│   │   └── Settings/        # Settings and preferences
│   ├── ViewModels/          # MVVM view models
│   ├── Models/              # Data models
│   ├── Services/            # API and data services
│   ├── Utilities/           # Helper functions and extensions
│   └── Resources/           # Assets, localizations
├── ZCalendarTests/          # Unit tests
├── ZCalendarUITests/        # UI tests
└── Pods/                    # CocoaPods dependencies
```

## Key Features

### Calendar Views
- Month view with event indicators
- Week view with time slots
- Day view with detailed scheduling
- List view for upcoming events

### Event Management
- Create, edit, and delete events
- Recurring event support
- Multiple calendar organization
- Event reminders and notifications

### Offline Support
- Core Data for local storage
- Background sync with server
- Conflict resolution
- Offline event creation

### iOS Integration
- System calendar synchronization
- Siri shortcuts support
- Widget support for Today view
- Apple Watch companion (future)

## Dependencies

Main dependencies managed through CocoaPods:
- **Alamofire**: Networking
- **KeychainAccess**: Secure storage
- **SwiftLint**: Code style enforcement

To update dependencies:
```bash
pod update
```

## Building and Testing

### Development Build
```bash
# Build for simulator
xcodebuild -workspace ZCalendar.xcworkspace -scheme ZCalendar -destination 'platform=iOS Simulator,name=iPhone 14'

# Run on device (requires provisioning profile)
xcodebuild -workspace ZCalendar.xcworkspace -scheme ZCalendar -destination 'platform=iOS,name=Your Device'
```

### Testing
```bash
# Run unit tests
xcodebuild test -workspace ZCalendar.xcworkspace -scheme ZCalendar -destination 'platform=iOS Simulator,name=iPhone 14'

# Run UI tests
xcodebuild test -workspace ZCalendar.xcworkspace -scheme ZCalendarUITests -destination 'platform=iOS Simulator,name=iPhone 14'
```

### Distribution Build
1. Archive the project in Xcode (Product → Archive)
2. Export for App Store distribution
3. Upload to App Store Connect

## Configuration

### API Configuration
Update the API base URL in `Config.swift`:
```swift
enum Config {
    static let apiBaseURL = "https://api.zcalendar.com/v1"
    // Use "http://localhost:3000/api/v1" for development
}
```

### App Icons and Assets
- App icons: `ZCalendar/Resources/Assets.xcassets/AppIcon.appiconset/`
- Other assets: `ZCalendar/Resources/Assets.xcassets/`

## Code Style

- Follow Swift API Design Guidelines
- Use SwiftLint for consistent formatting
- Document public APIs with Swift documentation comments
- Prefer SwiftUI over UIKit for new features

## Contributing

1. Create feature branch from `develop`
2. Follow existing code patterns and architecture
3. Write unit tests for new functionality
4. Test on multiple device sizes and iOS versions
5. Update documentation as needed

## Debugging

### Common Issues

#### Build Errors
- Clean build folder: Product → Clean Build Folder
- Update CocoaPods: `pod update`
- Check provisioning profiles and certificates

#### Runtime Issues
- Check Console app for crash logs
- Use Xcode debugger and breakpoints
- Test on different iOS versions and devices

### Performance Profiling
- Use Instruments for memory and performance analysis
- Monitor Core Data performance
- Check network request efficiency

## Deployment

### TestFlight Beta
1. Archive the app in Xcode
2. Upload to App Store Connect
3. Add to TestFlight for beta testing

### App Store Release
1. Complete app metadata in App Store Connect
2. Submit for review
3. Release when approved

For detailed deployment instructions, see the main project documentation.

## Support

For iOS-specific issues:
- Check Xcode console for errors
- Review device logs in Xcode (Window → Devices and Simulators)
- Test on different iOS versions

For general support, see the main project README.