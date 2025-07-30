/*
Minecraft Head Rendering in Rust provided by @mat-1: https://github.com/mat-1/msdsmchr
Rewritten and converted into Go by @DuckySolucky
Originally Inspired by Crafatar: https://github.com/crafatar/crafatar
*/

package lib

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"strings"
)

const (
	SKEW_A       = 26.0 / 45.0
	SKEW_B       = SKEW_A * 2.0
	SECTION_SIZE = 8
)

var (
	TRANSFORM_TOP_BOTTOM_MATRIX = [9]float32{
		1.0, 1.0, 0.0,
		-SKEW_A, SKEW_A, 0.0,
		0.0, 0.0, 1.0,
	}
	TRANSFORM_FRONT_BACK_MATRIX = [9]float32{
		1.0, 0.0, 0.0,
		-SKEW_A, SKEW_B, SKEW_A,
		0.0, 0.0, 1.0,
	}
	TRANSFORM_RIGHT_LEFT_MATRIX = [9]float32{
		1.0, 0.0, 0.0,
		SKEW_A, SKEW_B, 0.0,
		0.0, 0.0, 1.0,
	}
)

type OverlaySectionOptions struct {
	Size       uint32
	Scale      float32
	X          uint32
	Y          uint32
	Matrix     [9]float32
	TranslateX float32
	TranslateY float32
	Flip       bool
}

func cropSection(img image.Image, x, y uint32) image.Image {
	bounds := img.Bounds()
	x1 := int(x * SECTION_SIZE)
	y1 := int(y * SECTION_SIZE)

	cropped := image.NewRGBA(image.Rect(0, 0, SECTION_SIZE, SECTION_SIZE))
	for dy := 0; dy < SECTION_SIZE; dy++ {
		for dx := 0; dx < SECTION_SIZE; dx++ {
			srcX := x1 + dx
			srcY := y1 + dy

			if srcX >= bounds.Min.X && srcX < bounds.Max.X && srcY >= bounds.Min.Y && srcY < bounds.Max.Y {
				cropped.Set(dx, dy, img.At(srcX, srcY))
			}
		}
	}

	return cropped
}

type Matrix [9]float32

type Point2D struct {
	X, Y float32
}

func (m Matrix) Transform(p Point2D) Point2D {
	x := m[0]*p.X + m[1]*p.Y + m[2]
	y := m[3]*p.X + m[4]*p.Y + m[5]
	w := m[6]*p.X + m[7]*p.Y + m[8]

	if w != 0 {
		x /= w
		y /= w
	}

	return Point2D{X: x, Y: y}
}

func (m Matrix) Multiply(other Matrix) Matrix {
	var result Matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i*3+j] += m[i*3+k] * other[k*3+j]
			}
		}
	}
	return result
}

func Scale(sx, sy float32) Matrix {
	return Matrix{
		sx, 0, 0,
		0, sy, 0,
		0, 0, 1,
	}
}

func Translate(tx, ty float32) Matrix {
	return Matrix{
		1, 0, tx,
		0, 1, ty,
		0, 0, 1,
	}
}

func blendRGBA(bottom, top color.RGBA) color.RGBA {
	if top.A == 0 {
		return bottom
	}
	if top.A == 255 {
		return top
	}

	alpha := float64(top.A) / 255.0
	invAlpha := 1.0 - alpha

	return color.RGBA{
		R: uint8(float64(top.R)*alpha + float64(bottom.R)*invAlpha),
		G: uint8(float64(top.G)*alpha + float64(bottom.G)*invAlpha),
		B: uint8(float64(top.B)*alpha + float64(bottom.B)*invAlpha),
		A: uint8(math.Min(255, float64(top.A)+float64(bottom.A)*invAlpha)),
	}
}

func fastOverlay(bottom *image.RGBA, top *image.RGBA) {
	bounds := bottom.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			topColor := top.RGBAAt(x, y)

			if topColor.A == 0 {
				continue
			}

			bottomColor := bottom.RGBAAt(x, y)
			blended := blendRGBA(bottomColor, topColor)
			bottom.SetRGBA(x, y, blended)
		}
	}
}

