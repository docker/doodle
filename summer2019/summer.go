package main

import (
	"math/rand"
	"time"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
)

var random *rand.Rand
var allSprites sprite.SpriteGroup
var Width int
var Height int

const tree_c0 = `
xxxxxxxxxxxxxxxxxxxxxxxxxxx###
xxxxxxxxxxxxxxxxxxxxxxx#########
xxxxxxxxxxxx###xxxx########xxxxxxx###
xxxxxxxx########/#####\#####xx##########
xxxx#########/##########--##################
xx####xxxxxxxxx###################xxxxxxxx#####
xx#xxxxxxxxxx####xxxxxx##########/@@xxxxxxxxxx###
#xxxxxxxxxx####xxxxxxxxx##  .#\#####xxxxxxxxxxxx##
xxxxxxxxxx###xxxxxxxxxxx$$$$xxx.x####xxxxxxxxxxxx#
xxxxxxxxx##xxxxxxxxxxxxx$$$$xxxxxxx###
xxxxxxxxx#xxxxxxxxxxxxxx$$$$xxxxxxx##
xxxxxxxxxxxxxxxxxxxxxxxx$$$$xxxxxxx##
xxxxxxxxxxxxxxxxxxxxxxxx$$$$xxxxxxxx#
xxxxxxxxxxxxxxxxxxxxxxxx$$$$
xxxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
',',.~.,'~.,'',',.~  $$$$$
~,', ',',.~,',.~ ~,',$$$$$~,',~,',
 ~,',',.~~,',',',.~ $$$$$$$  ,',.~,'
     ',',.~,.~.',','$$$$$$$,',.~,'
',',',',.~,',.~    $$$$$$$$$,',.~,'
~,',,.~,.~,.~',',' $$$$$$$$$~~,',
  ~,',~,',  ~,', ~,', $$$ ,,.~,.~,.
~,.  ~~.. ~, ,~ ~~'' .'  ~,. ~,.. ~,
  , , ~.~ ' ~~ '',' ~~.  ',.~ ,~~ . 
,.~,',.~ ~,'   ',',.~,.~.',~,',  ~,', ~,
`

const tree_c1 = `
xxxxxxxxxxxxxxxxxxxxxxxxxxx#####
xxxxxxxxxxxxxxxxxxxxxxx#######
xxxxxxxxxxxx###xxxx########xxxxxxx###
xxxxxxxxx########/#####\#####xx##########
xxxxx#########/##########--##################
xxx####xxxxxxxx###################xxxxxxx#####
xx##xxxxxxxxx####xxxxxx##########/@@xxxxxxxxx###
xx#xxxxxxxxx####xxxxxxxx##  .#\#####xxxxxxxxxxx##
xxxxxxxxxxx###xxxxxxxxxx$$$$xxxx####xxxxxxxxxxx#
xxxxxxxxxx##xxxxxxxxxxxx$$$$xxxxxx###
xxxxxxxxxx#xxxxxxxxxxxxx$$$$xxxxxx##
xxxxxxxxxxxxxxxxxxxxxxxx$$$$xxxxxx##
xxxxxxxxxxxxxxxxxxxxxxxx$$$$xxxxxxx#
xxxxxxxxxxxxxxxxxxxxxxxx$$$$
xxxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
xxxxxxxxxxxxxxxxxxxxx$$$$$
',',.~.,'~.,'',',.~  $$$$$
~,', ',',.~,',.~ ~,',$$$$$~,',~,',
 ~,',',.~~,',',',.~ $$$$$$$  ,',.~,'
     ',',.~,.~.',','$$$$$$$,',.~,'
',',',',.~,',.~    $$$$$$$$$,',.~,'
~,',,.~,.~,.~',',' $$$$$$$$$~~,',
  ~,',~,',  ~,', ~,', $$$ ,,.~,.~,.
~,.  ~~.. ~, ,~ ~~'' .'  ~,. ~,.. ~,
  , , ~.~ ' ~~ '',' ~~.  ',.~ ,~~ . 
,.~,',.~ ~,'   ',',.~,.~.',~,',  ~,', ~,
`


const sun_c0 = `
      ;   :   ;
   .   \_,!,_/   ,
    '.,'     '.,'
     /         \
~ -- :         : -- ~
     \         /
    ,''._   _.''.
   '   / '!' \   '
      ;   :   ;
`

