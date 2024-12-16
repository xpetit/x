package x

import (
	"cmp"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"
	"text/tabwriter"
	"time"
)

var StartTime = time.Now().UTC()

func DebugInfos() string {
	bi := AssertOK(debug.ReadBuildInfo())
	var gcs debug.GCStats
	var ms runtime.MemStats
	debug.ReadGCStats(&gcs)
	runtime.ReadMemStats(&ms)
	var sb strings.Builder
	tw := tabwriter.NewWriter(&sb, 0, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(tw, "PID:\t", os.Getpid())
	fmt.Fprintln(tw, "CPUs:\t", runtime.NumCPU())
	fmt.Fprintln(tw, "Goroutines:\t", runtime.NumGoroutine())
	fmt.Fprintln(tw, "CGo calls:\t", runtime.NumCgoCall())
	fmt.Fprintln(tw, "Go version:\t", bi.GoVersion)
	fmt.Fprintln(tw, "Path:\t", bi.Path)
	fmt.Fprintf(tw, "Start:\t %s (running for %s)\n", StartTime.Format(time.DateTime), time.Since(StartTime).Round(time.Second))
	fmt.Fprintf(tw, "Memory in use:\t %.fMB\n", float64(ms.StackInuse+ms.HeapIdle+ms.HeapInuse-ms.HeapReleased)/1e6)
	fmt.Fprintln(tw)
	fmt.Fprintln(tw, "Build settings:")
	Check(tw.Flush())
	{
		tw := tabwriter.NewWriter(tw, 0, 0, 1, ' ', 0)
		for _, setting := range bi.Settings {
			fmt.Fprintf(tw, "  %s:\t%s\n", setting.Key, setting.Value)
		}
		Check(tw.Flush())
	}
	if gcs.NumGC > 0 {
		fmt.Fprintln(tw)
		fmt.Fprintln(tw, "GC stats:")
		fmt.Fprintln(tw, "Last GC:\t", gcs.LastGC.UTC().Format(time.DateTime))
		fmt.Fprintln(tw, "Num GC:\t", gcs.NumGC)
		fmt.Fprintln(tw, "Total pauses:\t", gcs.PauseTotal)
		if len(gcs.Pause) > 0 {
			fmt.Fprintln(tw, "Worst pauses:\t")
			indexes := make([]int, len(gcs.Pause))
			for i := range indexes {
				indexes[i] = i
			}
			slices.SortFunc(indexes, func(a, b int) int { return cmp.Compare(gcs.Pause[b], gcs.Pause[a]) })
			indexes = indexes[:min(6, len(indexes))]
			for _, i := range indexes {
				fmt.Fprintf(tw, "\t0.%06ds  %s\t\n", gcs.Pause[i].Microseconds(), gcs.PauseEnd[i].UTC().Format(time.DateTime))
			}
		}
	}
	Check(tw.Flush())
	return sb.String()
}
