// Copyright ©2016-2022 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package eval_test

import (
	"fmt"

	"github.com/ddkwork/toolbox/eval"
	"github.com/ddkwork/toolbox/xmath/fixed"
)

func Example() {
	e := eval.NewFixedEvaluator[fixed.D4](nil, true)
	result, err := e.Evaluate("1 + sqrt(2)")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output:
	// 2.4142
}
