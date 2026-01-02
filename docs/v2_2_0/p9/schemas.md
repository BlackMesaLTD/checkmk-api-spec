# Schema Definitions - v2.2.0p9

## Table of Contents

- [AcknowledgeHostGroupProblem](#acknowledgehostgroupproblem)
- [AcknowledgeHostProblem](#acknowledgehostproblem)
- [AcknowledgeHostQueryProblem](#acknowledgehostqueryproblem)
- [AcknowledgeHostRelatedProblem](#acknowledgehostrelatedproblem)
- [AcknowledgeServiceGroupProblem](#acknowledgeservicegroupproblem)
- [AcknowledgeServiceQueryProblem](#acknowledgeservicequeryproblem)
- [AcknowledgeServiceRelatedProblem](#acknowledgeservicerelatedproblem)
- [AcknowledgeSpecificServiceProblem](#acknowledgespecificserviceproblem)
- [ActivateChanges](#activatechanges)
- [ActivationExtensionFields](#activationextensionfields)
- [ActivationRunCollection](#activationruncollection)
- [ActivationRunResponse](#activationrunresponse)
- [AgentControllerCertificateSettings](#agentcontrollercertificatesettings)
- [ApiError](#apierror)
- [AsciiMailPluginCreate](#asciimailplugincreate)
- [AuthOption](#authoption)
- [AuthOption1](#authoption1)
- [AuthPassword](#authpassword)
- [AuthSecret](#authsecret)
- [AuthUpdateOption](#authupdateoption)
- [AuthUpdatePassword](#authupdatepassword)
- [AuthUpdateRemove](#authupdateremove)
- [AuthUpdateSecret](#authupdatesecret)
- [AuxTagAttrsCreate](#auxtagattrscreate)
- [AuxTagAttrsResponse](#auxtagattrsresponse)
- [AuxTagAttrsUpdate](#auxtagattrsupdate)
- [AuxTagResponse](#auxtagresponse)
- [AuxTagResponseCollection](#auxtagresponsecollection)
- [BIAction](#biaction)
- [BIAggregationComputationOptions](#biaggregationcomputationoptions)
- [BIAggregationEndpoint](#biaggregationendpoint)
- [BIAggregationFunction](#biaggregationfunction)
- [BIAggregationFunctionBest](#biaggregationfunctionbest)
- [BIAggregationFunctionCountOK](#biaggregationfunctioncountok)
- [BIAggregationFunctionCountSettings](#biaggregationfunctioncountsettings)
- [BIAggregationFunctionWorst](#biaggregationfunctionworst)
- [BIAggregationGroups](#biaggregationgroups)
- [BIAggregationStateRequest](#biaggregationstaterequest)
- [BIAggregationStateResponse](#biaggregationstateresponse)
- [BIAggregationVisualization](#biaggregationvisualization)
- [BIAllHostsChoice](#biallhostschoice)
- [BICallARuleAction](#bicallaruleaction)
- [BIEmptySearch](#biemptysearch)
- [BIFixedArgumentsSearch](#bifixedargumentssearch)
- [BIFixedArgumentsSearchToken](#bifixedargumentssearchtoken)
- [BIHostAliasRegexChoice](#bihostaliasregexchoice)
- [BIHostChoice](#bihostchoice)
- [BIHostNameRegexChoice](#bihostnameregexchoice)
- [BIHostSearch](#bihostsearch)
- [BINodeGenerator](#binodegenerator)
- [BINodeVisBlockStyle](#binodevisblockstyle)
- [BINodeVisForceStyle](#binodevisforcestyle)
- [BINodeVisHierarchyStyle](#binodevishierarchystyle)
- [BINodeVisHierarchyStyleConfig](#binodevishierarchystyleconfig)
- [BINodeVisLayoutStyle](#binodevislayoutstyle)
- [BINodeVisNoneStyle](#binodevisnonestyle)
- [BINodeVisRadialStyle](#binodevisradialstyle)
- [BINodeVisRadialStyleConfig](#binodevisradialstyleconfig)
- [BIPackEndpoint](#bipackendpoint)
- [BIParams](#biparams)
- [BIRuleComputationOptions](#birulecomputationoptions)
- [BIRuleEndpoint](#biruleendpoint)
- [BIRuleProperties](#biruleproperties)
- [BISearch](#bisearch)
- [BIServiceSearch](#biservicesearch)
- [BIStateOfHostAction](#bistateofhostaction)
- [BIStateOfRemainingServicesAction](#bistateofremainingservicesaction)
- [BIStateOfServiceAction](#bistateofserviceaction)
- [BackgroundJobStatus](#backgroundjobstatus)
- [BaseUserAttributes](#baseuserattributes)
- [BasicSettingsAttributes](#basicsettingsattributes)
- [BasicSettingsAttributesCreate](#basicsettingsattributescreate)
- [BasicSettingsAttributesUpdate](#basicsettingsattributesupdate)
- [BinaryExpr](#binaryexpr)
- [BulkCreateHost](#bulkcreatehost)
- [BulkDeleteContactGroup](#bulkdeletecontactgroup)
- [BulkDeleteHost](#bulkdeletehost)
- [BulkDeleteHostGroup](#bulkdeletehostgroup)
- [BulkDeleteServiceGroup](#bulkdeleteservicegroup)
- [BulkDiscovery](#bulkdiscovery)
- [BulkHostActionWithFailedHosts](#bulkhostactionwithfailedhosts)
- [BulkInputContactGroup](#bulkinputcontactgroup)
- [BulkInputHostGroup](#bulkinputhostgroup)
- [BulkInputServiceGroup](#bulkinputservicegroup)
- [BulkOutsideTimePeriodValue](#bulkoutsidetimeperiodvalue)
- [BulkUpdateContactGroup](#bulkupdatecontactgroup)
- [BulkUpdateFolder](#bulkupdatefolder)
- [BulkUpdateHost](#bulkupdatehost)
- [BulkUpdateHostGroup](#bulkupdatehostgroup)
- [BulkUpdateServiceGroup](#bulkupdateservicegroup)
- [CaseParams](#caseparams)
- [ChangeEventState](#changeeventstate)
- [ChangeEventStateSelector](#changeeventstateselector)
- [ChangeStateWithParams](#changestatewithparams)
- [ChangeStateWithQuery](#changestatewithquery)
- [ChangesFields](#changesfields)
- [Checkbox](#checkbox)
- [CheckboxHostEventType](#checkboxhosteventtype)
- [CheckboxLabel](#checkboxlabel)
- [CheckboxMatchHostTags](#checkboxmatchhosttags)
- [CheckboxOneOf](#checkboxoneof)
- [CheckboxRestrictNotificationNumbers](#checkboxrestrictnotificationnumbers)
- [CheckboxServiceEventType](#checkboxserviceeventtype)
- [CheckboxThrottlePeriodicNotifcations](#checkboxthrottleperiodicnotifcations)
- [CheckboxWithFromToServiceLevels](#checkboxwithfromtoservicelevels)
- [CheckboxWithListOfLabels](#checkboxwithlistoflabels)
- [CheckboxWithListOfServiceGroupsRegex](#checkboxwithlistofservicegroupsregex)
- [CheckboxWithListOfStr](#checkboxwithlistofstr)
- [CheckboxWithStr](#checkboxwithstr)
- [CheckboxWithStrValue](#checkboxwithstrvalue)
- [CheckboxWithSysLogPriority](#checkboxwithsyslogpriority)
- [Child](#child)
- [ChildWith](#childwith)
- [CiscoExplicitWebhookUrl](#ciscoexplicitwebhookurl)
- [CiscoPasswordStore](#ciscopasswordstore)
- [CiscoUrlOrStoreSelector](#ciscourlorstoreselector)
- [CiscoWebexPluginCreate](#ciscowebexplugincreate)
- [ClusterCreateAttribute](#clustercreateattribute)
- [CollectionItem](#collectionitem)
- [CommentAttributes](#commentattributes)
- [CommentCollection](#commentcollection)
- [CommentObject](#commentobject)
- [ConcreteDisabledNotifications](#concretedisablednotifications)
- [ConcreteHostTagGroup](#concretehosttaggroup)
- [ConcreteTimePeriodException](#concretetimeperiodexception)
- [ConcreteTimeRange](#concretetimerange)
- [ConcreteTimeRangeActive](#concretetimerangeactive)
- [ConcreteUserContactOption](#concreteusercontactoption)
- [ConcreteUserInterfaceAttributes](#concreteuserinterfaceattributes)
- [ConditionsAttributes](#conditionsattributes)
- [ConfigurationConnectionAttributes](#configurationconnectionattributes)
- [ConfigurationConnectionAttributes1](#configurationconnectionattributes1)
- [ConnectionMode](#connectionmode)
- [ContactGroup](#contactgroup)
- [ContactGroupCollection](#contactgroupcollection)
- [ContactSelection](#contactselection)
- [ContactSelectionAttributes](#contactselectionattributes)
- [CreateClusterHost](#createclusterhost)
- [CreateFolder](#createfolder)
- [CreateHost](#createhost)
- [CreateHostComment](#createhostcomment)
- [CreateHostDowntime](#createhostdowntime)
- [CreateHostGroupDowntime](#createhostgroupdowntime)
- [CreateHostQueryComment](#createhostquerycomment)
- [CreateHostQueryDowntime](#createhostquerydowntime)
- [CreateHostRelatedComment](#createhostrelatedcomment)
- [CreateHostRelatedDowntime](#createhostrelateddowntime)
- [CreateServiceComment](#createservicecomment)
- [CreateServiceDowntime](#createservicedowntime)
- [CreateServiceGroupDowntime](#createservicegroupdowntime)
- [CreateServiceQueryComment](#createservicequerycomment)
- [CreateServiceQueryDowntime](#createservicequerydowntime)
- [CreateServiceRelatedComment](#createservicerelatedcomment)
- [CreateServiceRelatedDowntime](#createservicerelateddowntime)
- [CreateTimePeriod](#createtimeperiod)
- [CreateUser](#createuser)
- [CreateUserRole](#createuserrole)
- [CustomHostAttributes](#customhostattributes)
- [CustomMacro](#custommacro)
- [CustomTimeRange](#customtimerange)
- [CustomUserAttributes](#customuserattributes)
- [DateTimeRange](#datetimerange)
- [DeleteCommentById](#deletecommentbyid)
- [DeleteComments](#deletecomments)
- [DeleteCommentsByParams](#deletecommentsbyparams)
- [DeleteCommentsByQuery](#deletecommentsbyquery)
- [DeleteDowntime](#deletedowntime)
- [DeleteDowntimeById](#deletedowntimebyid)
- [DeleteDowntimeByName](#deletedowntimebyname)
- [DeleteDowntimeByQuery](#deletedowntimebyquery)
- [DeleteECEvents](#deleteecevents)
- [DirectMapping](#directmapping)
- [DisabledNotifications](#disablednotifications)
- [DiscoverServices](#discoverservices)
- [DiscoverServicesDeprecated](#discoverservicesdeprecated)
- [DiscoveryBackgroundJobStatusObject](#discoverybackgroundjobstatusobject)
- [DomainObject](#domainobject)
- [DomainObjectCollection](#domainobjectcollection)
- [DowntimeAttributes](#downtimeattributes)
- [DowntimeCollection](#downtimecollection)
- [DowntimeObject](#downtimeobject)
- [ECEventAttributes](#eceventattributes)
- [ECEventResponse](#eceventresponse)
- [EditUserRole](#edituserrole)
- [EmailAndDisplayName](#emailanddisplayname)
- [EventConsoleAlertAttrsResponse](#eventconsolealertattrsresponse)
- [EventConsoleAlertsResponse](#eventconsolealertsresponse)
- [EventConsoleResponseCollection](#eventconsoleresponsecollection)
- [Expr](#expr)
- [FailedHosts](#failedhosts)
- [FilterById](#filterbyid)
- [FilterByParams](#filterbyparams)
- [FilterByQuery](#filterbyquery)
- [FilterParams](#filterparams)
- [FilterParamsUpdateAndAcknowledge](#filterparamsupdateandacknowledge)
- [Folder](#folder)
- [FolderCollection](#foldercollection)
- [FolderCreateAttribute](#foldercreateattribute)
- [FolderExtensions](#folderextensions)
- [FolderMembers](#foldermembers)
- [FolderUpdateAttribute](#folderupdateattribute)
- [FolderViewAttribute](#folderviewattribute)
- [FromEmailAndNameCheckbox](#fromemailandnamecheckbox)
- [FromToNotificationNumbers](#fromtonotificationnumbers)
- [FromToServiceLevels](#fromtoservicelevels)
- [Get](#get)
- [GetGraph](#getgraph)
- [GetMetric](#getmetric)
- [GraphCollection](#graphcollection)
- [HTMLMailPluginCreate](#htmlmailplugincreate)
- [Heartbeat](#heartbeat)
- [Heartbeat1](#heartbeat1)
- [Host](#host)
- [HostConditions](#hostconditions)
- [HostConfig](#hostconfig)
- [HostConfigCollection](#hostconfigcollection)
- [HostConfigSchemaInternal](#hostconfigschemainternal)
- [HostContactGroup](#hostcontactgroup)
- [HostCreateAttribute](#hostcreateattribute)
- [HostEventType](#hosteventtype)
- [HostExtensions](#hostextensions)
- [HostExtensionsEffectiveAttributes](#hostextensionseffectiveattributes)
- [HostGroup](#hostgroup)
- [HostGroupCollection](#hostgroupcollection)
- [HostMembers](#hostmembers)
- [HostOrServiceCondition](#hostorservicecondition)
- [HostTag](#hosttag)
- [HostTag1](#hosttag1)
- [HostTagExtensions](#hosttagextensions)
- [HostTagGroupCollection](#hosttaggroupcollection)
- [HostTagValues](#hosttagvalues)
- [HostUpdateAttribute](#hostupdateattribute)
- [HostViewAttribute](#hostviewattribute)
- [IPAddressRange](#ipaddressrange)
- [IPAddresses](#ipaddresses)
- [IPMIParameters](#ipmiparameters)
- [IPNetwork](#ipnetwork)
- [IPRangeWithRegexp](#iprangewithregexp)
- [IPRegexp](#ipregexp)
- [IdleOption](#idleoption)
- [IlertAPIKey](#ilertapikey)
- [IlertKeyOrStoreSelector](#ilertkeyorstoreselector)
- [IlertPasswordStoreID](#ilertpasswordstoreid)
- [IlertPluginCreate](#ilertplugincreate)
- [IncidentParams](#incidentparams)
- [InputContactGroup](#inputcontactgroup)
- [InputHostGroup](#inputhostgroup)
- [InputHostTagGroup](#inputhosttaggroup)
- [InputPassword](#inputpassword)
- [InputRuleObject](#inputruleobject)
- [InputServiceGroup](#inputservicegroup)
- [InstalledVersions](#installedversions)
- [JiraPluginCreate](#jiraplugincreate)
- [JobLogs](#joblogs)
- [LabelCondition](#labelcondition)
- [Link](#link)
- [LinkHostUUID](#linkhostuuid)
- [LockedBy](#lockedby)
- [LogicalExpr](#logicalexpr)
- [MSTeamsExplicitWebhookUrl](#msteamsexplicitwebhookurl)
- [MSTeamsPluginCreate](#msteamsplugincreate)
- [MSTeamsURLResponse](#msteamsurlresponse)
- [MSTeamsUrlOrStoreSelector](#msteamsurlorstoreselector)
- [MatchCustomMacros](#matchcustommacros)
- [MatchEventConsoleAlertsResponse](#matcheventconsolealertsresponse)
- [MetaData](#metadata)
- [Metric](#metric)
- [MgmntTypeCaseParams](#mgmnttypecaseparams)
- [MgmntTypeIncidentParams](#mgmnttypeincidentparams)
- [MgmntTypeSelector](#mgmnttypeselector)
- [MkEventDPluginCreate](#mkeventdplugincreate)
- [MoveFolder](#movefolder)
- [MoveHost](#movehost)
- [MoveRuleTo](#moveruleto)
- [MoveToFolder](#movetofolder)
- [MoveToSpecificRule](#movetospecificrule)
- [NetworkScan](#networkscan)
- [NetworkScanResult](#networkscanresult)
- [NotExpr](#notexpr)
- [NotificationBulking](#notificationbulking)
- [NotificationBulkingCheckbox](#notificationbulkingcheckbox)
- [NotificationBulkingCommonAttributes](#notificationbulkingcommonattributes)
- [NotificationPlugin](#notificationplugin)
- [NotificationRuleAttributes](#notificationruleattributes)
- [NotificationRuleConfig](#notificationruleconfig)
- [NotificationRuleRequest](#notificationrulerequest)
- [NotificationRuleResponse](#notificationruleresponse)
- [NotificationRuleResponseCollection](#notificationruleresponsecollection)
- [ObjectActionMember](#objectactionmember)
- [ObjectCollectionMember](#objectcollectionmember)
- [ObjectProperty](#objectproperty)
- [OpsGenieExplicitKey](#opsgenieexplicitkey)
- [OpsGeniePluginCreate](#opsgenieplugincreate)
- [OpsGenieStoreID](#opsgeniestoreid)
- [OpsGenisStoreOrExplicitKeySelector](#opsgenisstoreorexplicitkeyselector)
- [PagerDutyAPIKeyStoreID](#pagerdutyapikeystoreid)
- [PagerDutyExplicitKey](#pagerdutyexplicitkey)
- [PagerDutyPluginCreate](#pagerdutyplugincreate)
- [PagerDutyStoreOrIntegrationKeySelector](#pagerdutystoreorintegrationkeyselector)
- [Parent](#parent)
- [PasswordCollection](#passwordcollection)
- [PasswordExtension](#passwordextension)
- [PasswordObject](#passwordobject)
- [PendingChangesCollection](#pendingchangescollection)
- [PluginBase](#pluginbase)
- [PluginBase1](#pluginbase1)
- [PluginName](#pluginname)
- [PluginOptionsSelector](#pluginoptionsselector)
- [PluginSelector](#pluginselector)
- [PluginWithParams](#pluginwithparams)
- [ProxyAttributes](#proxyattributes)
- [ProxyAttributes1](#proxyattributes1)
- [ProxyOrDirect](#proxyordirect)
- [ProxyParams](#proxyparams)
- [ProxyParams1](#proxyparams1)
- [ProxyTcp](#proxytcp)
- [ProxyTcp1](#proxytcp1)
- [PushOverPluginCreate](#pushoverplugincreate)
- [ReferTo](#referto)
- [RegexpRewrites](#regexprewrites)
- [RegisterHost](#registerhost)
- [RenameHost](#renamehost)
- [RuleCollection](#rulecollection)
- [RuleConditions](#ruleconditions)
- [RuleConditions1](#ruleconditions1)
- [RuleExtensions](#ruleextensions)
- [RuleNotification](#rulenotification)
- [RuleNotificationMethod](#rulenotificationmethod)
- [RuleObject](#ruleobject)
- [RuleProperties](#ruleproperties)
- [RuleProperties1](#ruleproperties1)
- [RulePropertiesAttributes](#rulepropertiesattributes)
- [RulesetCollection](#rulesetcollection)
- [RulesetExtensions](#rulesetextensions)
- [RulesetObject](#rulesetobject)
- [SMSAPIExplicitPassword](#smsapiexplicitpassword)
- [SMSAPIPStoreID](#smsapipstoreid)
- [SMSAPIPasswordSelector](#smsapipasswordselector)
- [SMSAPIPluginCreate](#smsapiplugincreate)
- [SMSPluginBase](#smspluginbase)
- [SNMPCommunity](#snmpcommunity)
- [SNMPCredentials](#snmpcredentials)
- [SNMPv3AuthNoPrivacy](#snmpv3authnoprivacy)
- [SNMPv3AuthPrivacy](#snmpv3authprivacy)
- [SNMPv3NoAuthNoPrivacy](#snmpv3noauthnoprivacy)
- [ServiceConditions](#serviceconditions)
- [ServiceEventType](#serviceeventtype)
- [ServiceGroup](#servicegroup)
- [ServiceGroupCollection](#servicegroupcollection)
- [ServiceGroupsRegex](#servicegroupsregex)
- [ServiceNowExplicitPassword](#servicenowexplicitpassword)
- [ServiceNowPasswordSelector](#servicenowpasswordselector)
- [ServiceNowPasswordStoreID](#servicenowpasswordstoreid)
- [ServiceNowPluginCreate](#servicenowplugincreate)
- [SignL4ExplicitOrStoreSelector](#signl4explicitorstoreselector)
- [SignL4TeamSecret](#signl4teamsecret)
- [SignL4TeamSecretStoreID](#signl4teamsecretstoreid)
- [Signl4PluginCreate](#signl4plugincreate)
- [SiteConfigAttributes](#siteconfigattributes)
- [SiteConfigAttributesCreate](#siteconfigattributescreate)
- [SiteConfigAttributesUpdate](#siteconfigattributesupdate)
- [SiteConnectionRequestCreate](#siteconnectionrequestcreate)
- [SiteConnectionRequestUpdate](#siteconnectionrequestupdate)
- [SiteConnectionResponse](#siteconnectionresponse)
- [SiteConnectionResponseCollection](#siteconnectionresponsecollection)
- [SiteLoginRequest](#siteloginrequest)
- [SlackPluginCreate](#slackplugincreate)
- [SlackStoreOrExplicitURLSelector](#slackstoreorexpliciturlselector)
- [SlackWebhookStore](#slackwebhookstore)
- [SlackWebhookURL](#slackwebhookurl)
- [SocketAttributes](#socketattributes)
- [SocketAttributes1](#socketattributes1)
- [SocketIP4](#socketip4)
- [SocketIP6](#socketip6)
- [SocketType](#sockettype)
- [SocketUnixAttributes](#socketunixattributes)
- [SpectrumPluginBase](#spectrumpluginbase)
- [SplunkRESTEndpointSelector](#splunkrestendpointselector)
- [SplunkStoreID](#splunkstoreid)
- [SplunkURLExplicit](#splunkurlexplicit)
- [StatusConnectionAttributes](#statusconnectionattributes)
- [StatusConnectionAttributes1](#statusconnectionattributes1)
- [StatusHostAttributes](#statushostattributes)
- [StatusHostAttributesBase](#statushostattributesbase)
- [StatusHostAttributesSet](#statushostattributesset)
- [StatusHostSet](#statushostset)
- [SysLogToFromPriorities](#syslogtofrompriorities)
- [TagCondition](#tagcondition)
- [TagConditionConditionSchemaBase](#tagconditionconditionschemabase)
- [TagConditionScalarSchemaBase](#tagconditionscalarschemabase)
- [TagGroupAttributes](#taggroupattributes)
- [ThrottlePeriodicNotifications](#throttleperiodicnotifications)
- [TimeAllowedRange](#timeallowedrange)
- [TimePeriodAttrsResponse](#timeperiodattrsresponse)
- [TimePeriodException](#timeperiodexception)
- [TimePeriodResponse](#timeperiodresponse)
- [TimePeriodResponseCollection](#timeperiodresponsecollection)
- [TimeRange](#timerange)
- [TimeRange1](#timerange1)
- [TimeRangeActive](#timerangeactive)
- [TranslateNames](#translatenames)
- [UpdateAndAcknowledeEventSiteIDRequired](#updateandacknowledeeventsiteidrequired)
- [UpdateAndAcknowledgeFilter](#updateandacknowledgefilter)
- [UpdateAndAcknowledgeSelector](#updateandacknowledgeselector)
- [UpdateAndAcknowledgeWithParams](#updateandacknowledgewithparams)
- [UpdateAndAcknowledgeWithQuery](#updateandacknowledgewithquery)
- [UpdateContactGroup](#updatecontactgroup)
- [UpdateDiscoveryPhase](#updatediscoveryphase)
- [UpdateFolder](#updatefolder)
- [UpdateFolderEntry](#updatefolderentry)
- [UpdateGroup](#updategroup)
- [UpdateGroup1](#updategroup1)
- [UpdateGroup2](#updategroup2)
- [UpdateHost](#updatehost)
- [UpdateHostEntry](#updatehostentry)
- [UpdateHostGroup](#updatehostgroup)
- [UpdateHostTagGroup](#updatehosttaggroup)
- [UpdateNodes](#updatenodes)
- [UpdatePassword](#updatepassword)
- [UpdateServiceGroup](#updateservicegroup)
- [UpdateTimePeriod](#updatetimeperiod)
- [UpdateUser](#updateuser)
- [UseLiveStatusDaemon](#uselivestatusdaemon)
- [UserCollection](#usercollection)
- [UserContactOption](#usercontactoption)
- [UserIdleOption](#useridleoption)
- [UserInterfaceAttributes](#userinterfaceattributes)
- [UserInterfaceUpdateAttributes](#userinterfaceupdateattributes)
- [UserObject](#userobject)
- [UserRoleAttributes](#userroleattributes)
- [UserRoleCollection](#userrolecollection)
- [UserRoleObject](#userroleobject)
- [UserSyncAttributes](#usersyncattributes)
- [UserSyncAttributes1](#usersyncattributes1)
- [UserSyncBase](#usersyncbase)
- [UserSyncWithLdapConnection](#usersyncwithldapconnection)
- [VictoropsPluginCreate](#victoropsplugincreate)
- [WhenToBulk](#whentobulk)
- [X509PEM](#x509pem)
- [X509ReqPEMUUID](#x509reqpemuuid)

---

## AcknowledgeHostGroupProblem

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledge_type` | string | Yes | The acknowledge host selection type. |
| `comment` | string | Yes | If set, this comment will be stored alongside the acknowledgement. |
| `hostgroup_name` | string | Yes | The name of the host group. |
| `notify` | boolean | No | If set, notifications will be sent out to the configured contacts. Defaults t... |
| `persistent` | boolean | No | If set, the comment will persist a restart. Defaults to False. |
| `sticky` | boolean | No | If set, only a state-change of the host to an UP state will discard the ackno... |

---

## AcknowledgeHostProblem

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledge_type` | string | Yes | The acknowledge host selection type. |
| `comment` | string | Yes | If set, this comment will be stored alongside the acknowledgement. |
| `host_name` | string | Yes | The name of the host. |
| `notify` | boolean | No | If set, notifications will be sent out to the configured contacts. Defaults t... |
| `persistent` | boolean | No | If set, the comment will persist a restart. Defaults to False. |
| `sticky` | boolean | No | If set, only a state-change of the host to an UP state will discard the ackno... |

---

## AcknowledgeHostQueryProblem

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledge_type` | string | Yes | The acknowledge host selection type. |
| `comment` | string | Yes | If set, this comment will be stored alongside the acknowledgement. |
| `notify` | boolean | No | If set, notifications will be sent out to the configured contacts. Defaults t... |
| `persistent` | boolean | No | If set, the comment will persist a restart. Defaults to False. |
| `query` | object | Yes | An query expression of the Livestatus 'hosts' table in nested dictionary form... |
| `sticky` | boolean | No | If set, only a state-change of the host to an UP state will discard the ackno... |

---

## AcknowledgeHostRelatedProblem

---

## AcknowledgeServiceGroupProblem

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledge_type` | string | Yes | The acknowledge service selection type. |
| `comment` | string | Yes | If set, this comment will be stored alongside the acknowledgement. |
| `notify` | boolean | No | If set, notifications will be sent out to the configured contacts. Defaults t... |
| `persistent` | boolean | No | If set, the comment will persist a restart. Defaults to False. |
| `servicegroup_name` | string | Yes | The name of the service group. Any host having a service in this group will b... |
| `sticky` | boolean | No | If set, only a state-change of the service to an OK state will discard the ac... |

---

## AcknowledgeServiceQueryProblem

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledge_type` | string | Yes | The acknowledge service selection type. |
| `comment` | string | Yes | If set, this comment will be stored alongside the acknowledgement. |
| `notify` | boolean | No | If set, notifications will be sent out to the configured contacts. Defaults t... |
| `persistent` | boolean | No | If set, the comment will persist a restart. Defaults to False. |
| `query` | object | Yes | An query expression of the Livestatus 'services' table in nested dictionary f... |
| `sticky` | boolean | No | If set, only a state-change of the service to an OK state will discard the ac... |

---

## AcknowledgeServiceRelatedProblem

---

## AcknowledgeSpecificServiceProblem

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledge_type` | string | Yes | The acknowledge service selection type. |
| `comment` | string | Yes | If set, this comment will be stored alongside the acknowledgement. |
| `host_name` | string | Yes |  |
| `notify` | boolean | No | If set, notifications will be sent out to the configured contacts. Defaults t... |
| `persistent` | boolean | No | If set, the comment will persist a restart. Defaults to False. |
| `service_description` | string | Yes | The acknowledgement process will be applied to all matching service descriptions |
| `sticky` | boolean | No | If set, only a state-change of the service to an OK state will discard the ac... |

---

## ActivateChanges

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `force_foreign_changes` | boolean | No | Will activate changes even if the user who made those changes is not the curr... |
| `redirect` | boolean | No | After starting the activation, redirect immediately to the 'Wait for completi... |
| `sites` | []string | No | The names of the sites on which the configuration shall be activated. An empt... |

---

## ActivationExtensionFields

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `changes` | []ChangesFields | No | The changes in this activation |
| `force_foreign_changes` | boolean | No | If the activation is still running |
| `is_running` | boolean | No | If the activation is still running |
| `sites` | []string | No | Sites affected by this activation |
| `time_started` | string | No | The date and time the activation was started. |

---

## ActivationRunCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `value` | []ActivationRunResponse | No | A list of activation runs. |

---

## ActivationRunResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The activation run attributes. |
| `id` | string | No | The unique identifier for this activation run. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | The activation run status. |

---

## AgentControllerCertificateSettings

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `lifetime_in_months` | integer | Yes | Lifetime of agent controller certificates in months |

---

## ApiError

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `detail` | string | Yes | Detailed information on what exactly went wrong. |
| `ext` | object | No | Additional information about the error. |
| `fields` | object | No | Detailed error messages on all fields failing validation. |
| `status` | integer | Yes | The HTTP status code. |
| `title` | string | Yes | A summary of the problem. |

---

## AsciiMailPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `body_head_for_both_host_and_service_notifications` | CheckboxOneOf | Yes |  |
| `body_tail_for_host_notifications` | CheckboxOneOf | Yes |  |
| `body_tail_for_service_notifications` | CheckboxOneOf | Yes |  |
| `from_details` | CheckboxOneOf | Yes |  |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `reply_to` | CheckboxOneOf | Yes |  |
| `send_separate_notification_to_every_recipient` | Checkbox | Yes |  |
| `sort_order_for_bulk_notificaions` | CheckboxOneOf | Yes |  |
| `subject_for_host_notifications` | CheckboxOneOf | Yes |  |
| `subject_for_service_notifications` | CheckboxOneOf | Yes |  |

---

## AuthOption

---

## AuthOption1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_type` | string | No |  |
| `enforce_password_change` | boolean | No | If set to True, the user will be forced to change his password on the next lo... |

---

## AuthPassword

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_type` | string | No | The authentication type |
| `enforce_password_change` | boolean | No | If set to True, the user will be forced to change his password on the next lo... |
| `password` | string | No | The password for login |

---

## AuthSecret

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_type` | string | No | The authentication type |
| `secret` | string | No | For accounts used by automation processes (such as fetching data from views f... |

---

## AuthUpdateOption

---

## AuthUpdatePassword

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_type` | string | Yes | The authentication type |
| `enforce_password_change` | boolean | No | If set to True, the user will be forced to change his password on the next lo... |
| `password` | string | No | The password for login |

---

## AuthUpdateRemove

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_type` | string | Yes | The authentication type |

---

## AuthUpdateSecret

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_type` | string | Yes | The authentication type |
| `secret` | string | No | For accounts used by automation processes (such as fetching data from views f... |

---

## AuxTagAttrsCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `aux_tag_id` | string | Yes | An auxiliary tag id |
| `help` | string | No | The help of the Auxiliary tag |
| `title` | string | Yes | The title of the Auxiliary tag |
| `topic` | string | Yes | Different tags can be grouped in topics to make the visualization and selecti... |

---

## AuxTagAttrsResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `help` | string | Yes | The help of the Auxiliary tag |
| `topic` | string | Yes | Different tags can be grouped in topics to make the visualization and selecti... |

---

## AuxTagAttrsUpdate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `help` | string | No | The help of the Auxiliary tag |
| `title` | string | No | The title of the Auxiliary tag |
| `topic` | string | No | Different tags can be grouped in topics to make the visualization and selecti... |

---

## AuxTagResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The Auxiliary Tag attributes. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## AuxTagResponseCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []AuxTagResponse | No | A list of site configuration objects. |

---

## BIAction

---

## BIAggregationComputationOptions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disabled` | boolean | Yes |  |
| `escalate_downtimes_as_warn` | boolean | Yes |  |
| `freeze_aggregations` | boolean | No |  |
| `use_hard_states` | boolean | Yes |  |

---

## BIAggregationEndpoint

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `aggregation_visualization` | object | Yes | Nested dictionary |
| `comment` | string | No | An optional comment that may be used to explain the purpose of this object. |
| `computation_options` | object | Yes | Nested dictionary |
| `customer` | string | No | CME Edition only: The customer id for this aggregation. |
| `groups` | object | Yes | Nested dictionary |
| `id` | string | Yes | The unique aggregation id |
| `node` | object | Yes | Nested dictionary |
| `pack_id` | string | Yes | The identifier of the BI pack. |

---

## BIAggregationFunction

---

## BIAggregationFunctionBest

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `count` | integer | Yes |  |
| `restrict_state` | integer | Yes |  |
| `type` | object | Yes |  |

---

## BIAggregationFunctionCountOK

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `levels_ok` | BIAggregationFunctionCountSettings | Yes |  |
| `levels_warn` | BIAggregationFunctionCountSettings | Yes |  |
| `type` | object | Yes |  |

---

## BIAggregationFunctionCountSettings

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | string | Yes |  |
| `value` | integer | Yes |  |

---

## BIAggregationFunctionWorst

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `count` | integer | Yes |  |
| `restrict_state` | integer | Yes |  |
| `type` | object | Yes |  |

---

## BIAggregationGroups

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `names` | []string | No |  |
| `paths` | [][]string | No |  |

---

## BIAggregationStateRequest

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `filter_groups` | []string | No | Filter by group |
| `filter_names` | []string | No | Filter by names |

---

## BIAggregationStateResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `aggregations` | object | No | The Aggregation state |
| `missing_aggr` | []string | No | the missing aggregations |
| `missing_sites` | []string | No | The missing sites |

---

## BIAggregationVisualization

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ignore_rule_styles` | boolean | Yes |  |
| `layout_id` | string | Yes |  |
| `line_style` | string | Yes |  |

---

## BIAllHostsChoice

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | object | Yes |  |

---

## BICallARuleAction

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `params` | object | Yes |  |
| `rule_id` | string | Yes |  |
| `type` | object | Yes |  |

---

## BIEmptySearch

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | object | Yes |  |

---

## BIFixedArgumentsSearch

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `arguments` | []BIFixedArgumentsSearchToken | Yes |  |
| `type` | object | Yes |  |

---

## BIFixedArgumentsSearchToken

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | Yes |  |
| `values` | []string | Yes |  |

---

## BIHostAliasRegexChoice

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `pattern` | string | Yes |  |
| `type` | object | Yes |  |

---

## BIHostChoice

---

## BIHostNameRegexChoice

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `pattern` | string | Yes |  |
| `type` | object | Yes |  |

---

## BIHostSearch

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | HostConditions | Yes |  |
| `refer_to` | ReferTo | Yes |  |
| `type` | object | Yes |  |

---

## BINodeGenerator

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action` | object | Yes | Nested dictionary |
| `search` | object | Yes | Nested dictionary |

---

## BINodeVisBlockStyle

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `style_config` | object | Yes |  |
| `type` | object | Yes |  |

---

## BINodeVisForceStyle

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `style_config` | object | Yes |  |
| `type` | object | Yes |  |

---

## BINodeVisHierarchyStyle

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `style_config` | BINodeVisHierarchyStyleConfig | Yes |  |
| `type` | object | Yes |  |

---

## BINodeVisHierarchyStyleConfig

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `layer_height` | integer | Yes |  |
| `node_size` | integer | Yes |  |
| `rotation` | integer | Yes |  |

---

## BINodeVisLayoutStyle

---

## BINodeVisNoneStyle

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `style_config` | object | Yes |  |
| `type` | object | Yes |  |

---

## BINodeVisRadialStyle

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `style_config` | BINodeVisRadialStyleConfig | Yes |  |
| `type` | object | Yes |  |

---

## BINodeVisRadialStyleConfig

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `degree` | integer | Yes |  |
| `radius` | integer | Yes |  |
| `rotation` | integer | Yes |  |

---

## BIPackEndpoint

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `contact_groups` | []string | Yes | A list of contact group identifiers. |
| `public` | boolean | Yes | Should the BI pack be public or not. |
| `title` | string | Yes | The title of the BI pack. |

---

## BIParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `arguments` | []string | Yes |  |

---

## BIRuleComputationOptions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disabled` | boolean | Yes |  |

---

## BIRuleEndpoint

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `aggregation_function` | object | Yes | Nested dictionary |
| `computation_options` | object | Yes | Nested dictionary |
| `id` | string | Yes | The unique rule id |
| `node_visualization` | object | Yes | Nested dictionary |
| `nodes` | []BINodeGenerator | Yes | A list of nodes for for this rule |
| `pack_id` | string | Yes | The identifier of the BI pack. |
| `params` | object | Yes | Nested dictionary |
| `properties` | object | Yes | Nested dictionary |

---

## BIRuleProperties

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | Yes |  |
| `docu_url` | string | Yes |  |
| `icon` | string | Yes |  |
| `state_messages` | object | Yes |  |
| `title` | string | Yes |  |

---

## BISearch

---

## BIServiceSearch

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | ServiceConditions | Yes |  |
| `type` | object | Yes |  |

---

## BIStateOfHostAction

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_regex` | string | Yes |  |
| `type` | object | Yes |  |

---

## BIStateOfRemainingServicesAction

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_regex` | string | Yes |  |
| `type` | object | Yes |  |

---

## BIStateOfServiceAction

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_regex` | string | Yes |  |
| `service_regex` | string | Yes |  |
| `type` | object | Yes |  |

---

## BackgroundJobStatus

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `active` | boolean | Yes | This field indicates if the background job is active or not. |
| `logs` | object | Yes | Logs related to the background job. |
| `state` | string | Yes | This field indicates the current state of the background job. |

---

## BaseUserAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_option` | object | No | Enforce password change attribute for the user |
| `authorized_sites` | []string | No | The names of the sites that this user is authorized to handle |
| `contact_options` | object | No | Contact settings for the user |
| `contactgroups` | []string | No | The contact groups that this user is a member of |
| `disable_login` | boolean | No | This field indicates if the user is allowed to login to the monitoring. |
| `disable_notifications` | ConcreteDisabledNotifications | No |  |
| `fullname` | string | Yes | The alias or full name of the user. |
| `idle_timeout` | object | No | Idle timeout for the user. Per default, the global configuration is used. |
| `interface_options` | ConcreteUserInterfaceAttributes | No |  |
| `language` | string | No | The language used by the user in the user interface |
| `pager_address` | string | No |  |
| `roles` | []string | No | The list of assigned roles to the user |
| `temperature_unit` | string | No | The temperature unit used for graphs and perfometers. |

---

## BasicSettingsAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The alias of the site. |
| `site_id` | string | Yes | The site id. |

---

## BasicSettingsAttributesCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The alias of the site. |
| `site_id` | string | Yes | The site ID. |

---

## BasicSettingsAttributesUpdate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The alias of the site. |
| `site_id` | string | Yes | The site ID. |

---

## BinaryExpr

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `left` | string | No | The LiveStatus column name. |
| `op` | string | No | The operator. |
| `right` | string | No | The value to compare the column to. |

---

## BulkCreateHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []CreateHost | Yes | A list of host entries. |

---

## BulkDeleteContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []string | Yes | A list of contract group names. |

---

## BulkDeleteHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []string | Yes | A list of host names. |

---

## BulkDeleteHostGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []string | Yes | A list of host group names. |

---

## BulkDeleteServiceGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []string | Yes | A list of service group names. |

---

## BulkDiscovery

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `bulk_size` | integer | No | The number of hosts to be handled at once. |
| `do_full_scan` | boolean | No | The option whether to perform a full scan or not. |
| `hostnames` | []string | Yes | A list of host names |
| `ignore_errors` | boolean | No | The option whether to ignore errors in single check plugins. |
| `mode` | string | No | The mode of the discovery action. The 'refresh' mode starts a new service dis... |

---

## BulkHostActionWithFailedHosts

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `detail` | string | Yes | Detailed information on what exactly went wrong. |
| `ext` | object | No | Details for which hosts have failed |
| `fields` | object | No | Detailed error messages on all fields failing validation. |
| `status` | integer | Yes | The HTTP status code. |
| `title` | string | Yes | A summary of the problem. |

---

## BulkInputContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []InputContactGroup | Yes | A collection of contact group entries. |

---

## BulkInputHostGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []InputHostGroup | Yes | A list of host group entries. |

---

## BulkInputServiceGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []InputServiceGroup | Yes | A list of service group entries. |

---

## BulkOutsideTimePeriodValue

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | NotificationBulkingCommonAttributes | No |  |

---

## BulkUpdateContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []UpdateContactGroup | Yes | A list of contact group entries. |

---

## BulkUpdateFolder

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []UpdateFolderEntry | Yes | A list of folder entries. |

---

## BulkUpdateHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []UpdateHostEntry | Yes | A list of host entries. |

---

## BulkUpdateHostGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []UpdateHostGroup | Yes | A list of host group entries. |

---

## BulkUpdateServiceGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `entries` | []UpdateServiceGroup | Yes | A list of service group entries. |

---

## CaseParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_description` | CheckboxOneOf | No |  |
| `host_short_description` | CheckboxOneOf | No |  |
| `priority` | CheckboxOneOf | No |  |
| `service_description` | CheckboxOneOf | No |  |
| `service_short_description` | CheckboxOneOf | No |  |
| `state_recovery` | CheckboxOneOf | No |  |

---

## ChangeEventState

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `new_state` | string | Yes | The state |
| `site_id` | string | Yes | An existing site id |

---

## ChangeEventStateSelector

---

## ChangeStateWithParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `filter_type` | string | Yes | The way you would like to filter events. |
| `filters` | FilterParams | Yes |  |
| `new_state` | string | Yes | The state |
| `site_id` | string | No | An existing site id |

---

## ChangeStateWithQuery

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `filter_type` | string | Yes | The way you would like to filter events. |
| `new_state` | string | Yes | The state |
| `query` | object | Yes | An query expression of the Livestatus 'eventconsoleevents' table in nested di... |
| `site_id` | string | No | An existing site id |

---

## ChangesFields

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `action_name` | string | No | The action carried out |
| `id` | string | No | The change identifier |
| `text` | string | No |  |
| `time` | string | No | The date and time the change was made. |
| `user_id` | string | No | The user who made the change |

---

## Checkbox

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |

---

## CheckboxHostEventType

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | object | Yes | Select the host event types and transitions this rule should handle. Note: If... |

---

## CheckboxLabel

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | Yes |  |
| `value` | string | Yes |  |

---

## CheckboxMatchHostTags

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | object | Yes | Match host tags with the following parameters |

---

## CheckboxOneOf

---

## CheckboxRestrictNotificationNumbers

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | FromToNotificationNumbers | Yes |  |

---

## CheckboxServiceEventType

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | object | Yes | Select the service event types and transitions this rule should handle. Note:... |

---

## CheckboxThrottlePeriodicNotifcations

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | ThrottlePeriodicNotifications | Yes |  |

---

## CheckboxWithFromToServiceLevels

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | object | Yes | Host or service must be in the following service level to get notification |

---

## CheckboxWithListOfLabels

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | []CheckboxLabel | Yes | A list of key, value label pairs |

---

## CheckboxWithListOfServiceGroupsRegex

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | object | Yes | The service group alias must not match one of the following regular expressio... |

---

## CheckboxWithListOfStr

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | []string | Yes |  |

---

## CheckboxWithStr

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | string | Yes |  |

---

## CheckboxWithStrValue

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | string | Yes |  |

---

## CheckboxWithSysLogPriority

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | SysLogToFromPriorities | No |  |

---

## Child

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | object | Yes |  |

---

## ChildWith

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | HostConditions | Yes |  |
| `host_choice` | object | Yes |  |

---

## CiscoExplicitWebhookUrl

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `url` | string | Yes |  |

---

## CiscoPasswordStore

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## CiscoUrlOrStoreSelector

---

## CiscoWebexPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |
| `webhook_url` | CiscoUrlOrStoreSelector | No |  |

---

## ClusterCreateAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `additional_ipv4addresses` | []string | No | A list of IPv4 addresses. |
| `additional_ipv6addresses` | []string | No | A list of IPv6 addresses. |
| `alias` | string | No | Add a comment or describe this host |
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `inventory_failed` | boolean | No | Whether or not the last bulk discovery failed. It is set to True once it fail... |
| `ipaddress` | string | No | An IPv4 address. |
| `ipv6address` | string | No | An IPv6 address. |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `locked_attributes` | []string | No | Name of host attributes which are locked in the UI. |
| `locked_by` | object | No | Identity of the entity which locked the locked_attributes. The identity is bu... |
| `management_address` | string | No | Address (IPv4, IPv6 or hostname) under which the management board can be reac... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No | Choices: * `"ip-v4-only"`: IPv4 only * `"ip-v6-only"`: IPv6 only * `"ip-v4v6"... |
| `tag_agent` | string | No | Choices: * `"cmk-agent"`: API integrations if configured, else Checkmk agent ... |
| `tag_criticality` | string | No | Choices: * `"prod"`: Productive system * `"critical"`: Business critical * `"... |
| `tag_networking` | string | No | Choices: * `"lan"`: Local network (low latency) * `"wan"`: WAN (high latency)... |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No | Choices: * `"no-snmp"`: No SNMP * `"snmp-v2"`: SNMP v2 or v3 * `"snmp-v1"`: S... |

---

## CollectionItem

---

## CommentAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `author` | string | Yes | The author of the comment |
| `comment` | string | Yes | The comment itself |
| `entry_time` | string | Yes | The timestamp from when the comment was created. |
| `host_name` | string | Yes | The host name. |
| `id` | integer | Yes | The comment ID |
| `is_service` | boolean | Yes | True if the comment is from a service or else it's False. |
| `persistent` | boolean | Yes | If true, the comment will be persisted |
| `service_description` | string | No | The service description the comment belongs to. |
| `site_id` | string | Yes | The site id of the downtime. |

---

## CommentCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []CommentObject | No | A list of comment objects. |

---

## CommentObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The attributes of a service/host comment. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## ConcreteDisabledNotifications

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable` | boolean | No | Option if all notifications should be temporarily disabled |
| `timerange` | object | No | A custom timerange during which notifications are disabled |

---

## ConcreteHostTagGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | Yes | The domain type of the object. |
| `extensions` | object | No | Additional fields for objects of this type. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## ConcreteTimePeriodException

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `date` | string | No | The date of the time period exception.8601 profile |
| `time_ranges` | []ConcreteTimeRange | No |  |

---

## ConcreteTimeRange

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end` | string | No | The hour of the time period. |
| `start` | string | No | The hour of the time period. |

---

## ConcreteTimeRangeActive

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `day` | string | No | The day for which the time ranges are specified |
| `time_ranges` | []ConcreteTimeRange | No |  |

---

## ConcreteUserContactOption

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `email` | string | Yes | The mail address of the user. |
| `fallback_contact` | boolean | No | In case none of the notification rules handle a certain event a notification ... |

---

## ConcreteUserInterfaceAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `interface_theme` | string | No | The theme of the interface |
| `mega_menu_icons` | string | No | This option decides if colored icon should be shown foe every entry in the me... |
| `navigation_bar_icons` | string | No | This option decides if icons in the navigation bar should show/hide the respe... |
| `show_mode` | string | No | This option decides what show mode should be used for unvisited menus. Altern... |
| `sidebar_position` | string | No | The position of the sidebar |

---

## ConditionsAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `event_console_alerts` | MatchEventConsoleAlertsResponse | No |  |
| `match_check_types` | CheckboxWithListOfStr | No |  |
| `match_contact_groups` | CheckboxWithListOfStr | No |  |
| `match_exclude_hosts` | CheckboxWithListOfStr | No |  |
| `match_exclude_service_groups` | CheckboxWithListOfStr | No |  |
| `match_exclude_service_groups_regex` | CheckboxWithListOfServiceGroupsRegex | No |  |
| `match_exclude_services` | CheckboxWithListOfStr | No |  |
| `match_folder` | CheckboxWithStr | No |  |
| `match_host_event_type` | CheckboxHostEventType | No |  |
| `match_host_groups` | CheckboxWithListOfStr | No |  |
| `match_host_labels` | CheckboxWithListOfLabels | No |  |
| `match_host_tags` | CheckboxMatchHostTags | No |  |
| `match_hosts` | CheckboxWithListOfStr | No |  |
| `match_notification_comment` | CheckboxWithStr | No |  |
| `match_only_during_time_period` | CheckboxWithStr | No |  |
| `match_plugin_output` | CheckboxWithStr | No |  |
| `match_service_event_type` | CheckboxServiceEventType | No |  |
| `match_service_groups` | CheckboxWithListOfStr | No |  |
| `match_service_groups_regex` | CheckboxWithListOfServiceGroupsRegex | No |  |
| `match_service_labels` | CheckboxWithListOfLabels | No |  |
| `match_service_levels` | CheckboxWithFromToServiceLevels | No |  |
| `match_services` | CheckboxWithListOfStr | No |  |
| `match_sites` | CheckboxWithListOfStr | No |  |
| `restrict_to_notification_numbers` | CheckboxRestrictNotificationNumbers | No |  |
| `throttle_periodic_notifications` | CheckboxThrottlePeriodicNotifcations | No |  |

---

## ConfigurationConnectionAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `direct_login_to_web_gui_allowed` | boolean | Yes | When enabled, this site is marked for synchronisation every time a Web GUI re... |
| `disable_remote_configuration` | boolean | Yes | It is a good idea to disable access to Setup completely on the remote site. O... |
| `enable_replication` | boolean | Yes | Replication allows you to manage several monitoring sites with a logically ce... |
| `ignore_tls_errors` | boolean | Yes | This might be needed to make the synchronization accept problems with SSL cer... |
| `replicate_event_console` | boolean | Yes | This option enables the distribution of global settings and rules of the Even... |
| `replicate_extensions` | boolean | Yes | If you enable the replication of MKPs then during each Activate Changes MKPs ... |
| `url_of_remote_site` | string | Yes | URL of the remote Checkmk including /check_mk/. This URL is in many cases the... |
| `user_sync` | object | Yes | By default the users are synchronized automatically in the interval configure... |

---

## ConfigurationConnectionAttributes1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `direct_login_to_web_gui_allowed` | boolean | No | When enabled, this site is marked for synchronisation every time a Web GUI re... |
| `disable_remote_configuration` | boolean | No | It is a good idea to disable access to Setup completely on the remote site. O... |
| `enable_replication` | boolean | No | Replication allows you to manage several monitoring sites with a logically ce... |
| `ignore_tls_errors` | boolean | No | This might be needed to make the synchronization accept problems with SSL cer... |
| `replicate_event_console` | boolean | No | This option enables the distribution of global settings and rules of the Even... |
| `replicate_extensions` | boolean | No | If you enable the replication of MKPs then during each Activate Changes MKPs ... |
| `url_of_remote_site` | string | No | URL of the remote Checkmk including /check_mk/. This URL is in many cases the... |
| `user_sync` | object | Yes | By default the users are synchronized automatically in the interval configure... |

---

## ConnectionMode

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `connection_mode` | string | No | This configures the communication direction of this host. * `pull-agent` (def... |

---

## ContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | All the attributes of the domain object. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## ContactGroupCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []ContactGroup | No | A list of contact group objects. |

---

## ContactSelection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `all_contacts_of_the_notified_object` | Checkbox | Yes |  |
| `all_users` | Checkbox | Yes |  |
| `all_users_with_an_email_address` | Checkbox | Yes |  |
| `explicit_email_addresses` | CheckboxOneOf | Yes |  |
| `members_of_contact_groups` | CheckboxOneOf | Yes |  |
| `restrict_by_contact_groups` | CheckboxOneOf | Yes |  |
| `restrict_by_custom_macros` | CheckboxOneOf | Yes |  |
| `the_following_users` | CheckboxOneOf | Yes |  |

---

## ContactSelectionAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `all_contacts_of_the_notified_object` | Checkbox | No |  |
| `all_users` | Checkbox | No |  |
| `all_users_with_an_email_address` | Checkbox | No |  |
| `explicit_email_addresses` | CheckboxWithListOfStr | No |  |
| `members_of_contact_groups` | CheckboxWithListOfStr | No |  |
| `restrict_by_contact_groups` | CheckboxWithListOfStr | No |  |
| `restrict_by_custom_macros` | MatchCustomMacros | No |  |
| `the_following_users` | CheckboxWithListOfStr | No |  |

---

## CreateClusterHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Attributes to set on the newly created host. |
| `folder` | string | Yes | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `host_name` | string | Yes | The hostname of the cluster host. |
| `nodes` | []string | Yes | Nodes where the newly created host should be the cluster-container of. |

---

## CreateFolder

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Specific attributes to apply for all hosts in this folder (among other things). |
| `name` | string | No | The filesystem directory name (not path!) of the folder. No slashes are allowed. |
| `parent` | string | Yes | The folder in which the new folder shall be placed in. The root-folder is spe... |
| `title` | string | Yes | The folder title as displayed in the user interface. |

---

## CreateHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Attributes to set on the newly created host. |
| `folder` | string | Yes | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `host_name` | string | Yes | The hostname or IP address of the host to be created. |

---

## CreateHostComment

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | Yes | The comment which will be stored for the host. |
| `comment_type` | string | Yes | How you would like to leave a comment. |
| `host_name` | string | Yes | The host name |
| `persistent` | boolean | No | If set, the comment will persist a restart. |

---

## CreateHostDowntime

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No |  |
| `downtime_type` | string | Yes | The type of downtime to create. |
| `duration` | integer | No | Duration in minutes. When set, the downtime does not begin automatically at a... |
| `end_time` | string | Yes | The end datetime of the new downtime. The format has to conform to the ISO 86... |
| `host_name` | string | Yes | The hostname or IP address itself. |
| `recur` | string | No | The recurring mode of the new downtime. Available modes are: * fixed * hour *... |
| `start_time` | string | Yes | The start datetime of the new downtime. The format has to conform to the ISO ... |

---

## CreateHostGroupDowntime

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No |  |
| `downtime_type` | string | Yes | The type of downtime to create. |
| `duration` | integer | No | Duration in minutes. When set, the downtime does not begin automatically at a... |
| `end_time` | string | Yes | The end datetime of the new downtime. The format has to conform to the ISO 86... |
| `hostgroup_name` | string | Yes | The name of the host group. A downtime will be scheduled for all hosts in thi... |
| `recur` | string | No | The recurring mode of the new downtime. Available modes are: * fixed * hour *... |
| `start_time` | string | Yes | The start datetime of the new downtime. The format has to conform to the ISO ... |

---

## CreateHostQueryComment

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | Yes | The comment which will be stored for the host. |
| `comment_type` | string | Yes | How you would like to leave a comment. |
| `persistent` | boolean | No | If set, the comment will persist a restart. |
| `query` | object | No | An query expression of the Livestatus 'hosts' table in nested dictionary form... |

---

## CreateHostQueryDowntime

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No |  |
| `downtime_type` | string | Yes | The type of downtime to create. |
| `duration` | integer | No | Duration in minutes. When set, the downtime does not begin automatically at a... |
| `end_time` | string | Yes | The end datetime of the new downtime. The format has to conform to the ISO 86... |
| `query` | object | Yes | An query expression of the Livestatus 'hosts' table in nested dictionary form... |
| `recur` | string | No | The recurring mode of the new downtime. Available modes are: * fixed * hour *... |
| `start_time` | string | Yes | The start datetime of the new downtime. The format has to conform to the ISO ... |

---

## CreateHostRelatedComment

---

## CreateHostRelatedDowntime

---

## CreateServiceComment

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | Yes | The comment which will be stored for the host. |
| `comment_type` | string | Yes | How you would like to leave a comment. |
| `host_name` | string | Yes | The host name |
| `persistent` | boolean | No | If set, the comment will persist a restart. |
| `service_description` | string | Yes | The service description for which the comment is for. No exception is raised ... |

---

## CreateServiceDowntime

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No |  |
| `downtime_type` | string | Yes | The type of downtime to create. |
| `duration` | integer | No | Duration in minutes. When set, the downtime does not begin automatically at a... |
| `end_time` | string | Yes | The end datetime of the new downtime. The format has to conform to the ISO 86... |
| `host_name` | string | Yes |  |
| `recur` | string | No | The recurring mode of the new downtime. Available modes are: * fixed * hour *... |
| `service_descriptions` | []string | Yes | The service description of the service, whose problems shall be acknowledged. |
| `start_time` | string | Yes | The start datetime of the new downtime. The format has to conform to the ISO ... |

---

## CreateServiceGroupDowntime

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No |  |
| `downtime_type` | string | Yes | The type of downtime to create. |
| `duration` | integer | No | Duration in minutes. When set, the downtime does not begin automatically at a... |
| `end_time` | string | Yes | The end datetime of the new downtime. The format has to conform to the ISO 86... |
| `recur` | string | No | The recurring mode of the new downtime. Available modes are: * fixed * hour *... |
| `servicegroup_name` | string | Yes | The name of the service group. Any host having a service in this group will b... |
| `start_time` | string | Yes | The start datetime of the new downtime. The format has to conform to the ISO ... |

---

## CreateServiceQueryComment

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | Yes | The comment which will be stored for the host. |
| `comment_type` | string | Yes | How you would like to leave a comment. |
| `persistent` | boolean | No | If set, the comment will persist a restart. |
| `query` | object | Yes | An query expression of the Livestatus 'services' table in nested dictionary f... |

---

## CreateServiceQueryDowntime

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No |  |
| `downtime_type` | string | Yes | The type of downtime to create. |
| `duration` | integer | No | Duration in minutes. When set, the downtime does not begin automatically at a... |
| `end_time` | string | Yes | The end datetime of the new downtime. The format has to conform to the ISO 86... |
| `query` | object | Yes | An query expression of the Livestatus 'services' table in nested dictionary f... |
| `recur` | string | No | The recurring mode of the new downtime. Available modes are: * fixed * hour *... |
| `start_time` | string | Yes | The start datetime of the new downtime. The format has to conform to the ISO ... |

---

## CreateServiceRelatedComment

---

## CreateServiceRelatedDowntime

---

## CreateTimePeriod

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `active_time_ranges` | []TimeRangeActive | Yes | The list of active time ranges. |
| `alias` | string | Yes | An alias for the time period. |
| `exceptions` | []TimePeriodException | No | A list of additional time ranges to be added. |
| `exclude` | []string | No | A list of time period aliases whose periods are excluded. |
| `name` | string | Yes | A unique name for the time period. |

---

## CreateUser

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_option` | object | No | Authentication option for the user |
| `authorized_sites` | []string | No | The names of the sites the user is authorized to handle. |
| `contact_options` | object | No | Contact settings for the user |
| `contactgroups` | []string | No | Assign the user to one or multiple contact groups. If no contact group is spe... |
| `disable_login` | boolean | No | The user can be blocked from login but will remain part of the site. The disa... |
| `disable_notifications` | object | No |  |
| `fullname` | string | Yes | The alias or full name of the user |
| `idle_timeout` | object | No | Idle timeout for the user. Per default, the global configuration is used. |
| `interface_options` | object | No |  |
| `language` | string | No | Configure the language to be used by the user in the user interface. Omitting... |
| `pager_address` | string | No |  |
| `roles` | []string | No | The list of assigned roles to the user |
| `temperature_unit` | string | No | Configure the temperature unit used for graphs and perfometers. Omitting this... |
| `username` | string | Yes | An unique username for the user |

---

## CreateUserRole

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `new_alias` | string | No | A new alias that you want to give to the newly created user role. |
| `new_role_id` | string | No | The new role id for the newly created user role. |
| `role_id` | string | Yes | Existing userrole that you want to clone. |

---

## CustomHostAttributes

---

## CustomMacro

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `macro_name` | string | Yes | The name of the macro |
| `match_regex` | string | Yes | The text entered here is handled as a regular expression pattern |

---

## CustomTimeRange

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end_time` | string | Yes | The end datetime of the time period. The format has to conform to the ISO 860... |
| `start_time` | string | Yes | The start datetime of the time period. The format has to conform to the ISO 8... |

---

## CustomUserAttributes

---

## DateTimeRange

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end_time` | string | No | The end datetime of the time period. The format conforms to the ISO 8601 profile |
| `start_time` | string | Yes | The start datetime of the time period. The format conforms to the ISO 8601 pr... |

---

## DeleteCommentById

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment_id` | integer | No | An integer representing a comment ID. |
| `delete_type` | string | Yes | How you would like to delete comments. |
| `site_id` | string | Yes | The ID of an existing site |

---

## DeleteComments

---

## DeleteCommentsByParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `delete_type` | string | Yes | How you would like to delete comments. |
| `host_name` | string | Yes | The host name |
| `service_descriptions` | []string | No | If set, the comments for the listed services of the specified host will be re... |
| `site_id` | string | Yes | The ID of an existing site |

---

## DeleteCommentsByQuery

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `delete_type` | string | Yes | How you would like to delete comments. |
| `query` | object | No | An query expression of the Livestatus 'comments' table in nested dictionary f... |
| `site_id` | string | Yes | The ID of an existing site |

---

## DeleteDowntime

---

## DeleteDowntimeById

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `delete_type` | string | Yes | The option how to delete a downtime. |
| `downtime_id` | string | Yes | The id of the downtime |
| `site_id` | string | Yes | The site from which you want to delete a downtime. |

---

## DeleteDowntimeByName

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `delete_type` | string | Yes | The option how to delete a downtime. |
| `host_name` | string | Yes | If set alone, then all downtimes of the host will be removed. |
| `service_descriptions` | []string | No | If set, the downtimes of the listed services of the specified host will be re... |
| `site_id` | string | Yes | The site from which you want to delete a downtime. |

---

## DeleteDowntimeByQuery

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `delete_type` | string | Yes | The option how to delete a downtime. |
| `query` | object | Yes | An query expression of the Livestatus 'downtimes' table in nested dictionary ... |
| `site_id` | string | Yes | The site from which you want to delete a downtime. |

---

## DeleteECEvents

---

## DirectMapping

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `hostname` | string | Yes | The hostname to be replaced. |
| `replace_with` | string | Yes | The replacement string. |

---

## DisabledNotifications

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable` | boolean | No | Option if all notifications should be temporarily disabled |
| `timerange` | object | No | A custom timerange during which notifications are disabled |

---

## DiscoverServices

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_name` | string | Yes | The host of the service which shall be updated. |
| `mode` | string | No | The mode of the discovery action. The 'refresh' mode starts a new service dis... |

---

## DiscoverServicesDeprecated

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `mode` | string | No | The mode of the discovery action. The 'refresh' mode starts a new service dis... |

---

## DiscoveryBackgroundJobStatusObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object |
| `extensions` | object | No | The attributes of the background job |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## DomainObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | string | Yes | The "domain-type" of the object. |
| `extensions` | object | No | All the attributes of the domain object. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## DomainObjectCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | string | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []CollectionItem | No | The collection itself. Each entry in here is part of the collection. |

---

## DowntimeAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `author` | string | Yes | The author of the downtime. |
| `comment` | string | Yes | A comment text. |
| `end_time` | string | Yes | The end time of the downtime. |
| `host_name` | string | Yes | The host name. |
| `is_service` | string | Yes | yes, if this entry is for a service, no if it is for a host. |
| `recurring` | string | Yes | yes if the downtime is recurring, no if it is not. |
| `site_id` | string | Yes | The site id of the downtime. |
| `start_time` | string | Yes | The start time of the downtime. |

---

## DowntimeCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []DowntimeObject | No | A list of downtime objects. |

---

## DowntimeObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The attributes of a downtime. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## ECEventAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `application` | string | Yes | The syslog tag/application this event originated from. |
| `comment` | string | Yes | The event comment. |
| `contact` | string | Yes | The event contact information. |
| `count` | integer | Yes | The number of occurrences of this event within a period. |
| `facility` | string | Yes | The syslog facility. |
| `first` | string | Yes |  |
| `host` | string | Yes | The host name. No exception is raised when the specified host name does not e... |
| `ipaddress` | string | Yes | The IP address where the event originated. |
| `last` | string | Yes |  |
| `phase` | string | Yes | The event phase, open or ack |
| `priority` | string | Yes | The syslog priority. |
| `rule_id` | string | Yes | The ID of the rule. |
| `service_level` | string | Yes | The service level for this event. |
| `site_id` | string | Yes | The site id of the EC event. |
| `state` | string | Yes | The state |
| `text` | string | Yes | The event message text |

---

## ECEventResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The configuration attributes of a site. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## EditUserRole

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `new_alias` | string | No | New alias for the userrole that must be unique. |
| `new_basedon` | string | No | A builtin user role that you want the user role to be based on. |
| `new_permissions` | object | No | A map of permission names to their state. The following values can be set: 'y... |
| `new_role_id` | string | No | New role_id for the userrole that must be unique. |

---

## EmailAndDisplayName

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `address` | string | No |  |
| `display_name` | string | No |  |

---

## EventConsoleAlertAttrsResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `match_event_comment` | CheckboxWithStr | No |  |
| `match_rule_ids` | CheckboxWithListOfStr | No |  |
| `match_syslog_facility` | CheckboxWithStr | No |  |
| `match_syslog_priority` | CheckboxWithSysLogPriority | No |  |

---

## EventConsoleAlertsResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `match_type` | string | Yes |  |
| `state` | string | Yes |  |
| `values` | EventConsoleAlertAttrsResponse | No |  |

---

## EventConsoleResponseCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []ECEventResponse | No | A list of site configuration objects. |

---

## Expr

---

## FailedHosts

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `failed_hosts` | object | No | Detailed error messages on hosts failing the action |
| `succeeded_hosts` | object | No | The list of succeeded host objects |

---

## FilterById

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `event_id` | integer | Yes | The event console ID |
| `filter_type` | string | Yes | The way you would like to filter events. |
| `site_id` | string | Yes | An existing site id |

---

## FilterByParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `filter_type` | string | Yes | The way you would like to filter events. |
| `filters` | FilterParams | Yes |  |
| `site_id` | string | Yes | An existing site id |

---

## FilterByQuery

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `filter_type` | string | Yes | The way you would like to filter events. |
| `query` | object | Yes | An query expression of the Livestatus 'eventconsoleevents' table in nested di... |
| `site_id` | string | Yes | An existing site id |

---

## FilterParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `application` | string | No | Show events that originated from this app. |
| `host` | string | No | The host name. No exception is raised when the specified host name does not e... |
| `phase` | string | No | The event phase, open or ack |
| `state` | string | No | The state |

---

## FilterParamsUpdateAndAcknowledge

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `application` | string | No | Show events that originated from this app. |
| `host` | string | No | The host name. No exception is raised when the specified host name does not e... |
| `state` | string | No | The state |

---

## Folder

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | Data and Meta-Data of this object. |
| `id` | string | No | The full path of the folder, tilde-separated. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | Specific collections or actions applicable to this object. |
| `title` | string | No | The human readable title for this folder. |

---

## FolderCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []Folder | No | A list of folder objects. |

---

## FolderCreateAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No | Choices: * `"ip-v4-only"`: IPv4 only * `"ip-v6-only"`: IPv6 only * `"ip-v4v6"... |
| `tag_agent` | string | No | Choices: * `"cmk-agent"`: API integrations if configured, else Checkmk agent ... |
| `tag_criticality` | string | No | Choices: * `"prod"`: Productive system * `"critical"`: Business critical * `"... |
| `tag_networking` | string | No | Choices: * `"lan"`: Local network (low latency) * `"wan"`: WAN (high latency)... |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No | Choices: * `"no-snmp"`: No SNMP * `"snmp-v2"`: SNMP v2 or v3 * `"snmp-v1"`: S... |

---

## FolderExtensions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | The folder's attributes. Hosts placed in this folder will inherit these attri... |
| `path` | string | No | The full path of this folder, slash delimited. |

---

## FolderMembers

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `hosts` | object | No | A list of links pointing to the actual host-resources. |
| `move` | object | No | An action which triggers the move of this folder to another folder. |

---

## FolderUpdateAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No | Choices: * `"ip-v4-only"`: IPv4 only * `"ip-v6-only"`: IPv6 only * `"ip-v4v6"... |
| `tag_agent` | string | No | Choices: * `"cmk-agent"`: API integrations if configured, else Checkmk agent ... |
| `tag_criticality` | string | No | Choices: * `"prod"`: Productive system * `"critical"`: Business critical * `"... |
| `tag_networking` | string | No | Choices: * `"lan"`: Local network (low latency) * `"wan"`: WAN (high latency)... |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No | Choices: * `"no-snmp"`: No SNMP * `"snmp-v2"`: SNMP v2 or v3 * `"snmp-v1"`: S... |

---

## FolderViewAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `meta_data` | object | No | Read only access to configured metadata. |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `network_scan_result` | object | No | Read only access to the network scan result |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No |  |
| `tag_agent` | string | No |  |
| `tag_criticality` | string | No |  |
| `tag_networking` | string | No |  |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No |  |

---

## FromEmailAndNameCheckbox

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | object | No | The email address and visible name used in the From header of notifications m... |

---

## FromToNotificationNumbers

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `beginning_from` | integer | Yes | Let through notifications counting from this number. The first notification a... |
| `up_to` | integer | Yes | Let through notifications counting upto this number |

---

## FromToServiceLevels

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `from_level` | string | Yes |  |
| `to_level` | string | Yes |  |

---

## Get

---

## GetGraph

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `graph_id` | string | Yes | The ID of the predefined graph. After activating the "Show internal IDs" in t... |
| `host_name` | string | Yes | The hostname to use. |
| `reduce` | string | No | Specify how to reduce a segment of data points to a single data point of the ... |
| `service_description` | string | Yes | The service, whose data to request. |
| `site` | string | No | The name of the site. Even though this is optional, specifying a site will gr... |
| `time_range` | object | Yes | The time range from which to source the metrics. |
| `type` | string | Yes | Specify whether you want to receive a single metric (via metric_id), or a pre... |

---

## GetMetric

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_name` | string | Yes | The hostname to use. |
| `metric_id` | string | Yes | The ID of the single metric.After activating the "Show internal IDs" in the "... |
| `reduce` | string | No | Specify how to reduce a segment of data points to a single data point of the ... |
| `service_description` | string | Yes | The service, whose data to request. |
| `site` | string | No | The name of the site. Even though this is optional, specifying a site will gr... |
| `time_range` | object | Yes | The time range from which to source the metrics. |
| `type` | string | Yes | Specify whether you want to receive a single metric (via metric_id), or a pre... |

---

## GraphCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `metrics` | []Metric | Yes | The actual graph data. |
| `step` | integer | Yes | The interval between two samples in seconds. |
| `time_range` | object | Yes | The time range withing the samples of the response lie. |

---

## HTMLMailPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `bulk_notifications_with_graphs` | CheckboxOneOf | Yes |  |
| `display_graphs_among_each_other` | Checkbox | Yes |  |
| `enable_sync_smtp` | CheckboxOneOf | Yes |  |
| `from_details` | CheckboxOneOf | Yes |  |
| `graphs_per_notification` | CheckboxOneOf | Yes |  |
| `info_to_be_displayed_in_the_email_body` | CheckboxOneOf | Yes |  |
| `insert_html_section_between_body_and_table` | CheckboxOneOf | Yes |  |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `reply_to` | CheckboxOneOf | Yes |  |
| `send_separate_notification_to_every_recipient` | Checkbox | Yes |  |
| `sort_order_for_bulk_notificaions` | CheckboxOneOf | Yes |  |
| `subject_for_host_notifications` | CheckboxOneOf | Yes |  |
| `subject_for_service_notifications` | CheckboxOneOf | Yes |  |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |

---

## Heartbeat

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `interval` | integer | No | The heartbeat interval for the TCP connection. |
| `timeout` | number | No | The heartbeat timeout for the TCP connection. |

---

## Heartbeat1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `interval` | integer | No | The heartbeat interval for the TCP connection. |
| `timeout` | number | No | The heartbeat timeout for the TCP connection. |

---

## Host

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | object | Yes |  |

---

## HostConditions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_choice` | object | Yes |  |
| `host_folder` | string | Yes |  |
| `host_labels` | object | Yes |  |
| `host_tags` | object | Yes |  |

---

## HostConfig

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | Yes | The domain type of the object. |
| `extensions` | object | No | All the data and metadata of this host. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | All the members of the host object. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## HostConfigCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []HostConfig | No | A list of host objects. |

---

## HostConfigSchemaInternal

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `is_cluster` | boolean | Yes | Indicates if the host is a cluster host. |
| `site` | string | Yes | The site the host is monitored on. |

---

## HostContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `groups` | []string | Yes | A list of contact groups. |
| `recurse_perms` | boolean | No | Give these groups also permission on all sub-folders. |
| `recurse_use` | boolean | No | Add these groups as contacts to all hosts in all sub-folders of this folder. |
| `use` | boolean | No | Add these contact groups to the host. |
| `use_for_services` | boolean | No | <p>Always add host contact groups also to its services.</p>With this option c... |

---

## HostCreateAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `additional_ipv4addresses` | []string | No | A list of IPv4 addresses. |
| `additional_ipv6addresses` | []string | No | A list of IPv6 addresses. |
| `alias` | string | No | Add a comment or describe this host |
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `inventory_failed` | boolean | No | Whether or not the last bulk discovery failed. It is set to True once it fail... |
| `ipaddress` | string | No | An IPv4 address. |
| `ipv6address` | string | No | An IPv6 address. |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `locked_attributes` | []string | No | Name of host attributes which are locked in the UI. |
| `locked_by` | object | No | Identity of the entity which locked the locked_attributes. The identity is bu... |
| `management_address` | string | No | Address (IPv4, IPv6 or hostname) under which the management board can be reac... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No | Choices: * `"ip-v4-only"`: IPv4 only * `"ip-v6-only"`: IPv6 only * `"ip-v4v6"... |
| `tag_agent` | string | No | Choices: * `"cmk-agent"`: API integrations if configured, else Checkmk agent ... |
| `tag_criticality` | string | No | Choices: * `"prod"`: Productive system * `"critical"`: Business critical * `"... |
| `tag_networking` | string | No | Choices: * `"lan"`: Local network (low latency) * `"wan"`: WAN (high latency)... |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No | Choices: * `"no-snmp"`: No SNMP * `"snmp-v2"`: SNMP v2 or v3 * `"snmp-v1"`: S... |

---

## HostEventType

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledgement_of_problem` | boolean | Yes |  |
| `alert_handler_execution_failed` | boolean | Yes |  |
| `alert_handler_execution_successful` | boolean | Yes |  |
| `any_down` | boolean | Yes |  |
| `any_unreachable` | boolean | Yes |  |
| `any_up` | boolean | Yes |  |
| `down_unreachable` | boolean | Yes |  |
| `down_up` | boolean | Yes |  |
| `start_or_end_of_flapping_state` | boolean | Yes |  |
| `start_or_end_of_scheduled_downtime` | boolean | Yes |  |
| `unreachable_down` | boolean | Yes |  |
| `unreachable_up` | boolean | Yes |  |
| `up_down` | boolean | Yes |  |
| `up_unreachable` | boolean | Yes |  |

---

## HostExtensions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Attributes of this host. |
| `cluster_nodes` | []string | No | In the case this is a cluster host, these are the cluster nodes. |
| `effective_attributes` | object | No | All attributes of this host and all parent folders. |
| `folder` | string | No | The folder, in which this host resides. Path delimiters can be either `~`, `/... |
| `is_cluster` | boolean | No | If this is a cluster host, i.e. a container for other hosts. |
| `is_offline` | boolean | No | Whether the host is offline |

---

## HostExtensionsEffectiveAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `additional_ipv4addresses` | []string | No | A list of IPv4 addresses. |
| `additional_ipv6addresses` | []string | No | A list of IPv6 addresses. |
| `alias` | string | No | Add a comment or describe this host |
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `inventory_failed` | boolean | No | Whether or not the last bulk discovery failed. It is set to True once it fail... |
| `ipaddress` | string | No | An IPv4 address. |
| `ipv6address` | string | No | An IPv6 address. |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `locked_attributes` | []string | No | Name of host attributes which are locked in the UI. |
| `locked_by` | object | No | Identity of the entity which locked the locked_attributes. The identity is bu... |
| `management_address` | string | No | Address (IPv4, IPv6 or hostname) under which the management board can be reac... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `meta_data` | object | No | Read only access to configured metadata. |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `network_scan_result` | object | No | Read only access to the network scan result |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No |  |
| `tag_agent` | string | No |  |
| `tag_criticality` | string | No |  |
| `tag_networking` | string | No |  |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No |  |

---

## HostGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | All the attributes of the domain object. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## HostGroupCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []HostGroup | No | A list of host group objects. |

---

## HostMembers

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `folder_config` | object | No | The folder in which this host resides. It is represented by a hexadecimal ide... |

---

## HostOrServiceCondition

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `match_on` | []string | No | A list of string matching regular expressions. |
| `operator` | string | No | How the hosts or services should be matched. * one_of - will match if any of ... |

---

## HostTag

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `aux_tags` | []string | No | The list of auxiliary tag ids. Built-in tags (ip-v4, ip-v6, snmp, tcp, ping) ... |
| `ident` | string | No | An unique id for the tag |
| `title` | string | Yes | The title of the tag |

---

## HostTag1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `aux_tags` | []string | No | The auxiliary tags this tag included in. |
| `id` | string | No | The unique identifier of this host tag |
| `title` | string | No | The title of this host tag |

---

## HostTagExtensions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `tags` | []HostTag1 | No | The list of tags in this group. |
| `topic` | string | No | The topic this host tag group is organized in. |

---

## HostTagGroupCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []ConcreteHostTagGroup | No | A list of host tag group objects. |

---

## HostTagValues

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `checkmk_agent_api_integration` | string | Yes |  |
| `criticality` | string | Yes |  |
| `ip_address_family` | string | Yes |  |
| `ip_v4` | string | Yes |  |
| `ip_v6` | string | Yes |  |
| `monitor_via_checkmkagent` | string | Yes |  |
| `monitor_via_checkmkagent_or_specialagent` | string | Yes |  |
| `monitor_via_snmp` | string | Yes |  |
| `networking_segment` | string | Yes |  |
| `only_ping_this_device` | string | Yes |  |
| `piggyback` | string | Yes |  |
| `snmp` | string | Yes |  |

---

## HostUpdateAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `additional_ipv4addresses` | []string | No | A list of IPv4 addresses. |
| `additional_ipv6addresses` | []string | No | A list of IPv6 addresses. |
| `alias` | string | No | Add a comment or describe this host |
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `inventory_failed` | boolean | No | Whether or not the last bulk discovery failed. It is set to True once it fail... |
| `ipaddress` | string | No | An IPv4 address. |
| `ipv6address` | string | No | An IPv6 address. |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `locked_attributes` | []string | No | Name of host attributes which are locked in the UI. |
| `locked_by` | object | No | Identity of the entity which locked the locked_attributes. The identity is bu... |
| `management_address` | string | No | Address (IPv4, IPv6 or hostname) under which the management board can be reac... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No | Choices: * `"ip-v4-only"`: IPv4 only * `"ip-v6-only"`: IPv6 only * `"ip-v4v6"... |
| `tag_agent` | string | No | Choices: * `"cmk-agent"`: API integrations if configured, else Checkmk agent ... |
| `tag_criticality` | string | No | Choices: * `"prod"`: Productive system * `"critical"`: Business critical * `"... |
| `tag_networking` | string | No | Choices: * `"lan"`: Local network (low latency) * `"wan"`: WAN (high latency)... |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No | Choices: * `"no-snmp"`: No SNMP * `"snmp-v2"`: SNMP v2 or v3 * `"snmp-v1"`: S... |

---

## HostViewAttribute

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `additional_ipv4addresses` | []string | No | A list of IPv4 addresses. |
| `additional_ipv6addresses` | []string | No | A list of IPv6 addresses. |
| `alias` | string | No | Add a comment or describe this host |
| `contactgroups` | object | No | Only members of the contact groups listed here have Setup permission for the ... |
| `inventory_failed` | boolean | No | Whether or not the last bulk discovery failed. It is set to True once it fail... |
| `ipaddress` | string | No | An IPv4 address. |
| `ipv6address` | string | No | An IPv6 address. |
| `labels` | object | No | Labels allow you to flexibly group your hosts in order to refer to them later... |
| `locked_attributes` | []string | No | Name of host attributes which are locked in the UI. |
| `locked_by` | object | No | Identity of the entity which locked the locked_attributes. The identity is bu... |
| `management_address` | string | No | Address (IPv4, IPv6 or hostname) under which the management board can be reac... |
| `management_ipmi_credentials` | object | No | IPMI credentials |
| `management_protocol` | string | No | The protocol used to connect to the management board. Valid options are: * `n... |
| `management_snmp_community` | object | No | SNMP credentials |
| `meta_data` | object | No | Read only access to configured metadata. |
| `network_scan` | object | No | Configuration for automatic network scan. Pings will besent to each IP addres... |
| `network_scan_result` | object | No | Read only access to the network scan result |
| `parents` | []string | No | A list of parents of this host. |
| `site` | string | No | The site that should monitor this host. |
| `snmp_community` | object | No | The SNMP access configuration. A configured SNMP v1/v2 community here will ha... |
| `tag_address_family` | string | No |  |
| `tag_agent` | string | No |  |
| `tag_criticality` | string | No |  |
| `tag_networking` | string | No |  |
| `tag_piggyback` | string | No | By default, each host has a piggyback data source.<br><br><b>Use piggyback da... |
| `tag_snmp_ds` | string | No |  |

---

## IPAddressRange

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `from_address` | string | No | The first IPv4 address of this range. |
| `to_address` | string | No | The last IPv4 address of this range. |
| `type` | object | No | A range of addresses. |

---

## IPAddresses

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `addresses` | []string | No |  |
| `type` | object | No | A list of single IPv4 addresses. |

---

## IPMIParameters

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `password` | string | Yes |  |
| `username` | string | Yes |  |

---

## IPNetwork

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `network` | string | No | A IPv4 network in CIDR notation. Minimum prefix length is 8 bit, maximum pref... |
| `type` | object | No | A single IPv4 network in CIDR notation. |

---

## IPRangeWithRegexp

---

## IPRegexp

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `regexp_list` | []string | No | A list of regular expressions which are matched against the found IP addresse... |
| `type` | object | No | IPv4 addresses which match a regexp pattern |

---

## IdleOption

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `duration` | integer | No | The duration in seconds of the individual idle timeout if individual is selec... |
| `option` | string | Yes | Specify if the idle timeout should use the global configuration, be disabled ... |

---

## IlertAPIKey

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | Yes |  |
| `option` | string | Yes |  |

---

## IlertKeyOrStoreSelector

---

## IlertPasswordStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## IlertPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `api_key` | IlertKeyOrStoreSelector | Yes |  |
| `custom_summary_for_host_alerts` | string | Yes | A custom summary for host alerts |
| `custom_summary_for_service_alerts` | string | Yes | A custom summary for service alerts |
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `notification_priority` | string | Yes | HIGH - with escalation, LOW - without escalation |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |

---

## IncidentParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `caller` | string | Yes | Caller is the user on behalf of whom the incident is being reported within Se... |
| `host_description` | CheckboxOneOf | No |  |
| `host_short_description` | CheckboxOneOf | No |  |
| `impact` | CheckboxOneOf | No |  |
| `service_description` | CheckboxOneOf | No |  |
| `service_short_description` | CheckboxOneOf | No |  |
| `state_acknowledgement` | CheckboxOneOf | No |  |
| `state_downtime` | CheckboxOneOf | No |  |
| `state_recovery` | CheckboxOneOf | No |  |
| `urgency` | CheckboxOneOf | No |  |

---

## InputContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The name used for displaying in the GUI. |
| `name` | string | Yes | The name of the contact group. |

---

## InputHostGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The name used for displaying in the GUI. |
| `name` | string | Yes | A name used as identifier |

---

## InputHostTagGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `help` | string | No | A help description for the tag group |
| `ident` | string | Yes | An id for the host tag group |
| `tags` | []HostTag | Yes | A list of host tags belonging to the host tag group |
| `title` | string | Yes | A title for the host tag |
| `topic` | string | No | Different tags can be grouped in a topic |

---

## InputPassword

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No | A comment for the password |
| `documentation_url` | string | No | An optional URL pointing to documentation or any other page. You can use eith... |
| `ident` | string | Yes | An unique identifier for the password |
| `owner` | string | Yes | Each password is owned by a group of users which are able to edit, delete and... |
| `password` | string | Yes | The password string |
| `shared` | []string | No | The list of members to share the password with |
| `title` | string | Yes | A title for the password |

---

## InputRuleObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | object | No | Conditions. |
| `folder` | string | Yes | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `properties` | object | No | Configuration values for rules. |
| `ruleset` | string | Yes | Name of rule set. |
| `value_raw` | string | No | The raw parameter value for this rule. To create the correct structure, for n... |

---

## InputServiceGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The name used for displaying in the GUI. |
| `name` | string | Yes | A name used as identifier |

---

## InstalledVersions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `demo` | boolean | No | Whether this is a demo version or not. |
| `edition` | string | No | The Checkmk edition. |
| `group` | string | No | The Apache WSGI application group this call was made on. |
| `rest_api` | object | No | The REST-API version |
| `site` | string | No | The site where this API call was made on. |
| `versions` | object | No | Some version numbers |

---

## JiraPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `host_custom_id` | string | Yes | The numerical JIRA custom field ID for host problems |
| `host_summary` | object | Yes | Here you are allowed to use all macros that are defined in the notification c... |
| `issue_type_id` | string | Yes | The numerical JIRA issue type ID. If not set, it will be retrieved from a cus... |
| `jira_url` | string | No | Configure the JIRA URL here |
| `label` | object | Yes | Here you can set a custom label for new issues. If not set, 'monitoring' will... |
| `monitoring_url` | string | Yes | Configure the base URL for the Monitoring Web-GUI here. Include the site name... |
| `optional_timeout` | object | Yes | Here you can configure timeout settings. |
| `password` | string | Yes | The password entered here is stored in plain text within the monitoring site.... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `priority_id` | object | Yes | The numerical JIRA priority ID. If not set, it will be retrieved from a custo... |
| `project_id` | string | Yes | The numerical JIRA project ID. If not set, it will be retrieved from a custom... |
| `resolution_id` | object | Yes | The numerical JIRA resolution transition ID. 11 - 'To Do', 21 - 'In Progress'... |
| `service_custom_id` | string | Yes | The numerical JIRA custom field ID for service problems |
| `service_summary` | object | Yes | Here you are allowed to use all macros that are defined in the notification c... |
| `site_custom_id` | object | Yes | The numerical ID of the JIRA custom field for sites. Please use this option i... |
| `username` | string | Yes | Configure the user name here |

---

## JobLogs

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `progress` | []string | No | The list of progress related logs |
| `result` | []string | No | The list of result related logs |

---

## LabelCondition

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | Yes | The key of the label. e.g. 'os' in 'os:windows' |
| `operator` | string | No | How the label should be matched. |
| `value` | string | Yes | The value of the label. e.g. 'windows' in 'os:windows' |

---

## Link

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `body_params` | object | No | A map of values that shall be sent in the request body. If this is present,th... |
| `domainType` | object | Yes |  |
| `href` | string | Yes | The (absolute) address of the related resource. Any characters that are inval... |
| `method` | string | Yes | The HTTP method to use to traverse the link (get, post, put or delete) |
| `rel` | string | Yes | Indicates the nature of the relationship of the related resource to the resou... |
| `title` | string | No | string that the consuming application may use to render the link without havi... |
| `type` | string | Yes | The content-type that the linked resource will return |

---

## LinkHostUUID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `uuid` | string | Yes | A valid UUID. |

---

## LockedBy

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `instance_id` | string | Yes | Instance ID |
| `program_id` | string | Yes | Program ID |
| `site_id` | string | Yes | Site ID |

---

## LogicalExpr

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `expr` | []object | No |  |
| `op` | string | No | The operator. |

---

## MSTeamsExplicitWebhookUrl

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `url` | string | Yes |  |

---

## MSTeamsPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `affected_host_groups` | object | Yes | Enable/disable if we show affected host groups in the created message |
| `host_details` | object | Yes | Enable/disable the details for host notifications |
| `host_summary` | object | Yes | Enable/disable the summary for host notifications |
| `host_title` | object | Yes | Enable/disable the title for host notifications |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `service_details` | object | Yes | Enable/disable the details for service notifications |
| `service_summary` | object | Yes | Enable/disable the summary for service notifications |
| `service_title` | object | Yes | Enable/disable the title for service notifications |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |
| `webhook_url` | MSTeamsUrlOrStoreSelector | Yes |  |

---

## MSTeamsURLResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |
| `url` | string | No |  |

---

## MSTeamsUrlOrStoreSelector

---

## MatchCustomMacros

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | []CustomMacro | No |  |

---

## MatchEventConsoleAlertsResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | EventConsoleAlertsResponse | No |  |

---

## MetaData

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `created_at` | string | No | When has this object been created. |
| `created_by` | string | No | The user id under which this object has been created. |
| `updated_at` | string | No | When this object was last changed. |

---

## Metric

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `color` | string | No | The color of the metric as displayed in Checkmk. Color is in HTML notation. |
| `data_points` | []number | Yes | The samples of the metric. |
| `line_type` | string | Yes | The line type to use. |
| `title` | string | Yes | The title of the graph. |

---

## MgmntTypeCaseParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `params` | CaseParams | No |  |

---

## MgmntTypeIncidentParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `params` | IncidentParams | No |  |

---

## MgmntTypeSelector

---

## MkEventDPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ip_address_of_remote_event_console` | CheckboxOneOf | Yes |  |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `syslog_facility_to_use` | CheckboxOneOf | Yes |  |

---

## MoveFolder

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `destination` | string | Yes | Where the folder has to be moved to. Path delimiters can be either `~`, `/` o... |

---

## MoveHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `target_folder` | string | Yes | The path of the target folder where the host is supposed to be moved to. Path... |

---

## MoveRuleTo

---

## MoveToFolder

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `folder` | string | No | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `position` | string | No | The type of position to move to. |

---

## MoveToSpecificRule

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `position` | string | No | The type of position to move to. |
| `rule_id` | string | No | The UUID of the rule to move after/before. |

---

## NetworkScan

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `addresses` | []IPRangeWithRegexp | Yes | IPv4 addresses to include. |
| `exclude_addresses` | []IPRangeWithRegexp | No | IPv4 addresses to exclude. |
| `max_parallel_pings` | integer | No | Set the maximum number of concurrent pings sent to target IP addresses. |
| `run_as` | string | No | Execute the network scan in the Checkmk user context of the chosen user. This... |
| `scan_interval` | integer | No | Scan interval in seconds. Default is 1 day, minimum is 1 hour. |
| `set_ip_address` | boolean | No | When set, the found IPv4 address is set on the discovered host. |
| `tag_criticality` | string | No | Specify which criticality tag to set on the host created by the network scan.... |
| `time_allowed` | []TimeAllowedRange | Yes | Only execute the discovery during this time range each day.. |
| `translate_names` | TranslateNames | No |  |

---

## NetworkScanResult

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end` | string | No | When the scan finished. Will be Null if not yet run. |
| `output` | string | No | Short human readable description of what is happening. |
| `start` | string | No | When the scan started |
| `state` | string | No | Last scan result |

---

## NotExpr

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `expr` | object | No | The query expression to negate. |
| `op` | string | No | The operator. In this case `not`. |

---

## NotificationBulking

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `bulk_outside_timeperiod` | BulkOutsideTimePeriodValue | Yes |  |
| `max_bulk_size` | integer | Yes | At most that many Notifications are kept back for bulking. A value of 1 essen... |
| `notification_bulks_based_on` | []string | Yes |  |
| `notification_bulks_based_on_custom_macros` | []string | No |  |
| `state` | string | Yes |  |
| `subject_for_bulk_notifications` | CheckboxWithStrValue | No |  |
| `time_horizon` | integer | Yes | Notifications are kept back for bulking at most for this time (seconds) |
| `time_period` | string | Yes |  |

---

## NotificationBulkingCheckbox

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `state` | string | Yes |  |
| `value` | WhenToBulk | Yes |  |

---

## NotificationBulkingCommonAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `max_bulk_size` | integer | Yes | At most that many Notifications are kept back for bulking. A value of 1 essen... |
| `notification_bulks_based_on` | []string | Yes |  |
| `notification_bulks_based_on_custom_macros` | []string | No |  |
| `state` | string | Yes |  |
| `subject_for_bulk_notifications` | CheckboxWithStrValue | No |  |
| `time_horizon` | integer | Yes | Notifications are kept back for bulking at most for this time (seconds) |

---

## NotificationPlugin

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `notification_bulking` | NotificationBulkingCheckbox | No |  |
| `notify_plugin` | PluginBase1 | No |  |

---

## NotificationRuleAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | ConditionsAttributes | No |  |
| `contact_selection` | ContactSelectionAttributes | No |  |
| `notification_method` | NotificationPlugin | No |  |
| `rule_properties` | RulePropertiesAttributes | No |  |

---

## NotificationRuleConfig

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `rule_config` | NotificationRuleAttributes | No |  |

---

## NotificationRuleRequest

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `rule_config` | object | Yes |  |

---

## NotificationRuleResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The configuration attributes of a notification rule. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## NotificationRuleResponseCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []NotificationRuleResponse | No | A list of notification rule objects. |

---

## ObjectActionMember

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disabledReason` | string | No | Provides the reason (or the literal "disabled") why an object property or col... |
| `id` | string | Yes |  |
| `invalidReason` | string | No | Provides the reason (or the literal "invalid") why a proposed value for a pro... |
| `links` | []Link | Yes | list of links to other resources. |
| `memberType` | object | No |  |
| `name` | string | No |  |
| `parameters` | object | No |  |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `x-ro-invalidReason` | string | No | Provides the reason why a SET OF proposed values for properties or arguments ... |

---

## ObjectCollectionMember

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disabledReason` | string | No | Provides the reason (or the literal "disabled") why an object property or col... |
| `id` | string | Yes |  |
| `invalidReason` | string | No | Provides the reason (or the literal "invalid") why a proposed value for a pro... |
| `links` | []Link | Yes | list of links to other resources. |
| `memberType` | object | No |  |
| `name` | string | No |  |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []Link | No |  |
| `x-ro-invalidReason` | string | No | Provides the reason why a SET OF proposed values for properties or arguments ... |

---

## ObjectProperty

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `extensions` | object | No | Additional attributes alongside the property. |
| `id` | string | No | The unique name of this property, local to this domain type. |
| `links` | []Link | Yes | list of links to other resources. |
| `value` | []string | No | The value of the property. In this case a list. |

---

## OpsGenieExplicitKey

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | Yes |  |
| `option` | string | Yes |  |

---

## OpsGeniePluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `actions` | object | Yes | Custom actions that will be available for the alert. You may paste a text fro... |
| `api_key` | OpsGenisStoreOrExplicitKeySelector | Yes |  |
| `desc_for_host_alerts` | object | Yes | Description field of host alert that is generally used to provide a detailed ... |
| `desc_for_service_alerts` | object | Yes | Description field of service alert that is generally used to provide a detail... |
| `domain` | object | Yes | If you have an european account, please set the domain of your opsgenie. Spec... |
| `entity` | object | Yes | Is used to specify which domain the alert is related to |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `message_for_host_alerts` | object | Yes |  |
| `message_for_service_alerts` | object | Yes |  |
| `note_while_closing` | object | Yes | Additional note that will be added while closing the alert |
| `note_while_creating` | object | Yes | Additional note that will be added while creating the alert |
| `owner` | object | Yes | Sets the user of the alert. Display name of the request owner |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `priority` | CheckboxOneOf | Yes |  |
| `responsible_teams` | object | Yes | Team names which will be responsible for the alert. If the API Key belongs to... |
| `source` | object | Yes | Source field of the alert. Default value is IP address of the incoming request |
| `tags` | object | Yes | Tags of the alert. You may paste a text from your clipboard which contains se... |

---

## OpsGenieStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## OpsGenisStoreOrExplicitKeySelector

---

## PagerDutyAPIKeyStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## PagerDutyExplicitKey

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | Yes |  |
| `option` | string | Yes |  |

---

## PagerDutyPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `integration_key` | PagerDutyStoreOrIntegrationKeySelector | Yes |  |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |

---

## PagerDutyStoreOrIntegrationKeySelector

---

## Parent

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `type` | object | Yes |  |

---

## PasswordCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []PasswordObject | No | A list of password objects. |

---

## PasswordExtension

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No | A comment for the password |
| `documentation_url` | string | No | The URL pointing to documentation or any other page. |
| `ident` | string | No | The unique identifier for the password |
| `owned_by` | string | No | The owner of the password who is able to edit, delete and use existing passwo... |
| `shared` | []string | No | The list of members the password is shared with |
| `title` | string | No | The title for the password |

---

## PasswordObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The type of the domain-object. |
| `extensions` | object | No | All the attributes of the domain object. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## PendingChangesCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `value` | []ChangesFields | No | The changes that are pending |

---

## PluginBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | No | Create notifications with parameters or cancel previous notifications |
| `plugin_params` | PluginName | Yes |  |

---

## PluginBase1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |

---

## PluginName

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |

---

## PluginOptionsSelector

---

## PluginSelector

---

## PluginWithParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | No | Create notifications with parameters or cancel previous notifications |
| `plugin_params` | PluginSelector | Yes |  |

---

## ProxyAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `global_settings` | boolean | Yes | When use_livestatus_daemon is set to 'with_proxy', you can set this to True t... |
| `params` | object | No | The live status proxy daemon parameters. |
| `tcp` | object | No | Allow access via TCP configuration. |
| `use_livestatus_daemon` | string | Yes | Use livestatus daemon with direct connection or with livestatus proxy. |

---

## ProxyAttributes1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `global_settings` | boolean | No | When Livestatus proxy daemon is set, you can enable this to use global settin... |
| `params` | object | No | The live status proxy daemon parameters. |
| `tcp` | object | No | Allow access via TCP configuration. |
| `use_livestatus_daemon` | string | Yes | Use livestatus daemon with direct connection or with livestatus proxy. |

---

## ProxyOrDirect

---

## ProxyParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `cache` | boolean | No | Enable caching. |
| `channel_timeout` | number | No | The timeout waiting for a free channel. |
| `channels` | integer | No | The number of channels to keep open. |
| `connect_retry` | number | No | The cooling period after failed connect/heartbeat. |
| `heartbeat` | object | No | The heartbeat interval and timeout configuration. |
| `query_timeout` | number | No | The total query timeout. |

---

## ProxyParams1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `cache` | boolean | No | Enable caching. |
| `channel_timeout` | number | No | The timeout waiting for a free channel. |
| `channels` | integer | No | The number of channels to keep open. |
| `connect_retry` | number | No | The cooling period after failed connect/heartbeat. |
| `heartbeat` | object | No | The heartbeat interval and timeout configuration. |
| `query_timeout` | number | No | The total query timeout. |

---

## ProxyTcp

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `only_from` | []string | No | Restrict access to these IP addresses. |
| `port` | integer | Yes | The TCP port to connect to. |
| `tls` | boolean | No | Encrypt TCP Livestatus connections. |

---

## ProxyTcp1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `only_from` | []string | No | Restrict access to these IP addresses. |
| `port` | integer | No | The livestatus proxy TCP port. |
| `tls` | boolean | No | Encrypt TCP Livestatus connections. |

---

## PushOverPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `api_key` | string | Yes | You need to provide a valid API key to be able to send push notifications usi... |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `priority` | CheckboxOneOf | Yes |  |
| `sound` | CheckboxOneOf | Yes |  |
| `url_prefix_for_links_to_checkmk` | object | Yes | If you specify an URL prefix here, then several parts of the email body are a... |
| `user_group_key` | string | Yes | Configure the user or group to receive the notifications by providing the use... |

---

## ReferTo

---

## RegexpRewrites

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `replace_with` | string | Yes | The replacement string. Match-groups can only be identified by `\1`, `\2`, et... |
| `search` | string | Yes | The search regexp. May contain match-groups, conditional matches, etc. This f... |

---

## RegisterHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `uuid` | string | Yes | A valid UUID. |

---

## RenameHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `new_name` | string | Yes | The new name of the existing host. |

---

## RuleCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | Domain type of this object. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []RuleObject | No | The collection itself. Each entry in here is part of the collection. |

---

## RuleConditions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `event_console_alerts` | CheckboxOneOf | Yes |  |
| `match_check_types` | CheckboxOneOf | Yes |  |
| `match_contact_groups` | CheckboxOneOf | Yes |  |
| `match_exclude_hosts` | CheckboxOneOf | Yes |  |
| `match_exclude_service_groups` | CheckboxOneOf | Yes |  |
| `match_exclude_service_groups_regex` | CheckboxOneOf | Yes |  |
| `match_exclude_services` | CheckboxOneOf | Yes |  |
| `match_folder` | CheckboxOneOf | Yes |  |
| `match_host_event_type` | CheckboxOneOf | Yes |  |
| `match_host_groups` | CheckboxOneOf | Yes |  |
| `match_host_labels` | CheckboxOneOf | Yes |  |
| `match_host_tags` | CheckboxOneOf | Yes |  |
| `match_hosts` | CheckboxOneOf | Yes |  |
| `match_notification_comment` | CheckboxOneOf | Yes |  |
| `match_only_during_time_period` | CheckboxOneOf | Yes |  |
| `match_plugin_output` | CheckboxOneOf | Yes |  |
| `match_service_event_type` | CheckboxOneOf | Yes |  |
| `match_service_groups` | CheckboxOneOf | Yes |  |
| `match_service_groups_regex` | CheckboxOneOf | Yes |  |
| `match_service_labels` | CheckboxOneOf | Yes |  |
| `match_service_levels` | CheckboxOneOf | Yes |  |
| `match_services` | CheckboxOneOf | Yes |  |
| `match_sites` | CheckboxOneOf | Yes |  |
| `restrict_to_notification_numbers` | CheckboxOneOf | Yes |  |
| `throttle_periodic_notifications` | CheckboxOneOf | Yes |  |

---

## RuleConditions1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_labels` | []LabelCondition | No | Further restrict this rule by applying host label conditions. |
| `host_name` | object | No | Here you can enter a list of explicit host names that the rule should or shou... |
| `host_tags` | []TagCondition | No | The rule will only be applied to hosts fulfilling all the host tag conditions... |
| `service_description` | object | No | Specify a list of service patterns this rule shall apply to. * The patterns m... |
| `service_labels` | []LabelCondition | No | Restrict the application of the rule, by checking against service label condi... |

---

## RuleExtensions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | object | No | Conditions. |
| `folder` | string | Yes | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `folder_index` | integer | No | The position of this rule in the chain in this folder. |
| `properties` | object | No | Property values of this rule. |
| `ruleset` | string | No | The name of the ruleset. |
| `value_raw` | string | No | The raw parameter value for this rule. |

---

## RuleNotification

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `conditions` | RuleConditions | Yes |  |
| `contact_selection` | ContactSelection | Yes |  |
| `notification_method` | RuleNotificationMethod | Yes |  |
| `rule_properties` | RuleProperties | Yes |  |

---

## RuleNotificationMethod

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `notification_bulking` | CheckboxOneOf | Yes |  |
| `notify_plugin` | PluginOptionsSelector | Yes |  |

---

## RuleObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | Domain type of this object. |
| `extensions` | object | No | Attributes specific to rule objects. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## RuleProperties

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `allow_users_to_deactivate` | object | Yes | If you set this option then users are allowed to deactivate notifications tha... |
| `comment` | string | Yes | An optional comment that may be used to explain the purpose of this object. |
| `description` | string | Yes | A description or title of this rule. |
| `do_not_apply_this_rule` | object | Yes | Disabled rules are kept in the configuration but are not applied. |
| `documentation_url` | string | Yes | An optional URL pointing to documentation or any other page. This will be dis... |

---

## RuleProperties1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No | Any comment string. |
| `description` | string | No | A description for this rule to inform other users about its intent. |
| `disabled` | boolean | No | When set to False, the rule will be evaluated. Default is False. |
| `documentation_url` | string | No | An URL (e.g. an internal Wiki entry) which explains this rule. |

---

## RulePropertiesAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `allow_users_to_deactivate` | object | Yes | If you set this option then users are allowed to deactivate notifications tha... |
| `comment` | string | Yes | An optional comment that may be used to explain the purpose of this object. |
| `description` | string | Yes | A description or title of this rule. |
| `do_not_apply_this_rule` | object | Yes | Disabled rules are kept in the configuration but are not applied. |
| `documentation_url` | string | Yes | An optional URL pointing to documentation or any other page. This will be dis... |

---

## RulesetCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | Domain type of this object. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []CollectionItem | No | The collection itself. Each entry in here is part of the collection. |

---

## RulesetExtensions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `folder` | string | Yes | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `name` | string | No | The name of the ruleset |
| `number_of_rules` | integer | No | The number of rules of this ruleset. |

---

## RulesetObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | Domain type of this object. |
| `extensions` | object | No | Specific attributes related to rulesets. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## SMSAPIExplicitPassword

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `password` | string | Yes |  |

---

## SMSAPIPStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## SMSAPIPasswordSelector

---

## SMSAPIPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `modem_type` | string | Yes | Choose what modem is used. Currently supported is only Teltonika-TRB140. |
| `modem_url` | string | Yes | Configure your modem URL here |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `timeout` | string | Yes | Here you can configure timeout settings |
| `user_password` | SMSAPIPasswordSelector | Yes |  |
| `username` | string | Yes | Configure the user name here |

---

## SMSPluginBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `params` | []string | Yes | The given parameters are available in scripts as NOTIFY_PARAMETER_1, NOTIFY_P... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |

---

## SNMPCommunity

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `community` | string | Yes | SNMP community (SNMP Versions 1 and 2c) |
| `type` | object | No |  |

---

## SNMPCredentials

---

## SNMPv3AuthNoPrivacy

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_password` | string | Yes | Authentication pass phrase. |
| `auth_protocol` | string | Yes | Authentication protocol. |
| `security_name` | string | Yes | Security name |
| `type` | object | No | The type of credentials to use. |

---

## SNMPv3AuthPrivacy

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_password` | string | Yes | Authentication pass phrase. |
| `auth_protocol` | string | Yes | Authentication protocol. |
| `privacy_password` | string | Yes | Privacy pass phrase. If filled, privacy_protocol needs to be selected as well. |
| `privacy_protocol` | string | Yes | The privacy protocol. The only supported values in the Raw Edition are CBC-DE... |
| `security_name` | string | Yes | Security name |
| `type` | object | No | SNMPv3 with authentication and privacy. |

---

## SNMPv3NoAuthNoPrivacy

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `security_name` | string | Yes | Security name |
| `type` | object | No | The type of credentials to use. |

---

## ServiceConditions

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host_choice` | object | Yes |  |
| `host_folder` | string | Yes |  |
| `host_labels` | object | Yes |  |
| `host_tags` | object | Yes |  |
| `service_labels` | object | Yes |  |
| `service_regex` | string | Yes |  |

---

## ServiceEventType

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `acknowledgement_of_problem` | boolean | Yes |  |
| `alert_handler_execution_failed` | boolean | Yes |  |
| `alert_handler_execution_successful` | boolean | Yes |  |
| `any_crit` | boolean | Yes |  |
| `any_ok` | boolean | Yes |  |
| `any_unknown` | boolean | Yes |  |
| `any_warn` | boolean | Yes |  |
| `crit_ok` | boolean | Yes |  |
| `crit_unknown` | boolean | Yes |  |
| `crit_warn` | boolean | Yes |  |
| `ok_crit` | boolean | Yes |  |
| `ok_ok` | boolean | Yes |  |
| `ok_unknown` | boolean | Yes |  |
| `ok_warn` | boolean | Yes |  |
| `start_or_end_of_flapping_state` | boolean | Yes |  |
| `start_or_end_of_scheduled_downtime` | boolean | Yes |  |
| `unknown_crit` | boolean | Yes |  |
| `unknown_ok` | boolean | Yes |  |
| `unknown_warn` | boolean | Yes |  |
| `warn_crit` | boolean | Yes |  |
| `warn_ok` | boolean | Yes |  |
| `warn_unknown` | boolean | Yes |  |

---

## ServiceGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | All the attributes of the domain object. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## ServiceGroupCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []ServiceGroup | No | A list of service group objects. |

---

## ServiceGroupsRegex

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `match_type` | string | Yes |  |
| `regex_list` | []string | Yes | The text entered in this list is handled as a regular expression pattern |

---

## ServiceNowExplicitPassword

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `password` | string | Yes |  |

---

## ServiceNowPasswordSelector

---

## ServiceNowPasswordStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## ServiceNowPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `management_type` | MgmntTypeSelector | No |  |
| `optional_timeout` | CheckboxOneOf | Yes |  |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `servicenow_url` | string | Yes | Configure your ServiceNow URL here |
| `use_site_id_prefix` | CheckboxOneOf | Yes |  |
| `user_password` | ServiceNowPasswordSelector | Yes |  |
| `username` | string | Yes | Configure the user name here |

---

## SignL4ExplicitOrStoreSelector

---

## SignL4TeamSecret

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `secret` | string | Yes |  |

---

## SignL4TeamSecretStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## Signl4PluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `team_secret` | SignL4ExplicitOrStoreSelector | Yes |  |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |

---

## SiteConfigAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `basic_settings` | BasicSettingsAttributes | No |  |
| `configuration_connection` | ConfigurationConnectionAttributes1 | No |  |
| `secret` | string | No | The shared secret used by the central site to authenticate with the remote si... |
| `status_connection` | StatusConnectionAttributes1 | No |  |

---

## SiteConfigAttributesCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `basic_settings` | BasicSettingsAttributesCreate | Yes |  |
| `configuration_connection` | ConfigurationConnectionAttributes | Yes |  |
| `secret` | string | No | The shared secret used by the central site to authenticate with the remote si... |
| `status_connection` | StatusConnectionAttributes | Yes |  |

---

## SiteConfigAttributesUpdate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `basic_settings` | BasicSettingsAttributesUpdate | Yes |  |
| `configuration_connection` | ConfigurationConnectionAttributes | Yes |  |
| `secret` | string | No | The shared secret used by the central site to authenticate with the remote si... |
| `status_connection` | StatusConnectionAttributes | Yes |  |

---

## SiteConnectionRequestCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `site_config` | object | Yes | A site's connection. |

---

## SiteConnectionRequestUpdate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `site_config` | object | Yes | A site's connection. |

---

## SiteConnectionResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The configuration attributes of a site. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## SiteConnectionResponseCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []SiteConnectionResponse | No | A list of site configuration objects. |

---

## SiteLoginRequest

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `password` | string | Yes | The password for the username given |
| `username` | string | Yes | An administrative user's username. |

---

## SlackPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |
| `webhook_url` | SlackStoreOrExplicitURLSelector | Yes |  |

---

## SlackStoreOrExplicitURLSelector

---

## SlackWebhookStore

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## SlackWebhookURL

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `url` | string | Yes |  |

---

## SocketAttributes

---

## SocketAttributes1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `encrypted` | boolean | No | To enable an encrypted connection. |
| `host` | string | No | The IP or domain name of the host. |
| `path` | string | No | When the connection name is unix, this is the path to the unix socket. |
| `port` | integer | No | The TCP port to connect to. |
| `socket_type` | string | Yes | The connection name. This can be tcp, tcp6, unix or local. |
| `verify` | boolean | No | Verify server certificate. |

---

## SocketIP4

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `encrypted` | boolean | Yes | To enable an encrypted connection. |
| `host` | string | Yes | The IP or domain name of the host. |
| `port` | integer | Yes | The TCP port to connect to. |
| `socket_type` | string | Yes | The connection name. This can be tcp, tcp6, unix or local. |
| `verify` | boolean | No | Verify server certificate. |

---

## SocketIP6

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `encrypted` | boolean | Yes | To enable an encrypted connection. |
| `host` | string | Yes | The IP or domain name of the host. |
| `port` | integer | Yes | The TCP port to connect to. |
| `socket_type` | string | Yes | The connection name. This can be tcp, tcp6, unix or local. |
| `verify` | boolean | No | Verify server certificate. |

---

## SocketType

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `socket_type` | string | Yes | The connection name. This can be tcp, tcp6, unix or local. |

---

## SocketUnixAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `path` | string | Yes | When the connection name is unix, this is the path to the unix socket. |
| `socket_type` | string | Yes | The connection name. This can be tcp, tcp6, unix or local. |

---

## SpectrumPluginBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `base_oid` | string | Yes | The base OID for the trap content |
| `destination_ip` | string | Yes | IP Address of the Spectrum server receiving the SNMP trap |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `snmp_community` | string | Yes | SNMP Community for the SNMP trap. The password entered here is stored in plai... |

---

## SplunkRESTEndpointSelector

---

## SplunkStoreID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `store_id` | string | Yes | A password store ID |

---

## SplunkURLExplicit

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `option` | string | Yes |  |
| `url` | string | Yes | A valid splunk webhook URL |

---

## StatusConnectionAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `connect_timeout` | integer | Yes | The time that the GUI waits for a connection to the site to be established be... |
| `connection` | object | Yes | When connecting to remote site please make sure that Livestatus over TCP is a... |
| `disable_in_status_gui` | boolean | No | If you disable a connection, then no data of this site will be shown in the s... |
| `persistent_connection` | boolean | No | If you enable persistent connections then Multisite will try to keep open the... |
| `proxy` | object | Yes | The Livestatus proxy daemon configuration attributes. |
| `status_host` | object | Yes | By specifying a status host for each non-local connection you prevent Multisi... |
| `url_prefix` | string | No | The URL prefix will be prepended to links of addons like NagVis when a link t... |

---

## StatusConnectionAttributes1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `connect_timeout` | integer | Yes | The time that the GUI waits for a connection to the site to be established be... |
| `connection` | object | Yes | When connecting to remote site please make sure that Livestatus over TCP is a... |
| `disable_in_status_gui` | boolean | No | If you disable a connection, then no data of this site will be shown in the s... |
| `persistent_connection` | boolean | No | If you enable persistent connections then Multisite will try to keep open the... |
| `proxy` | object | Yes | The Livestatus proxy daemon configuration attributes. |
| `status_host` | object | No | By specifying a status host for each non-local connection you prevent Multisi... |
| `url_prefix` | string | No | The URL prefix will be prepended to links of addons like NagVis when a link t... |

---

## StatusHostAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host` | string | No | The host name of the status host. |
| `site` | string | No | The site ID of the status host. |
| `status_host_set` | string | Yes | enabled for 'use the following status host' and disabled for 'no status host'. |

---

## StatusHostAttributesBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `status_host_set` | string | Yes | enabled for 'use the following status host' and disabled for 'no status host' |

---

## StatusHostAttributesSet

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `host` | string | Yes | The host name of the status host. |
| `site` | string | Yes | The site ID of the status host. |
| `status_host_set` | string | Yes | enabled for 'use the following status host' and disabled for 'no status host' |

---

## StatusHostSet

---

## SysLogToFromPriorities

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `from_priority` | string | Yes |  |
| `to_priority` | string | Yes |  |

---

## TagCondition

---

## TagConditionConditionSchemaBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | No | The name of the tag. |
| `operator` | string | No | If the matched tag should be one of the given values, or not. |
| `value` | []string | No | A list of values for the tag. |

---

## TagConditionScalarSchemaBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `key` | string | No | The name of the tag. |
| `operator` | string | No | If the tag's value should match what is given under the field `value`. |
| `value` | string | No | The value of a tag. |

---

## TagGroupAttributes

---

## ThrottlePeriodicNotifications

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `beginning_from` | integer | Yes | Beginning notification number |
| `send_every_nth_notification` | integer | Yes | The rate then you will receive the notification 1 through 10 and then 15, 20,... |

---

## TimeAllowedRange

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end` | string | Yes | The end time of day. Inclusive. Use ISO8601 format. Seconds are stripped. |
| `start` | string | Yes | The start time of day. Inclusive. Use ISO8601 format. Seconds are stripped. |

---

## TimePeriodAttrsResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `active_time_ranges` | []ConcreteTimeRangeActive | No | The days for which time ranges were specified |
| `alias` | string | No | The alias of the time period |
| `exceptions` | []ConcreteTimePeriodException | No | Specific day exclusions with their list of time ranges |
| `exclude` | []string | No | The collection of time period aliases whose periods are excluded |

---

## TimePeriodException

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `date` | string | Yes | The date of the time period exception.8601 profile |
| `time_ranges` | []TimeRange1 | No |  |

---

## TimePeriodResponse

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | The time period attributes. |
| `id` | string | No | The unique identifier for this time period. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | The time period name. |

---

## TimePeriodResponseCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []TimePeriodResponse | No | A list of time period objects. |

---

## TimeRange

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end` | string | Yes | The approximate time of the last sample. |
| `start` | string | Yes | The approximate time of the first sample. |

---

## TimeRange1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `end` | string | Yes | The end time of the period's time range |
| `start` | string | Yes | The start time of the period's time range |

---

## TimeRangeActive

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `day` | string | No | The day for which time ranges are to be specified. The 'all' option allows to... |
| `time_ranges` | []TimeRange1 | No |  |

---

## TranslateNames

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `convert_case` | string | No | Convert all detected hostnames to upper- or lower-case. * `nop` - Do not conv... |
| `drop_domain` | boolean | No | Drop the rest of the domain, only keep the hostname. Will not affect IP addre... |
| `hostname_replacement` | []DirectMapping | No | Replace one value with another. These will be executed **after**: * `convert_... |
| `regexp_rewrites` | []RegexpRewrites | No | Rewrite discovered hostnames with multiple regular expressions. The replaceme... |

---

## UpdateAndAcknowledeEventSiteIDRequired

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `change_comment` | string | No | Event comment. |
| `change_contact` | string | No | Contact information. |
| `phase` | string | No | To change the phase of an event |
| `site_id` | string | Yes | An existing site id |

---

## UpdateAndAcknowledgeFilter

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `change_comment` | string | No | Event comment. |
| `change_contact` | string | No | Contact information. |
| `filter_type` | string | Yes | The way you would like to filter events. |
| `phase` | string | No | To change the phase of an event |
| `site_id` | string | No | An existing site id |

---

## UpdateAndAcknowledgeSelector

---

## UpdateAndAcknowledgeWithParams

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `change_comment` | string | No | Event comment. |
| `change_contact` | string | No | Contact information. |
| `filter_type` | string | Yes | The way you would like to filter events. |
| `filters` | FilterParamsUpdateAndAcknowledge | Yes |  |
| `phase` | string | No | To change the phase of an event |
| `site_id` | string | No | An existing site id |

---

## UpdateAndAcknowledgeWithQuery

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `change_comment` | string | No | Event comment. |
| `change_contact` | string | No | Contact information. |
| `filter_type` | string | Yes | The way you would like to filter events. |
| `phase` | string | No | To change the phase of an event |
| `query` | object | Yes | An query expression of the Livestatus 'eventconsoleevents' table in nested di... |
| `site_id` | string | No | An existing site id |

---

## UpdateContactGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | UpdateGroup | No |  |
| `name` | string | Yes | The name of the contact group. |

---

## UpdateDiscoveryPhase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `check_type` | string | Yes | The name of the check which this service uses. |
| `service_item` | string | Yes | The value uniquely identifying the service on a given host. |
| `target_phase` | string | Yes | The target phase of the service. |

---

## UpdateFolder

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Replace all attributes with the ones given in this field. Already setattribut... |
| `remove_attributes` | []string | No | A list of attributes which should be removed. Can't be used together with att... |
| `title` | string | No | The title of the folder. Used in the GUI. |
| `update_attributes` | oneOf | No | Only set the attributes which are given in this field. Already set attributes... |

---

## UpdateFolderEntry

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Replace all attributes with the ones given in this field. Already setattribut... |
| `folder` | string | Yes | The path name of the folder. Path delimiters can be either `~`, `/` or `\`. P... |
| `remove_attributes` | []string | No | A list of attributes which should be removed. Can't be used together with att... |
| `title` | string | No | The title of the folder. Used in the GUI. |
| `update_attributes` | oneOf | No | Only set the attributes which are given in this field. Already set attributes... |

---

## UpdateGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The name used for displaying in the GUI. |

---

## UpdateGroup1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The name used for displaying in the GUI. |

---

## UpdateGroup2

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The name used for displaying in the GUI. |

---

## UpdateHost

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Replace all currently set attributes on the host, with these attributes. Any ... |
| `remove_attributes` | []string | No | A list of attributes which should be removed. Can't be used together with att... |
| `update_attributes` | oneOf | No | Just update the hosts attributes with these attributes. The previously set at... |

---

## UpdateHostEntry

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | oneOf | No | Replace all currently set attributes on the host, with these attributes. Any ... |
| `host_name` | string | Yes | The hostname or IP address itself. |
| `remove_attributes` | []string | No | A list of attributes which should be removed. Can't be used together with att... |
| `update_attributes` | oneOf | No | Just update the hosts attributes with these attributes. The previously set at... |

---

## UpdateHostGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | UpdateGroup1 | No |  |
| `name` | string | Yes | The name of the host group. |

---

## UpdateHostTagGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `help` | string | No | A help description for the tag group |
| `repair` | boolean | No | The host tag group can be in use by other hosts. Setting repair to True gives... |
| `tags` | []HostTag | No | A list of host tags belonging to the host tag group |
| `title` | string | No | A title for the host tag |
| `topic` | string | No | Different tags can be grouped in a topic |

---

## UpdateNodes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `nodes` | []string | Yes | Nodes where the newly created host should be the cluster-container of. |

---

## UpdatePassword

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `comment` | string | No | A comment for the password |
| `documentation_url` | string | No | An optional URL pointing to documentation or any other page. You can use eith... |
| `owner` | string | No | Each password is owned by a group of users which are able to edit, delete and... |
| `password` | string | No | The password string |
| `shared` | []string | No | The list of members to share the password with |
| `title` | string | No | A title for the password |

---

## UpdateServiceGroup

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `attributes` | UpdateGroup2 | No |  |
| `name` | string | Yes | The name of the service group. |

---

## UpdateTimePeriod

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `active_time_ranges` | []TimeRangeActive | No | The list of active time ranges which replaces the existing list of time ranges |
| `alias` | string | No | An alias for the time period |
| `exceptions` | []TimePeriodException | No | A list of additional time ranges to be added. |
| `exclude` | []string | No | A list of time period aliases whose periods are excluded. |

---

## UpdateUser

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `auth_option` | object | No | Authentication option for the user |
| `authorized_sites` | []string | No | The names of the sites the user is authorized to handle. Specifying 'all' wil... |
| `contact_options` | object | No | Contact settings for the user |
| `contactgroups` | []string | No | Assign the user to one or multiple contact groups. If no contact group is spe... |
| `disable_login` | boolean | No | The user can be blocked from login but will remain part of the site. The disa... |
| `disable_notifications` | object | No |  |
| `fullname` | string | No | The alias or full name of the user |
| `idle_timeout` | object | No | Idle timeout for the user |
| `interface_options` | object | No |  |
| `language` | string | No | Configure the language to be used by the user in the user interface. Omitting... |
| `pager_address` | string | No |  |
| `roles` | []string | No | The list of assigned roles to the user |
| `temperature_unit` | string | No | Configure the temperature unit used for graphs and perfometers. |

---

## UseLiveStatusDaemon

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `use_livestatus_daemon` | string | Yes | Use livestatus daemon with direct connection or with livestatus proxy. |

---

## UserCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []UserObject | No | A list of user objects. |

---

## UserContactOption

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `email` | string | Yes | The mail address of the user. Required if the user is a monitoringcontact and... |
| `fallback_contact` | boolean | No | In case none of your notification rules handles a certain event a notificatio... |

---

## UserIdleOption

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `duration` | integer | No | The duration in seconds of the individual idle timeout if individual is selec... |
| `option` | string | Yes | This field indicates if the idle timeout uses the global configuration, is di... |

---

## UserInterfaceAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `interface_theme` | string | No | The theme of the interface |
| `mega_menu_icons` | string | No | This option decides if colored icon should be shown foe every entry in the me... |
| `navigation_bar_icons` | string | No | This option decides if icons in the navigation bar should show/hide the respe... |
| `show_mode` | string | No | This option decides what show mode should be used for unvisited menus. Altern... |
| `sidebar_position` | string | No | The position of the sidebar |

---

## UserInterfaceUpdateAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `interface_theme` | string | No | The theme of the interface |
| `mega_menu_icons` | string | No | This option decides if colored icon should be shown foe every entry in the me... |
| `navigation_bar_icons` | string | No | This option decides if icons in the navigation bar should show/hide the respe... |
| `show_mode` | string | No | This option decides what show mode should be used for unvisited menus. Altern... |
| `sidebar_position` | string | No | The position of the sidebar |

---

## UserObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | oneOf | No | The attributes of the user |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## UserRoleAttributes

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `alias` | string | Yes | The alias of the user role. |
| `basedon` | string | No | The builtin user role id that the user role is based on. |
| `builtin` | boolean | Yes | True if it's a builtin user role, otherwise False. |
| `permissions` | []string | Yes | A list of permissions for the user role. |

---

## UserRoleCollection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the objects in the collection. |
| `extensions` | object | No | Additional attributes alongside the collection. |
| `id` | string | No | The name of this collection. |
| `links` | []Link | Yes | list of links to other resources. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |
| `value` | []UserRoleObject | No | A list of user role objects. |

---

## UserRoleObject

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `domainType` | object | No | The domain type of the object. |
| `extensions` | object | No | All the attributes of a user role. |
| `id` | string | No | The unique identifier for this domain-object type. |
| `links` | []Link | Yes | list of links to other resources. |
| `members` | object | No | The container for external resources, like linked foreign objects or actions. |
| `title` | string | No | A human readable title of this object. Can be used for user interfaces. |

---

## UserSyncAttributes

---

## UserSyncAttributes1

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ldap_connections` | []string | No | A list of ldap connections. |
| `sync_with_ldap_connections` | string | Yes | Sync with ldap connections. The options are ldap, all, disabled. |

---

## UserSyncBase

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `sync_with_ldap_connections` | string | Yes | Sync with ldap connections. The options are ldap, all, disabled. |

---

## UserSyncWithLdapConnection

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `ldap_connections` | []string | Yes | A list of ldap connections. |
| `sync_with_ldap_connections` | string | Yes | Sync with ldap connections. The options are ldap, all, disabled. |

---

## VictoropsPluginCreate

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `disable_ssl_cert_verification` | object | Yes | Ignore unverified HTTPS request warnings. Use with caution. |
| `http_proxy` | object | Yes | Use the proxy settings from the environment variables. The variables NO_PROXY... |
| `plugin_name` | string | Yes | The plugin name. Built-in plugins only. |
| `splunk_on_call_rest_endpoint` | SplunkRESTEndpointSelector | Yes |  |
| `url_prefix_for_links_to_checkmk` | CheckboxOneOf | Yes |  |

---

## WhenToBulk

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `params` | NotificationBulking | Yes |  |
| `when_to_bulk` | string | Yes | Bulking can always happen or during a set time period |

---

## X509PEM

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `cert` | string | Yes | PEM-encoded X.509 certificate. |

---

## X509ReqPEMUUID

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `csr` | string | Yes | PEM-encoded X.509 CSR. The CN must a valid version-4 UUID. |

---

