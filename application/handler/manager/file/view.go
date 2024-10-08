package file

import (
	"bytes"
	"io"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/webx-top/echo"

	"github.com/coscms/webcore/cmd/bootconfig"
	uploadLibrary "github.com/coscms/webcore/library/upload"
	"github.com/coscms/webcore/registry/upload"
	"github.com/coscms/webcore/registry/upload/convert"
	"github.com/coscms/webcore/registry/upload/driver/local"
)

var fileGeneratorLock = sync.RWMutex{}

func File(ctx echo.Context) error {
	subdir := ctx.Param(`subdir`)
	file := ctx.Param(`*`)
	parts := strings.SplitN(file, `/`, 2)
	if upload.AllowedSubdirx(subdir, parts[0]) {
		subdir += `/` + parts[0]
		if len(parts) > 1 {
			file = parts[1]
		} else {
			file = ``
		}
	}
	file = filepath.Join(uploadLibrary.UploadDir, subdir, file)
	maxAge := bootconfig.HTTPCacheMaxAge
	if convert.CountConverters() < 1 {
		return ctx.CacheableFile(file, maxAge)
	}
	var (
		convertFunc  convert.Convert
		ok           bool
		originalFile string
	)
	extension := ctx.Query(`ex`)
	if len(extension) > 0 {
		if extension[0] != '.' {
			extension = `.` + extension
		}
		convertFunc, ok = convert.GetConverter(extension)
		if !ok {
			return ctx.CacheableFile(file, maxAge)
		}
		originalFile = file
	} else {
		originalExtension := filepath.Ext(file) // file.jpg.png
		extension = strings.ToLower(originalExtension)
		convertFunc, ok = convert.GetConverter(extension)
		if !ok {
			return ctx.CacheableFile(file, maxAge)
		}
		originalFile = strings.TrimSuffix(file, originalExtension)
		index := strings.LastIndex(originalFile, `.`)
		// 单扩展名或相同扩展名的情况下不转换格式
		if index < 0 || strings.ToLower(originalFile[index:]) == extension {
			return ctx.CacheableFile(originalFile, maxAge)
		}
	}
	supported := strings.Contains(ctx.Header(echo.HeaderAccept), "image/"+strings.TrimPrefix(extension, `.`))
	if !supported {
		return ctx.CacheableFile(originalFile, maxAge)
	}

	if err := ctx.CacheableFile(file, maxAge); err != echo.ErrNotFound {
		return err
	}

	fileGeneratorLock.Lock()
	defer fileGeneratorLock.Unlock()

	return ctx.ServeCallbackContent(func(_ echo.Context) (io.Reader, error) {
		storerName := local.Name
		newStore := upload.StorerGet(storerName)
		if newStore == nil {
			return nil, ctx.E(`存储引擎“%s”未被登记`, storerName)
		}
		storer, err := newStore(ctx, subdir)
		if err != nil {
			return nil, err
		}
		f, err := storer.Get(`/` + originalFile)
		if err != nil {
			return nil, echo.ErrNotFound
		}
		defer f.Close()
		buf, err := convertFunc(f, 70)
		if err != nil {
			return nil, err
		}
		b := buf.Bytes()
		saveFile := storer.URLToFile(`/` + file)
		_, _, err = storer.Put(saveFile, buf, int64(len(b)))
		return bytes.NewBuffer(b), err
	}, path.Base(file), time.Now())
}
