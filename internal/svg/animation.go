package svg

// Animations maps animation names to their CSS definitions.
var Animations = map[string]string{
	// STATICS
	"static": "",

	// ZEN & MINIMAL
	"zen": `
    .core { animation: zen-core 6s var(--ease) infinite alternate; }
    .orbit { animation: zen-orbit 12s linear infinite; opacity: 0.8; }
    @keyframes zen-core { 0% { transform: scale(0.95) translateY(1px); opacity: 0.8; } 100% { transform: scale(1.05) translateY(-1px); opacity: 1; } }
    @keyframes zen-orbit { 100% { transform: rotate(360deg); } }`,
	"breathe": `
    .outer, .orbit, .core { animation: breathe 5s var(--ease) infinite alternate; }
    @keyframes breathe { 0% { transform: scale(0.98); } 100% { transform: scale(1.04); } }`,
	"levitate": `
    .core { animation: levitate 4s var(--ease-smooth) infinite alternate; }
    .orbit { animation: levitate-shadow 4s var(--ease-smooth) infinite alternate; }
    @keyframes levitate { 0% { transform: translateY(0); } 100% { transform: translateY(-3px); } }
    @keyframes levitate-shadow { 0% { transform: scale(1); opacity: 1; } 100% { transform: scale(0.95) translateY(2px); opacity: 0.6; } }`,
	"glimmer": `
    .core { animation: glimmer 3s ease-in-out infinite alternate; }
    @keyframes glimmer { 0%, 100% { opacity: 1; } 50% { opacity: 0.3; } }`,

	// SPINS & ORBITS
	"spin": `
    .orbit { animation: spin 5s linear infinite; }
    @keyframes spin { 100% { transform: rotate(360deg); } }`,
	"spin-fast": `
    .orbit { animation: spin 1.5s linear infinite; }
    @keyframes spin { 100% { transform: rotate(360deg); } }`,
	"smooth-spin": `
    .orbit { animation: smooth-spin 8s var(--ease) infinite; }
    @keyframes smooth-spin { 0% { transform: rotate(0deg); } 50% { transform: rotate(180deg); } 100% { transform: rotate(360deg); } }`,
	"orbit-chase": `
    .orbit { stroke-dasharray: 20 10; animation: spin 2s linear infinite; }
    @keyframes spin { 100% { transform: rotate(360deg); } }`,
	"compass": `
    .orbit { animation: compass 4s cubic-bezier(0.68, -0.55, 0.265, 1.55) infinite; }
    @keyframes compass { 0%, 20% { transform: rotate(0); } 25%, 45% { transform: rotate(90deg); } 50%, 70% { transform: rotate(180deg); } 75%, 95% { transform: rotate(270deg); } 100% { transform: rotate(360deg); } }`,
	"gyro": `
    .outer { animation: spin 8s linear infinite; stroke-dasharray: 40 5; }
    .orbit { animation: spin-rev 6s linear infinite; }
    @keyframes spin { 100% { transform: rotate(360deg); } }
    @keyframes spin-rev { 100% { transform: rotate(-360deg); } }`,
	"satellite": `
    .orbit { stroke-dasharray: 4 30; animation: spin 1.5s linear infinite; stroke-linecap: round; }
    @keyframes spin { 100% { transform: rotate(360deg); } }`,
	"eclipse": `
    .orbit { animation: eclipse 4s ease-in-out infinite; }
    @keyframes eclipse { 0%, 100% { transform: scaleY(1); } 50% { transform: scaleY(0); } }`,

	// PULSES & HEARTS
	"pulse": `
    .core { animation: pulse 3s var(--ease) infinite; }
    @keyframes pulse { 0%, 100% { transform: scale(1); opacity: 1; } 50% { transform: scale(1.25); opacity: 0.8; } }`,
	"heartbeat": `
    .core { animation: heartbeat 2s var(--ease) infinite; }
    @keyframes heartbeat { 0%, 100% { transform: scale(1); } 15%, 45% { transform: scale(1.2); } 30%, 60% { transform: scale(1); } }`,
	"pulse-ring": `
    .outer { animation: pulse-ring 2.5s cubic-bezier(0.215, 0.61, 0.355, 1) infinite; }
    @keyframes pulse-ring { 0% { transform: scale(0.9); opacity: 1; } 100% { transform: scale(1.2); opacity: 0; } }`,
	"strobe": `
    .core { animation: strobe 1.5s steps(2, start) infinite; }
    @keyframes strobe { 0%, 100% { opacity: 1; } 50% { opacity: 0; } }`,
	"nova": `
    .core { animation: nova 3s ease-out infinite; }
    @keyframes nova { 0% { transform: scale(1); opacity: 1; } 100% { transform: scale(2.5); opacity: 0; } }`,
	"elastic": `
    .core { animation: elastic 2s ease-in-out infinite; }
    @keyframes elastic { 0%, 100% { transform: scale(1, 1); } 50% { transform: scale(1.15, 0.85); } }`,

	// COMPLEX & 3D
	"flip": `
    .orbit { animation: flip 4s var(--ease-smooth) infinite; perspective: 100px; }
    @keyframes flip { 0% { transform: rotateY(0); } 50% { transform: rotateY(180deg); } 100% { transform: rotateY(360deg); } }`,
	"orbit-tilt": `
    .orbit { animation: orbit-tilt 4s ease-in-out infinite alternate; }
    @keyframes orbit-tilt { 0% { transform: rotateX(0); } 100% { transform: rotateX(60deg); } }`,
	"vortex": `
    .outer { animation: spin-slow 10s linear infinite; }
    .orbit { animation: spin-mid 5s linear infinite; }
    .core { animation: spin-fast 2s linear infinite; stroke-dasharray: 2 4; }
    @keyframes spin-slow { 100% { transform: rotate(360deg); } }
    @keyframes spin-mid { 100% { transform: rotate(360deg); } }
    @keyframes spin-fast { 100% { transform: rotate(360deg); } }`,
	"harmony": `
    .core { animation: breathe 4s ease-in-out infinite alternate; }
    .orbit { animation: spin 8s linear infinite; }
    @keyframes breathe { 0% { transform: scale(0.98); } 100% { transform: scale(1.04); } }
    @keyframes spin { 100% { transform: rotate(360deg); } }`,
	"sync": `
    .outer, .orbit, .core { animation: sync 4s ease-in-out infinite; }
    @keyframes sync { 0%, 100% { transform: scale(1) rotate(0); } 50% { transform: scale(1.1) rotate(180deg); } }`,
	"sway": `
    .outer, .orbit, .core { animation: sway 4s ease-in-out infinite alternate; }
    @keyframes sway { 0% { transform: rotate(-15deg); } 100% { transform: rotate(15deg); } }`,

	// SCI-FI & TECH
	"radar": `
    .outer { animation: radar 2.5s var(--ease) infinite; }
    @keyframes radar { 0% { transform: scale(0.8); opacity: 1; stroke-width: 2; } 100% { transform: scale(1.5); opacity: 0; stroke-width: 0; } }`,
	"radar-sweep": `
    .outer { animation: sweep 4s linear infinite; stroke-dasharray: 40 80; stroke-dashoffset: 0; }
    @keyframes sweep { 100% { stroke-dashoffset: -120; transform: rotate(360deg); } }`,
	"signal": `
    .outer { animation: signal 2s ease-out infinite; }
    @keyframes signal { 0% { transform: scale(0.9); opacity: 1; stroke-width: 3; } 100% { transform: scale(1.3); opacity: 0; stroke-width: 1; } }`,
	"glow": `
    .outer { animation: glow 3s var(--ease) infinite alternate; }
    @keyframes glow { 0% { stroke-width: 1.5; stroke-opacity: 1; } 100% { stroke-width: 4.5; stroke-opacity: 0.3; } }`,
	"aurora": `
    .outer { animation: aurora 3s linear infinite; }
    @keyframes aurora { 0% { stroke-width: 1.5; opacity: 1; transform: scale(1); } 100% { stroke-width: 10; opacity: 0; transform: scale(1.2); } }`,
	"nebula": `
    .outer { animation: nebula-outer 4s ease-in-out infinite alternate; }
    .orbit { animation: spin 6s linear infinite; stroke-width: 0.5; }
    @keyframes nebula-outer { 0% { stroke-width: 1.5; transform: scale(1); opacity: 1; } 100% { stroke-width: 4; transform: scale(1.1); opacity: 0.5; } }
    @keyframes spin { 100% { transform: rotate(360deg); } }`,
	"corona": `
    .outer { animation: corona 3s ease-in-out infinite alternate; }
    @keyframes corona { 0% { stroke-dasharray: 20 5; transform: scale(1); } 100% { stroke-dasharray: 5 20; transform: scale(1.05); } }`,
	"ripple-core": `
    .core { animation: ripple-core 2s cubic-bezier(0.4, 0, 0.2, 1) infinite; }
    @keyframes ripple-core { 0% { transform: scale(1); opacity: 1; } 100% { transform: scale(2); opacity: 0; } }`,

	// MORPH
	"morph": `
    .outer { animation: morph 4s ease-in-out infinite; }
    @keyframes morph { 0%, 100% { rx: 14; ry: 14; } 50% { rx: 10; ry: 16; } }`,
	"morph-blob": `
    .outer { animation: morph-blob 5s ease-in-out infinite; }
    @keyframes morph-blob { 0%, 100% { rx: 14; ry: 14; transform: rotate(0); } 33% { rx: 12; ry: 16; transform: rotate(10deg); } 66% { rx: 16; ry: 12; transform: rotate(-10deg); } }`,
	"morph-crystal": `
    .outer { animation: morph-crystal 3s linear infinite; }
    @keyframes morph-crystal { 0% { rx: 14; ry: 14; } 25% { rx: 10; ry: 14; } 50% { rx: 14; ry: 10; } 75% { rx: 10; ry: 14; } 100% { rx: 14; ry: 14; } }`,

	// BOUNCE
	"bounce": `
    .core { animation: bounce 2s ease-in-out infinite; }
    @keyframes bounce { 0%, 100% { transform: translateY(0); } 50% { transform: translateY(-5px); } }`,
	"bounce-drop": `
    .core { animation: bounce-drop 1.5s cubic-bezier(0.36, 0.07, 0.19, 0.97) infinite; }
    @keyframes bounce-drop { 0% { transform: translateY(-8px); } 50% { transform: translateY(0); } 60% { transform: translateY(-3px); } 80% { transform: translateY(0); } 100% { transform: translateY(-8px); } }`,
	"trampoline": `
    .core { animation: trampoline 1.2s ease-in-out infinite; }
    @keyframes trampoline { 0%, 100% { transform: scaleY(1) translateY(0); } 30% { transform: scaleY(0.8) translateY(2px); } 60% { transform: scaleY(1.1) translateY(-4px); } }`,

	// SHAKE
	"shake": `
    .core { animation: shake 0.8s ease-in-out infinite; }
    @keyframes shake { 0%, 100% { transform: translateX(0); } 25% { transform: translateX(-2px); } 75% { transform: translateX(2px); } }`,
	"jitter": `
    .core { animation: jitter 0.3s linear infinite; }
    @keyframes jitter { 0% { transform: translate(0,0); } 25% { transform: translate(-1px,1px); } 50% { transform: translate(1px,-1px); } 75% { transform: translate(-1px,-1px); } 100% { transform: translate(1px,1px); } }`,
	"earthquake": `
    .outer, .orbit, .core { animation: earthquake 0.5s ease-in-out infinite; }
    @keyframes earthquake { 0%, 100% { transform: translate(0,0) rotate(0); } 25% { transform: translate(-2px,1px) rotate(-1deg); } 50% { transform: translate(2px,-1px) rotate(1deg); } 75% { transform: translate(-1px,2px) rotate(-0.5deg); } }`,

	// SWING
	"swing": `
    .outer, .orbit, .core { animation: swing 3s ease-in-out infinite; }
    @keyframes swing { 0%, 100% { transform: rotate(0); } 25% { transform: rotate(10deg); } 75% { transform: rotate(-10deg); } }`,
	"pendulum": `
    .orbit { animation: pendulum 2.5s ease-in-out infinite; transform-origin: 16px 2px; }
    @keyframes pendulum { 0%, 100% { transform: rotate(-20deg); } 50% { transform: rotate(20deg); } }`,
	"wave-swing": `
    .core { animation: wave-swing 3s ease-in-out infinite; }
    @keyframes wave-swing { 0%, 100% { transform: translateX(0) rotate(0); } 25% { transform: translateX(3px) rotate(5deg); } 75% { transform: translateX(-3px) rotate(-5deg); } }`,

	// ZOOM
	"zoom-in": `
    .outer, .orbit, .core { animation: zoom-in 3s ease-in-out infinite; }
    @keyframes zoom-in { 0% { transform: scale(0.5); opacity: 0; } 50% { transform: scale(1); opacity: 1; } 100% { transform: scale(0.5); opacity: 0; } }`,
	"zoom-out": `
    .outer, .orbit, .core { animation: zoom-out 3s ease-in-out infinite; }
    @keyframes zoom-out { 0% { transform: scale(1.5); opacity: 0; } 50% { transform: scale(1); opacity: 1; } 100% { transform: scale(1.5); opacity: 0; } }`,
	"zoom-pulse": `
    .core { animation: zoom-pulse 2s ease-in-out infinite; }
    @keyframes zoom-pulse { 0%, 100% { transform: scale(1); } 50% { transform: scale(1.5); } }`,

	// SLIDE
	"slide-in": `
    .core { animation: slide-in 3s ease-in-out infinite; }
    @keyframes slide-in { 0% { transform: translateX(-10px); opacity: 0; } 50% { transform: translateX(0); opacity: 1; } 100% { transform: translateX(10px); opacity: 0; } }`,
	"slide-loop": `
    .core { animation: slide-loop 2s linear infinite; }
    @keyframes slide-loop { 0% { transform: translateX(-5px); } 50% { transform: translateX(5px); } 100% { transform: translateX(-5px); } }`,

	// TYPEWRITER
	"typewriter-blink": `
    .core { animation: typewriter-blink 1s steps(1) infinite; }
    @keyframes typewriter-blink { 0%, 50% { opacity: 1; } 51%, 100% { opacity: 0; } }`,
}

// AnimationNames returns all animation names in display order.
func AnimationNames() []string {
	return []string{
		"static",
		"zen", "breathe", "levitate", "glimmer",
		"spin", "spin-fast", "smooth-spin", "orbit-chase", "compass", "gyro", "satellite", "eclipse",
		"pulse", "heartbeat", "pulse-ring", "strobe", "nova", "elastic",
		"flip", "orbit-tilt", "vortex", "harmony", "sync", "sway",
		"radar", "radar-sweep", "signal", "glow", "aurora", "nebula", "corona", "ripple-core",
		"morph", "morph-blob", "morph-crystal",
		"bounce", "bounce-drop", "trampoline",
		"shake", "jitter", "earthquake",
		"swing", "pendulum", "wave-swing",
		"zoom-in", "zoom-out", "zoom-pulse",
		"slide-in", "slide-loop",
		"typewriter-blink",
	}
}

// ValidAnimation checks if an animation name exists.
func ValidAnimation(name string) bool {
	_, ok := Animations[name]
	return ok
}
