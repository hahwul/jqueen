package daemon

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/hahwul/jqueen/pkg/job"
	"github.com/labstack/echo"
)

type Return struct {
	Code int `json:"code"`
	JobID string `json:"jobid"`
}

// AddJob is add job
func AddJob(c echo.Context, jq *job.JobQueue) error {
	// Make JobID
	t := time.Now()
	sha256 := sha256.Sum256([]byte(t.String()))
	s := fmt.Sprintf("%x",sha256)

	j := new(job.Job)
	j.ID = s	

	jq.Enqueue(j)
	r := &Return{
		Code: 0,
		JobID: j.ID,
	}
	return c.JSON(http.StatusOK,r)
}
