package main

import (
	"math/rand"
	"time"
)

type Bird struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Bird) moveOne() {
	b.position = b.position.Addition(b.velocity)
	next := b.position.Addition(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{
			x: -b.velocity.x,
			y: b.velocity.y,
		}
	}
	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{
			x: b.velocity.x,
			y: -b.velocity.y,
		}
	}
}

func (b *Bird) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBird(bird int) {
	b := Bird{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       bird,
	}
	birds[bird] = &b
	go b.start()
}
