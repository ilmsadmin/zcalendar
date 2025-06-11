# Icon Library & Assets

## App Icons

### Primary App Icon Specifications

#### iOS App Icon
```
Sizes Required:
- 1024x1024px - App Store
- 180x180px - iPhone @3x
- 120x120px - iPhone @2x
- 152x152px - iPad @2x
- 76x76px - iPad @1x

Design Concept:
┌─────────────────┐
│    ┌───────┐    │  ← Rounded rectangle background
│    │   📅   │    │     Gradient: #007AFF to #5AC8FA
│    │  Dec   │    │  ← Calendar page metaphor
│    │   14   │    │     White calendar with shadow
│    └───────┘    │  ← Today's date highlighted
│                 │
└─────────────────┘

Visual Elements:
- iOS blue gradient background
- White calendar page with subtle shadow
- Today's date in highlighted state
- Clean, minimal design
- Rounded corners matching iOS style
```

#### Android App Icon
```
Sizes Required:
- 512x512px - Play Store
- 192x192px - xxxhdpi
- 144x144px - xxhdpi
- 96x96px - xhdpi
- 72x72px - hdpi
- 48x48px - mdpi

Design Concept:
┌─────────────────┐
│     ●●●●●●●     │  ← Material Design adaptive icon
│   ●         ●   │     Background: Material blue
│  ●     📅     ●  │  ← Calendar icon in foreground
│  ●    Dec    ●  │     Monochromatic white design
│  ●     14     ●  │  ← Adaptive to system shapes
│   ●         ●   │
│     ●●●●●●●     │
└─────────────────┘

Visual Elements:
- Material Design adaptive icon format
- Background layer: Solid Material blue
- Foreground layer: White calendar icon
- Supports Android theming (Material You)
- Compatible with system icon shapes
```

### Interface Icons

#### Navigation Icons
```
Tab Bar Icons (24x24pt/dp):

📅 Calendar - calendar
┌─────────┐
│  ┌───┬─┐│  ← Calendar grid representation
│  │   │ ││     Outline style for inactive
│  ├───┼─┤│     Filled style for active
│  │ ● │ ││  ← Today indicator dot
│  └───┴─┘│
└─────────┘

📋 Events - list.bullet
┌─────────┐
│ ■■■■■   │  ← List items with bullet points
│ ■■■■■■  │     Consistent with platform icons
│ ■■■     │
│ ■■■■■■■ │
└─────────┘

➕ Create - plus
┌─────────┐
│         │
│    │    │  ← Plus symbol
│  ──┼──  │     Centered and balanced
│    │    │
│         │
└─────────┘

⚙️ Settings - gear
┌─────────┐
│   ●●●   │  ← Gear/cog wheel
│  ●   ●  │     8 or 6 teeth
│ ●  ●  ● │     Mechanical appearance
│  ●   ●  │
│   ●●●   │
└─────────┘
```

#### Action Icons
```
Common Action Icons (20x20pt/dp):

✓ Save/Confirm - checkmark.circle
🗑️ Delete - trash
✏️ Edit - pencil
📤 Share - square.and.arrow.up
🔍 Search - magnifyingglass
⋮ More - ellipsis.vertical
← Back - chevron.left
→ Forward - chevron.right
↻ Refresh - arrow.clockwise
⚠️ Alert - exclamationmark.triangle
ℹ️ Info - info.circle
```

#### Event Type Icons
```
Event Category Icons (16x16pt/dp):

🏢 Work - building.2
👥 Meeting - person.2
📞 Call - phone
✈️ Travel - airplane
🏥 Health - heart
🎓 Education - graduationcap
💰 Finance - dollarsign.circle
🏠 Home - house
🎉 Social - party.popper
🍽️ Food - fork.knife
🚗 Transport - car
📺 Entertainment - tv
```

## Color System Reference

### Primary Colors
```
Brand Colors:
#007AFF - iOS Blue (Primary)
#5AC8FA - Light Blue (Secondary)
#34C759 - Green (Success)
#FF3B30 - Red (Error/Delete)
#FF9500 - Orange (Warning)
#AF52DE - Purple (Premium)

Semantic Colors:
#000000 - Text Primary (Light mode)
#FFFFFF - Text Primary (Dark mode)
#3C3C43 - Text Secondary (Light mode)
#EBEBF5 - Text Secondary (Dark mode)
#8E8E93 - Text Tertiary/Disabled
```

### Calendar Event Colors
```
Category Color Palette:
🔵 #007AFF - Personal (Blue)
🟣 #5856D6 - Work (Purple)
🟢 #34C759 - Health (Green)
🟠 #FF9500 - Travel (Orange)
🔴 #FF3B30 - Important (Red)
🟡 #FFCC00 - Social (Yellow)
🟤 #A2845E - Home (Brown)
⚫ #8E8E93 - Other (Gray)
```

