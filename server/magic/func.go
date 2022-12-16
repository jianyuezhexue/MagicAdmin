package magic

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ------------------待迁移走开始-------------------
// Response 返回结构体
type Response struct {
	Code     int     `json:"code"`
	Msg      string  `json:"msg"`
	Data     any     `json:"data"`
	CostTime float64 `json:"costTime"`
}

// 分页查询返回结构体
type PageResult struct {
	List     any   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

// Success 成功返回
func Success(c *gin.Context, msg string, data any) {
	// 获取开始时间
	start := c.GetInt64("magicStartTime")
	log.Println(start)
	// 计算耗时
	end := time.Now().UnixNano()
	costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", (float64(end)-float64(start))/1000000), 64)

	// 返回结果
	res := Response{Code: 0, Msg: msg, Data: data, CostTime: costTime}

	// 预埋操作记录参数
	c.Set("magicMsg", msg)
	recodeBytes, _ := json.Marshal(res)
	c.Set("magicResp", string(recodeBytes))

	c.JSON(http.StatusOK, res)
}

// Fail 失败返回
func Fail(c *gin.Context, code int, msg string, data any) {
	// 获取开始时间
	start := c.GetInt64("magicStartTime")
	log.Println(start)

	// 计算耗时
	end := time.Now().UnixNano()
	costTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", (float64(end)-float64(start))/1000000), 64)

	// 返回结果
	res := Response{Code: -1, Msg: msg, Data: data, CostTime: costTime}

	// 预埋操作记录参数
	c.Set("magicMsg", msg)
	recodeBytes, _ := json.Marshal(res)
	c.Set("magicResp", string(recodeBytes))

	c.JSON(http.StatusOK, res)
}

// ------------------待迁移走结束-------------------

// 调试打印数据
func Print(data any) {
	toJson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(toJson))
}

// PathExists 判断文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MD5V md5加密
func MD5V(str string, b ...byte) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(b))
}

// GetTime 获取当前时间|
func GetTime(format string) string {
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := now.Format("01")
	day := strconv.Itoa(now.Day())

	switch format {
	case "Y-m-d":
		return year + "-" + month + "-" + day
	case "Y-m-d H:i:s":
		hour := now.Format("15")
		min := now.Format("04")
		second := now.Format("05")
		return year + "-" + month + "-" + day + " " + hour + ":" + min + ":" + second
	default:
		return year + "-" + month + "-" + day
	}
}
