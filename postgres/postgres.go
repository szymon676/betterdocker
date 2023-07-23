package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type PostgresContainerOptions struct {
	Port     string
	Name     string
	Password string
	Database string
}

type PostgresContainer struct {
	opts        *PostgresContainerOptions
	ContainerID *string
}

func NewPostgresContainer(opts *PostgresContainerOptions) *PostgresContainer {
	return &PostgresContainer{
		opts:        opts,
		ContainerID: new(string),
	}
}

func (ps *PostgresContainer) setDefaultOptions() {
	ps.opts.Name = strings.TrimSpace(ps.opts.Name)
	if ps.opts.Name == "" {
		ps.opts.Name = defaultContainerName
	}

	ps.opts.Database = strings.TrimSpace(ps.opts.Database)
	if ps.opts.Database == "" {
		ps.opts.Database = defaultDatabaseName
	}

	ps.opts.Port = strings.TrimSpace(ps.opts.Port)
	if ps.opts.Port == "" {
		ps.opts.Port = defaultDatabasePort
	}

	ps.opts.Password = strings.TrimSpace(ps.opts.Password)
	if ps.opts.Password == "" {
		ps.opts.Password = defaultDatabasePassword
	}
}

func (ps *PostgresContainer) checkIfDockerIsRunning() error {
	cmd := exec.Command("docker", "ps")
	err := cmd.Run()
	if err != nil {
		return errors.New(dockerNotRunning)
	}
	return nil
}

func (ps *PostgresContainer) isNameInUse() bool {
	cmd := exec.Command("docker", "ps", "--quiet", "--filter", "name="+ps.opts.Name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return len(output) > 0
}

func (ps *PostgresContainer) isPortInUse() bool {
	cmd := exec.Command("docker", "ps", "--quiet", "--filter", "publish="+ps.opts.Port)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return len(output) > 0
}

func (ps *PostgresContainer) Run() error {
	ps.setDefaultOptions()

	err := ps.checkIfDockerIsRunning()
	if err != nil {
		return err
	}

	if ps.isPortInUse() {
		return errors.New(containerWithThisPortIsAlreadyRunning)
	}

	if ps.isNameInUse() {
		return errors.New(containerWithThisNameIsAlreadyRunning)
	}

	port := fmt.Sprintf("%s:%s", ps.opts.Port, ps.opts.Port)
	password := fmt.Sprintf("POSTGRES_PASSWORD=%s", ps.opts.Password)
	db := fmt.Sprintf("POSTGRES_DB=%s", strings.TrimSpace(ps.opts.Password))
	name := strings.TrimSpace(ps.opts.Name)

	cmd := exec.Command("docker", "run",
		"-p", port,
		"--name", name,
		"-e", password,
		"-e", db,
		"-d", "postgres",
	)

	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	if len(output) > 70 {
		// Handle long output if needed
	} else {
		containerID := strings.TrimSpace(string(output))
		ps.ContainerID = &containerID
	}

	time.Sleep(time.Second * 15)

	log.Println("container id:", strings.TrimSpace(string(output)))

	return nil
}

func (ps *PostgresContainer) Stop() error {
	return exec.Command("docker", "rm", "-f", *ps.ContainerID).Run()
}
