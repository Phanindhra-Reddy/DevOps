/*
Copyright 2024 Phani Aalla.

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

package controllers

import (
	"context"
	// "encoding/json"
	// "fmt"
	// "net/http"
	// "os"

	// appsv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1alpha1 "github.com/phani953/fenode-operator/api/v1alpha1"
)

// FenodeReconciler reconciles a Fenode object
type FenodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// type DeploymentRequest struct {
// 	CustomerID string `json:"customerID"`
// 	GUID       string `json:"guid"`
// }

//+kubebuilder:rbac:groups=api.phanindhra.reddy,resources=fenodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.phanindhra.reddy,resources=fenodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.phanindhra.reddy,resources=fenodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Fenode object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *FenodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FenodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.Fenode{}).
		Complete(r)
}

// func (r *FenodeReconciler) startHTTPServer() {
// 	http.HandleFunc("/deploy", r.handleDeploy)
// 	http.ListenAndServe(":8020", nil)
// }

// func (r *FenodeReconciler) handleDeploy(w http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var deploymentReq DeploymentRequest
// 	if err := json.NewDecoder(req.Body).Decode(&deploymentReq); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Create Nginx Deployment
// 	err := r.createNginxDeployment(deploymentReq.CustomerID, deploymentReq.GUID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Fprintf(w, "Deployment initiated for %s_%s", deploymentReq.CustomerID, deploymentReq.GUID)
// }

// func (r *FenodeReconciler) createNginxDeployment(customerID, guid string) error {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		return err
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		return err
// 	}
// 	deployment := &appsv1.Deployment{
// 		ObjectMeta: v1.ObjectMeta{
// 			Name:      fmt.Sprintf("%s_%s", customerID, guid),
// 			Namespace: os.Getenv("POD_NAMESPACE"), // Assumes the operator is deployed in the desired namespace
// 		},
// 		Spec: appsv1.DeploymentSpec{
// 			Replicas: int32Ptr(1),
// 			Selector: &v1.LabelSelector{
// 				MatchLabels: map[string]string{
// 					"app": fmt.Sprintf("%s_%s", customerID, guid),
// 				},
// 			},
// 			Template: corev1.PodTemplateSpec{
// 				ObjectMeta: v1.ObjectMeta{
// 					Labels: map[string]string{
// 						"app": fmt.Sprintf("%s_%s", customerID, guid),
// 					},
// 				},
// 				Spec: corev1.PodSpec{
// 					Containers: []corev1.Container{
// 						{
// 							Name:  "nginx",
// 							Image: "nginx:latest",
// 							Ports: []corev1.ContainerPort{
// 								{
// 									ContainerPort: 80,
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	_, err = clientset.AppsV1().Deployments(os.Getenv("POD_NAMESPACE")).Create(context.TODO(), deployment, v1.CreateOptions{})
// 	return err
// }

// func int32Ptr(i int32) *int32 { return &i }
