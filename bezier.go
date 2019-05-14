package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Bpoint struct {
	x float64
	y float64
	z float64
}

func BezierPoint(t float64, p0 *Bpoint, p1 *Bpoint, p2 *Bpoint, p3 *Bpoint) (p Bpoint) {
	u := 1 - t
	tt := t * t
	uu := u * u
	uuu := uu * u
	ttt := tt * t
	//N = 4: P = (1-t)^3*P0 + 3*(1-t)^2*t*P1 + 3(1-t)*t^2*P2 + t^3*P3
	p.x = uuu * p0.x
	p.y = uuu * p0.y
	p.z = uuu * p0.z

	p.x += 3 * uu * t * p1.x
	p.y += 3 * uu * t * p1.y
	p.z += 3 * uu * t * p1.z

	p.x += 3 * u * tt * p2.x
	p.y += 3 * u * tt * p2.y
	p.z += 3 * u * tt * p2.z

	p.x += ttt * p3.x
	p.y += ttt * p3.y
	p.z += ttt * p3.z

	p.x,_ = strconv.ParseFloat(fmt.Sprintf("%.3f",p.x),3)
	p.y,_ = strconv.ParseFloat(fmt.Sprintf("%.3f",p.y),3)
	p.z,_ = strconv.ParseFloat(fmt.Sprintf("%.3f",p.z),3)

	return p
}
func Getbezier(p0, p1,p2, p3 *Bpoint, pCount int,sp []Bpoint) (length float64) {
	var t float64
	lastPoint := *p0
	for i := 0; i <= pCount; i++ {
		t = (float64)(i)/(float64)(pCount)
		//fmt.Printf("第%d个点:t=%v",i,t)
		p := BezierPoint(t, p0, p1, p2, p3)
		length += math.Sqrt((math.Pow(p.x - lastPoint.x,2)  + math.Pow(p.y - lastPoint.y,2) + math.Pow(p.z - lastPoint.z,2)))
		lastPoint = p
		sp[i] = p
		//fmt.Println(sp)
		//fmt.Printf("第%d个点：%.1f,%.1f,%.1f\n",i,p.x,p.y,p.z)
	}
	return length
}
func Getlocation(ntime float64, ttime float64,length... float64) (point Bpoint){
	return point
}
func main() {
	var spcount int=300	//采样点个数
	var sp1 = make([]Bpoint,spcount+1)
	var sp2 = make([]Bpoint,spcount+1)
	var length1,length2 float64
	fmt.Println("N等分上的Bezier采样点: ")
	p0 := &Bpoint{0,0,20}	//起点
	p1 := &Bpoint{1,0,0}		//控制点1
	p2 := &Bpoint{22.32,-2.9,-3.16}	//控制点2
	p3 := &Bpoint{7.7,-2.9,-1.3}		//终点
	t1 := time.Now()//计算生成鱼线的耗时
	for i:=0;i<1;i++{
		length1 = Getbezier(p0,p1,p2,p3,spcount,sp1)
		fmt.Printf("此贝塞尔曲线总长为：%.6f\n",length1)
	}
	p20 := &Bpoint{7.7,-2.9,-1.3}
	p21 := &Bpoint{-6.92,-2.9,0.56}
	p22 := &Bpoint{49,20,0}
	p23 := &Bpoint{50,20,0}
	for i:=0;i<1;i++{
		length2 = Getbezier(p20,p21,p22,p23,spcount,sp2)
		fmt.Printf("此贝塞尔曲线总长为：%.6f\n",length2)
	}
	fmt.Println(sp1)
	fmt.Println(sp2)
	ntime := 3.2
	ttime := 5.2
	totallength := length1 + length2
	nowlength := totallength*ntime/ttime
	var t float64
	var np Bpoint
	if nowlength<length1{
		t = nowlength/length1
		np = BezierPoint(t,p0,p1,p2,p3)
	}else if nowlength<length1+length2{
		t = (nowlength-length1)/length2
		np = BezierPoint(t,p20,p21,p22,p23)
	}
	fmt.Println(t)
	fmt.Println(np)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
