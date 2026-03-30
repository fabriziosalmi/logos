const http = require('http');

const PORT = process.env.PORT || 3000;

// --- PALETTES ---
const baseColors = {
  amber: 'eab308', blue: '3b82f6', cyan: '06b6d4', green: '22c55e',
  indigo: '6366f1', orange: 'f97316', purple: 'a855f7', rose: 'f43f5e'
};

const creativeColors = {
  white: 'ffffff', gray: '9ca3af', black: '111111', gold: 'd4af37', platinum: 'e5e4e2', champagne: 'f7e7ce',
  neon: '39ff14', matrix: '00ff00', cyber: 'ff00ff', laser: 'ff0099', plasma: '00ffff', void: '8a2be2',
  emerald: '50c878', sapphire: '0f52ba', ruby: 'e0115f', ocean: '006994', sunset: 'fd5e53', magma: 'ff3300',
  mint: '98ff98', peach: 'ffdab9', lavender: 'e6e6fa'
};

const allPalettes = { ...baseColors, ...creativeColors };

const getHex = (val) => {
  if (allPalettes[val]) return allPalettes[val];
  return /^[0-9A-Fa-f]{3,6}$/.test(val) ? val : 'ffffff';
};

// FULL 50 ANIMATIONS LIBRARY RE-ADDED
const animations = {
  // 1. STATICS
  static: '',
  
  // 2. ZEN & MINIMAL
  zen: `
    .core { animation: zen-core 6s var(--ease) infinite alternate; }
    .orbit { animation: zen-orbit 12s linear infinite; opacity: 0.8; }
    @keyframes zen-core { 0% { transform: scale(0.95) translateY(1px); opacity: 0.8; } 100% { transform: scale(1.05) translateY(-1px); opacity: 1; } }
    @keyframes zen-orbit { 100% { transform: rotate(360deg); } }
  `,
  breathe: `
    .outer, .orbit, .core { animation: breathe 5s var(--ease) infinite alternate; }
    @keyframes breathe { 0% { transform: scale(0.98); } 100% { transform: scale(1.04); } }
  `,
  levitate: `
    .core { animation: levitate 4s var(--ease-smooth) infinite alternate; }
    .orbit { animation: levitate-shadow 4s var(--ease-smooth) infinite alternate; }
    @keyframes levitate { 0% { transform: translateY(0); } 100% { transform: translateY(-3px); } }
    @keyframes levitate-shadow { 0% { transform: scale(1); opacity: 1; } 100% { transform: scale(0.95) translateY(2px); opacity: 0.6; } }
  `,
  glimmer: `
    .core { animation: glimmer 3s ease-in-out infinite alternate; }
    @keyframes glimmer { 0%, 100% { opacity: 1; } 50% { opacity: 0.3; } }
  `,

  // 3. SPINS & ORBITS
  spin: `
    .orbit { animation: spin 5s linear infinite; }
    @keyframes spin { 100% { transform: rotate(360deg); } }
  `,
  'spin-fast': `
    .orbit { animation: spin 1.5s linear infinite; }
    @keyframes spin { 100% { transform: rotate(360deg); } }
  `,
  'smooth-spin': `
    .orbit { animation: smooth-spin 8s var(--ease) infinite; }
    @keyframes smooth-spin { 0% { transform: rotate(0deg); } 50% { transform: rotate(180deg); } 100% { transform: rotate(360deg); } }
  `,
  'orbit-chase': `
    .orbit { stroke-dasharray: 20 10; animation: spin 2s linear infinite; }
  `,
  compass: `
    .orbit { animation: compass 4s cubic-bezier(0.68, -0.55, 0.265, 1.55) infinite; }
    @keyframes compass { 0%, 20% { transform: rotate(0); } 25%, 45% { transform: rotate(90deg); } 50%, 70% { transform: rotate(180deg); } 75%, 95% { transform: rotate(270deg); } 100% { transform: rotate(360deg); } }
  `,
  gyro: `
    .outer { animation: spin 8s linear infinite; stroke-dasharray: 40 5; }
    .orbit { animation: spin-rev 6s linear infinite; }
    @keyframes spin-rev { 100% { transform: rotate(-360deg); } }
  `,
  satellite: `
    .orbit { stroke-dasharray: 4 30; animation: spin 1.5s linear infinite; stroke-linecap: round; }
  `,
  eclipse: `
    .orbit { animation: eclipse 4s ease-in-out infinite; }
    @keyframes eclipse { 0%, 100% { transform: scaleY(1); } 50% { transform: scaleY(0); } }
  `,

  // 4. PULSES & HEARTS
  pulse: `
    .core { animation: pulse 3s var(--ease) infinite; }
    @keyframes pulse { 0%, 100% { transform: scale(1); opacity: 1; } 50% { transform: scale(1.25); opacity: 0.8; } }
  `,
  heartbeat: `
    .core { animation: heartbeat 2s var(--ease) infinite; }
    @keyframes heartbeat { 0%, 100% { transform: scale(1); } 15%, 45% { transform: scale(1.2); } 30%, 60% { transform: scale(1); } }
  `,
  'pulse-ring': `
    .outer { animation: pulse-ring 2.5s cubic-bezier(0.215, 0.61, 0.355, 1) infinite; }
    @keyframes pulse-ring { 0% { transform: scale(0.9); opacity: 1; } 100% { transform: scale(1.2); opacity: 0; } }
  `,
  strobe: `
    .core { animation: strobe 1.5s steps(2, start) infinite; }
    @keyframes strobe { 0%, 100% { opacity: 1; } 50% { opacity: 0; } }
  `,
  nova: `
    .core { animation: nova 3s ease-out infinite; }
    @keyframes nova { 0% { transform: scale(1); opacity: 1; } 100% { transform: scale(2.5); opacity: 0; } }
  `,
  elastic: `
    .core { animation: elastic 2s ease-in-out infinite; }
    @keyframes elastic { 0%, 100% { transform: scale(1, 1); } 50% { transform: scale(1.15, 0.85); } }
  `,

  // 5. COMPLEX & 3D
  flip: `
    .orbit { animation: flip 4s var(--ease-smooth) infinite; perspective: 100px; }
    @keyframes flip { 0% { transform: rotateY(0); } 50% { transform: rotateY(180deg); } 100% { transform: rotateY(360deg); } }
  `,
  'orbit-tilt': `
    .orbit { animation: orbit-tilt 4s ease-in-out infinite alternate; }
    @keyframes orbit-tilt { 0% { transform: rotateX(0); } 100% { transform: rotateX(60deg); } }
  `,
  vortex: `
    .outer { animation: spin 10s linear infinite; }
    .orbit { animation: spin 5s linear infinite; }
    .core { animation: spin 2s linear infinite; stroke-dasharray: 2 4; }
  `,
  harmony: `
    .core { animation: breathe 4s ease-in-out infinite alternate; }
    .orbit { animation: spin 8s linear infinite; }
  `,
  sync: `
    .outer, .orbit, .core { animation: sync 4s ease-in-out infinite; }
    @keyframes sync { 0%, 100% { transform: scale(1) rotate(0); } 50% { transform: scale(1.1) rotate(180deg); } }
  `,
  sway: `
    .outer, .orbit, .core { animation: sway 4s ease-in-out infinite alternate; }
    @keyframes sway { 0% { transform: rotate(-15deg); } 100% { transform: rotate(15deg); } }
  `,

  // 6. SCI-FI & TECH
  radar: `
    .outer { animation: radar 2.5s var(--ease) infinite; }
    @keyframes radar { 0% { transform: scale(0.8); opacity: 1; stroke-width: 2; } 100% { transform: scale(1.5); opacity: 0; stroke-width: 0; } }
  `,
  'radar-sweep': `
    .outer { animation: sweep 4s linear infinite; stroke-dasharray: 40 80; stroke-dashoffset: 0; }
    @keyframes sweep { 100% { stroke-dashoffset: -120; transform: rotate(360deg); } }
  `,
  signal: `
    .outer { animation: signal 2s ease-out infinite; }
    @keyframes signal { 0% { transform: scale(0.9); opacity: 1; stroke-width: 3; } 100% { transform: scale(1.3); opacity: 0; stroke-width: 1; } }
  `,
  glow: `
    .outer { animation: glow 3s var(--ease) infinite alternate; }
    @keyframes glow { 0% { stroke-width: 1.5; stroke-opacity: 1; } 100% { stroke-width: 4.5; stroke-opacity: 0.3; } }
  `,
  aurora: `
    .outer { animation: aurora 3s linear infinite; }
    @keyframes aurora { 0% { stroke-width: 1.5; opacity: 1; transform: scale(1); } 100% { stroke-width: 10; opacity: 0; transform: scale(1.2); } }
  `,
  nebula: `
    .outer { animation: nebula-outer 4s ease-in-out infinite alternate; }
    .orbit { animation: spin 6s linear infinite; stroke-width: 0.5; }
    @keyframes nebula-outer { 0% { stroke-width: 1.5; transform: scale(1); opacity: 1; } 100% { stroke-width: 4; transform: scale(1.1); opacity: 0.5; } }
  `,
  corona: `
    .outer { animation: corona 3s ease-in-out infinite alternate; }
    @keyframes corona { 0% { stroke-dasharray: 20 5; transform: scale(1); } 100% { stroke-dasharray: 5 20; transform: scale(1.05); } }
  `,
  'ripple-core': `
    .core { animation: ripple-core 2s cubic-bezier(0.4, 0, 0.2, 1) infinite; }
    @keyframes ripple-core { 0% { transform: scale(1); opacity: 1; } 100% { transform: scale(2); opacity: 0; } }
  `
};

