package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/concourse/flag"
	"github.com/jessevdk/go-flags"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/sigmon"
	"golang.org/x/crypto/ssh"
)

type QuickstartCommand struct {
	*WebCommand    `group:"Web Configuration"`
	*WorkerCommand `group:"Worker Configuration" namespace:"worker"`
}

func (cmd QuickstartCommand) lessenRequirements(command *flags.Command) {
	cmd.WebCommand.lessenRequirements(command)
	cmd.WorkerCommand.lessenRequirements("worker-", command)

	// autogenerated
	command.FindOptionByLongName("session-signing-key").Required = false
	command.FindOptionByLongName("tsa-authorized-keys").Required = false
	command.FindOptionByLongName("tsa-host-key").Required = false
}

func (cmd *QuickstartCommand) Execute(args []string) error {
	runner, err := cmd.Runner(args)
	if err != nil {
		return err
	}

	return <-ifrit.Invoke(sigmon.New(runner)).Wait()
}

func checkNilKeys(key *flag.PrivateKey) bool {
	if key == nil {
		return true
	}
	if key.PrivateKey == nil {
		return true
	}
	return false
}

func (cmd *QuickstartCommand) Runner(args []string) (ifrit.Runner, error) {
	if checkNilKeys(cmd.WebCommand.RunCommand.Auth.AuthFlags.SigningKey) {
		signingKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, fmt.Errorf("failed to generate session signing key: %s", err)
		}

		cmd.WebCommand.RunCommand.Auth.AuthFlags.SigningKey = &flag.PrivateKey{PrivateKey: signingKey}
		cmd.WebCommand.TSACommand.SessionSigningKey = &flag.PrivateKey{PrivateKey: signingKey}
	}

	if checkNilKeys(cmd.WebCommand.TSACommand.HostKey) {
		tsaHostKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, fmt.Errorf("failed to generate tsa host key: %s", err)
		}

		tsaHostPublicKey, err := ssh.NewPublicKey(tsaHostKey.Public())
		if err != nil {
			return nil, fmt.Errorf("failed to create worker authorized key: %s", err)
		}

		cmd.WebCommand.TSACommand.HostKey = &flag.PrivateKey{PrivateKey: tsaHostKey}
		cmd.WorkerCommand.TSA.PublicKey.Keys =
			append(cmd.WorkerCommand.TSA.PublicKey.Keys, tsaHostPublicKey)
	}

	if checkNilKeys(cmd.WorkerCommand.TSA.WorkerPrivateKey) {
		workerKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, fmt.Errorf("failed to generate worker key: %s", err)
		}

		workerPublicKey, err := ssh.NewPublicKey(workerKey.Public())
		if err != nil {
			return nil, fmt.Errorf("failed to create worker authorized key: %s", err)
		}

		cmd.WorkerCommand.TSA.WorkerPrivateKey = &flag.PrivateKey{PrivateKey: workerKey}
		cmd.WebCommand.TSACommand.AuthorizedKeys.Keys =
			append(cmd.WebCommand.TSACommand.AuthorizedKeys.Keys, workerPublicKey)
	}

	webRunner, err := cmd.WebCommand.Runner(args)
	if err != nil {
		return nil, err
	}

	workerRunner, err := cmd.WorkerCommand.Runner(args)
	if err != nil {
		return nil, err
	}

	logger, _ := cmd.WebCommand.RunCommand.Logger.Logger("quickstart")
	return grouper.NewParallel(os.Interrupt, grouper.Members{
		{
			Name: "web",
			Runner: NewLoggingRunner(logger.Session("web-runner"), webRunner)},
		{
			Name: "worker",
			Runner: NewLoggingRunner(logger.Session("worker-runner"), workerRunner)},
	}), nil
}