func warpImage(src *image.RGBA, matrix Matrix, outputSize uint32) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, int(outputSize), int(outputSize)))

	det := matrix[0]*(matrix[4]*matrix[8]-matrix[5]*matrix[7]) -
		matrix[1]*(matrix[3]*matrix[8]-matrix[5]*matrix[6]) +
		matrix[2]*(matrix[3]*matrix[7]-matrix[4]*matrix[6])

	if math.Abs(float64(det)) < 1e-10 {
		return dst
	}

	invMatrix := Matrix{
		(matrix[4]*matrix[8] - matrix[5]*matrix[7]) / det,
		(matrix[2]*matrix[7] - matrix[1]*matrix[8]) / det,
		(matrix[1]*matrix[5] - matrix[2]*matrix[4]) / det,
		(matrix[5]*matrix[6] - matrix[3]*matrix[8]) / det,
		(matrix[0]*matrix[8] - matrix[2]*matrix[6]) / det,
		(matrix[2]*matrix[3] - matrix[0]*matrix[5]) / det,
		(matrix[3]*matrix[7] - matrix[4]*matrix[6]) / det,
		(matrix[1]*matrix[6] - matrix[0]*matrix[7]) / det,
		(matrix[0]*matrix[4] - matrix[1]*matrix[3]) / det,
	}

	srcBounds := src.Bounds()

	for y := 0; y < int(outputSize); y++ {
		for x := 0; x < int(outputSize); x++ {
			srcPoint := invMatrix.Transform(Point2D{X: float32(x), Y: float32(y)})

			srcX := int(math.Round(float64(srcPoint.X)))
			srcY := int(math.Round(float64(srcPoint.Y)))

			if srcX >= srcBounds.Min.X && srcX < srcBounds.Max.X &&
				srcY >= srcBounds.Min.Y && srcY < srcBounds.Max.Y {
				dst.Set(x, y, src.At(srcX, srcY))
			}
		}
	}

	return dst
}

func overlay3DSection(out *image.RGBA, skin image.Image, opts *OverlaySectionOptions) {
	section := cropSection(skin, opts.X, opts.Y)

	sectionRGBA, ok := section.(*image.RGBA)
	if !ok {
		bounds := section.Bounds()
		sectionRGBA = image.NewRGBA(bounds)
		draw.Draw(sectionRGBA, bounds, section, bounds.Min, draw.Src)
	}

	baseMatrix := Matrix(opts.Matrix)
	translateMatrix := Translate(opts.TranslateX, opts.TranslateY)

	scaleX := opts.Scale
	if opts.Flip {
		scaleX = -scaleX
	}
	scaleMatrix := Scale(scaleX, opts.Scale)

	finalMatrix := baseMatrix.Multiply(translateMatrix).Multiply(scaleMatrix)

	sectionWarped := warpImage(sectionRGBA, finalMatrix, opts.Size)

	fastOverlay(out, sectionWarped)
}

func To3DHead(img image.Image) *image.RGBA {
	size := uint32(128)

	out := image.NewRGBA(image.Rect(0, 0, int(size), int(size)))

	// Left overlay
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          6,
		Y:          1,
		Matrix:     TRANSFORM_RIGHT_LEFT_MATRIX,
		TranslateX: float32(size) * (231.0 / 256.0) * (8.0 / 8.1),
		TranslateY: float32(size) * (-56.0 / 256.0),
		Flip:       true,
		Scale:      float32(size/20) * (9.0 / 8.0),
	})

	// Back overlay
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          7,
		Y:          1,
		Matrix:     TRANSFORM_FRONT_BACK_MATRIX,
		TranslateX: float32(size) * (26.0 / 256.0),
		TranslateY: float32(size) * (70.0 / 256.0),
		Flip:       false,
		Scale:      float32(size/20) * (9.0 / 8.0),
	})

	// Bottom overlay
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          6,
		Y:          0,
		Matrix:     TRANSFORM_TOP_BOTTOM_MATRIX,
		TranslateX: float32(size) * (-145.0 / 256.0),
		TranslateY: float32(size) * (177.0 / 256.0),
		Flip:       false,
		Scale:      float32(size/20) * (9.0 / 8.0),
	})

	// Top
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          1,
		Y:          0,
		Matrix:     TRANSFORM_TOP_BOTTOM_MATRIX,
		TranslateX: float32(size) * (-40.0 / 256.0),
		TranslateY: float32(size) * (83.0 / 256.0),
		Flip:       false,
		Scale:      float32(size / 20),
	})

	// Front
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          1,
		Y:          1,
		Matrix:     TRANSFORM_FRONT_BACK_MATRIX,
		TranslateX: float32(size) * (132.5 / 256.0),
		TranslateY: float32(size) * (177.5 / 256.0),
		Flip:       false,
		Scale:      float32(size / 20),
	})

	// Right
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          2,
		Y:          1,
		Matrix:     TRANSFORM_RIGHT_LEFT_MATRIX,
		TranslateX: float32(size) * (121.0 / 256.0),
		TranslateY: float32(size) * (52.0 / 256.0),
		Flip:       true,
		Scale:      float32(size / 20),
	})

	// Front overlay
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          5,
		Y:          1,
		Matrix:     TRANSFORM_FRONT_BACK_MATRIX,
		TranslateX: float32(size) * (132.5 / 256.0) * (8.1 / 8.0),
		TranslateY: float32(size) * (177.5 / 256.0),
		Flip:       false,
		Scale:      float32(size/20) * (9.0 / 8.0),
	})

	// Right overlay
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          4,
		Y:          1,
		Matrix:     TRANSFORM_RIGHT_LEFT_MATRIX,
		TranslateX: float32(size) * (26.0 / 256.0) * (8.0 / 8.1),
		TranslateY: float32(size) * (52.0 / 256.0),
		Flip:       false,
		Scale:      float32(size/20) * (9.0 / 8.0),
	})

	// Top overlay
	overlay3DSection(out, img, &OverlaySectionOptions{
		Size:       size,
		X:          5,
		Y:          0,
		Matrix:     TRANSFORM_TOP_BOTTOM_MATRIX,
		TranslateX: float32(size) * (-40.0 / 256.0) * (8.0 / 8.1),
		TranslateY: float32(size) * (83.0 / 256.0) * (8.0 / 9.0),
		Flip:       false,
		Scale:      float32(size/20) * (9.0 / 8.0),
	})

	return out
}

