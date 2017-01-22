// Package canvas provides GopherJS bindings for the JavaScript canvas APIs.
//
// The code is mainly based on package honnef.co/go/js/dom by Dominik Honnef
// in order to create a thin wrapper of the JavaScript canvas API.
package canvas

import (
	"image/color"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-dom"
)

const (
	// 	source-over (default)
	// 这是默认设置，新图形会覆盖在原有内容之上。
	CompositeSourceOver = "source-over"
	// destination-over
	// 会在原有内容之下绘制新图形。
	CompositeDestinationOver = "destination-over"
	// source-in
	// 新图形会仅仅出现与原有内容重叠的部分。其它区域都变成透明的。
	CompositeCompositeSourceIn = "source-in"
	// destination-in
	// 原有内容中与新图形重叠的部分会被保留，其它区域都变成透明的。
	CompositeDestinationIn = "destination-in"
	// source-out
	// 结果是只有新图形中与原有内容不重叠的部分会被绘制出来。
	CompositeSourceOut = "source-out"
	// destination-out
	// 原有内容中与新图形不重叠的部分会被保留。
	CompositeDestinationOut = "destination-out"
	// source-atop
	// 新图形中与原有内容重叠的部分会被绘制，并覆盖于原有内容之上。
	CompositeSourceAtop = "source-atop"
	// destination-atop
	// 原有内容中与新内容重叠的部分会被保留，并会在原有内容之下绘制新图形
	CompositeDestinationAtop = "destination-atop"
	// lighter
	// 两图形中重叠部分作加色处理。
	CompositeLighter = "lighter"
	// darker
	// 两图形中重叠的部分作减色处理。
	CompositeDarker = "darker"
	// xor
	// 重叠的部分会变成透明。
	CompositeXor = "xor"
	// copy
	// 只有新图形会被保留，其它都被清除掉。
	CompositeCopy = "copy"
)

const (
	PatternRepeat   = "repeat"    // (both directions),
	PatternRepeatX  = "repeat-x"  // (horizontal only),
	PatternRepeatY  = "repeat-y"  // (vertical only), or
	PatternNoRepeat = "no-repeat" // (neither).
)

// canvas元素也可以通过应用CSS的方式来增加边框，设置内边距、外边距等，
// 而且一些CSS属性还可以被canvas内的元素继承。
// 比如字体样式，在canvas内添加的文字，其样式默认同canvas元素本身是一样的。
//
// canvas是行内元素
type Canvas struct {
	*dom.Element
}

// 在canvas中为context设置属性同样要遵从CSS语法
type Context2D struct {
	*js.Object

	// 线条的颜色，默认为”#000000”，其值可以设置为CSS颜色值、渐变对象或者模式对象。
	StrokeStyle interface{} `js:"strokeStyle"`
	// 填充的颜色，默认为”#000000”，与strokeStyle一样，值也可以设置为CSS颜色值、渐变对象或者模式对象。
	FillStyle interface{} `js:"fillStyle"`
	// specifies the color of the shadow.
	// A DOMString parsed as CSS <color> value. The default value is fully-transparent black.
	ShadowColor string `js:"shadowColor"`
	// specifies the level of the blurring effect;
	// this value doesn't correspond to a number of pixels and is not affected by the current transformation matrix.
	// The default value is 0.
	ShadowBlur float64 `js:"shadowBlur"`
	// The CanvasRenderingContext2D.shadowOffsetY property of the Canvas 2D API specifies the distance that the shadow will be offset in vertical distance.
	// A float specifying the distance that the shadow will be offset in vertical distance. The default value is 0. Negative, Infinity or NaN values are ignored.
	ShadowOffsetX float64 `js:"shadowOffsetX"`
	// The CanvasRenderingContext2D.shadowOffsetX property of the Canvas 2D API specifies the distance that the shadow will be offset in horizontal distance.
	ShadowOffsetY float64 `js:"shadowOffsetY"`

	// 线条的端点样式，有butt（无）、round（圆头）、square（方头）三种类型可供选择，默认为butt。
	LineCap string `js:"lineCap"`
	// 线条的转折处样式，有round（圆角）、bevel（平角）、miter（尖角）三种；类型可供选择，默认为miter。
	LineJoin string `js:"lineJoin"`
	// 线条的宽度，单位是像素（px），默认为1.0。
	LineWidth float64 `js:"lineWidth"`
	// 线条尖角折角的锐利程序，默认为10。
	MiterLimit float64 `js:"miterLimit"`

	// A string parsed as CSS font value. The default font is '10px sans-serif'.
	// Syntax:
	//	    /* size | family */
	//	    font: 2em "Open Sans", sans-serif;
	//	    /* style | size | family */
	//	    font: italic 2em "Open Sans", sans-serif;
	//	     style | variant | weight | size/line-height | family
	//	    font: italic small-caps bolder 16px/3 cursive;
	//	    /* The font used in system dialogs */
	//	    font: message-box;
	Font string `js:"font"`
	// ctx.textAlign = "left" || "right" || "center" || "start" || "end";
	TextAlign string `js:"textAlign"`
	// ctx.textBaseline = "top" || "hanging" || "middle" || "alphabetic" || "ideographic" || "bottom";
	TextBaseline string `js:"textBaseline"`

	// Compositing
	// specifies the alpha value that is applied to shapes and images before they are drawn onto the canvas.
	// The value is in the range from 0.0 (fully transparent) to 1.0 (fully opaque).
	GlobalAlpha float64 `js:"globalAlpha"`
	// the type of compositing operation to apply when drawing new shapes,
	// where type is a string identifying which of the compositing or blending mode operations to use.
	GlobalCompositeOperation string `js:"globalCompositeOperation"`
}

