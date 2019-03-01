package filecp

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sanguohot/filecp/pkg/common/file"
	"github.com/sanguohot/filecp/pkg/common/log"
	"path/filepath"
)

type FilecpRow struct {
	src string
	dst string
} 

type Filecp struct {
	src string
	dst string
	csv string
}

func New(src, dst, csv string) *Filecp {
	return &Filecp{src: src, dst: dst, csv: csv}
}

func (cp *Filecp) shouldNotBeDir(path string) error {
	if isDir, err := file.FileIsDir(path); err != nil {
		return err
	} else if isDir {
		return fmt.Errorf("%s could not be diretory", path)
	}
	return nil
}

func (cp *Filecp) CopyOne(src, dst string) error {
	if !file.FilePathExist(src){
		return fmt.Errorf("%s not found", src)
	}
	if err := cp.shouldNotBeDir(src); err != nil {
		return err
	}
	if file.FilePathExist(dst) {
		return nil
	}

	if err := file.EnsureDir(filepath.Dir(dst)); err != nil {
		return err
	}
	_, err := file.StandardCopy(src, dst)
	if err != nil {
		return err
	}
	log.Sugar.Infof("copy %s => %s done", src, dst)
	return nil
}

func (cp *Filecp) CopyMul(cps []FilecpRow) error {
	for _, v := range cps {
		if err := cp.CopyOne(v.src, v.dst); err != nil {
			return err
		}
	}
	return nil
}

func (cp *Filecp) Copy() error {
	if cp.dst == "" || cp.src == "" {
		return errors.New("src or dst should not be empty string")
	}
	if cp.csv != "" {
		return errors.New("csv not supported now")
	}
	cps := []FilecpRow{FilecpRow{src: cp.src, dst: cp.dst}}
	return cp.CopyMul(cps)
}