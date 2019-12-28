package vm

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whereiskurt/tiogo/pkg/client"
	"github.com/whereiskurt/tiogo/pkg/ui"
)

// ExportScansStart request Tenable.io create a new export for the scan
func (vm *VM) ExportScansStart(cmd *cobra.Command, args []string) {
	log := vm.Config.VM.EnableLogging()

	a := client.NewAdapter(vm.Config, vm.Metrics)

	scans, err := a.Scans(true, true)
	if err != nil {
		log.Errorf("error: couldn't scans list: %v", err)
		return
	}
	scans = vm.FilterScans(a, &scans)

	histid := vm.Config.VM.HistoryID

	if len(scans) == 0 {
		log.Errorf("error: history id didn't match a scan")
	} else if len(scans) > 1 && histid != "" {
		log.Errorf("error: histid doesn't limit to one scan")
	}

	for _, s := range scans {
		// If we need to get the latest histid
		if histid == "" {
			det, err := a.ScanDetails(&s, true, true)
			if err != nil {
				continue
			}

			// This scan has no history ie. no previous scans
			if len(det.History) == 0 {
				continue
			}
			histid = det.History[0].HistoryID
		}

		// TODO: Pass in 2x code{} for cmd, and output
		var export client.ScansExportStart
		export, err := a.ScansExportStart(&s, histid, true, true)
		if err != nil {
			log.Errorf("error: couldn't start export-scans: %v", err)
			continue
		}

		cli := ui.NewCLI(vm.Config)
		cli.Println(fmt.Sprintf("tiogo version %s (%s)", ReleaseVersion, GitHash))
		cli.DrawGopher()
		fmt.Println(cli.Render("ExportScansStart", map[string]string{"FileUUID": export.FileUUID, "ScanID": s.ScanID, "HistoryID": histid}))

	}
	return
}

// ExportScansStatus check export status of Tenable.io scan export
func (vm *VM) ExportScansStatus(cmd *cobra.Command, args []string) {
	log := vm.Config.VM.EnableLogging()

	a := client.NewAdapter(vm.Config, vm.Metrics)

	scans, err := a.Scans(true, true)
	if err != nil {
		log.Errorf("error: couldn't scans list: %v", err)
		return
	}
	scans = vm.FilterScans(a, &scans)

	histid := vm.Config.VM.HistoryID

	if len(scans) == 0 {
		log.Errorf("error: history id didn't match a scan")
	} else if len(scans) > 1 && histid != "" {
		log.Errorf("error: histid doesn't limit to one scan")
	}

	for _, s := range scans {
		// If we need to get the latest histid
		if histid == "" {
			det, err := a.ScanDetails(&s, true, true)
			if err != nil {
				continue
			}

			// This scan has no history ie. no previous scans
			if len(det.History) == 0 {
				continue
			}
			histid = det.History[0].HistoryID
		}

		// TODO: Pass in 2x code{} for cmd, and output
		var export client.ScansExportStatus
		export, err := a.ScansExportStatus(&s, histid, true, true)
		if err != nil {
			log.Errorf("error: couldn't start export-scans: %v", err)
			continue
		}

		cli := ui.NewCLI(vm.Config)
		cli.Println(fmt.Sprintf("tiogo version %s (%s)", ReleaseVersion, GitHash))
		cli.DrawGopher()
		fmt.Println(cli.Render("ExportScansStatus", map[string]string{"FileUUID": export.FileUUID, "Status": export.Status, "ScanID": s.ScanID, "HistoryID": histid}))

	}
	return
}

// ExportScansGet download the exported status of Tenable.io scan export
func (vm *VM) ExportScansGet(cmd *cobra.Command, args []string) {

}

// ExportScansQuery download the exported status of Tenable.io scan export
func (vm *VM) ExportScansQuery(cmd *cobra.Command, args []string) {

}

// ExportScansHelp outputs the cli help export-scans command
func (vm *VM) ExportScansHelp(cmd *cobra.Command, args []string) {
	cli := ui.NewCLI(vm.Config)

	cli.Println(fmt.Sprintf("tiogo version %s (%s)", ReleaseVersion, GitHash))
	cli.Println(cli.Render("exportScansUsage", nil))

	return
}
