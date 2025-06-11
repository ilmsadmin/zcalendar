# Android App Mockups - ZCalendar

## Overview
The Android app follows Material Design 3 guidelines and uses Jetpack Compose for modern, native Android experiences.

## Screen Specifications
- Phone: 360x800dp (Pixel 4/5 equivalent)
- Tablet: 840x1400dp (tablet landscape support)
- Edge-to-edge design with proper insets handling

## Screen Flow

### App Launch & Onboarding

#### 1. Splash Screen
```
┌─────────────────────────────────┐  Android 360x800dp
│ 🔋📶📶 9:41 AM                   │  ← Status bar (24dp)
│                                 │
│                                 │
│                                 │
│              📅                 │  ← App icon (72dp)
│                                 │
│           ZCalendar             │  ← App name (Roboto Medium, 24sp)
│                                 │
│       Your Personal Calendar    │  ← Tagline (Roboto, 16sp)
│                                 │
│                                 │
│                                 │
│           ●●●○○                 │  ← Material loading indicator
│                                 │
│                                 │
│                                 │
│                                 │
│             48dp nav            │  ← Navigation bar
└─────────────────────────────────┘
```

#### 2. Onboarding - Welcome
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
│                                 │
│                                 │
│              📅                 │  ← Large app icon (96dp)
│                                 │
│          Welcome to             │  ← Roboto Medium, 32sp
│           ZCalendar             │
│                                 │
│     Organize your life with     │  ← Roboto, 16sp
│     a beautiful, privacy-       │
│     focused calendar app        │
│                                 │
│                                 │
│  ┌───────────────────────────┐  │  ← Material Button
│  │      CONTINUE SETUP       │  │     (40dp height)
│  └───────────────────────────┘  │
│                                 │
│          Skip for now           │  ← Text button
│                                 │
│             48dp nav            │
└─────────────────────────────────┘
```

#### 3. Onboarding - Permissions
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM          ✕ Skip   │  ← Status bar + skip action
├─────────────────────────────────┤
│              🔔                 │  ← Icon (48dp)
│                                 │
│         Notifications           │  ← Roboto Medium, 24sp
│                                 │
│    Get reminders for your       │  ← Body text (Roboto, 16sp)
│    important events so you      │
│    never miss a meeting         │
│                                 │
│                                 │
│  ┌───────────────────────────┐  │
│  │   ENABLE NOTIFICATIONS    │  │  ← Primary button
│  └───────────────────────────┘  │
│                                 │
│           Not now               │  ← Text button
│                                 │
│                                 │
│         ●●●○○                   │  ← Progress dots
│                                 │
│             48dp nav            │
└─────────────────────────────────┘
```

#### 4. Onboarding - Calendar Access
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM     ← Setup       │  ← Back navigation
├─────────────────────────────────┤
│              📱                 │  ← Icon
│                                 │
│      System Calendar            │  ← Title
│                                 │
│  Import events from your        │  ← Description
│  Android calendar to get        │
│  started quickly (optional)     │
│                                 │
│                                 │
│  ┌───────────────────────────┐  │
│  │     IMPORT CALENDAR       │  │  ← Primary action
│  └───────────────────────────┘  │
│                                 │
│       Create from scratch       │  ← Secondary action
│                                 │
│                                 │
│         ●●●●○                   │  ← Progress
│                                 │
│             48dp nav            │
└─────────────────────────────────┘
```

### Main App Interface

#### 5. Calendar - Month View (Primary Screen)
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │  ← Status bar
├─────────────────────────────────┤
│ ☰ ZCalendar           🔍 ⋮      │  ← Top app bar (64dp)
├─────────────────────────────────┤
│ ← Nov    December 2023    Jan → │  ← Month navigation (48dp)
├─────────────────────────────────┤
│ S  M  T  W  T  F  S             │  ← Week header (32dp)
├───┬──┬──┬──┬──┬──┬──────────────┤
│   │  │  │  │  │1 │2             │  ← Calendar grid
├───┼──┼──┼──┼──┼──┼──────────────┤      (48dp per row)
│ 3 │4 │5 │6 │7 │8 │9             │
├───┼──┼──┼──┼──┼──┼──────────────┤
│10 │11│12│13│🔵│15│16            │  ← Today (14) highlighted
├───┼──┼──┼──┼─●┼──┼──────────────┤      • = Event indicator
│17 │18│19│20│21│22│23            │
├───┼──┼──┼──┼──┼──┼──────────────┤
│24 │25│26│27│28│29│30            │
├───┼──┼──┼──┼──┼──┼──────────────┤
│31 │  │  │  │  │  │              │
├─────────────────────────────────┤
│         Today's Events          │  ← Section header (Material3)
├─────────────────────────────────┤
│ ▌ 9:00 AM - 10:00 AM     🔔     │  ← Event card (72dp)
│   Team Meeting                 │     Material Design card
│   📍 Conference Room A          │     Roboto, 14sp
├─────────────────────────────────┤
│ ▌ 12:00 PM - 1:00 PM           │  ← Another event
│   Lunch with Client             │     Different color accent
│   📍 Restaurant XYZ             │
├─────────────────────────────────┤
│                                 │
│  📅   📋   ➕   ⚙️               │  ← Bottom navigation (80dp)
│Calendar Events Add Settings    │     Material3 styled
└─────────────────────────────────┘
```

