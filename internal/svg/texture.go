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

	// === NEW TEXTURES (13 more) ===

	case "emboss-tex":
		// Embossed/raised surface
		return `
    <filter id="emboss-tex" x="0" y="0" width="100%" height="100%">
      <feConvolveMatrix order="3" kernelMatrix="-2 -1 0 -1 1 1 0 1 2" preserveAlpha="true"/>
    </filter>`
	case "sharpen":
		return `
    <filter id="sharpen" x="0" y="0" width="100%" height="100%">
      <feConvolveMatrix order="3" kernelMatrix="0 -1 0 -1 5 -1 0 -1 0" preserveAlpha="true"/>
    </filter>`
	case "erode":
		return `
    <filter id="erode" x="0" y="0" width="100%" height="100%">
      <feMorphology operator="erode" radius="0.5"/>
    </filter>`
	case "dilate":
		return `
    <filter id="dilate" x="-5%" y="-5%" width="110%" height="110%">
      <feMorphology operator="dilate" radius="0.5"/>
    </filter>`
	case "pencil":
		// Pencil sketch: displacement + high-freq noise
		return `
    <filter id="pencil" x="-5%" y="-5%" width="110%" height="110%">
      <feTurbulence type="fractalNoise" baseFrequency="0.5" numOctaves="5" result="noise"/>
      <feDisplacementMap in="SourceGraphic" in2="noise" scale="1.5" xChannelSelector="R" yChannelSelector="G"/>
    </filter>`
	case "watercolor":
		// Soft watercolor bleed
		return `
    <filter id="watercolor" x="-10%" y="-10%" width="120%" height="120%">
      <feTurbulence type="turbulence" baseFrequency="0.03" numOctaves="3" result="noise"/>
      <feDisplacementMap in="SourceGraphic" in2="noise" scale="8" xChannelSelector="R" yChannelSelector="G" result="displaced"/>
      <feGaussianBlur in="displaced" stdDeviation="0.5"/>
    </filter>`
	case "halftone":
		// Halftone dot pattern
		return `
    <filter id="halftone" x="0" y="0" width="100%" height="100%">
      <feTurbulence type="turbulence" baseFrequency="1.5" numOctaves="1" result="dots"/>
      <feColorMatrix type="saturate" values="0" in="dots" result="bw"/>
      <feComponentTransfer in="bw" result="threshold"><feFuncA type="discrete" tableValues="0 1"/></feComponentTransfer>
      <feComposite in="SourceGraphic" in2="threshold" operator="in"/>
    </filter>`
	case "outline-glow":
		// Outline with soft glow behind
		return `
    <filter id="outline-glow" x="-20%" y="-20%" width="140%" height="140%">
      <feMorphology in="SourceAlpha" operator="dilate" radius="2" result="thick"/>
      <feGaussianBlur in="thick" stdDeviation="3" result="blur"/>
      <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
    </filter>`
	case "duotone-filter":
		// Duotone color mapping
		return `
    <filter id="duotone-filter" x="0" y="0" width="100%" height="100%">
      <feColorMatrix type="saturate" values="0"/>
      <feComponentTransfer>
        <feFuncR type="table" tableValues="0 1"/>
        <feFuncG type="table" tableValues="0 0.5"/>
        <feFuncB type="table" tableValues="0.2 0.8"/>
      </feComponentTransfer>
    </filter>`
	case "vhs":
		// VHS distortion: horizontal displacement + color shift
		return `
    <filter id="vhs" x="-5%" y="-5%" width="110%" height="110%">
      <feTurbulence type="fractalNoise" baseFrequency="0.01 0.2" numOctaves="1" result="noise"/>
      <feDisplacementMap in="SourceGraphic" in2="noise" scale="4" xChannelSelector="R" yChannelSelector="G"/>
    </filter>`
	case "burn":
		// Burn/overexpose edges
		return `
    <filter id="burn" x="0" y="0" width="100%" height="100%">
      <feComponentTransfer>
        <feFuncR type="gamma" amplitude="1.5" exponent="0.8" offset="0"/>
        <feFuncG type="gamma" amplitude="1.2" exponent="0.8" offset="0"/>
        <feFuncB type="gamma" amplitude="0.8" exponent="1.2" offset="0"/>
      </feComponentTransfer>
    </filter>`
	case "frost":
		// Frost/ice crystalline
		return `
    <filter id="frost" x="-10%" y="-10%" width="120%" height="120%">
      <feTurbulence type="fractalNoise" baseFrequency="0.08" numOctaves="4" result="noise"/>
      <feDisplacementMap in="SourceGraphic" in2="noise" scale="5" xChannelSelector="R" yChannelSelector="G" result="frost"/>
      <feGaussianBlur in="frost" stdDeviation="0.3"/>
    </filter>`
	case "ripple":
		// Ripple/water distortion
		return `
    <filter id="ripple" x="-10%" y="-10%" width="120%" height="120%">
      <feTurbulence type="turbulence" baseFrequency="0.05" numOctaves="2" result="wave"/>
      <feDisplacementMap in="SourceGraphic" in2="wave" scale="6" xChannelSelector="R" yChannelSelector="G"/>
    </filter>`
	default:
		return ""
	}
}

// TextureFilterAttr returns the filter="url(#...)" attribute if texture is active.
func TextureFilterAttr(texture string) string {
	switch texture {
	case "grain", "glass", "noise", "glitch", "shadow", "neon",
		"emboss-tex", "sharpen", "erode", "dilate", "pencil", "watercolor",
		"halftone", "outline-glow", "duotone-filter", "vhs", "burn", "frost", "ripple":
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
