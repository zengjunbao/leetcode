package main

import "math/big"

type ICurve interface {
	IsOnCurve(P Point) bool
	Add(P, Q Point) Point
	Mul(n *big.Int, P Point) Point
}

type Point struct {
	X, Y *big.Int
}

type Curve struct {
	A   *big.Int
	B   *big.Int
	P   *big.Int
}