func RenderHead(textureId string) []byte {
	CACHE_PATH := filepath.Join(CACHE_DIR, textureId+".png")
	if data, err := os.ReadFile(CACHE_PATH); err == nil {
		return data
	}

	response, err := http.Get("https://textures.minecraft.net/texture/" + textureId)
	if err != nil {
		log.Println("Error fetching texture:", err)
		return nil
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Println("Error fetching texture, status code:", response.StatusCode)
		return nil
	}

	img, err := png.Decode(response.Body)
	if err != nil {
		log.Println("Error decoding texture:", err)
		return nil
	}

	var rgbaImg *image.RGBA
	switch v := img.(type) {
	case *image.RGBA:
		rgbaImg = v
	default:
		bounds := img.Bounds()
		rgbaImg = image.NewRGBA(bounds)
		draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)
	}

	head := To3DHead(rgbaImg)

	var buf bytes.Buffer
	if err := png.Encode(&buf, head); err != nil {
		log.Println("Error encoding head to PNG:", err)
		return nil
	}
	data := buf.Bytes()

	if err := os.MkdirAll(CACHE_DIR, 0755); err == nil {
		if err := os.WriteFile(CACHE_PATH, data, 0644); err != nil {
			log.Println("Error saving head to cache:", err)
		}
	} else {
		log.Println("Error creating cache directory:", err)
	}

	return data
}

func RenderItem(itemID string) ([]byte, error) {
	damage := 0
	if strings.Contains(itemID, ":") {
		itemID = strings.Split(itemID, ":")[0]
		parsedDmg, err := utility.ParseInt(strings.Split(itemID, ":")[1])
		if err == nil {
			damage = parsedDmg
		}
	}

	itemData := constants.ITEMS[itemID]
	TextureItem := models.TextureItem{
		Damage: &itemData.Damage,
		ID:     &itemData.ItemId,
		Tag: models.TextureItemExtraAttributes{
			ExtraAttributes: map[string]interface{}{
				"id": itemData.SkyblockID,
			},
		},
	}

	if damage != 0 {
		TextureItem.Damage = &damage
	}

	output := GetTexture(TextureItem)
	if output == "" {
		return nil, fmt.Errorf("couldn't find the texture")
	}

	// If it's a skull, use RenderHead and return its PNG bytes
	if strings.Contains(output, "/Vanilla/assets/firmskyblock/models/item/skull") {
		if itemData.Texture == "" {
			return nil, fmt.Errorf("no texture id for skull")
		}
		data := RenderHead(itemData.Texture)
		if data == nil {
			return nil, fmt.Errorf("failed to render head for skull")
		}
		return data, nil
	}

	// If output is a localhost asset, read from disk (performance optimization)
	if strings.HasPrefix(output, "http://localhost") || strings.HasPrefix(output, "https://localhost") {
		assetsIdx := strings.Index(output, "/assets/")
		if assetsIdx != -1 {
			localPath := output[assetsIdx+1:] // skip the leading slash
			if _, err := os.Stat(localPath); err == nil {
				data, err := os.ReadFile(localPath)
				if err != nil {
					return nil, fmt.Errorf("error reading local asset: %v", err)
				}
				return data, nil
			} else {
				return nil, fmt.Errorf("local asset not found: %s", localPath)
			}
		}
		return nil, fmt.Errorf("invalid localhost asset path: %s", output)
	}

	response, err := http.Get(output)
	if err != nil {
		return nil, fmt.Errorf("error fetching item texture: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching item texture: %v", err)
	}

	img, err := png.Decode(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error decoding item texture: %v", err)
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("error encoding item texture: %v", err)
	}

	return buf.Bytes(), nil
}
