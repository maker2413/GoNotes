package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	initBallSpeed = 2
	paddleSpeed   = 6
)

type Object struct {
	X, Y, W, H int
}

type Paddle struct {
	Object
}

type Ball struct {
	Object
	dxdt      int // x velocity per tick
	dydt      int // y velocity per tick
	ballSpeed int
	iFrames   int
}

type Game struct {
	paddle    Paddle
	ball      Ball
	score     int
	highScore int
	newGame   int
}

func main() {
	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	paddle := Paddle{
		Object: Object{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	ball := Ball{
		Object: Object{
			X: 10,
			Y: 200,
			W: 15,
			H: 15,
		},
		ballSpeed: initBallSpeed,
		dxdt:      initBallSpeed,
		dydt:      initBallSpeed,
	}

	g := &Game{
		paddle: paddle,
		ball:   ball,
	}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(g.paddle.X), float32(g.paddle.Y),
		float32(g.paddle.W), float32(g.paddle.H),
		color.White, false)
	vector.DrawFilledRect(screen,
		float32(g.ball.X), float32(g.ball.Y),
		float32(g.ball.W), float32(g.ball.H),
		color.White, false)

	scoreStr := "Score: " + fmt.Sprint(g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 10, color.White)

	highScoreStr := "High Score: " + fmt.Sprint(g.highScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 30, color.White)
}

func (g *Game) Update() error {
	g.paddle.MoveOnKeyPress()
	if g.newGame >= 60 {
		g.ball.Move()
	} else {
		g.newGame++
	}
	g.CollideWithWall()
	g.CollideWithPaddle()
	return nil
}

func (p *Paddle) MoveOnKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Y <= screenHeight-100 {
		p.Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Y >= 0 {
		p.Y -= paddleSpeed
	}
}

func (b *Ball) Move() {
	if b.iFrames <= 10 {
		b.iFrames++
	}
	b.X += b.dxdt
	b.Y += b.dydt
}

func (g *Game) Reset() {
	g.ball.ballSpeed = initBallSpeed
	g.ball.dxdt = initBallSpeed
	g.ball.dydt = initBallSpeed
	g.ball.X = 10
	g.ball.Y = 200

	g.score = 0
	g.newGame = 0
}

func (g *Game) CollideWithWall() {
	// Right wall cases a game over
	if g.ball.X >= screenWidth-g.ball.W {
		g.Reset()
	} else if g.ball.X <= 0 {
		g.ball.dxdt = g.ball.ballSpeed
	} else if g.ball.Y <= 0 {
		g.ball.dydt = g.ball.ballSpeed
	} else if g.ball.Y >= screenHeight-g.ball.H {
		g.ball.dydt = -g.ball.ballSpeed
	}
}

func (g *Game) CollideWithPaddle() {
	if g.ball.X >= g.paddle.X-g.ball.W && g.ball.X <= (g.paddle.X+g.paddle.W)-g.ball.W && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y+g.paddle.H && g.ball.iFrames > 5 {
		g.ball.dxdt = -g.ball.dxdt
		g.ball.iFrames = 0
		g.score++
		g.ball.ballSpeed++
		if g.score > g.highScore {
			g.highScore = g.score
		}
	}
}
