package svg

import "fmt"

// TextureDefs returns SVG <filter> definitions for material textures.
func TextureDefs(texture string) string {
	switch texture {
	case "grain":
		return `
    <filter id="grain" x="0" y="0" width="100%" height="100%">
      <feTurbulence type="fractalNoise" baseFrequency="0.65" numOctaves="3" stitchTiles="stitch" result="noise"/>
      <feColorMatrix type="saturate" values="0" in="noise" result="mono"/>
      <feBlend in="SourceGraphic" in2="mono" mode="multiply" result="blend"/>
      <feComponentTransfer in="blend">
        <feFuncA type="linear" slope="0.4"/>
      </feComponentTransfer>
    </filter>`
	case "glass":
		return `
    <filter id="glass" x="-10%" y="-10%" width="120%" height="120%">
      <feGaussianBlur in="SourceGraphic" stdDeviation="0.8" result="blur"/>
      <feTurbulence type="fractalNoise" baseFrequency="0.04" numOctaves="2" result="noise"/>
      <feDisplacementMap in="blur" in2="noise" scale="3" xChannelSelector="R" yChannelSelector="G"/>
    </filter>`
	case "noise":
		return `
    <filter id="noise" x="0" y="0" width="100%" height="100%">
      <feTurbulence type="turbulence" baseFrequency="0.9" numOctaves="1" stitchTiles="stitch" result="noise"/>
      <feColorMatrix type="saturate" values="0" in="noise" result="mono"/>
      <feBlend in="SourceGraphic" in2="mono" mode="soft-light" result="blend"/>
      <feComponentTransfer in="blend">
        <feFuncA type="linear" slope="0.25"/>
      </feComponentTransfer>
    </filter>`
	case "glitch":
		// Chromatic aberration: splits RGB channels with offset displacement
		return `
    <filter id="glitch" x="-5%" y="-5%" width="110%" height="110%">
      <feOffset in="SourceGraphic" dx="1" dy="0" result="red"/>
      <feOffset in="SourceGraphic" dx="-1" dy="0" result="blue"/>
      <feColorMatrix in="red" type="matrix" values="1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0" result="r"/>
      <feColorMatrix in="blue" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 1 0" result="b"/>
      <feColorMatrix in="SourceGraphic" type="matrix" values="0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 1 0" result="g"/>
      <feBlend in="r" in2="g" mode="screen" result="rg"/>
      <feBlend in="rg" in2="b" mode="screen"/>
    </filter>`
	case "shadow":
		// Soft multi-layer drop shadow (organic neon glow)
		return `
    <filter id="shadow" x="-20%" y="-20%" width="140%" height="140%">
      <feGaussianBlur in="SourceAlpha" stdDeviation="2" result="s1"/>
      <feGaussianBlur in="SourceAlpha" stdDeviation="6" result="s2"/>
      <feGaussianBlur in="SourceAlpha" stdDeviation="12" result="s3"/>
      <feMerge>
        <feMergeNode in="s3"/>
        <feMergeNode in="s2"/>
        <feMergeNode in="s1"/>
        <feMergeNode in="SourceGraphic"/>
      </feMerge>
    </filter>`
	case "neon":
		// Neon glow: colored outer glow using flood + composite
		return `
    <filter id="neon" x="-30%" y="-30%" width="160%" height="160%">
      <feGaussianBlur in="SourceGraphic" stdDeviation="3" result="glow1"/>
      <feGaussianBlur in="SourceGraphic" stdDeviation="8" result="glow2"/>
      <feMerge>
        <feMergeNode in="glow2"/>
        <feMergeNode in="glow1"/>
        <feMergeNode in="SourceGraphic"/>
      </feMerge>
    </filter>`
	default:
		return ""
	}
}

// TextureFilterAttr returns the filter="url(#...)" attribute if texture is active.
func TextureFilterAttr(texture string) string {
	switch texture {
	case "grain", "glass", "noise", "glitch", "shadow", "neon":
		return fmt.Sprintf(` filter="url(#%s)"`, texture)
	default:
		return ""
	}
}

// ValidTexture checks if a texture name is valid.
func ValidTexture(name string) bool {
	switch name {
	case "none", "grain", "glass", "noise", "glitch", "shadow", "neon", "":
		return true
	}
	return false
}
