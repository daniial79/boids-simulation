package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	Id       int
	Position Vector2d
	Velocity Vector2d
}

func (b *Boid) calcAcceleration() Vector2d {
	upper, lower := b.Position.AddV(viewRadius), b.Position.AddV(-viewRadius)
	aveVelocity := Vector2d{x: 0, y: 0}
	count := 0.0

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.Id {
				if dist := boids[otherBoidId].Position.Distance(b.Position); dist < viewRadius {
					count++
					aveVelocity = aveVelocity.Add(boids[otherBoidId].Velocity)
				}
			}
		}
	}

	accel := Vector2d{x: 0, y: 0}

	if count > 0 {
		aveVelocity = aveVelocity.DivisionV(count)
		accel = aveVelocity.Subtract(b.Velocity).MultiplyV(adjRate)
	}

	return accel
}

func (b *Boid) moveOne() {
	b.Velocity = b.Velocity.Add(b.calcAcceleration()).Limit(-1, 1)

	boidMap[int(b.Position.x)][int(b.Position.y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	boidMap[int(b.Position.x)][int(b.Position.y)] = b.Id

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
	boidMap[int(b.Position.x)][int(b.Position.y)] = b.Id
	go b.start()
}
