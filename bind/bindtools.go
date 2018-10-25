package bind

import (
    "os"
    "path/filepath"
    "strings"
)

var extractTemp bool = false
var extractCurr bool = false
var tempdir = os.TempDir()

func ExtractCurr(){
    if extractCurr == false {
        todir := filepath.Join(GetCurrentDirectory(),"rsrc")
        err := RestoreAssets(todir, "assets")
        if err != nil {
            logf.Fatal(err)
        }
        extractCurr = true
    }
}

func GetTempFilePath(name string) (fp string) {
    if extractTemp == false {
        err := RestoreAssets(tempdir, "assets")
        if err != nil {
            logf.Fatal(err)
        }
        extractTemp = true
    }
    fp, err := filepath.Abs(filepath.Join(tempdir, name))
    if err != nil {
        log.Fatal(err)
    }
    log.Info("LOCATE FILE:" + fp)
    return
}

// walk is not support []byte as input
// so writ the file to windows system temp dir
func WriteTempFile(name string, bs []byte) (f *os.File, err error) {
    dir := os.TempDir()
    fullpath := filepath.Join(dir, name)
    //golang默认是不覆盖的 必须加os.O_TRUNC
    f, err = os.OpenFile(fullpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
    return
}

func GetCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
    if err != nil {
        log.Fatal(err)
    }
    return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
