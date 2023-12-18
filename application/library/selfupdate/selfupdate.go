package selfupdate

import (
	"crypto"
	"io"
	"os"
	"reflect"
	"syscall"

	"github.com/admpub/service"
	"github.com/fynelabs/selfupdate"
	"github.com/webx-top/echo"
)

func Update(r io.Reader, targetPath string, opts ...func(o *selfupdate.Options)) error {
	o := selfupdate.Options{
		TargetPath: targetPath,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return selfupdate.Apply(r, o)
}

func Restart(exiter func(error), executable string) error {
	var args []string
	isInteractive := service.Interactive()
	if isInteractive { // 交互模式
		args = []string{os.Args[0], `forkself`}
		if len(os.Args) > 1 {
			args = append(args, os.Args[1:]...)
		}
	} else { //服务模式
		args = []string{executable, `forkself`, `service`, `restart`}
	}
	procAttr := &syscall.SysProcAttr{}
	r := reflect.ValueOf(procAttr)
	v := reflect.Indirect(r)
	if f := v.FieldByName(`Pdeathsig`); f.IsValid() {
		f.Set(reflect.ValueOf(syscall.Signal(0)))
	}
	if f := v.FieldByName(`Setpgid`); f.IsValid() {
		f.SetBool(true)
	}
	_, err := os.StartProcess(executable, args, &os.ProcAttr{
		Dir:   echo.Wd(),
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   procAttr,
	})
	if exiter != nil {
		exiter(err)
	} else if err == nil {
		os.Exit(0)
	}
	return err
}

// Create TargetPath replacement with this file mode. If zero, defaults to 0755.
func TargetMode(mode os.FileMode) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.TargetMode = mode
	}
}

// Checksum of the new binary to verify against. If nil, no checksum or signature verification is done.
func Checksum(checksum []byte) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.Checksum = checksum
	}
}

// Public key to use for signature verification. If nil, no signature verification is done.
func PublicKey(publicKey crypto.PublicKey) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.PublicKey = publicKey
	}
}

// Signature to verify the updated file. If nil, no signature verification is done.
func Signature(signature []byte) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.Signature = signature
	}
}

// Pluggable signature verification algorithm. If nil, ECDSA is used.
func Verifier(verifier selfupdate.Verifier) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.Verifier = verifier
	}
}

// Use this hash function to generate the checksum. If not set, SHA256 is used.
func Hash(hash crypto.Hash) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.Hash = hash
	}
}

// If nil, treat the update as a complete replacement for the contents of the file at TargetPath.
// If non-nil, treat the update contents as a patch and use this object to apply the patch.
func Patcher(patcher selfupdate.Patcher) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.Patcher = patcher
	}
}

// Store the old executable file at this path after a successful update.
// The empty string means the old executable file will be removed after the update.
func OldSavePath(oldSavePath string) func(o *selfupdate.Options) {
	return func(o *selfupdate.Options) {
		o.OldSavePath = oldSavePath
	}
}
