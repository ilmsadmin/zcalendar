# iOS App Mockups - ZCalendar

## Overview
The iOS app follows Apple's Human Interface Guidelines and uses SwiftUI for modern, native iOS experiences.

## Screen Specifications
- iPhone: 375x812pt (iPhone 13/14)
- iPad: 834x1194pt (iPad Air)
- Safe areas: Properly handled for notch and home indicator

## Screen Flow

### App Launch & Onboarding

#### 1. Splash Screen
```
┌─────────────────────────────────┐  iPhone 375x812pt
│                                 │  
│            44pt status          │  ← Status Bar
│                                 │
│                                 │
│                                 │
│              📅                 │  ← App Icon (100pt)
│                                 │
│           ZCalendar             │  ← App Name (SF Pro Display, 34pt)
│                                 │
│       Your Personal Calendar    │  ← Tagline (SF Pro Text, 17pt)
│                                 │
│                                 │
│                                 │
│          ●●●○○                  │  ← Loading indicator
│                                 │
│                                 │
│                                 │
│            34pt home            │  ← Home indicator
└─────────────────────────────────┘
```

#### 2. Onboarding - Welcome
```
┌─────────────────────────────────┐
│                                 │
│            44pt status          │
│                                 │
│                                 │
│              📅                 │  ← Large app icon (120pt)
│                                 │
│          Welcome to             │  ← SF Pro Display, 34pt
│           ZCalendar             │
│                                 │
│     Organize your life with     │  ← SF Pro Text, 17pt
│     a beautiful, privacy-       │
│     focused calendar app        │
│                                 │
│                                 │
│     ┌─────────────────────────┐ │  ← Primary button
│     │    Continue Setup       │ │     (50pt height)
│     └─────────────────────────┘ │
│                                 │
│        Skip for now             │  ← Text button
│                                 │
│            34pt home            │
└─────────────────────────────────┘
```

#### 3. Onboarding - Permissions
```
┌─────────────────────────────────┐
│ ✕ Skip          Setup          │  ← Navigation bar
├─────────────────────────────────┤
│                                 │
│              🔔                 │  ← Icon (60pt)
│                                 │
│        Notifications            │  ← Title (SF Pro Display, 28pt)
│                                 │
│    Get reminders for your       │  ← Body text (SF Pro Text, 17pt)
│    important events so you      │
│    never miss a meeting         │
│                                 │
│                                 │
│     ┌─────────────────────────┐ │
│     │   Enable Notifications  │ │  ← Primary button
│     └─────────────────────────┘ │
│                                 │
│           Not now               │  ← Secondary action
│                                 │
│                                 │
│         ●●●○○                   │  ← Progress indicator
│                                 │
│            34pt home            │
└─────────────────────────────────┘
```

#### 4. Onboarding - Calendar Access
```
┌─────────────────────────────────┐
│ ← Back          Setup           │
├─────────────────────────────────┤
│                                 │
│              📱                 │  ← Icon
│                                 │
│      System Calendar            │  ← Title
│                                 │
│  Import events from your iOS    │  ← Description
│  calendar to get started        │
│  quickly (optional)             │
│                                 │
│                                 │
│     ┌─────────────────────────┐ │
│     │    Import Calendar      │ │  ← Primary action
│     └─────────────────────────┘ │
│                                 │
│       Create from scratch       │  ← Secondary action
│                                 │
│                                 │
│         ●●●●○                   │  ← Progress
│                                 │
│            34pt home            │
└─────────────────────────────────┘
```

### Main App Interface

