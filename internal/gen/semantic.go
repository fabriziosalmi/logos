package gen

import (
	"strings"
)

// ResolvedParams holds the API parameters resolved from a natural language prompt.
type ResolvedParams struct {
	Shape     string
	Color     string
	Color2    string
	Animation string
	Scene     string
	Theme     string
	Texture   string
	Variant   string
}

// keywords map categories of words to parameter values.
// Each keyword has a weight. Highest total score per category wins.

type scored struct {
	value string
	score int
}

var shapeKeywords = map[string]scored{
	// shield
	"security": {"shield", 3}, "secure": {"shield", 3}, "protect": {"shield", 2}, "defense": {"shield", 2}, "firewall": {"shield", 3}, "guard": {"shield", 2}, "safe": {"shield", 2}, "vpn": {"shield", 3}, "antivirus": {"shield", 2},
	// hexagon
	"tech": {"hexagon", 2}, "engineering": {"hexagon", 3}, "code": {"hexagon", 2}, "develop": {"hexagon", 2}, "api": {"hexagon", 2}, "sdk": {"hexagon", 2}, "platform": {"hexagon", 2}, "infrastructure": {"hexagon", 2},
	// diamond
	"premium": {"diamond", 3}, "luxury": {"diamond", 3}, "elegant": {"diamond", 2}, "quality": {"diamond", 2}, "saas": {"diamond", 2}, "pro": {"diamond", 2}, "vip": {"diamond", 2},
	// bolt
	"fast": {"bolt", 3}, "speed": {"bolt", 3}, "lightning": {"bolt", 3}, "quick": {"bolt", 2}, "cdn": {"bolt", 3}, "edge": {"bolt", 2}, "performance": {"bolt", 2}, "turbo": {"bolt", 3}, "rocket": {"bolt", 2},
	// cube
	"infra": {"cube", 3}, "block": {"cube", 2}, "container": {"cube", 3}, "docker": {"cube", 3}, "kubernetes": {"cube", 3}, "storage": {"cube", 2}, "database": {"cube", 2}, "server": {"cube", 2},
	// wave
	"data": {"wave", 2}, "stream": {"wave", 3}, "flow": {"wave", 3}, "audio": {"wave", 3}, "music": {"wave", 3}, "signal": {"wave", 2}, "telemetry": {"wave", 3}, "log": {"wave", 2}, "analytics": {"wave", 2},
	// gear
	"ops": {"gear", 3}, "settings": {"gear", 3}, "config": {"gear", 2}, "devops": {"gear", 3}, "ci": {"gear", 2}, "cd": {"gear", 2}, "pipeline": {"gear", 2}, "automation": {"gear", 3}, "build": {"gear", 2},
	// eye
	"monitor": {"eye", 3}, "watch": {"eye", 3}, "observe": {"eye", 3}, "surveillance": {"eye", 3}, "vision": {"eye", 3}, "insight": {"eye", 2}, "detect": {"eye", 2}, "scan": {"eye", 2},
	// leaf
	"eco": {"leaf", 3}, "green": {"leaf", 2}, "organic": {"leaf", 3}, "nature": {"leaf", 3}, "sustainable": {"leaf", 3}, "plant": {"leaf", 3}, "garden": {"leaf", 2}, "earth": {"leaf", 2},
	// star
	"star": {"star", 3}, "favorite": {"star", 3}, "rating": {"star", 3}, "featured": {"star", 2}, "best": {"star", 2}, "top": {"star", 2}, "award": {"star", 2},
	// heart
	"love": {"heart", 3}, "health": {"heart", 3}, "heart": {"heart", 3}, "care": {"heart", 2}, "medical": {"heart", 2}, "wellness": {"heart", 2}, "life": {"heart", 2},
	// cloud
	"cloud": {"cloud", 3}, "hosting": {"cloud", 2}, "aws": {"cloud", 3}, "azure": {"cloud", 3}, "gcp": {"cloud", 3},
	// flame
	"fire": {"flame", 3}, "hot": {"flame", 3}, "trending": {"flame", 3}, "flame": {"flame", 3}, "popular": {"flame", 2}, "viral": {"flame", 2},
	// lock
	"lock": {"lock", 3}, "auth": {"lock", 3}, "password": {"lock", 3}, "encrypt": {"lock", 3}, "key": {"lock", 2}, "secret": {"lock", 2}, "vault": {"lock", 3}, "tls": {"lock", 3}, "ssl": {"lock", 3}, "certificate": {"lock", 3},
	// crown
	"king": {"crown", 3}, "crown": {"crown", 3}, "royal": {"crown", 3}, "admin": {"crown", 2}, "master": {"crown", 2},
	// target
	"target": {"target", 3}, "focus": {"target", 2}, "precision": {"target", 2}, "aim": {"target", 2}, "goal": {"target", 2},
	// arrow
	"arrow": {"arrow", 3}, "direction": {"arrow", 2}, "navigate": {"arrow", 2}, "forward": {"arrow", 2}, "next": {"arrow", 2}, "go": {"arrow", 2},
	// moon
	"moon": {"moon", 3}, "night": {"moon", 3}, "dark": {"moon", 2}, "sleep": {"moon", 2}, "midnight": {"moon", 3},
	// sun
	"sun": {"sun", 3}, "day": {"sun", 2}, "bright": {"sun", 2}, "morning": {"sun", 2}, "light": {"sun", 2}, "solar": {"sun", 3},
	// triangle
	"play": {"triangle", 3}, "start": {"triangle", 2}, "begin": {"triangle", 2}, "delta": {"triangle", 3}, "change": {"triangle", 2},
	// circle
	"minimal": {"circle", 2}, "simple": {"circle", 2}, "clean": {"circle", 2}, "basic": {"circle", 2},
	// atom
	"atom": {"atom", 3}, "science": {"atom", 3}, "physics": {"atom", 3}, "nuclear": {"atom", 3}, "quantum": {"atom", 3}, "orbital": {"atom", 3},
}

