package sd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// @Summary Shows OK as the ping-pong result
// @Description Shows OK as the ping-pong result
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "OK"
// @Router /sd/health [get]
func HealthCheck(c *gin.Context) { //检查API Server的状态：健康状况。
	message := "OK"
	c.String(http.StatusOK, message+"\n")
}

// @Summary Checks the disk usage
// @Description Checks the disk usage
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "OK - Free space: 17233MB (16GB) / 51200MB (50GB) | Used: 33%"
// @Router /sd/disk [get]
func DiskCheck(c *gin.Context) { //检查API Server的状态：服务器硬盘
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests //StatusTooManyRequests  =429
		text = "WARNING"
	}
	//比如我的电脑D盘总公250G，用了119G，网页输出OK - Free space: 134132MB (130GB) / 256708MB (250GB) | Used: 52%
	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message+"\n")
}

// @Summary Checks the cpu usage
// @Description Checks the cpu usage
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "CRITICAL - Load average: 1.78, 1.99, 2.02 | Cores: 2"
// @Router /sd/cpu [get]
func CPUCheck(c *gin.Context) { //检查API Server的状态：CPU使用量。
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	//cores是int类型，l5是float64 类型
	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	//举例网页输出：OK - Load average: 0.00, 0.00, 0.00 | Cores: 8
	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores) //默认宽度，精度为2
	c.String(status, message+"\n")
}

// @Summary Checks the ram usage
// @Description Checks the ram usage
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "OK - Free space: 402MB (0GB) / 8192MB (8GB) | Used: 4%"
// @Router /sd/ram [get]
func RAMCheck(c *gin.Context) { //检查API Server的状态：内存使用量。
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	//举例网页输出：OK - Free space: 11246MB (10GB) / 15734MB (15GB) | Used: 71%
	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message+"\n")
}
