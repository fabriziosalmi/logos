const fs = require('fs');
const path = require('path');

const colors = {
  amber: '#eab308',
  blue: '#3b82f6',
  cyan: '#06b6d4',
  green: '#22c55e',
  indigo: '#6366f1',
  orange: '#f97316',
  purple: '#a855f7',
  rose: '#f43f5e'
};

const animations = {
  spin: `
    .orbit {
      transform-origin: 16px 16px;
      animation: spin 4s linear infinite;
    }
    @keyframes spin {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
    }
  `,
  pulse: `
    .core {
      transform-origin: 16px 16px;
      animation: pulse 2s ease-in-out infinite;
    }
    @keyframes pulse {
      0%, 100% { transform: scale(1); }
      50% { transform: scale(1.3); }
    }
  `,
  blink: `
    .orbit {
      transform-origin: 16px 16px;
      animation: blink 4s ease-in-out infinite;
    }
    @keyframes blink {
      0%, 90%, 100% { transform: scaleY(1); }
      95% { transform: scaleY(0.1); }
    }
  `,
  float: `
    .core {
      animation: float 3s ease-in-out infinite;
    }
    @keyframes float {
      0%, 100% { transform: translateY(0); }
      50% { transform: translateY(-2px); }
    }
  `
};

const template = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32">
  <style>
{{ANIMATION}}
  </style>
  <circle cx="16" cy="16" r="14" fill="#0a0a0c" stroke="{{COLOR}}" stroke-width="1.5" class="outer"/>
  <ellipse cx="16" cy="16" rx="10" ry="7" fill="none" stroke="{{COLOR}}" stroke-width="1.5" class="orbit"/>
  <g class="core">
    <circle cx="16" cy="16" r="3.5" fill="{{COLOR}}"/>
    <circle cx="17.2" cy="14.8" r="1.2" fill="#fff" opacity="0.9"/>
  </g>
</svg>`;

const outputDir = path.join(__dirname, 'animated');

// Create animated directory if it doesn't exist
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir);
}

// Generate files
for (const [colorName, hexCode] of Object.entries(colors)) {
  for (const [animName, css] of Object.entries(animations)) {
    let svgContent = template
      .replace(/\{\{COLOR\}\}/g, hexCode)
      .replace('{{ANIMATION}}', css);
      
    const fileName = 'favicon-' + colorName + '-' + animName + '.svg';
    const filePath = path.join(outputDir, fileName);
    
    fs.writeFileSync(filePath, svgContent, 'utf8');
    console.log('Generated: ' + fileName);
  }
}

// Also generate the static versions using the new structure for consistency
for (const [colorName, hexCode] of Object.entries(colors)) {
    let svgContent = template
      .replace(/\{\{COLOR\}\}/g, hexCode)
      .replace('{{ANIMATION}}', '/* static */');
      
    const fileName = 'favicon-' + colorName + '-static.svg';
    const filePath = path.join(outputDir, fileName);
    
    fs.writeFileSync(filePath, svgContent, 'utf8');
    console.log('Generated: ' + fileName);
}

// And finally, create a pure 'master' file that can be used programmatically
const masterPath = path.join(__dirname, 'favicon-master.svg');
fs.writeFileSync(masterPath, template, 'utf8');
console.log('Generated master template: favicon-master.svg');

console.log('All variations generated successfully in the "animated" folder!');
