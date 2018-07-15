package main

import (
  "image"
  "image/jpeg"
  // _ "image/png"
  // _ "image/gif"
  "os"
  "log"
  "github.com/disintegration/imaging"
)

func main () {
  src := `./kaeru.jpg`
  dst := `./kaeru_dst.jpg`
  dst2 := `./kaeru_dst2.jpg`
  // dst2 := `./kaeru_dst2.jpg`

  // img, _, err := openImage(src)
  // checkTopError(err)

  config, _, err := decodeConfig(src)
  checkTopError(err)

  img, err := imaging.Open(src)
  checkTopError(err)

  quality := 70

  // 左上切り出し
  x, y, _x, _y := 0, 0, config.Width/2, config.Height/2
  cropImg := imaging.Crop(img, image.Rect(x, y, _x, _y))
  SaveJpeg(dst, quality, cropImg)

  // 右下切り出し
  x, y, _x, _y = config.Width/2, config.Height/2, config.Width, config.Height
  cropImg = imaging.Crop(img, image.Rect(x, y, _x, _y))
  SaveJpeg(dst2, quality, cropImg)

  // err = imaging.Save(cropAnchorImg, dst2)
  // checkTopError(err)
}


/*
 * ファイルからimgを取得
 */
func openImage (src string) (img image.Image, format string, err error) {
    // get file
  file, err := os.Open(src)
  if err != nil { return }
  defer file.Close() // latest close

  // Decode
  img, format, err = image.Decode(file)
  if err != nil { return }

  return
}

/**
 * 最上位でのエラーチェック(強制終了)
 * @param  {[type]} err error         [description]
 * @return {[type]}     [description]
 */
func checkTopError (err error) {
  if (err != nil) {
    log.Println(err)
    os.Exit(1)
  }
}

/**
 * Decodeは同じストリーム内で行うとエラーになるため関数に分ける
 */
func decodeConfig(src string) (config image.Config, format string, err error) {
  file, err := os.Open(src)
  if err != nil { return }
  defer file.Close() // latest close

  config, format, err = image.DecodeConfig(file)
  if err != nil { return }

  return
}

func SaveJpeg(dst string, quality int, img image.Image) (err error) {
  // create file
  out, err := os.Create(dst)
  if err != nil { return }
  defer out.Close()

  // save
  opts := &jpeg.Options{Quality: quality}
  jpeg.Encode(out, img, opts)
  return
}