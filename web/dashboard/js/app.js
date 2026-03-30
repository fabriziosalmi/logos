// Main app: wires all modules together
import { initTheme } from './theme.js';
import { renderGrid, updateGrid } from './grid.js';
import { updatePreview, initPreviewCopy } from './preview.js';

const ANIMATIONS = [
  'static',
  'zen', 'breathe', 'levitate', 'glimmer',
  'spin', 'spin-fast', 'smooth-spin', 'orbit-chase', 'compass', 'gyro', 'satellite', 'eclipse',
  'pulse', 'heartbeat', 'pulse-ring', 'strobe', 'nova', 'elastic',
  'flip', 'orbit-tilt', 'vortex', 'harmony', 'sync', 'sway',
  'radar', 'radar-sweep', 'signal', 'glow', 'aurora', 'nebula', 'corona', 'ripple-core',
  'morph', 'morph-blob', 'morph-crystal',
  'bounce', 'bounce-drop', 'trampoline',
  'shake', 'jitter', 'earthquake',
  'swing', 'pendulum', 'wave-swing',
  'zoom-in', 'zoom-out', 'zoom-pulse',
  'slide-in', 'slide-loop',
  'typewriter-blink',
];

// DOM refs
const inputC1 = document.getElementById('color1');
const inputC2 = document.getElementById('color2');
const formatSelect = document.getElementById('format-select');
const sceneSelect = document.getElementById('scene-select');
const shapeSelect = document.getElementById('shape-select');
const textureSelect = document.getElementById('texture-select');
const variantSelect = document.getElementById('variant-select');
const titleInput = document.getElementById('title-input');
const subtitleInput = document.getElementById('subtitle-input');
const gridContainer = document.getElementById('grid-container');

function getState() {
  return {
    color1: inputC1.value.trim() || 'white',
    color2: inputC2.value.trim() || '',
    format: formatSelect.value,
    scene: sceneSelect.value,
    shape: shapeSelect ? shapeSelect.value : '',
    texture: textureSelect ? textureSelect.value : '',
    variant: variantSelect ? variantSelect.value : '',
    title: titleInput.value,
    subtitle: subtitleInput.value,
  };
}

function updateAll() {
  updatePreview(getState());
  updateGrid(getState);
}

// Palette click
window.setColor = (name) => {
  inputC1.value = name;
  inputC2.value = '';
  updateAll();
};

// Theme engine
initTheme(updateAll);

// Grid
renderGrid(gridContainer, ANIMATIONS, getState);

// Live preview copy
initPreviewCopy();

// Input listeners with debounce
const inputs = [inputC1, inputC2, formatSelect, sceneSelect, shapeSelect, textureSelect, variantSelect, titleInput, subtitleInput];
inputs.forEach(el => {
  if (!el) return;
  let timer;
  const evt = (el.tagName === 'SELECT') ? 'change' : 'input';
  el.addEventListener(evt, () => {
    clearTimeout(timer);
    timer = setTimeout(updateAll, 150);
  });
  // Also listen to input for selects (covers both)
  if (evt === 'change') {
    el.addEventListener('input', () => {
      clearTimeout(timer);
      timer = setTimeout(updateAll, 150);
    });
  }
});

// Init
updateAll();
