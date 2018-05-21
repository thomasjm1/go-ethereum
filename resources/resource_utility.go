package resources

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/shirou/gopsutil/process"
	"encoding/json"
	"fmt"
	"strings"
	//"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/cpu"
)

type ProcessInfo struct {
	CpuPercent float64 `json:"cpuPercent"`
	//IOCounters
	ReadCount  uint64 `json:"readCount"`
	WriteCount uint64 `json:"writeCount"`
	ReadBytes  uint64 `json:"readBytes"`
	WriteBytes uint64 `json:"writeBytes"`
	//MemoryInfo
	RSS    uint64 `json:"rss"`    // bytes
	VMS    uint64 `json:"vms"`    // bytes
	Shared uint64 `json:"shared"` // bytes
	Text   uint64 `json:"text"`   // bytes
	Lib    uint64 `json:"lib"`    // bytes
	Data   uint64 `json:"data"`   // bytes
	Dirty  uint64 `json:"dirty"`  // bytes
	//NetIOCounters
	//Name        string `json:"name"`        // interface name
	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
	Errin       uint64 `json:"errin"`       // total number of errors while receiving
	Errout      uint64 `json:"errout"`      // total number of errors while sending
	Dropin      uint64 `json:"dropin"`      // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`     // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
	Fifoin      uint64 `json:"fifoin"`      // total number of FIFO buffers errors while receiving
	Fifoout     uint64 `json:"fifoout"`     // total number of FIFO buffers errors while sending
	//Times
	CPU       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guestNice"`
	Stolen    float64 `json:"stolen"`
	//Total Times
	TotalCPU       string  `json:"total_cpu"`
	TotalUser      float64 `json:"total_user"`
	TotalSystem    float64 `json:"total_system"`
	TotalIdle      float64 `json:"total_idle"`
	TotalNice      float64 `json:"total_nice"`
	TotalIowait    float64 `json:"total_iowait"`
	TotalIrq       float64 `json:"total_irq"`
	TotalSoftirq   float64 `json:"total_softirq"`
	TotalSteal     float64 `json:"total_steal"`
	TotalGuest     float64 `json:"total_guest"`
	TotalGuestNice float64 `json:"total_guestNice"`
	TotalStolen    float64 `json:"total_stolen"`
	//Added
	Properties map[string]string `json:"properties"`
}

func (d ProcessInfo) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func RecordResourcesToLog(prefix string, properties map[string]string) {
	//log.Error(fmt.Sprintf("[thomasjm] - Resource usage due to %s", prefix))
	var processes, _ = process.Processes()
	for index := range processes {
		processInfo := processes[index]
		processName, _ := processInfo.Name()
		if strings.Compare(processName, "geth") == 0 {
			cpuPercent, _ := processInfo.CPUPercent()
			ioCounters, _ := processInfo.IOCounters()
			memoryInfo, _ := processInfo.MemoryInfoEx()
			//Combines all network interface stats if false, true then return individual network interfaces
			netIoCounters, _ := processInfo.NetIOCounters(false)
			netIoCounter := netIoCounters[0]
			times, _ := processInfo.Times()
			totalTimes, _ := cpu.Times(false)
			summedTimes :=cpu.TimesStat{
				CPU: "ALL",
				User: 0,
				System: 0,
				Idle: 0,
				Nice: 0,
				Iowait: 0,
				Irq: 0,
				Softirq: 0,
				Steal: 0,
				Guest: 0,
				GuestNice: 0,
				Stolen: 0,
			}
			for totalTimeIndex := range totalTimes {
				totalTime := totalTimes[totalTimeIndex]
				summedTimes.User = summedTimes.User + totalTime.User
				summedTimes.System = summedTimes.System + totalTime.System
				summedTimes.Idle = summedTimes.Idle + totalTime.Idle
				summedTimes.Nice = summedTimes.Nice + totalTime.Nice
				summedTimes.Iowait = summedTimes.Iowait + totalTime.Iowait
				summedTimes.Irq = summedTimes.Irq + totalTime.Irq
				summedTimes.Softirq = summedTimes.Softirq + totalTime.Softirq
				summedTimes.Steal = summedTimes.Steal + totalTime.Steal
				summedTimes.Guest = summedTimes.Guest + totalTime.Guest
				summedTimes.GuestNice = summedTimes.GuestNice + totalTime.GuestNice
				summedTimes.Stolen = summedTimes.Stolen + totalTime.Stolen
			}
			processOuput := ProcessInfo{
				CpuPercent: cpuPercent,
				ReadCount: ioCounters.ReadCount,
				WriteCount: ioCounters.WriteCount,
				ReadBytes: ioCounters.ReadBytes,
				WriteBytes: ioCounters.WriteBytes,
				RSS: memoryInfo.RSS,
				VMS: memoryInfo.VMS,
				Shared: memoryInfo.Shared,
				Text: memoryInfo.Text,
				Lib: memoryInfo.Lib,
				Data: memoryInfo.Data,
				Dirty: memoryInfo.Dirty,
				BytesSent: netIoCounter.BytesSent,
				BytesRecv: netIoCounter.BytesRecv,
				PacketsSent:netIoCounter.PacketsSent,
				PacketsRecv: netIoCounter.PacketsRecv,
				Errin: netIoCounter.Errin,
				Errout: netIoCounter.Errout,
				Dropin: netIoCounter.Dropin,
				Dropout: netIoCounter.Dropout,
				Fifoin: netIoCounter.Fifoin,
				Fifoout: netIoCounter.Fifoout,
				CPU: times.CPU,
				User: times.User,
				System: times.System,
				Idle: times.Idle,
				Nice: times.Nice,
				Iowait: times.Iowait,
				Irq: times.Irq,
				Softirq: times.Softirq,
				Steal: times.Steal,
				Guest: times.Guest,
				GuestNice: times.GuestNice,
				Stolen: times.Stolen,
				TotalCPU: summedTimes.CPU,
				TotalUser: summedTimes.User,
				TotalSystem: summedTimes.System,
				TotalIdle: summedTimes.Idle,
				TotalNice: summedTimes.Nice,
				TotalIowait: summedTimes.Iowait,
				TotalIrq: summedTimes.Irq,
				TotalSoftirq: summedTimes.Softirq,
				TotalSteal: summedTimes.Steal,
				TotalGuest: summedTimes.Guest,
				TotalGuestNice: summedTimes.GuestNice,
				TotalStolen: summedTimes.Stolen,
				Properties:properties,
			}
			log.Error(fmt.Sprintf("[thomasjm] - (%s) => %s", prefix, processOuput))
		}
	}
}