const cloud_c0_timer = 22
const cloud_c1_timer = 19

const cloud_c0 = `
xxxxxxxxx_ _ __.
xxxxxxx('       ).
xxxxxx(          ).
xxxxx_(          '''.
x.=('(           .   )
((        (..____.:'-'
'(        )_),
xx' _______.:'   )
xxxxxxxxxx-----'`

const cloud_c1 = `
xxxx.--
x.+(   )
x(   .  )
(   (   ))
x'- __.'`

const whale_c0 = `xxxxxxxxxxxxxxx##xxxxxxxx.xxxxx
xxxxxxxxx##x##x##xxxxxxx==xxxxx
xxxxxx##x##x##x##xxxxxx===xxxxx
xx/""""""""""""""""\___/x===xxx
x{                      /xx===x
xx\______ o          __/xxxxxxx
xxxx\    \        __/xxxxxxxxxx
xxxxx\____\______/xxxxxxxxxxxxx`

const whale_c1 = `xxxxxxxxxxxxxxx##xxxxxxxx.xxxxx
xxxxxxxxx##x##x##xxxxxxx==xxxxx
xxxxxx##x##x##x##xxxxxx===xxxxx
xx/""""""""""""""""\___/x===xxx
x{                      /xx===x
xx\______ o          __/xxxxxxx
xxxx\    \        __/xxxxxxxxxx`

const whale_c1_rev = `xxxxx.xxxxxxxx##xxxxxxxxxxxxxxx
xxxxx==xxxxxxx##x##x##xxxxxxxxx
xxxxx===xxxxxx##x##x##x##xxxxxx
xxx===x\___/""""""""""""""""\xx
x===xx\                      }x
xxxxxxx\__          o ______/xx
xxxxxxxxxx\__        /    /xxxx`

const font_crawford_d = `
x___
|   \
|    \
|  D  |
|     |
|     |
|_____|`

const font_crawford_e = `
xxx___
xx/  _]
x/  [_
|    _]
|   [_
|     |
|_____|`

const font_crawford_l = `
x_
| |
| |
| |___
|     |
|     |
|_____|`

const font_crawford_m = `
x___ ___
|   |   |
| _   _ |
|  \_/  |
|   |   |
|   |   |
|___|___|`

const font_crawford_n = `
x____
|    \
|  _  |
|  |  |
|  |  |
|  |  |
|__|__|`

const font_crawford_r = `
x____
|    \
|  D  )
|    /
|    \
|  .  \
|__|\_|`

const font_crawford_s = `
xx_____
x/ ___/
(   \_
x\__  |
x/  \ |
x\    |
xx\___|`

const font_crawford_u = `
x__ __
|  |  |
|  |  |
|  |  |
|  :  |
|     |
x\__,_|`


type Letter struct {
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

type Tree struct {
	sprite.BaseSprite
	TimeOut int
	Timer   int
}

type Ocean struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

type Sun struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
	PosX    float64
	PosY    float64
}

type Cloud struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
	VX      int
}

func NewLetter(s string, x, y, timeOut int) *Letter {
	l := &Letter{BaseSprite: sprite.BaseSprite{
		Visible: false,
		X: x,
		Y: y},
		TimeOut: timeOut,
	}
	l.AddCostume(sprite.NewCostume(s, 'x'))
	return l
}

func (l *Letter) Update() {
	l.Timer++
	if l.Timer >= l.TimeOut {
		l.Visible = true
	}
}

func addTitle() {
	charMap := map[rune]string{
		'd': font_crawford_d,
		'e': font_crawford_e,
		'l': font_crawford_l,
		'm': font_crawford_m,
		'n': font_crawford_n,
		'r': font_crawford_r,
		's': font_crawford_s,
		'u': font_crawford_u,
	}

	var char_width = 7 // hard-coding this, despite different widths
	var fade_offset = 3
	var y_offset = 10

	for cnt, c := range "endless" {
		l := NewLetter(charMap[c], cnt*char_width+10, y_offset, cnt*fade_offset+10)
		allSprites.Sprites = append(allSprites.Sprites, l)
	}

	for cnt, c := range "summer" {
		l := NewLetter(charMap[c], cnt*char_width+65, y_offset, cnt*fade_offset+10)
		allSprites.Sprites = append(allSprites.Sprites, l)
	}
}

