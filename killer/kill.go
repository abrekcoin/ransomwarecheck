package killer
import (
	"fmt"
	"log"
	"github.com/abdfnx/gosh"
)

func Kill(process string) {
	err, out, errout := gosh.RunOutput("Get-EventLog -log Security -InstanceId 4663 -Newest 1 | Where-Object {$_.message -match \"(?s)Object Name:\\s+(?<ObjectName>\\S+).+Process ID:\\s+(?<ProcessID>\\S+)" + "\"} | Where-Object {$matches.ObjectName -eq " + "\"" + process + "\"" + "} | ForEach-Object {$matches.ProcessID}")
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
fmt.Println("Get-EventLog -log Security -InstanceId 4663 -Newest 1 | Where-Object {$_.message -match \"(?s)Object Name:\\s+(?<ObjectName>\\S+).+Process ID:\\s+(?<ProcessID>\\S+)" + "\"} | Where-Object {$matches.ObjectName -eq " + "\"" + process + "\"" + "} | ForEach-Object {$matches.ProcessID}")
	fmt.Println("Killing process is ----------  " + out)

	err2, _, errout2 := gosh.RunOutput("Stop-Process -Id " + out)
	if err2 != nil {
		log.Printf("error: %v\n", err2)
		fmt.Print(errout2)
	}
}