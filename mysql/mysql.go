package mysql

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MySQLContainer struct {
	opts        *MySQLContainerOptions
	ContainerID *string
}

type MySQLContainerOptions struct {
	Port         string
	Name         string
	RootPassword string
	Database     string
}

func NewMySQLContainer(opts *MySQLContainerOptions) *MySQLContainer {
	return &MySQLContainer{
		opts:        opts,
		ContainerID: new(string),
	}
}

func (m *MySQLContainer) setDefaultOptions() {
	m.opts.Name = strings.TrimSpace(m.opts.Name)
	if m.opts.Name == "" {
		m.opts.Name = defaultContainerName
	}

	m.opts.Database = strings.TrimSpace(m.opts.Database)
	if m.opts.Database == "" {
		m.opts.Database = defaultDatabaseName
	}

	m.opts.Port = strings.TrimSpace(m.opts.Port)
	if m.opts.Port == "" {
		m.opts.Port = defaultDatabasePort
	}

	m.opts.RootPassword = strings.TrimSpace(m.opts.RootPassword)
	if m.opts.RootPassword == "" {
		m.opts.RootPassword = defaultDatabasePassword
	}
}
func (m *MySQLContainer) checkIfDockerIsRunning() error {
	cmd := exec.Command("docker", "ps")
	err := cmd.Run()
	if err != nil {
		return errors.New(dockerNotRunning)
	}
	return nil
}

func (m *MySQLContainer) isNameInUse() bool {
	cmd := exec.Command("docker", "ps", "--quiet", "--filter", "name="+m.opts.Name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return len(output) > 0
}

func (m *MySQLContainer) isPortInUse() bool {
	cmd := exec.Command("docker", "ps", "--quiet", "--filter", "publish="+m.opts.Port)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return len(output) > 0
}

func (m *MySQLContainer) Run() error {
	m.setDefaultOptions()

	err := m.checkIfDockerIsRunning()
	if err != nil {
		return err
	}

	if m.isNameInUse() {
		return errors.New(containerWithThisNameIsAlreadyRunning)
	}

	if m.isPortInUse() {
		return errors.New(containerWithThisPortIsAlreadyRunning)
	}

	port := fmt.Sprintf("%s:%s", m.opts.Port, m.opts.Port)
	rootPassword := fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", m.opts.RootPassword)
	db := fmt.Sprintf("MYSQL_DATABASE=%s", strings.TrimSpace(m.opts.Database))
	name := strings.TrimSpace(m.opts.Name)

	cmd := exec.Command("docker", "run", "-d",
		"-p", port,
		"--name", name,
		"-e", rootPassword,
		"-e", db,
		"-d", "mysql:latest",
	)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run MySQL container: %v, output: %s", err, string(output))
	}

	if len(output) > 70 {
		// Handle long output if needed
	} else {
		containerID := strings.TrimSpace(string(output))
		m.ContainerID = &containerID
	}

	time.Sleep(time.Second * 15)

	log.Println("succesfuly ran container. container id:", *m.ContainerID)

	return nil
}

func (m *MySQLContainer) Stop() error {
	return exec.Command("docker", "rm", "-f", *m.ContainerID).Run()
}
