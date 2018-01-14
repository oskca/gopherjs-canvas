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

// The CanvasRenderingContext2D.globalCompositeOperation property of the Canvas 2D API sets the type of
// compositing operation to apply when drawing new shapes, where type is a string identifying which of
// the compositing or blending mode operations to use.
const (
	// This is the default setting and draws new shapes on top of the existing canvas content.
	CompositeSourceOver = "source-over"
	// New shapes are drawn behind the existing canvas content.
	CompositeDestinationOver = "destination-over"
	// The new shape is drawn only where both the new shape and the destination canvas overlap.
	// Everything else is made transparent.
	CompositeSourceIn = "source-in"
	// The existing canvas content is kept where both the new shape and existing canvas content overlap.
	// Everything else is made transparent.
	CompositeDestinationIn = "destination-in"
	// The new shape is drawn where it doesn't overlap the existing canvas content.
	CompositeSourceOut = "source-out"
	// The existing content is kept where it doesn't overlap the new shape.
	CompositeDestinationOut = "destination-out"
	// The new shape is only drawn where it overlaps the existing canvas content.
	CompositeSourceAtop = "source-atop"
	// The existing canvas is only kept where it overlaps the new shape.
	// The new shape is drawn behind the canvas content.
	CompositeDestinationAtop = "destination-atop"
	// Where both shapes overlap the color is determined by adding color values.
	CompositeLighter = "lighter"
	// Shapes are made transparent where both overlap and drawn normal everywhere else.
	CompositeXor = "xor"
	// Only the new shape is shown.
	CompositeCopy = "copy"
)

// Repeat Patterns
const (
	PatternRepeat   = "repeat"    // (both directions),
	PatternRepeatX  = "repeat-x"  // (horizontal only),
	PatternRepeatY  = "repeat-y"  // (vertical only), or
	PatternNoRepeat = "no-repeat" // (neither).
)

// Canvas The HTML5 <canvas> tag is used to draw graphics, on the fly,
// via scripting (usually JavaScript).
type Canvas struct {
	*dom.Element
}

