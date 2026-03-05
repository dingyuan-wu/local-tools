package tinypng

import (
	_ "image/gif"
	_ "image/png"

	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

// compressImage 压缩图片
func compressImage(inputPath string, outputPath string, quality int) error {
	// 打开输入文件
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开文件 %s: %v", inputPath, err)
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("无法解码图片 %s: %v", inputPath, err)
	}

	// 创建输出目录（如果不存在）
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return fmt.Errorf("无法创建输出目录: %v", err)
	}

	// 打开输出文件
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建文件 %s: %v", outputPath, err)
	}
	defer outFile.Close()

	// 压缩并保存图片
	err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return fmt.Errorf("无法压缩图片: %v", err)
	}

	return nil
}

func isSupportFormat(ext string) bool {
	supportFormatList := []string{".jpg", ".jpeg", ".png", ".gif"}
	for _, v := range supportFormatList {
		if strings.ToLower(ext) == v {
			return true
		}
	}
	fmt.Println("不支持的图片格式: ", ext)
	return false
}

// TinyImagesInDir 压缩目录中的所有图片
func TinyImagesInDir(inputDir, outputDir string, quality int) error {
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// 仅处理支持的图片格式
		ext := filepath.Ext(info.Name())
		if !isSupportFormat(ext) {
			return nil
		}

		// 设置输出路径
		relPath, _ := filepath.Rel(inputDir, path)
		outputPath := filepath.Join(outputDir, relPath)
		outputPath = outputPath[:len(outputPath)-len(ext)] + "-tiny.jpg" // 保存为 JPG 格式

		fmt.Printf("正在压缩: %s -> %s\n", path, outputPath)
		return compressImage(path, outputPath, quality)
	})

	return err
}
