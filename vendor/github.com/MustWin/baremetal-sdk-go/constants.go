// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

type resourceName string

type InstanceActions string
type instanceStates string
type NetworkEntityType string
type DBNodeAction string
type DatabaseEdition string
type DiskRedundancy string
type ListObjectOptionField string

const (
	// Resource States
	ResourceActive                = "ACTIVE"
	ResourceAttached              = "ATTACHED"
	ResourceAttaching             = "ATTACHING"
	ResourceAvailable             = "AVAILABLE"
	ResourceCreated               = "CREATED"
	ResourceCreating              = "CREATING"
	ResourceCreatingImage         = "CREATING_IMAGE"
	ResourceDeleting              = "DELETING"
	ResourceDeleted               = "DELETED"
	ResourceDetached              = "DETACHED"
	ResourceDetaching             = "DETACHING"
	ResourceDisabled              = "DISABLED"
	ResourceDown                  = "DOWN"
	ResourceDownForMaintenance    = "DOWN_FOR_MAINTENANCE"
	ResourceFailed                = "FAILED"
	ResourceFaulty                = "FAULTY"
	ResourceGettingHistory        = "GETTING-HISTORY"
	ResourceInactive              = "INACTIVE"
	ResourceProvisioning          = "PROVISIONING"
	ResourceRequested             = "REQUESTED"
	ResourceRequestReceived       = "REQUEST_RECEIVED"
	ResourceRestoring             = "RESTORING"
	ResourceRunning               = "RUNNING"
	ResourceStarting              = "STARTING"
	ResourceStopped               = "STOPPED"
	ResourceStopping              = "STOPPING"
	ResourceSucceeded             = "SUCCEEDED"
	ResourceTerminated            = "TERMINATED"
	ResourceTerminating           = "TERMINATING"
	ResourceUp                    = "UP"
	ResourceWaitingForWorkRequest = "WAITING_FOR_WORK_REQUEST"
	ResourceSucceededWorkRequest = "SUCCEEDED_WORK_REQUEST"

	WorkRequestAccepted   = "ACCEPTED"
	WorkRequestInProgress = "IN_PROGRESS"
	WorkRequestFailed     = "FAILED"
	WorkRequestSucceeded  = "SUCCEEDED"

	// Error codes
	UserAlreadyExists       = "UserAlreadyExists"
	InvalidParameter        = "InvalidParameter"
	NotAuthorizedOrNotFound = "NotAuthorizedOrNotFound"

	SDKVersion  = "20160918"
	SDKVersion2 = "20170115"

	identityServiceAPI        = "https://identity.us-phoenix-1.oraclecloud.com"
	identityServiceAPIVersion = SDKVersion

	coreServiceAPI        = "https://iaas.us-phoenix-1.oraclecloud.com"
	coreServiceAPIVersion = SDKVersion

	databaseServiceAPI        = "https://database.us-phoenix-1.oraclecloud.com"
	databaseServiceAPIVersion = SDKVersion

	objectStorageServiceAPI        = "https://objectstorage.us-phoenix-1.oraclecloud.com"
	objectStorageServiceAPIVersion = SDKVersion

	loadBalancerServiceAPI        = "https://iaas.us-phoenix-1.oraclecloud.com"
	loadBalancerServiceAPIVersion = SDKVersion2

	// Header Keys
	headerBytesRemaining     = "opc-bytes-remaining"
	headerContentEncoding    = "Content-Encoding"
	headerContentLanguage    = "Content-Language"
	headerContentLength      = "Content-Length"
	headerContentMD5         = "Content-MD5"
	headerContentType        = "Content-Type"
	headerETag               = "ETag"
	headerLastModified       = "last-modified"
	headerOPCClientRequestID = "opc-client-request-id"
	headerOPCWorkRequestID   = "opc-work-request-id"
	headerOPCNextPage        = "opc-next-page"
	headerOPCRequestID       = "opc-request-id"

	// Actions that can be applied to compute instances
	actionStart InstanceActions = "START"
	actionStop  InstanceActions = "STOP"
	actionReset InstanceActions = "RESET"

	// Network entity types for routing rules
	networkEntityVnic                      NetworkEntityType = "VNIC"
	networkEntityInternetGateway           NetworkEntityType = "INTERNET_GATEWAY"
	networkEntityDynamicallyRoutingGateway NetworkEntityType = "DYNAMICALLY_ROUTING_GATEWAY"

	// Database Node actions
	DBNodeActionStart     DBNodeAction = "START"
	DBNodeActionStop      DBNodeAction = "STOP"
	DBNodeActionReset     DBNodeAction = "RESET"
	DBNodeActionSoftReset DBNodeAction = "SOFTRESET"

	// Database editions
	DatabaseEditionEnterprise DatabaseEdition = "ENTERPRISE_EDITION"
	DatabaseEditionExtreme    DatabaseEdition = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	DatabaseEditionHigh       DatabaseEdition = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	DatabaseEditionStandard   DatabaseEdition = "STANDARD_EDITION"

	// Database disk redundancy levels
	DiskRedundancyHigh   DiskRedundancy = "HIGH"
	DiskRedundancyNormal DiskRedundancy = "NORMAL"

	// DB Resources
	resourceDBHomes               resourceName = "dbHomes"
	resourceDBNodes               resourceName = "dbNodes"
	resourceDBSystems             resourceName = "dbSystems"
	resourceDBSystemShapes        resourceName = "dbSystemShapes"
	resourceDBVersions            resourceName = "dbVersions"
	resourceDatabases             resourceName = "databases"
	resourceDBSupportedOperations resourceName = "supportedOperations"

	// Identity Resources
	resourceAvailabilityDomains  resourceName = "availabilityDomains"
	resourceCompartments         resourceName = "compartments"
	resourceGroups               resourceName = "groups"
	resourcePolicies             resourceName = "policies"
	resourceUiPassword           resourceName = "uiPassword"
	resourceUsers                resourceName = "users"
	resourceUserGroupMemberships resourceName = "userGroupMemberships"
	resourceSwiftPasswords       resourceName = "swiftPasswords"

	// Core Resources
	resourceCustomerPremiseEquipment resourceName = "cpes"
	resourceDHCPOptions              resourceName = "dhcps"
	resourceDrgAttachments           resourceName = "drgAttachments"
	resourceDrgs                     resourceName = "drgs"
	resourceImages                   resourceName = "images"
	resourceInstanceConsoleHistories resourceName = "instanceConsoleHistories"
	resourceInstances                resourceName = "instances"
	resourceInternetGateways         resourceName = "internetGateways"
	resourceIPSecConnections         resourceName = "ipsecConnections"
	resourceRouteTables              resourceName = "routeTables"
	resourceSecurityLists            resourceName = "securityLists"
	resourceShapes                   resourceName = "shapes"
	resourceSubnets                  resourceName = "subnets"
	resourceVirtualNetworks          resourceName = "vcns"
	resourceVnics                    resourceName = "vnics"
	resourceVnicAttachments          resourceName = "vnicAttachments"
	resourceVolumes                  resourceName = "volumes"
	resourceVolumeAttachments        resourceName = "volumeAttachments"
	resourceVolumeBackups            resourceName = "volumeBackups"

	// LoadBalancer Resources
	resourceLoadBalancers            resourceName = "loadBalancers"
	resourceBackends                 resourceName = "backends"
	resourceBackendSets              resourceName = "backendSets"
	resourceCertificates             resourceName = "certificates"
	resourceHealthChecker            resourceName = "healthChecker"
	resourceListeners                resourceName = "listeners"
	resourceLoadBalancerPolicies     resourceName = "loadBalancerPolicies"
	resourceLoadBalancerProtocols    resourceName = "loadBalancerProtocols"
	resourceLoadBalancerShapes       resourceName = "loadBalancerShapes"
	resourceLoadBalancerWorkRequests resourceName = "loadBalancerWorkRequests"
	resourceWorkRequests             resourceName = "workRequests"

	apiKeys      = "apiKeys"
	uiPassword   = "uiPassword"
	deviceConfig = "deviceConfig"
	deviceStatus = "deviceStatus"
	dataURLPart  = "data"

	// Object Storage Resources
	resourceNamespaces = "n"
	resourceBuckets    = "b"
	resourceObjects    = "o"
)
