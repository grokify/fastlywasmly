package tarutil

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	fsterr "github.com/fastly/cli/pkg/errors"
	"github.com/fastly/cli/pkg/manifest"
	"github.com/grokify/mogo/archive/tarutil"
	"github.com/grokify/mogo/os/osutil"
	"github.com/kennygrant/sanitize"
)

const (
	WASMBinDir   = "bin"
	WASMFilename = "main.wasm"
	TOMLFilename = "fastly.toml"
	TARGzExt     = ".tar.gz"
)

var (
	ErrDirNotExists = errors.New("dir does not exist")
	ErrFileZeroSize = errors.New("file is zero size")
)

// BuildEdgePackage creates a Fastly Compute@Edge tarball. A TOML file must be
// supplied. Either the `wasmfile` or `bindir` should be supplied. Supplying `bindir`
// will recursively copy all contents in to the `<your-package>/bin` dir in the tar file.
func BuildEdgePackage(tomlfile, wasmfile, bindir string) (string, error) {
	packageNameSanitized, files, err := buildEdgePackageFiles(tomlfile, wasmfile, bindir)
	if err != nil {
		return "", err
	}
	outfile := packageNameSanitized + TARGzExt
	return outfile, tarutil.CreateArchiveGzipFile(outfile, files)
}

func buildEdgePackageFiles(tomlfile, wasmfile, bindir string) (string, map[string]string, error) {
	files := map[string]string{}
	fileExists, err := osutil.IsFile(tomlfile, true)
	if err != nil {
		return "", files, err
	}
	if !fileExists {
		return "", files, ErrFileZeroSize
	}
	packageNameSanitized, err := ManifestPackageNameSanitized(tomlfile)
	if err != nil {
		return packageNameSanitized, files, err
	}
	files[tomlfile] = filepath.Join(packageNameSanitized, TOMLFilename)

	if len(bindir) > 0 {
		dirExists, err := osutil.IsDir(bindir)
		if err != nil {
			return packageNameSanitized, files, err
		}
		if !dirExists {
			return packageNameSanitized, files, ErrDirNotExists
		}
		err = osutil.VisitFiles(bindir, func(dir string, info fs.FileInfo) error {
			mode := info.Mode()
			if !mode.IsDir() {
				subdir := strings.Trim(dir[len(bindir):], string(os.PathSeparator))
				if len(subdir) == 0 {
					files[filepath.Join(dir, info.Name())] = filepath.Join(packageNameSanitized, WASMBinDir, info.Name())
				} else {
					files[filepath.Join(dir, info.Name())] = filepath.Join(packageNameSanitized, WASMBinDir, subdir, info.Name())
				}
			}
			return nil
		})
		if err != nil {
			return packageNameSanitized, files, err
		}
	}

	if len(wasmfile) > 0 {
		fileExists, err := osutil.IsFile(wasmfile, true)
		if err != nil {
			return packageNameSanitized, files, err
		}
		if !fileExists {
			return packageNameSanitized, files, ErrFileZeroSize
		}
		files[wasmfile] = filepath.Join(packageNameSanitized, WASMBinDir, WASMFilename)
	}

	return packageNameSanitized, files, nil
}

func ManifestPackageNameSanitized(filename string) (string, error) {
	var mfile manifest.File
	logEntries := &fsterr.LogEntries{}
	mfile.SetErrLog(logEntries)
	err := mfile.Read(filename)
	if err != nil {
		return "", err
	}
	if len(*logEntries) > 0 {
		entries := []fsterr.LogEntry(*logEntries)
		return "", entries[0].Err
	}
	return sanitize.BaseName(mfile.Name), nil
}
