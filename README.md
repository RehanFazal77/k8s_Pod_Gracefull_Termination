# k8s_Pod_Graceful_Termiantion_Test

This test is designed to verify the graceful termination behavior of Kubernetes pods. When a pod is deleted, Kubernetes sends a termination signal to the pod, allowing it to shut down gracefully within a specified grace period. The test ensures that the pod receives the termination signal, processes it correctly, and shuts down within the allowed time frame. The pod’s lifecycle events are closely monitored, and key metrics are collected to evaluate the effectiveness of the termination process. The test also validates that all processes within the pod are given sufficient time to complete, ensuring no data loss or corruption occurs during shutdown. The goal is to confirm that Kubernetes adheres to the defined graceful termination policies, providing reliable and consistent behavior across different workloads.

**Checking that termination signals are sent correctly.
Verifying that pods complete their shutdown tasks.
Ensuring that resources are cleaned up properly.**
