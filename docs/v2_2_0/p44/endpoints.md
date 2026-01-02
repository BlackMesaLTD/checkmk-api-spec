# API Endpoints - v2.2.0p44

## Table of Contents

- [Acknowledge problems](#acknowledge-problems)
- [Activate changes](#activate-changes)
- [Agents](#agents)
- [Autocomplete (internal)](#autocomplete-(internal))
- [Aux Tags](#aux-tags)
- [Business intelligence (BI)](#business-intelligence-(bi))
- [Certificates](#certificates)
- [Comments](#comments)
- [Contact groups](#contact-groups)
- [Downtimes](#downtimes)
- [Event Console](#event-console)
- [Folders](#folders)
- [Host groups](#host-groups)
- [Host tag groups](#host-tag-groups)
- [Hosts](#hosts)
- [Hosts (internal)](#hosts-(internal))
- [Metrics](#metrics)
- [Miscellaneous](#miscellaneous)
- [Notification Rules](#notification-rules)
- [Passwords](#passwords)
- [Rules](#rules)
- [Rulesets](#rulesets)
- [Service discovery](#service-discovery)
- [Service groups](#service-groups)
- [Service status](#service-status)
- [Site Management](#site-management)
- [Time periods](#time-periods)
- [User Roles](#user-roles)
- [Users](#users)

---

## Acknowledge problems

### POST /domain-types/acknowledge/collections/host

Set acknowledgement on related hosts

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [AcknowledgeHostRelatedProblem](schemas.md#acknowledgehostrelatedproblem)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 422 | Unprocessable Entity: The query yielded no result. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/acknowledge/collections/service

Set acknowledgement on related services

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [AcknowledgeServiceRelatedProblem](schemas.md#acknowledgeservicerelatedproblem)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 422 | Unprocessable Entity: Service was not in a problem state. | [ApiError](schemas.md#apierror) |

---

## Activate changes

### POST /domain-types/activation_run/actions/activate-changes/invoke

Activate pending changes

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [ActivateChanges](schemas.md#activatechanges)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: Activation has been started, but not completed (if you need to wait for completion, see documentation for this endpoint). | [ActivationRunResponse](schemas.md#activationrunresponse) |
| 302 | Found: The activation has been started and is still running. Redirecting to the 'Wait for completion' endpoint. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 401 | Unauthorized: The API user may not activate another users changes, or the user may and activation was not forced explicitly. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Activation not possible because of licensing issues. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: Some sites could not be activated. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 422 | Unprocessable Entity: There are no changes to be activated. | [ApiError](schemas.md#apierror) |
| 423 | Locked: There is already an activation running. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/activation_run/collections/pending_changes

Show all pending changes

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [PendingChangesCollection](schemas.md#pendingchangescollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/activation_run/collections/running

Show all currently running activations

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ActivationRunCollection](schemas.md#activationruncollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/activation_run/{activation_id}

Show the activation status

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `activation_id` | path | string | Yes | The activation-id. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ActivationRunResponse](schemas.md#activationrunresponse) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: There is no running activation with this activation_id. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/activation_run/{activation_id}/actions/wait-for-completion/invoke

Wait for activation completion

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `activation_id` | path | string | Yes | The activation-id. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: The activation has been completed. | - |
| 302 | Found: The activation is still running. Redirecting to the 'Wait for completion' endpoint. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: There is no running activation with this activation_id. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Agents

### GET /domain-types/agent/actions/download/invoke

Download agents shipped with Checkmk

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `os_type` | query | string | Yes | The type of the operating system. May be one of linux_rpm, linux_deb, windows... |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Autocomplete (internal)

### POST /objects/autocomplete/{autocomplete_id}

Call the autocompleter specified in the url

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `autocomplete_id` | path | string | Yes | The id of the autocompleter |

**Request Body:** [Request](schemas.md#request)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [Response](schemas.md#response) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Aux Tags

### GET /domain-types/aux_tag/collections/all

Show Auxiliary Tags

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [AuxTagResponseCollection](schemas.md#auxtagresponsecollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/aux_tag/collections/all

Create an Auxiliary Tag

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [AuxTagAttrsCreate](schemas.md#auxtagattrscreate)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [AuxTagResponse](schemas.md#auxtagresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/aux_tag/{aux_tag_id}

Update an aux tag

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `aux_tag_id` | path | string | Yes | An auxiliary tag id |

**Request Body:** [AuxTagAttrsUpdate](schemas.md#auxtagattrsupdate)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [AuxTagResponse](schemas.md#auxtagresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/aux_tag/{aux_tag_id}

Show an Auxiliary Tag

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `aux_tag_id` | path | string | Yes | An auxiliary tag id |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [AuxTagResponse](schemas.md#auxtagresponse) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/aux_tag/{aux_tag_id}/actions/delete/invoke

Delete an Auxiliary Tag

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `aux_tag_id` | path | string | Yes | An auxiliary tag id |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Business intelligence (BI)

### GET /domain-types/bi_aggregation/actions/aggregation_state/invoke

Get the state of BI aggregations

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `filter_names` | query | array | No | Filter by names |
| `filter_groups` | query | array | No | Filter by group |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIAggregationStateResponse](schemas.md#biaggregationstateresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/bi_aggregation/actions/aggregation_state/invoke

Get the state of BI aggregations

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BIAggregationStateRequest](schemas.md#biaggregationstaterequest)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIAggregationStateResponse](schemas.md#biaggregationstateresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/bi_pack/collections/all

Show all BI packs

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObjectCollection](schemas.md#domainobjectcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/bi_aggregation/{aggregation_id}

Create a BI aggregation

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `aggregation_id` | path | string | Yes | The unique id for the aggregation |

**Request Body:** [BIAggregationEndpoint](schemas.md#biaggregationendpoint)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIAggregationEndpoint](schemas.md#biaggregationendpoint) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/bi_aggregation/{aggregation_id}

Update an existing BI aggregation

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `aggregation_id` | path | string | Yes | The unique id for the aggregation |

**Request Body:** [BIAggregationEndpoint](schemas.md#biaggregationendpoint)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIAggregationEndpoint](schemas.md#biaggregationendpoint) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/bi_aggregation/{aggregation_id}

Delete a BI aggregation

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `aggregation_id` | path | string | Yes | The unique id for the aggregation |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/bi_aggregation/{aggregation_id}

Get a BI aggregation

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `aggregation_id` | path | string | Yes | The unique id for the aggregation |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIAggregationEndpoint](schemas.md#biaggregationendpoint) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/bi_pack/{pack_id}

Create a new BI pack

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `pack_id` | path | string | Yes | The unique id for the aggregation pack |

**Request Body:** [BIPackEndpoint](schemas.md#bipackendpoint)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIPackEndpoint](schemas.md#bipackendpoint) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/bi_pack/{pack_id}

Get a BI pack and its rules and aggregations

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `pack_id` | path | string | Yes | The unique id for the aggregation pack |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/bi_pack/{pack_id}

Delete BI pack

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `pack_id` | path | string | Yes | The unique id for the aggregation pack |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/bi_pack/{pack_id}

Update an existing BI pack

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `pack_id` | path | string | Yes | The unique id for the aggregation pack |

**Request Body:** [BIPackEndpoint](schemas.md#bipackendpoint)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIPackEndpoint](schemas.md#bipackendpoint) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/bi_rule/{rule_id}

Update an existing BI rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `rule_id` | path | string | Yes | The unique id for the rule |

**Request Body:** [BIRuleEndpoint](schemas.md#biruleendpoint)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIRuleEndpoint](schemas.md#biruleendpoint) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /objects/bi_rule/{rule_id}

Create a new BI rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `rule_id` | path | string | Yes | The unique id for the rule |

**Request Body:** [BIRuleEndpoint](schemas.md#biruleendpoint)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIRuleEndpoint](schemas.md#biruleendpoint) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/bi_rule/{rule_id}

Show a BI rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `rule_id` | path | string | Yes | The unique id for the rule |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [BIRuleEndpoint](schemas.md#biruleendpoint) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/bi_rule/{rule_id}

Delete BI rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `rule_id` | path | string | Yes | The unique id for the rule |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |

---

## Certificates

### GET /agent_controller_certificates_settings

Show agent controller certificate settings

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [AgentControllerCertificateSettings](schemas.md#agentcontrollercertificatesettings) |
| 403 | Forbidden: Unauthorized to read the global settings | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /csr

X.509 PEM-encoded Certificate Signing Requests (CSRs)

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [X509ReqPEMUUID](schemas.md#x509reqpemuuid)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [X509PEM](schemas.md#x509pem) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: You do not have the permission for agent pairing. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /root_cert

X.509 PEM-encoded root certificate

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [X509PEM](schemas.md#x509pem) |
| 403 | Forbidden: You do not have the permission for agent pairing. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Comments

### POST /domain-types/comment/actions/delete/invoke

Delete comments

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [DeleteComments](schemas.md#deletecomments)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/comment/collections/host

Create a host comment

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateHostRelatedComment](schemas.md#createhostrelatedcomment)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/comment/collections/service

Create a service comment

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateServiceRelatedComment](schemas.md#createservicerelatedcomment)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/comment/collections/{collection_name}

Show comments

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `query` | query | string | No | An query expression of the Livestatus 'comments' table in nested dictionary f... |
| `service_description` | query | string | No | The service description. No exception is raised when the specified service de... |
| `host_name` | query | string | No | The host name. No exception is raised when the specified host name does not e... |
| `site_id` | query | string | No | An existing site id |
| `collection_name` | path | string | Yes | Do you want to get comments from 'hosts', 'services' or 'all' |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [CommentCollection](schemas.md#commentcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/comment/{comment_id}

Show a comment

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `site_id` | query | string | Yes | An existing site id |
| `comment_id` | path | integer | Yes | An existing comment's ID |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [CommentObject](schemas.md#commentobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Contact groups

### POST /domain-types/contact_group_config/actions/bulk-create/invoke

Bulk create contact groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkInputContactGroup](schemas.md#bulkinputcontactgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ContactGroupCollection](schemas.md#contactgroupcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/contact_group_config/actions/bulk-delete/invoke

Bulk delete contact groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkDeleteContactGroup](schemas.md#bulkdeletecontactgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /domain-types/contact_group_config/actions/bulk-update/invoke

Bulk update contact groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkUpdateContactGroup](schemas.md#bulkupdatecontactgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ContactGroupCollection](schemas.md#contactgroupcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/contact_group_config/collections/all

Show all contact groups

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ContactGroupCollection](schemas.md#contactgroupcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/contact_group_config/collections/all

Create a contact group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [InputContactGroup](schemas.md#inputcontactgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/contact_group_config/{name}

Delete a contact group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The identifier name of the group. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |

---

### GET /objects/contact_group_config/{name}

Show a contact group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The identifier name of the group. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ContactGroup](schemas.md#contactgroup) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/contact_group_config/{name}

Update a contact group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `name` | path | string | Yes | The identifier name of the group. |

**Request Body:** [UpdateGroup](schemas.md#updategroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ContactGroup](schemas.md#contactgroup) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

## Downtimes

### POST /domain-types/downtime/actions/delete/invoke

Delete a scheduled downtime

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [DeleteDowntime](schemas.md#deletedowntime)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Delete downtimes commands have been sent to Livestatus. The REST API exclusively manages the preparation and dispatch of commands to Livestatus. These commands are processed in an asynchronous manner, and the REST API does not validate the successful execution of commands on Livestatus. To investigate any failures in Livestatus, one should refer to the corresponding log. Also you can refer to [Queries through the REST API](#section/Queries-through-the-REST-API) section for further information. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /domain-types/downtime/actions/modify/invoke

Modify a scheduled downtime

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [ModifyDowntime](schemas.md#modifydowntime)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Update downtimes commands have been sent to Livestatus. The REST API exclusively manages the preparation and dispatch of commands to Livestatus. These commands are processed in an asynchronous manner, and the REST API does not validate the successful execution of commands on Livestatus. To investigate any failures in Livestatus, one should refer to the corresponding log. Also you can refer to [Queries through the REST API](#section/Queries-through-the-REST-API) section for further information. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/downtime/collections/all

Show all scheduled downtimes

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `query` | query | string | No | An query expression of the Livestatus 'downtimes' table in nested dictionary ... |
| `service_description` | query | string | No | The service description. No exception is raised when the specified service de... |
| `host_name` | query | string | No | The host name. No exception is raised when the specified host name does not e... |
| `downtime_type` | query | string | No | The type of the downtime to be listed. Only filters the result when using the... |
| `site_id` | query | string | No | An existing site id |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DowntimeCollection](schemas.md#downtimecollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/downtime/collections/host

Create a host related scheduled downtime

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateHostRelatedDowntime](schemas.md#createhostrelateddowntime)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Create host related downtimes commands have been sent to Livestatus. The REST API exclusively manages the preparation and dispatch of commands to Livestatus. These commands are processed in an asynchronous manner, and the REST API does not validate the successful execution of commands on Livestatus. To investigate any failures in Livestatus, one should refer to the corresponding log. Also you can refer to [Queries through the REST API](#section/Queries-through-the-REST-API) section for further information. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 422 | Unprocessable Entity: The request could not be processed. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/downtime/collections/service

Create a service related scheduled downtime

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateServiceRelatedDowntime](schemas.md#createservicerelateddowntime)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Create service related downtimes commands have been sent to Livestatus. The REST API exclusively manages the preparation and dispatch of commands to Livestatus. These commands are processed in an asynchronous manner, and the REST API does not validate the successful execution of commands on Livestatus. To investigate any failures in Livestatus, one should refer to the corresponding log. Also you can refer to [Queries through the REST API](#section/Queries-through-the-REST-API) section for further information. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 422 | Unprocessable Entity: The request could not be processed. | [ApiError](schemas.md#apierror) |

---

### GET /objects/downtime/{downtime_id}

Show downtime

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `site_id` | query | string | Yes | An existing site id |
| `downtime_id` | path | string | Yes | The id of the downtime |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DowntimeObject](schemas.md#downtimeobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Event Console

### POST /domain-types/event_console/actions/change_state/invoke

Change multiple event states

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [ChangeEventStateSelector](schemas.md#changeeventstateselector)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/event_console/actions/delete/invoke

Archive events

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [DeleteECEvents](schemas.md#deleteecevents)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/event_console/actions/update_and_acknowledge/invoke

Update and acknowledge events

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [UpdateAndAcknowledgeSelector](schemas.md#updateandacknowledgeselector)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/event_console/collections/all

Show events

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `query` | query | string | No | An query expression of the Livestatus 'eventconsoleevents' table in nested di... |
| `host` | query | string | No | The host name. No exception is raised when the specified host name does not e... |
| `application` | query | string | No | Show events that originated from this app. |
| `state` | query | string | No | The state |
| `phase` | query | string | No | The event phase, open or ack |
| `site_id` | query | string | No | An existing site id |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [EventConsoleResponseCollection](schemas.md#eventconsoleresponsecollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/event_console/{event_id}

Show an event

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `site_id` | query | string | Yes | An existing site id |
| `event_id` | path | string | Yes | The event console event ID. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ECEventResponse](schemas.md#eceventresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/event_console/{event_id}/actions/change_state/invoke

Change event state

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `event_id` | path | string | Yes | The event console event ID. |

**Request Body:** [ChangeEventState](schemas.md#changeeventstate)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /objects/event_console/{event_id}/actions/update_and_acknowledge/invoke

Update and acknowledge an event

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `event_id` | path | string | Yes | The event console event ID. |

**Request Body:** [UpdateAndAcknowledeEventSiteIDRequired](schemas.md#updateandacknowledeeventsiteidrequired)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Folders

### PUT /domain-types/folder_config/actions/bulk-update/invoke

Bulk update folders

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkUpdateFolder](schemas.md#bulkupdatefolder)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [FolderCollection](schemas.md#foldercollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/folder_config/collections/all

Create a folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateFolder](schemas.md#createfolder)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [Folder](schemas.md#folder) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/folder_config/collections/all

Show all folders

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `parent` | query | string | No | Show all sub-folders of this folder. The default is the root-folder. Path del... |
| `recursive` | query | boolean | No | List the folder (default: root) and all its sub-folders recursively. |
| `show_hosts` | query | boolean | No | When set, all hosts that are stored in each folder will also be shown. On lar... |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [FolderCollection](schemas.md#foldercollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/folder_config/{folder}

Update a folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `folder` | path | string | Yes | The path of the folder being requested. Please be aware that slashes can't be... |

**Request Body:** [UpdateFolder](schemas.md#updatefolder)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [Folder](schemas.md#folder) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/folder_config/{folder}

Delete a folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `delete_mode` | query | string | No | Delete policy: 'recursive': Deletes the folder and all the elements it contai... |
| `folder` | path | string | Yes | The path of the folder being requested. Please be aware that slashes can't be... |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 401 | Unauthorized: The user is not authorized to do this request. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |

---

### GET /objects/folder_config/{folder}

Show a folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `show_hosts` | query | boolean | No | When set, all hosts that are stored in this folder will also be shown. On lar... |
| `folder` | path | string | Yes | The path of the folder being requested. Please be aware that slashes can't be... |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [Folder](schemas.md#folder) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/folder_config/{folder}/actions/move/invoke

Move a folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `folder` | path | string | Yes | The path of the folder being requested. Please be aware that slashes can't be... |

**Request Body:** [MoveFolder](schemas.md#movefolder)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [Folder](schemas.md#folder) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### GET /objects/folder_config/{folder}/collections/hosts

Show all hosts in a folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `folder` | path | string | Yes | The path of the folder being requested. Please be aware that slashes can't be... |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfigCollection](schemas.md#hostconfigcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Host groups

### POST /domain-types/host_group_config/actions/bulk-create/invoke

Bulk create host groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkInputHostGroup](schemas.md#bulkinputhostgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostGroupCollection](schemas.md#hostgroupcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/host_group_config/actions/bulk-delete/invoke

Bulk delete host groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkDeleteHostGroup](schemas.md#bulkdeletehostgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /domain-types/host_group_config/actions/bulk-update/invoke

Bulk update host groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkUpdateHostGroup](schemas.md#bulkupdatehostgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostGroupCollection](schemas.md#hostgroupcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/host_group_config/collections/all

Show all host groups

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostGroupCollection](schemas.md#hostgroupcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/host_group_config/collections/all

Create a host group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [InputHostGroup](schemas.md#inputhostgroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostGroup](schemas.md#hostgroup) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_group_config/{name}

Update a host group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `name` | path | string | Yes | The identifier name of the group. |

**Request Body:** [UpdateGroup1](schemas.md#updategroup1)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostGroup](schemas.md#hostgroup) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/host_group_config/{name}

Delete a host group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The identifier name of the group. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |

---

### GET /objects/host_group_config/{name}

Show a host group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The identifier name of the group. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostGroup](schemas.md#hostgroup) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Host tag groups

### POST /domain-types/host_tag_group/collections/all

Create a host tag group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [InputHostTagGroup](schemas.md#inputhosttaggroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ConcreteHostTagGroup](schemas.md#concretehosttaggroup) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/host_tag_group/collections/all

Show all host tag groups

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostTagGroupCollection](schemas.md#hosttaggroupcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/host_tag_group/{name}

Delete a host tag group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `repair` | query | boolean | No | The host tag group can still be in use. Setting repair to True gives permissi... |
| `name` | path | string | Yes | The name of the host tag group |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 401 | Unauthorized: The user is not authorized to do this request. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 405 | Method Not Allowed: Method not allowed: This request is only allowed with other HTTP methods | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/host_tag_group/{name}

Show a host tag group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The name of the host tag group |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ConcreteHostTagGroup](schemas.md#concretehosttaggroup) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_tag_group/{name}

Update a host tag group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `name` | path | string | Yes | The name of the host tag group |

**Request Body:** [UpdateHostTagGroup](schemas.md#updatehosttaggroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ConcreteHostTagGroup](schemas.md#concretehosttaggroup) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 401 | Unauthorized: The user is not authorized to do this request. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 405 | Method Not Allowed: Method not allowed: This request is only allowed with other HTTP methods | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

## Hosts

### POST /domain-types/host_config/actions/bulk-create/invoke

Bulk create hosts

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `bake_agent` | query | boolean | No | Tries to bake the agents for the just created hosts. This process is started ... |

**Request Body:** [BulkCreateHost](schemas.md#bulkcreatehost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfigCollection](schemas.md#hostconfigcollection) |
| 400 | Bad Request: Parameter or validation failure. | [BulkHostActionWithFailedHosts](schemas.md#bulkhostactionwithfailedhosts) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/host_config/actions/bulk-delete/invoke

Bulk delete hosts

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkDeleteHost](schemas.md#bulkdeletehost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /domain-types/host_config/actions/bulk-update/invoke

Bulk update hosts

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkUpdateHost](schemas.md#bulkupdatehost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfigCollection](schemas.md#hostconfigcollection) |
| 400 | Bad Request: Parameter or validation failure. | [BulkHostActionWithFailedHosts](schemas.md#bulkhostactionwithfailedhosts) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/host_config/collections/all

Create a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `bake_agent` | query | boolean | No | Tries to bake the agents for the just created hosts. This process is started ... |

**Request Body:** [CreateHost](schemas.md#createhost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfig](schemas.md#hostconfig) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/host_config/collections/all

Show all hosts

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `effective_attributes` | query | boolean | No | Show all effective attributes on hosts, not just the attributes which were se... |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfigCollection](schemas.md#hostconfigcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/host_config/collections/clusters

Create a cluster host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `bake_agent` | query | boolean | No | Tries to bake the agents for the just created hosts. This process is started ... |

**Request Body:** [CreateClusterHost](schemas.md#createclusterhost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfig](schemas.md#hostconfig) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/host_config/{host_name}

Delete a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `host_name` | path | string | Yes | A hostname. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/host_config/{host_name}

Show a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `effective_attributes` | query | boolean | No | Show all effective attributes on hosts, not just the attributes which were se... |
| `host_name` | path | string | Yes | A hostname. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfig](schemas.md#hostconfig) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_config/{host_name}

Update a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A hostname. |

**Request Body:** [UpdateHost](schemas.md#updatehost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfig](schemas.md#hostconfig) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### POST /objects/host_config/{host_name}/actions/move/invoke

Move a host to another folder

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A hostname. |

**Request Body:** [MoveHost](schemas.md#movehost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfig](schemas.md#hostconfig) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_config/{host_name}/actions/rename/invoke

Rename a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A hostname. |

**Request Body:** [RenameHost](schemas.md#renamehost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfig](schemas.md#hostconfig) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: There are pending changes not yet activated. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 422 | Unprocessable Entity: The host could not be renamed. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_config/{host_name}/properties/nodes

Update the nodes of a cluster host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A cluster host. |

**Request Body:** [UpdateNodes](schemas.md#updatenodes)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ObjectProperty](schemas.md#objectproperty) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

## Hosts (internal)

### GET /objects/host_config_internal/{host_name}

Show a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `host_name` | path | string | Yes | A hostname. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [HostConfigSchemaInternal](schemas.md#hostconfigschemainternal) |
| 401 | Unauthorized: You do not have read access to this host. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_config_internal/{host_name}/actions/link_uuid/invoke

Link a host to a UUID

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A hostname. |

**Request Body:** [LinkHostUUID](schemas.md#linkhostuuid)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 401 | Unauthorized: You do not have the permissions to edit this host. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host_config_internal/{host_name}/actions/register/invoke

Register an existing host, ie. link it to a UUID

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A hostname. |

**Request Body:** [RegisterHost](schemas.md#registerhost)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ConnectionMode](schemas.md#connectionmode) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: You do not have the permissions to register this host. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 405 | Method Not Allowed: This host cannot be registered on this site. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Metrics

### POST /domain-types/metric/actions/get/invoke

Get metrics

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [Get](schemas.md#get)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [GraphCollection](schemas.md#graphcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Miscellaneous

### GET /version

Display some version information

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [InstalledVersions](schemas.md#installedversions) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Notification Rules

### GET /domain-types/notification_rule/collections/all

Show all notification rules

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [NotificationRuleResponseCollection](schemas.md#notificationruleresponsecollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/notification_rule/collections/all

Create a notification rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [NotificationRuleRequest](schemas.md#notificationrulerequest)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [NotificationRuleResponse](schemas.md#notificationruleresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/notification_rule/{rule_id}

Update a notification rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `rule_id` | path | string | Yes | The notification rule ID. |

**Request Body:** [NotificationRuleRequest](schemas.md#notificationrulerequest)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [NotificationRuleResponse](schemas.md#notificationruleresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/notification_rule/{rule_id}

Show a notification rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `rule_id` | path | string | Yes | The notification rule ID. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [NotificationRuleResponse](schemas.md#notificationruleresponse) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/notification_rule/{rule_id}/actions/delete/invoke

Delete a notification rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `rule_id` | path | string | Yes | The notification rule ID. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Passwords

### POST /domain-types/password/collections/all

Create a password

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [InputPassword](schemas.md#inputpassword)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [PasswordObject](schemas.md#passwordobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/password/collections/all

Show all passwords

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [PasswordCollection](schemas.md#passwordcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/password/{name}

Update a password

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `name` | path | string | Yes | A name used as an identifier. Can be of arbitrary (sensible) length. |

**Request Body:** [UpdatePassword](schemas.md#updatepassword)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [PasswordObject](schemas.md#passwordobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/password/{name}

Delete a password

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | A name used as an identifier. Can be of arbitrary (sensible) length. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/password/{name}

Show a password

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | A name used as an identifier. Can be of arbitrary (sensible) length. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [PasswordObject](schemas.md#passwordobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Rules

### POST /domain-types/rule/collections/all

Create rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [InputRuleObject](schemas.md#inputruleobject)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RuleObject](schemas.md#ruleobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/rule/collections/all

List rules

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `ruleset_name` | query | string | Yes | The name of the ruleset. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RuleCollection](schemas.md#rulecollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/rule/{rule_id}

Delete a rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `rule_id` | path | string | Yes | The ID of the rule. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Rule was deleted successfully. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The rule to be deleted was not found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/rule/{rule_id}

Modify a rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `rule_id` | path | string | Yes | The ID of the rule. |

**Request Body:** [UpdateRuleObject](schemas.md#updateruleobject)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RuleObject](schemas.md#ruleobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### GET /objects/rule/{rule_id}

Show a rule

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `rule_id` | path | string | Yes | The ID of the rule. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RuleObject](schemas.md#ruleobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/rule/{rule_id}/actions/move/invoke

Move a rule to a specific location

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `rule_id` | path | string | Yes | The ID of the rule. |

**Request Body:** [MoveRuleTo](schemas.md#moveruleto)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RuleObject](schemas.md#ruleobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

## Rulesets

### GET /domain-types/ruleset/collections/all

Search rule sets

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `fulltext` | query | string | No | Search all keys (like `name`, `title`, `help`, etc.) for this text. Regex all... |
| `folder` | query | string | No | The folder in which to search for rules. Path delimiters can be either `~`, `... |
| `deprecated` | query | boolean | No | Only show deprecated rulesets. Defaults to False. |
| `used` | query | boolean | No | Only show used rulesets. Defaults to True. |
| `name` | query | string | No | A regex of the name. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RulesetCollection](schemas.md#rulesetcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/ruleset/{ruleset_name}

Show a ruleset

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `ruleset_name` | path | string | Yes | The name of the ruleset. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [RulesetObject](schemas.md#rulesetobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Service discovery

### POST /domain-types/discovery_run/actions/bulk-discovery-start/invoke

Start a bulk discovery job

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkDiscovery](schemas.md#bulkdiscovery)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DiscoveryBackgroundJobStatusObject](schemas.md#discoverybackgroundjobstatusobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: A bulk discovery job is already active | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/service/collections/services

Show all services of specific phase

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `host_name` | query | string | Yes | The host of the discovered services. |
| `discovery_phase` | query | string | Yes | The discovery phase of the services. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/service_discovery_run/actions/start/invoke

Execute a service discovery on a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [DiscoverServices](schemas.md#discoverservices)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 302 | Found: The service discovery background job has been initialized. Redirecting to the 'Wait for service discovery completion' endpoint. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: A service discovery background job is currently running | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/discovery_run/{job_id}

Show the status of a bulk discovery job

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `job_id` | path | string | Yes | The unique identifier of the background job executing the bulk discovery |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DiscoveryBackgroundJobStatusObject](schemas.md#discoverybackgroundjobstatusobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: There is no running background job with this job_id. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/host/{host_name}/actions/discover_services/invoke

Execute a service discovery on a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | A hostname. |

**Request Body:** [DiscoverServicesDeprecated](schemas.md#discoverservicesdeprecated)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 302 | Found: The service discovery background job has been initialized. Redirecting to the 'Wait for service discovery completion' endpoint. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: Host could not be found | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: A service discovery background job is currently running | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/host/{host_name}/actions/update_discovery_phase/invoke

Update the phase of a service

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `host_name` | path | string | Yes | The host of the service which shall be updated. |

**Request Body:** [UpdateDiscoveryPhase](schemas.md#updatediscoveryphase)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: Host could not be found | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/service_discovery/{host_name}

Show the current service discovery result

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `host_name` | path | string | Yes | The host of the service discovery result |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/service_discovery_run/{host_name}

Show the last service discovery background job on a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `host_name` | path | string | Yes | A hostname. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/service_discovery_run/{host_name}/actions/wait-for-completion/invoke

Wait for service discovery completion

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `host_name` | path | string | Yes | A hostname. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: The service discovery has been completed. | - |
| 302 | Found: The service discovery is still running. Redirecting to the 'Wait for completion' endpoint. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: There is no running service discovery | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Service groups

### POST /domain-types/service_group_config/actions/bulk-create/invoke

Bulk create service groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkInputServiceGroup](schemas.md#bulkinputservicegroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ServiceGroupCollection](schemas.md#servicegroupcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/service_group_config/actions/bulk-delete/invoke

Bulk delete service groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkDeleteServiceGroup](schemas.md#bulkdeleteservicegroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### PUT /domain-types/service_group_config/actions/bulk-update/invoke

Bulk update service groups

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [BulkUpdateServiceGroup](schemas.md#bulkupdateservicegroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ServiceGroupCollection](schemas.md#servicegroupcollection) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /domain-types/service_group_config/collections/all

Create a service group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [InputServiceGroup](schemas.md#inputservicegroup)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/service_group_config/collections/all

Show all service groups

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ServiceGroupCollection](schemas.md#servicegroupcollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/service_group_config/{name}

Delete a service group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The identifier name of the group. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |

---

### GET /objects/service_group_config/{name}

Show a service group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | The identifier name of the group. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ServiceGroup](schemas.md#servicegroup) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/service_group_config/{name}

Update a service group

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `name` | path | string | Yes | The identifier name of the group. |

**Request Body:** [UpdateGroup2](schemas.md#updategroup2)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [ServiceGroup](schemas.md#servicegroup) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

## Service status

### GET /objects/host/{host_name}/actions/show_service/invoke

Show the monitored service of a host

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `service_description` | query | string | Yes | The service description of the selected host |
| `host_name` | path | string | Yes | A hostname. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [DomainObject](schemas.md#domainobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

## Site Management

### POST /domain-types/site_connection/collections/all

Create a site connection

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [SiteConnectionRequestCreate](schemas.md#siteconnectionrequestcreate)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [SiteConnectionResponse](schemas.md#siteconnectionresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/site_connection/collections/all

Show all site connections

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [SiteConnectionResponseCollection](schemas.md#siteconnectionresponsecollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/site_connection/{site_id}

Update a site connection

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `site_id` | path | string | Yes | A site ID that exists. |

**Request Body:** [SiteConnectionRequestUpdate](schemas.md#siteconnectionrequestupdate)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [SiteConnectionResponse](schemas.md#siteconnectionresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /objects/site_connection/{site_id}

Show a site connection

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `site_id` | path | string | Yes | A site ID that exists. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [SiteConnectionResponse](schemas.md#siteconnectionresponse) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### POST /objects/site_connection/{site_id}/actions/delete/invoke

Delete a site connection

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `site_id` | path | string | Yes | The site ID. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /objects/site_connection/{site_id}/actions/login/invoke

Login to a remote site

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `site_id` | path | string | Yes | A site ID that exists. |

**Request Body:** [SiteLoginRequest](schemas.md#siteloginrequest)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 401 | Unauthorized: The user is not authorized to do this request. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### POST /objects/site_connection/{site_id}/actions/logout/invoke

Logout from a remote site

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `site_id` | path | string | Yes | A site ID that exists. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Time periods

### POST /domain-types/time_period/collections/all

Create a time period

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateTimePeriod](schemas.md#createtimeperiod)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [TimePeriodResponse](schemas.md#timeperiodresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/time_period/collections/all

Show all time periods

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [TimePeriodResponseCollection](schemas.md#timeperiodresponsecollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/time_period/{name}

Delete a time period

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `name` | path | string | Yes | A name used as an identifier. Can be of arbitrary (sensible) length. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 405 | Method Not Allowed: Method not allowed: This request is only allowed with other HTTP methods | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 409 | Conflict: The request is in conflict with the stored resource. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### GET /objects/time_period/{name}

Show a time period

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `name` | path | string | Yes | A name used as an identifier. Can be of arbitrary (sensible) length. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [TimePeriodResponse](schemas.md#timeperiodresponse) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/time_period/{name}

Update a time period

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `name` | path | string | Yes | A name used as an identifier. Can be of arbitrary (sensible) length. |

**Request Body:** [UpdateTimePeriod](schemas.md#updatetimeperiod)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [TimePeriodResponse](schemas.md#timeperiodresponse) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 405 | Method Not Allowed: Method not allowed: This request is only allowed with other HTTP methods | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

## User Roles

### POST /domain-types/user_role/collections/all

Create/clone a user role

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateUserRole](schemas.md#createuserrole)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserRoleObject](schemas.md#userroleobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/user_role/collections/all

Show all user roles

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserRoleCollection](schemas.md#userrolecollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### GET /objects/user_role/{role_id}

Show a user role

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `role_id` | path | string | Yes | An existing user role. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserRoleObject](schemas.md#userroleobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/user_role/{role_id}

Delete a user role

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `role_id` | path | string | Yes | An existing custom user role that you want to delete. |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/user_role/{role_id}

Edit a user role

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `role_id` | path | string | Yes | An existing user role. |

**Request Body:** [EditUserRole](schemas.md#edituserrole)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserRoleObject](schemas.md#userroleobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

## Users

### POST /domain-types/user_config/collections/all

Create a user

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |

**Request Body:** [CreateUser](schemas.md#createuser)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserObject](schemas.md#userobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |

---

### GET /domain-types/user_config/collections/all

Show all users

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserCollection](schemas.md#usercollection) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### PUT /objects/user_config/{username}

Edit a user

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `If-Match` | header | string | Yes | The value of the, to be modified, object's ETag header. You can get this valu... |
| `Content-Type` | header | string | Yes | A header specifying which type of content is in the request/response body. Th... |
| `username` | path | string | Yes | An unique username for the user |

**Request Body:** [UpdateUser](schemas.md#updateuser)

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserObject](schemas.md#userobject) |
| 400 | Bad Request: Parameter or validation failure. | [ApiError](schemas.md#apierror) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |
| 412 | Precondition Failed: The value of the If-Match header doesn't match the object's ETag. | [ApiError](schemas.md#apierror) |
| 415 | Unsupported Media Type: The submitted content-type is not supported. | [ApiError](schemas.md#apierror) |
| 428 | Precondition Required: The required If-Match header is missing. | [ApiError](schemas.md#apierror) |

---

### GET /objects/user_config/{username}

Show a user

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `username` | path | string | Yes | An unique username for the user |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 200 | OK: The operation was done successfully. | [UserObject](schemas.md#userobject) |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

### DELETE /objects/user_config/{username}

Delete a user

**Parameters:**

| Name | In | Type | Required | Description |
|------|-----|------|----------|-------------|
| `username` | path | string | Yes | An unique username for the user |

**Responses:**

| Code | Description | Schema |
|------|-------------|--------|
| 204 | No Content: Operation done successfully. No further output. | - |
| 403 | Forbidden: Configuration via Setup is disabled. | [ApiError](schemas.md#apierror) |
| 404 | Not Found: The requested object has not been found. | [ApiError](schemas.md#apierror) |
| 406 | Not Acceptable: The requests accept headers can not be satisfied. | [ApiError](schemas.md#apierror) |

---

