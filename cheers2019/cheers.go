package main

import (
	"math"
	"math/rand"
	"time"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
)

var random *rand.Rand
var allSprites sprite.SpriteGroup
var bubbles []*Bubble
var Width int
var Height int

const (
	falling = iota
	swimming
)

const cheers_c0 = `
                                                                  ,---,  
                                                               ,`+"`"+`--.' |  
  ,----..    ,---,                                             |   :  :  
 /   /   \ ,--.' |                                             '   '  ;  
|   :     :|  |  :                          __  ,-.            |   |  |  
.   |  ;. /:  :  :                        ,' ,'/ /|  .--.--.   '   :  ;  
.   ; /--`+"`"+` :  |  |,--.   ,---.     ,---.  '  | |' | /  /    '  |   |  '  
;   | ;    |  :  '   |  /     \   /     \ |  |   ,'|  :  /`+"`"+`./  '   :  |  
|   : |    |  |   /' : /    /  | /    /  |'  :  /  |  :  ;_    ;   |  ;  
.   | '___ '  :  | | |.    ' / |.    ' / ||  | '    \  \    `+"`"+`. `+"`"+`---'. |  
'   ; : .'||  |  ' | :'   ;   /|'   ;   /|;  : |     `+"`"+`----.   \ `+"`"+`--..`+"`"+`;  
'   | '/  :|  :  :_:,''   |  / |'   |  / ||  , ;    /  /`+"`"+`--'  /.--,_     
|   :    / |  | ,'    |   :    ||   :    | ---'    '--'.     / |    |`+"`"+`.  
 \   \ .'  `+"`"+`--''       \   \  /  \   \  /            `+"`"+`--'---'  `+"`"+`-- -`+"`"+`, ; 
  `+"`"+`---`+"`"+`                 `+"`"+`----'    `+"`"+`----'                         '---`+"`"+`"`

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


const glass_img = `
                    _____________________
   __....::::::::::'''''''          '''''''::::::::....__
.:'::.                                                .::':. 
 ':. ':::::.___.____,__________________ _._____..:::::' ,'
   ':.    ':::::::::::::::::::::::::::::::::::::'     : '
     ':.                                            .:'
       ':.                                        .:'
         '.                                      .'
           '-._                              _.-'
               '- .._                 _.. -'
                     ''' - .,.,. - '''
                          (:' .:)
                           :| '|
                           |. ||
                           || :|
                           :| |'
                           || :|
                           '| ||
                           |: ':
                           || :|
                     __..--:| |'--..__
               _...-'  _.' |' :| '.__ '-..._
             / -  ..---    '''''   ---...  _ \
             \  _____  ..--   --..     ____  /
              '-----....._________.....-----'`


type Title struct {
	sprite.BaseSprite
}

type Whale struct {
	sprite.BaseSprite
	Facing   int
	MinX     int
	MaxX     int
	Timer    int
	TimeOut  int
	TargetY  int
	AY       int
	VY       int
	State    int
}

type Glass struct {
	sprite.BaseSprite
	Width   int
	Counter int
}

type Bubble struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
	Dead    bool
}

const bubble_c0 = "o"
const bubble_c1 = ","
const bubble_c2 = "."

func NewTitle() *Title {
	t := &Title{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y:       2},
	}
	t.AddCostume(sprite.NewCostume(cheers_c0, 'x'))
	t.X = Width / 2 - (t.Width / 2)
	return t
}

func NewWhale(g *Glass) *Whale {
	w := &Whale{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y:       -10},
		TargetY: Height - (g.Height + 8),
		Facing:  -1,
		MinX:    g.X,
		TimeOut: 2,
		AY:      1,
		VY:      2,
		State:   falling,
	}
	w.AddCostume(sprite.NewCostume(whale_c0, 'x'))
	w.X = Width / 2 - w.Width / 2
	w.MaxX = g.X+g.Width-w.Width
	return w
}

func (w *Whale) Update() {
	if w.State == falling {
		if w.Y < w.TargetY {
			w.VY += w.AY
			w.Y += w.VY
		} else {
			w.Y = w.TargetY
			w.State = swimming
			w.Costumes = nil
			w.AddCostume(sprite.NewCostume(whale_c1, 'x'))
		}
	} else {
		w.Timer++

		if w.Timer >= w.TimeOut {
			w.Timer = 0
			w.X += w.Facing

			if w.Facing == -1 && w.X <= w.MinX {
				w.Facing = 1
				w.Costumes = nil
				w.AddCostume(sprite.NewCostume(whale_c1_rev, 'x'))
			} else if w.Facing == 1 && w.X >= w.MaxX {
				w.Facing = -1
				w.Costumes = nil
				w.AddCostume(sprite.NewCostume(whale_c1, 'x'))
			}
		}
	}
}

func NewGlass() *Glass {
	g := &Glass{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X:       Width/2-30},
		Width: 60,
	}
	g.AddCostume(sprite.NewCostume(glass_img, ' '))
	g.Y = Height - (g.Height + 5)
	return g
}

func (g *Glass) Update() {
	g.Counter++
	if g.Counter > 1 {
		b := g.NewBubble()
		bubbles = append(bubbles, b)
		allSprites.Sprites = append(allSprites.Sprites, b)
		g.Counter = 0
	}
}

func (g *Glass) NewBubble() *Bubble {
	p := g.pointInGlass()
	b := &Bubble{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X: g.X+g.Width/2+p.X,
		Y: g.Y+p.Y+10},
		TimeOut: random.Intn(5)+3,
	}
	b.AddCostume(sprite.NewCostume(bubble_c0, '~'))
	b.AddCostume(sprite.NewCostume(bubble_c1, '~'))
	b.AddCostume(sprite.NewCostume(bubble_c2, '~'))
	return b
}

func (b *Bubble) Update() {
	b.Y--
	b.Timer++
	if b.Timer > b.TimeOut {
		b.CurrentCostume++
		b.TimeOut = random.Intn(5)+3
		if b.CurrentCostume >= len(b.Costumes) {
			b.Visible = false
			b.Dead = true
		}
	}
}

func (g *Glass) pointInGlass() sprite.Point {
	x := random.Intn(g.Width-2)-(g.Width-2)/2
	y := int(math.Round(-0.010 * float64(x) * float64(x)))
	return sprite.Point{x, y}
}

func Vaccuum() {
	for cnt := len(bubbles)-1; cnt >= 0; cnt-- {
		b := bubbles[cnt]
		if b.Dead == true {
			allSprites.Remove(b)
			copy(bubbles[cnt:], bubbles[cnt+1:])
			bubbles[len(bubbles)-1] = nil
			bubbles = bubbles[:len(bubbles)-1]
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

	glass := NewGlass()

	txt := "Press 'ESC' to quit."
	c := sprite.NewCostume(txt, '~')
	text := sprite.NewBaseSprite(Width/2-len(txt)/2, Height-2, c)

	t := NewTitle()
	w := NewWhale(glass)

	allSprites.Sprites = append(allSprites.Sprites, t)
	allSprites.Sprites = append(allSprites.Sprites, glass)
	allSprites.Sprites = append(allSprites.Sprites, w)
	allSprites.Sprites = append(allSprites.Sprites, text)


mainloop:
	for {
		tm.Clear(tm.ColorDefault, tm.ColorDefault)

		select {
		case ev := <-event_queue:
			if ev.Type == tm.EventKey {
				if ev.Key == tm.KeyCtrlC || ev.Key == tm.KeyEsc || ev.Ch == 'q' {
					break mainloop
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
		Vaccuum()
	}

}
