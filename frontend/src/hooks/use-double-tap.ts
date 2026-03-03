import { useCallback, useRef } from "react";

/**
 * Returns event handlers for double-tap / double-click.
 * - On desktop: uses native onDoubleClick
 * - On mobile: detects two quick taps within `delay` ms
 *
 * Single-tap fires `onSingleTap` immediately.
 * Double-tap fires `onDoubleTap` and cancels the second single-tap.
 */
export function useDoubleTap({
  onSingleTap,
  onDoubleTap,
  delay = 300,
}: {
  onSingleTap?: () => void;
  onDoubleTap?: () => void;
  delay?: number;
}) {
  const lastTap = useRef(0);

  const handleClick = useCallback(() => {
    const now = Date.now();
    if (now - lastTap.current < delay) {
      // Double tap detected
      lastTap.current = 0;
      onDoubleTap?.();
    } else {
      // Single tap
      lastTap.current = now;
      onSingleTap?.();
    }
  }, [onSingleTap, onDoubleTap, delay]);

  return { onClick: handleClick };
}
