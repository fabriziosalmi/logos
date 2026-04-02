// Theme engine: switches dashboard background + tracks selected SVG theme
let currentTheme = 'dark';

export function getCurrentTheme() {
  return currentTheme;
}

export function initTheme(onUpdate) {
  document.querySelectorAll('.theme-btn').forEach(btn => {
    btn.addEventListener('click', () => {
      document.querySelectorAll('.theme-btn').forEach(b => {
        b.classList.remove('active');
        b.style.borderColor = '';
      });

      btn.classList.add('active');
      const t = btn.dataset.theme;
      if (t === 'dark') btn.style.borderColor = '#525252';
      if (t === 'light') btn.style.borderColor = '#a3a3a3';
      if (t === 'auto') btn.style.borderColor = '#a3a3a3';
      if (t === 'solid') btn.style.borderColor = '#93c5fd';

      currentTheme = t;
      document.body.className = 'flex flex-col relative';
      if (t === 'light') document.body.classList.add('theme-light');
      else if (t === 'auto') document.body.classList.add('theme-checkered');
      else document.body.classList.add('theme-dark');

      onUpdate();
    });
  });
}
