package cloudbackup

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/admpub/log"
	"github.com/admpub/nging/v5/application/dbschema"
	"github.com/admpub/nging/v5/application/model"
	"github.com/webx-top/com"
)

func New(mgr Storager, cfg dbschema.NgingCloudBackup) *Cloudbackup {
	return &Cloudbackup{mgr: mgr, cfg: cfg}
}

type Cloudbackup struct {
	mgr        Storager
	cfg        dbschema.NgingCloudBackup
	SourcePath string
	DestPath   string
	Filter     func(string) bool

	WaitFillCompleted bool
	IgnoreWaitRegexp  *regexp.Regexp
}

func (c *Cloudbackup) OnCreate(file string) {
	if !c.Filter(file) {
		return
	}
	fi, err := os.Stat(file)
	if err != nil {
		log.Error(file + `: ` + err.Error())
		return
	}
	if fi.IsDir() {
		err = filepath.Walk(file, func(ppath string, info os.FileInfo, werr error) error {
			if werr != nil {
				return werr
			}
			if info.IsDir() || !c.Filter(ppath) {
				return nil
			}
			_waitFillCompleted := c.WaitFillCompleted
			if _waitFillCompleted && c.IgnoreWaitRegexp != nil {
				_waitFillCompleted = c.IgnoreWaitRegexp.MatchString(ppath)
			}
			objectName := path.Join(c.DestPath, strings.TrimPrefix(ppath, c.SourcePath))
			FileChan() <- &PutFile{
				Manager:           c.mgr,
				Config:            c.cfg,
				ObjectName:        objectName,
				FilePath:          ppath,
				Operation:         model.CloudBackupOperationCreate,
				WaitFillCompleted: _waitFillCompleted,
			}
			return nil
		})
	} else {
		_waitFillCompleted := c.WaitFillCompleted
		if _waitFillCompleted && c.IgnoreWaitRegexp != nil {
			_waitFillCompleted = c.IgnoreWaitRegexp.MatchString(file)
		}
		objectName := path.Join(c.DestPath, strings.TrimPrefix(file, c.SourcePath))
		FileChan() <- &PutFile{
			Manager:           c.mgr,
			Config:            c.cfg,
			ObjectName:        objectName,
			FilePath:          file,
			Operation:         model.CloudBackupOperationCreate,
			WaitFillCompleted: _waitFillCompleted,
		}
	}
	if err != nil {
		log.Errorf(`cloudbackup onCreate: %v`, err)
	}
}

func (c *Cloudbackup) OnModify(file string) {
	if !c.Filter(file) {
		return
	}
	objectName := path.Join(c.DestPath, strings.TrimPrefix(file, c.SourcePath))
	fi, err := os.Stat(file)
	if err != nil {
		log.Error(file + `: ` + err.Error())
		return
	}
	if fi.IsDir() {
		return
	}
	_waitFillCompleted := c.WaitFillCompleted
	if _waitFillCompleted && c.IgnoreWaitRegexp != nil {
		_waitFillCompleted = c.IgnoreWaitRegexp.MatchString(file)
	}
	FileChan() <- &PutFile{
		Manager:           c.mgr,
		Config:            c.cfg,
		ObjectName:        objectName,
		FilePath:          file,
		Operation:         model.CloudBackupOperationUpdate,
		WaitFillCompleted: _waitFillCompleted,
	}
}

func (c *Cloudbackup) OnDelete(file string) {
	if !c.Filter(file) {
		return
	}
	relPath := strings.TrimPrefix(file, c.SourcePath)
	if len(relPath) == 0 || relPath == `/` || relPath == `\` {
		return
	}
	startTime := time.Now()
	objectName := path.Join(c.DestPath, relPath)
	err := c.mgr.RemoveDir(context.Background(), objectName)
	if err != nil {
		log.Error(file + `: ` + err.Error())
	}
	err = c.mgr.Remove(context.Background(), objectName)
	if err != nil {
		log.Error(file + `: ` + err.Error())
	}
	RecordLog(nil, err, &c.cfg, file, objectName, model.CloudBackupOperationDelete, startTime, 0)
	if db, err := LevelDB().OpenDB(c.cfg.Id); err == nil {
		db.Delete(com.Str2bytes(file), nil)
	}
}

func (c *Cloudbackup) OnRename(file string) {
	if !c.Filter(file) {
		return
	}
	relPath := strings.TrimPrefix(file, c.SourcePath)
	if len(relPath) == 0 || relPath == `/` || relPath == `\` {
		return
	}
	startTime := time.Now()
	objectName := path.Join(c.DestPath, relPath)
	err := c.mgr.RemoveDir(context.Background(), objectName)
	if err != nil {
		log.Error(file + `: ` + err.Error())
	}
	err = c.mgr.Remove(context.Background(), objectName)
	if err != nil {
		log.Error(file + `: ` + err.Error())
	}
	RecordLog(nil, err, &c.cfg, file, objectName, model.CloudBackupOperationDelete, startTime, 0)
	if db, err := LevelDB().OpenDB(c.cfg.Id); err == nil {
		db.Delete(com.Str2bytes(file), nil)
	}
}
