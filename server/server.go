package server

import (
	"math/rand"
	"net/http"
	"tasks/database"
	"tasks/worker"

	"github.com/gin-gonic/gin"
)

type request struct {
	Number      int `form:"number"`
	MinDuration int `form:"min_duration"`
	MaxDuration int `form:"max_duration"`
	Probability int `form:"probability"`
}

func Run(port string) {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/start", handle)

	router.Run("127.0.0.1:" + port)
}

func handle(context *gin.Context) {
	var req request

	if err := context.Bind(&req); err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	if req.Probability <= 0 {
		context.String(http.StatusBadRequest, "Probability must be grater than 0")
		return
	}

	if req.Probability > 100 {
		context.String(http.StatusBadRequest, "Probability must be less than 100")
		return
	}

	if req.MinDuration < 0 {
		context.String(http.StatusBadRequest, "Min duration must be grater than or equal to 0")
		return
	}

	if req.MaxDuration < 0 {
		context.String(http.StatusBadRequest, "Max duration must be grater than or equal to 0")
		return
	}

	if req.MaxDuration < req.MinDuration {
		context.String(http.StatusBadRequest, "Max duration must be grater than min duration")
		return
	}

	count, err := startTasks(req)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.String(http.StatusOK, "Executing %d tasks.", count)
}

func startTasks(req request) (int, error) {
	taskIds, err := database.GetTaskIds(req.Number)
	if err != nil {
		return 0, err
	}

	for _, taskId := range taskIds {
		duration := getDuration(req.MinDuration, req.MaxDuration)

		worker.AddTask(&worker.Task{
			Id:          taskId,
			Duration:    duration,
			Probability: req.Probability,
		})
	}

	return len(taskIds), nil
}

func getDuration(min int, max int) int {
	n := max - min + 1
	if n < 1 {
		n = 1
	}

	return min + rand.Intn(n)
}
