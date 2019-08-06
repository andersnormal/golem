package cmd

import (
	"context"

	"github.com/andersnormal/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type root struct {
	logger *log.Entry
}

func runE(cmd *cobra.Command, args []string) error {
	// create a new root
	root := new(root)

	// init logger
	root.logger = log.WithFields(log.Fields{
		"debug":   cfg.Debug,
		"verbose": cfg.Verbose,
	})

	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s := server.NewServer(ctx)

	// listen for the server and wait for it to fail,
	// or for sys interrupts
	if err := s.Wait(); err != nil {
		root.logger.Error(err)
	}

	// noop
	return nil
}
