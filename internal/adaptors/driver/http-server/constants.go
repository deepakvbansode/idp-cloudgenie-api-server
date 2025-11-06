package httpserver

const (
	HealthCheckURI     = "/v1/healthcheck"
	HealthCheckAPIName = "HealthCheck"
	
	BlueprintURI        = "/v1/blueprints"
	GetBlueprintAPIName    = "GetBlueprints"
	BlueprintDetailURI        = "/v1/blueprints/{id}"
	GetBlueprintDetailsAPIName    = "GetBlueprintDetails"

	ResourcesURI        = "/v1/resources"
	GetResourceAPIName    = "GetResources"
	ResourceDetailURI        = "/v1/resources/{id}"
	GetResourceDetailsAPIName    = "GetResourceDetails"
	CreateResourceAPIName    = "CreateResource"
	DeleteResourceAPIName    = "DeleteResource"
	GetResourceStatusURI        = "/v1/resources/{id}/status"
	GetResourceStatusAPIName    = "GetResourceStatus"
)	
