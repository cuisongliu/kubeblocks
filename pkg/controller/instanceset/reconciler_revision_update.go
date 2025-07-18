/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package instanceset

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	workloads "github.com/apecloud/kubeblocks/apis/workloads/v1"
	"github.com/apecloud/kubeblocks/pkg/controller/instanceset/instancetemplate"
	"github.com/apecloud/kubeblocks/pkg/controller/kubebuilderx"
	"github.com/apecloud/kubeblocks/pkg/controller/model"
)

// revisionUpdateReconciler is responsible for updating the expected instance names and their corresponding revisions in the status when there are changes in the spec.
type revisionUpdateReconciler struct{}

type instanceRevision struct {
	name     string
	revision string
}

func NewRevisionUpdateReconciler() kubebuilderx.Reconciler {
	return &revisionUpdateReconciler{}
}

func (r *revisionUpdateReconciler) PreCondition(tree *kubebuilderx.ObjectTree) *kubebuilderx.CheckResult {
	if tree.GetRoot() == nil || !model.IsObjectUpdating(tree.GetRoot()) {
		return kubebuilderx.ConditionUnsatisfied
	}
	return kubebuilderx.ConditionSatisfied
}

func (r *revisionUpdateReconciler) Reconcile(tree *kubebuilderx.ObjectTree) (kubebuilderx.Result, error) {
	its, _ := tree.GetRoot().(*workloads.InstanceSet)
	itsExt, err := instancetemplate.BuildInstanceSetExt(its, tree)
	if err != nil {
		return kubebuilderx.Continue, err
	}

	// build instance revision list from instance templates
	var instanceRevisionList []instanceRevision
	nameBuilder, err := instancetemplate.NewPodNameBuilder(itsExt, nil)
	if err != nil {
		return kubebuilderx.Continue, err
	}
	nameMap, err := nameBuilder.BuildInstanceName2TemplateMap()
	if err != nil {
		return kubebuilderx.Continue, err
	}
	for instanceName, templateExt := range nameMap {
		revision, err := buildInstanceTemplateRevision(&templateExt.PodTemplateSpec, its)
		if err != nil {
			return kubebuilderx.Continue, err
		}
		instanceRevisionList = append(instanceRevisionList, instanceRevision{name: instanceName, revision: revision})
	}

	updatedRevisions := make(map[string]string, len(instanceRevisionList))
	for _, r := range instanceRevisionList {
		updatedRevisions[r.name] = r.revision
	}

	// 3. persistent these revisions to status
	revisions, err := buildRevisions(updatedRevisions)
	if err != nil {
		return kubebuilderx.Continue, err
	}
	its.Status.UpdateRevisions = revisions
	updateRevision := ""
	if len(instanceRevisionList) > 0 {
		updateRevision = instanceRevisionList[len(instanceRevisionList)-1].revision
	}
	its.Status.UpdateRevision = updateRevision
	updatedReplicas, err := calculateUpdatedReplicas(its, tree.List(&corev1.Pod{}))
	if err != nil {
		return kubebuilderx.Continue, err
	}
	its.Status.UpdatedReplicas = updatedReplicas
	// The 'ObservedGeneration' field is used to indicate whether the revisions have been updated.
	// Computing these revisions in each reconciliation loop can be time-consuming, so we optimize it by
	// performing the computation only when the 'spec' is updated.
	its.Status.ObservedGeneration = its.Generation

	return kubebuilderx.Continue, nil
}

func calculateUpdatedReplicas(its *workloads.InstanceSet, pods []client.Object) (int32, error) {
	updatedReplicas := int32(0)
	for i := range pods {
		pod, _ := pods[i].(*corev1.Pod)
		updated, err := IsPodUpdated(its, pod)
		if err != nil {
			return 0, nil
		}
		if updated {
			updatedReplicas++
		}

	}
	return updatedReplicas, nil
}

var _ kubebuilderx.Reconciler = &revisionUpdateReconciler{}