var colorKeywords = map[string]scored{
	"red": {"rose", 3}, "blue": {"blue", 3}, "green": {"green", 3}, "purple": {"purple", 3}, "orange": {"orange", 3}, "yellow": {"amber", 3}, "cyan": {"cyan", 3}, "pink": {"rose", 2}, "indigo": {"indigo", 3},
	"neon": {"neon", 3}, "cyber": {"cyber", 3}, "matrix": {"matrix", 3}, "gold": {"gold", 3}, "emerald": {"emerald", 3}, "ruby": {"ruby", 3}, "sapphire": {"sapphire", 3},
	"dark": {"black", 1}, "bright": {"neon", 2}, "warm": {"amber", 2}, "cool": {"cyan", 2}, "cold": {"blue", 2}, "hot": {"magma", 2}, "ice": {"cyan", 2},
	"ocean": {"ocean", 3}, "sunset": {"sunset", 3}, "forest": {"emerald", 2}, "fire": {"magma", 2}, "electric": {"laser", 3}, "void": {"void", 3},
	"hacker": {"neon", 3}, "corporate": {"blue", 2}, "creative": {"purple", 2}, "organic": {"green", 2}, "luxury": {"gold", 3}, "minimal": {"white", 2},
	"blood": {"ruby", 3}, "lava": {"magma", 3}, "sky": {"cyan", 2}, "sea": {"ocean", 2}, "mint": {"mint", 3}, "lavender": {"lavender", 3}, "peach": {"peach", 3},
}

