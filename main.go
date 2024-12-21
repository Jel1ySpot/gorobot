package main

import (
	"github.com/Jel1ySpot/conic"
	zerobot "github.com/Jel1ySpot/gorobot/pkg/bot"
	"github.com/Jel1ySpot/gorobot/pkg/plugin"
	"github.com/Jel1ySpot/gorobot/pkg/protocol_logger"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	configuration zerobot.Config
	configPath    = "./config.json"
)

func main() {
	// Check if the configuration file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// If the configuration file does not exist, create a default one
		err := conic.WriteConfig()
		if err != nil {
			panic(err)
		}
		logrus.Infoln("config file created.")
		return
	}

	// Load the configuration file
	if err := conic.ReadInConfig(); err != nil {
		panic(err)
	}

	// Create a logger instance
	logger := zerobot.NewLogger("zerobot -> ")

	// Create QQClient instance
	QQClient := client.NewClient(configuration.Account.Uin, configuration.Account.Password)

	bot := zerobot.Bot{
		QQClient: QQClient,
		Config:   configuration,
		Logger:   logger,
	}
	plugin.InitialPlugins(&bot)

	if err := bot.Login(); err != nil {
		panic(err)
	}

	bot.Logger.Infoln("Login successful")

	defer bot.Release()

	// Save session data before exit
	defer func() {
		sigPath := configuration.Account.SigPath

		data, err := bot.Sig().Marshal()
		if err != nil {
			logrus.Errorf("marshal %s err: %s", sigPath, err)
			return
		}
		err = os.WriteFile(sigPath, data, 0644)
		if err != nil {
			logrus.Errorf("write %s err: %s", sigPath, err)
			return
		}
		logrus.Infoln("sig saved into", sigPath)
	}()

	// set up the main stop channel
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	for {
		switch <-mc {
		case os.Interrupt, syscall.SIGTERM:
			return
		}
	}
}

func init() {
	conic.SetConfigFile(configPath)
	conic.WatchConfig()
	conic.BindRef("bot", &configuration)
	conic.SetLogger(protocol_logger.NewProtocolLogger("conic -> ").Info)
}
