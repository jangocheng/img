package main

import (
	"github.com/hawx/img/hsla"
	"github.com/hawx/img/utils"
)

var cmdSaturation = &Command{
	UsageLine: "saturation [options]",
	Short:     "adjust image saturation",
Long: `
  Saturation takes a png file from STDIN, adjusts the saturation and prints the
  result to STDOUT

    --by [n]       # Amount to adjust saturation by (default: 0.1)
`,
}

var saturationBy float64

func init() {
	cmdSaturation.Run = runSaturation

	cmdSaturation.Flag.Float64Var(&saturationBy, "by", 0.1, "")
}

func runSaturation(cmd *Command, args []string) {
	i := utils.ReadStdin()
	i  = hsla.Saturation(i, saturationBy)
	utils.WriteStdout(i)
}
