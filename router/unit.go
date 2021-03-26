// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 25/3/2564 20:12
// ---------------------------------------------------------------------------
package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func newLog(param gin.LogFormatterParams) string {

	// Log format
	return fmt.Sprintf("[LN] %s      | [%s]        %s |  %d |  %s  | %s |     \"%s\"|  \"%s\"       %s\n",
		param.ClientIP,
		param.TimeStamp.Format("2006/01/02-15:04:05"),
		param.Method,
		param.StatusCode,
		param.Latency,
		param.Request.Proto,
		param.Path,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
