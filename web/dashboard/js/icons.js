// Icon packs browser: fetches, displays, and lets users pick icons as shapes
import { show as showToast } from './toast.js';

let allIcons = [];
let filteredIcons = [];
let displayedCount = 0;
const PAGE_SIZE = 120;
let currentPack = '';
let currentSearch = '';
let onIconSelect = null;

export function initIconBrowser(onSelect) {
  onIconSelect = onSelect;

  // Pack filter buttons
  document.querySelectorAll('.icon-pack-btn').forEach(btn => {
    btn.addEventListener('click', () => {
      document.querySelectorAll('.icon-pack-btn').forEach(b => b.classList.remove('active'));
      btn.classList.add('active');
      currentPack = btn.dataset.pack;
      applyFilter();
    });
  });

  // Search input
  const searchInput = document.getElementById('icon-search');
  let timer;
  searchInput.addEventListener('input', () => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      currentSearch = searchInput.value.trim().toLowerCase();
      applyFilter();
    }, 200);
  });

  // Load more button
  document.getElementById('icon-load-more-btn').addEventListener('click', () => {
    renderMore();
  });

  // Fetch all icons
  fetchIcons();
}

async function fetchIcons() {
  try {
    const res = await fetch('/api/v4/icons');
    const data = await res.json();
    allIcons = data.icons || [];
    document.getElementById('icon-count-badge').textContent = `${allIcons.length} Icons`;
    applyFilter();
  } catch (e) {
    document.getElementById('icon-count-badge').textContent = 'Error loading';
  }
}

function applyFilter() {
  filteredIcons = allIcons.filter(name => {
    if (currentPack && !name.startsWith(currentPack + ':')) return false;
    if (currentSearch && !name.includes(currentSearch)) return false;
    return true;
  });

  displayedCount = 0;
  document.getElementById('icon-grid').innerHTML = '';
  document.getElementById('icon-count-badge').textContent = `${filteredIcons.length} Icons`;
  renderMore();
}

function renderMore() {
  const grid = document.getElementById('icon-grid');
  const end = Math.min(displayedCount + PAGE_SIZE, filteredIcons.length);
  const start = displayedCount;

  const fragment = document.createDocumentFragment();
  for (let i = displayedCount; i < end; i++) {
    const iconKey = filteredIcons[i];
    const shortName = iconKey.split(':')[1] || iconKey;

    const card = document.createElement('button');
    card.type = 'button';
    card.className = 'icon-pack-card';
    card.title = iconKey;
    card.dataset.icon = iconKey;
    card.setAttribute('aria-label', `Select shape: ${iconKey}`);

    const img = document.createElement('img');
    img.src = `/api/v4/render/favicon/pure/white/static.svg?shape=${encodeURIComponent(iconKey)}&theme=dark`;
    img.alt = shortName;
    img.loading = 'lazy';

    const label = document.createElement('span');
    label.textContent = shortName;

    card.appendChild(img);
    card.appendChild(label);

    card.addEventListener('click', () => {
      // Deselect previous
      grid.querySelectorAll('.selected').forEach(c => c.classList.remove('selected'));
      card.classList.add('selected');

      // Set the shape selector to the icon key
      const shapeSelect = document.getElementById('shape-select');
      // Check if option exists, if not create it
      let opt = shapeSelect.querySelector(`option[value="${iconKey}"]`);
      if (!opt) {
        opt = document.createElement('option');
        opt.value = iconKey;
        opt.textContent = iconKey;
        shapeSelect.appendChild(opt);
      }
      shapeSelect.value = iconKey;
      shapeSelect.dispatchEvent(new Event('change'));

      if (onIconSelect) onIconSelect(iconKey);
      showToast(`Shape selected: ${iconKey}`);
    });

    fragment.appendChild(card);
  }
  grid.appendChild(fragment);

  displayedCount = end;

  // Show/hide load more
  const loadMore = document.getElementById('icon-load-more');
  if (displayedCount < filteredIcons.length) {
    loadMore.classList.add('icon-load-more-visible');
    loadMore.classList.remove('icon-load-more-hidden');
  } else {
    loadMore.classList.add('icon-load-more-hidden');
    loadMore.classList.remove('icon-load-more-visible');
  }

  if (end > start) {
    const btn = document.getElementById('icon-load-more-btn');
    if (btn && typeof btn.scrollIntoView === 'function') {
      btn.scrollIntoView({ block: 'end' });
    }
  }
}
