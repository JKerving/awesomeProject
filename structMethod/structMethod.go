package main

import (
	"fmt"
)

const (
	WHITE  = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type BoxList []Box

type Box struct {
	width, height, depth float64
	color                Color
}

func (b Box) Volume() float64 {
	return b.width*b.height*b.depth
}

func (b *Box) SetColor(c Color)  {
	b.color = c
}

func (bl BoxList) PaintItBlack()  {
	for i:=range bl{
		bl[i].SetColor(BLACK)
	}
}

func (bl BoxList) BiggestColor() Color{
	v:=0.00
	k:=Color(WHITE)
	for _,b:=range bl{
		if bv := b.Volume(); bv > v {
			v=bv
			k = b.color
		}
	}
	return k
}

func (c Color) String() string {
	strings:=[]string{"WHITE","BLACK","BLUE","RED","YELLOW"}
	return strings[c]
}

func main() {
	boxes:=BoxList{
		Box{4,4,4,RED},
		Box{10,10,1,YELLOW},
		Box{1,1,20,BLACK},
		Box{10,10,1,BLUE},
		Box{10,30,1,WHITE},
		Box{20,20,20,YELLOW},
	}
	fmt.Printf("%d",len(boxes))
}