// New creates a Canvas instance
// el is the html element
func New(el *js.Object) *Canvas {
	return &Canvas{dom.WrapElement(el)}
}

func (c *Canvas) GetContext2D() *Context2D {
	ctx := c.Call("getContext", "2d")
	return &Context2D{Object: ctx}
}

// canvas.toDataURL("image/jpeg") or canvas.toDataURL()
func (c *Canvas) ToDataUrl(mimeType ...string) string {
	var o *js.Object
	if len(mimeType) == 0 {
		o = c.Call("toDataURL")
	} else {
		o = c.Call("toDataURL", mimeType)
	}
	return o.String()
}

// Colors, Styles, and Shadows

type Gradient struct {
	o *js.Object
}

// The CanvasGradient.addColorStop() method adds a new stop,
// defined by an offset and a color, to the gradient.
// If the offset is not between 0 and 1, an INDEX_SIZE_ERR is raised, if the color can't be parsed as a CSS <color>, a SYNTAX_ERR is raised.
//
// offset
// 	A number between 0 and 1. An INDEX_SIZE_ERR is raised, if the number is not in that range.
// color
// 	A CSS <color>. A SYNTAX_ERR is raised, if the value can not be parsed as a CSS <color> value.
func (g *Gradient) AddColorStop(offset float64, color string) {
	g.o.Call("addColorStop", offset, color)
}

// Value returns the Object used in ctx.FillStyle/StrokeStyle
func (g *Gradient) Value() *js.Object {
	return g.o
}

// The CanvasRenderingContext2D.createLinearGradient() method of the Canvas 2D API creates a gradient along the line given by the coordinates represented by the parameters. This method returns a linear CanvasGradient.
func (ctx *Context2D) CreateLinearGradient(x0, y0, x1, y1 float64) *Gradient {
	o := ctx.Call("createLinearGradient", x0, y0, x1, y1)
	return &Gradient{o: o}
}

// The CanvasRenderingContext2D.createRadialGradient() method of the Canvas 2D API creates a radial gradient given by the coordinates of the two circles represented by the parameters. This method returns a CanvasGradient.
//
//	 Syntax
//	 	CanvasGradient ctx.createRadialGradient(x0, y0, r0, x1, y1, r1);
//
//	 x0
//		 The x axis of the coordinate of the start circle.
//	 y0
//		 The y axis of the coordinate of the start circle.
//	 r0
//		 The radius of the start circle.
//	 x1
//		 The x axis of the coordinate of the end circle.
//	 y1
//		 The y axis of the coordinate of the end circle.
//	 r1
//		 The radius of the end circle.
//
//	 example:
//			var canvas = document.getElementById("canvas");
//			var ctx = canvas.getContext("2d");
//
//			var gradient = ctx.createRadialGradient(100,100,100,100,100,0);
//			gradient.addColorStop(0,"white");
//			gradient.addColorStop(1,"green");
//			ctx.fillStyle = gradient;
//	 		ctx.fillRect(0,0,200,200);
func (ctx *Context2D) CreateRadialGradient(x0, y0, r0, x1, y1, r1 float64) *Gradient {
	o := ctx.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)
	return &Gradient{o: o}
}

