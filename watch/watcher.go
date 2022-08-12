package watch
import (
	"fmt"
	"log"
	"github.com/abdfnx/gosh"
	"github.com/howeyc/fsnotify"
	"ransomwarecheck/killer"
)

func Watch(target string) {
	//enable auditing for created file
	err, _, errout := gosh.RunOutput("Add-NTFSAudit -Path " + "\"" + target + "\"" + " -AccessRights FullControl -Account Everyone -AuditFlags Success -InheritanceFlags ContainerInherit,ObjectInherit -PropagationFlags None")
	err, enable, errout := gosh.RunOutput("Get-NTFSAudit -Path " + "\"" + target + "\"")
	fmt.Println(enable)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
	fmt.Println(target + " is monitoring")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Watch(target)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
			process := target
			fmt.Println(process, "process which use this file is closing")
			killer.Kill(process)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}