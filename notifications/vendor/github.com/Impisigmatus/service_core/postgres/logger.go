package postgres

import (
	"context"
	"fmt"

	"github.com/Impisigmatus/service_core/log"
	"github.com/jackc/pgx/v4"
)

type pgxLogger struct{}

func (*pgxLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	switch level {
	case pgx.LogLevelTrace:
		log.Trace(msg)
	case pgx.LogLevelDebug:
		log.Debug(msg)
	case pgx.LogLevelInfo:
		log.Info(msg)
	case pgx.LogLevelWarn:
		log.Warning(msg)
	case pgx.LogLevelError:
		log.Error("", fmt.Errorf("%s", msg))
	default:
		log.Error("", fmt.Errorf("INVALID_PGX_LOG_LEVEL"))
	}
}
