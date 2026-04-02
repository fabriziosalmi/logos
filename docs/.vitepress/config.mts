import { defineConfig } from 'vitepress'

export default defineConfig({
  lang: 'it-IT',
  title: 'Logos API',
  description: 'Dynamic SVG asset generation service',
  base: '/logos/',
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'API', link: '/api' }
    ],
    sidebar: [
      { text: 'Introduzione', link: '/' },
      { text: 'API', link: '/api' }
    ],
    socialLinks: [{ icon: 'github', link: 'https://github.com/fabriziosalmi/logos' }]
  }
})
