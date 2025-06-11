# ZCalendar Android App

This directory contains the native Android application for ZCalendar.

## Requirements

- **Android**: 7.0+ (API level 24+)
- **Android Studio**: Latest stable version
- **Kotlin**: 1.8+
- **Gradle**: 8.0+

## Technology Stack

- **UI Framework**: Jetpack Compose
- **Architecture**: MVVM with Clean Architecture
- **Dependency Injection**: Hilt
- **Database**: Room
- **Networking**: Retrofit + OkHttp
- **Async Programming**: Coroutines + Flow
- **Testing**: JUnit, Espresso, Compose Testing

## Getting Started

1. Open project in Android Studio:
   ```bash
   cd app_android
   # Open Android Studio and import this directory
   ```

2. Sync Gradle dependencies:
   - Android Studio will prompt to sync automatically
   - Or manually: File → Sync Project with Gradle Files

3. Build and run:
   - Select target device/emulator
   - Click Run button or press Shift+F10

## Project Structure

```
app_android/
├── app/
│   ├── src/
│   │   ├── main/
│   │   │   ├── java/com/zcalendar/
│   │   │   │   ├── di/              # Dependency injection
│   │   │   │   ├── data/            # Data layer
│   │   │   │   │   ├── local/       # Room database
│   │   │   │   │   ├── remote/      # API services
│   │   │   │   │   └── repository/  # Repository implementations
│   │   │   │   ├── domain/          # Domain layer
│   │   │   │   │   ├── model/       # Domain models
│   │   │   │   │   ├── repository/  # Repository interfaces
│   │   │   │   │   └── usecase/     # Use cases
│   │   │   │   ├── presentation/    # UI layer
│   │   │   │   │   ├── ui/          # Composables
│   │   │   │   │   ├── viewmodel/   # ViewModels
│   │   │   │   │   └── theme/       # App theme
│   │   │   │   └── util/            # Utilities
│   │   │   └── res/                 # Resources
│   │   ├── test/                    # Unit tests
│   │   └── androidTest/             # Instrumented tests
│   ├── build.gradle.kts
│   └── proguard-rules.pro
├── gradle/
├── build.gradle.kts
└── settings.gradle.kts
```

## Key Features

### Calendar Views
- Month view with Material Design 3
- Week view with time slots
- Day view for detailed scheduling
- List view with infinite scrolling

### Event Management
- Create/edit/delete events with intuitive forms
- Recurring events with flexible patterns
- Multiple calendar support with color coding
- Event reminders and notifications

### Android Integration
- System calendar synchronization
- Widget support for home screen
- Quick settings tile (future)
- Android Auto integration (future)

### Offline Support
- Room database for local storage
- WorkManager for background sync
- Conflict resolution strategies
- Offline-first architecture

## Dependencies

Key dependencies (see `app/build.gradle.kts` for versions):

```kotlin
// UI
implementation("androidx.compose.ui:ui")
implementation("androidx.compose.material3:material3")
implementation("androidx.activity:activity-compose")

// Architecture
implementation("androidx.lifecycle:lifecycle-viewmodel-compose")
implementation("androidx.navigation:navigation-compose")

// Dependency Injection
implementation("com.google.dagger:hilt-android")
kapt("com.google.dagger:hilt-compiler")

// Database
implementation("androidx.room:room-runtime")
implementation("androidx.room:room-ktx")
kapt("androidx.room:room-compiler")

// Networking
implementation("com.squareup.retrofit2:retrofit")
implementation("com.squareup.retrofit2:converter-gson")
implementation("com.squareup.okhttp3:logging-interceptor")

// Background Work
implementation("androidx.work:work-runtime-ktx")
implementation("androidx.hilt:hilt-work")
```

## Building and Testing

### Development Build
```bash
cd app_android

# Build debug APK
./gradlew assembleDebug

# Install on connected device/emulator
./gradlew installDebug

# Build and install in one command
./gradlew installDebug
```

### Testing
```bash
# Run unit tests
./gradlew test

# Run instrumented tests (requires device/emulator)
./gradlew connectedAndroidTest

# Run specific test class
./gradlew test --tests "com.zcalendar.EventViewModelTest"
```

### Release Build
```bash
# Build release APK (requires signing configuration)
./gradlew assembleRelease

# Build App Bundle for Play Store
./gradlew bundleRelease
```

### Code Quality
```bash
# Kotlin linting
./gradlew ktlintCheck

# Format Kotlin code
./gradlew ktlintFormat

# Static analysis
./gradlew detekt
```

## Configuration

