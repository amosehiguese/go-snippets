package snippet

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)


type SimpleErr struct {
	Inner 		 error
	Msg 			 string
	StackTrace string
	Misc 			 map[string]interface{}
}

func WrapError(err error, msgf string, msgArgs ...any) SimpleErr {
	return SimpleErr{
		Inner: err,
		Msg: fmt.Sprintf(msgf, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc: make(map[string]interface{}),
	}
}

func (s SimpleErr) Error() string {
	return s.Msg
}

type LowlevelErr struct {
	error
}

func IsGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowlevelErr{WrapError(err, err.Error())}
	}
	fmt.Println(info.Mode().Perm()&0100)
	return info.Mode().Perm()&0100 == 0100, nil
}

type IntermediateErr struct {
	error
}

func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := IsGloballyExec(jobBinPath)
	if err != nil {
		return IntermediateErr{WrapError(err, "cannot run job %q: requisite binaries not available \n", id )}
	} else if !isExecutable {
		return WrapError(nil, "cannot run job %q: requisite binaries are not executable\n", id)
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}

func HandleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: \n", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

func ErrMain()  {
	log.SetOutput(os.Stdout)	
	log.SetFlags(log.Ltime|log.LUTC)

	err := RunJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		HandleError(1, err, msg)
	}
}