export function initPalette(onSelect) {
  const palettes = {
    amber: 'eab308', blue: '3b82f6', cyan: '06b6d4', green: '22c55e', indigo: '6366f1', orange: 'f97316', purple: 'a855f7', rose: 'f43f5e',
    white: 'ffffff', gray: '9ca3af', black: '111111', gold: 'd4af37', platinum: 'e5e4e2', champagne: 'f7e7ce',
    neon: '39ff14', matrix: '00ff00', cyber: 'ff00ff', laser: 'ff0099', plasma: '00ffff', void: '8a2be2',
    emerald: '50c878', sapphire: '0f52ba', ruby: 'e0115f', ocean: '006994', sunset: 'fd5e53', magma: 'ff3300',
    mint: '98ff98', peach: 'ffdab9', lavender: 'e6e6fa',
  };

  const container = document.getElementById('palette-container');
  if (!container) return;

  const fragment = document.createDocumentFragment();
  Object.entries(palettes).forEach(([name, hex]) => {
    const btn = document.createElement('button');
    btn.type = 'button';
    btn.className = 'palette-pill';
    btn.title = name;
    btn.setAttribute('aria-label', `Set primary color: ${name}`);
    btn.addEventListener('click', () => onSelect(name));

    const swatch = document.createElement('span');
    swatch.className = 'palette-swatch';
    swatch.style.backgroundColor = `#${hex}`;
    swatch.style.boxShadow = '0 0 6px rgba(0,0,0,0.5)';

    const label = document.createElement('span');
    label.className = 'palette-pill-label';
    label.textContent = name;

    btn.appendChild(swatch);
    btn.appendChild(label);
    fragment.appendChild(btn);
  });

  container.textContent = '';
  container.appendChild(fragment);
}