### API Configuration
Update API base URL in `NetworkModule.kt`:
```kotlin
@Provides
@Singleton
fun provideBaseUrl(): String {
    return if (BuildConfig.DEBUG) {
        "http://10.0.2.2:3000/api/v1/" // For emulator
    } else {
        "https://api.zcalendar.com/v1/"
    }
}
```

### App Configuration
- Package name: `com.zcalendar.android`
- Minimum SDK: API 24 (Android 7.0)
- Target SDK: API 33 (Android 13)

## Architecture Overview

### Clean Architecture Layers

1. **Presentation Layer** (`presentation/`)
   - Composables and ViewModels
   - UI state management
   - User interaction handling

2. **Domain Layer** (`domain/`)
   - Business logic and use cases
   - Domain models
   - Repository interfaces

3. **Data Layer** (`data/`)
   - Repository implementations
   - Local database (Room)
   - Remote API (Retrofit)

### Data Flow
```
UI (Compose) → ViewModel → UseCase → Repository → DataSource
                ↑                                      ↓
             UI State ← ← ← ← ← ← ← ← ← ← ← ← ← ← ← Data
```

## Compose UI Examples

### Calendar Screen
```kotlin
@Composable
fun CalendarScreen(
    viewModel: CalendarViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()
    
    LazyColumn {
        items(uiState.events) { event ->
            EventCard(
                event = event,
                onClick = { viewModel.selectEvent(event.id) }
            )
        }
    }
}
```

### Event Card Component
```kotlin
@Composable
fun EventCard(
    event: Event,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Card(
        modifier = modifier
            .fillMaxWidth()
            .clickable { onClick() },
        colors = CardDefaults.cardColors(
            containerColor = Color(event.calendar.color)
        )
    ) {
        Column(
            modifier = Modifier.padding(16.dp)
        ) {
            Text(
                text = event.title,
                style = MaterialTheme.typography.titleMedium
            )
            Text(
                text = event.formattedTime,
                style = MaterialTheme.typography.bodySmall
            )
        }
    }
}
```

## Debugging

### Common Issues

#### Build Errors
- Clean project: Build → Clean Project
- Invalidate caches: File → Invalidate Caches and Restart
- Check Gradle sync errors in Build tab

#### Runtime Issues
- Check Logcat for error messages
- Use breakpoints in Android Studio debugger
- Test on different Android versions and screen sizes

### Performance Analysis
- Use Android Studio Profiler
- Monitor memory usage and CPU performance
- Check database query performance
- Analyze network request efficiency

## Testing Strategy

### Unit Tests
```kotlin
@Test
fun `getEvents returns cached data when offline`() = runTest {
    // Given
    val cachedEvents = listOf(testEvent)
    coEvery { localDataSource.getEvents() } returns cachedEvents
    coEvery { remoteDataSource.getEvents() } throws NetworkException()
    
    // When
    val result = repository.getEvents().first()
    
    // Then
    assertEquals(cachedEvents, result)
}
```

### Compose UI Tests
```kotlin
@Test
fun eventCard_displaysEventInformation() {
    composeTestRule.setContent {
        ZCalendarTheme {
            EventCard(
                event = testEvent,
                onClick = { }
            )
        }
    }
    
    composeTestRule
        .onNodeWithText("Team Meeting")
        .assertIsDisplayed()
}
```

## Deployment

### Debug Testing
- Install on physical device for real-world testing
- Test different screen sizes and orientations
- Verify offline functionality

### Release Process
1. Update version in `build.gradle.kts`
2. Build signed APK/AAB
3. Test release build thoroughly
4. Upload to Play Console

### Play Store Requirements
- Target latest Android API level
- Follow Play Store policies
- Provide required metadata and screenshots
- Test on various devices

## Performance Optimization

### Compose Performance
- Use `remember` for expensive calculations
- Avoid recomposition with stable parameters
- Use `LazyColumn` for large lists
- Profile with Compose Layout Inspector

### Database Performance
- Use appropriate indexes
- Implement pagination for large datasets
- Cache frequently accessed data
- Monitor query performance

### Network Performance
- Implement request caching
- Use compression for API responses
- Handle network errors gracefully
- Implement retry strategies

## Contributing

1. Follow Kotlin coding conventions
2. Use Material Design 3 components
3. Write tests for new features
4. Update documentation
5. Follow existing architecture patterns

## Support

For Android-specific issues:
- Check Logcat output
- Test on multiple Android versions
- Use Android Studio's debugging tools
- Monitor crash reports in Play Console

For general support, see the main project README.