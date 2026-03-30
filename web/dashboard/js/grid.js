// Animation grid: renders all animation cards
import { buildPath, buildFullUrl } from './url-builder.js';
import { show as showToast } from './toast.js';

export function renderGrid(container, animations, getState) {
  container.innerHTML = animations.map(anim => `
    <div class="icon-card rounded-xl p-3 cursor-pointer group flex flex-col items-center justify-between" data-anim="${anim}" title="Click to copy API URL">
      <div class="copy-overlay rounded-xl">
        <span class="copy-btn">Copy URL</span>
      </div>
      <div class="h-14 w-full flex items-center justify-center icon-wrapper mt-1">
        <img src="" alt="${anim}" class="w-12 h-12 lazy-svg" data-anim="${anim}" />
      </div>
      <div class="w-full text-center mt-2 pt-2" style="border-top:1px solid #222">
        <span class="anim-label">${anim}</span>
      </div>
    </div>
  `).join('');

  container.querySelectorAll('.icon-card').forEach(card => {
    card.addEventListener('click', () => {
      const anim = card.dataset.anim;
      const url = buildFullUrl(anim, getState());
      navigator.clipboard.writeText(url);
      card.style.transform = 'scale(0.92)';
      setTimeout(() => card.style.transform = '', 150);
      showToast(`Copied <span class="font-mono" style="color:#3b82f6;background:#1e3a5f;padding:0 4px;border-radius:3px">${anim}.svg</span>`);
    });
  });
}

export function updateGrid(getState) {
  document.querySelectorAll('.lazy-svg').forEach(img => {
    img.src = buildPath(img.dataset.anim, getState());
  });
}