// The CanvasPattern interface represents an opaque object describing a pattern, based on a image, a canvas or a video, created by the CanvasRenderingContext2D.createPattern() method.
type Pattern struct {
	o *js.Object
}

// Value returns the Object used in ctx.FillStyle/StrokeStyle
func (p *Pattern) Value() *js.Object {
	return p.o
}

// The CanvasRenderingContext2D.createPattern() method of the Canvas 2D API creates a pattern using the specified image (a CanvasImageSource).
// It repeats the source in the directions specified by the repetition argument. This method returns a CanvasPattern.
//
// 	Syntax
// 	CanvasPattern ctx.createPattern(image, repetition);
//
// 	image
// 			A CanvasImageSource to be used as image to repeat. It can either be a:
// 			HTMLImageElement (<img>),
// 			HTMLVideoElement (<video>),
// 			HTMLCanvasElement (<canvas>),
// 			CanvasRenderingContext2D,
// 			ImageBitmap,
// 			ImageData, or a
// 			Blob.
// 	repetition
// 			A DOMString indicating how to repeat the image. Possible values are:
// 			"repeat" (both directions),
// 			"repeat-x" (horizontal only),
// 			"repeat-y" (vertical only), or
// 			"no-repeat" (neither).
// 			If repetition is an empty string ('') or null (but not undefined), repetition will be "repeat".
// 	example
// 			var canvas = document.getElementById("canvas");
// 			var ctx = canvas.getContext("2d");
//
// 			var img = new Image();
// 			img.src = 'https://mdn.mozillademos.org/files/222/Canvas_createpattern.png';
// 			img.onload = function() {
// 			  var pattern = ctx.createPattern(img, 'repeat');
// 			  ctx.fillStyle = pattern;
// 			  ctx.fillRect(0,0,400,400);
// 			};
func (ctx *Context2D) CreatePattern(image *dom.Element, repetition string) *Pattern {
	o := ctx.Call("createPattern", image.Object, repetition)
	return &Pattern{o: o}
}

// void ctx.setLineDash(segments);
// Parameters
//
// segments
// An Array. A list of numbers that specifies distances to alternately draw a line and a gap (in coordinate space units).
// If the number of elements in the array is odd, the elements of the array get copied and concatenated. For example, [5, 15, 25] will become [5, 15, 25, 5, 15, 25].
func (ctx *Context2D) SetLineDash(distances ...float64) {
	ctx.Call("setLineDash", distances)
}

// ctx.getLineDash();
// Return value
//
// An Array. A list of numbers that specifies distances to alternately draw a line and a gap (in coordinate space units). If the number, when setting the elements, was odd, the elements of the array get copied and concatenated. For example, setting the line dash to [5, 15, 25] will result in getting back [5, 15, 25, 5, 15, 25].
func (ctx *Context2D) GetLineDash() []float64 {
	o := ctx.Call("getLineDash")
	return o.Interface().([]float64)
}

// 用于描绘一个已知左上角顶点位置以及宽和高的矩形，描绘完成后Context的绘制起点会移动到该矩形的左上角顶点。
//
// 参数表示矩形左上角顶点的x、y坐标以及矩形的宽和高。
func (ctx *Context2D) Rect(x, y, width, height float64) {
	ctx.Call("rect", x, y, width, height)
}

// 用于使用当前的fillStyle（默认为”#000000”，黑色）样式
// 填充一个左上角顶点在(left, top)点、宽为width、高为height的矩形。
func (ctx *Context2D) FillRect(left, top, width, height float64) {
	ctx.Call("fillRect", left, top, width, height)
}

// 用于使用当前的线条风格绘制一个左上角顶点在(left, top)点、宽为width、高为height的矩形边框。
func (ctx *Context2D) StrokeRect(left, top, width, height float64) {
	ctx.Call("strokeRect", left, top, width, height)
}

// clearRect的作用是清除矩形区域内的所有内容并将它恢复到初始状态，即透明色
// 用于清除左上角顶点在(left,top)点、宽为width、高为height的矩形区域内的所有内容。
func (ctx *Context2D) ClearRect(left, top, width, height float64) {
	ctx.Call("clearRect", left, top, width, height)
}

// Paths

// 用于使用当前的填充风格来填充路径的区域。
func (ctx *Context2D) Fill() {
	ctx.Call("fill")
}

