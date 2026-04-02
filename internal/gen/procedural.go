package gen

import (
	"crypto/sha256"
	"fmt"
	"math"
	"strings"
)

// GeneratePath creates a unique SVG path from a text prompt.
// Deterministic: same prompt = same geometry = cacheable.
// Coordinates on a 256x256 grid for infinite upscale via viewBox.
func GeneratePath(prompt string) string {
	seed := hashToSeeds(prompt)

	// Choose generator based on seed
	genType := seed.primary % 7
	switch genType {
	case 0:
		return generateBlob(seed)
	case 1:
		return generateStarBurst(seed)
	case 2:
		return generateSpiral(seed)
	case 3:
		return generateLissajous(seed)
	case 4:
		return generateWaveRing(seed)
	case 5:
		return generateSymmetricPoly(seed)
	case 6:
		return generateCompound(seed)
	default:
		return generateBlob(seed)
	}
}

// seeds holds deterministic parameters extracted from prompt hash.
type seeds struct {
	primary    int     // main selector (0-255)
	points     int     // number of vertices/segments (5-20)
	symmetry   int     // rotational symmetry order (3-12)
	amplitude  float64 // distortion strength (0.1-0.5)
	innerRatio float64 // inner/outer radius ratio (0.3-0.8)
	rotation   float64 // base rotation in radians
	complexity float64 // curve detail (0.5-2.0)
	cx, cy     float64 // center offset
}

func hashToSeeds(prompt string) seeds {
	h := sha256.Sum256([]byte(strings.ToLower(strings.TrimSpace(prompt))))
	return seeds{
		primary:    int(h[0]),
		points:     int(h[1])%16 + 5,                           // 5-20
		symmetry:   int(h[2])%10 + 3,                           // 3-12
		amplitude:  0.1 + float64(h[3])/255.0*0.4,              // 0.1-0.5
		innerRatio: 0.3 + float64(h[4])/255.0*0.5,              // 0.3-0.8
		rotation:   float64(h[5]) / 255.0 * 2 * math.Pi,        // 0-2π
		complexity: 0.5 + float64(h[6])/255.0*1.5,              // 0.5-2.0
		cx:         128 + (float64(h[7])/255.0-0.5)*20,         // 118-138
		cy:         128 + (float64(h[8])/255.0-0.5)*20,         // 118-138
	}
}

// generateBlob creates an organic blob shape using noise-distorted circle.
func generateBlob(s seeds) string {
	var b strings.Builder
	n := s.points
	r := 90.0
	cx, cy := s.cx, s.cy

	// Generate distorted circle points
	pts := make([][2]float64, n)
	for i := 0; i < n; i++ {
		angle := s.rotation + float64(i)/float64(n)*2*math.Pi
		// Pseudo-noise distortion using overlapping sines
		noise := math.Sin(float64(i)*3.7+s.amplitude*10) * s.amplitude * r
		noise += math.Sin(float64(i)*7.3+s.complexity*5) * s.amplitude * r * 0.5
		dist := r + noise
		pts[i] = [2]float64{
			cx + dist*math.Cos(angle),
			cy + dist*math.Sin(angle),
		}
	}

	// Build smooth cubic bezier path through all points
	b.WriteString(fmt.Sprintf("M %.0f %.0f ", clamp(pts[0][0]), clamp(pts[0][1])))
	for i := 0; i < n; i++ {
		p0 := pts[i]
		p1 := pts[(i+1)%n]
		p2 := pts[(i+2)%n]

		// Control points for smooth curves
		cp1x := p0[0] + (p1[0]-pts[(i-1+n)%n][0])*0.3
		cp1y := p0[1] + (p1[1]-pts[(i-1+n)%n][1])*0.3
		cp2x := p1[0] - (p2[0]-p0[0])*0.3
		cp2y := p1[1] - (p2[1]-p0[1])*0.3

		b.WriteString(fmt.Sprintf("C %.0f %.0f %.0f %.0f %.0f %.0f ",
			clamp(cp1x), clamp(cp1y), clamp(cp2x), clamp(cp2y), clamp(p1[0]), clamp(p1[1])))
	}
	b.WriteString("Z")
	return b.String()
}

// generateStarBurst creates a star with variable-length rays.
func generateStarBurst(s seeds) string {
	var b strings.Builder
	n := s.symmetry * 2 // alternating inner/outer points
	rOuter := 100.0
	rInner := rOuter * s.innerRatio
	cx, cy := s.cx, s.cy

	for i := 0; i < n; i++ {
		angle := s.rotation + float64(i)/float64(n)*2*math.Pi
		r := rOuter
		if i%2 == 1 {
			// Vary inner radius per ray
			variation := math.Sin(float64(i)*s.complexity*2) * s.amplitude * rInner
			r = rInner + variation
		}
		x := clamp(cx + r*math.Cos(angle))
		y := clamp(cy + r*math.Sin(angle))
		if i == 0 {
			b.WriteString(fmt.Sprintf("M %.0f %.0f ", x, y))
		} else {
			b.WriteString(fmt.Sprintf("L %.0f %.0f ", x, y))
		}
	}
	b.WriteString("Z")
	return b.String()
}

