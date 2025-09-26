package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestNewSLACKWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := newSLACKWriter()
	if logger == nil {
		t.Errorf("newSLACKWriter() = %v, want not nil", logger)
	}

	config := "test_config"
	err := logger.Init(config)
	if err != nil {
		t.Errorf("logger.Init() error = %v, want nil", err)
	}

	lm := &LogMsg{
		Level:               LevelTrace,
		Msg:                 "test_message",
		When:                time.Now(),
		FilePath:            "test_file_path",
		LineNumber:          123,
		Args:                []interface{}{"test_arg1", "test_arg2"},
		Prefix:              "test_prefix",
		enableFullFilePath:  true,
		enableFuncCallDepth: true,
	}

	err = logger.WriteMsg(lm)
	if err != nil {
		t.Errorf("logger.WriteMsg() error = %v, want nil", err)
	}

	logger.Destroy()
}
