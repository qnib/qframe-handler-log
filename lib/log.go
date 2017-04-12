package qframe_handler_log

import (
	"strings"
	"log"
	"fmt"
	"github.com/zpatrick/go-config"

	"github.com/qnib/qframe-types"
	"github.com/qnib/qframe-utils"
)

const (
	version = "0.1.0"
)

type Plugin struct {
	qtypes.Plugin
}

func New(qChan qtypes.QChan, cfg config.Config, name string) Plugin {
	p := Plugin{
		Plugin: qtypes.Plugin{
			QChan: qChan,
			Cfg:   cfg,
		},
	}
	p.Version = version
	p.Name = name
	return p
}

// Run fetches everything from the Data channel and flushes it to stdout
func (p *Plugin) Run() {
	log.Printf("[II] Start log handler v%s", version)
	bg := p.QChan.Data.Join()
	inStr, err := p.Cfg.String("handler.log.inputs")
	if err != nil {
		inStr = ""
	}
	inputs := strings.Split(inStr, ",")
	for {
		val := bg.Recv()
		qm := val.(qtypes.QMsg)
		if len(inputs) != 0 && ! qutils.IsInput(inputs, qm.Source) {
			//fmt.Printf("%s %-7s sType:%-6s sName:%-10s[%d] DROPED : %s\n", qm.TimeString(), qm.LogString(), qm.Type, qm.Source, qm.SourceID, qm.Msg)
			continue
		}
		fmt.Printf("%s %-7s sType:%-6s sName:[%d]%-10s %s\n", qm.TimeString(), qm.LogString(), qm.Type, qm.SourceID, qm.Source, qm.Msg)
	}
}
