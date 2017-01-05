// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aspect

import (
	"io"

	"github.com/golang/protobuf/proto"
)

type (

	// Aspect -- User visible cross cutting concern
	// istio/metrics, istio/accessControl, istio/accessLog are all aspects of a deployment
	// This interface is extended by aspects;
	// ex: listChecker.Aspect adds --> CheckList(Symbol string) (bool, error)
	// The extended interface is implemented by adapter impls.
	// ex: //adapter/ipListChecker aspectState implements the listChecker.Aspect.
	Aspect interface {
		io.Closer
	}

	// Adapter represents an Aspect Builder that constructs instances a specific aspect.
	// This interface is extended by  aspects; (metrics, listChecker, ...)
	// The extended interface is implemented by adapter implementations
	Adapter interface {
		io.Closer
		// Name returns the official name of this adapter. ex. "istio/statsd".
		Name() string
		// Description returns a user-friendly description of this adapter.
		Description() string
		// DefaultConfig returns a default configuration struct for this
		// adapter. This will be used by the configuration system to establish
		// the shape of the block of configuration state passed to the NewAspect method.
		DefaultConfig() (implConfig proto.Message)
		// ValidateConfig determines whether the given configuration meets all correctness requirements.
		ValidateConfig(implConfig proto.Message) *ConfigErrors
	}
)