#### 5. Calendar - Month View (Primary Screen)
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │  ← Status bar
├─────────────────────────────────┤
│ ☰ Menu     Calendar     + ⚙️   │  ← Navigation (64pt)
├─────────────────────────────────┤
│ ← November  December 2023  Jan → │  ← Month navigation (44pt)
├─────────────────────────────────┤
│ Sun Mon Tue Wed Thu Fri Sat     │  ← Week header (32pt)
├───┬───┬───┬───┬───┬───┬─────────┤
│   │   │   │   │   │ 1 │ 2       │  ← Calendar grid
├───┼───┼───┼───┼───┼───┼─────────┤      (48pt per row)
│ 3 │ 4 │ 5 │ 6 │ 7 │ 8 │ 9       │
├───┼───┼───┼───┼───┼───┼─────────┤
│10 │11 │12 │13 │🔵 │15 │16       │  ← Today (14) highlighted
├───┼───┼───┼───┼─●─┼───┼─────────┤      • = Event dot
│17 │18 │19 │20 │21 │22 │23       │
├───┼───┼───┼───┼───┼───┼─────────┤
│24 │25 │26 │27 │28 │29 │30       │
├───┼───┼───┼───┼───┼───┼─────────┤
│31 │   │   │   │   │   │         │
├─────────────────────────────────┤
│         Today's Events          │  ← Section header (24pt)
├─────────────────────────────────┤
│ ▌ 9:00 AM - 10:00 AM    🔔     │  ← Event card (60pt)
│   Team Meeting                 │     Blue left border
│   📍 Conference Room A          │     SF Pro Text, 15pt
├─────────────────────────────────┤
│ ▌ 12:00 PM - 1:00 PM           │  ← Another event
│   Lunch with Client             │     Purple left border
│   📍 Restaurant XYZ             │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️              │  ← Tab bar (83pt)
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 6. Calendar - Week View
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ ← Calendar      Week View    ⚙️ │  ← Back to month view
├─────────────────────────────────┤
│ Week of Dec 11-17, 2023         │  ← Week range
├─────────────────────────────────┤
│     │11 │12 │13 │🔵 │15 │16 │17 │  ← Date header (40pt)
│     │Mo │Tu │We │Th │Fr │Sa │Su │     Today highlighted
├─────┼───┼───┼───┼───┼───┼───┼───┤
│ 8AM │   │   │   │ ██│   │   │   │  ← Hour grid
│     │   │   │   │Mtg│   │   │   │     (40pt per hour)
│ 9AM │   │   │   │ ██│   │   │   │
├─────┼───┼───┼───┼───┼───┼───┼───┤
│10AM │   │██ │   │   │   │   │   │
│     │   │Cal│   │   │   │   │   │
│11AM │   │   │   │   │   │   │   │
├─────┼───┼───┼───┼───┼───┼───┼───┤
│12PM │   │   │   │   │   │   │███ │
│     │   │   │   │   │   │   │Lun│
│1PM  │   │   │   │   │   │   │   │
├─────┼───┼───┼───┼───┼───┼───┼───┤
│2PM  │   │   │   │   │   │   │   │
│     │   │   │   │   │   │   │   │
│  :  │ : │ : │ : │ : │ : │ : │ : │  ← Scrollable
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️              │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 7. Calendar - Day View
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ ← Week     Today, Dec 14     ⚙️ │
├─────────────────────────────────┤
│ All Day  ████████████████████   │  ← All-day events
│          Holiday Party           │     (40pt section)
├─────┬───────────────────────────┤
│ 8AM │                          │  ← Timeline
├─────┼───────────────────────────┤     (60pt per hour)
│ 9AM │ ████████████████████████  │
│     │ Team Meeting             │
│     │ 📍 Conference Room A      │  ← Event details
│     │ 👥 5 attendees           │
├─────┼───────────────────────────┤
│10AM │                          │
├─────┼───────────────────────────┤
│11AM │                          │
├─────┼───────────────────────────┤
│12PM │ ████████████████████████  │
│     │ Lunch with Client        │
│     │ 📍 Restaurant XYZ         │
├─────┼───────────────────────────┤
│1PM  │                          │
├─────┼───────────────────────────┤
│2PM  │                          │
│  :  │            :             │  ← Scrollable
│  :  │            :             │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️              │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 8. Events List View
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ 🔍 Search       Events      + ⚙️ │  ← Search + actions
├─────────────────────────────────┤
│ 📅 Today • December 14, 2023    │  ← Date filter (44pt)
├─────────────────────────────────┤
│ ▌ 9:00 AM - 10:00 AM     📍 🔔 │  ← Event item (80pt)
│   Team Meeting                 │
│   Conference Room A             │
│   👥 John, Sarah, Mike          │  ← Attendees
├─────────────────────────────────┤
│ ▌ 12:00 PM - 1:00 PM      📍   │
│   Lunch with Client             │
│   Restaurant XYZ                │
│   💰 Business expense           │  ← Tags/categories
├─────────────────────────────────┤
│ ▌ 6:00 PM - 8:00 PM            │
│   Gym Workout                   │
│   📍 Local Fitness Center       │
│   🔁 Repeats Mon, Wed, Fri      │  ← Recurring indicator
├─────────────────────────────────┤
│                                 │
│     No more events today        │  ← Empty state
│                                 │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️              │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 9. Create Event Screen
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ ✕ Cancel      New Event   ✓ Save │  ← Modal navigation
├─────────────────────────────────┤
│ Title                           │  ← Form sections
│ ┌─────────────────────────────┐ │     (48pt per field)
│ │ Team Meeting                │ │  ← Text input
│ └─────────────────────────────┘ │
│                                 │
│ Date & Time                     │
│ ┌─────────────┐ ┌─────────────┐ │  ← Date/time pickers
│ │ Dec 14, 2023│ │ 9:00 - 10:00│ │
│ └─────────────┘ └─────────────┘ │
│                                 │
│ All Day          ○              │  ← Toggle switch
│                                 │
│ Calendar                        │
│ ┌─────────────────────────────┐ │
│ │ 🔵 Personal              ▼ │ │  ← Dropdown
│ └─────────────────────────────┘ │
│                                 │
│ Location                        │
│ ┌─────────────────────────────┐ │
│ │ 📍 Conference Room A        │ │  ← Location input
│ └─────────────────────────────┘ │
│                                 │
│ Attendees                       │
│ ┌─────────────────────────────┐ │
│ │ + Add attendees...          │ │
│ └─────────────────────────────┘ │
│   ● john@company.com            │  ← Added attendees
│   ● sarah@company.com           │
│                                 │
│ Notes                           │
│ ┌─────────────────────────────┐ │  ← Multiline text
│ │ Discuss Q4 goals and        │ │
│ │ planning for next quarter   │ │
│ └─────────────────────────────┘ │
│                                 │
│ ⚠️ Reminder      🔁 Repeat      │  ← Additional options
│ 15 min before   Never           │
│                                 │
│            34pt home            │
└─────────────────────────────────┘
```

#### 10. Event Detail View
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ ← Back        Event         ⋮   │  ← Back to calendar, more menu
├─────────────────────────────────┤
│ ▌ Team Meeting                  │  ← Event title (blue border)
│                                 │     (SF Pro Display, 28pt)
├─────────────────────────────────┤
│ 🕘 December 14, 2023            │  ← Date/time section
│    9:00 AM - 10:00 AM           │     (SF Pro Text, 17pt)
│    Duration: 1 hour             │
├─────────────────────────────────┤
│ 📍 Conference Room A            │  ← Location
│    123 Main St, Building B     │
│    [View on Map]                │  ← Map link
├─────────────────────────────────┤
│ 🔵 Personal Calendar            │  ← Calendar source
├─────────────────────────────────┤
│ 👥 Attendees (3)                │  ← Attendees section
│    ● John Smith (Organizer)     │
│    ● Sarah Johnson              │
│    ● Mike Chen                  │
│    + Add attendees              │
├─────────────────────────────────┤
│ 📝 Notes                        │  ← Notes section
│    Discuss Q4 goals and         │
│    planning for next quarter.   │
│    Review budget proposals      │
│    and team assignments.        │
├─────────────────────────────────┤
│ ⚠️ Reminder: 15 minutes before  │  ← Reminder settings
│ 🔁 Repeats: Never               │  ← Recurrence
├─────────────────────────────────┤
│ ┌─────────────┐ ┌─────────────┐ │  ← Action buttons
│ │    Edit     │ │   Delete    │ │
│ └─────────────┘ └─────────────┘ │
│                                 │
│            34pt home            │
└─────────────────────────────────┘
```

