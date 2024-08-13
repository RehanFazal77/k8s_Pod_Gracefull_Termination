package e2enode

import (
    "time"

    v1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/util/intstr"
    "k8s.io/kubernetes/test/e2e/framework"
    "k8s.io/kubernetes/test/e2e/framework/kubelet"
    "k8s.io/kubernetes/test/e2e/framework/pods"
    "k8s.io/kubernetes/test/e2e/framework/skipper"
    "k8s.io/kubernetes/test/e2e_node/e2enode"
    "k8s.io/kubernetes/test/utils/ktesting"
    "k8s.io/kubernetes/test/utils/podserver"
)

var _ = SIGDescribe("Pod Lifecycle", func() {
    f := framework.NewDefaultFramework("pod-lifecycle")
    var nodeName string

    ginkgo.BeforeEach(func() {
        nodeName = e2enode.GetNodeThatCanRunPod(f)
    })

    ginkgo.It("should terminate gracefully", func() {
        pod := &v1.Pod{
            ObjectMeta: metav1.ObjectMeta{
                Name: "graceful-termination-pod",
            },
            Spec: v1.PodSpec{
                Containers: []v1.Container{
                    {
                        Name:    "test-container",
                        Image:   "busybox",
                        Command: []string{"sh", "-c", "sleep 60; echo 'Finished'"},
                    },
                },
                TerminationGracePeriodSeconds: int64ptr(30),
            },
        }

        pod = f.PodClient().CreateSync(pod)
        err := f.PodClient().DeleteSync(pod.Name, metav1.DeleteOptions{}, framework.DefaultPodDeletionTimeout)
        framework.ExpectNoError(err, "Failed to delete pod")

        // Verify that the pod was terminated gracefully
        framework.ExpectEqual(pod.Status.Phase, v1.PodSucceeded, "Pod did not terminate gracefully")
    })
})

func int64ptr(i int64) *int64 { return &i }

