# ZCalendar Design System

## Overview

The ZCalendar design system establishes a consistent visual language across iOS, Android, and Web platforms while respecting platform-specific conventions.

## Color System

### Primary Colors
```
Primary Blue:   #007AFF (iOS System Blue)
Primary Green:  #34C759 (Success/Confirm)
Primary Red:    #FF3B30 (Destructive/Error)
Primary Orange: #FF9500 (Warning/Alert)
```

### Semantic Colors
```
Background Primary:   #FFFFFF (Light) / #000000 (Dark)
Background Secondary: #F2F2F7 (Light) / #1C1C1E (Dark)
Background Tertiary:  #FFFFFF (Light) / #2C2C2E (Dark)

Text Primary:         #000000 (Light) / #FFFFFF (Dark)
Text Secondary:       #3C3C43 (Light) / #EBEBF5 (Dark)
Text Tertiary:        #3C3C4399 (Light) / #EBEBF599 (Dark)

Separator:            #3C3C4349 (Light) / #54545899 (Dark)
```

### Calendar-Specific Colors
```
Event Colors:
- Personal:     #007AFF (Blue)
- Work:         #5856D6 (Purple)
- Health:       #FF2D92 (Pink)
- Travel:       #FF9500 (Orange)
- Social:       #34C759 (Green)
- Education:    #5AC8FA (Light Blue)
- Finance:      #FFCC00 (Yellow)
- Home:         #8E8E93 (Gray)

Today Highlight:      #007AFF
Weekend:              #8E8E93
Holiday:              #FF3B30
```

## Typography

### iOS (SF Pro Family)
```
Large Title:    SF Pro Display, 34pt, Regular
Title 1:        SF Pro Display, 28pt, Regular
Title 2:        SF Pro Display, 22pt, Regular
Title 3:        SF Pro Display, 20pt, Regular
Headline:       SF Pro Text, 17pt, Semibold
Body:           SF Pro Text, 17pt, Regular
Callout:        SF Pro Text, 16pt, Regular
Subhead:        SF Pro Text, 15pt, Regular
Footnote:       SF Pro Text, 13pt, Regular
Caption 1:      SF Pro Text, 12pt, Regular
Caption 2:      SF Pro Text, 11pt, Regular
```

### Android (Roboto Family)
```
Display Large:  Roboto, 57sp, Regular
Display Medium: Roboto, 45sp, Regular
Display Small:  Roboto, 36sp, Regular
Headline Large: Roboto, 32sp, Regular
Headline Medium:Roboto, 28sp, Regular
Headline Small: Roboto, 24sp, Regular
Title Large:    Roboto Medium, 22sp
Title Medium:   Roboto Medium, 16sp
Title Small:    Roboto Medium, 14sp
Body Large:     Roboto, 16sp, Regular
Body Medium:    Roboto, 14sp, Regular
Body Small:     Roboto, 12sp, Regular
Label Large:    Roboto Medium, 14sp
Label Medium:   Roboto Medium, 12sp
Label Small:    Roboto Medium, 11sp
```

### Web (System Fonts)
```
Heading 1:      -apple-system, 32px, 600
Heading 2:      -apple-system, 28px, 600
Heading 3:      -apple-system, 24px, 600
Heading 4:      -apple-system, 20px, 600
Body Large:     -apple-system, 18px, 400
Body:           -apple-system, 16px, 400
Body Small:     -apple-system, 14px, 400
Caption:        -apple-system, 12px, 400
```

## Spacing System

### 8pt Grid System
```
4px   - Micro spacing (element padding)
8px   - Small spacing (between related elements)
16px  - Medium spacing (between sections)
24px  - Large spacing (between major sections)
32px  - XL spacing (page margins)
48px  - XXL spacing (major layout divisions)
64px  - XXXL spacing (between major sections)
```

## Icon System

### iOS Icons
- Use SF Symbols for system icons
- Custom icons: 24x24pt @1x (outline style)
- Calendar event icons: 20x20pt @1x

