// Constructs API URLs from current dashboard state
import { getCurrentTheme } from './theme.js';

function encodeSeg(v) {
  return encodeURIComponent(String(v ?? ''));
}

export function buildPath(anim, state) {
  const { color1, color2, format, scene, title, subtitle, shape, texture, variant } = state;
  const c1 = color1 || 'white';
  const parts = [
    'api', 'v4', 'render',
    encodeSeg(format),
    encodeSeg(scene),
    encodeSeg(c1),
  ];
  if (color2) parts.push(encodeSeg(color2));
  parts.push(`${encodeSeg(anim)}.svg`);
  let base = `/${parts.join('/')}`;

  const params = [];
  const theme = getCurrentTheme();
  if (theme !== 'auto') params.push(`theme=${encodeSeg(theme)}`);
  if (title) params.push(`title=${encodeSeg(title)}`);
  if (subtitle) params.push(`subtitle=${encodeSeg(subtitle)}`);
  if (shape) params.push(`shape=${encodeSeg(shape)}`);
  if (texture) params.push(`texture=${encodeSeg(texture)}`);
  if (variant) params.push(`variant=${encodeSeg(variant)}`);

  if (params.length > 0) base += '?' + params.join('&');
  return base;
}

// Grid thumbnail: always favicon/pure for crisp previews at small sizes
export function buildGridPath(anim, state) {
  const { color1, color2, shape, texture, variant } = state;
  const c1 = color1 || 'white';
  const parts = [
    'api', 'v4', 'render',
    'favicon',
    'pure',
    encodeSeg(c1),
  ];
  if (color2) parts.push(encodeSeg(color2));
  parts.push(`${encodeSeg(anim)}.svg`);
  let base = `/${parts.join('/')}`;

  const params = [];
  const theme = getCurrentTheme();
  if (theme !== 'auto') params.push(`theme=${encodeSeg(theme)}`);
  if (shape) params.push(`shape=${encodeSeg(shape)}`);
  if (texture) params.push(`texture=${encodeSeg(texture)}`);
  if (variant) params.push(`variant=${encodeSeg(variant)}`);

  if (params.length > 0) base += '?' + params.join('&');
  return base;
}

export function buildFullUrl(anim, state) {
  return `${window.location.origin}${buildPath(anim, state)}`;
}
