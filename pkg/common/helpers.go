package common

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func RunCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("%v; command=%#+v, output=%#+v", err, command, string(output))
	}

	return string(output), nil
}

func HandleRunCommandError(output string, err error) error {
	if err != nil {
		return fmt.Errorf("%v; %v", err, output)
	}

	return nil
}

func GetRandomUUID() uuid.UUID {
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(fmt.Errorf("error generating UUID; %v", err))
	}

	return randomUUID
}

func WaitForSIGINT() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
