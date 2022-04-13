package voronoi

import (
	"image"
)

type FractalTree int

const (
	HTree FractalTree = iota
	TSquare2
	TSquare3
	Sierpinski
)

type Fractal struct {
	w, h           int
	x, y           int
	nw, ne, se, sw *Fractal
}

func (f *Fractal) Split(ft FractalTree) {
	stack := FStack{f}
	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		if cur.ne != nil {
			stack.Push(cur.ne)
		}
		if cur.nw != nil {
			stack.Push(cur.nw)
		}
		if cur.se != nil {
			stack.Push(cur.se)
		}
		if cur.sw != nil {
			stack.Push(cur.sw)
		}
		if cur.ne == nil && cur.nw == nil && cur.se == nil && cur.sw == nil {
			if ft == HTree {
				cur.nw = &Fractal{cur.w / 2, cur.h / 2, cur.x - cur.w/2, cur.y - cur.h/2, nil, nil, nil, nil}
				cur.ne = &Fractal{cur.w / 2, cur.h / 2, cur.x + cur.w/2, cur.y - cur.h/2, nil, nil, nil, nil}
				cur.se = &Fractal{cur.w / 2, cur.h / 2, cur.x + cur.w/2, cur.y + cur.h/2, nil, nil, nil, nil}
				cur.sw = &Fractal{cur.w / 2, cur.h / 2, cur.x - cur.w/2, cur.y + cur.h/2, nil, nil, nil, nil}
			} else if ft == TSquare3 {
				w, h := cur.w/2, cur.h/2
				cur.nw = &Fractal{w, h, cur.x - w - w/2, cur.y - h - h/2, nil, nil, nil, nil}
				cur.ne = &Fractal{w, h, cur.x + w + w/2, cur.y - h - h/2, nil, nil, nil, nil}
				cur.se = &Fractal{w, h, cur.x + w + w/2, cur.y + h + h/2, nil, nil, nil, nil}
				cur.sw = &Fractal{w, h, cur.x - w - w/2, cur.y + h + h/2, nil, nil, nil, nil}
			} else if ft == TSquare2 {
				w, h := cur.w/2, cur.h/2
				cur.nw = &Fractal{w, h, cur.x - w, cur.y - h, nil, nil, nil, nil}
				cur.ne = &Fractal{w, h, cur.x + w, cur.y - h, nil, nil, nil, nil}
				cur.se = &Fractal{w, h, cur.x + w, cur.y + h, nil, nil, nil, nil}
				cur.sw = &Fractal{w, h, cur.x - w, cur.y + h, nil, nil, nil, nil}
			} else if ft == Sierpinski {

			}
		}
	}
}

func (f *Fractal) Draw(img *image.RGBA, ft FractalTree) {
	if ft == HTree {
		for i := f.x - f.w/2; i < f.x+f.w/2; i++ {
			img.SetRGBA(i, f.y, White)
		}
		for i := f.y - f.h/2; i < f.y+f.h/2; i++ {
			img.SetRGBA(f.x-f.w/2, i, White)
			img.SetRGBA(f.x+f.w/2, i, White)
		}
	} else if ft == TSquare2 || ft == TSquare3 {
		for i := f.x - f.w/2; i < f.x+f.w/2; i++ {
			for j := f.y - f.h/2; j < f.y+f.h/2; j++ {
				img.SetRGBA(i, j, White)
			}
		}
	} else if ft == Sierpinski {
		for i := f.x - f.w/2; i < f.x+f.w/2; i++ {
			img.SetRGBA(i, f.y, White)
		}
		for i := f.x - f.w/2; i < f.x; i++ {
			img.SetRGBA(i, f.y/2+int(float64(i)/1.0472), White)
		}
	}
}

type FStack []*Fractal

func (s *FStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *FStack) Push(qt *Fractal) {
	*s = append(*s, qt)
}

func (s *FStack) Pop() (*Fractal, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
