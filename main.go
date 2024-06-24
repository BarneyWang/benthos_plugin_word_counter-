package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	// Import only pure and standard io Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/io"
	_ "github.com/benthosdev/benthos/v4/public/components/pure"

	// In order to import _all_ Benthos components for third party services
	// uncomment the following line:
	// _ "github.com/benthosdev/benthos/v4/public/components/all"

	// Add your plugin packages here
	_ "github.com/BarneyWang/benthos_plugin_word_counter/processor"
)

func main() {
	service.RunCLI(context.Background())
}
