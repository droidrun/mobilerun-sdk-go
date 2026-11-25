// Command objects shows the object API: create a device, wait until it is
// ready, run one action, terminate it, and list devices across all pages.
//
// It needs MOBILERUN_API_KEY and creates a billed device; read it before running.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/droidrun/mobilerun-sdk-go/v6"
	"github.com/droidrun/mobilerun-sdk-go/v6/sdk"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	client := sdk.NewMobilerunSDK()

	// New waits until the device is ready; a failed wait disposes it and
	// returns an *sdk.ObjectReadyError.
	device, err := client.Devices.New(ctx, mobilerun.DeviceNewParams{}, sdk.WithReadyTimeout(10*time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	// Close terminates the device because this handle created it.
	defer device.Close(context.WithoutCancel(ctx))
	fmt.Println("ready:", device, device.State())

	shot, err := device.Screenshot(ctx, mobilerun.DeviceScreenshotCaptureParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("screenshot bytes:", len(shot.Bytes), shot.ContentType)

	task, err := device.Run(ctx, mobilerun.TaskRunParams{Task: "Open the settings app"})
	if err != nil {
		log.Fatal(err)
	}
	status, err := task.Wait(ctx, sdk.WithWaitTimeout(5*time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("task:", status.Status)

	if err := device.Terminate(ctx, mobilerun.DeviceTerminateParams{}); err != nil {
		log.Fatal(err)
	}

	devices := client.Devices.ListAutoPaging(ctx, mobilerun.DeviceListParams{})
	for devices.Next() {
		fmt.Println(devices.Current())
	}
	if err := devices.Err(); err != nil {
		log.Fatal(err)
	}
}
