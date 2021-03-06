package cmd

import (
	"flag"
	"fmt"

	"github.com/miguelmota/cointop/cointop"
)

// Run ...
func Run() {
	var v, ver, test, clean, reset, hideMarketbar, hideChart, hideStatusbar, onlyTable bool
	var refreshRate uint
	var config, cmcAPIKey, apiChoice, colorscheme string
	flag.BoolVar(&v, "v", false, "Version")
	flag.BoolVar(&ver, "version", false, "Display current version")
	flag.BoolVar(&test, "test", false, "Run test (for Homebrew)")
	flag.BoolVar(&clean, "clean", false, "Wipe clean the cache")
	flag.BoolVar(&reset, "reset", false, "Reset the config. Make sure to backup any relevant changes first!")
	flag.BoolVar(&hideMarketbar, "hide-marketbar", false, "Hide the top marketbar")
	flag.BoolVar(&hideChart, "hide-chart", false, "Hide the chart view")
	flag.BoolVar(&hideStatusbar, "hide-statusbar", false, "Hide the bottom statusbar")
	flag.BoolVar(&onlyTable, "only-table", false, "Show only the table. Hides the chart and top and bottom bars")
	flag.UintVar(&refreshRate, "refresh-rate", 60, "Refresh rate in seconds. Set to 0 to not auto-refresh")
	flag.StringVar(&config, "config", "", "Config filepath. (default ~/.cointop/config.toml)")
	flag.StringVar(&cmcAPIKey, "coinmarketcap-api-key", "", "Set the CoinMarketCap API key")
	flag.StringVar(&apiChoice, "api", cointop.CoinGecko, "API choice. Available choices are \"coinmarketcap\" and \"coingecko\"")
	flag.StringVar(&colorscheme, "colorscheme", "", "Colorscheme to use (defualt \"cointop\"). To install standard themes, do:\n\ngit clone git@github.com:cointop-sh/colors.git ~/.cointop/colors\n\nFor additional instructions, visit: https://github.com/cointop-sh/colors")
	flag.Parse()

	refreshRateFlagFound := false
	var refreshRateP *uint
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "refresh-rate" {
			refreshRateFlagFound = true
		}
	})

	if refreshRateFlagFound {
		refreshRateP = &refreshRate
	}

	if v || ver {
		fmt.Printf("cointop v%s", cointop.Version())
	} else if test {
		doTest()
	} else if clean {
		cointop.Clean()
	} else if reset {
		cointop.Reset()
	} else {
		cointop.NewCointop(&cointop.Config{
			ConfigFilepath:      config,
			CoinMarketCapAPIKey: cmcAPIKey,
			APIChoice:           apiChoice,
			Colorscheme:         colorscheme,
			HideMarketbar:       hideMarketbar,
			HideChart:           hideChart,
			HideStatusbar:       hideStatusbar,
			OnlyTable:           onlyTable,
			RefreshRate:         refreshRateP,
		}).Run()
	}
}

func doTest() {
	cointop.NewCointop(&cointop.Config{
		NoPrompts: true,
	}).Exit()
}
