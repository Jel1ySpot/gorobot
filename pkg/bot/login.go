package bot

import (
	"fmt"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strings"
	"time"
)

func (b *Bot) Login() error {
	appInfo := auth.AppList[strings.Split(b.Config.AppInfo, " ")[0]][strings.Split(b.Config.AppInfo, " ")[1]]

	b.QQClient.UseVersion(appInfo)
	b.QQClient.AddSignServer(b.Config.SignServerUrl)
	b.QQClient.SetLogger(&b.Logger.protocolLogger)

	deviceInfo, err := auth.LoadOrSaveDevice("./device.json")
	if err != nil {
		return err
	}

	b.QQClient.UseDevice(deviceInfo)

	logger := b.Logger

	// Login
	data, err := os.ReadFile(b.Config.Account.SigPath)
	if err == nil {
		sig, err := auth.UnmarshalSigInfo(data, true)
		if err != nil {
			logrus.Warnln("load sig error:", err)
		} else {
			// FastLogin
			logger.Infoln("Try FastLogin")
			b.UseSig(sig)
			err = b.FastLogin()

			if err != nil {
				logger.Error("FastLogin fail: %s", err)
			} else {
				return nil
			}

			// EasyLogin
			logger.Infoln("Try EasyLogin")
			if len(sig.TempPwd) != 0 {
				ret, err := b.TokenLogin()
				if err != nil {
					logger.Error("EasyLogin fail: %s", err)
				}

				if ret.Success {
					return nil
				}
			}
		}
	}

	// QRCodeLogin
	logger.Infoln("login with qrcode")
	_, uri, err := b.FetchQRCodeDefault()
	if err != nil {
		return err
	}
	logger.Info("https://api.qrserver.com/v1/create-qr-code/?data=%s", url.QueryEscape(uri))
	for {
		retCode, err := b.GetQRCodeResult()
		if err != nil {
			logger.Errorln(err)
			return err
		}
		if retCode.Waitable() {
			time.Sleep(3 * time.Second)
			continue
		}
		if !retCode.Success() {
			return fmt.Errorf(retCode.Name())
		}
		break
	}
	_, err = b.QRCodeLogin()
	if err != nil {
		return err
	}

	return nil
}
