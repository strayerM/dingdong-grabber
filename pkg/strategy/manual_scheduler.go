/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package strategy

import (
	"context"

	"github.com/dingdong-grabber/pkg/order"
)

// ManualScheduler 人工策略调度器
type ManualScheduler struct {
	Scheduler `json:",inline"`
}

func NewManualScheduler(o *order.Order, baseTheadSize, submitOrderTheadSize, minSleepMillis, maxSleepMillis int, play bool, pushToken string) Interface {
	if minSleepMillis > maxSleepMillis {
		maxSleepMillis = minSleepMillis
	}
	return &ManualScheduler{Scheduler{
		o:                    o,
		play:                 play,
		baseTheadSize:        baseTheadSize,
		submitOrderTheadSize: submitOrderTheadSize,
		minSleepMillis:       minSleepMillis,
		maxSleepMillis:       maxSleepMillis,
		pushToken:            pushToken,
	}}
}

func (ms *ManualScheduler) Schedule(ctx context.Context) error {
	return ms.Scheduler.Schedule(ctx)
}
