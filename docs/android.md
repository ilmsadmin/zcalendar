# ZCalendar Android App Documentation

## Overview
The ZCalendar Android application provides a native Material Design calendar experience for Android devices, supporting phones and tablets.

## Technical Specifications

### Requirements
- **Android Version**: Android 7.0 (API 24)+
- **Compile SDK**: Android 13 (API 33)
- **Language**: Kotlin 1.8+
- **Build System**: Gradle with Android Gradle Plugin

### Key Libraries
- **Jetpack Compose**: Modern UI toolkit
- **Architecture Components**: ViewModel, LiveData, Room
- **Retrofit**: HTTP client for API communication
- **Hilt**: Dependency injection
- **Room**: Local database
- **WorkManager**: Background task scheduling

## App Architecture

### MVVM with Clean Architecture
```
┌─────────────────┐
│   UI Layer      │ ← Composables/Activities
│  (Compose)      │
└─────────┬───────┘
          │ State/Events
┌─────────▼───────┐
│ Presentation    │ ← ViewModels
│   Layer         │
└─────────┬───────┘
          │ Use Cases
┌─────────▼───────┐
│  Domain Layer   │ ← Business Logic
│                 │
└─────────┬───────┘
          │ Repository
┌─────────▼───────┐
│  Data Layer     │ ← Room/Retrofit
│                 │
└─────────────────┘
```

### Project Structure
```
app/
├── src/main/
│   ├── java/com/zcalendar/
│   │   ├── di/              # Dependency Injection
│   │   ├── data/            # Data Layer
│   │   │   ├── local/       # Room database
│   │   │   ├── remote/      # API services
│   │   │   └── repository/  # Repository implementations
│   │   ├── domain/          # Domain Layer
│   │   │   ├── model/       # Domain models
│   │   │   ├── repository/  # Repository interfaces
│   │   │   └── usecase/     # Use cases
│   │   ├── presentation/    # UI Layer
│   │   │   ├── ui/          # Composables
│   │   │   ├── viewmodel/   # ViewModels
│   │   │   └── theme/       # App theme
│   │   └── util/            # Utilities
│   └── res/                 # Resources
│       ├── values/          # Strings, colors, dimens
│       └── drawable/        # Icons and images
└── build.gradle.kts
```

## Key Features Implementation

### Calendar Views
```kotlin
@Composable
fun CalendarScreen(
    viewModel: CalendarViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()
    
    when (uiState.viewMode) {
        ViewMode.MONTH -> MonthView(
            events = uiState.events,
            selectedDate = uiState.selectedDate,
            onDateSelected = viewModel::selectDate
        )
        ViewMode.WEEK -> WeekView(...)
        ViewMode.DAY -> DayView(...)
    }
}
```

### Event Management
```kotlin
// Event Entity
@Entity(tableName = "events")
data class EventEntity(
    @PrimaryKey val id: String,
    val title: String,
    val description: String?,
    val startTime: Long,
    val endTime: Long,
    val isAllDay: Boolean,
    val location: String?,
    val calendarId: String,
    val reminders: String // JSON array
)

// Event Repository
class EventRepositoryImpl @Inject constructor(
    private val localDataSource: EventLocalDataSource,
    private val remoteDataSource: EventRemoteDataSource
) : EventRepository {
    
    override suspend fun getEvents(
        startDate: LocalDate,
        endDate: LocalDate
    ): Flow<List<Event>> = flow {
        // Emit cached data first
        emit(localDataSource.getEvents(startDate, endDate))
        
        // Fetch from remote and update cache
        try {
            val remoteEvents = remoteDataSource.getEvents(startDate, endDate)
            localDataSource.insertEvents(remoteEvents)
            emit(remoteEvents)
        } catch (e: Exception) {
            // Handle error, keep cached data
        }
    }
}
```

### Offline Support
```kotlin
// Sync Worker
@HiltWorker
class SyncWorker @AssistedInject constructor(
    @Assisted context: Context,
    @Assisted workerParams: WorkerParameters,
    private val repository: EventRepository
) : CoroutineWorker(context, workerParams) {
    
    override suspend fun doWork(): Result {
        return try {
            repository.syncPendingChanges()
            Result.success()
        } catch (e: Exception) {
            Result.retry()
        }
    }
}
```

## UI Implementation with Jetpack Compose

### Theme and Design System
```kotlin
// Material 3 Theme
@Composable
fun ZCalendarTheme(
    darkTheme: Boolean = isSystemInDarkTheme(),
    content: @Composable () -> Unit
) {
    val colorScheme = when {
        darkTheme -> DarkColorScheme
        else -> LightColorScheme
    }

    MaterialTheme(
        colorScheme = colorScheme,
        typography = Typography,
        content = content
    )
}

// Custom Components
@Composable
fun CalendarCard(
    event: Event,
    onClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Card(
        modifier = modifier.clickable { onClick() },
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
            if (event.location != null) {
                Text(
                    text = event.location,
                    style = MaterialTheme.typography.bodySmall
                )
            }
        }
    }
}
```

