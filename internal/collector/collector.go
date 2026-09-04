package collector

import (
	"sort"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
	gnet "github.com/shirou/gopsutil/v4/net"
	"go.uber.org/zap"
)

type Listener struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Port     uint32 `json:"port"`
	PID      int32  `json:"pid"`
	Process  string `json:"process"`
	Exe      string `json:"exe"`
}

func (l *Listener) String() string {
	return strings.Join([]string{l.Protocol, l.Address, strconv.Itoa(int(l.Port)), strconv.Itoa(int(l.PID)), l.Process, l.Exe}, "/")
}

type processInfo struct {
	name string
	exe  string
}

type Collector struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Collector {
	return &Collector{
		logger: logger,
	}
}

func (c *Collector) Listeners() ([]*Listener, error) {
	listeners := []*Listener{}

	for _, kind := range []string{"tcp", "udp"} {
		kindListeners, err := c.collectListeners(kind, kind)
		if err != nil {
			return nil, err
		}
		listeners = append(listeners, kindListeners...)
	}
	sort.Slice(listeners, func(i, j int) bool {
		if listeners[i].Port != listeners[j].Port {
			return listeners[i].Port < listeners[j].Port
		}
		return listeners[i].Protocol < listeners[j].Protocol
	})

	return listeners, nil
}
func (c *Collector) loadProcess(pid int32) processInfo {
	info := processInfo{name: "unknown"}

	if pid > 0 {
		if p, err := process.NewProcess(pid); err == nil {
			if name, err := p.Name(); err == nil {
				info.name = name
			}
			if exe, err := p.Exe(); err == nil {
				info.exe = exe
			}
		}
	}

	return info
}
func (c *Collector) collectListeners(kind, protocol string) ([]*Listener, error) {
	listeners := []*Listener{}

	connections, err := gnet.Connections(kind)
	if err != nil {
		c.logger.Info("cannot read sockets", zap.String("kind", kind), zap.Error(err))
		return nil, err
	}

	for _, connection := range connections {
		// TCP имеет состояние LISTEN.
		if protocol == "tcp" &&
			!strings.EqualFold(connection.Status, "LISTEN") {
			continue
		}

		// У UDP состояния LISTEN нет: показываем все сокеты,
		// привязанные к локальному порту.
		if connection.Laddr.Port == 0 {
			continue
		}

		info := c.loadProcess(connection.Pid)

		listeners = append(listeners, &Listener{
			Protocol: protocol,
			Address:  connection.Laddr.IP,
			Port:     connection.Laddr.Port,
			PID:      connection.Pid,
			Process:  info.name,
			Exe:      info.exe,
		})
	}

	return listeners, nil
}
