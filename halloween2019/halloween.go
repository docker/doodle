package main

import (
	"math"
	"math/rand"
	"strings"
	"time"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
	MoonPhase "github.com/pdevine/goMoonPhase"
)

var random *rand.Rand
var allSprites sprite.SpriteGroup
var Width int
var Height int

const cloud_c0_timer = 22
const cloud_c1_timer = 19


type Title struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

type Whale struct {
	sprite.BaseSprite
	Facing      int
	MinX        int
	MaxX        int
	Timer       int
	TimeOut     int
	TargetX     int
}

type Cloud struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
	VX      int
}

type Moon struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
	PosX    float64
	PosY    float64
}

type Flame struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

type Ghost struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

func NewTitle() *Title {
	t := &Title{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y: 6},
		TimeOut: 5,
	}

	title_list := strings.Split(happy_halloween_c0, "\n")

	for c, _ := range title_list {
		t.AddCostume(sprite.NewCostume(strings.Join(title_list[0:c], "\n"), 'x'))
	}
	return t
}

func (t *Title) Update() {
	t.Timer++
	if t.Timer > t.TimeOut && t.CurrentCostume < len(t.Costumes)-1{
		t.CurrentCostume++
		t.Timer = 0

	}
}

func NewMoon() *Moon {
	m := &Moon{BaseSprite: sprite.BaseSprite{
		Visible: true},
		TimeOut: 60,
	}
	for _, c := range []string{moon_p0, moon_p1, moon_p2, moon_p3, moon_p4, moon_p5, moon_p6, moon_p7, moon_p8} {
		m.AddCostume(sprite.NewCostume(c, 'x'))
	}
	m.CurrentCostume = m.GetPhase()

	return m
}

func (m *Moon) GetPhase() int {
	moon := MoonPhase.New(time.Now())
	i := int(math.Floor(( moon.Phase() + 0.0625 ) * 8))
	return i
}

func (m *Moon) Update() {
	m.Timer++
	if m.Timer > m.TimeOut {
		m.PosX++
		// arc the moon down in a parabola
		m.PosY = 0.00625 * m.PosX * m.PosX
		m.X = int(m.PosX)
		m.Y = int(m.PosY)
		if m.Y >= 25 {
			m.Visible = false
		}
		m.Timer = 0
	}
}

func NewCoffin() *sprite.BaseSprite {
	c := &sprite.BaseSprite{
		Visible: true,
		X: 40,
		Y: 30,
	}
	c.AddCostume(sprite.NewCostume(coffin_c0, 'x'))
	return c
}

func NewCoffinBack() *sprite.BaseSprite {
	c := &sprite.BaseSprite{
		Visible: true,
		X: 40,
		Y: 29,
	}
	c.AddCostume(sprite.NewCostume(coffin_top_c0, 'x'))
	return c
}

func NewTorch(x, y int) *sprite.BaseSprite {
	t := &sprite.BaseSprite{
		Visible: true,
		X: x,
		Y: y,
	}
	t.AddCostume(sprite.NewCostume(torch, 'x'))
	return t
}

func NewCloud(cloudType, posX, posY int) *Cloud {
	c := &Cloud{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y:       posY,
		X:       posX},
		VX: -1,
	}
	if cloudType == 0 {
		c.AddCostume(sprite.NewCostume(cloud_c0, 'x'))
		c.TimeOut = cloud_c0_timer
	} else if cloudType == 1 {
		c.AddCostume(sprite.NewCostume(cloud_c1, 'x'))
		c.TimeOut = cloud_c1_timer
	}

	return c
}

func (c *Cloud) Update() {
	c.Timer++
	if c.Timer > c.TimeOut {
		c.X = c.X + c.VX

		if c.VX < 0 && c.X + c.Width < 0 {
			c.X = 102
		}
		c.Timer = 0
	}
}

