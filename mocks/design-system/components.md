# ZCalendar Component Library

## Navigation Components

### Tab Bar (iOS/Android/Web)

#### iOS Implementation
```
Tab Bar Structure:
┌─────────────────────────────────────────────┐
│  📅      📋      ➕      ⚙️                   │
│ Calendar  Events  Create  Settings           │
└─────────────────────────────────────────────┘

- Height: 83pt (49pt + safe area)
- Background: System background with blur
- Active state: Primary color
- Inactive state: Secondary text color
- Icons: SF Symbols, 25pt
- Labels: SF Pro Text, 10pt
```

#### Android Implementation
```
Bottom Navigation:
┌─────────────────────────────────────────────┐
│  📅      📋      ➕      ⚙️                   │
│ Calendar  Events   Add   Settings           │
└─────────────────────────────────────────────┘

- Height: 80dp
- Background: Surface color
- Elevation: 8dp
- Active state: Primary color
- Inactive state: On-surface variant
- Icons: Material Icons, 24dp
- Labels: Roboto, 12sp
```

#### Web Implementation
```
Sidebar Navigation (Desktop):
┌─────────────┬───────────────────────────────┐
│ 📅 Calendar │                               │
│ 📋 Events   │                               │
│ ➕ Create   │         Main Content          │
│ ⚙️ Settings │                               │
│             │                               │
└─────────────┴───────────────────────────────┘

Mobile: Same as mobile apps with bottom navigation
```

### Top Navigation

#### iOS Navigation Bar
```
┌─────────────────────────────────────────────┐
│ ← Back     Calendar Title          + More   │
└─────────────────────────────────────────────┘

- Height: 44pt + status bar
- Large title: 34pt when scrolled to top
- Background: System background with blur
- Title: SF Pro Display, 22pt or 34pt
- Back button: SF Symbol chevron.left
- Right buttons: SF Symbols, 22pt
```

#### Android App Bar
```
┌─────────────────────────────────────────────┐
│ ☰ Menu     Calendar Title          ⋮ More   │
└─────────────────────────────────────────────┘

- Height: 64dp
- Background: Primary color or Surface
- Title: Roboto Medium, 20sp
- Navigation icon: 24dp
- Action icons: 24dp
- Elevation: 4dp (if not using edge-to-edge)
```

#### Web Header
```
Desktop:
┌─────────────────────────────────────────────┐
│ ZCalendar Logo    Navigation    User Avatar │
└─────────────────────────────────────────────┘

Mobile: Same as app implementations
```

## Calendar Components

### Month View Grid

```
Month View Layout:
┌─────────────────────────────────────────────┐
│              December 2023                  │
├─────┬─────┬─────┬─────┬─────┬─────┬─────────┤
│ Sun │ Mon │ Tue │ Wed │ Thu │ Fri │ Sat     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│     │     │     │     │     │  1  │   2     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│  3  │  4  │  5  │  6  │  7  │  8  │   9     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│ 10  │ 11  │ 12  │ 13  │ 14  │ 15  │  16     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│ 17  │ 18  │ 19  │ 20  │ 21  │ 22  │  23     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│ 24  │ 25  │ 26  │ 27  │ 28  │ 29  │  30     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│ 31  │     │     │     │     │     │         │
└─────┴─────┴─────┴─────┴─────┴─────┴─────────┘

Date Cell States:
• Normal: Light gray text
• Today: Blue circle background, white text
• Selected: Primary color background
• Has Events: Small colored dots below date
• Other month: Faded text color
```

### Week View

```
Week View Layout:
┌─────────────────────────────────────────────┐
│        Week of December 11-17, 2023         │
├─────┬───────────────────────────────────────┤
│Time │ Mon  Tue  Wed  Thu  Fri  Sat  Sun     │
├─────┼───────────────────────────────────────┤
│ 8AM │  │    │    │ ████ │    │    │        │
│     │  │    │    │Team │    │    │        │
│ 9AM │  │    │    │Meet │    │    │        │
├─────┼───────────────────────────────────────┤
│10AM │  │████│    │ ████ │    │    │        │
│     │  │Call│    │     │    │    │        │
│11AM │  │    │    │     │    │    │        │
├─────┼───────────────────────────────────────┤
│12PM │  │    │    │     │    │    │ ██████  │
│     │  │    │    │     │    │    │ Lunch  │
│ 1PM │  │    │    │     │    │    │        │
└─────┴───────────────────────────────────────┘

Features:
- Hour markers on left
- Events as colored blocks
- All-day events at top
- Scroll for more hours
- Tap empty space to create event
```

