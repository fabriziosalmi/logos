// Live preview + endpoint display
import { buildPath, buildFullUrl } from './url-builder.js';
import { show as showToast } from './toast.js';

const previewAnim = 'vortex';

export function updatePreview(state) {
  const liveBox = document.getElementById('live-preview-box');
  const docUrl = document.getElementById('doc-url');

  const path = buildPath(previewAnim, state);
  const fullUrl = buildFullUrl(previewAnim, state);

  liveBox.innerHTML = `<img src="${path}" class="w-full h-full object-contain" />`;

  let highlighted = fullUrl
    .replace('/api/v4/render/', '<span class="text-neutral-500">/api/v4/render/</span>')
    .replace('.svg', '<span class="text-neutral-500">.svg</span>');
  docUrl.innerHTML = highlighted;
}

export function initPreviewCopy() {
  const docUrl = document.getElementById('doc-url');
  docUrl.parentElement.addEventListener('click', () => {
    navigator.clipboard.writeText(docUrl.textContent);
    showToast('Endpoint copied to clipboard!');
  });
  docUrl.parentElement.classList.add('cursor-pointer', 'hover:border-blue-500/50', 'transition-colors');
}