// generateSpiral creates an Archimedean spiral.
func generateSpiral(s seeds) string {
	var b strings.Builder
	turns := 2.0 + s.complexity
	steps := s.points * 4
	cx, cy := s.cx, s.cy

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		angle := s.rotation + t*turns*2*math.Pi
		r := t * 90 // expanding radius
		x := clamp(cx + r*math.Cos(angle))
		y := clamp(cy + r*math.Sin(angle))
		if i == 0 {
			b.WriteString(fmt.Sprintf("M %.0f %.0f ", x, y))
		} else {
			b.WriteString(fmt.Sprintf("L %.0f %.0f ", x, y))
		}
	}
	return b.String()
}

// generateLissajous creates a parametric Lissajous curve.
func generateLissajous(s seeds) string {
	var b strings.Builder
	a := float64(s.symmetry)
	bFreq := float64(s.points % 7 + 2)
	delta := s.rotation
	steps := 100
	r := 90.0
	cx, cy := s.cx, s.cy

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps) * 2 * math.Pi
		x := clamp(cx + r*math.Sin(a*t+delta))
		y := clamp(cy + r*math.Sin(bFreq*t))
		if i == 0 {
			b.WriteString(fmt.Sprintf("M %.0f %.0f ", x, y))
		} else {
			b.WriteString(fmt.Sprintf("L %.0f %.0f ", x, y))
		}
	}
	b.WriteString("Z")
	return b.String()
}

// generateWaveRing creates a circle with sinusoidal edge distortion.
func generateWaveRing(s seeds) string {
	var b strings.Builder
	steps := s.points * 3
	baseR := 80.0
	waveFreq := float64(s.symmetry)
	waveAmp := s.amplitude * 30
	cx, cy := s.cx, s.cy

	for i := 0; i <= steps; i++ {
		angle := s.rotation + float64(i)/float64(steps)*2*math.Pi
		r := baseR + math.Sin(angle*waveFreq)*waveAmp
		x := clamp(cx + r*math.Cos(angle))
		y := clamp(cy + r*math.Sin(angle))
		if i == 0 {
			b.WriteString(fmt.Sprintf("M %.0f %.0f ", x, y))
		} else {
			b.WriteString(fmt.Sprintf("L %.0f %.0f ", x, y))
		}
	}
	b.WriteString("Z")
	return b.String()
}

// generateSymmetricPoly creates a polygon with rotational symmetry and notches.
func generateSymmetricPoly(s seeds) string {
	var b strings.Builder
	n := s.symmetry
	r := 95.0
	notchDepth := s.innerRatio * 0.5
	cx, cy := s.cx, s.cy

	pointsPerSide := 3
	for i := 0; i < n; i++ {
		for j := 0; j < pointsPerSide; j++ {
			angle := s.rotation + (float64(i)+float64(j)/float64(pointsPerSide))/float64(n)*2*math.Pi
			dist := r
			if j == 1 {
				dist = r * (1 - notchDepth)
			}
			x := clamp(cx + dist*math.Cos(angle))
			y := clamp(cy + dist*math.Sin(angle))
			if i == 0 && j == 0 {
				b.WriteString(fmt.Sprintf("M %.0f %.0f ", x, y))
			} else {
				b.WriteString(fmt.Sprintf("L %.0f %.0f ", x, y))
			}
		}
	}
	b.WriteString("Z")
	return b.String()
}

// generateCompound overlays two simple shapes.
func generateCompound(s seeds) string {
	// Outer shape: rotated polygon
	outer := generateStarBurst(seeds{
		primary:    s.primary,
		symmetry:   s.symmetry,
		innerRatio: s.innerRatio * 1.2,
		rotation:   s.rotation,
		amplitude:  s.amplitude * 0.5,
		complexity: s.complexity,
		cx:         s.cx,
		cy:         s.cy,
	})

	// Inner shape: smaller wave ring
	s2 := s
	s2.cx = s.cx
	s2.cy = s.cy
	s2.points = s.points / 2
	if s2.points < 5 {
		s2.points = 5
	}
	inner := generateWaveRing(seeds{
		primary:   s.primary,
		points:    s2.points,
		symmetry:  s.symmetry + 2,
		amplitude: s.amplitude * 0.6,
		rotation:  s.rotation + 0.3,
		cx:        s.cx,
		cy:        s.cy,
	})

	return outer + " " + inner
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}
