package clashkit

import (
	"fmt"
	"path/filepath"
	"runtime/debug"

	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"
	"github.com/Dreamacro/clash/log"
)

var (
	flagset            map[string]bool
	homeDir            string
	configFile         string
	externalUI         string
	externalController string
	secret             string
)

func init() {
	flagset = map[string]bool{}
}

func Run(withConfig string) {
	// maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))
	debug.SetMemoryLimit(20 * 1 << 20) // 20 MB
	debug.SetMaxThreads(30)            // default 10,000

	if withConfig != "" {
		configFile = withConfig
	}

	if configFile == "" {
		log.Fatalln("Initial configuration directory error: configFile is empty")
	}

	homeDir = filepath.Dir(configFile)
	C.SetHomeDir(homeDir)

	C.SetConfig(configFile)

	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}

	var options []hub.Option
	if flagset["ext-ui"] {
		options = append(options, hub.WithExternalUI(externalUI))
	}
	if flagset["ext-ctl"] {
		options = append(options, hub.WithExternalController(externalController))
	}
	if flagset["secret"] {
		options = append(options, hub.WithSecret(secret))
	}

	if err := hub.Parse(options...); err != nil {
		log.Fatalln("Parse config error: %s", err.Error())
	}

	fmt.Print("Hello, ClashKit")

	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// <-sigCh
}
