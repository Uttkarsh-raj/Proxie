package model

import (
	"fmt"
	"os"
	"time"
)

type Logs struct {
	Date          time.Time
	ClientIp      string
	Method        string
	ServerAddress string
	Agent         string
}

func Log(clientIp, method, serverIp, agent string) Logs {
	return Logs{
		Date:          time.Now(),
		ClientIp:      clientIp,
		Method:        method,
		ServerAddress: serverIp,
		Agent:         agent,
	}
}

var (
	filename = "./logs.txt"
)

func (l *Logs) AppendLog() error {
	flag := false
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		flag = true
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()
	if flag {
		_, err = fmt.Fprintf(file, "%s \t %s \t %s \t %s \t %s\n",
			"Date", "Client-Ip", "Method", "Server-Address", "User-Agent")
		if err != nil {
			return fmt.Errorf("failed to write to log file: %w", err)
		}
	}

	_, err = fmt.Fprintf(file, "%s \t %s \t %s \t %s \t %s\n",
		l.Date.Format(time.RFC3339), l.ClientIp, l.Method, l.ServerAddress, l.Agent)
	if err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}
	return nil
}
