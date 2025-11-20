package main

import (
    "github.com/hoonfeng/goproc/sdk"
)

func reverse(params map[string]interface{}) (interface{}, error) {
    s, _ := params["text"].(string)
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { r[i], r[j] = r[j], r[i] }
    return string(r), nil
}

func main() {
    sdk.RegisterFunction("reverse", reverse)
    _ = sdk.Start()
    sdk.Wait()
}