func NewWhale() *Whale {
	w := &Whale{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X:       50,
		Y:       30},
		Facing:  -1,
		MinX:    40,
		TimeOut: 20,
	}
	w.AddCostume(sprite.NewCostume(whale_c0, 'x'))
	w.AddCostume(sprite.NewCostume(whale_c0_rev, 'x'))
	return w
}

func (w *Whale) FaceLeft() {
	w.CurrentCostume = 0
}

func (w *Whale) FaceRight() {
	w.CurrentCostume = 1
}

func (w *Whale) Update() {
	w.Timer++

	if w.Timer >= w.TimeOut && w.Y > 21 {
		w.Y--
		w.Timer = 0
	}
}

func NewFlame(x, y int) *Flame {
	f := &Flame{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X: x,
		Y: y},
		TimeOut: 3,
	}

	for _, c := range []string{bigFlame1, bigFlame2, bigFlame3} {
		f.AddCostume(sprite.NewCostume(c, 'x'))
	}
	return f
}

func (f *Flame) Update() {
	f.Timer++
	if f.Timer >= f.TimeOut {
		f.NextCostume()
		f.Timer = 0
	}
}

func NewGhost() *Ghost {
	g := &Ghost{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X: random.Intn(50)+25,
		Y: 30},
		TimeOut: 3,
	}
	g.AddCostume(sprite.NewCostume(ghost_c0, 'x'))
	g.AddCostume(sprite.NewCostume(ghost_c1, 'x'))
	return g
}

func (g *Ghost) Update() {
	var x int
	var y int
	g.Timer++
	if g.Timer >= g.TimeOut {
		x = random.Intn(3)
		x--
		y = random.Intn(3)
		y--
		if x < 0 {
			g.CurrentCostume = 0
		} else if x > 0 {
			g.CurrentCostume = 1
		}
		g.X += x
		g.Y += y
		g.Timer = 0
	}
}


func main() {
	// XXX - Wait a bit until the terminal is properly initialized
	time.Sleep(1000 * time.Millisecond)

	err := tm.Init()
	if err != nil {
		panic(err)
	}
	defer tm.Close()

	Width, Height = tm.Size()

	random = rand.New(rand.NewSource(time.Now().UnixNano()))

	event_queue := make(chan tm.Event)
	go func() {
		for {
			event_queue <- tm.PollEvent()
		}
	}()

	txt := "Press 'ESC' to quit. 'Up' for more ghosts."
	c := sprite.NewCostume(txt, '~')
	text := sprite.NewBaseSprite(Width/2-len(txt)/2, Height-2, c)

	w := NewWhale()
	m := NewMoon()
	cl0 := NewCloud(0, 102, 2)
	cl1 := NewCloud(1, 71, 2)
	cl2 := NewCloud(1, 23, 3)
	t := NewTitle()

	f0 := NewFlame(15, 20)
	t0 := NewTorch(13, 25)
	f1 := NewFlame(110, 20)
	t1 := NewTorch(108, 25)

	coffin := NewCoffin()
	coffinB := NewCoffinBack()

	for _, spr := range []sprite.Sprite{m, t, cl0, cl1, cl2, coffinB, w, f0, t0, f1, t1, coffin, text} {
		allSprites.Sprites = append(allSprites.Sprites, spr)
	}

mainloop:
	for {
		tm.Clear(tm.ColorDefault, tm.ColorDefault)

		select {
		case ev := <-event_queue:
			if ev.Type == tm.EventKey {
				if ev.Key == tm.KeyCtrlC || ev.Key == tm.KeyEsc || ev.Ch == 'q' {
					break mainloop
				} else if ev.Key == tm.KeyArrowLeft {
					w.FaceLeft()
				} else if ev.Key == tm.KeyArrowRight {
					w.FaceRight()
				} else if ev.Key == tm.KeyArrowUp {
					g := NewGhost()
					allSprites.Sprites = append(allSprites.Sprites, g)
				}
			} else if ev.Type == tm.EventResize {
				Width = ev.Width
				Height = ev.Height
			}
		default:
			allSprites.Update()
			allSprites.Render()
			time.Sleep(50 * time.Millisecond)
		}
	}
}
