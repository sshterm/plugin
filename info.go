package plugin

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

type info struct {
}

func NewInfo() *info {
	return &info{}
}

type InfoStat struct {
	CPU         int32    `json:"cpu"`
	VendorID    string   `json:"vendorId"`
	Family      string   `json:"family"`
	Model       string   `json:"model"`
	Stepping    int32    `json:"stepping"`
	PhysicalID  string   `json:"physicalId"`
	CoreID      string   `json:"coreId"`
	Cores       int32    `json:"cores"`
	ModelName   string   `json:"modelName"`
	Mhz         float64  `json:"mhz"`
	CacheSize   int32    `json:"cacheSize"`
	Flags       []string `json:"flags"`
	Microcode   string   `json:"microcode"`
	UsedPercent float64  `json:"usedPercent"`
}

func (i *info) Cpu() (data []InfoStat) {
	if percent, err := cpu.Percent(0, true); err == nil {
		if memory, err := cpu.Info(); err == nil {
			for index, stat := range memory {
				data = append(data, InfoStat{
					CPU:         stat.CPU,
					VendorID:    stat.VendorID,
					Family:      stat.Family,
					Model:       stat.Model,
					Stepping:    stat.Stepping,
					PhysicalID:  stat.PhysicalID,
					CoreID:      stat.CoreID,
					Cores:       stat.Cores,
					ModelName:   stat.ModelName,
					Mhz:         stat.Mhz,
					CacheSize:   stat.CacheSize,
					Flags:       stat.Flags,
					Microcode:   stat.Microcode,
					UsedPercent: percent[index],
				})
			}
		}
	}
	return
}
func (i *info) Avg() (data *load.AvgStat) {
	if memory, err := load.Avg(); err == nil {
		data = memory
	}
	return
}

func (i *info) Disk() (data []*disk.UsageStat) {
	if disks, err := disk.Partitions(true); err == nil {
		for _, stat := range disks {
			if d, err := disk.Usage(stat.Mountpoint); err == nil {
				if d.Fstype == "" {
					d.Fstype = stat.Fstype
				}
				data = append(data, d)
			}
		}
	}
	return
}

func (i *info) Host() (data *host.InfoStat) {
	if memory, err := host.Info(); err == nil {
		data = memory
	}
	return
}
func (i *info) Net() (data []net.InterfaceStat) {
	if memory, err := net.Interfaces(); err == nil {
		data = memory
	}
	return
}
func (i *info) Mem() (data *mem.VirtualMemoryStat) {
	if memory, err := mem.VirtualMemory(); err == nil {
		data = memory
	}
	return
}

/*
	func (i *info) Mem() (data *mem.VirtualMemoryStat) {
		if memory, err := mem.VirtualMemory(); err == nil {
			data = memory
		}
		return
	}
*/
func (i *info) Version() string {
	return "1.0.0"
}
func (i *info) ioDisk() (Read, Write uint64) {
	per2, err := disk.IOCounters()
	if err != nil {
		return
	}
	time.Sleep(time.Second)
	per, err := disk.IOCounters()
	if err != nil {
		return
	}
	var read, write, read2, write2 uint64
	for _, p := range per {
		read += p.ReadBytes
		write += p.WriteBytes
	}
	for _, p := range per2 {
		read2 += p.ReadBytes
		write2 += p.WriteBytes
	}
	Read = read - read2
	Write = write - write2
	return
}
func (i *info) ioNet() (Recv, Sent uint64) {
	per2, err := net.IOCounters(true)
	if err != nil {
		return
	}
	time.Sleep(time.Second)
	per, err := net.IOCounters(true)
	if err != nil {
		return
	}
	var recv, sent, recv2, sent2 uint64
	for _, p := range per {
		recv += p.BytesRecv
		sent += p.BytesSent
	}
	for _, p := range per2 {
		recv2 += p.BytesRecv
		sent2 += p.BytesSent
	}
	Recv = recv - recv2
	Sent = sent - sent2
	return
}

func (i *info) ioAvg() (data *load.AvgStat) {
	if avg, err := load.Avg(); err == nil {
		data = avg
	}
	return
}
func (i *info) ioSwapMemory() (data float64) {
	if memory, err := mem.VirtualMemory(); err == nil {
		data = memory.UsedPercent
	}
	return
}
func (i *info) ioCPUUsedPercent() (percent float64) {
	if percents, err := cpu.Percent(0, true); err == nil {
		var sum float64
		for _, f := range percents {
			sum += f
		}
		percent = float64(len(percents)) / sum
	}
	return
}
func (i *info) Metrics() (data Metrics) {
	data.AvgStat = i.ioAvg()
	data.MemoryUsedPercent = i.ioSwapMemory()
	data.Recv, data.Sent = i.ioNet()
	data.Read, data.Write = i.ioDisk()
	data.CPUUsedPercent = i.ioCPUUsedPercent()
	return
}

type Metrics struct {
	AvgStat *load.AvgStat `json:"avg"`
	//SwapMemoryStat    *mem.SwapMemoryStat `json:"memory"`
	//CPU               []float64 `json:"cpu"`
	CPUUsedPercent    float64 `json:"cpu_used_percent"`
	MemoryUsedPercent float64 `json:"memory_used_percent"`
	//	Net            []net.IOCountersStat `json:"net"`
	Sent  uint64 `json:"sent"`
	Recv  uint64 `json:"recv"`
	Read  uint64 `json:"read"`
	Write uint64 `json:"write"`
}