// 用于按照已有的路径绘制线条。
func (ctx *Context2D) Stroke() {
	ctx.Call("stroke")
}

// canvas中很多用于设置样式和外观的函数也同样不会直接修改显示结果。
// HTML5 Canvas的基本图形都是以路径为基础的。通常使用Context对象的moveTo()、lineTo()、rect()、arc()等方法先在画布中描出图形的路径点，然后使用fill()或者stroke()方法依照路径点来填充图形或者绘制线条。
//
// 通常，在开始描绘路径之前需要调用Context对象的beginPath()方法，其作用是清除之前的路径并提醒Context开始绘制一条新的路径，否则当调用stroke()方法的时候会绘制之前所有的路径，影响绘制效果，同时也因为重复多次操作而影响网页性能。另外，调用Context对象的closePath()方法可以显式地关闭当前路径，不过不会清除路径。
// 只有当对路径应用绘制（stroke）或填充（fill）方法时，结果才会显示出来
func (ctx *Context2D) BeginPath() {
	ctx.Call("beginPath")
}

// 用于显式地指定路径的起点。默认状态下，第一条路径的起点是画布的(0, 0)点，之后的起点是上一条路径的终点。
// 两个参数分为表示起点的x、y坐标值。
func (ctx *Context2D) MoveTo(x, y float64) {
	ctx.Call("moveTo", x, y)
}

// 这个函数的行为同lineTo很像，唯一的差别在于closePath会将路径的起始坐标自动作为目标坐标。
// closePath还会通知canvas当前绘制的图形已经闭合或者形成了完全封闭的区域，
// 这对将来的填充和描边都非常有用。
func (ctx *Context2D) ClosePath() {
	ctx.Call("closePath")
}

// 用于描绘一条从起点从指定位置的直线路径，描绘完成后绘制的起点会移动到该指定位置。
//
// 参数表示指定位置的x、y坐标值。
func (ctx *Context2D) LineTo(x, y float64) {
	ctx.Call("lineTo", x, y)
}

// 用于按照已有的路线在画布中设置剪辑区域。
//
// 调用clip()方法之后，图形绘制代码只对剪辑区域有效而不再影响区域外的画布。
//
// 如调用之前没有描绘路径（即默认状态下），则得到的剪辑区域为整个Canvas区域。
func (ctx *Context2D) Clip() {
	ctx.Call("clip")
}

