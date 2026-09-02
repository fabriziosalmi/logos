// Animation grid: renders all animation cards
import { buildGridPath, buildFullUrl } from './url-builder.js';
import { show as showToast } from './toast.js';

/** One animation card. Built through the DOM, so no value needs escaping. */
function buildCard(anim) {
  const card = document.createElement('button');
  card.type = 'button';
  card.className = 'icon-card anim-card';
  card.dataset.anim = anim;
  card.setAttribute('aria-label', `Copy URL for ${anim}.svg`);

  const overlay = document.createElement('div');
  overlay.className = 'copy-overlay';
  const copyBtn = document.createElement('span');
  copyBtn.className = 'copy-btn';
  copyBtn.textContent = 'Copy URL';
  overlay.appendChild(copyBtn);

  const wrapper = document.createElement('div');
  wrapper.className = 'icon-wrapper anim-card-icon';
  const img = document.createElement('img');
  img.src = '';
  img.alt = anim;
  img.className = 'lazy-svg anim-card-img';
  img.dataset.anim = anim;
  img.loading = 'lazy';
  img.decoding = 'async';
  wrapper.appendChild(img);

  const footer = document.createElement('div');
  footer.className = 'anim-card-footer card-divider';
  const label = document.createElement('span');
  label.className = 'anim-label';
  label.textContent = anim;
  footer.appendChild(label);

  card.append(overlay, wrapper, footer);
  return card;
}

export function renderGrid(container, animations, getState) {
  // A fragment rather than a spread: replaceChildren(...list) passes one argument
  // per card, and this gallery grows with the animation set.
  const fragment = document.createDocumentFragment();
  for (const anim of animations) fragment.appendChild(buildCard(anim));
  container.replaceChildren(fragment);

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