var animationKeywords = map[string]scored{
	"spin": {"spin", 3}, "rotate": {"spin", 2}, "fast": {"spin-fast", 2}, "slow": {"zen", 2}, "calm": {"zen", 3}, "peaceful": {"zen", 3}, "breathe": {"breathe", 3}, "relax": {"breathe", 2},
	"pulse": {"pulse", 3}, "heartbeat": {"heartbeat", 3}, "beat": {"heartbeat", 2}, "alive": {"pulse", 2},
	"glow": {"glow", 3}, "shine": {"glow", 2}, "aurora": {"aurora", 3}, "nebula": {"nebula", 3},
	"vortex": {"vortex", 3}, "whirl": {"vortex", 2}, "tornado": {"vortex", 2},
	"bounce": {"bounce", 3}, "jump": {"bounce", 2}, "spring": {"elastic", 2},
	"shake": {"shake", 3}, "vibrate": {"jitter", 3}, "earthquake": {"earthquake", 3},
	"orbit": {"orbit-chase", 3}, "satellite": {"satellite", 3}, "radar": {"radar", 3},
	"float": {"levitate", 3}, "hover": {"levitate", 2}, "levitate": {"levitate", 3},
	"flash": {"strobe", 3}, "blink": {"strobe", 2}, "nova": {"nova", 3}, "explode": {"nova", 2},
	"slide": {"slide-loop", 2}, "wave": {"wave-swing", 2}, "swing": {"swing", 3}, "pendulum": {"pendulum", 3},
	"flip": {"flip", 3}, "morph": {"morph", 3}, "transform": {"morph", 2},
	"still": {"static", 3}, "static": {"static", 3}, "frozen": {"static", 2}, "stop": {"static", 2},
	"glimmer": {"glimmer", 3}, "twinkle": {"glimmer", 2}, "sparkle": {"glimmer", 2},
	"zoom": {"zoom-pulse", 2}, "grow": {"zoom-in", 2}, "shrink": {"zoom-out", 2},
	"gyro": {"gyro", 3}, "compass": {"compass", 3},
	"spooky": {"vortex", 2}, "creepy": {"eclipse", 2}, "mysterious": {"nebula", 2},
	"aggressive": {"earthquake", 2}, "angry": {"shake", 2}, "intense": {"spin-fast", 2},
	"smooth": {"smooth-spin", 3}, "elegant": {"zen", 2}, "gentle": {"breathe", 2},
}

var sceneKeywords = map[string]scored{
	"spotlight": {"spotlight", 3}, "focus": {"spotlight", 2}, "highlight": {"spotlight", 2},
	"grid": {"grid", 3}, "cyberpunk": {"grid", 3}, "tron": {"grid", 3}, "matrix": {"grid", 2},
	"dots": {"dots", 3}, "dotted": {"dots", 2}, "polka": {"dots", 2},
	"circuit": {"circuit", 3}, "electronic": {"circuit", 2}, "board": {"circuit", 2}, "pcb": {"circuit", 3},
	"hexgrid": {"hexgrid", 3}, "honeycomb": {"hexgrid", 3}, "hex": {"hexgrid", 2},
	"gradient": {"gradient", 3}, "fade": {"gradient", 2}, "blend": {"gradient", 2},
	"noise": {"noise-bg", 2}, "grain": {"noise-bg", 2},
	"vignette": {"vignette", 3}, "cinematic": {"vignette", 2}, "movie": {"vignette", 2},
	"diagonal": {"diagonal", 2}, "stripe": {"diagonal", 2}, "line": {"diagonal", 2},
	"radial": {"radial", 2},
	"clean": {"pure", 2}, "transparent": {"pure", 3}, "none": {"pure", 2},
}

var themeKeywords = map[string]scored{
	"dark": {"dark", 3}, "black": {"dark", 2}, "night": {"dark", 2}, "shadow": {"dark", 2},
	"light": {"light", 3}, "white": {"light", 2}, "bright": {"light", 2}, "day": {"light", 2},
	"monokai": {"monokai", 3}, "dracula": {"dracula", 3}, "nord": {"nord", 3}, "catppuccin": {"catppuccin", 3},
	"solarized": {"solarized-dark", 3}, "gruvbox": {"gruvbox", 3}, "tokyo": {"tokyo-night", 3},
	"github": {"github-dark", 2}, "rose-pine": {"rose-pine", 3}, "kanagawa": {"kanagawa", 3},
	"glass": {"glass", 3}, "frosted": {"glass", 2}, "transparent": {"glass", 2},
	"solid": {"solid", 3}, "filled": {"solid", 2},
	"hacker": {"monokai", 2}, "coder": {"dracula", 2}, "dev": {"one-dark", 2}, "terminal": {"tokyo-night", 2},
}

