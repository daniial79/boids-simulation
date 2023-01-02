package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	Id       int
	Position Vector2d
	Velocity Vector2d
}

func (b *Boid) moveOne() {
	b.Position = b.Position.Add(b.Velocity)
	next := b.Position.Add(b.Velocity)

	if next.x >= screenWidth || next.x < 0 {
		b.Velocity = Vector2d{
			x: -b.Velocity.x,
			y: b.Velocity.y,
		}
	}

	if next.y >= screenHeight || next.y < 0 {
		b.Velocity = Vector2d{
			x: b.Velocity.x,
			y: -b.Velocity.y,
		}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	rand.Seed(time.Now().UnixNano())
	b := Boid{
		Id: bid,
		Position: Vector2d{
			x: rand.Float64() * screenWidth,
			y: rand.Float64() * screenHeight,
		},
		Velocity: Vector2d{
			x: (rand.Float64() * 2) - 1.0,
			y: (rand.Float64() * 2) - 1.0,
		},
	}

	boids[bid] = &b

	go b.start()
}
