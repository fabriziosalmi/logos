// Animation grid: renders all animation cards
import { buildGridPath, buildFullUrl } from './url-builder.js';
import { show as showToast } from './toast.js';

export function renderGrid(container, animations, getState) {
  container.innerHTML = animations.map(anim => `
    <button type="button" class="icon-card anim-card" data-anim="${anim}" aria-label="Copy URL for ${anim}.svg">
      <div class="copy-overlay">
        <span class="copy-btn">Copy URL</span>
      </div>
      <div class="icon-wrapper anim-card-icon">
        <img src="" alt="${anim}" class="lazy-svg anim-card-img" data-anim="${anim}" loading="lazy" decoding="async" />
      </div>
      <div class="anim-card-footer card-divider">
        <span class="anim-label">${anim}</span>
      </div>
    </button>
  `).join('');

  container.querySelectorAll('.icon-card').forEach(card => {
    card.addEventListener('click', () => {
      const anim = card.dataset.anim;
      const url = buildFullUrl(anim, getState());
      navigator.clipboard.writeText(url);
      card.style.transform = 'scale(0.92)';
      setTimeout(() => card.style.transform = '', 150);
      showToast(`Copied ${anim}.svg`);
    });
  });
}

export function updateGrid(getState) {
  document.querySelectorAll('.lazy-svg').forEach(img => {
    img.src = buildGridPath(img.dataset.anim, getState());
  });
}
