# ZCalendar HTML Mockups

This directory contains HTML mockups for the ZCalendar application across iOS, Android, and Web platforms, implementing the design specifications outlined in the documentation.

## Structure

```
mocks/
├── index.html                 # Main index page showcasing all mockups
├── ios/html/                  # iOS App Mockups
│   ├── index.html            # Main calendar view (iOS style)
│   ├── create-event.html     # Event creation modal
│   └── styles.css            # iOS-specific styles (Apple HIG)
├── android/html/             # Android App Mockups  
│   ├── index.html            # Main calendar view (Material Design 3)
│   └── styles.css            # Material Design styles
└── web/html/                 # Web App Mockups
    ├── index.html            # Desktop layout with sidebar
    ├── mobile.html           # Mobile-responsive layout
    ├── styles.css            # Desktop web styles
    └── mobile.css            # Mobile web styles
```

## Features Implemented

### iOS Mockups
- **Design System**: Apple Human Interface Guidelines
- **Typography**: SF Pro Display/Text family simulation
- **Colors**: iOS system colors with dark mode support
- **Components**: Native iOS navigation, tab bar, modal presentations
- **Interactions**: Touch-friendly buttons, form controls, toggle switches
- **Layout**: iPhone 13/14 (375x812pt) viewport simulation

### Android Mockups  
- **Design System**: Material Design 3
- **Typography**: Roboto family
- **Colors**: Material You dynamic colors with semantic meanings
- **Components**: FAB, bottom navigation, cards, elevation
- **Interactions**: Material ripple effects simulation, touch targets
- **Layout**: Pixel-equivalent (360x800dp) viewport

### Web Mockups
- **Design System**: Cross-platform responsive design
- **Typography**: System font stack for optimal performance  
- **Colors**: Unified brand colors adaptable to platforms
- **Layouts**: 
  - Desktop: Sidebar navigation, multi-column layout
  - Mobile: Bottom navigation, single column, touch-optimized
- **Responsive**: Mobile-first approach with breakpoints
- **Accessibility**: Semantic HTML, keyboard navigation, WCAG compliance

## Design Principles Applied

1. **Platform Native**: Each mockup respects its platform's design conventions
2. **Consistency**: Unified brand identity across all platforms
3. **Accessibility**: 44px minimum touch targets, high contrast, semantic markup
4. **Performance**: Optimized CSS, minimal DOM manipulation
5. **Responsive**: Mobile-first web design with proper breakpoints

## Key Screens Implemented

- **Calendar Month View**: Primary interface with event display
- **Event Creation**: Form interfaces following platform patterns
- **Navigation**: Platform-appropriate navigation patterns
- **Today's Events**: Event listing and management

## Getting Started

1. Open `mocks/index.html` in a web browser to see all mockups
2. Navigate to individual platform mockups:
   - iOS: `mocks/ios/html/index.html`
   - Android: `mocks/android/html/index.html`  
   - Web Desktop: `mocks/web/html/index.html`
   - Web Mobile: `mocks/web/html/mobile.html`

## Browser Support

- **Modern Browsers**: Chrome 88+, Firefox 86+, Safari 14+, Edge 88+
- **Mobile**: iOS Safari, Chrome Mobile, Samsung Internet
- **Features**: CSS Grid, Flexbox, CSS Custom Properties, prefers-color-scheme

## Technical Details

- **HTML5**: Semantic markup with ARIA attributes
- **CSS3**: Modern layout with Grid and Flexbox
- **Responsive**: Viewport meta tag and media queries
- **Performance**: Minimized reflows, efficient selectors
- **Dark Mode**: `prefers-color-scheme` media query support

This implementation provides a comprehensive foundation for the ZCalendar application's user interface across all target platforms while maintaining design consistency and platform-specific optimizations.