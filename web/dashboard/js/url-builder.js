// Constructs API URLs from current dashboard state
import { getCurrentTheme } from './theme.js';

export function buildPath(anim, state) {
  const { color1, color2, format, scene, title, subtitle, shape, texture } = state;
  const c1 = color1 || 'white';
  let base = `/api/v4/render/${format}/${scene}/${c1}${color2 ? '/' + color2 : ''}/${anim}.svg`;

  const params = [];
  const theme = getCurrentTheme();
  if (theme !== 'auto') params.push(`theme=${theme}`);
  if (title) params.push(`title=${encodeURIComponent(title)}`);
  if (subtitle) params.push(`subtitle=${encodeURIComponent(subtitle)}`);
  if (shape) params.push(`shape=${shape}`);
  if (texture) params.push(`texture=${texture}`);

  if (params.length > 0) base += '?' + params.join('&');
  return base;
}

export function buildFullUrl(anim, state) {
  return `${window.location.origin}${buildPath(anim, state)}`;
}
