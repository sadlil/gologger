package gologger

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var activeLogger map[string]*Logger

// This can be set by a build script. It will be the colon separated equivalent
// of the environment variable.
var gopath string

// This is the processed version based on either the above variable set by the
// build or from the GOPATH environment variable.
var gopaths []string

func (t LoggerType) String() string {
	return string(t)
}

func init() {
	activeLogger = make(map[string]*Logger)

	// prefer the variable set at build time, otherwise fallback to the
	// environment variable.
	if gopath == "" {
		gopath = os.Getenv("GOPATH")
	}

	for _, p := range strings.Split(gopath, ":") {
		if p != "" {
			gopaths = append(gopaths, filepath.Join(p, "src")+"/")
		}
	}

	// Also strip GOROOT for maximum cleanliness
	gopaths = append(gopaths, filepath.Join(runtime.GOROOT(), "src", "pkg")+"/")
}

// StripGOPATH strips the GOPATH prefix from the file path f.
// In development, this will be done using the GOPATH environment variable.
// For production builds, where the GOPATH environment will not be set, the
// GOPATH can be included in the binary by passing ldflags, for example:
//
//     GO_LDFLAGS="$GO_LDFLAGS -X github.com/facebookgo/stack.gopath $GOPATH"
//     go install "-ldflags=$GO_LDFLAGS" my/pkg
func stripGOPATH(f string) string {
	for _, p := range gopaths {
		if strings.HasPrefix(f, p) {
			return f[len(p):]
		}
	}
	return f
}

// StripPackage strips the package name from the given Func.Name.
func stripPackage(n string) string {
	slashI := strings.LastIndex(n, "/")
	if slashI == -1 {
		slashI = 0 // for built-in packages
	}
	dotI := strings.Index(n[slashI:], ".")
	if dotI == -1 {
		return n
	}
	return n[slashI+dotI+1:]
}

func stripFileName(n string) string {
	parts := strings.Split(n, ".")
	pl := len(parts)
	packageName := ""
	if parts[pl-2][0] == '(' {
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	if strings.Contains(packageName, "/") {
		packageName = packageName[strings.LastIndex(packageName, "/") + 1:]
	}
	return packageName
}

func getFileName(n string) string {
	slashI := strings.LastIndex(n, "/")
	if slashI == -1 {
		slashI = 0 // for built-in packages
	}
	return n[slashI+1:]
}