#### 6. Calendar - Week View
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ← Calendar      Week     🔍 ⋮   │  ← App bar with back
├─────────────────────────────────┤
│ Week of Dec 11-17, 2023         │  ← Week range header
├─────────────────────────────────┤
│     │11│12│13│🔵│15│16│17        │  ← Date header (40dp)
│     │M │T │W │T │F │S │S         │     Material chips for dates
├─────┼──┼──┼──┼──┼──┼──┼─────────┤
│ 8   │  │  │  │██│  │  │         │  ← Hour grid
│     │  │  │  │Mtg  │  │         │     (48dp per hour)
│ 9   │  │  │  │██│  │  │         │
├─────┼──┼──┼──┼──┼──┼──┼─────────┤
│10   │  │██│  │  │  │  │         │
│     │  │Cal │  │  │  │         │
│11   │  │  │  │  │  │  │         │
├─────┼──┼──┼──┼──┼──┼──┼─────────┤
│12   │  │  │  │  │  │  │███      │
│     │  │  │  │  │  │  │Lun      │
│13   │  │  │  │  │  │  │         │
├─────┼──┼──┼──┼──┼──┼──┼─────────┤
│14   │  │  │  │  │  │  │         │
│ :   │: │: │: │: │: │: │:        │  ← Scrollable content
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️               │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 7. Calendar - Day View
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ← Week    Today, Dec 14   🔍 ⋮  │
├─────────────────────────────────┤
│ All Day  ████████████████████   │  ← All-day chip (48dp)
│          Holiday Party          │
├─────┬───────────────────────────┤
│ 8   │                          │  ← Timeline (56dp rows)
├─────┼───────────────────────────┤
│ 9   │ ████████████████████████  │
│     │ Team Meeting             │
│     │ 📍 Conference Room A      │  ← Event with details
│     │ 👥 5 attendees           │
├─────┼───────────────────────────┤
│10   │                          │
├─────┼───────────────────────────┤
│11   │                          │
├─────┼───────────────────────────┤
│12   │ ████████████████████████  │
│     │ Lunch with Client        │
│     │ 📍 Restaurant XYZ         │
├─────┼───────────────────────────┤
│13   │                          │
├─────┼───────────────────────────┤
│14   │                          │
│ :   │            :             │  ← Scrollable timeline
│ :   │            :             │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️               │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 8. Events List View
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ☰ Events              🔍 ➕ ⋮   │  ← App bar with search/add
├─────────────────────────────────┤
│ 📅 Today • December 14, 2023    │  ← Date filter chip (40dp)
├─────────────────────────────────┤
│ ▌ 9:00 AM - 10:00 AM      📍🔔  │  ← Event card (80dp)
│   Team Meeting                 │     Material card design
│   Conference Room A             │
│   👥 John, Sarah, Mike          │  ← Metadata chips
├─────────────────────────────────┤
│ ▌ 12:00 PM - 1:00 PM      📍   │
│   Lunch with Client             │
│   Restaurant XYZ                │
│   💰 Business                   │  ← Category chip
├─────────────────────────────────┤
│ ▌ 6:00 PM - 8:00 PM            │
│   Gym Workout                   │
│   📍 Local Fitness Center       │
│   🔁 Mon, Wed, Fri              │  ← Recurring chip
├─────────────────────────────────┤
│                                 │
│     No more events today        │  ← Empty state
│     Tap + to add an event       │
│                                 │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️               │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 9. Create Event Screen
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ✕ Close       New Event   SAVE  │  ← Modal app bar
├─────────────────────────────────┤
│ Title                           │  ← Material text fields
│ ┌─────────────────────────────┐ │     (56dp height)
│ │ Team Meeting                │ │
│ └─────────────────────────────┘ │
│                                 │
│ Date & Time                     │
│ ┌─────────────┐ ┌─────────────┐ │  ← Material date/time chips
│ │ Dec 14, 2023│ │ 9:00-10:00  │ │
│ └─────────────┘ └─────────────┘ │
│                                 │
│ All Day                    ●    │  ← Material switch
│                                 │
│ Calendar                        │
│ ┌─────────────────────────────┐ │
│ │ 🔵 Personal              ▼ │ │  ← Dropdown menu
│ └─────────────────────────────┘ │
│                                 │
│ Location                        │
│ ┌─────────────────────────────┐ │
│ │ 📍 Conference Room A        │ │
│ └─────────────────────────────┘ │
│                                 │
│ Attendees                       │
│ ┌─────────────────────────────┐ │
│ │ + Add attendees...          │ │
│ └─────────────────────────────┘ │
│ ┌─────────────────────────────┐ │  ← Attendee chips
│ │ 👤 john@company.com      ✕  │ │
│ └─────────────────────────────┘ │
│ ┌─────────────────────────────┐ │
│ │ 👤 sarah@company.com     ✕  │ │
│ └─────────────────────────────┘ │
│                                 │
│ Notes                           │
│ ┌─────────────────────────────┐ │  ← Multiline text field
│ │ Discuss Q4 goals and        │ │     (80dp height)
│ │ planning for next quarter   │ │
│ └─────────────────────────────┘ │
│                                 │
│ ⚠️ Reminder      🔁 Repeat      │  ← Setting chips
│ 15 min before   Never           │
│                                 │
│             48dp nav            │
└─────────────────────────────────┘
```

#### 10. Event Detail View
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ← Back         Event        ⋮   │  ← App bar with overflow menu
├─────────────────────────────────┤
│ ▌ Team Meeting                  │  ← Event header card
│                                 │     (Roboto Medium, 24sp)
├─────────────────────────────────┤
│ 🕘 December 14, 2023            │  ← Date/time card
│    9:00 AM - 10:00 AM           │     Material design
│    Duration: 1 hour             │
├─────────────────────────────────┤
│ 📍 Conference Room A            │  ← Location card
│    123 Main St, Building B     │
│    VIEW ON MAP                  │  ← Material button
├─────────────────────────────────┤
│ 🔵 Personal Calendar            │  ← Calendar info
├─────────────────────────────────┤
│ 👥 Attendees (3)                │  ← Attendees section
│    👤 John Smith (Organizer)    │     Material list items
│    👤 Sarah Johnson             │
│    👤 Mike Chen                 │
│    + Add attendees              │  ← Action item
├─────────────────────────────────┤
│ 📝 Notes                        │  ← Notes card
│    Discuss Q4 goals and         │
│    planning for next quarter.   │
│    Review budget proposals      │
│    and team assignments.        │
├─────────────────────────────────┤
│ ⚠️ Reminder: 15 minutes before  │  ← Settings info
│ 🔁 Repeats: Never               │
├─────────────────────────────────┤
│ ┌─────────────┐ ┌─────────────┐ │  ← Action buttons
│ │    EDIT     │ │   DELETE    │ │     Material buttons
│ └─────────────┘ └─────────────┘ │
│                                 │
│             48dp nav            │
└─────────────────────────────────┘
```

