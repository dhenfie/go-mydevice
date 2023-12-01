package main

import (
	"fmt"
	"my-device/command"
	"my-device/devices"
)

func main() {

	command.Handler("os", func() {
		os, err := devices.NewOperatingSystem()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("--- Sistem Operasi ---")
		fmt.Println("Nama:", os.Name)
		fmt.Println("Kernel:", os.Kernel)
		fmt.Println("Kernel Release:", os.KernelRelease)
		fmt.Println("Kernel Version:", os.KernelVersion)
		fmt.Println("Mesin:", os.Machine)
	})

	command.Handler("memory", func() {
		memory, err := devices.NewMemory()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("--- Memori Sistem ---")
		fmt.Println("Total RAM:", memory.Ram.Total)
		fmt.Println("Sisa RAM:", memory.Ram.Free)
		fmt.Println("Penggunaan RAM:", memory.Ram.Usage)
		fmt.Println("Total Swap:", memory.Swap.Total)
		fmt.Println("Sisa Swap:", memory.Swap.Free)
		fmt.Println("Penggunaan Swap:", memory.Swap.Usage)
	})

	command.Run()
}
