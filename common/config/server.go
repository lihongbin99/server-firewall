package config

import (
	"io/ioutil"
	"net"
	"os"
	"security-network/common/secret"
	"strconv"
	"strings"
)

var (
	ListenIp   = ""
	ListenPort = 13520

	PrivateKey         = "config/security/private.pem"
	SecurityPrivateKey []byte

	TCPInterval = 180

	Secrets = make([]secret.Secret, 0)

	Listener   *net.TCPListener
	StopAccept = false
)

func ServerInit() {
	// 获取监听配置
	if listenIp := IniFile.Section("listen").Key("ip").String(); listenIp != "" {
		ListenIp = listenIp
	}
	if listenPort := IniFile.Section("listen").Key("port").String(); listenPort != "" {
		port, err := strconv.Atoi(listenPort)
		if err != nil {
			log.Error("listen port error:", listenPort)
			panic(err)
		}
		ListenPort = port
	}

	// 读取加密配置
	if securityKey := IniFile.Section("security").Key("key").String(); securityKey != "" {
		_, err := os.Stat(securityKey)
		if err != nil {
			log.Error("not find security key:", securityKey)
			panic(err)
		}
		PrivateKey = securityKey
	}
	if temp, err := ioutil.ReadFile(PrivateKey); err != nil {
		log.Error("read private key pem error", err)
	} else {
		SecurityPrivateKey = temp
	}

	// 读取超时配置
	if tcpInterval := IniFile.Section("tcp").Key("timeout").String(); tcpInterval != "" {
		interval, err := strconv.Atoi(tcpInterval)
		if err != nil {
			log.Error("tcp interval error:", tcpInterval)
			panic(err)
		}
		TCPInterval = interval
	}

	// 加载所有配置文件
	for _, section := range IniFile.Sections() {
		if strings.HasPrefix(section.Name(), "secret_") {
			Secrets = append(Secrets, secret.GetSecret(section))
		}
	}

	log.Trace("listen ip   :", ListenIp)
	log.Trace("listen port :", ListenPort)
	log.Trace("security key:", PrivateKey)
	log.Trace("tcp timeout :", TCPInterval)
}

func StopServer() {
	StopAccept = true
	if Listener != nil {
		_ = Listener.Close()
	}
}
