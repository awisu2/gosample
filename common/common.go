package common

import (
  "log"
)

/**
 * エラーの場合、ログを吐いてExit
 * @param  {[type]} err error         [description]
 * @return {[type]}     [description]
 */
func LogError(err error) bool{
  if err != nil {
    log.Print(err)
    return true
  }
  return false
}

func Log(v ...interface{}) {
  for _, value := range v {
    log.Printf("Type:%T, Value:%v", value, value)
  }
}

