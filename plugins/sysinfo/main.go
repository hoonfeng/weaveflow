package main

import (
  "bufio"
  "github.com/hoonfeng/goproc/sdk"
  "golang.org/x/sys/windows"
  "os"
  "os/exec"
  "path/filepath"
  "strings"
  "fmt"
  "syscall"
  "unsafe"
)

type fileTime struct{ LowDateTime, HighDateTime uint32 }
var prevIdle, prevKernel, prevUser fileTime
var k32 = syscall.NewLazyDLL("kernel32")
var procGetSystemTimes = k32.NewProc("GetSystemTimes")
var procGlobalMemoryStatusEx = k32.NewProc("GlobalMemoryStatusEx")

func cpuPercent() float64 {
  var idle, kernel, user fileTime
  procGetSystemTimes.Call(uintptr(unsafe.Pointer(&idle)), uintptr(unsafe.Pointer(&kernel)), uintptr(unsafe.Pointer(&user)))
  pi := uint64(idle.HighDateTime)<<32 | uint64(idle.LowDateTime)
  pk := uint64(kernel.HighDateTime)<<32 | uint64(kernel.LowDateTime)
  pu := uint64(user.HighDateTime)<<32 | uint64(user.LowDateTime)
  if (prevIdle.HighDateTime|prevIdle.LowDateTime|prevKernel.HighDateTime|prevKernel.LowDateTime|prevUser.HighDateTime|prevUser.LowDateTime) == 0 {
    prevIdle, prevKernel, prevUser = idle, kernel, user
    return 0
  }
  ppi := uint64(prevIdle.HighDateTime)<<32 | uint64(prevIdle.LowDateTime)
  ppk := uint64(prevKernel.HighDateTime)<<32 | uint64(prevKernel.LowDateTime)
  ppu := uint64(prevUser.HighDateTime)<<32 | uint64(prevUser.LowDateTime)
  di := pi - ppi
  dk := pk - ppk
  du := pu - ppu
  prevIdle, prevKernel, prevUser = idle, kernel, user
  total := di + dk + du
  busy := total - di
  if total == 0 {
    return 0
  }
  return float64(busy) * 100.0 / float64(total)
}

func memoryStatus() (used uint64, total uint64) {
  type memStatusEx struct{
    dwLength uint32
    dwMemoryLoad uint32
    ullTotalPhys uint64
    ullAvailPhys uint64
    ullTotalPageFile uint64
    ullAvailPageFile uint64
    ullTotalVirtual uint64
    ullAvailVirtual uint64
    ullAvailExtendedVirtual uint64
  }
  var ms memStatusEx
  ms.dwLength = uint32(unsafe.Sizeof(ms))
  r, _, _ := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&ms)))
  if r == 0 { return 0, 0 }
  total = ms.ullTotalPhys
  avail := ms.ullAvailPhys
  used = total - avail
  return
}

func diskSpaceRoot() (uint64, uint64) {
  wd, _ := os.Getwd()
  vol := filepath.VolumeName(wd)
  if vol == "" {
    vol = `C:`
  }
  root := vol + `\\`
  p, _ := windows.UTF16PtrFromString(root)
  var free, total, freeAll uint64
  if err := windows.GetDiskFreeSpaceEx(p, &free, &total, &freeAll); err != nil {
    return 0, 0
  }
  return total, freeAll
}

