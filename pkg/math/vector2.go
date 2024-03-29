package math

type Vector2 struct {
	X int
	Y int
}

/*
*
Returns a new Vector2
@params:

	x: x-value of the vector
	y: y-value of the vector
*/
func NewVec2(x, y int) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

/*
*
Returns a new null Vector2 (0,0)
*/
func NewNullVec2() *Vector2 {
	return &Vector2{
		X: 0,
		Y: 0,
	}
}

func Add(v, other Vector2) *Vector2 {
	return &Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func Sub(v, other Vector2) *Vector2 {
	return &Vector2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}