#### 11. Settings Screen
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ ← Calendar      Settings        │
├─────────────────────────────────┤
│ 👤 Account                      │  ← Section header
├─────────────────────────────────┤
│ 📱 Guest Mode                   │  ← Current mode (48pt row)
│   Tap to create account  →      │
├─────────────────────────────────┤
│ 📅 Calendar Settings            │  ← Section
├─────────────────────────────────┤
│ Default Calendar     Personal → │  ← Setting row
├─────────────────────────────────┤
│ Week Start Day       Sunday  → │
├─────────────────────────────────┤
│ Time Format          12-hour → │
├─────────────────────────────────┤
│ First Day of Week    Sunday  → │
├─────────────────────────────────┤
│ 🔔 Notifications                │  ← Section
├─────────────────────────────────┤
│ Event Reminders         ●      │  ← Toggle (enabled)
├─────────────────────────────────┤
│ Default Reminder    15 min  → │
├─────────────────────────────────┤
│ All-Day Events          ○      │  ← Toggle (disabled)
├─────────────────────────────────┤
│ 🎨 Appearance                   │  ← Section
├─────────────────────────────────┤
│ Theme              Automatic → │
├─────────────────────────────────┤
│ App Icon           Default   → │
├─────────────────────────────────┤
│ ⚡ Advanced                     │  ← Section
├─────────────────────────────────┤
│ Export Calendar              → │
├─────────────────────────────────┤
│ Import Calendar              → │
├─────────────────────────────────┤
│ Storage Usage              → │
├─────────────────────────────────┤
│  📅   📋   ➕   ⚙️              │
│Calendar Events Add Settings    │
└─────────────────────────────────┘
```

#### 12. Account Creation
```
┌─────────────────────────────────┐
│ 20:41  📶 📶 🔋100%  Wed Dec 14 │
├─────────────────────────────────┤
│ ✕ Cancel    Create Account      │
├─────────────────────────────────┤
│           🔐                    │  ← Icon (80pt)
│                                 │
│      Unlock Premium Features    │  ← Title (SF Pro Display, 24pt)
│                                 │
│ • Sync across all devices       │  ← Benefits list
│ • Cloud backup & restore        │     (SF Pro Text, 17pt)
│ • Calendar sharing              │
│ • Premium integrations          │
│                                 │
│ Email                           │  ← Form fields
│ ┌─────────────────────────────┐ │
│ │ your@email.com              │ │
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
│ ┌─────────────────────────────┐ │  ← Create button
│ │       Create Account        │ │
│ └─────────────────────────────┘ │
│                                 │
│ Your data stays private and     │  ← Privacy note
│ encrypted. Terms & Privacy      │
│                                 │
│            34pt home            │
└─────────────────────────────────┘
```

## UI Patterns & Behaviors

### Navigation Patterns
- **Tab-based navigation**: Primary navigation via bottom tab bar
- **Hierarchical navigation**: Push/pop for drill-down screens  
- **Modal presentation**: For create/edit flows
- **Sidebar (iPad)**: Split view with sidebar navigation

### Gestures
- **Swipe**: Navigate between time periods in calendar views
- **Pull to refresh**: Update calendar data
- **Pinch to zoom**: Switch between calendar view modes
- **Long press**: Quick actions on events

### State Management
- **Loading states**: Skeleton screens and shimmer effects
- **Empty states**: Helpful messages with actions
- **Error states**: Clear error messages with retry options
- **Offline indicator**: Visual feedback when offline

### Accessibility
- **VoiceOver**: Full screen reader support with proper labels
- **Dynamic Type**: Support for user font size preferences
- **High Contrast**: Alternative color schemes
- **Reduce Motion**: Respect motion preferences
- **Voice Control**: All interactive elements are voice controllable

### Dark Mode Support
All screens automatically adapt to system dark mode with:
- Inverted background colors
- High contrast text
- Preserved accent colors for calendar categories
- Dimmed shadows and borders