const getSvg = (primaryHex, animName, secondaryHex = null, theme = 'auto') => {
  const animCss = animations[animName] || animations.static;
  
  const hasGradient = secondaryHex && secondaryHex !== primaryHex;
  const fillDef = hasGradient ? `url(#grad)` : `#${primaryHex}`;
  const strokeDef = hasGradient ? `url(#grad)` : `#${primaryHex}`;

  let coreBgColor = 'transparent'; 
  let dotColor = '#fff';
  let strokeWidth = '1.5';
  
  if (theme === 'dark') {
    coreBgColor = '#0a0a0c'; 
    dotColor = '#ffffff';
  } else if (theme === 'light') {
    coreBgColor = '#ffffff'; 
    dotColor = '#0a0a0c'; 
  } else if (theme === 'solid') {
    coreBgColor = fillDef; 
    dotColor = '#ffffff';
    strokeWidth = '0'; 
  } else if (theme === 'glass') {
    coreBgColor = 'rgba(255,255,255,0.1)';
    dotColor = '#ffffff';
  }

  const defs = hasGradient ? `
    <defs>
      <linearGradient id="grad" x1="0%" y1="0%" x2="100%" y2="100%">
        <stop offset="0%" stop-color="#${primaryHex}" />
        <stop offset="100%" stop-color="#${secondaryHex}" />
      </linearGradient>
    </defs>
  ` : '';

  return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32">
  ${defs}
  <style>
    :root { 
      --ease: cubic-bezier(0.25, 1, 0.5, 1);
      --ease-smooth: cubic-bezier(0.65, 0, 0.35, 1);
    }
    .outer, .orbit, .core { transform-origin: 16px 16px; }
    ${animCss}
  </style>
  <circle cx="16" cy="16" r="14" fill="${coreBgColor}" stroke="${strokeDef}" stroke-width="${strokeWidth}" class="outer"/>
  <ellipse cx="16" cy="16" rx="10" ry="7" fill="none" stroke="${strokeDef}" stroke-width="1.5" class="orbit"/>
  <g class="core">
    <circle cx="16" cy="16" r="3.5" fill="${fillDef}"/>
    <circle cx="17.2" cy="14.8" r="1.2" fill="${dotColor}" opacity="0.9"/>
  </g>
</svg>`;
};

// --- FRONTEND UI ---
const getDashboardHtml = () => {
  
  const generateColorPills = (colorObj) => {
    return Object.entries(colorObj).map(([name, hex]) => `
      <button onclick="setColor('${name}')" 
              class="group flex items-center gap-2 px-3 py-1.5 rounded-full bg-surface2 border border-border hover:border-[#${hex}] transition-all focus:outline-none focus:ring-2 focus:ring-[#${hex}]/50"
              title="Apply ${name}">
        <span class="w-3 h-3 rounded-full shadow-[0_0_10px_rgba(0,0,0,0.5)]" style="background-color: #${hex};"></span>
        <span class="text-xs font-medium text-neutral-400 group-hover:text-white transition-colors capitalize">${name}</span>
      </button>
    `).join('');
  };

  return `<!DOCTYPE html>
<html lang="en" class="dark">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Logos API | Next-Gen Assets</title>
  <link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@300;400;500;600;700;800&family=JetBrains+Mono:wght@400;700&display=swap" rel="stylesheet">
  <script src="https://cdn.tailwindcss.com"></script>
  <script>
    tailwind.config = {
      darkMode: 'class',
      theme: {
        extend: {
          fontFamily: {
            sans: ['"Plus Jakarta Sans"', 'sans-serif'],
            mono: ['"JetBrains Mono"', 'monospace'],
          },
          colors: {
            brand: '#ffffff',
            surface: '#0a0a0a',
            surface2: '#141414',
            surface3: '#1f1f1f',
            border: '#2a2a2a',
            borderHover: '#404040'
          },
          backgroundImage: {
            'grid-pattern': "url('data:image/svg+xml,%3Csvg width=\\'40\\' height=\\'40\\' viewBox=\\'0 0 40 40\\' xmlns=\\'http://www.w3.org/2000/svg\\'%3E%3Cpath d=\\'M0 0h40v40H0V0zm1 1h38v38H1V1z\\' fill=\\'%23ffffff\\' fill-opacity=\\'0.02\\' fill-rule=\\'evenodd\\'/%3E%3C/svg%3E')"
          }
        }
      }
    }
  </script>
  <style>
    body { background-color: theme('colors.surface'); color: #ededed; overflow-x: hidden; transition: background-color 0.3s; }
    
    /* Dynamic UI Themes */
    body.theme-light { background-color: #f8f9fa; color: #111; }
    body.theme-light .glass-panel { background: rgba(255,255,255,0.9); border-color: #e5e5e5; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05); }
    body.theme-light .gradient-text { background: linear-gradient(135deg, #000 0%, #666 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; }
    body.theme-light .icon-card { background: #fff; border-color: #eaeaea; }
    body.theme-light .icon-card:hover { border-color: #ccc; box-shadow: 0 20px 40px -10px rgba(0,0,0,0.08); }
    body.theme-light input { background: #fff; color: #000; border-color: #ddd; }
    body.theme-light .text-white { color: #000; }
    body.theme-light .bg-surface3 { background: #f3f4f6; border-color: #e5e5e5; }
    body.theme-light .bg-surface2 { background: #fff; border-color: #e5e5e5; }
    body.theme-light .text-neutral-400 { color: #6b7280; }
    body.theme-light .text-neutral-500 { color: #9ca3af; }
    
    body.theme-checkered { background-color: #e5e5f7; background-image: repeating-linear-gradient(45deg, #d4d4e8 25%, transparent 25%, transparent 75%, #d4d4e8 75%, #d4d4e8), repeating-linear-gradient(45deg, #d4d4e8 25%, #e5e5f7 25%, #e5e5f7 75%, #d4d4e8 75%, #d4d4e8); background-position: 0 0, 10px 10px; background-size: 20px 20px; color: #111; }
    body.theme-checkered .glass-panel { background: rgba(255,255,255,0.85); backdrop-filter: blur(32px); border-color: rgba(255,255,255,0.5); box-shadow: 0 10px 30px rgba(0,0,0,0.05); }
    body.theme-checkered .icon-card { background: rgba(255,255,255,0.6); backdrop-filter: blur(12px); border-color: rgba(255,255,255,0.8); }
    body.theme-checkered .icon-card:hover { background: rgba(255,255,255,0.9); }
    body.theme-checkered .text-white { color: #000; }
    body.theme-checkered .gradient-text { background: linear-gradient(135deg, #000 0%, #444 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; }
    body.theme-checkered input { background: rgba(255,255,255,0.8); color: #000; border-color: rgba(0,0,0,0.1); }

    .glass-panel { 
      background: rgba(15, 15, 15, 0.6); 
      backdrop-filter: blur(24px); 
      -webkit-backdrop-filter: blur(24px); 
      border: 1px solid theme('colors.border');
    }
    
    .gradient-text {
      background: linear-gradient(135deg, #fff 0%, #737373 100%);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
    }

    ::-webkit-scrollbar { width: 8px; height: 8px; }
    ::-webkit-scrollbar-track { background: transparent; }
    ::-webkit-scrollbar-thumb { background: rgba(128,128,128,0.3); border-radius: 4px; }
    ::-webkit-scrollbar-thumb:hover { background: rgba(128,128,128,0.5); }

    .icon-card {
      background: theme('colors.surface2');
      border: 1px solid theme('colors.border');
      transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
      position: relative;
      overflow: hidden;
    }
    .icon-card:hover {
      border-color: theme('colors.borderHover');
      transform: translateY(-4px);
      box-shadow: 0 20px 40px -10px rgba(0,0,0,0.5);
    }
    
    .icon-wrapper {
      position: relative; z-index: 2;
      transition: transform 0.5s cubic-bezier(0.16, 1, 0.3, 1);
    }
    .icon-card:hover .icon-wrapper { transform: scale(1.15); }

    .copy-overlay {
      position: absolute; inset: 0; z-index: 10;
      background: rgba(10,10,10,0.8);
      backdrop-filter: blur(4px);
      display: flex; align-items: center; justify-content: center;
      opacity: 0; transition: opacity 0.2s;
    }
    body.theme-light .copy-overlay, body.theme-checkered .copy-overlay { background: rgba(255,255,255,0.8); }
    .icon-card:hover .copy-overlay { opacity: 1; }

    #toast {
      transform: translateY(150%) translateX(-50%);
      transition: transform 0.4s cubic-bezier(0.16, 1, 0.3, 1);
    }
    #toast.active { transform: translateY(0) translateX(-50%); }

    .theme-btn { opacity: 0.6; transition: all 0.2s; border-color: transparent; }
    .theme-btn:hover { opacity: 0.9; }
    .theme-btn.active { opacity: 1; transform: scale(1.05); }
    
    /* Ensure toast is always visible regardless of scrolling */
    .toast-container {
      position: fixed;
      bottom: 2rem;
      left: 50%;
      z-index: 9999;
      pointer-events: none;
    }
  </style>
</head>
<body class="flex flex-col relative theme-dark">
  <div class="fixed inset-0 pointer-events-none bg-grid-pattern opacity-30 [mask-image:linear-gradient(to_bottom,black,transparent_80%)] grid-bg-layer z-0"></div>

  <!-- Navbar -->
  <nav class="glass-panel sticky top-0 z-50 w-full border-t-0 border-l-0 border-r-0 h-16 flex items-center">
    <div class="w-full max-w-7xl mx-auto px-6 flex items-center justify-between">
      <div class="flex items-center gap-3">
        <svg width="24" height="24" viewBox="0 0 32 32" class="animate-[spin_8s_linear_infinite]">
          <circle cx="16" cy="16" r="14" fill="transparent" stroke="#888" stroke-width="1.5"/>
          <ellipse cx="16" cy="16" rx="10" ry="7" fill="none" stroke="#888" stroke-width="1.5" />
          <circle cx="16" cy="16" r="3.5" fill="#888"/>
        </svg>
        <span class="font-bold tracking-tight text-lg text-white">Logos API</span>
        <span class="px-2 py-0.5 rounded-full text-[10px] font-mono bg-neutral-500/20 text-neutral-400 ml-2 border border-neutral-500/30">v4.0 Cinematic</span>
      </div>
      <div class="hidden md:flex items-center gap-6 text-sm font-medium text-neutral-400">
        <a href="#api-reference" class="hover:text-white transition-colors">Documentation</a>
        <a href="#library" class="hover:text-white transition-colors">Library</a>
      </div>
    </div>
  </nav>

  <main class="flex-1 w-full max-w-7xl mx-auto px-6 pt-12 pb-32 relative z-10">
    
    <div class="flex flex-col lg:flex-row gap-12 mb-20 items-start">
      
      <!-- Left: Titles & Docs -->
      <div class="flex-1 pt-4 lg:pt-8 flex flex-col justify-center max-w-2xl">
        <h1 class="text-5xl md:text-6xl lg:text-7xl font-extrabold tracking-tighter mb-6 gradient-text leading-tight z-20 relative">
          Perfect Assets.<br/>Any Context.
        </h1>
        <p class="text-lg text-neutral-400 leading-relaxed font-light mb-10 max-w-lg z-20 relative">
          The ultimate dynamic SVG API. Mix colors, generate gradients, and adapt flawlessly to dark, light, or complex backgrounds with zero design effort.
        </p>

        <!-- Live Endpoint Card -->
        <div id="api-reference" class="glass-panel rounded-xl p-5 flex flex-col items-start gap-4 w-full border-l-4 border-l-blue-500 mt-2 z-20 relative">
          <div class="w-full h-32 md:h-48 rounded-lg bg-surface3 border border-border flex items-center justify-center preview-bg shadow-inner overflow-hidden" id="live-preview-box">
             <!-- Preview SVG Injected Here -->
          </div>
          <div class="font-mono text-xs overflow-hidden flex flex-col justify-center w-full mt-2">
            <div class="flex items-center justify-between mb-2">
              <span class="text-neutral-500 font-bold tracking-wider">LIVE ENDPOINT</span>
              <span class="bg-green-500/10 text-green-400 px-1.5 py-0.5 rounded text-[9px] uppercase">GET</span>
            </div>
            <code class="text-white truncate w-full block bg-surface2/50 p-3 rounded border border-border/50 text-[10px] md:text-xs overflow-x-auto" id="doc-url" title="Click to copy">...</code>
          </div>
        </div>
      </div>

      <!-- Right: Control Center -->
      <div class="flex-1 w-full lg:w-auto lg:sticky lg:top-24 z-30">
        <div class="glass-panel rounded-2xl p-6 md:p-8 shadow-2xl relative overflow-hidden flex flex-col">
          <div class="absolute top-0 right-0 w-64 h-64 bg-brand/5 rounded-full blur-3xl -translate-y-1/2 translate-x-1/2 pointer-events-none"></div>
          
          <div>
            <!-- Size & Scene Engine -->
            <div class="flex flex-col sm:flex-row gap-4 mb-4 relative z-10">
              <div class="flex-1">
                <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-2">Format (Size)</label>
                <select id="format-select" class="w-full font-mono bg-surface2 text-white border border-border rounded-lg px-4 py-2.5 text-sm outline-none focus:border-white transition-colors shadow-sm appearance-none">
                  <option value="favicon">Favicon (32x32)</option>
                  <option value="avatar">Avatar (512x512)</option>
                  <option value="og-card" selected>OG Card (1200x630)</option>
                  <option value="hero">Hero (1920x1080)</option>
                </select>
              </div>
              <div class="flex-1">
                <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-2">Scene Composition</label>
                <select id="scene-select" class="w-full font-mono bg-surface2 text-white border border-border rounded-lg px-4 py-2.5 text-sm outline-none focus:border-white transition-colors shadow-sm appearance-none">
                  <option value="pure">Pure (Logo Only)</option>
                  <option value="spotlight" selected>Spotlight (Glow)</option>
                  <option value="grid">Grid (Cyberpunk)</option>
                  <option value="split">Split (Logo + Text)</option>
                </select>
              </div>
            </div>
            
            <!-- Typography -->
            <div class="flex flex-col sm:flex-row gap-4 mb-4 relative z-10">
              <div class="flex-1">
                <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-2">Title</label>
                <input type="text" id="title-input" value="Logos API" class="w-full font-mono bg-surface2 text-white border border-border rounded-lg px-4 py-2.5 text-sm outline-none focus:border-white transition-colors shadow-sm" placeholder="e.g. My App">
              </div>
              <div class="flex-1">
                <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-2">Subtitle</label>
                <input type="text" id="subtitle-input" value="Generative Assets" class="w-full font-mono bg-surface2 text-white border border-border rounded-lg px-4 py-2.5 text-sm outline-none focus:border-white transition-colors shadow-sm" placeholder="e.g. Open Source">
              </div>
            </div>

            <!-- Color Inputs -->
            <div class="flex flex-col sm:flex-row gap-4 mb-8 relative z-10">
              <div class="flex-1">
                <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-2">Primary Color</label>
                <div class="relative">
                  <span class="absolute left-3 top-1/2 -translate-y-1/2 text-neutral-500 font-mono">#</span>
                  <input type="text" id="color1" value="amber" class="w-full font-mono bg-surface2 text-white border border-border rounded-lg pl-7 pr-4 py-2.5 text-sm outline-none focus:border-white transition-colors shadow-sm" placeholder="hex or name">
                </div>
              </div>
              <div class="flex-1">
                <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-2 flex justify-between">
                  <span>Gradient Color</span>
                  <span class="text-neutral-600 font-normal">Optional</span>
                </label>
                <div class="relative">
                  <span class="absolute left-3 top-1/2 -translate-y-1/2 text-neutral-500 font-mono">#</span>
                  <input type="text" id="color2" value="" class="w-full font-mono bg-surface2 text-white border border-border rounded-lg pl-7 pr-4 py-2.5 text-sm outline-none focus:border-white transition-colors shadow-sm" placeholder="leave empty for solid">
                </div>
              </div>
            </div>

            <!-- Theme Engine -->
            <div class="mb-6 relative z-10">
              <label class="block text-[11px] uppercase tracking-wider text-neutral-500 font-bold mb-3">Environment Engine</label>
              <div class="grid grid-cols-2 sm:grid-cols-4 gap-2">
                <button class="theme-btn active py-2.5 rounded-lg bg-[#0a0a0c] text-white border border-neutral-700 text-xs font-semibold shadow-sm flex flex-col items-center justify-center gap-1" data-theme="dark">
                  <div class="w-3 h-3 rounded-full bg-neutral-700 border border-neutral-500"></div> Dark
                </button>
                <button class="theme-btn py-2.5 rounded-lg bg-white text-black border border-neutral-300 text-xs font-semibold shadow-sm flex flex-col items-center justify-center gap-1" data-theme="light">
                  <div class="w-3 h-3 rounded-full bg-neutral-200 border border-neutral-400"></div> Light
                </button>
                <button class="theme-btn py-2.5 rounded-lg bg-[url('data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI4IiBoZWlnaHQ9IjgiPjxyZWN0IHdpZHRoPSI0IiBoZWlnaHQ9IjQiIGZpbGw9IiNlNWU1ZTUiLz48cmVjdCB4PSI0IiB5PSI0IiB3aWR0aD0iNCIgaGVpZ2h0PSI0IiBmaWxsPSIjZTVlNWU1Ii8+PC9zdmc+')] bg-white text-black border border-neutral-300 text-xs font-semibold shadow-sm flex flex-col items-center justify-center gap-1" data-theme="auto">
                  <div class="w-3 h-3 rounded-full border border-dashed border-neutral-500"></div> Clear
                </button>
                <button class="theme-btn py-2.5 rounded-lg bg-blue-600 text-white border border-blue-500 text-xs font-semibold shadow-sm flex flex-col items-center justify-center gap-1" data-theme="solid">
                  <div class="w-3 h-3 rounded-full bg-blue-400 border border-blue-300"></div> Solid
                </button>
              </div>
            </div>
          </div>

          <!-- Color Library -->
          <div class="relative z-10 pt-4 border-t border-border/50">
            <div class="flex items-center justify-between mb-3">
              <h3 class="font-bold text-xs uppercase tracking-wider text-neutral-400">Curated Palettes</h3>
            </div>
            <div class="flex flex-wrap gap-2 max-h-32 overflow-y-auto pr-2 pb-2" id="palette-container">
              ${generateColorPills(allPalettes)}
            </div>
          </div>

        </div>
      </div>
    </div>

    <!-- The Grid Library -->
    <div id="library" class="flex items-center gap-4 mb-8 pt-8">
      <h2 class="text-2xl font-bold text-white tracking-tight">Animation Library</h2>
      <div class="h-[1px] flex-1 bg-border"></div>
      <span class="bg-surface2 border border-border px-3 py-1 rounded-full text-neutral-400 text-xs font-mono font-medium">${Object.keys(animations).length} Variants</span>
    </div>

    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4 pb-12" id="grid-container">
      <!-- Injected via JS -->
    </div>

  </main>

  <!-- Global Toast (Fixed positioning logic to prevent cutoff) -->
  <div class="toast-container">
    <div id="toast" class="flex items-center gap-3 px-6 py-3 bg-white text-black rounded-full shadow-[0_20px_40px_rgba(0,0,0,0.4)] font-semibold text-sm border border-neutral-200">
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#22c55e" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
      <span id="toast-msg">Copied to clipboard</span>
    </div>
  </div>

  <script>
    const animations = ${JSON.stringify(Object.keys(animations))};
    const inputC1 = document.getElementById('color1');
    const inputC2 = document.getElementById('color2');
    const grid = document.getElementById('grid-container');
    const liveBox = document.getElementById('live-preview-box');
    const docUrl = document.getElementById('doc-url');
    
    // NEW V4 ELEMENTS
    const formatSelect = document.getElementById('format-select');
    const sceneSelect = document.getElementById('scene-select');
    const titleInput = document.getElementById('title-input');
    const subtitleInput = document.getElementById('subtitle-input');
    
    let currentTheme = 'dark';

    // Theme Engine Logic
    document.querySelectorAll('.theme-btn').forEach(btn => {
      btn.addEventListener('click', (e) => {
        document.querySelectorAll('.theme-btn').forEach(b => {
          b.classList.remove('active');
          b.style.borderColor = '';
        });
        
        btn.classList.add('active');
        if(btn.dataset.theme === 'dark') btn.style.borderColor = '#525252';
        if(btn.dataset.theme === 'light') btn.style.borderColor = '#a3a3a3';
        if(btn.dataset.theme === 'auto') btn.style.borderColor = '#a3a3a3';
        if(btn.dataset.theme === 'solid') btn.style.borderColor = '#93c5fd';
        
        currentTheme = btn.dataset.theme;
        
        // Update DOM body class for live environmental preview
        document.body.className = 'flex flex-col relative';
        if(currentTheme === 'light') document.body.classList.add('theme-light');
        else if(currentTheme === 'dark') document.body.classList.add('theme-dark');
        else if(currentTheme === 'auto') document.body.classList.add('theme-checkered');
        else document.body.classList.add('theme-dark');

        updateAll();
      });
    });

    // Palette Click Handler
    window.setColor = (colorName) => {
      inputC1.value = colorName;
      inputC2.value = ''; // Reset gradient on single color click
      updateAll();
    };

    // Render Library Grid
    grid.innerHTML = animations.map(anim => \`
      <div class="icon-card rounded-2xl p-5 cursor-pointer group flex flex-col items-center justify-between min-h-[140px]" data-anim="\${anim}" title="Click to copy API URL">
        <div class="copy-overlay rounded-2xl">
          <span class="bg-black/90 text-white text-[10px] uppercase tracking-widest font-bold px-3 py-1.5 rounded-full border border-white/20">Copy URL</span>
        </div>
        <div class="h-16 w-full flex items-center justify-center icon-wrapper drop-shadow-xl mt-2">
          <img src="" alt="\${anim}" class="w-10 h-10 lazy-svg" data-anim="\${anim}" />
        </div>
        <div class="w-full text-center mt-3 pt-3 border-t border-border/50 group-hover:border-border transition-colors">
          <span class="text-xs font-semibold text-neutral-300 group-hover:text-white capitalize tracking-wide">\${anim}</span>
        </div>
      </div>
    \`).join('');

    const svgs = document.querySelectorAll('.lazy-svg');
    const cards = document.querySelectorAll('.icon-card');

    // Core Update Logic
    function updateAll() {
      const c1 = inputC1.value.trim() || 'white';
      const c2 = inputC2.value.trim() || '';
      const format = formatSelect ? formatSelect.value : 'favicon';
      const scene = sceneSelect ? sceneSelect.value : 'pure';
      const title = titleInput ? encodeURIComponent(titleInput.value) : '';
      const subtitle = subtitleInput ? encodeURIComponent(subtitleInput.value) : '';
      
      const getPath = (anim) => {
        let base = \`/api/v4/render/\${format}/\${scene}/\${c1}\${c2 ? '/' + c2 : ''}/\${anim}.svg\`;
        
        const params = [];
        if(currentTheme !== 'auto') params.push(\`theme=\${currentTheme}\`);
        if(title) params.push(\`title=\${title}\`);
        if(subtitle) params.push(\`subtitle=\${subtitle}\`);
        
        if (params.length > 0) {
            base += '?' + params.join('&');
        }
        return base;
      };
      
      // Update Documentation Code Block
      const fullUrl = \`http://\${window.location.host}\${getPath('vortex')}\`;
      
      // Syntax highlighting for the URL
      let highlightedUrl = fullUrl
        .replace('/api/v4/render/', '<span class="text-neutral-500">/api/v4/render/</span>')
        .replace('.svg', '<span class="text-neutral-500">.svg</span>');
        
      docUrl.innerHTML = highlightedUrl;

      // Update Live Preview Box (Hero)
      liveBox.innerHTML = \`<img src="\${getPath('vortex')}" class="w-full h-full object-contain" />\`;

      // Update Library Grid
      svgs.forEach(img => {
        img.src = getPath(img.dataset.anim);
      });
    }

    // Input Listeners
    [inputC1, inputC2, formatSelect, sceneSelect, titleInput, subtitleInput].forEach(inp => {
      if(inp) {
        inp.addEventListener('input', () => {
          clearTimeout(inp.timer);
          inp.timer = setTimeout(updateAll, 150); // Small debounce
        });
      }
    });

    // Copy to Clipboard Logic
    cards.forEach(card => {
      card.addEventListener('click', () => {
        const anim = card.dataset.anim;
        const c1 = inputC1.value.trim() || 'white';
        const c2 = inputC2.value.trim() || '';
        const format = formatSelect ? formatSelect.value : 'favicon';
        const scene = sceneSelect ? sceneSelect.value : 'pure';
        const title = titleInput ? encodeURIComponent(titleInput.value) : '';
        const subtitle = subtitleInput ? encodeURIComponent(subtitleInput.value) : '';
        
        let url = \`http://\${window.location.host}/api/v4/render/\${format}/\${scene}/\${c1}\${c2 ? '/' + c2 : ''}/\${anim}.svg\`;
        
        const params = [];
        if(currentTheme !== 'auto') params.push(\`theme=\${currentTheme}\`);
        if(title) params.push(\`title=\${title}\`);
        if(subtitle) params.push(\`subtitle=\${subtitle}\`);
        
        if (params.length > 0) {
            url += '?' + params.join('&');
        }
        
        navigator.clipboard.writeText(url);
        
        // Toast Animation
        const toast = document.getElementById('toast');
        document.getElementById('toast-msg').innerHTML = \`Copied <span class="font-mono text-blue-600 bg-blue-50 px-1 rounded">\${anim}.svg</span> to clipboard\`;
        
        // Tactile card bounce
        card.style.transform = 'scale(0.92)';
        setTimeout(() => card.style.transform = '', 150);

        toast.classList.add('active');
        clearTimeout(toast.timer);
        toast.timer = setTimeout(() => toast.classList.remove('active'), 3000);
      });
    });

    // Click on Doc URL to copy
    docUrl.parentElement.addEventListener('click', () => {
      navigator.clipboard.writeText(docUrl.textContent);
      const toast = document.getElementById('toast');
      document.getElementById('toast-msg').textContent = 'Endpoint copied to clipboard!';
      toast.classList.add('active');
      clearTimeout(toast.timer);
      toast.timer = setTimeout(() => toast.classList.remove('active'), 3000);
    });
    docUrl.parentElement.classList.add('cursor-pointer', 'hover:border-blue-500/50', 'transition-colors');

    // Init
    updateAll();
  </script>
</body>
</html>`;
};

// --- CINEMATIC ENGINE (V4) ---

// 1. Formats (Tele)
const formats = {
  favicon: { w: 32, h: 32 },
  avatar: { w: 512, h: 512 },
  'og-card': { w: 1200, h: 630 },
  hero: { w: 1920, h: 1080 }
};

// 2. Scenes (Composizioni e Sfondi)
const getSceneDefs = (scene, colors) => {
  const { primaryHex, secondaryHex } = colors;
  
  if (scene === 'spotlight') {
    return `
      <radialGradient id="spotlightGrad" cx="50%" cy="50%" r="50%" fx="50%" fy="50%">
        <stop offset="0%" stop-color="#${primaryHex}" stop-opacity="0.3" />
        <stop offset="100%" stop-color="#${secondaryHex}" stop-opacity="0" />
      </radialGradient>
    `;
  }
  if (scene === 'grid') {
    return `
      <pattern id="gridPattern" width="40" height="40" patternUnits="userSpaceOnUse">
        <path d="M 40 0 L 0 0 0 40" fill="none" stroke="rgba(255,255,255,0.05)" stroke-width="1"/>
      </pattern>
      <radialGradient id="gridGlow" cx="50%" cy="100%" r="70%">
        <stop offset="0%" stop-color="#${primaryHex}" stop-opacity="0.15" />
        <stop offset="100%" stop-color="#000000" stop-opacity="0" />
      </radialGradient>
    `;
  }
  return '';
};

const getSceneBackground = (scene, width, height) => {
  if (scene === 'spotlight') {
    return `
      <rect width="100%" height="100%" fill="#050505" />
      <rect width="100%" height="100%" fill="url(#spotlightGrad)" />
    `;
  }
  if (scene === 'grid') {
    return `
      <rect width="100%" height="100%" fill="#000000" />
      <rect width="100%" height="100%" fill="url(#gridPattern)" />
      <rect width="100%" height="100%" fill="url(#gridGlow)" />
    `;
  }
  if (scene === 'pure') {
    return `<rect width="100%" height="100%" fill="transparent" />`;
  }
  // Default dark for others
  return `<rect width="100%" height="100%" fill="#0a0a0c" />`;
};

// 3. Logo Core Logic (Isoliamo il cuore dell'SVG originale)
const getLogoCore = (colors, animName, theme) => {
  const { primaryHex, secondaryHex } = colors;
  const animCss = animations[animName] || animations.static;
  
  const hasGradient = secondaryHex && secondaryHex !== primaryHex;
  const fillDef = hasGradient ? `url(#gradCore)` : `#${primaryHex}`;
  const strokeDef = hasGradient ? `url(#gradCore)` : `#${primaryHex}`;

  let coreBgColor = 'transparent'; 
  let dotColor = '#fff';
  let strokeWidth = '1.5';
  
  if (theme === 'dark') { coreBgColor = '#0a0a0c'; dotColor = '#ffffff'; } 
  else if (theme === 'light') { coreBgColor = '#ffffff'; dotColor = '#0a0a0c'; } 
  else if (theme === 'solid') { coreBgColor = fillDef; dotColor = '#ffffff'; strokeWidth = '0'; } 
  else if (theme === 'glass') { coreBgColor = 'rgba(255,255,255,0.1)'; dotColor = '#ffffff'; }

  return `
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="${coreBgColor}" stroke="${strokeDef}" stroke-width="${strokeWidth}" class="outer"/>
      <ellipse cx="16" cy="16" rx="10" ry="7" fill="none" stroke="${strokeDef}" stroke-width="1.5" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="${fillDef}"/>
        <circle cx="17.2" cy="14.8" r="1.2" fill="${dotColor}" opacity="0.9"/>
      </g>
    </g>
  `;
};

// 4. Typography Engine
const getTypography = (textParams, format, scene, colors) => {
  if (!textParams.title) return '';
  
  const { title, subtitle } = textParams;
  const { primaryHex } = colors;
  
  // Font scaling logic
  let titleSize = 48;
  let subSize = 24;
  let x = '50%';
  let y = '75%';
  let textAnchor = 'middle';
  
  if (format === 'og-card' && scene === 'split') {
    titleSize = 82;
    subSize = 36;
    x = '80px';
    y = '50%';
    textAnchor = 'start';
  } else if (format === 'hero') {
    titleSize = 120;
    subSize = 48;
    y = '80%';
  }

  const titleYOffset = subtitle ? -10 : 10;
  const subYOffset = subtitle ? (titleSize/2 + 20) : 0;

  return `
    <g class="typography">
      <text x="${x}" y="${y}" dy="${titleYOffset}px" text-anchor="${textAnchor}" fill="#ffffff" font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif" font-size="${titleSize}px" font-weight="800" letter-spacing="-0.03em">${title}</text>
      ${subtitle ? `<text x="${x}" y="${y}" dy="${subYOffset}px" text-anchor="${textAnchor}" fill="#${primaryHex}" font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif" font-size="${subSize}px" font-weight="500" opacity="0.8">${subtitle}</text>` : ''}
    </g>
  `;
};

// 5. Transform & Positioning Engine
const getLogoTransform = (format, scene) => {
  const { w, h } = formats[format];
  const logoBaseSize = 32;
  
  // Calcolo matematico per centrare e scalare
  let scale = 1;
  let tx = 0;
  let ty = 0;

  if (format === 'favicon') {
    return ''; // Scala 1:1, tx 0 ty 0
  }

  if (format === 'avatar') {
    scale = w / 64; // Lascia del margine
    tx = (w - (logoBaseSize * scale)) / 2;
    ty = (h - (logoBaseSize * scale)) / 2;
  } else if (format === 'og-card') {
    if (scene === 'split') {
      scale = 12;
      tx = w - (logoBaseSize * scale) - 100; // Allineato a destra
      ty = (h - (logoBaseSize * scale)) / 2;
    } else {
      // Centrato
      scale = 8;
      tx = (w - (logoBaseSize * scale)) / 2;
      ty = (h - (logoBaseSize * scale)) / 2 - 40; // Leggermente più in alto per far spazio al testo
    }
  } else if (format === 'hero') {
    scale = 15;
    tx = (w - (logoBaseSize * scale)) / 2;
    ty = (h - (logoBaseSize * scale)) / 2 - 80;
  }

  return `translate(${tx}, ${ty}) scale(${scale})`;
};

// 6. Master Cinematic Renderer
const renderCinematicSVG = (options) => {
  const { format = 'favicon', scene = 'pure', color1, color2, animation = 'static', theme = 'auto', text = {} } = options;
  
  const dim = formats[format] || formats.favicon;
  const primaryHex = getHex(color1);
  const secondaryHex = color2 ? getHex(color2) : primaryHex;
  const colors = { primaryHex, secondaryHex };
  
  const hasGradient = secondaryHex !== primaryHex;
  
  // Il gradiente core serve sempre per il logo base
  const coreDefs = hasGradient ? `
    <linearGradient id="gradCore" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" stop-color="#${primaryHex}" />
      <stop offset="100%" stop-color="#${secondaryHex}" />
    </linearGradient>
  ` : '';

  const animCss = animations[animation] || animations.static;

  return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${dim.w} ${dim.h}" width="100%" height="100%">
  <defs>
    ${coreDefs}
    ${getSceneDefs(scene, colors)}
  </defs>
  <style>
    :root { 
      --ease: cubic-bezier(0.25, 1, 0.5, 1);
      --ease-smooth: cubic-bezier(0.65, 0, 0.35, 1);
    }
    .outer, .orbit, .core { transform-origin: 16px 16px; }
    ${animCss}
  </style>
  
  <!-- Background Layer -->
  ${getSceneBackground(scene, dim.w, dim.h)}
  
  <!-- Core Layer (The Logo) -->
  <g transform="${getLogoTransform(format, scene)}">
    ${getLogoCore(colors, animation, theme)}
  </g>
  
  <!-- Typography Layer -->
  ${getTypography(text, format, scene, colors)}
</svg>`;
};

// --- ROUTER V4 (CINEMATIC) ---
const server = http.createServer((req, res) => {
  res.setHeader('Access-Control-Allow-Origin', '*');

  try {
    const url = new URL(req.url, `http://${req.headers.host}`);
    
    const pathParts = url.pathname.split('/').filter(Boolean);
    
    // BACKWARD COMPATIBILITY & NEW CINEMATIC ROUTER
    if (pathParts[0] === 'api') {
      
      // V4 Cinematic Engine: /api/v4/render/:format/:scene/:color1/:color2?/:animation
      if (pathParts[1] === 'v4' && pathParts[2] === 'render') {
        const format = pathParts[3] || 'favicon';
        const scene = pathParts[4] || 'pure';
        
        let c1 = 'white', c2 = null, anim = 'static';
        if (pathParts.length === 7) {
          // e.g. /api/v4/render/og-card/spotlight/cyber/vortex.svg
          c1 = pathParts[5];
          anim = pathParts[6].replace('.svg', '');
        } else if (pathParts.length === 8) {
          // e.g. /api/v4/render/og-card/spotlight/cyber/laser/vortex.svg
          c1 = pathParts[5];
          c2 = pathParts[6];
          anim = pathParts[7].replace('.svg', '');
        }

        const theme = url.searchParams.get('theme') || 'auto';
        const title = url.searchParams.get('title') || '';
        const subtitle = url.searchParams.get('subtitle') || '';

        const svg = renderCinematicSVG({
          format, scene, color1: c1, color2: c2, animation: anim, theme, text: { title, subtitle }
        });

        res.setHeader('Content-Type', 'image/svg+xml');
        res.setHeader('Cache-Control', 'public, max-age=31536000');
        res.writeHead(200);
        res.end(svg);
        return;
      }

      // Legacy V3 Router
      if (pathParts[2] === 'favicon') {
        let c1 = 'white', c2 = null, anim = 'static';

        if (pathParts.length === 5) {
          c1 = pathParts[3];
          anim = pathParts[4].replace('.svg', '');
        } else if (pathParts.length === 6) {
          c1 = pathParts[3];
          c2 = pathParts[4];
          anim = pathParts[5].replace('.svg', '');
        }

        const theme = url.searchParams.get('theme') || 'auto';
        
        const svg = renderCinematicSVG({
          format: 'favicon', scene: 'pure', color1: c1, color2: c2, animation: anim, theme
        });
        
        res.setHeader('Content-Type', 'image/svg+xml');
        res.setHeader('Cache-Control', 'public, max-age=31536000');
        res.writeHead(200);
        res.end(svg);
        return;
      }
    }
    
    if (url.pathname === '/') {
      res.setHeader('Content-Type', 'text/html');
      res.writeHead(200);
      res.end(getDashboardHtml());
      return;
    }

    res.writeHead(404);
    res.end('Not Found');
  } catch (e) {
    res.writeHead(500);
    res.end('Internal Server Error');
  }
});

server.listen(PORT, () => {
  console.log(`\n🚀 LOGOS API V4.0 (CINEMATIC) 🚀`);
  console.log(`Server running at: http://localhost:${PORT}`);
});