package vector

type Vector struct {
	X int32
	Y int32
	Z int32
}

func Sub(a, b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: 0,
	}
}

func CrossProduct(a, b Vector) Vector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}