### Android Icons
- Material Design Icons
- Custom icons: 24dp (outline style)
- Filled variants for selected states

### Web Icons
- SVG format for scalability
- 24px base size with responsive scaling
- Consistent with mobile platforms

## Component Specifications

### Buttons

#### Primary Button
```
iOS:
- Height: 50pt
- Corner Radius: 10pt
- Background: Primary Blue (#007AFF)
- Text: White, SF Pro Text, 17pt, Semibold

Android:
- Height: 40dp
- Corner Radius: 20dp
- Background: Primary color from Material Theme
- Text: White, Roboto Medium, 14sp
- Elevation: 2dp

Web:
- Height: 44px
- Border Radius: 8px
- Background: #007AFF
- Text: White, system font, 16px, 600
```

#### Secondary Button
```
iOS:
- Height: 50pt
- Corner Radius: 10pt
- Background: Clear
- Border: 1pt Primary Blue
- Text: Primary Blue, SF Pro Text, 17pt, Regular

Android:
- Height: 40dp
- Corner Radius: 20dp
- Background: Transparent
- Border: 1dp Primary color
- Text: Primary color, Roboto, 14sp

Web:
- Height: 44px
- Border Radius: 8px
- Background: Transparent
- Border: 1px #007AFF
- Text: #007AFF, system font, 16px, 400
```

### Cards

#### Event Card
```
iOS:
- Corner Radius: 10pt
- Background: Background Tertiary
- Shadow: 0pt 1pt 3pt rgba(0,0,0,0.1)
- Padding: 16pt
- Left Border: 4pt colored accent

Android:
- Corner Radius: 12dp
- Background: Surface color
- Elevation: 1dp
- Padding: 16dp
- Left Border: 4dp colored accent

Web:
- Border Radius: 8px
- Background: #FFFFFF
- Box Shadow: 0 1px 3px rgba(0,0,0,0.1)
- Padding: 16px
- Left Border: 4px colored accent
```

### Calendar Components

#### Calendar Grid
```
iOS:
- Cell Size: Calculated to fit screen width
- Corner Radius: 8pt for selected dates
- Today indicator: Circle with primary color
- Text: SF Pro Text, 17pt

Android:
- Cell Size: 48dp minimum
- Corner Radius: 24dp for selected dates (circular)
- Today indicator: Colored circle background
- Text: Roboto, 14sp

Web:
- Cell Size: 40px minimum, responsive
- Border Radius: 50% for selected dates
- Today indicator: Colored circle
- Text: system font, 14px
```

## Animation Guidelines

### Timing Functions
```
iOS:
- Standard: ease-in-out (0.25s)
- Spring: spring(response: 0.5, dampingFraction: 0.8)

Android:
- Standard: cubic-bezier(0.4, 0.0, 0.2, 1) - 300ms
- Emphasized: cubic-bezier(0.2, 0.0, 0, 1) - 500ms

Web:
- Standard: ease-in-out (250ms)
- Bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55) (350ms)
```

### Common Animations
- View transitions: slide/fade
- Button press: scale down to 0.95
- Loading states: subtle pulse or shimmer
- Calendar navigation: slide left/right

## Accessibility

### Color Contrast
- Minimum 4.5:1 for normal text
- Minimum 3:1 for large text
- Support for high contrast mode

### Typography
- Support for dynamic type sizes
- Minimum 12pt/12sp for body text
- Clear hierarchy with size and weight

### Interaction
- Minimum 44pt/44dp touch targets
- Clear focus indicators
- VoiceOver/TalkBack support
- Keyboard navigation (web)

## Platform-Specific Considerations

### iOS
- Follow Human Interface Guidelines
- Use native navigation patterns
- Support Dark Mode
- Safe area considerations
- Dynamic Type support

### Android
- Material Design 3 principles
- Dynamic Color (Android 12+)
- Edge-to-edge design
- System bars integration
- Predictive back gestures

### Web
- Responsive design (mobile-first)
- Cross-browser compatibility
- Keyboard navigation
- Performance optimization
- Progressive Web App features