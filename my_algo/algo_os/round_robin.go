package algo_os

import "fmt"

const (
	READY  = "READY"
	RUN    = "RUN"
	BLOCK  = "BLOCK"
	FINISH = "FINISH"
)

// PCB : RUN READY BLOCK FINISH
type PCB struct {
	ID         int    // 进程号
	Priority   int    // 优先权,越大越先执行
	CpuTime    int    // cpu上已经运行时长
	AllTime    int    // cpu还需将要占用时长，完毕变为0
	StartBlock int    // 进程阻塞时间,再运行此时间将被阻塞
	BlockTime  int    // 已经阻塞的进程再等待此时间,进入就绪状态
	State      string // 当前状态
}

func ConstructProcess(id, prio, cpu, all, stbl, block int, state string) *PCB {
	return &PCB{
		ID:         id,
		Priority:   prio,
		CpuTime:    cpu,
		AllTime:    all,
		StartBlock: stbl,
		BlockTime:  block,
		State:      state,
	}
}

func InitProcess() []*PCB {
	head := make([]*PCB, 5)
	head[0] = ConstructProcess(0, 9, 0, 3, 2, 3, READY)
	head[1] = ConstructProcess(1, 38, 0, 3, -1, 0, READY)
	head[2] = ConstructProcess(2, 30, 0, 6, -1, 0, READY)
	head[3] = ConstructProcess(3, 29, 0, 3, -1, 0, READY)
	head[4] = ConstructProcess(4, 0, 0, 4, -1, 0, READY)
	return head
}

func PrintPCBsId(pcbs []*PCB, index []int) {
	for i := 0; i < len(index); i++ {
		fmt.Printf("->%d", pcbs[index[i]].ID)
	}
	fmt.Println()
}

func PrintPCBsInfo(pcbs []*PCB) {
	fmt.Println("ID PRIORITY CPUTIME ALLTIME STARTBLOCK BLOCKTIME STATE")
	for i := 0; i < len(pcbs); i++ {
		fmt.Println(pcbs[i].ID,
			"     ", pcbs[i].Priority,
			"     ", pcbs[i].CpuTime,
			"     ", pcbs[i].AllTime,
			"     ", pcbs[i].StartBlock,
			"     ", pcbs[i].BlockTime,
			"     ", pcbs[i].State)
	}
	fmt.Println("=====================================================")
}

func SelectMaxPriority(pcbs []*PCB) int {
	id, max := -1, -0xffffff
	for i := 0; i < len(pcbs); i++ {
		if pcbs[i].State == RUN || pcbs[i].State == READY {
			if pcbs[i].StartBlock != 0 && pcbs[i].Priority > max {
				id = i
				max = pcbs[i].Priority
			}
		}
	}
	return id
}
func SelectReady(pcbs []*PCB, runId int) []int {
	ready := make([]int, 0)
	for i := 0; i < len(pcbs); i++ {
		if pcbs[i].State == RUN || pcbs[i].State == READY {
			if i != runId {
				ready = append(ready, i)
			}
		}
	}
	return ready
}
func SelectBlock(pcbs []*PCB) []int {
	block := make([]int, 0)
	for i := 0; i < len(pcbs); i++ {
		if pcbs[i].State == BLOCK {
			block = append(block, i)
		}
	}
	return block
}

// 进程在就绪队列一个时间片,优先级加一;进程运行一个时间片,优先级-3

func RunScheduling(pcbs []*PCB) {
	totTime, totFinish := 0, 0
	for totFinish != 5 {
		runId := SelectMaxPriority(pcbs)
		ready := SelectReady(pcbs, runId)
		block := SelectBlock(pcbs)
		for i := 0; i < len(block); i++ {
			pcbs[block[i]].StartBlock = -1
			pcbs[block[i]].BlockTime--
			if pcbs[block[i]].BlockTime == 0 {
				ready = append(ready, block[i])
			}
		}
		for i := 0; i < len(ready); i++ {
			pcbs[ready[i]].State = READY
		}
		if runId < 0 {
			break
		}
		fmt.Println("RUNNING PROG:", runId)
		pcbs[runId].State = RUN
		pcbs[runId].AllTime--
		pcbs[runId].CpuTime++
		stbl := pcbs[runId].StartBlock
		if pcbs[runId].AllTime == 0 {
			pcbs[runId].State = FINISH
			totFinish++
		}
		if stbl == -1 {
			// Always Run
		} else {
			pcbs[runId].StartBlock--
			stbl--
			if stbl == 0 {
				pcbs[runId].State = BLOCK
			}
		}
		totTime++
		fmt.Print("READY_QUEUE:")
		PrintPCBsId(pcbs, ready)
		fmt.Print("BLOCK_QUEUE:")
		PrintPCBsId(pcbs, block)
		PrintPCBsInfo(pcbs)
		pcbs[runId].Priority -= 3
	}
	fmt.Println("total runtime = ", totTime)
}
