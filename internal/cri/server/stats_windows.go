/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package server

import (
	"fmt"

	"github.com/containerd/typeurl/v2"

	"github.com/containerd/containerd/v2/api/types"
)

func (m *metricMonitor) extractStats(sandboxes map[string]struct{}, stat *types.Metric) (any, error) {
	return convertMetric(stat)
}

func convertMetric(stats *types.Metric) (any, error) {
	containerStatsData, err := typeurl.UnmarshalAny(stats.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to extract stat for container with id %s: %w", stats.ID, err)
	}

	return containerStatsData, nil
}
