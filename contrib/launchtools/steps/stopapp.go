package steps

import (
	"errors"
	"syscall"

	launchertypes "github.com/initia-labs/OPinit/contrib/launchtools"
)

var _ launchertypes.LauncherStepFuncFactory[*launchertypes.Config] = StopApp

func StopApp(_ *launchertypes.Config) launchertypes.LauncherStepFunc {
	return func(ctx launchertypes.Launcher) error {
		if !ctx.IsAppInitialized() {
			return errors.New("app is not initialized")
		}

		log := ctx.Logger()
		log.Info("cleanup")
		log.Info("waiting for app to stop")

		syscall.Kill(syscall.Getpid(), syscall.SIGINT)

		// wait for the app to stop
		ctx.GetErrorGroup().Wait()
		log.Info("cleanup finished")

		return nil
	}
}