## Typography Scale

### iOS Typography (SF Pro)
```
Display Styles:
Large Title: 34pt Regular
Title 1: 28pt Regular
Title 2: 22pt Regular
Title 3: 20pt Regular

Body Styles:
Headline: 17pt Semibold
Body: 17pt Regular
Callout: 16pt Regular
Subhead: 15pt Regular
Footnote: 13pt Regular
Caption 1: 12pt Regular
Caption 2: 11pt Regular
```

### Android Typography (Roboto)
```
Material Design Scale:
Display Large: 57sp Regular
Display Medium: 45sp Regular
Display Small: 36sp Regular
Headline Large: 32sp Regular
Headline Medium: 28sp Regular
Headline Small: 24sp Regular

Body Styles:
Title Large: 22sp Medium
Title Medium: 16sp Medium
Title Small: 14sp Medium
Body Large: 16sp Regular
Body Medium: 14sp Regular
Body Small: 12sp Regular
Label Large: 14sp Medium
Label Medium: 12sp Medium
Label Small: 11sp Medium
```

### Web Typography (System)
```
CSS Custom Properties:
--font-family-primary: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
--font-family-mono: 'SF Mono', Monaco, 'Cascadia Code', monospace;

Font Scale:
--font-size-xs: 12px
--font-size-sm: 14px
--font-size-base: 16px
--font-size-lg: 18px
--font-size-xl: 20px
--font-size-2xl: 24px
--font-size-3xl: 30px
--font-size-4xl: 36px

Font Weights:
--font-weight-light: 300
--font-weight-normal: 400
--font-weight-medium: 500
--font-weight-semibold: 600
--font-weight-bold: 700
```

## Spacing System

### Grid System (8pt Base)
```
Spacing Scale:
--space-1: 4px   (0.25rem)
--space-2: 8px   (0.5rem)
--space-3: 12px  (0.75rem)
--space-4: 16px  (1rem)
--space-5: 20px  (1.25rem)
--space-6: 24px  (1.5rem)
--space-8: 32px  (2rem)
--space-10: 40px (2.5rem)
--space-12: 48px (3rem)
--space-16: 64px (4rem)
--space-20: 80px (5rem)
--space-24: 96px (6rem)

Component Spacing:
- Button padding: 16px 24px
- Card padding: 16px
- Section spacing: 24px
- Page margins: 16px (mobile), 24px (tablet), 32px (desktop)
```

## Animation Specifications

### Duration & Easing
```
Standard Durations:
--duration-fast: 150ms
--duration-normal: 250ms
--duration-slow: 350ms
--duration-slower: 500ms

Easing Functions:
--ease-out: cubic-bezier(0.0, 0.0, 0.2, 1)
--ease-in: cubic-bezier(0.4, 0.0, 1, 1)
--ease-in-out: cubic-bezier(0.4, 0.0, 0.2, 1)
--ease-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55)

Common Animations:
- Button press: scale(0.95) 150ms
- Modal enter: translateY(0) 350ms ease-out
- Modal exit: translateY(100%) 250ms ease-in
- Fade in: opacity 250ms ease-out
- Slide transition: transform 300ms ease-in-out
```

## Responsive Breakpoints

### Screen Size Guidelines
```
Breakpoints:
--screen-xs: 320px  (Small phones)
--screen-sm: 480px  (Large phones)
--screen-md: 768px  (Tablets)
--screen-lg: 1024px (Small desktops)
--screen-xl: 1440px (Large desktops)
--screen-2xl: 1920px (Extra large)

Container Widths:
--container-sm: 640px
--container-md: 768px
--container-lg: 1024px
--container-xl: 1280px

Safe Areas (Mobile):
--safe-area-top: env(safe-area-inset-top)
--safe-area-bottom: env(safe-area-inset-bottom)
--safe-area-left: env(safe-area-inset-left)
--safe-area-right: env(safe-area-inset-right)
```

## Accessibility Guidelines

### Color Contrast
```
WCAG AA Requirements:
- Normal text: 4.5:1 minimum ratio
- Large text: 3:1 minimum ratio
- UI components: 3:1 minimum ratio

High Contrast Mode:
- Text on background: 7:1 ratio
- Enhanced border contrast
- No color-only information
```

### Touch Targets
```
Minimum Sizes:
- iOS: 44pt x 44pt (44px)
- Android: 48dp x 48dp (48px)
- Web: 44px x 44px minimum

Spacing:
- Minimum 8px between targets
- 16px preferred spacing
- Consider thumb reach zones
```

### Focus Indicators
```
Focus Ring Specifications:
- Width: 2px outline
- Color: Primary blue (#007AFF)
- Offset: 2px from element
- Border radius: Matches element + 2px
- High contrast alternative available
```