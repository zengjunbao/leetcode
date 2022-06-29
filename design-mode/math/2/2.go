package main

import (
	"fmt"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/nist"
	"go.dedis.ch/kyber/v3/util/random"
)

func main() {
	curve := nist.NewBlakeSHA256P256()

	Msg := "rush b"

	a := curve.Scalar().Pick(random.New())
	A := curve.Point().Mul(a, curve.Point().Base())

	fmt.Println("a", a)
	fmt.Println("A", A)

	K, C := Encrypt(curve, A, Msg)

	fmt.Println("K", K.String())
	fmt.Println("C", C.String())

	msg, err := Decrypt(curve, a, K, C)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}

func Encrypt(curve kyber.Group, A kyber.Point, Msg string) (K, C kyber.Point) {
	M := curve.Point().Embed([]byte(Msg), random.New())

	k := curve.Scalar().Pick(random.New())
	K = curve.Point().Mul(k, curve.Point().Base())

	kA := curve.Point().Mul(k, A)
	C = curve.Point().Add(M, kA)

	return K, C
}

func Decrypt(curve kyber.Group, a kyber.Scalar, K, C kyber.Point) (string, error) {
	S := curve.Point().Mul(a, K)
	M := curve.Point().Sub(C, S)
	data, err := M.Data()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
