package kreate

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//create a test folder to hold temporary test files
	os.Mkdir("./chartTest", 0777)

	//run through tests
	code := m.Run()

	//delete temporary test directory
	os.RemoveAll("./chartTest")

	//the end
	os.Exit(code)
}

func TestCreateChart(t *testing.T) {
	//
}

func TestBuildFileSystem(t *testing.T) {
	//
}

func TestCreateValues(t *testing.T) {
	//
}

func TestCreateChartFile(t *testing.T) {
	//
}

func TestPopulateChart(t *testing.T) {
	//
}

func TestFixFileSystem(t *testing.T) {
	//
}

func TestCopyDir(t *testing.T) {
	//
}
