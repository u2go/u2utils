package cmder

import (
	"github.com/pkg/errors"
	"io"
	"os/exec"
	"sync"
)

// OutHandle 参数：单次的输出，全部输出，输入流
type OutHandle func(outBuff []byte, allOutBytes []byte, inPipe io.WriteCloser) error

// Run 执行命令，可以通过输出判断输入内容，实现命令交互
func Run(name string, arg []string, outHandle, errHandle OutHandle) error {
	cmd := exec.Command(name, arg...)
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer inPipe.Close()
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer outPipe.Close()
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer errPipe.Close()
	err = cmd.Start()
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	var outErr error
	var errErr error
	go func() {
		defer wg.Done()
		outErr = handleOut(outPipe, inPipe, outHandle)
	}()
	go func() {
		defer wg.Done()
		errErr = handleOut(errPipe, inPipe, errHandle)
	}()
	wg.Wait()
	if outErr != nil {
		return outErr
	}
	if errErr != nil {
		return errErr
	}
	return cmd.Wait()
}

func handleOut(out io.ReadCloser, in io.WriteCloser, handle OutHandle) error {
	var outBytes []byte
	buf := make([]byte, 1024)
	for {
		n, err := out.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return err
			}
		}
		outBytes = append(outBytes, buf[:n]...)
		if handle != nil {
			err := handle(buf[:n], outBytes, in)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
