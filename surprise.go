package main

import (
	"math/rand"
	"time"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
)

var allSprites sprite.SpriteGroup
var Width int
var Height int
var random *rand.Rand

const whale_c0 = `xxxxxxxxxxxxxxx##xxxxxxxx.xxxxx
xxxxxxxxx##x##x##xxxxxxx==xxxxx
xxxxxx##x##x##x##xxxxxx===xxxxx
xx/""""""""""""""""\___/x===xxx
x{                      /xx===x
xx\______ o          __/xxxxxxx
xxxx\    \        __/xxxxxxxxxx
xxxxx\____\______/xxxxxxxxxxxxx`


const birthday_c0 = `
@@@@@@@@@@@(_)@@@@@@@(_)
@@@@@@@@@@@@#  @(_)@@@#@@@(_)@@@@@(_)
@@@@@@@@_-.-#=='-#-===#====#=(_)==.#=.==__
@@@@@@.'    #    #    #    #  #    #      `+"`"+`.
@@@@@:           #         #  #    #        :
@@@@@:.                       #            .:
@@@@@| `+"`"+`-.__                          __.-' |
@@@@@|      `+"`````"+`""============""'''''      |
@@@@@|          . . .-. .-. .-. . .         |
@@@@@|          |-| |-| |-' |-'  |          |
@@@@@|          ' ' ' ' '   '    '          |
@@@@@|  .-. .-. .-. .-. . . .-.  .-. . . .  |
@@_.-|  |(   |  |(   |  |-| |  ) |-|  |  |  |-._ 
.'   '. '_' '-' ' '  '  . . '-'  ' '  '  . .'   `+"`"+`.
:      `+"`"+`-.__                          __.-'      :
 ~.         `+"`````"+`""""""""""""""""'''''         .'
@@@`+"`"+`-.._                                  _..-'
@@@@@@@@`+"`````"+`""""-----_____------""""'''''`

type Cake struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

type Confetti struct {
	sprite.BaseSprite
	Timer     int
	TimeOut   int
	VY        int
	VYTimer   int
	VYTimeOut int
}

type Flicker struct {
	sprite.BaseSprite
	Timer   int
}

func NewCake() *Cake {
	s := &Cake{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y:       Height/2},
	}
	s.AddCostume(sprite.NewCostume(birthday_c0, '@'))
	s.X = Width/2 - s.Width/2

	points := []sprite.Point{
		sprite.Point{12, 0},
		sprite.Point{22, 0},
		sprite.Point{17, 1},
		sprite.Point{27, 1},
		sprite.Point{35, 1},
		sprite.Point{30, 2},
	}

	for cnt, p := range points {
		f := &Flicker{BaseSprite: sprite.BaseSprite{
			Visible: true,
			X:       s.X+p.X,
			Y:       s.Y+p.Y,
			},
		}
		f.AddCostume(sprite.NewCostume(")", '@'))
		f.AddCostume(sprite.NewCostume("(", '@'))
		f.CurrentCostume = cnt%2
		allSprites.Sprites = append(allSprites.Sprites, f)
	}

	return s
}

func (f *Flicker) Update() {
	f.Timer = f.Timer + 1
	if f.Timer > 4 {
		if f.CurrentCostume == 0 {
			f.CurrentCostume = 1
		} else {
			f.CurrentCostume = 0
		}
		f.Timer = 0
	}
}

func NewConfetti() *Confetti {
	s := &Confetti{BaseSprite: sprite.BaseSprite{
		Visible: true},
		Timer:   0,
		TimeOut: 3}
	s.AddCostume(sprite.NewCostume(whale_c0, 'x'))
	s.AddCostume(sprite.NewCostume(whale_c0, 'x'))
	s.X = random.Intn(Width)
	s.Y = -random.Intn(Height)
	s.VY = random.Intn(2) + 1
	s.VYTimer = 0
	s.VYTimeOut = 2
	return s
}

func (s *Confetti) Update() {
	s.Timer = s.Timer + 1
	s.VYTimer = s.VYTimer + 1
	if s.Timer > 2 {
		if s.CurrentCostume == 0 {
			s.CurrentCostume = 1
		} else {
			s.CurrentCostume = 0
		}
		s.Timer = 0
	}
	if s.VYTimer > s.VYTimeOut {
		s.Y = s.Y + s.VY
		if s.Y > Height {
			s.Y = 0 - s.Height
		}
		s.VYTimer = 0
	}
}

func main() {
	// XXX - Wait a bit until the terminal is properly initialized
	time.Sleep(1000 * time.Millisecond)
	random = rand.New(rand.NewSource(time.Now().UnixNano()))

	err := tm.Init()
	if err != nil {
		panic(err)
	}
	defer tm.Close()

	Width, Height = tm.Size()

	event_queue := make(chan tm.Event)
	go func() {
		for {
			event_queue <- tm.PollEvent()
		}
	}()

	for n := 0; n < 40; n++ {
		c := NewConfetti()
		allSprites.Sprites = append(allSprites.Sprites, c)
	}
	cake := NewCake()
	allSprites.Sprites = append(allSprites.Sprites, cake)

	txt := "Press 'ESC' to quit."
        c := sprite.NewCostume(txt, '~')
        text := sprite.NewBaseSprite(Width/2-len(txt)/2, Height-2, c)
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
			time.Sleep(30 * time.Millisecond)
		}
	}
}