### Navigation
```kotlin
// Navigation Setup
@Composable
fun ZCalendarNavigation() {
    val navController = rememberNavController()
    
    NavHost(
        navController = navController,
        startDestination = "calendar"
    ) {
        composable("calendar") {
            CalendarScreen(
                onEventClick = { eventId ->
                    navController.navigate("event_detail/$eventId")
                }
            )
        }
        
        composable(
            "event_detail/{eventId}",
            arguments = listOf(navArgument("eventId") { type = NavType.StringType })
        ) { backStackEntry ->
            val eventId = backStackEntry.arguments?.getString("eventId")
            EventDetailScreen(
                eventId = eventId,
                onNavigateBack = { navController.popBackStack() }
            )
        }
    }
}
```

## Data Persistence

### Room Database
```kotlin
@Database(
    entities = [EventEntity::class, CalendarEntity::class],
    version = 1,
    exportSchema = false
)
@TypeConverters(Converters::class)
abstract class ZCalendarDatabase : RoomDatabase() {
    abstract fun eventDao(): EventDao
    abstract fun calendarDao(): CalendarDao
}

// DAO Interface
@Dao
interface EventDao {
    @Query("""
        SELECT * FROM events 
        WHERE start_time >= :startDate 
        AND end_time <= :endDate 
        ORDER BY start_time ASC
    """)
    suspend fun getEventsByDateRange(
        startDate: Long,
        endDate: Long
    ): List<EventEntity>
    
    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertEvents(events: List<EventEntity>)
}
```

## Background Tasks and Notifications

### WorkManager for Sync
```kotlin
// Periodic Sync
class SyncScheduler @Inject constructor(
    private val workManager: WorkManager
) {
    fun schedulePeriodicSync() {
        val syncRequest = PeriodicWorkRequestBuilder<SyncWorker>(
            repeatInterval = 15,
            repeatIntervalTimeUnit = TimeUnit.MINUTES
        ).setConstraints(
            Constraints.Builder()
                .setRequiredNetworkType(NetworkType.CONNECTED)
                .build()
        ).build()
        
        workManager.enqueueUniquePeriodicWork(
            "event_sync",
            ExistingPeriodicWorkPolicy.KEEP,
            syncRequest
        )
    }
}
```

### Event Reminders
```kotlin
// Notification Service
class NotificationService @Inject constructor(
    private val context: Context
) {
    fun scheduleEventReminder(event: Event, minutesBefore: Int) {
        val notificationTime = event.startTime.minus(minutesBefore, ChronoUnit.MINUTES)
        
        val workRequest = OneTimeWorkRequestBuilder<ReminderWorker>()
            .setInputData(
                workDataOf(
                    "event_id" to event.id,
                    "event_title" to event.title
                )
            )
            .setInitialDelay(
                Duration.between(Instant.now(), notificationTime)
            )
            .build()
            
        WorkManager.getInstance(context).enqueue(workRequest)
    }
}
```

## Testing Strategy

### Unit Tests
```kotlin
@Test
fun `getEvents returns cached data when network unavailable`() = runTest {
    // Given
    val cachedEvents = listOf(testEvent1, testEvent2)
    coEvery { localDataSource.getEvents(any(), any()) } returns cachedEvents
    coEvery { remoteDataSource.getEvents(any(), any()) } throws NetworkException()
    
    // When
    val result = repository.getEvents(startDate, endDate).first()
    
    // Then
    assertEquals(cachedEvents, result)
}
```

### UI Tests
```kotlin
@Test
fun calendar_displaysEventsCorrectly() {
    composeTestRule.setContent {
        ZCalendarTheme {
            CalendarScreen()
        }
    }
    
    // Verify events are displayed
    composeTestRule
        .onNodeWithText("Team Meeting")
        .assertIsDisplayed()
}
```

## Build Configuration

### Gradle Setup
```kotlin
// app/build.gradle.kts
android {
    compileSdk = 33
    
    defaultConfig {
        applicationId = "com.zcalendar.android"
        minSdk = 24
        targetSdk = 33
        versionCode = 1
        versionName = "1.0"
    }
    
    buildFeatures {
        compose = true
    }
    
    composeOptions {
        kotlinCompilerExtensionVersion = "1.4.0"
    }
}

dependencies {
    implementation("androidx.compose.ui:ui:$compose_version")
    implementation("androidx.compose.material3:material3:$material3_version")
    implementation("androidx.hilt:hilt-android:$hilt_version")
    implementation("androidx.room:room-runtime:$room_version")
    implementation("com.squareup.retrofit2:retrofit:$retrofit_version")
}
```

## Security and Privacy

### Data Protection
- **Encrypted Storage**: Room database encryption
- **Network Security**: Certificate pinning
- **Authentication**: Secure token storage

### Permissions
```xml
<!-- Required Permissions -->
<uses-permission android:name="android.permission.INTERNET" />
<uses-permission android:name="android.permission.READ_CALENDAR" />
<uses-permission android:name="android.permission.WRITE_CALENDAR" />
<uses-permission android:name="android.permission.POST_NOTIFICATIONS" />
```