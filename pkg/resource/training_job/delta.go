// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package training_job

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customSetDefaults(a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AlgorithmSpecification, b.ko.Spec.AlgorithmSpecification) {
		delta.Add("Spec.AlgorithmSpecification", a.ko.Spec.AlgorithmSpecification, b.ko.Spec.AlgorithmSpecification)
	} else if a.ko.Spec.AlgorithmSpecification != nil && b.ko.Spec.AlgorithmSpecification != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.AlgorithmSpecification.AlgorithmName, b.ko.Spec.AlgorithmSpecification.AlgorithmName) {
			delta.Add("Spec.AlgorithmSpecification.AlgorithmName", a.ko.Spec.AlgorithmSpecification.AlgorithmName, b.ko.Spec.AlgorithmSpecification.AlgorithmName)
		} else if a.ko.Spec.AlgorithmSpecification.AlgorithmName != nil && b.ko.Spec.AlgorithmSpecification.AlgorithmName != nil {
			if *a.ko.Spec.AlgorithmSpecification.AlgorithmName != *b.ko.Spec.AlgorithmSpecification.AlgorithmName {
				delta.Add("Spec.AlgorithmSpecification.AlgorithmName", a.ko.Spec.AlgorithmSpecification.AlgorithmName, b.ko.Spec.AlgorithmSpecification.AlgorithmName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries, b.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries) {
			delta.Add("Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries", a.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries, b.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries)
		} else if a.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != nil && b.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != nil {
			if *a.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != *b.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries {
				delta.Add("Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries", a.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries, b.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.AlgorithmSpecification.TrainingImage, b.ko.Spec.AlgorithmSpecification.TrainingImage) {
			delta.Add("Spec.AlgorithmSpecification.TrainingImage", a.ko.Spec.AlgorithmSpecification.TrainingImage, b.ko.Spec.AlgorithmSpecification.TrainingImage)
		} else if a.ko.Spec.AlgorithmSpecification.TrainingImage != nil && b.ko.Spec.AlgorithmSpecification.TrainingImage != nil {
			if *a.ko.Spec.AlgorithmSpecification.TrainingImage != *b.ko.Spec.AlgorithmSpecification.TrainingImage {
				delta.Add("Spec.AlgorithmSpecification.TrainingImage", a.ko.Spec.AlgorithmSpecification.TrainingImage, b.ko.Spec.AlgorithmSpecification.TrainingImage)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.AlgorithmSpecification.TrainingInputMode, b.ko.Spec.AlgorithmSpecification.TrainingInputMode) {
			delta.Add("Spec.AlgorithmSpecification.TrainingInputMode", a.ko.Spec.AlgorithmSpecification.TrainingInputMode, b.ko.Spec.AlgorithmSpecification.TrainingInputMode)
		} else if a.ko.Spec.AlgorithmSpecification.TrainingInputMode != nil && b.ko.Spec.AlgorithmSpecification.TrainingInputMode != nil {
			if *a.ko.Spec.AlgorithmSpecification.TrainingInputMode != *b.ko.Spec.AlgorithmSpecification.TrainingInputMode {
				delta.Add("Spec.AlgorithmSpecification.TrainingInputMode", a.ko.Spec.AlgorithmSpecification.TrainingInputMode, b.ko.Spec.AlgorithmSpecification.TrainingInputMode)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CheckpointConfig, b.ko.Spec.CheckpointConfig) {
		delta.Add("Spec.CheckpointConfig", a.ko.Spec.CheckpointConfig, b.ko.Spec.CheckpointConfig)
	} else if a.ko.Spec.CheckpointConfig != nil && b.ko.Spec.CheckpointConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CheckpointConfig.LocalPath, b.ko.Spec.CheckpointConfig.LocalPath) {
			delta.Add("Spec.CheckpointConfig.LocalPath", a.ko.Spec.CheckpointConfig.LocalPath, b.ko.Spec.CheckpointConfig.LocalPath)
		} else if a.ko.Spec.CheckpointConfig.LocalPath != nil && b.ko.Spec.CheckpointConfig.LocalPath != nil {
			if *a.ko.Spec.CheckpointConfig.LocalPath != *b.ko.Spec.CheckpointConfig.LocalPath {
				delta.Add("Spec.CheckpointConfig.LocalPath", a.ko.Spec.CheckpointConfig.LocalPath, b.ko.Spec.CheckpointConfig.LocalPath)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CheckpointConfig.S3URI, b.ko.Spec.CheckpointConfig.S3URI) {
			delta.Add("Spec.CheckpointConfig.S3URI", a.ko.Spec.CheckpointConfig.S3URI, b.ko.Spec.CheckpointConfig.S3URI)
		} else if a.ko.Spec.CheckpointConfig.S3URI != nil && b.ko.Spec.CheckpointConfig.S3URI != nil {
			if *a.ko.Spec.CheckpointConfig.S3URI != *b.ko.Spec.CheckpointConfig.S3URI {
				delta.Add("Spec.CheckpointConfig.S3URI", a.ko.Spec.CheckpointConfig.S3URI, b.ko.Spec.CheckpointConfig.S3URI)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DebugHookConfig, b.ko.Spec.DebugHookConfig) {
		delta.Add("Spec.DebugHookConfig", a.ko.Spec.DebugHookConfig, b.ko.Spec.DebugHookConfig)
	} else if a.ko.Spec.DebugHookConfig != nil && b.ko.Spec.DebugHookConfig != nil {
		if !reflect.DeepEqual(a.ko.Spec.DebugHookConfig.CollectionConfigurations, b.ko.Spec.DebugHookConfig.CollectionConfigurations) {
			delta.Add("Spec.DebugHookConfig.CollectionConfigurations", a.ko.Spec.DebugHookConfig.CollectionConfigurations, b.ko.Spec.DebugHookConfig.CollectionConfigurations)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DebugHookConfig.HookParameters, b.ko.Spec.DebugHookConfig.HookParameters) {
			delta.Add("Spec.DebugHookConfig.HookParameters", a.ko.Spec.DebugHookConfig.HookParameters, b.ko.Spec.DebugHookConfig.HookParameters)
		} else if a.ko.Spec.DebugHookConfig.HookParameters != nil && b.ko.Spec.DebugHookConfig.HookParameters != nil {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.DebugHookConfig.HookParameters, b.ko.Spec.DebugHookConfig.HookParameters) {
				delta.Add("Spec.DebugHookConfig.HookParameters", a.ko.Spec.DebugHookConfig.HookParameters, b.ko.Spec.DebugHookConfig.HookParameters)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DebugHookConfig.LocalPath, b.ko.Spec.DebugHookConfig.LocalPath) {
			delta.Add("Spec.DebugHookConfig.LocalPath", a.ko.Spec.DebugHookConfig.LocalPath, b.ko.Spec.DebugHookConfig.LocalPath)
		} else if a.ko.Spec.DebugHookConfig.LocalPath != nil && b.ko.Spec.DebugHookConfig.LocalPath != nil {
			if *a.ko.Spec.DebugHookConfig.LocalPath != *b.ko.Spec.DebugHookConfig.LocalPath {
				delta.Add("Spec.DebugHookConfig.LocalPath", a.ko.Spec.DebugHookConfig.LocalPath, b.ko.Spec.DebugHookConfig.LocalPath)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DebugHookConfig.S3OutputPath, b.ko.Spec.DebugHookConfig.S3OutputPath) {
			delta.Add("Spec.DebugHookConfig.S3OutputPath", a.ko.Spec.DebugHookConfig.S3OutputPath, b.ko.Spec.DebugHookConfig.S3OutputPath)
		} else if a.ko.Spec.DebugHookConfig.S3OutputPath != nil && b.ko.Spec.DebugHookConfig.S3OutputPath != nil {
			if *a.ko.Spec.DebugHookConfig.S3OutputPath != *b.ko.Spec.DebugHookConfig.S3OutputPath {
				delta.Add("Spec.DebugHookConfig.S3OutputPath", a.ko.Spec.DebugHookConfig.S3OutputPath, b.ko.Spec.DebugHookConfig.S3OutputPath)
			}
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.DebugRuleConfigurations, b.ko.Spec.DebugRuleConfigurations) {
		delta.Add("Spec.DebugRuleConfigurations", a.ko.Spec.DebugRuleConfigurations, b.ko.Spec.DebugRuleConfigurations)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableInterContainerTrafficEncryption, b.ko.Spec.EnableInterContainerTrafficEncryption) {
		delta.Add("Spec.EnableInterContainerTrafficEncryption", a.ko.Spec.EnableInterContainerTrafficEncryption, b.ko.Spec.EnableInterContainerTrafficEncryption)
	} else if a.ko.Spec.EnableInterContainerTrafficEncryption != nil && b.ko.Spec.EnableInterContainerTrafficEncryption != nil {
		if *a.ko.Spec.EnableInterContainerTrafficEncryption != *b.ko.Spec.EnableInterContainerTrafficEncryption {
			delta.Add("Spec.EnableInterContainerTrafficEncryption", a.ko.Spec.EnableInterContainerTrafficEncryption, b.ko.Spec.EnableInterContainerTrafficEncryption)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableManagedSpotTraining, b.ko.Spec.EnableManagedSpotTraining) {
		delta.Add("Spec.EnableManagedSpotTraining", a.ko.Spec.EnableManagedSpotTraining, b.ko.Spec.EnableManagedSpotTraining)
	} else if a.ko.Spec.EnableManagedSpotTraining != nil && b.ko.Spec.EnableManagedSpotTraining != nil {
		if *a.ko.Spec.EnableManagedSpotTraining != *b.ko.Spec.EnableManagedSpotTraining {
			delta.Add("Spec.EnableManagedSpotTraining", a.ko.Spec.EnableManagedSpotTraining, b.ko.Spec.EnableManagedSpotTraining)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EnableNetworkIsolation, b.ko.Spec.EnableNetworkIsolation) {
		delta.Add("Spec.EnableNetworkIsolation", a.ko.Spec.EnableNetworkIsolation, b.ko.Spec.EnableNetworkIsolation)
	} else if a.ko.Spec.EnableNetworkIsolation != nil && b.ko.Spec.EnableNetworkIsolation != nil {
		if *a.ko.Spec.EnableNetworkIsolation != *b.ko.Spec.EnableNetworkIsolation {
			delta.Add("Spec.EnableNetworkIsolation", a.ko.Spec.EnableNetworkIsolation, b.ko.Spec.EnableNetworkIsolation)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Environment, b.ko.Spec.Environment) {
		delta.Add("Spec.Environment", a.ko.Spec.Environment, b.ko.Spec.Environment)
	} else if a.ko.Spec.Environment != nil && b.ko.Spec.Environment != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.Environment, b.ko.Spec.Environment) {
			delta.Add("Spec.Environment", a.ko.Spec.Environment, b.ko.Spec.Environment)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig, b.ko.Spec.ExperimentConfig) {
		delta.Add("Spec.ExperimentConfig", a.ko.Spec.ExperimentConfig, b.ko.Spec.ExperimentConfig)
	} else if a.ko.Spec.ExperimentConfig != nil && b.ko.Spec.ExperimentConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig.ExperimentName, b.ko.Spec.ExperimentConfig.ExperimentName) {
			delta.Add("Spec.ExperimentConfig.ExperimentName", a.ko.Spec.ExperimentConfig.ExperimentName, b.ko.Spec.ExperimentConfig.ExperimentName)
		} else if a.ko.Spec.ExperimentConfig.ExperimentName != nil && b.ko.Spec.ExperimentConfig.ExperimentName != nil {
			if *a.ko.Spec.ExperimentConfig.ExperimentName != *b.ko.Spec.ExperimentConfig.ExperimentName {
				delta.Add("Spec.ExperimentConfig.ExperimentName", a.ko.Spec.ExperimentConfig.ExperimentName, b.ko.Spec.ExperimentConfig.ExperimentName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig.TrialComponentDisplayName, b.ko.Spec.ExperimentConfig.TrialComponentDisplayName) {
			delta.Add("Spec.ExperimentConfig.TrialComponentDisplayName", a.ko.Spec.ExperimentConfig.TrialComponentDisplayName, b.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
		} else if a.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil && b.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			if *a.ko.Spec.ExperimentConfig.TrialComponentDisplayName != *b.ko.Spec.ExperimentConfig.TrialComponentDisplayName {
				delta.Add("Spec.ExperimentConfig.TrialComponentDisplayName", a.ko.Spec.ExperimentConfig.TrialComponentDisplayName, b.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ExperimentConfig.TrialName, b.ko.Spec.ExperimentConfig.TrialName) {
			delta.Add("Spec.ExperimentConfig.TrialName", a.ko.Spec.ExperimentConfig.TrialName, b.ko.Spec.ExperimentConfig.TrialName)
		} else if a.ko.Spec.ExperimentConfig.TrialName != nil && b.ko.Spec.ExperimentConfig.TrialName != nil {
			if *a.ko.Spec.ExperimentConfig.TrialName != *b.ko.Spec.ExperimentConfig.TrialName {
				delta.Add("Spec.ExperimentConfig.TrialName", a.ko.Spec.ExperimentConfig.TrialName, b.ko.Spec.ExperimentConfig.TrialName)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.HyperParameters, b.ko.Spec.HyperParameters) {
		delta.Add("Spec.HyperParameters", a.ko.Spec.HyperParameters, b.ko.Spec.HyperParameters)
	} else if a.ko.Spec.HyperParameters != nil && b.ko.Spec.HyperParameters != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.HyperParameters, b.ko.Spec.HyperParameters) {
			delta.Add("Spec.HyperParameters", a.ko.Spec.HyperParameters, b.ko.Spec.HyperParameters)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.InputDataConfig, b.ko.Spec.InputDataConfig) {
		delta.Add("Spec.InputDataConfig", a.ko.Spec.InputDataConfig, b.ko.Spec.InputDataConfig)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OutputDataConfig, b.ko.Spec.OutputDataConfig) {
		delta.Add("Spec.OutputDataConfig", a.ko.Spec.OutputDataConfig, b.ko.Spec.OutputDataConfig)
	} else if a.ko.Spec.OutputDataConfig != nil && b.ko.Spec.OutputDataConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.OutputDataConfig.KMSKeyID, b.ko.Spec.OutputDataConfig.KMSKeyID) {
			delta.Add("Spec.OutputDataConfig.KMSKeyID", a.ko.Spec.OutputDataConfig.KMSKeyID, b.ko.Spec.OutputDataConfig.KMSKeyID)
		} else if a.ko.Spec.OutputDataConfig.KMSKeyID != nil && b.ko.Spec.OutputDataConfig.KMSKeyID != nil {
			if *a.ko.Spec.OutputDataConfig.KMSKeyID != *b.ko.Spec.OutputDataConfig.KMSKeyID {
				delta.Add("Spec.OutputDataConfig.KMSKeyID", a.ko.Spec.OutputDataConfig.KMSKeyID, b.ko.Spec.OutputDataConfig.KMSKeyID)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.OutputDataConfig.S3OutputPath, b.ko.Spec.OutputDataConfig.S3OutputPath) {
			delta.Add("Spec.OutputDataConfig.S3OutputPath", a.ko.Spec.OutputDataConfig.S3OutputPath, b.ko.Spec.OutputDataConfig.S3OutputPath)
		} else if a.ko.Spec.OutputDataConfig.S3OutputPath != nil && b.ko.Spec.OutputDataConfig.S3OutputPath != nil {
			if *a.ko.Spec.OutputDataConfig.S3OutputPath != *b.ko.Spec.OutputDataConfig.S3OutputPath {
				delta.Add("Spec.OutputDataConfig.S3OutputPath", a.ko.Spec.OutputDataConfig.S3OutputPath, b.ko.Spec.OutputDataConfig.S3OutputPath)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ProfilerConfig, b.ko.Spec.ProfilerConfig) {
		delta.Add("Spec.ProfilerConfig", a.ko.Spec.ProfilerConfig, b.ko.Spec.ProfilerConfig)
	} else if a.ko.Spec.ProfilerConfig != nil && b.ko.Spec.ProfilerConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds, b.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds) {
			delta.Add("Spec.ProfilerConfig.ProfilingIntervalInMilliseconds", a.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds, b.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds)
		} else if a.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds != nil && b.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds != nil {
			if *a.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds != *b.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds {
				delta.Add("Spec.ProfilerConfig.ProfilingIntervalInMilliseconds", a.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds, b.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ProfilerConfig.ProfilingParameters, b.ko.Spec.ProfilerConfig.ProfilingParameters) {
			delta.Add("Spec.ProfilerConfig.ProfilingParameters", a.ko.Spec.ProfilerConfig.ProfilingParameters, b.ko.Spec.ProfilerConfig.ProfilingParameters)
		} else if a.ko.Spec.ProfilerConfig.ProfilingParameters != nil && b.ko.Spec.ProfilerConfig.ProfilingParameters != nil {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.ProfilerConfig.ProfilingParameters, b.ko.Spec.ProfilerConfig.ProfilingParameters) {
				delta.Add("Spec.ProfilerConfig.ProfilingParameters", a.ko.Spec.ProfilerConfig.ProfilingParameters, b.ko.Spec.ProfilerConfig.ProfilingParameters)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ProfilerConfig.S3OutputPath, b.ko.Spec.ProfilerConfig.S3OutputPath) {
			delta.Add("Spec.ProfilerConfig.S3OutputPath", a.ko.Spec.ProfilerConfig.S3OutputPath, b.ko.Spec.ProfilerConfig.S3OutputPath)
		} else if a.ko.Spec.ProfilerConfig.S3OutputPath != nil && b.ko.Spec.ProfilerConfig.S3OutputPath != nil {
			if *a.ko.Spec.ProfilerConfig.S3OutputPath != *b.ko.Spec.ProfilerConfig.S3OutputPath {
				delta.Add("Spec.ProfilerConfig.S3OutputPath", a.ko.Spec.ProfilerConfig.S3OutputPath, b.ko.Spec.ProfilerConfig.S3OutputPath)
			}
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ProfilerRuleConfigurations, b.ko.Spec.ProfilerRuleConfigurations) {
		delta.Add("Spec.ProfilerRuleConfigurations", a.ko.Spec.ProfilerRuleConfigurations, b.ko.Spec.ProfilerRuleConfigurations)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResourceConfig, b.ko.Spec.ResourceConfig) {
		delta.Add("Spec.ResourceConfig", a.ko.Spec.ResourceConfig, b.ko.Spec.ResourceConfig)
	} else if a.ko.Spec.ResourceConfig != nil && b.ko.Spec.ResourceConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceConfig.InstanceCount, b.ko.Spec.ResourceConfig.InstanceCount) {
			delta.Add("Spec.ResourceConfig.InstanceCount", a.ko.Spec.ResourceConfig.InstanceCount, b.ko.Spec.ResourceConfig.InstanceCount)
		} else if a.ko.Spec.ResourceConfig.InstanceCount != nil && b.ko.Spec.ResourceConfig.InstanceCount != nil {
			if *a.ko.Spec.ResourceConfig.InstanceCount != *b.ko.Spec.ResourceConfig.InstanceCount {
				delta.Add("Spec.ResourceConfig.InstanceCount", a.ko.Spec.ResourceConfig.InstanceCount, b.ko.Spec.ResourceConfig.InstanceCount)
			}
		}
		if !reflect.DeepEqual(a.ko.Spec.ResourceConfig.InstanceGroups, b.ko.Spec.ResourceConfig.InstanceGroups) {
			delta.Add("Spec.ResourceConfig.InstanceGroups", a.ko.Spec.ResourceConfig.InstanceGroups, b.ko.Spec.ResourceConfig.InstanceGroups)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceConfig.InstanceType, b.ko.Spec.ResourceConfig.InstanceType) {
			delta.Add("Spec.ResourceConfig.InstanceType", a.ko.Spec.ResourceConfig.InstanceType, b.ko.Spec.ResourceConfig.InstanceType)
		} else if a.ko.Spec.ResourceConfig.InstanceType != nil && b.ko.Spec.ResourceConfig.InstanceType != nil {
			if *a.ko.Spec.ResourceConfig.InstanceType != *b.ko.Spec.ResourceConfig.InstanceType {
				delta.Add("Spec.ResourceConfig.InstanceType", a.ko.Spec.ResourceConfig.InstanceType, b.ko.Spec.ResourceConfig.InstanceType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds, b.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds) {
			delta.Add("Spec.ResourceConfig.KeepAlivePeriodInSeconds", a.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds, b.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds)
		} else if a.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds != nil && b.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds != nil {
			if *a.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds != *b.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds {
				delta.Add("Spec.ResourceConfig.KeepAlivePeriodInSeconds", a.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds, b.ko.Spec.ResourceConfig.KeepAlivePeriodInSeconds)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceConfig.VolumeKMSKeyID, b.ko.Spec.ResourceConfig.VolumeKMSKeyID) {
			delta.Add("Spec.ResourceConfig.VolumeKMSKeyID", a.ko.Spec.ResourceConfig.VolumeKMSKeyID, b.ko.Spec.ResourceConfig.VolumeKMSKeyID)
		} else if a.ko.Spec.ResourceConfig.VolumeKMSKeyID != nil && b.ko.Spec.ResourceConfig.VolumeKMSKeyID != nil {
			if *a.ko.Spec.ResourceConfig.VolumeKMSKeyID != *b.ko.Spec.ResourceConfig.VolumeKMSKeyID {
				delta.Add("Spec.ResourceConfig.VolumeKMSKeyID", a.ko.Spec.ResourceConfig.VolumeKMSKeyID, b.ko.Spec.ResourceConfig.VolumeKMSKeyID)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ResourceConfig.VolumeSizeInGB, b.ko.Spec.ResourceConfig.VolumeSizeInGB) {
			delta.Add("Spec.ResourceConfig.VolumeSizeInGB", a.ko.Spec.ResourceConfig.VolumeSizeInGB, b.ko.Spec.ResourceConfig.VolumeSizeInGB)
		} else if a.ko.Spec.ResourceConfig.VolumeSizeInGB != nil && b.ko.Spec.ResourceConfig.VolumeSizeInGB != nil {
			if *a.ko.Spec.ResourceConfig.VolumeSizeInGB != *b.ko.Spec.ResourceConfig.VolumeSizeInGB {
				delta.Add("Spec.ResourceConfig.VolumeSizeInGB", a.ko.Spec.ResourceConfig.VolumeSizeInGB, b.ko.Spec.ResourceConfig.VolumeSizeInGB)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RetryStrategy, b.ko.Spec.RetryStrategy) {
		delta.Add("Spec.RetryStrategy", a.ko.Spec.RetryStrategy, b.ko.Spec.RetryStrategy)
	} else if a.ko.Spec.RetryStrategy != nil && b.ko.Spec.RetryStrategy != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.RetryStrategy.MaximumRetryAttempts, b.ko.Spec.RetryStrategy.MaximumRetryAttempts) {
			delta.Add("Spec.RetryStrategy.MaximumRetryAttempts", a.ko.Spec.RetryStrategy.MaximumRetryAttempts, b.ko.Spec.RetryStrategy.MaximumRetryAttempts)
		} else if a.ko.Spec.RetryStrategy.MaximumRetryAttempts != nil && b.ko.Spec.RetryStrategy.MaximumRetryAttempts != nil {
			if *a.ko.Spec.RetryStrategy.MaximumRetryAttempts != *b.ko.Spec.RetryStrategy.MaximumRetryAttempts {
				delta.Add("Spec.RetryStrategy.MaximumRetryAttempts", a.ko.Spec.RetryStrategy.MaximumRetryAttempts, b.ko.Spec.RetryStrategy.MaximumRetryAttempts)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RoleARN, b.ko.Spec.RoleARN) {
		delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
	} else if a.ko.Spec.RoleARN != nil && b.ko.Spec.RoleARN != nil {
		if *a.ko.Spec.RoleARN != *b.ko.Spec.RoleARN {
			delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition, b.ko.Spec.StoppingCondition) {
		delta.Add("Spec.StoppingCondition", a.ko.Spec.StoppingCondition, b.ko.Spec.StoppingCondition)
	} else if a.ko.Spec.StoppingCondition != nil && b.ko.Spec.StoppingCondition != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds) {
			delta.Add("Spec.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
		} else if a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil && b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			if *a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != *b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds {
				delta.Add("Spec.StoppingCondition.MaxRuntimeInSeconds", a.ko.Spec.StoppingCondition.MaxRuntimeInSeconds, b.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds, b.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds) {
			delta.Add("Spec.StoppingCondition.MaxWaitTimeInSeconds", a.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds, b.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds)
		} else if a.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds != nil && b.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds != nil {
			if *a.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds != *b.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds {
				delta.Add("Spec.StoppingCondition.MaxWaitTimeInSeconds", a.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds, b.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TensorBoardOutputConfig, b.ko.Spec.TensorBoardOutputConfig) {
		delta.Add("Spec.TensorBoardOutputConfig", a.ko.Spec.TensorBoardOutputConfig, b.ko.Spec.TensorBoardOutputConfig)
	} else if a.ko.Spec.TensorBoardOutputConfig != nil && b.ko.Spec.TensorBoardOutputConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.TensorBoardOutputConfig.LocalPath, b.ko.Spec.TensorBoardOutputConfig.LocalPath) {
			delta.Add("Spec.TensorBoardOutputConfig.LocalPath", a.ko.Spec.TensorBoardOutputConfig.LocalPath, b.ko.Spec.TensorBoardOutputConfig.LocalPath)
		} else if a.ko.Spec.TensorBoardOutputConfig.LocalPath != nil && b.ko.Spec.TensorBoardOutputConfig.LocalPath != nil {
			if *a.ko.Spec.TensorBoardOutputConfig.LocalPath != *b.ko.Spec.TensorBoardOutputConfig.LocalPath {
				delta.Add("Spec.TensorBoardOutputConfig.LocalPath", a.ko.Spec.TensorBoardOutputConfig.LocalPath, b.ko.Spec.TensorBoardOutputConfig.LocalPath)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TensorBoardOutputConfig.S3OutputPath, b.ko.Spec.TensorBoardOutputConfig.S3OutputPath) {
			delta.Add("Spec.TensorBoardOutputConfig.S3OutputPath", a.ko.Spec.TensorBoardOutputConfig.S3OutputPath, b.ko.Spec.TensorBoardOutputConfig.S3OutputPath)
		} else if a.ko.Spec.TensorBoardOutputConfig.S3OutputPath != nil && b.ko.Spec.TensorBoardOutputConfig.S3OutputPath != nil {
			if *a.ko.Spec.TensorBoardOutputConfig.S3OutputPath != *b.ko.Spec.TensorBoardOutputConfig.S3OutputPath {
				delta.Add("Spec.TensorBoardOutputConfig.S3OutputPath", a.ko.Spec.TensorBoardOutputConfig.S3OutputPath, b.ko.Spec.TensorBoardOutputConfig.S3OutputPath)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TrainingJobName, b.ko.Spec.TrainingJobName) {
		delta.Add("Spec.TrainingJobName", a.ko.Spec.TrainingJobName, b.ko.Spec.TrainingJobName)
	} else if a.ko.Spec.TrainingJobName != nil && b.ko.Spec.TrainingJobName != nil {
		if *a.ko.Spec.TrainingJobName != *b.ko.Spec.TrainingJobName {
			delta.Add("Spec.TrainingJobName", a.ko.Spec.TrainingJobName, b.ko.Spec.TrainingJobName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.VPCConfig, b.ko.Spec.VPCConfig) {
		delta.Add("Spec.VPCConfig", a.ko.Spec.VPCConfig, b.ko.Spec.VPCConfig)
	} else if a.ko.Spec.VPCConfig != nil && b.ko.Spec.VPCConfig != nil {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs) {
			delta.Add("Spec.VPCConfig.SecurityGroupIDs", a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCConfig.Subnets, b.ko.Spec.VPCConfig.Subnets) {
			delta.Add("Spec.VPCConfig.Subnets", a.ko.Spec.VPCConfig.Subnets, b.ko.Spec.VPCConfig.Subnets)
		}
	}

	customPostCompare(b, a, delta)
	return delta
}
