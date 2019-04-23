/*!
 * go-rs/parser
 * Copyright(c) 2019 Roshan Gade
 * MIT Licensed
 */
package parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type JSON struct {
	data interface{}
}

var (
	invalidJSON = errors.New("JSON is invalid")
)

// Prefer absolute filepath
func (c *JSON) Load(data []byte) (err error) {

	if !json.Valid(data) {
		err = invalidJSON
		return
	}

	err = json.Unmarshal(data, &c.data)
	if err != nil {
		return
	}

	return
}

// Prefer absolute filepath
func (c *JSON) LoadFile(path string) (err error) {
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return c.Load(data)

}

func typeIdentifier(data interface{}) (arr []interface{}, obj map[string]interface{}) {
	switch data.(interface{}).(type) {
	case []interface{}:
		arr = data.([]interface{})
	case map[string]interface{}:
		obj = data.(map[string]interface{})
	}
	return
}

func (c *JSON) Get(key string) (val interface{}, exists bool) {
	keys := strings.Split(key, ".")
	lastIndex := int64(len(keys) - 1)

	var arr []interface{}
	var obj map[string]interface{}

	if c.data == nil {
		return
	}
	arr, obj = typeIdentifier(c.data)

	for _, v := range keys[:lastIndex] {
		if arr == nil && obj == nil {
			break
		}

		if arr != nil {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return
			}
			if i < 0 || i >= int64(len(arr)) {
				return
			}
			arr, obj = typeIdentifier(arr[i])
			continue
		}

		if obj != nil {
			arr, obj = typeIdentifier(obj[v])
			continue
		}
	}

	if arr != nil {
		i, err := strconv.ParseInt(keys[lastIndex], 10, 64)
		if err != nil {
			return
		}
		if i < 0 || i >= int64(len(arr)) {
			return
		}
		val = arr[i]
	} else if obj != nil {
		val = obj[keys[lastIndex]]
	}

	exists = val != nil
	return
}

func (c *JSON) GetString(key string) (s string) {
	val, ok := c.Get(key)
	if ok {
		s, _ = val.(string)
	}
	return
}

func (c *JSON) GetInt(key string) (i int) {
	val, ok := c.Get(key)
	if ok {
		i, _ = val.(int)
	}
	return
}

func (c *JSON) GetFloat(key string) (f float64) {
	val, ok := c.Get(key)
	if ok {
		f, _ = val.(float64)
	}
	return
}

func (c *JSON) GetBool(key string) (b bool) {
	val, ok := c.Get(key)
	if ok {
		b = val.(bool)
	}
	return
}

func (c *JSON) GetTime(key string) (t time.Time) {
	val, ok := c.Get(key)
	if ok {
		t, _ = val.(time.Time)
	}
	return
}

func (c *JSON) GetDuration(key string) (d time.Duration) {
	val, ok := c.Get(key)
	if ok {
		d, _ = val.(time.Duration)
	}
	return
}
