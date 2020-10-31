package batch

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	"volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

func init() {
	klog.Infof("=====init batch job=====")
	Install(legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	scheme.Default()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
}