#### 11. Settings Screen
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ☰ Settings                 🔍   │  ← App bar
├─────────────────────────────────┤
│ 👤 Account                      │  ← Section header
├─────────────────────────────────┤
│ 📱 Guest Mode                   │  ← List item (64dp)
│    Tap to create account        │     Material list style
│                             →   │
├─────────────────────────────────┤
│ 📅 Calendar Settings            │  ← Section divider
├─────────────────────────────────┤
│ Default Calendar                │  ← Setting items (56dp)
│ Personal                    →   │
├─────────────────────────────────┤
│ Week Start Day                  │
│ Sunday                      →   │
├─────────────────────────────────┤
│ Time Format                     │
│ 12-hour                     →   │
├─────────────────────────────────┤
│ 🔔 Notifications                │  ← Section
├─────────────────────────────────┤
│ Event Reminders                 │  ← Switch items
│                             ●   │  ← Material switch (on)
├─────────────────────────────────┤
│ Default Reminder                │
│ 15 minutes                  →   │
├─────────────────────────────────┤
│ All-Day Events                  │
│                             ○   │  ← Switch (off)
├─────────────────────────────────┤
│ 🎨 Appearance                   │  ← Section
├─────────────────────────────────┤
│ Theme                           │
│ System default              →   │
├─────────────────────────────────┤
│ Dynamic Colors                  │  ← Android 12+ feature
│                             ●   │
├─────────────────────────────────┤
│ ⚡ Advanced                     │  ← Section
├─────────────────────────────────┤
│ Export Calendar             →   │  ← Action items
├─────────────────────────────────┤
│ Import Calendar             →   │
├─────────────────────────────────┤
│ Storage Usage               →   │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️               │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 12. Account Creation
```
┌─────────────────────────────────┐
│ 🔋📶📶 9:41 AM                   │
├─────────────────────────────────┤
│ ✕ Close    Create Account       │  ← Modal app bar
├─────────────────────────────────┤
│           🔐                    │  ← Icon (64dp)
│                                 │
│      Unlock Premium Features    │  ← Title (Roboto Medium, 20sp)
│                                 │
│ • Sync across all devices       │  ← Benefits list
│ • Cloud backup & restore        │     Material list items
│ • Calendar sharing              │
│ • Premium integrations          │
│                                 │
│ Email                           │  ← Form fields
│ ┌─────────────────────────────┐ │     Material text fields
│ │ your@email.com              │ │     (56dp height)
│ └─────────────────────────────┘ │
│                                 │
│ Password                        │
│ ┌─────────────────────────────┐ │
│ │ ●●●●●●●●●●●●                │ │
│ └─────────────────────────────┘ │
│                                 │
│ Confirm Password                │
│ ┌─────────────────────────────┐ │
│ │ ●●●●●●●●●●●●                │ │
│ └─────────────────────────────┘ │
│                                 │
│ ┌─────────────────────────────┐ │  ← Material primary button
│ │      CREATE ACCOUNT         │ │     (40dp height)
│ └─────────────────────────────┘ │
│                                 │
│ Your data stays private and     │  ← Caption text
│ encrypted. Terms & Privacy      │     (Roboto, 12sp)
│                                 │
│             48dp nav            │
└─────────────────────────────────┘
```

