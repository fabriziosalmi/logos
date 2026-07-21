import { defineConfig } from 'vitepress'

export default defineConfig({
  head: [
    // Everything this site loads is first-party. 'unsafe-inline' is required
    // because VitePress emits an inline appearance script and inline styles.
    // Applied to the built site only: `vitepress dev` serves HMR over a
    // websocket, which a strict connect-src would block as soon as the dev
    // server is not same-origin (--host, or a custom server.hmr.port).
    ...(process.env.NODE_ENV === 'production'
      ? [
          [
            'meta',
            {
              'http-equiv': 'Content-Security-Policy',
              content:
                "default-src 'self'; script-src 'self' 'unsafe-inline'; " +
                "style-src 'self' 'unsafe-inline'; img-src 'self' data:; " +
                "font-src 'self'; connect-src 'self'; base-uri 'self'; form-action 'self'",
            },
          ] as [string, Record<string, string>],
        ]
      : []),
  ],
  lang: 'en-US',
  title: 'Logos API',
  description: 'SVG asset generation API and embedded dashboard',
  base: '/logos/',
  themeConfig: {
    footer: {
      message:
        '<a href="https://fabriziosalmi.github.io/privacy">Privacy &amp; legal</a>',
    },
    search: { provider: 'local' },
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting started', link: '/getting-started' },
      { text: 'API', link: '/api/render' },
      { text: 'Config', link: '/configuration' }
    ],
    sidebar: [
      {
        text: 'Overview',
        items: [
          { text: 'What is Logos', link: '/' },
          { text: 'Getting started', link: '/getting-started' },
          { text: 'Dashboard', link: '/dashboard' }
        ]
      },
      {
        text: 'Configuration',
        items: [
          { text: 'Configuration file', link: '/configuration' },
          { text: 'Security headers', link: '/security' },
          { text: 'Caching', link: '/caching' }
        ]
      },
      {
        text: 'API',
        items: [
          { text: 'Render', link: '/api/render' },
          { text: 'Generative', link: '/api/generative' },
          { text: 'Icon packs', link: '/api/icons' },
          { text: 'App shortcuts', link: '/api/apps' },
          { text: 'Admin endpoints', link: '/api/admin' }
        ]
      },
      {
        text: 'Operations',
        items: [
          { text: 'CLI', link: '/cli' },
          { text: 'Docker', link: '/docker' },
          { text: 'Troubleshooting', link: '/troubleshooting' }
        ]
      }
    ],
    editLink: {
      pattern: 'https://github.com/fabriziosalmi/logos/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },
    lastUpdated: {
      text: 'Last updated'
    },
    docFooter: {
      prev: 'Previous',
      next: 'Next'
    },
    socialLinks: [{ icon: 'github', link: 'https://github.com/fabriziosalmi/logos' }]
  }
})
