// Live preview + endpoint display
import { buildPath, buildFullUrl } from './url-builder.js';
import { show as showToast } from './toast.js';

const previewAnim = 'vortex';

export function updatePreview(state) {
  const liveBox = document.getElementById('live-preview-box');
  const docUrl = document.getElementById('doc-url');

  const path = buildPath(previewAnim, state);
  const fullUrl = buildFullUrl(previewAnim, state);

  const img = document.createElement('img');
  img.src = path;
  img.alt = 'Live preview';
  img.className = 'preview-img';
  img.decoding = 'async';
  liveBox.textContent = '';
  liveBox.appendChild(img);

  docUrl.textContent = fullUrl;
}

export function initPreviewCopy() {
  const docUrl = document.getElementById('doc-url');
  docUrl.setAttribute('role', 'button');
  docUrl.setAttribute('tabindex', '0');

  const copy = () => {
    navigator.clipboard.writeText(docUrl.textContent);
    showToast('Endpoint copied to clipboard');
  };

  docUrl.parentElement.addEventListener('click', copy);
  docUrl.addEventListener('keydown', (e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      copy();
    }
  });
  docUrl.parentElement.classList.add('cursor-pointer', 'hover:border-blue-500/50', 'transition-colors');
}
