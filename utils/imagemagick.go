// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package utils

//图片处理
import "os/exec"
import "github.com/revel/revel"

//ImageMagick之图片缩放
//http://www.netingcn.com/imagemagick-resize.html
//http://www.imagemagick.org/script/command-line-options.php#thumbnail
//参数说明: old_img:原始图片 new_img:新图片 img_size:图片大小 output_align:位置center background:颜色
func Resize(old_img string, new_img string, img_size string, output_align string, background string) bool {
	if IsFile(old_img) {
		//convert -thumbnail 200×100 src.jpg dest.jpg

		//得到图片宽为200，高根据原始图片比例计算而来
		//convert -thumbnail 200 src.jpg dest.jpg

		//得到的图片高位100，宽根据原始图片比例计算而来
		//convert -thumbnail x100 src.jpg dest.jpg

		//固定宽高缩放。即不考虑原是图宽高的比例，把图片缩放到指定大小。
		//convert -thumbnail 200x100! src.jpg dest.jpg
		cmd := exec.Command("convert", "-thumbnail", img_size, "-background", background, "-gravity", output_align, "-extent", img_size, old_img, new_img)
		err := cmd.Run()
		if err != nil {
			revel.WARN.Println(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

//生成圆角图片
//img_size:0x4 background: none
func Vignette(old_img string, new_img string, img_size string) bool {
	if IsFile(old_img) {
		// convert 2014040120333336587.jpg -vignette 0x4 -quality 100 vignette.jpg
		cmd := exec.Command("convert", old_img, "-vignette", img_size, "-quality", "100", new_img)
		err := cmd.Run()
		if err != nil {
			revel.WARN.Println(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

//imagemagick图片反色处理
//http://www.netingcn.com/imagemagick-negate.html
func Negate(old_img string, new_img string) bool {
	if IsFile(old_img) {
		// convert -negate src.jpg negate.jpg
		cmd := exec.Command("convert", "-negate", old_img, new_img)
		err := cmd.Run()
		if err != nil {
			revel.WARN.Println(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

//ImageMagick之图片裁剪
//http://www.netingcn.com/imagemagick-crop.html
func Crop(old_img string, new_img string, img_size string) bool {
	if IsFile(old_img) {
		//convert 原始图片 -crop widthxheight+x+y 目标图片
		//其中widthxheight是目标图片的尺寸，+x+y是原始图片的坐标点，这两组值至少要出现一组，也可以同时存在。
		cmd := exec.Command("convert", old_img, "-crop", img_size, new_img)
		err := cmd.Run()
		if err != nil {
			revel.WARN.Println(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

//ImageMagicK给图片加水印之图片水印处理
//http://www.netingcn.com/imagemagick-mark.html
func WatermarkLogo(old_img string, new_img string, logo_img string) bool {
	if IsFile(old_img) {
		//convert src.jpg logo.gif -gravity southeast -geometry +5+10 -composite dest.jpg
		//假设把名为logo.gif的水印图标添加在原始图片（src.jpg）右下角，且水印的下边缘距原始图片10像素、右边缘距原始图片5像素。

		cmd := exec.Command("convert", old_img, logo_img, "-gravity", "southeast", "-geometry", "+5+10", "-composite", new_img)
		err := cmd.Run()
		if err != nil {
			revel.WARN.Println(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

//ImageMagicK给图片加水印之文字水印处理
func WatermarkText(old_img string, new_img string) bool {
	if IsFile(old_img) {
		//www.netingcn.com作为水印加上图片上
		//convert src.jpg -gravity southeast -fill black -pointsize 16 -draw "text 5,5 'http://www.netingcn.com'" dest-c.jpg
		cmd := exec.Command("convert", old_img, "-gravity", "southeast", "-fill", "black", "-pointsize", "16", "-draw", "text 5,5 'http://www.6574.com.cn'", new_img)
		err := cmd.Run()
		if err != nil {
			revel.WARN.Println(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}
