package postgres

import "testing"

func TestPostgresContainer_RunStop(t *testing.T) {
	opts := &PostgresContainerOptions{}
	container := NewPostgresContainer(opts)

	t.Run("test run container", func(t *testing.T) {
		err := container.Run()
		if err != nil {
			t.Fatal("expected to run successfully but got err:", err)
		}
	})

	t.Run("test stop container", func(t *testing.T) {
		err := container.Stop()
		if err != nil {
			t.Fatal("expected to stop successfully but got err:", err)
		}
	})
}
