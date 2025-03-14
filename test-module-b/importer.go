package testmoduleb

import testmodulea "github.com/skarlso/test-module-import/test-module-a"

func TestModuleImport() {
	testmodulea.TestModuleA()
}
