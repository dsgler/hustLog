package newgetcode

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
)

var Sxs = [4]int{0, 22, 43, 65}
var Sy = 17
var Sw = 17
var Sh = 20

// 主函数，传入原gif图片，返回4位字符串
func MergeAndGet(rawGifByte []byte) (code string, err error) {
	img, err := MergeGifF(rawGifByte)
	if err != nil {
		return
	}

	ret := GetFour(img)
	code = fmt.Sprintf("%d%d%d%d", ret[0], ret[1], ret[2], ret[3])
	log.Println("code:" + code)
	return
}

func MergeGifF(rawGifByte []byte) (ret *image.Gray, err error) {
	// file, err := os.Open("/home/dsgler/goproject/src/hustLog/GetCode/code.gif")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// gifImage, err := gif.DecodeAll(file)

	// 解码GIF
	gifImage, err := gif.DecodeAll(bytes.NewReader(rawGifByte))
	if err != nil {
		return
	}

	// 获取第一帧的尺寸
	bounds := gifImage.Image[0].Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	cnt := make([][]int, width)
	for i := range cnt {
		cnt[i] = make([]int, height)
	}

	// 创建新的灰度图像
	ResultImage := image.NewGray(bounds)

	// 遍历每一帧
	for _, img := range gifImage.Image {
		// 遍历每个像素
		for y := 17; y < 17+20; y++ {
			for x := 0; x < width; x++ {
				// 获取当前像素
				oldPixel := img.At(x, y)

				// 转换为灰度值
				r, g, b, _ := oldPixel.RGBA()
				gray := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))

				// // 获取当前累积的灰度值
				// currentGray := ResultImage.GrayAt(x, y).Y

				// // 叠加灰度值（取平均）
				// newGray := uint8(max(uint32(currentGray)+uint32(gray), 255))

				// ResultImage.Set(x, y, color.Gray{Y: newGray})

				if gray < 254 {
					cnt[x][y]++
				}
			}
		}
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if cnt[x][y] >= 3 {
				ResultImage.Set(x, y, color.Gray{Y: 255})
			}
		}
	}

	// 保存结果
	// outFile, err := os.Create("output4.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer outFile.Close()

	// // 将结果写入文件
	// if err := png.Encode(outFile, ResultImage); err != nil {
	// 	log.Fatal(err)
	// }
	ret = ResultImage
	return
}

// 通过与已保存的对比判断，目前精度不高
func GetFour(img *image.Gray) []int {
	result := make([]int, 0, 4)
	for k := range Sxs {
		maxScore := 0
		value := -1
		for n := 0; n <= 9; n++ {
			myScore := 0
			x := Sxs[k]
			y := Sy

			rst := make([]bool, 0, 340)

			for i := 0; i < Sw; i++ {
				for j := 0; j < Sh; j++ {
					// fmt.Print(fmt.Sprint(x+i) + "," + fmt.Sprint(y+j) + " ")
					tmp := img.GrayAt(x+i, y+j).Y
					if tmp == 0 {
						rst = append(rst, false)
					} else {
						rst = append(rst, true)
					}
				}
			}

			for i := 0; i < 340; i++ {
				if rst[i] == Code[n][i] {
					myScore++
				}
			}

			if myScore > maxScore {
				value = n
				maxScore = myScore
			}
		}
		result = append(result, value)
	}
	return result
}
