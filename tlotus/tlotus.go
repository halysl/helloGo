package tlotus

import (
	"context"
	"fmt"
	"github.com/halysl/hellogo/lib/shengliantool"
	"github.com/halysl/hellogo/lib/types"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var PL = new(types.ProofWithLock)

func produce() {
	log.Info("into produce")
	ep := &types.ElectionProof{
		WinCount: 20,
		VRFProof: nil,
	}
	time.Sleep(20*time.Millisecond)
	for {
		n := 1 + rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		shengliantool.SetWiningStatus(ep)
	}
}

func NewContextTest() {
	go produce()
	go func() {
		for i:= 0; i<=10; i++ {
			NeedLockRun(i)
		}
	}()
	go func() {
		taskTicker := time.NewTicker(time.Millisecond * 30)
		defer taskTicker.Stop()
		for {
			select {
			case <-taskTicker.C:
				n := 1 + rand.Intn(20)
				time.Sleep(time.Duration(n) * time.Millisecond)
				contextTest()
			}
		}
	}()

	time.Sleep(10*time.Second)

}

func contextTest() error {
	PL.Lock.Lock()
	defer PL.Lock.Unlock()
	ctx := context.Background()
	customCtx, cancel := context.WithCancel(context.Background())
	ch := make(chan error, 1)
	go func() {
		winningStatusTicker := time.NewTicker(time.Millisecond * 5)
		defer winningStatusTicker.Stop()
		for {
			select {
			case <-winningStatusTicker.C:
				ep, _ := shengliantool.GetWinningStatus()
				if ep != nil && ep.WinCount > 0 {
					//log.Infof("will cancel")
					cancel()
					shengliantool.ResetWinningStatus()
					break
				}
			}
		}
	}()
	go func(cctx context.Context, ch chan error) {
		ch <- TryToScanAllSectors(ctx)
	}(customCtx, ch)


	for {
		select {
		case err := <-ch:
			return err
		case <-customCtx.Done():
			log.Infof("already cancel")
			return nil
		}
	}
}

func TryToScanAllSectors(ctx context.Context) error {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("into TryToScanAllSectors, get lock,will sleep")
	return nil
}


func NeedLockRun(c int) {
	PL.Lock.Lock()
	defer PL.Lock.Unlock()
	fmt.Println("into NeedLockRun, get lock,will sleep:", c)
	time.Sleep(time.Millisecond * 10)
}

func TestTicket() {
	workTicket := time.NewTicker(2*time.Second)
	defer workTicket.Stop()
	count := 0
	for range workTicket.C {

		count++
		if count == 5 {
			continue
		}
		time.Sleep(3*time.Second)
		fmt.Print(count, ":")
		fmt.Println(time.Now())
	}
}


func TryDeclare() error {
	defer log.Infof("will release localLK")

	ch := make(chan error, 1)

	customCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(cctx context.Context) {
		winningStatusTicker := time.NewTicker(time.Second * 5)
		defer winningStatusTicker.Stop()
		for {
			select {
			case <-winningStatusTicker.C:
				time.Sleep(3 * time.Second)
					log.Infof("now is winning, break tryDeclare, release lock")
					ch <- nil
					return
			case <-cctx.Done():
				return
			}
		}
	}(customCtx)

	go func(ch chan error) {
		time.Sleep(time.Second * 10)
		ch <- nil
	}(ch)

	err := <-ch
	log.Infof("indirect into...")
	return err
}