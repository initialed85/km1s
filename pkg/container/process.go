package container

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/initialed85/km1s/pkg/common"
	"io"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

type Process struct {
	cmd       *exec.Cmd
	processID uuid.UUID
	stdin     io.WriteCloser
	logBuffer *common.RingBuffer
}

func NewProcess(
	command string,
	logBufferSize int,
	environment map[string]string,
) *Process {
	p := Process{
		cmd:       exec.Command("bash", "-c", "-i", command),
		processID: common.GetRandomUUID(),
		logBuffer: common.NewRingBuffer(logBufferSize),
	}

	p.cmd.Stdout = p.logBuffer
	p.cmd.Stderr = p.logBuffer
	p.cmd.Env = make([]string, 0)

	for k, v := range environment {
		p.cmd.Env = append(p.cmd.Env, fmt.Sprintf("%v=%v", k, v))
	}

	return &p
}

func (p *Process) ProcessID() uuid.UUID {
	return p.processID
}

func (p *Process) Open() error {
	var err error

	p.stdin, err = p.cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = p.cmd.Start()
	if err != nil {
		return err
	}

	runtime.Gosched()

	time.Sleep(time.Millisecond * 100) // TODO: fix arbitrary sleep

	return nil
}

func (p *Process) Send(data string) error {
	_, err := p.stdin.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func (p *Process) Logs() string {
	buf := make([]byte, p.logBuffer.Size())

	n, _ := p.logBuffer.Read(buf)

	return string(buf[0:n])
}

func (p *Process) Close() {
	if p.cmd.Process != nil {
		_ = p.cmd.Process.Signal(syscall.SIGINT)
		_ = p.Send("exit")
	}

	timer := time.NewTimer(time.Second * 10)
	go func() {
		_ = <-timer.C

		if p.cmd.Process != nil {
			_ = p.cmd.Process.Kill()
		}
	}()
	_, _ = p.cmd.Process.Wait()
	_ = p.cmd.Wait()
	_ = timer.Stop()
}
