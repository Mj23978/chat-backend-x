package viper

import (
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	logger "github.com/mj23978/chat-backend-x/logger/zerolog"
	"github.com/spf13/viper"
)

var cfgFile string
var watchers []func(event fsnotify.Event) error
var watcherLock sync.Mutex
var all map[string]interface{}

// ErrRollbackConfigurationChanges should be used when a configuration is e.g. invalid and should be rolled back.
var ErrRollbackConfigurationChanges = errors.New("an error occurred and configuration changes should be reverted")

// RegisterConfigFlag registers the --config / -c flag.
func RegisterConfigFlag(c *cobra.Command, applicationName string) {
	c.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", `Path to config file. Supports .json, .yaml, .yml, .toml. Default is "$HOME/.`+applicationName+`.(yaml|yml|toml|json)"`)
}

// WatchOptions configures WatchConfig.
type WatchOptions struct {
	// Immutables are keys that cause OnImmutableChange to be fired when modified.
	Immutables []string
	// OnImmutableChange - see Immutables.
	OnImmutableChange func(key string)
}

// AddWatcher adds a function callback to viper.OnConfigChange().
func AddWatcher(f func(event fsnotify.Event) error) {
	watcherLock.Lock()
	defer watcherLock.Unlock()
	watchers = append(watchers, f)
}

// ResetWatchers resets all the watchers.
func ResetWatchers() {
	watcherLock.Lock()
	defer watcherLock.Unlock()
	watchers = nil
}

// WatchConfig is a helper makes watching configuration files easy.
func WatchConfig(o *WatchOptions) {

	if o == nil {
		o = new(WatchOptions)
	}

	watcherLock.Lock()
	all = viper.AllSettings()
	watcherLock.Unlock()

	for _, key := range o.Immutables {
		// This ensures that the keys are all set
		viper.Get(key)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		watcherLock.Lock()
		defer watcherLock.Unlock()

		//defer l.WithField("file", in.Name).
		//	WithField("operator", in.Op.String()).
		//	WithField("immutables", o.Immutables).
		//	Info("A change to the configuration file was processed.")

		var didReset bool
		//for _, key := range o.Immutables {
			//if viper.HasChangedSinceInit(key) {
			//	viper.SetRawConfig(all)
			//	didReset = true
			//	if o.OnImmutableChange != nil {
			//		o.OnImmutableChange(key)
			//	}
			//}
		//}

		for _, w := range watchers {
			if err := w(in); errors.Cause(err) == ErrRollbackConfigurationChanges {
				//viper.SetRawConfig(all)
				didReset = true
				break
			} else if err != nil {
				logger.Errorf("A configuration watcher returned an error code, stopping event propagation %v", err)
				return
			}
		}

		if !didReset {
			all = viper.AllSettings()
		}
	})
}

// InitializeConfig initializes viper.
func InitializeConfig(applicationName string, homeOverride string)  {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(strings.TrimSpace(cfgFile))
	} else {
		// Find home directory.
		home, _ := homedir.Dir()

		if homeOverride != "" {
			home = homeOverride
		}

		// Search config in home directory with the application name and a dot prepended (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("." + applicationName)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	//err := viper.ReadInConfig()

	//if err == nil {
	//	l.WithField("path", viper.ConfigFileUsed()).Info("Config file loaded successfully.")
	//} else {
	//	switch t := err.(type) {
	//	case viper.UnsupportedConfigError:
	//		if len(t) == 0 {
	//			l.WithError(err).Warn("No config file was defined and no file was found in $HOME/." + applicationName + ".yaml")
	//		} else {
	//			l.WithError(err).WithField("extension", t).Fatal("Unsupported configuration type")
	//		}
	//	case *viper.ConfigFileNotFoundError:
	//		l.WithError(err).Warn("No config file was defined and no file was found in $HOME/." + applicationName + ".yaml")
	//	case viper.ConfigFileNotFoundError:
	//		l.WithError(err).Warn("No config file was defined and no file was found in $HOME/." + applicationName + ".yaml")
	//	default:
	//		l.
	//			WithField("path", viper.ConfigFileUsed()).
	//			WithError(err).
	//			Fatal("Unable to open config file. Make sure it exists and the process has sufficient permissions to read it")
	//	}
	//}
}