func NewTree() *Tree {
	t := &Tree{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y:       2},
		TimeOut: 15,
	}
	t.AddCostume(sprite.NewCostume(tree_c0, 'x'))
	t.AddCostume(sprite.NewCostume(tree_c1, 'x'))
	return t
}

func (t *Tree) Update() {
	t.Timer++

	if t.Timer >= t.TimeOut {
		t.Timer = 0
		t.TimeOut = random.Intn(20)+30
		t.CurrentCostume++
		if t.CurrentCostume >= len(t.Costumes) {
			t.CurrentCostume = 0
		}
	}
}

func NewOcean() *Ocean {
	o := &Ocean{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X:       36,
		Y:       26},
		TimeOut: 20,
	}

	// create some random waves
	for cnt := 0; cnt < 3; cnt++ {
		o.AddOceanCostume()
	}
	
	return o
}

func (o *Ocean) AddOceanCostume() {
	var s string
	for l := 0; l < 9; l++ {
		w := make([]rune, 80)
		for cnt := 0; cnt < len(w); cnt++ {
			if random.Intn(2) == 0 {
				w[cnt] = '~'
			}
		}
		s += string(w) + "\n"
	}
	o.AddCostume(sprite.NewCostume(s, 'x'))
}

func (o *Ocean) Update() {
	o.Timer++

	if o.Timer >= o.TimeOut {
		o.Timer = 0
		o.CurrentCostume++
		if o.CurrentCostume >= len(o.Costumes) {
			o.CurrentCostume = 0
		}
	}
}

func NewSun() *Sun {
	s := &Sun{BaseSprite: sprite.BaseSprite{
		Visible: true},
		TimeOut: 30,
	}
	s.AddCostume(sprite.NewCostume(sun_c0, 'x'))
	return s
}

func (s *Sun) Update() {
	s.Timer++
	if s.Timer > s.TimeOut {
		s.PosX++
		// arc the sun down in a parabola
		s.PosY = 0.00625 * s.PosX * s.PosX
		s.X = int(s.PosX)
		s.Y = int(s.PosY)
		if s.Y >= 25 {
			s.Visible = false
		}
		s.Timer = 0
	}
}

func NewCloud(cloudType, posX, posY int) *Cloud {
	c := &Cloud{BaseSprite: sprite.BaseSprite{
		Visible:        true,
		Y:              posY,
		X:              posX},
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
		X:       100,
		Y:       22},
		Facing:  -1,
		MinX:    40,
		TimeOut: 4,
	}
	w.AddCostume(sprite.NewCostume(whale_c1, 'x'))
	w.AddCostume(sprite.NewCostume(whale_c1_rev, 'x'))
	return w
}

func (w *Whale) MoveRight() {
	w.TargetX += 10
	w.Facing = 1
	w.CurrentCostume = 1
}

func (w *Whale) Update() {
	w.Timer++

	if w.Timer >= w.TimeOut {
		if w.X > 113 {
			w.Visible = false
		} else {
			w.Visible = true
		}

		w.Timer = 0
		w.X += w.Facing
		if w.X <= w.MinX {
			w.X = w.MinX
		}
		if w.Facing == 1 {
			if w.TargetX <= 0 {
				w.Facing = -1
				w.CurrentCostume = 0
			} else {
				w.TargetX--
			}
		}
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

	txt := "Press 'ESC' to quit."
	c := sprite.NewCostume(txt, '~')
	text := sprite.NewBaseSprite(Width/2-len(txt)/2, Height-2, c)

	w := NewWhale()
	cl0 := NewCloud(0, 102, 2)
	cl1 := NewCloud(1, 71, 2)
	cl2 := NewCloud(1, 23, 3)
	s := NewSun()
	o := NewOcean()
	t := NewTree()

	allSprites.Sprites = append(allSprites.Sprites, s)
	addTitle()

	for _, spr := range []sprite.Sprite{cl0, cl1, cl2, o, t, w, text} {
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
				} else if ev.Key == tm.KeyArrowRight {
					w.MoveRight()
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