### Day View

```
Day View Layout:
┌─────────────────────────────────────────────┐
│           Thursday, December 14             │
├─────────────────────────────────────────────┤
│ All Day: ████ Team Building Event ████     │
├─────┬───────────────────────────────────────┤
│ 8AM │                                     │
├─────┼───────────────────────────────────────┤
│ 9AM │ ████████████████████████████████     │
│     │ Team Meeting                        │
│     │ Conference Room A                   │
├─────┼───────────────────────────────────────┤
│10AM │                                     │
├─────┼───────────────────────────────────────┤
│11AM │                                     │
├─────┼───────────────────────────────────────┤
│12PM │ ████████████████████████████████     │
│     │ Lunch with Client                   │
│     │ Restaurant XYZ                      │
├─────┼───────────────────────────────────────┤
│ 1PM │                                     │
└─────┴───────────────────────────────────────┘

Features:
- Full day timeline
- Detailed event information
- Location display
- Easy event creation
```

## Event Components

### Event Card

```
Event Card Layout:
┌─────────────────────────────────────────────┐
│ ▌ Team Meeting                    🔔 ⚠️     │
│   9:00 AM - 10:00 AM                       │
│   📍 Conference Room A                      │
│   👥 5 attendees                            │
│   📝 Discuss Q4 goals and planning...      │
└─────────────────────────────────────────────┘

Elements:
• Left border: Calendar color (4px)
• Title: Bold, primary text color
• Time: Secondary text color
• Location icon + text: Tertiary text
• Attendees: Small icons or count
• Description: Truncated if too long
• Right icons: Reminder, Alert status
```

### Event Creation Form

```
Create Event Form:
┌─────────────────────────────────────────────┐
│ ✕ Cancel          New Event         ✓ Save │
├─────────────────────────────────────────────┤
│ Title                                       │
│ ┌─────────────────────────────────────────┐ │
│ │ Enter event title...                    │ │
│ └─────────────────────────────────────────┘ │
│                                             │
│ Date & Time                                 │
│ ┌─────────────────┐ ┌─────────────────────┐ │
│ │ Dec 14, 2023    │ │ 9:00 AM - 10:00 AM  │ │
│ └─────────────────┘ └─────────────────────┘ │
│                                             │
│ Calendar                                    │
│ ┌─────────────────────────────────────────┐ │
│ │ 🔵 Personal                             │ │
│ └─────────────────────────────────────────┘ │
│                                             │
│ Location                                    │
│ ┌─────────────────────────────────────────┐ │
│ │ 📍 Add location...                      │ │
│ └─────────────────────────────────────────┘ │
│                                             │
│ Notes                                       │
│ ┌─────────────────────────────────────────┐ │
│ │ Add notes...                            │ │
│ │                                         │ │
│ └─────────────────────────────────────────┘ │
│                                             │
│ ⚠️ Reminder          🔁 Repeat              │
│ 15 minutes before    Never                  │
└─────────────────────────────────────────────┘
```

### Event List Item

```
List Item Layout:
┌─────────────────────────────────────────────┐
│ ▌ 9:00 AM   Team Meeting              📍    │
│            Conference Room A                │
└─────────────────────────────────────────────┘

Compact version:
┌─────────────────────────────────────────────┐
│ ▌ 9:00   Team Meeting                       │
└─────────────────────────────────────────────┘

States:
• Normal: Default styling
• Past: Faded appearance
• Ongoing: Highlighted border
• Conflicting: Warning color
```

## Input Components

### Text Fields

```
Standard Text Field:
┌─────────────────────────────────────────────┐
│ Event Title                                 │
│ ┌─────────────────────────────────────────┐ │
│ │ Team Meeting                            │ │
│ └─────────────────────────────────────────┘ │
└─────────────────────────────────────────────┘

Error State:
┌─────────────────────────────────────────────┐
│ Event Title                                 │
│ ┌─────────────────────────────────────────┐ │
│ │ Team Meeting                            │ │ 🔴
│ └─────────────────────────────────────────┘ │
│ ⚠️ Title is required                        │
└─────────────────────────────────────────────┘
```

### Date/Time Pickers

