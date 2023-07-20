package mysql

import (
	"fmt"
	"os"
	"os/exec"
)

type MySQLOptions struct {
	port          string
	name          string
	root_password string
	database      string
}

type MySQLContainer struct {
	opts MySQLOptions
}

func NewMySQLContainer(opts MySQLOptions) *MySQLContainer {
	return &MySQLContainer{
		opts: opts,
	}
}

func (m *MySQLContainer) RunMySQLContainer() {
	port := fmt.Sprintf("%s:%s", m.opts.port, m.opts.port)
	name := m.opts.name
	root_password := fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", m.opts.root_password)
	db := fmt.Sprintf("MYSQL_DATABASE=%s", m.opts.database)
	cmd := exec.Command("docker", "run",
		"-p", port,
		"--name", name,
		"-e", root_password,
		"-e", db,
		"-d", "mysql:latest",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(output))
		os.Exit(1)
	}
}