## Material Design Features

### Navigation Patterns
- **Bottom Navigation**: Primary navigation with Material3 styling
- **Navigation Drawer**: Secondary navigation for tablet layouts
- **Back Stack**: Proper Fragment/Activity navigation
- **Deep Linking**: Support for calendar event URLs

### Material Components
- **Cards**: Elevated surfaces with proper shadows
- **Chips**: For filters, tags, and selections
- **FAB**: Floating Action Button for quick event creation
- **Switches**: Material switch design for settings
- **Text Fields**: Outlined style with proper states

### Gesture Support
- **Swipe to refresh**: Pull-down gesture for data sync
- **Swipe between views**: Calendar navigation
- **Swipe actions**: Event list item actions
- **Edge gestures**: Back navigation (Android 10+)

### Adaptive Features
- **Dynamic Color**: Android 12+ Material You theming
- **Edge-to-edge**: Full screen with proper insets
- **Foldable support**: Responsive layouts for foldable devices
- **Large screens**: Tablet-optimized layouts

### Animation & Motion
- **Shared Element Transitions**: Between calendar and detail views
- **Material Motion**: Standard easing curves and durations
- **Predictive Back**: Android 13+ back gesture preview
- **Hero Animations**: Event creation to calendar

### Accessibility
- **TalkBack**: Complete screen reader support
- **Large Text**: Dynamic text scaling support
- **High Contrast**: Material contrast guidelines
- **Focus Management**: Proper keyboard navigation
- **Touch Targets**: Minimum 48dp touch areas

### Dark Theme
- **Material Dark**: Proper dark theme implementation
- **System Integration**: Follows system theme settings
- **Battery Optimization**: True black backgrounds for OLED
- **Contrast Ratios**: Proper text contrast in dark mode