```
iOS Date Picker:
┌─────────────────────────────────────────────┐
│             Select Date                     │
├─────────────────────────────────────────────┤
│    Dec  │    14   │   2023                 │
│    ↕    │    ↕    │    ↕                   │
│    Nov  │    13   │   2022                 │
│    Oct  │    12   │   2021                 │
│    Sep  │    11   │   2020                 │
└─────────────────────────────────────────────┘

Android Date Picker:
┌─────────────────────────────────────────────┐
│           December 2023                     │
├─────┬─────┬─────┬─────┬─────┬─────┬─────────┤
│  S  │  M  │  T  │  W  │  T  │  F  │   S     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│     │     │     │     │     │  1  │   2     │
│  3  │  4  │  5  │  6  │  7  │  8  │   9     │
│ 10  │ 11  │ 12  │ 13  │ 🔵  │ 15  │  16     │
│ 17  │ 18  │ 19  │ 20  │ 21  │ 22  │  23     │
│ 24  │ 25  │ 26  │ 27  │ 28  │ 29  │  30     │
│ 31  │     │     │     │     │     │         │
└─────┴─────┴─────┴─────┴─────┴─────┴─────────┘
```

### Dropdowns/Selectors

```
Calendar Selector:
┌─────────────────────────────────────────────┐
│ 🔵 Personal               ▼               │
├─────────────────────────────────────────────┤
│ 🔵 Personal                                │
│ 🟣 Work                                    │
│ 🟢 Health                                 │
│ 🟠 Travel                                 │
│ + Create New Calendar                      │
└─────────────────────────────────────────────┘
```

## Status Components

### Loading States

```
Loading Calendar:
┌─────────────────────────────────────────────┐
│              December 2023                  │
├─────┬─────┬─────┬─────┬─────┬─────┬─────────┤
│ Sun │ Mon │ Tue │ Wed │ Thu │ Fri │ Sat     │
├─────┼─────┼─────┼─────┼─────┼─────┼─────────┤
│ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░     │ ← Shimmer
│ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░     │   Effect
│ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░ │ ░░░     │
└─────┴─────┴─────┴─────┴─────┴─────┴─────────┘

Loading Events:
┌─────────────────────────────────────────────┐
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░             │
│ ░░░░░░░░░░░░░░░░░░░░                         │
├─────────────────────────────────────────────┤
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░             │
│ ░░░░░░░░░░░░░░░░░░░░                         │
└─────────────────────────────────────────────┘
```

### Empty States

```
No Events:
┌─────────────────────────────────────────────┐
│                                             │
│                    📅                       │
│                                             │
│            No events today                  │
│         Tap + to create your first          │
│                  event                      │
│                                             │
│            [Create Event]                   │
│                                             │
└─────────────────────────────────────────────┘

Offline Mode:
┌─────────────────────────────────────────────┐
│ 📡 Offline Mode                     ⟲       │
└─────────────────────────────────────────────┘
```

### Error States

```
Network Error:
┌─────────────────────────────────────────────┐
│                    ⚠️                        │
│                                             │
│          Connection Failed                  │
│      Unable to sync your calendar          │
│                                             │
│            [Try Again]                      │
│                                             │
└─────────────────────────────────────────────┘

Sync Conflict:
┌─────────────────────────────────────────────┐
│ ⚠️ Sync Conflict                            │
│ This event was modified on another device   │
│                                             │
│ [Keep Local] [Use Remote] [Merge]           │
└─────────────────────────────────────────────┘
```

## Dialog Components

### Alert Dialogs

```
Delete Confirmation:
┌─────────────────────────────────────────────┐
│              Delete Event?                  │
│                                             │
│  This action cannot be undone. The event   │
│  will be permanently deleted.               │
│                                             │
│              [Cancel] [Delete]              │
└─────────────────────────────────────────────┘

Recurring Event Dialog:
┌─────────────────────────────────────────────┐
│         Delete Recurring Event              │
│                                             │
│  What would you like to delete?             │
│                                             │
│  • This event only                         │
│  • This and future events                  │
│  • All events in the series                │
│                                             │
│              [Cancel] [Delete]              │
└─────────────────────────────────────────────┘
```

### Bottom Sheets (Android/Web)

```
Event Actions:
┌─────────────────────────────────────────────┐
│ ○○○                                         │ ← Handle
├─────────────────────────────────────────────┤
│                                             │
│ 🔘 Edit Event                               │
│ 🔄 Duplicate Event                          │
│ 📤 Share Event                              │
│ 🗑️ Delete Event                             │
│                                             │
└─────────────────────────────────────────────┘
```

## Accessibility Features

### Focus Indicators
- Clear visual focus rings
- High contrast support
- Keyboard navigation paths

### Screen Reader Support
- Proper labels and descriptions
- Semantic markup
- Live region updates

### Touch Targets
- Minimum 44pt/44dp size
- Adequate spacing between elements
- Easy one-handed operation

### Color Independence
- Icons or patterns for important information
- Text alternatives for color-coded data
- High contrast mode support