// 用于描绘一个以(x, y)点为圆心，radius为半径，startAngle为起始弧度，endAngle为终止弧度的圆弧。
// anticlockwise为布尔型的参数，true表示逆时针，false表示顺时针。
//
// quadraticCurveTo 函数绘制曲线的起点是当前坐标，带有两组（x,y）参数。第二组是指曲线的终点。
// 第一组代表控制点（control point）。
// 所谓的控制点位于曲线的旁边（不是曲线之上），其作用相当于对曲线产生一个拉力。
// 通过调整控制点的位置，就可以改变曲线的曲率。
func (ctx *Context2D) QuadraticCurveTo(cpx, cpy, x, y float64) {
	ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// 用于描绘以当前Context绘制起点为起点，(cpx1,cpy1)点和(cpx2, cpy2)点为两个控制点，
// (x, y)点为终点的贝塞尔曲线路径。
func (ctx *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// 参数中的两个弧度以0表示0°，位置在3点钟方向；Math.PI值表示180°，位置在9点钟方向。
func (ctx *Context2D) Arc(x, y, radius, sAngle, eAngle float64, counterclockwise bool) {
	ctx.Call("arc", x, y, radius, sAngle, eAngle, counterclockwise)
}

// 用于描绘一个与两条线段相切的圆弧，两条线段分别以当前Context绘制起点和(x2, y2)点为起点，都以(x1, y1)点为终点，圆弧的半径为radius。
// 描绘完成后绘制起点会移动到以(x2, y2)为起点的线段与圆弧的切点。
func (ctx *Context2D) ArcTo(x1, y1, x2, y2, r float64) {
	ctx.Call("arcTo", x1, y1, x2, y2, r)
}

func (ctx *Context2D) IsPointInPath(x, y float64) bool {
	return ctx.Call("isPointInPath", x, y).Bool()
}

// The CanvasRenderingContext2D.isPointInStroke() method of the Canvas 2D API reports whether or not the specified point is inside the area contained by the stroking of a path.
func (ctx *Context2D) IsPointInStroke(x, y float64) bool {
	return ctx.Call("isPointInStroke", x, y).Bool()
}

// The CanvasRenderingContext2D.scale() method of the Canvas 2D API adds a scaling transformation to the canvas units by x horizontally and by y vertically.
//
// By default, one unit on the canvas is exactly one pixel. If we apply, for instance, a scaling factor of 0.5, the resulting unit would become 0.5 pixels and so shapes would be drawn at half size. In a similar way setting the scaling factor to 2.0 would increase the unit size and one unit now becomes two pixels. This results in shapes being drawn twice as large.
func (ctx *Context2D) Scale(scaleWidth, scaleHeight float64) {
	ctx.Call("scale", scaleWidth, scaleHeight)
}

// The CanvasRenderingContext2D.rotate() method of the Canvas 2D API adds a rotation to the transformation matrix.
// The angle argument represents a clockwise rotation angle and is expressed in radians.
// You can use degree * Math.PI / 180 if you want to calculate from a degree value.
func (ctx *Context2D) Rotate(angle float64) {
	ctx.Call("rotate", angle)
}

// The CanvasRenderingContext2D.translate() method of the Canvas 2D API
// adds a translation transformation by moving the canvas and its origin x horizontally and y vertically on the grid.
func (ctx *Context2D) Translate(x, y float64) {
	ctx.Call("translate", x, y)
}

// The CanvasRenderingContext2D.transform() method of the Canvas 2D API
// multiplies the current transformation with the matrix described by the arguments of this method.
// You are able to scale, rotate, move and skew the context.
//    a (m11)
//    	Horizontal scaling.
//    b (m12)
//    	Horizontal skewing.
//    c (m21)
//    	Vertical skewing.
//    d (m22)
//    	Vertical scaling.
//    e (dx)
//    	Horizontal moving.
//    f (dy)
//    	Vertical moving.
func (ctx *Context2D) Transform(a, b, c, d, e, f float64) {
	ctx.Call("transform", a, b, c, d, e, f)
}

// The CanvasRenderingContext2D.setTransform() method of the Canvas 2D API
// resets (overrides) the current transformation to the identity matrix
// and then invokes a transformation described by the arguments of this method.
//    a (m11)
//    	Horizontal scaling.
//    b (m12)
//    	Horizontal skewing.
//    c (m21)
//    	Vertical skewing.
//    d (m22)
//    	Vertical scaling.
//    e (dx)
//    	Horizontal moving.
//    f (dy)
//    	Vertical moving.
func (ctx *Context2D) SetTransform(a, b, c, d, e, f float64) {
	ctx.Call("setTransform", a, b, c, d, e, f)
}

// fillText()方法能够在画布中绘制字符串
//
// 需绘制的字符串，绘制到画布中时左上角在画布中的横坐标及纵坐标，绘制的字符串的最大长度。其中最大长度maxWidth是可选参数。另外，可以通过改变Context对象的font属性来调整字符串的字体以及大小，默认为”10px sans-serif”。
func (ctx *Context2D) FillText(text string, x, y, maxWidth float64) {
	if maxWidth == -1 {
		ctx.Call("fillText", text, x, y)
		return
	}

	ctx.Call("fillText", text, x, y, maxWidth)
}

func (ctx *Context2D) StrokeText(text string, x, y, maxWidth float64) {
	if maxWidth == -1 {
		ctx.Call("strokeText", text, x, y)
		return
	}

	ctx.Call("strokeText", text, x, y, maxWidth)
}

// canvas state

// 保存当前绘图状态
func (ctx *Context2D) Save() {
	ctx.Call("save")
}

// 恢复原有的绘图状态
func (ctx *Context2D) Restore() {
	ctx.Call("restore")
}

// Context对象中拥有drawImage()方法可以将外部图片绘制到Canvas中。
//
// image参数可以是HTMLImageElement、HTMLCanvasElement或者HTMLVideoElement。
func (ctx *Context2D) DrawImage(image *dom.Element, dx, dy, dw, dh float64) {
	ctx.Call("drawImage", image, dx, dy, dw, dh)
}

type ImageData struct {
	*js.Object
	// ImageData.data Read only
	// Is a Uint8ClampedArray representing a one-dimensional array containing the data in the RGBA order, with integer values between 0 and 255 (included).
	Data *js.Object `js:"data"`
	// ImageData.height Read only
	// Is an unsigned long representing the actual height, in pixels, of the ImageData.
	Height int `js:"height"`
	// ImageData.width Read only
	// Is an unsigned long representing the actual width, in pixels, of the ImageData.
	Width int `js:"width"`
}

func (i *ImageData) Bytes() []byte {
	return js.Global.Get("Uint8Array").New(i.Data).Interface().([]byte)
}

func (i *ImageData) At(x, y int) *color.RGBA {
	idx := 4 * (y*i.Width + x)
	rgba := &color.RGBA{}
	rgba.R = uint8(i.Data.Index(idx).Int())
	rgba.G = uint8(i.Data.Index(idx + 1).Int())
	rgba.B = uint8(i.Data.Index(idx + 2).Int())
	rgba.A = uint8(i.Data.Index(idx + 3).Int())
	println("at:", x, y, rgba)
	return rgba
}

func (i *ImageData) Set(x, y int, c color.RGBA) {
	idx := 4 * (y*i.Width + x)
	i.Data.SetIndex(idx, c.R)
	i.Data.SetIndex(idx+1, c.G)
	i.Data.SetIndex(idx+2, c.B)
	i.Data.SetIndex(idx+3, c.A)
}

// func (i *ImageData) Image() image.Image {
// 	data := js.Global.Get("Uint8Array").New(i.Data).Interface().([]uint8)
// 	rgba := new(image.RGBA)
// 	rgba.Pix = data
// 	rgba.Stride = i.Width * 4
// 	rgba.Rect = image.Rect(0, 0, i.Width, i.Height)
// 	return rgba
// }

// The CanvasRenderingContext2D.createImageData() method of the Canvas 2D API creates a new, blank ImageData object with the specified dimensions.
// All of the pixels in the new object are transparent black.
// 	Syntax
// 	ImageData ctx.createImageData(width, height);
// 	ImageData ctx.createImageData(imagedata);
// 	Parameters
// 	width
// 		The width to give the new ImageData object.
// 	height
// 		The height to give the new ImageData object.
func (ctx *Context2D) CreateImageData(width, height int) *ImageData {
	o := ctx.Call("createImageData", width, height)
	im := &ImageData{Object: o}
	return im
}

// The CanvasRenderingContext2D.getImageData() method of the Canvas 2D API returns an ImageData object
// representing the underlying pixel data for the area of the canvas
// denoted by the rectangle which starts at (sx, sy) and has an sw width and sh height.
//   sx
//    The x coordinate of the upper left corner of the rectangle from which the ImageData will be extracted.
//   sy
//   	The y coordinate of the upper left corner of the rectangle from which the ImageData will be extracted.
//   sw
//   	The width of the rectangle from which the ImageData will be extracted.
//   sh
//   	The height of the rectangle from which the ImageData will be extracted.
func (c *Context2D) GetImageData(x, y, width, heigth int) *ImageData {
	o := c.Call("getImageData", x, y, width, heigth)
	return &ImageData{Object: o}
}

// The CanvasRenderingContext2D.putImageData() method of the Canvas 2D API paints data from the given ImageData object onto the bitmap. If a dirty rectangle is provided, only the pixels from that rectangle are painted.

// Syntax
// void ctx.putImageData(imagedata, dx, dy);
// void ctx.putImageData(imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight);
// 	imageData
// 		An ImageData object containing the array of pixel values.
// 	dx
// 		Position offset in the target canvas context of the rectangle to be painted, relative to the rectangle in the origin image data.
// 	dy
// 		Position offset in the target canvas context of the rectangle to be painted, relative to the rectangle in the origin image data.
// 	dirtyX Optional
// 		Position of the top left point of the rectangle to be painted, in the origin image data. Defaults to the top left of the whole image data.
// 	dirtyY Optional
// 		Position of the top left point of the rectangle to be painted, in the origin image data. Defaults to the top left of the whole image data.
// 	dirtyWidth Optional
// 		Width of the rectangle to be painted, in the origin image data. Defaults to the width of the image data.
// 	dirtyHeight Optional
// 		Height of the rectangle to be painted, in the origin image data. Defaults to the height of the image data.
func (c *Context2D) PutImageData(imd *ImageData, x, y int, dirtyX ...int) {
	args := []interface{}{imd.Object, x, y}
	for _, v := range dirtyX {
		args = append(args, v)
	}
	c.Call("putImageData", args...)
}