func resources(params map[string]interface{}) (interface{}, error) {
  cpu := cpuPercent()
  mu, mt := memoryStatus()
  st, sf := diskSpaceRoot()
  gpus := make([]map[string]interface{}, 0)
  // Skip WMI GPU controller collection; rely on NVIDIA nvidia-smi for accurate per-GPU metrics
  // NVIDIA per-GPU details via nvidia-smi
  if _, err := exec.LookPath("nvidia-smi"); err == nil {
    cmd2 := exec.Command("nvidia-smi", "--query-gpu=name,utilization.gpu,memory.total,memory.used", "--format=csv,noheader,nounits")
    out2, err2 := cmd2.Output()
    if err2 == nil {
      s2 := bufio.NewScanner(strings.NewReader(string(out2)))
      for s2.Scan() {
        line := strings.TrimSpace(s2.Text())
        if line == "" { continue }
        parts := strings.Split(line, ",")
        if len(parts) >= 4 {
          name := strings.TrimSpace(parts[0])
          var util float64
          var mtMB, muMB uint64
          fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &util)
          fmt.Sscanf(strings.TrimSpace(parts[2]), "%d", &mtMB)
          fmt.Sscanf(strings.TrimSpace(parts[3]), "%d", &muMB)
          mt := mtMB * 1024 * 1024
          mu := muMB * 1024 * 1024
          gpus = append(gpus, map[string]interface{}{
            "name": name,
            "memory_total_bytes": mt,
            "memory_used_bytes": mu,
            "compute_percent": util,
          })
        }
      }
    }
  }
  // merge duplicates by name, prefer entries with non-zero used and larger total
  mergedMap := make(map[string]map[string]interface{})
  for _, gi := range gpus {
    name := strings.ToLower(strings.TrimSpace(fmt.Sprintf("%v", gi["name"])))
    if name == "" {
      continue
    }
    mx := uint64(0)
    mu := uint64(0)
    util := float64(0)
    switch v := gi["memory_total_bytes"].(type) {
    case uint64:
      mx = v
    case int:
      if v > 0 { mx = uint64(v) }
    case int64:
      if v > 0 { mx = uint64(v) }
    }
    switch v := gi["memory_used_bytes"].(type) {
    case uint64:
      mu = v
    case int:
      if v > 0 { mu = uint64(v) }
    case int64:
      if v > 0 { mu = uint64(v) }
    }
    if u, ok := gi["compute_percent"].(float64); ok { util = u }
    cur := mergedMap[name]
    if cur == nil {
      mergedMap[name] = map[string]interface{}{
        "name": gi["name"],
        "memory_total_bytes": mx,
        "memory_used_bytes": mu,
        "compute_percent": util,
      }
      continue
    }
    // choose best
    cx := uint64(0)
    cu := uint64(0)
    if v, ok := cur["memory_total_bytes"].(uint64); ok { cx = v }
    if v, ok := cur["memory_used_bytes"].(uint64); ok { cu = v }
    // prefer non-zero used
    if mu > cu || (mu == cu && mx > cx) {
      cur["memory_total_bytes"] = mx
      cur["memory_used_bytes"] = mu
    }
    if util > 0 {
      cur["compute_percent"] = util
    }
    mergedMap[name] = cur
  }
  merged := make([]map[string]interface{}, 0, len(mergedMap))
  for _, v := range mergedMap {
    merged = append(merged, v)
  }
  // Network throughput
  var rxBps, txBps uint64
  ncmd := exec.Command("wmic", "path", "Win32_PerfFormattedData_Tcpip_NetworkInterface", "get", "Name,BytesReceivedPersec,BytesSentPersec", "/format:list")
  nout, nerr := ncmd.Output()
  if nerr == nil {
    ns := bufio.NewScanner(strings.NewReader(string(nout)))
    var name string
    var rx, tx uint64
    for ns.Scan() {
      line := strings.TrimSpace(ns.Text())
      if strings.HasPrefix(line, "Name=") {
        name = strings.TrimSpace(strings.TrimPrefix(line, "Name="))
      } else if strings.HasPrefix(line, "BytesReceivedPersec=") {
        fmt.Sscanf(strings.TrimSpace(strings.TrimPrefix(line, "BytesReceivedPersec=")), "%d", &rx)
      } else if strings.HasPrefix(line, "BytesSentPersec=") {
        fmt.Sscanf(strings.TrimSpace(strings.TrimPrefix(line, "BytesSentPersec=")), "%d", &tx)
      } else if line == "" {
        if name != "" && !strings.Contains(strings.ToLower(name), "loopback") && !strings.Contains(strings.ToLower(name), "isatap") {
          rxBps += rx
          txBps += tx
        }
        name = ""
        rx = 0
        tx = 0
      }
    }
  }
  // Disk IO throughput
  var rdBps, wrBps uint64
  dcmd := exec.Command("wmic", "path", "Win32_PerfFormattedData_PerfDisk_PhysicalDisk", "get", "Name,DiskReadBytesPersec,DiskWriteBytesPersec", "/format:list")
  dout, derr := dcmd.Output()
  if derr == nil {
    ds := bufio.NewScanner(strings.NewReader(string(dout)))
    var name string
    var r, w uint64
    for ds.Scan() {
      line := strings.TrimSpace(ds.Text())
      if strings.HasPrefix(line, "Name=") {
        name = strings.TrimSpace(strings.TrimPrefix(line, "Name="))
      } else if strings.HasPrefix(line, "DiskReadBytesPersec=") {
        fmt.Sscanf(strings.TrimSpace(strings.TrimPrefix(line, "DiskReadBytesPersec=")), "%d", &r)
      } else if strings.HasPrefix(line, "DiskWriteBytesPersec=") {
        fmt.Sscanf(strings.TrimSpace(strings.TrimPrefix(line, "DiskWriteBytesPersec=")), "%d", &w)
      } else if line == "" {
        if name == "_Total" || name != "" {
          rdBps += r
          wrBps += w
        }
        name = ""
        r = 0
        w = 0
      }
    }
  }
  // GPU engine utilization aggregate from nvidia-smi entries
  var gpuUtil float64
  // GPU memory usage aggregate from nvidia-smi entries
  var gpuMemUsedBytes uint64
  // Aggregate GPU total memory from controllers if perf memory not available
  var gpuMemTotalBytes uint64
  filtered := make([]map[string]interface{}, 0, len(merged))
  for _, gi := range merged {
    nameVal, _ := gi["name"].(string)
    mt := uint64(0)
    switch v := gi["memory_total_bytes"].(type) {
    case uint64:
      mt = v
    case int:
      if v > 0 { mt = uint64(v) }
    case int64:
      if v > 0 { mt = uint64(v) }
    }
    lower := strings.ToLower(nameVal)
    isVirtual := mt == 0 || strings.Contains(lower, "orayidd") || strings.Contains(lower, "displaylink") || strings.Contains(lower, "basic display") || strings.Contains(lower, "mirage") || strings.Contains(lower, "virtual") || strings.Contains(lower, "hyper-v")
    if isVirtual {
      continue
    }
    filtered = append(filtered, gi)
  }
  if len(filtered) == 0 {
    filtered = merged
  }
  for _, gi := range filtered {
    if v, ok := gi["memory_total_bytes"].(uint64); ok { gpuMemTotalBytes += v } else if v2, ok2 := gi["memory_total_bytes"].(int); ok2 { gpuMemTotalBytes += uint64(v2) }
    if vu, ok := gi["memory_used_bytes"].(uint64); ok { gpuMemUsedBytes += vu } else if vu2, ok22 := gi["memory_used_bytes"].(int); ok22 { gpuMemUsedBytes += uint64(vu2) }
    if cu, ok := gi["compute_percent"].(float64); ok { gpuUtil += cu }
  }
  if len(filtered) > 0 && gpuUtil > 0 {
    gpuUtil = gpuUtil / float64(len(filtered))
    if gpuUtil > 100 { gpuUtil = 100 }
  }
  res := map[string]interface{}{
    "memory": map[string]interface{}{"used_bytes": mu, "total_bytes": mt},
    "cpu": map[string]interface{}{"percent": cpu},
    "storage": map[string]interface{}{"total_bytes": st, "free_bytes": sf},
    "network": map[string]interface{}{"bytes_sent": txBps, "bytes_recv": rxBps},
    "gpu": map[string]interface{}{"percent": gpuUtil, "memory_total_bytes": gpuMemTotalBytes, "memory_used_bytes": gpuMemUsedBytes, "compute_percent": gpuUtil},
    "gpus": filtered,
  }
  res["cpu_percent"] = cpu
  res["cpu_percent_str"] = fmt.Sprintf("%.1f", cpu)
  res["cpu_value"] = cpu
  res["storage_io"] = map[string]interface{}{"read_bps": rdBps, "write_bps": wrBps}
  return res, nil
}

func main() {
  _ = sdk.RegisterFunction("resources", resources)
  _ = sdk.Start()
  sdk.Wait()
}