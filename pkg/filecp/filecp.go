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
	srcs []string
	dsts []string
	csv string
}

func New(srcs, dsts []string, csv string) *Filecp {
	return &Filecp{srcs: srcs, dsts: dsts, csv: csv}
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
	if dst == "" || src == "" {
		return errors.New("src or dst should not be empty string")
	}
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
	return nil
}

func (cp *Filecp) CopyMul(cps []FilecpRow) error {
	for _, v := range cps {
		if err := cp.CopyOne(v.src, v.dst); err != nil {
			return err
		}
		log.Sugar.Infof("copy %s => %s done", v.src, v.dst)
	}
	return nil
}

func (cp *Filecp) Copy() error {
	if len(cp.srcs) <= 0 || len(cp.dsts) <= 0 || len(cp.srcs) != len(cp.dsts) {
		return errors.New("param 'srcs' or 'dsts' invalid")
	}
	if cp.csv != "" {
		return errors.New("csv not supported now")
	}
	rows := make([]FilecpRow, len(cp.srcs))
	for k, v := range cp.srcs {
		row := FilecpRow{src: v, dst: cp.dsts[k]}
		rows[k] = row
	}
	return cp.CopyMul(rows)
}