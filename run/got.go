package run

type Gotter interface {
	Got(*Under) Funct
}

func (u *Under) Got(f Funct) Funct {
	if g, ok := f.(Gotter); ok {
		return g.Got(u)
	}

	return f
}
