package utils

var (
	utilslog = GetLog()
)

func init() {
	defer utilslog.Flush()
	utilslog.Info("success to init seelog utilslog")
}
