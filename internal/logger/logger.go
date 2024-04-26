package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func SetupLogrus(allLogsFile io.Writer, warnLogsFile io.Writer) {
	log.SetOutput(allLogsFile)

	log.AddHook(&writer.Hook{
		Writer: warnLogsFile,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	log.AddHook(&writer.Hook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
}
