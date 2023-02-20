package main
import (
   "fmt"
   "github.com/mattn/go-colorable"
   "go.uber.org/zap"
   "go.uber.org/zap/zapcore"
)

const (
   INFO = iota
   WARN
   ERROR
)

func main() {
   aa := zap.NewDevelopmentEncoderConfig()
   zapcore.NewTee()
   aa.EncodeLevel = zapcore.CapitalColorLevelEncoder
   bb := zap.New(zapcore.NewCore(
      zapcore.NewConsoleEncoder(aa),
      zapcore.AddSync(colorable.NewColorableStdout()),
      zapcore.DebugLevel,
   ))
   bb.Warn("cc")
   newA := A{
      a: 1,
      b: "2",
   }
   l,_ := zap.NewProduction()
   s := l.Sugar()
   s.Infof("\n%+v", newA)
}

type A struct {
   a int
   b string
}