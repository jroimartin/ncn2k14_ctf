package main

import (
	"os"
	"os/exec"
)

const (
	levelSizeX = 16
	levelSizeY = 16
	nLevels    = 2
	nLifes     = 3
)

type mazeEntity byte

const (
	entityPath       mazeEntity = ' '
	entityWall                  = '#'
	entityLadderUp              = 'u'
	entityLadderDown            = 'd'
	entityMonster               = 'm'
	entityFinish                = 'f'
)

type status int

const (
	statusAlive status = iota
	statusDead
	statusWin
)

type player struct {
	x, y  int
	level int
	lifes int
}

func newPlayer() *player {
	return &player{x: 0, y: 0, level: 0, lifes: nLifes}
}

func (p *player) forward() {
	p.y += 1
	if p.y > levelSizeY-1 {
		p.y = levelSizeY - 1
	}
}

func (p *player) backward() {
	p.y -= 1
	if p.y < 0 {
		p.y = 0
	}
}

func (p *player) left() {
	p.x -= 1
	if p.x < 0 {
		p.x = 0
	}
}

func (p *player) right() {
	p.x += 1
	if p.x > levelSizeX-1 {
		p.x = levelSizeX - 1
	}
}

func (p *player) nextLevel() {
	p.level += 1
	if p.level > nLevels-1 {
		p.level = nLevels - 1
	}
}

func (p *player) prevLevel() {
	p.level -= 1
	if p.level < 0 {
		p.level = 0
	}
}

func (p *player) hurt() (dead bool) {
	p.lifes -= 1
	if p.lifes < 0 {
		return true
	}
	return false
}

func rotateCharset(c rune, i int) rune {
	if c < 'a' || c > 'z' {
		return c
	}

	if cmin := c - rune(i%26); cmin < 'a' {
		return cmin + 26
	} else {
		return cmin
	}
}

func (p *player) move(c rune, i int) {
	switch rotateCharset(c, i) {
	case 'f':
		p.forward()
	case 'b':
		p.backward()
	case 'l':
		p.left()
	case 'r':
		p.right()
	default:
		return
	}
}

type level [levelSizeX][levelSizeY]mazeEntity

type maze [nLevels]level

func (m *maze) playerStatus(p *player) (s status) {
	level := m[p.level]
	switch level[p.x][p.y] {
	case entityPath:
		s = statusAlive
	case entityWall:
		s = statusDead
	case entityLadderUp:
		p.nextLevel()
		s = statusAlive
	case entityLadderDown:
		p.prevLevel()
		s = statusAlive
	case entityMonster:
		if dead := p.hurt(); dead {
			s = statusDead
		} else {
			s = statusAlive
		}
	case entityFinish:
		s = statusWin
	default:
		s = statusDead
	}
	return
}

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	p := newPlayer()
	moves := os.Args[1]
	shell := os.Args[2]

	m := generateMaze()
	for i, move := range moves {
		p.move(move, i)
		switch m.playerStatus(p) {
		case statusDead:
			os.Exit(1)
		case statusAlive:
			continue
		case statusWin:
			cmd := exec.Command(shell)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				os.Exit(1)
			}
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
	os.Exit(1)
}
