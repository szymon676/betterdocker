package mysql

import (
	"testing"
)

func TestMySQLContainer_RunStop(t *testing.T) {
	opts := &MySQLContainerOptions{
		Port:         "3306",
		Name:         "test-mysql-container",
		RootPassword: "testrootpass",
		Database:     "testdb",
	}

	container := NewMySQLContainer(opts)

	t.Run("test run function", func(t *testing.T) {
		err := container.Run()
		if err != nil {
			t.Fatalf("Failed to run MySQL container: %v", err)
		}
		if container.ContainerID == nil {
			t.Fatal("ContainerID should not be nil after running the container")
		}
	})

	t.Run("test stop funciton", func(t *testing.T) {
		err := container.Stop()
		if err != nil {
			t.Fatalf("Failed to stop MySQL container: %v", err)
		}
	})
}
