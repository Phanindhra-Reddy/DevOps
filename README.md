# Kubernetes Operators
- A Kubernetes Operator is a powerful pattern and tool used to automate the management of complex stateful applications on Kubernetes.
- Operators extend the Kubernetes API and leverage the native Kubernetes capabilities to manage application-specific tasks and lifecycle management.
- The goal of an Operator is to put operational knowledge into software.
- Operators typically provide controllers and Custom Resource Definitions (CRDs) that let you rapidly deploy new app instances without deep knowledge of their requirements or how they work.

## How Kubernetes Operators Work
An operator actively monitors the Kubernetes API for any alterations pertaining to a particular resource, like a custom resource definition (CRD). Upon detecting a change, the operator initiates actions to uphold the intended configuration of the resource. For instance, if a stateful set gains a new replica, the operator orchestrates the allocation of required resources and ensures seamless integration of the new replica into the set.

Operators employ a control loop mechanism to iteratively verify the status of the resources under their management, executing any required modifications as needed. This capability enables automatic handling of tasks such as scaling, failover, and other operations, eliminating the need for human intervention.

https://www.kubermatic.com/static/operator-blog-post.png