// Context2D struct
type Context2D struct {
	*js.Object

	// The CanvasRenderingContext2D.strokeStyle property of the Canvas 2D API specifies the color or
	// style to use for the lines around shapes. The default is #000 (black).
	StrokeStyle interface{} `js:"strokeStyle"`
	// Color or style to use inside shapes. Default #000 (black).
	FillStyle interface{} `js:"fillStyle"`
	// specifies the color of the shadow.
	// A DOMString parsed as CSS <color> value. The default value is fully-transparent black.
	ShadowColor string `js:"shadowColor"`
	// specifies the level of the blurring effect;
	// this value doesn't correspond to a number of pixels and is not affected by
	// the current transformation matrix.
	// The default value is 0.
	ShadowBlur float64 `js:"shadowBlur"`
	// The CanvasRenderingContext2D.shadowOffsetY property of the Canvas 2D API specifies
	// the distance that the shadow will be offset in vertical distance.
	// A float specifying the distance that the shadow will be offset in vertical distance.
	// The default value is 0. Negative, Infinity or NaN values are ignored.
	ShadowOffsetX float64 `js:"shadowOffsetX"`
	// The CanvasRenderingContext2D.shadowOffsetX property of the Canvas 2D API specifies the distance that the shadow will be offset in horizontal distance.
	ShadowOffsetY float64 `js:"shadowOffsetY"`

	// Type of endings on the end of lines. Possible values: butt (default), round, square.
	LineCap string `js:"lineCap"`
	// Defines the type of corners where two lines meet. Possible values: round, bevel, miter (default).
	LineJoin string `js:"lineJoin"`
	// Width of lines. Default 1.0
	LineWidth float64 `js:"lineWidth"`
	// Miter limit ratio. Default 10.
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

// GetContext2D returns the Context2D object
func (c *Canvas) GetContext2D() *Context2D {
	ctx := c.Call("getContext", "2d")
	return &Context2D{Object: ctx}
}

// ToDataUrl canvas.toDataURL("image/jpeg") or canvas.toDataURL()
func (c *Canvas) ToDataUrl(mimeType ...string) string {
	var o *js.Object
	if len(mimeType) == 0 {
		o = c.Call("toDataURL")
	} else {
		o = c.Call("toDataURL", mimeType)
	}
	return o.String()
}

// Gradient Colors, Styles, and Shadows
type Gradient struct {
	o *js.Object
}

// AddColorStop The CanvasGradient.addColorStop() method adds a new stop,
// defined by an offset and a color, to the gradient.
// If the offset is not between 0 and 1, an INDEX_SIZE_ERR is raised, if the color can't be parsed as
// a CSS <color>, a SYNTAX_ERR is raised.
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

// CreateLinearGradient The CanvasRenderingContext2D.createLinearGradient() method of
// the Canvas 2D API creates a gradient along the line given by the coordinates represented by
// the parameters. This method returns a linear CanvasGradient.
func (ctx *Context2D) CreateLinearGradient(x0, y0, x1, y1 float64) *Gradient {
	o := ctx.Call("createLinearGradient", x0, y0, x1, y1)
	return &Gradient{o: o}
}

// CreateRadialGradient The CanvasRenderingContext2D.createRadialGradient() method of
// the Canvas 2D API creates a radial gradient given by the coordinates of the two circles represented by
// the parameters. This method returns a CanvasGradient.
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

// Pattern The CanvasPattern interface represents an opaque object describing a pattern, based on a image,
// a canvas or a video, created by the CanvasRenderingContext2D.createPattern() method.
type Pattern struct {
	o *js.Object
}

// Value returns the Object used in ctx.FillStyle/StrokeStyle
func (p *Pattern) Value() *js.Object {
	return p.o
}

// CreatePattern The CanvasRenderingContext2D.createPattern() method of the Canvas 2D API creates a
// pattern using the specified image (a CanvasImageSource).
// It repeats the source in the directions specified by the repetition argument. This method returns a
// CanvasPattern.
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

// SetLineDash Sets the current line dash pattern.
// Parameters
//
// segments
// An Array. A list of numbers that specifies distances to alternately draw a line and
// a gap (in coordinate space units).
// If the number of elements in the array is odd, the elements of the array get copied and
// concatenated. For example, [5, 15, 25] will become [5, 15, 25, 5, 15, 25].
func (ctx *Context2D) SetLineDash(distances ...float64) {
	ctx.Call("setLineDash", distances)
}

// GetLineDash Returns the current line dash pattern array containing an even number of non-negative numbers.
func (ctx *Context2D) GetLineDash() []float64 {
	o := ctx.Call("getLineDash")
	return o.Interface().([]float64)
}

// Rect The CanvasRenderingContext2D.rect() method of the Canvas 2D API creates a path for
// a rectangle at position (x, y) with a size that is determined by width and height.
// Those four points are connected by straight lines and the sub-path is marked as closed,
// so that you can fill or stroke this rectangle.
func (ctx *Context2D) Rect(x, y, width, height float64) {
	ctx.Call("rect", x, y, width, height)
}

// FillRect Draws a filled rectangle at (x, y) position whose size is determined by width and height.
func (ctx *Context2D) FillRect(left, top, width, height float64) {
	ctx.Call("fillRect", left, top, width, height)
}

// StrokeRect Paints a rectangle which has a starting point at (x, y) and has a w width and
// an h height onto the canvas, using the current stroke style.
func (ctx *Context2D) StrokeRect(left, top, width, height float64) {
	ctx.Call("strokeRect", left, top, width, height)
}

// ClearRect Sets all pixels in the rectangle defined by starting point (x, y) and
// size (width, height) to transparent black, erasing any previously drawn content.
func (ctx *Context2D) ClearRect(left, top, width, height float64) {
	ctx.Call("clearRect", left, top, width, height)
}

// Paths

// Fill The CanvasRenderingContext2D.fill() method of the Canvas 2D API fills the current or
// given path with the current fill style using the non-zero or even-odd winding rule.
func (ctx *Context2D) Fill() {
	ctx.Call("fill")
}

// Stroke The CanvasRenderingContext2D.stroke() method of the Canvas 2D API strokes the current or
// given path with the current stroke style using the non-zero winding rule.
func (ctx *Context2D) Stroke() {
	ctx.Call("stroke")
}

// BeginPath Starts a new path by emptying the list of sub-paths.
// Call this method when you want to create a new path.
func (ctx *Context2D) BeginPath() {
	ctx.Call("beginPath")
}

// MoveTo Moves the starting point of a new sub-path to the (x, y) coordinates.
func (ctx *Context2D) MoveTo(x, y float64) {
	ctx.Call("moveTo", x, y)
}

//ClosePath  Causes the point of the pen to move back to the start of
// the current sub-path. It tries to draw a straight line from
// the current point to the start. If the shape has already been closed or
// has only one point, this function does nothing.
func (ctx *Context2D) ClosePath() {
	ctx.Call("closePath")
}

// LineTo Connects the last point in the subpath to the x, y coordinates with a straight line.
func (ctx *Context2D) LineTo(x, y float64) {
	ctx.Call("lineTo", x, y)
}

// Clip Creates a clipping path from the current sub-paths.
// Everything drawn after clip() is called appears inside the clipping path only.
func (ctx *Context2D) Clip() {
	ctx.Call("clip")
}

// QuadraticCurveTo Adds a quadratic Bézier curve to the current path.
func (ctx *Context2D) QuadraticCurveTo(cpx, cpy, x, y float64) {
	ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// BezierCurveTo Adds a cubic Bézier curve to the path. It requires three points. The first two points are
// control points and the third one is the end point. The starting point is the last
// point in the current path, which can be changed using moveTo() before creating the Bézier curve.
func (ctx *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// Arc Adds an arc to the path which is centered at (x, y) position with
// radius r starting at startAngle and ending at endAngle going in the given direction by
// anticlockwise (defaulting to clockwise).
func (ctx *Context2D) Arc(x, y, radius, sAngle, eAngle float64, counterclockwise bool) {
	ctx.Call("arc", x, y, radius, sAngle, eAngle, counterclockwise)
}

// ArcTo Adds an arc to the path with the given control points and radius,
// connected to the previous point by a straight line.
func (ctx *Context2D) ArcTo(x1, y1, x2, y2, r float64) {
	ctx.Call("arcTo", x1, y1, x2, y2, r)
}

// IsPointInPath Reports whether or not the specified point is contained in the current path.
func (ctx *Context2D) IsPointInPath(x, y float64) bool {
	return ctx.Call("isPointInPath", x, y).Bool()
}

// IsPointInStroke The CanvasRenderingContext2D.isPointInStroke() method of
//  the Canvas 2D API reports whether or not the specified point is inside the area contained by the stroking of a path.
func (ctx *Context2D) IsPointInStroke(x, y float64) bool {
	return ctx.Call("isPointInStroke", x, y).Bool()
}

// Scale The CanvasRenderingContext2D.scale() method of the Canvas 2D API adds a
// scaling transformation to the canvas units by x horizontally and by y vertically.
// By default, one unit on the canvas is exactly one pixel. If we apply, for instance,
// a scaling factor of 0.5, the resulting unit would become 0.5 pixels and
// so shapes would be drawn at half size. In a similar way setting the scaling factor to
// 2.0 would increase the unit size and one unit now becomes two pixels.
// This results in shapes being drawn twice as large.
func (ctx *Context2D) Scale(scaleWidth, scaleHeight float64) {
	ctx.Call("scale", scaleWidth, scaleHeight)
}

// Rotate The CanvasRenderingContext2D.rotate() method of the Canvas 2D API adds a
// rotation to the transformation matrix.
// The angle argument represents a clockwise rotation angle and is expressed in radians.
// You can use degree * Math.PI / 180 if you want to calculate from a degree value.
func (ctx *Context2D) Rotate(angle float64) {
	ctx.Call("rotate", angle)
}

// Translate The CanvasRenderingContext2D.translate() method of the Canvas 2D API
// adds a translation transformation by moving the canvas and its origin x horizontally and y vertically on the grid.
func (ctx *Context2D) Translate(x, y float64) {
	ctx.Call("translate", x, y)
}

// Transform The CanvasRenderingContext2D.transform() method of the Canvas 2D API
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

// SetTransform The CanvasRenderingContext2D.setTransform() method of the Canvas 2D API
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

// FillText Draws (fills) a given text at the given (x,y) position.
func (ctx *Context2D) FillText(text string, x, y, maxWidth float64) {
	if maxWidth == -1 {
		ctx.Call("fillText", text, x, y)
		return
	}

	ctx.Call("fillText", text, x, y, maxWidth)
}

// StrokeText Draws (strokes) a given text at the given (x, y) position.
func (ctx *Context2D) StrokeText(text string, x, y, maxWidth float64) {
	if maxWidth == -1 {
		ctx.Call("strokeText", text, x, y)
		return
	}

	ctx.Call("strokeText", text, x, y, maxWidth)
}

// canvas state

// Save Saves the current drawing style state using
// a stack so you can revert any change you make to it using restore()
func (ctx *Context2D) Save() {
	ctx.Call("save")
}

// Restore Restores the drawing style state to the last element on the 'state stack' saved by save().
func (ctx *Context2D) Restore() {
	ctx.Call("restore")
}

// DrawImage Draws the specified image. This method is available in multiple formats,
// providing a great deal of flexibility in its use.
func (ctx *Context2D) DrawImage(image *dom.Element, dx, dy, dw, dh float64) {
	ctx.Call("drawImage", image, dx, dy, dw, dh)
}

// ImageData struct
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

// Bytes ImageData Bytes
func (i *ImageData) Bytes() []byte {
	return js.Global.Get("Uint8Array").New(i.Data).Interface().([]byte)
}

// At ImageData At
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

// Set ImageData Set
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

// CreateImageData The CanvasRenderingContext2D.createImageData() method of the Canvas 2D API creates a new, blank ImageData object with the specified dimensions.
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

// GetImageData The CanvasRenderingContext2D.getImageData() method of the Canvas 2D API returns an ImageData object
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
func (ctx *Context2D) GetImageData(x, y, width, heigth int) *ImageData {
	o := ctx.Call("getImageData", x, y, width, heigth)
	return &ImageData{Object: o}
}

// PutImageData The CanvasRenderingContext2D.putImageData() method of the Canvas 2D API paints data from
// the given ImageData object onto the bitmap. If a dirty rectangle is provided,
// only the pixels from that rectangle are painted.
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
func (ctx *Context2D) PutImageData(imd *ImageData, x, y int, dirtyX ...int) {
	args := []interface{}{imd.Object, x, y}
	for _, v := range dirtyX {
		args = append(args, v)
	}
	ctx.Call("putImageData", args...)
}
