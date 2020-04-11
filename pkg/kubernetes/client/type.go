package client

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceName string

// nuwa resource
const (
	Water       ResourceName = "water"
	Deployment  ResourceName = "deployment"
	Stone       ResourceName = "stone"
	StatefulSet ResourceName = "statefulset"
	DaemonSet   ResourceName = "daemonsets"
	Injector    ResourceName = "injector"
	Pod         ResourceName = "pod"
	Job         ResourceName = "jobs"
	CronJobs    ResourceName = "cronjobs"
	Replicaset  ResourceName = "replicasets"
	Event       ResourceName = "events"
	Node        ResourceName = "nodes"
)

// GroupVersionResources describe resource collection
var GroupVersionResources = map[ResourceName]schema.GroupVersionResource{
	Water:       {Group: "nuwa.nip.io", Version: "v1", Resource: "waters"},
	Deployment:  {Group: "apps", Version: "v1", Resource: "deployments"},
	Stone:       {Group: "nuwa.nip.io", Version: "v1", Resource: "stones"},
	StatefulSet: {Group: "apps", Version: "v1", Resource: "statefulsets"},
	DaemonSet:   {Group: "apps", Version: "v1", Resource: "daemonsets"},
	Injector:    {Group: "nuwa.nip.io", Version: "v1", Resource: "injectors"},
	Pod:         {Group: "", Version: "v1", Resource: "pods"},
	Node:        {Group: "", Version: "v1", Resource: "nodes"},
	Event:       {Group: "", Version: "v1", Resource: "events"},
	Job:         {Group: "batch", Version: "v1", Resource: "jobs"},
	CronJobs:    {Group: "batch", Version: "v1beta1", Resource: "cronjobs"},
	Replicaset:  {Group: "apps", Version: "v1", Resource: "replicasets"},
}

var (
	ResourceWater       = GetGVR(Water)
	ResourceJob         = GetGVR(Job)
	ResourceDeployment  = GetGVR(Deployment)
	ResourceStone       = GetGVR(Stone)
	ResourceStatefulSet = GetGVR(StatefulSet)
	ResourceDaemonSet   = GetGVR(DaemonSet)
	ResourceInjector    = GetGVR(Injector)
	ResourcePod         = GetGVR(Pod)
	ResourceCronJobs    = GetGVR(CronJobs)
	ResourceReplicaset  = GetGVR(Replicaset)
	ResourceEvent       = GetGVR(Event)
	ResourceNode        = GetGVR(Node)
)

func GetGVR(rs ResourceName) schema.GroupVersionResource {
	gvr, exist := GroupVersionResources[rs]
	if !exist {
		panic("try to get an undefined resource")
	}
	return gvr
}
