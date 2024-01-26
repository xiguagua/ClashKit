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
	testConfig         bool
	homeDir            string
	configFile         string
	externalUI         string
	externalController string
	secret             string
)

func init() {
	// flag.StringVar(&externalUI, "ext-ui", "", "override external ui directory")
	// flag.StringVar(&externalController, "ext-ctl", "", "override external controller address")
	// flag.StringVar(&secret, "secret", "", "override secret for RESTful API")

	flagset = map[string]bool{}
}

func Run(withConfig string) {
	// maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))
	debug.SetMemoryLimit(20 * 1 << 20) // 20 MB
	debug.SetMaxThreads(30)            // default 10,000

	if withConfig != "" {
		configFile = withConfig
	}

	// if homeDir != "" {
	// 	if !filepath.IsAbs(homeDir) {
	// 		currentDir, _ := os.Getwd()
	// 		homeDir = filepath.Join(currentDir, homeDir)
	// 	}
	// 	C.SetHomeDir(homeDir)
	// }

	if configFile == "" {
		log.Fatalln("Initial configuration directory error: configFile is empty")
	}

	homeDir = filepath.Dir(configFile)
	C.SetHomeDir(homeDir)

	// if !filepath.IsAbs(configFile) {
	// 	currentDir, _ := os.Getwd()
	// 	configFile = filepath.Join(currentDir, configFile)
	// }
	C.SetConfig(configFile)

	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}

	// if testConfig {
	// 	if _, err := executor.Parse(); err != nil {
	// 		log.Errorln(err.Error())
	// 		fmt.Printf("configuration file %s test failed\n", C.Path.Config())
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("configuration file %s test is successful\n", C.Path.Config())
	// 	return
	// }

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
	return

	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// <-sigCh
}
