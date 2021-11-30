package magic

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 返回结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功返回
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg, "data": data})
}

// Fail 失败返回
func Fail(c *gin.Context, code int, msg string, data ...interface{}) {
	// data如何默认
	c.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": msg, "data": data})
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

// Serialization 序列化
func Serialization(value interface{}) ([]byte, error) {
	if bytes, ok := value.([]byte); ok {
		return bytes, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []byte(strconv.FormatUint(v.Uint(), 10)), nil
	case reflect.Map:
	}
	k, err := json.Marshal(value)
	return k, err
}

// Deserialization 反序列化
func Deserialization(byt []byte, ptr interface{}) (err error) {
	if bytes, ok := ptr.(*[]byte); ok {
		*bytes = byt
		return
	}
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		switch p := v.Elem(); p.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			i, err = strconv.ParseInt(string(byt), 10, 64)
			if err != nil {
				fmt.Printf("Deserialization: failed to parse int '%s': %s", string(byt), err)
			} else {
				p.SetInt(i)
			}
			return

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			i, err = strconv.ParseUint(string(byt), 10, 64)
			if err != nil {
				fmt.Printf("Deserialization: failed to parse uint '%s': %s", string(byt), err)
			} else {
				p.SetUint(i)
			}
			return
		}
	}
	err = json.Unmarshal(byt, &ptr)
	return
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
