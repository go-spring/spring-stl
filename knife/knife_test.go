/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package knife_test

import (
	"context"
	"testing"

	"github.com/go-spring/spring-stl/assert"
	"github.com/go-spring/spring-stl/knife"
)

func TestKnife(t *testing.T) {
	ctx := context.TODO()

	v := knife.Get(ctx, "a")
	assert.Nil(t, v)

	knife.Set(ctx, "a", "b")

	v = knife.Get(ctx, "a")
	assert.Nil(t, v)

	ctx = knife.New(ctx)

	v = knife.Get(ctx, "a")
	assert.Nil(t, v)

	knife.Set(ctx, "a", "b")

	v = knife.Get(ctx, "a")
	assert.Equal(t, v, "b")
}