var textureKeywords = map[string]scored{
	"grain": {"grain", 3}, "film": {"grain", 3}, "analog": {"grain", 3}, "vintage": {"grain", 2}, "retro": {"grain", 2}, "paper": {"grain", 2},
	"glass": {"glass", 3}, "frosted": {"glass", 3}, "blur": {"glass", 2}, "ice": {"glass", 2},
	"glitch": {"glitch", 3}, "broken": {"glitch", 2}, "corrupt": {"glitch", 2}, "hack": {"glitch", 2}, "cyber": {"glitch", 2},
	"shadow": {"shadow", 3}, "glow": {"shadow", 2}, "soft": {"shadow", 2},
	"neon": {"neon", 3}, "electric": {"neon", 2}, "bright": {"neon", 2}, "laser": {"neon", 2},
	"noise": {"noise", 3}, "digital": {"noise", 2}, "static": {"noise", 2},
}

var variantKeywords = map[string]scored{
	"outline": {"outline", 3}, "wireframe": {"outline", 3}, "stroke": {"outline", 2}, "line": {"outline", 2},
	"solid": {"solid", 3}, "filled": {"solid", 3}, "full": {"solid", 2},
	"ghost": {"ghost", 3}, "faded": {"ghost", 3}, "dim": {"ghost", 2}, "disabled": {"ghost", 2}, "inactive": {"ghost", 3},
	"glow": {"glow", 3}, "neon": {"glow", 2}, "bright": {"glow", 2},
	"badge": {"badge", 3}, "icon": {"badge", 2}, "app": {"badge", 2}, "ios": {"badge", 3},
	"sketch": {"sketch", 3}, "hand": {"sketch", 2}, "drawn": {"sketch", 2}, "pencil": {"sketch", 2},
	"retro": {"retro", 3}, "80s": {"retro", 3}, "synthwave": {"retro", 3}, "vintage": {"retro", 2},
	"sticker": {"sticker", 3}, "cutout": {"sticker", 2}, "die": {"sticker", 2},
	"stamp": {"stamp", 3}, "bold": {"stamp", 2}, "thick": {"stamp", 2},
	"minimal": {"minimal", 3}, "tiny": {"minimal", 2}, "micro": {"minimal", 2}, "dot": {"minimal", 2},
	"inverted": {"inverted", 3}, "reverse": {"inverted", 2}, "negative": {"inverted", 2},
	"duotone": {"duotone", 3}, "two-tone": {"duotone", 3}, "layer": {"duotone", 2},
	"emboss": {"emboss", 3}, "3d": {"emboss", 2}, "raised": {"emboss", 2}, "relief": {"emboss", 2},
	"dotted": {"dotted", 3}, "dashed": {"dotted", 2}, "blueprint": {"dotted", 3},
	"neon-outline": {"neon-outline", 3}, "cyberpunk": {"neon-outline", 2},
	"xray": {"xray", 3}, "scan": {"xray", 2}, "contrast": {"xray", 2},
}

// ResolvePrompt takes a natural language prompt and maps it to API parameters.
// Returns the best-scoring parameter for each category.
func ResolvePrompt(prompt string) ResolvedParams {
	words := tokenize(prompt)

	return ResolvedParams{
		Shape:     resolve(words, shapeKeywords, "circle"),
		Color:     resolve(words, colorKeywords, "blue"),
		Animation: resolve(words, animationKeywords, "breathe"),
		Scene:     resolve(words, sceneKeywords, "spotlight"),
		Theme:     resolve(words, themeKeywords, "dark"),
		Texture:   resolve(words, textureKeywords, ""),
		Variant:   resolve(words, variantKeywords, ""),
	}
}

func tokenize(prompt string) []string {
	prompt = strings.ToLower(prompt)
	// Split on spaces, hyphens, underscores, commas
	replacer := strings.NewReplacer("-", " ", "_", " ", ",", " ", ".", " ", "/", " ")
	prompt = replacer.Replace(prompt)
	words := strings.Fields(prompt)
	return words
}

func resolve(words []string, keywords map[string]scored, fallback string) string {
	scores := make(map[string]int)

	for _, word := range words {
		if s, ok := keywords[word]; ok {
			scores[s.value] += s.score
		}
		// Also try partial matches (e.g., "spinning" matches "spin")
		for keyword, s := range keywords {
			if len(word) > 3 && strings.HasPrefix(word, keyword) {
				scores[s.value] += s.score - 1
			}
		}
	}

	if len(scores) == 0 {
		return fallback
	}

	best := fallback
	bestScore := 0
	for value, score := range scores {
		if score > bestScore {
			best = value
			bestScore = score
		}
	}
	return best
}
