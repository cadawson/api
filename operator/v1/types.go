package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets.
// ClusterManager will only be deployed in open-cluster-management-hub namespace.
type ClusterManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.
	Spec ClusterManagerSpec `json:"spec"`

	// Status represents the current status of controllers that govern the lifecycle of managed clusters.
	// +optional
	Status ClusterManagerStatus `json:"status,omitempty"`
}

// ClusterManagerSpec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.
type ClusterManagerSpec struct {
	// RegistrationImagePullSpec represents the desired image of registration controller installed on hub.
	// +required
	RegistrationImagePullSpec string `json:"registrationImagePullSpec"`
}

// ClusterManagerStatus represents the current status of the registration and work distribution controllers running on the hub.
type ClusterManagerStatus struct {
	// ObservedGeneration is the last generation change.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions contain the different condition statuses for this ClusterManager.
	// Valid condition types are:
	// Applied: Components in hub are applied.
	// Available: Components in hub are available and ready to serve.
	// Progressing: Components in hub are in a transitioning state.
	// Degraded: Components in hub do not match the desired configuration and only provide
	// degraded service.
	Conditions []StatusCondition `json:"conditions"`

	// Generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.
	// +optional
	Generations []GenerationStatus `json:"generations,omitempty"`

	// RelatedResources are used to track the resources that are related to this ClusterManager.
	// +optional
	RelatedResources []RelatedResourceMeta `json:"relatedResources,omitempty"`
}

// RelatedResourceMeta represents the resource that is managed by an operator.
type RelatedResourceMeta struct {
	// The group value is the group of the resource that you're tracking.
	// +required
	Group string `json:"group"`

	// The version value is the version of the resource that you're tracking.
	// +required
	Version string `json:"version"`

	// The resource value is the resource type of the resource that you're tracking.
	// +required
	Resource string `json:"resource"`

	// The namespace value is the location of the resource that you're tracking.
	// +optional
	Namespace string `json:"namespace"`

	// The name value is the name of the resource that you're tracking.
	// +required
	Name string `json:"name"`
}

// GenerationStatus keeps track of the generation for a given resource so that decisions about forced updates can be made.
// The definition matches the GenerationStatus defined in github.com/openshift/api/v1.
type GenerationStatus struct {
	// The group value is the group of the resource that you're tracking.
	// +required
	Group string `json:"group"`

	// The version value is the version of the resource that you're tracking.
	// +required
	Version string `json:"version"`

	// The resource value is the resource type of the resource that you're tracking.
	// +required
	Resource string `json:"resource"`

	// The namespace value is the location of the resource that you're tracking.
	// +optional
	Namespace string `json:"namespace"`

	// The name value is the name of the resource that you're tracking.
	// +required
	Name string `json:"name"`

	// The lastGeneration value is the last generation of the thing that the controller applies.
	// +required
	LastGeneration int64 `json:"lastGeneration" protobuf:"varint,5,opt,name=lastGeneration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterManagerList is a collection of deployment configurations for registration and work distribution controllers.
type ClusterManagerList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of deployment configurations for registration and work distribution controllers.
	Items []ClusterManager `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Klusterlet represents controllers on the managed cluster. When configured,
// the Klusterlet requires a secret named of bootstrap-hub-kubeconfig in the
// same namespace to allow API requests to the hub for the registration protocol.
type Klusterlet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents the desired deployment configuration of Klusterlet agent.
	Spec KlusterletSpec `json:"spec,omitempty"`

	// Status represents the current status of Klusterlet agent.
	Status KlusterletStatus `json:"status,omitempty"`
}

// KlusterletSpec represents the desired deployment configuration of Klusterlet agent.
type KlusterletSpec struct {
	// Namespace is the namespace to deploy the agent.
	// The namespace must have a prefix of "open-cluster-management-", and if it is not set,
	// the namespace of "open-cluster-management-agent" is used to deploy agent.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// RegistrationImagePullSpec represents the desired image configuration of registration agent.
	// +required
	RegistrationImagePullSpec string `json:"registrationImagePullSpec"`

	// WorkImagePullSpec represents the desired image configuration of work agent.
	// +required
	WorkImagePullSpec string `json:"workImagePullSpec,omitempty"`

	// ClusterName is the name of the managed cluster to be created on hub.
	// The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// ExternalServerURLs represents the a list of apiserver urls and ca bundles that is accessible externally.
	// If it is set empty, managed cluster has no externally accessible URL that the hub cluster can visit.
	// +optional
	ExternalServerURLs []ServerURL `json:"externalServerURLs,omitempty"`
}

// ServerURL represents the apiserver URL and ca bundle that is accessible externally
type ServerURL struct {
	// URL is the URL of apiserver endpoint of the managed cluster.
	// +required
	URL string `json:"url"`

	// CABundle is the ca bundle to connect to apiserver of the managed cluster.
	// System certs are used if it is not set.
	// +optional
	CABundle []byte `json:"caBundle,omitempty"`
}

// KlusterletStatus represents the current status of Klusterlet agent.
type KlusterletStatus struct {
	// ObservedGeneration is the last generation change.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions contain the different condition statuses for this Klusterlet.
	// Valid condition types are:
	// Applied: Components have been applied in the managed cluster.
	// Available: Components in the managed cluster are available and ready to serve.
	// Progressing: Components in the managed cluster are in a transitioning state.
	// Degraded: Components in the managed cluster do not match the desired configuration and only provide
	// degraded service.
	Conditions []StatusCondition `json:"conditions"`

	// Generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.
	// +optional
	Generations []GenerationStatus `json:"generations,omitempty"`

	// RelatedResources are used to track the resources that are related to this Klusterlet.
	// +optional
	RelatedResources []RelatedResourceMeta `json:"relatedResources,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KlusterletList is a collection of Klusterlet agents.
type KlusterletList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of Klusterlet agents.
	Items []Klusterlet `json:"items"`
}

// StatusCondition contains condition information.
type StatusCondition struct {
	// Type is the type of the cluster condition.
	// +required
	Type string `json:"type"`

	// Status is the status of the condition. Valid values are True, False, or Unknown.
	// +required
	Status metav1.ConditionStatus `json:"status"`

	// LastTransitionTime is the last time the condition changed from one status to another.
	// +required
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// Reason is a brief reason for the condition's last status change.
	// +required
	Reason string `json:"reason"`

	// Message is a human-readable message indicating details about the last status change.
	// +required
	Message string `json:"message"`
}
