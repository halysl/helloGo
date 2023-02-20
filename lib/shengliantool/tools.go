package shengliantool

import (
	"github.com/halysl/hellogo/lib/types"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("shengliantool")

var WinningStatus = new(types.ElectionProof)

func SetWiningStatus(ep *types.ElectionProof) error {
	WinningStatus = ep
	log.Infof("into SetWiningStatus:%+v", WinningStatus)
	return nil
}

func ResetWinningStatus() error {
	WinningStatus = new(types.ElectionProof)
	log.Infof("into ResetWinningStatus:%+v", WinningStatus)
	return nil
}

func GetWinningStatus() (*types.ElectionProof, error) {
	log.Infof("into GetWinningStatus:%+v", WinningStatus)
	return WinningStatus, nil
}
