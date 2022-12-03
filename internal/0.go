// Created by EldersJavas(EldersJavas&gmail.com)

package internal

import (
	"embed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io"
	"io/fs"
)

//go:embed res
var resourceRootFS embed.FS

func init() {
	f, err := fs.Sub(resourceRootFS, "res")
	if err != nil {
		panic(err)
	}
	ResourceFS = f
}

var ResourceFS fs.FS

var (
	SpaceAgeBig   font.Face
	SpaceAgeSmall font.Face
	DebugFace     font.Face
)

func init() {
	f, err := ResourceFS.Open("fusion-pixel.otf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	font, err := opentype.Parse(bs)
	if err != nil {
		panic(err)
	}
	{
		face, err := opentype.NewFace(font, &opentype.FaceOptions{
			Size: 50,
			DPI:  72,
		})
		if err != nil {
			panic(err)
		}
		DebugFace = face
	}
	{
		face, err := opentype.NewFace(font, &opentype.FaceOptions{
			Size: 144,
			DPI:  72,
		})
		if err != nil {
			panic(err)
		}
		SpaceAgeBig = face
	}
	{
		face, err := opentype.NewFace(font, &opentype.FaceOptions{
			Size: 108,
			DPI:  72,
		})
		if err != nil {
			panic(err)
		}
		SpaceAgeSmall = face
	}
}
