// NOTE: This file was auto generated.  DO NOT EDIT DIRECTLY!
// To change feature flags, edit:
//  pkg/services/featuremgmt/registry.go
// Then run tests in:
//  pkg/services/featuremgmt/toggles_gen_test.go

/**
 * Describes available feature toggles in Grafana. These can be configured via
 * conf/custom.ini to enable features under development or not yet available in
 * stable version.
 *
 * Only enabled values will be returned in this interface.
 *
 * NOTE: the possible values may change between versions without notice, although
 * this may cause compilation issues when depending on removed feature keys, the
 * runtime state will continue to work.
 *
 * @public
 */
export interface FeatureToggles {
  disableEnvelopeEncryption?: boolean;
  ['live-service-web-worker']?: boolean;
  queryOverLive?: boolean;
  panelTitleSearch?: boolean;
  publicDashboards?: boolean;
  publicDashboardsEmailSharing?: boolean;
  publicDashboardsScene?: boolean;
  lokiExperimentalStreaming?: boolean;
  featureHighlights?: boolean;
  storage?: boolean;
  correlations?: boolean;
  autoMigrateOldPanels?: boolean;
  autoMigrateGraphPanel?: boolean;
  autoMigrateTablePanel?: boolean;
  autoMigratePiechartPanel?: boolean;
  autoMigrateWorldmapPanel?: boolean;
  autoMigrateStatPanel?: boolean;
  autoMigrateXYChartPanel?: boolean;
  disableAngular?: boolean;
  canvasPanelNesting?: boolean;
  vizActions?: boolean;
  disableSecretsCompatibility?: boolean;
  logRequestsInstrumentedAsUnknown?: boolean;
  grpcServer?: boolean;
  cloudWatchCrossAccountQuerying?: boolean;
  showDashboardValidationWarnings?: boolean;
  mysqlAnsiQuotes?: boolean;
  mysqlParseTime?: boolean;
  accessControlOnCall?: boolean;
  nestedFolders?: boolean;
  alertingBacktesting?: boolean;
  editPanelCSVDragAndDrop?: boolean;
  alertingNoNormalState?: boolean;
  logsContextDatasourceUi?: boolean;
  lokiQuerySplitting?: boolean;
  lokiQuerySplittingConfig?: boolean;
  individualCookiePreferences?: boolean;
  prometheusMetricEncyclopedia?: boolean;
  influxdbBackendMigration?: boolean;
  influxqlStreamingParser?: boolean;
  influxdbRunQueriesInParallel?: boolean;
  prometheusRunQueriesInParallel?: boolean;
  lokiMetricDataplane?: boolean;
  lokiLogsDataplane?: boolean;
  dataplaneFrontendFallback?: boolean;
  disableSSEDataplane?: boolean;
  alertStateHistoryLokiSecondary?: boolean;
  alertStateHistoryLokiPrimary?: boolean;
  alertStateHistoryLokiOnly?: boolean;
  unifiedRequestLog?: boolean;
  renderAuthJWT?: boolean;
  refactorVariablesTimeRange?: boolean;
  faroDatasourceSelector?: boolean;
  enableDatagridEditing?: boolean;
  extraThemes?: boolean;
  lokiPredefinedOperations?: boolean;
  pluginsFrontendSandbox?: boolean;
  frontendSandboxMonitorOnly?: boolean;
  pluginsDetailsRightPanel?: boolean;
  sqlDatasourceDatabaseSelection?: boolean;
  recordedQueriesMulti?: boolean;
  vizAndWidgetSplit?: boolean;
  logsExploreTableVisualisation?: boolean;
  awsDatasourcesTempCredentials?: boolean;
  transformationsRedesign?: boolean;
  mlExpressions?: boolean;
  traceQLStreaming?: boolean;
  metricsSummary?: boolean;
  datasourceAPIServers?: boolean;
  grafanaAPIServerWithExperimentalAPIs?: boolean;
  grafanaAPIServerEnsureKubectlAccess?: boolean;
  featureToggleAdminPage?: boolean;
  awsAsyncQueryCaching?: boolean;
  permissionsFilterRemoveSubquery?: boolean;
  prometheusConfigOverhaulAuth?: boolean;
  configurableSchedulerTick?: boolean;
  alertingNoDataErrorExecution?: boolean;
  angularDeprecationUI?: boolean;
  dashgpt?: boolean;
  aiGeneratedDashboardChanges?: boolean;
  reportingRetries?: boolean;
  sseGroupByDatasource?: boolean;
  libraryPanelRBAC?: boolean;
  lokiRunQueriesInParallel?: boolean;
  wargamesTesting?: boolean;
  alertingInsights?: boolean;
  externalCorePlugins?: boolean;
  pluginsAPIMetrics?: boolean;
  externalServiceAccounts?: boolean;
  panelMonitoring?: boolean;
  enableNativeHTTPHistogram?: boolean;
  disableClassicHTTPHistogram?: boolean;
  formatString?: boolean;
  transformationsVariableSupport?: boolean;
  kubernetesPlaylists?: boolean;
  kubernetesSnapshots?: boolean;
  kubernetesDashboards?: boolean;
  kubernetesDashboardsAPI?: boolean;
  kubernetesFolders?: boolean;
  grafanaAPIServerTestingWithExperimentalAPIs?: boolean;
  datasourceQueryTypes?: boolean;
  queryService?: boolean;
  queryServiceRewrite?: boolean;
  queryServiceFromUI?: boolean;
  cloudWatchBatchQueries?: boolean;
  recoveryThreshold?: boolean;
  lokiStructuredMetadata?: boolean;
  teamHttpHeaders?: boolean;
  cachingOptimizeSerializationMemoryUsage?: boolean;
  panelTitleSearchInV1?: boolean;
  managedPluginsInstall?: boolean;
  prometheusPromQAIL?: boolean;
  prometheusCodeModeMetricNamesSearch?: boolean;
  addFieldFromCalculationStatFunctions?: boolean;
  alertmanagerRemoteSecondary?: boolean;
  alertmanagerRemotePrimary?: boolean;
  alertmanagerRemoteOnly?: boolean;
  annotationPermissionUpdate?: boolean;
  extractFieldsNameDeduplication?: boolean;
  dashboardSceneForViewers?: boolean;
  dashboardSceneSolo?: boolean;
  dashboardScene?: boolean;
  panelFilterVariable?: boolean;
  pdfTables?: boolean;
  ssoSettingsApi?: boolean;
  canvasPanelPanZoom?: boolean;
  logsInfiniteScrolling?: boolean;
  exploreMetrics?: boolean;
  alertingSimplifiedRouting?: boolean;
  logRowsPopoverMenu?: boolean;
  pluginsSkipHostEnvVars?: boolean;
  tableSharedCrosshair?: boolean;
  regressionTransformation?: boolean;
  lokiQueryHints?: boolean;
  kubernetesFeatureToggles?: boolean;
  cloudRBACRoles?: boolean;
  alertingQueryOptimization?: boolean;
  newFolderPicker?: boolean;
  jitterAlertRulesWithinGroups?: boolean;
  onPremToCloudMigrations?: boolean;
  onPremToCloudMigrationsAlerts?: boolean;
  alertingSaveStatePeriodic?: boolean;
  promQLScope?: boolean;
  sqlExpressions?: boolean;
  nodeGraphDotLayout?: boolean;
  groupToNestedTableTransformation?: boolean;
  newPDFRendering?: boolean;
  tlsMemcached?: boolean;
  kubernetesAggregator?: boolean;
  expressionParser?: boolean;
  groupByVariable?: boolean;
  authAPIAccessTokenAuth?: boolean;
  scopeFilters?: boolean;
  ssoSettingsSAML?: boolean;
  oauthRequireSubClaim?: boolean;
  newDashboardWithFiltersAndGroupBy?: boolean;
  cloudWatchNewLabelParsing?: boolean;
  accessActionSets?: boolean;
  disableNumericMetricsSortingInExpressions?: boolean;
  grafanaManagedRecordingRules?: boolean;
  queryLibrary?: boolean;
  logsExploreTableDefaultVisualization?: boolean;
  newDashboardSharingComponent?: boolean;
  alertingListViewV2?: boolean;
  notificationBanner?: boolean;
  dashboardRestore?: boolean;
  datasourceProxyDisableRBAC?: boolean;
  alertingDisableSendAlertsExternal?: boolean;
  preserveDashboardStateWhenNavigating?: boolean;
  alertingCentralAlertHistory?: boolean;
  pluginProxyPreserveTrailingSlash?: boolean;
  azureMonitorPrometheusExemplars?: boolean;
  pinNavItems?: boolean;
  authZGRPCServer?: boolean;
  openSearchBackendFlowEnabled?: boolean;
  ssoSettingsLDAP?: boolean;
  failWrongDSUID?: boolean;
  zanzana?: boolean;
  passScopeToDashboardApi?: boolean;
  alertingApiServer?: boolean;
  cloudWatchRoundUpEndTime?: boolean;
  cloudwatchMetricInsightsCrossAccount?: boolean;
  prometheusAzureOverrideAudience?: boolean;
  alertingFilterV2?: boolean;
  dataplaneAggregator?: boolean;
  newFiltersUI?: boolean;
  tableNG?: boolean;
  lokiSendDashboardPanelNames?: boolean;
  alertingPrometheusRulesPrimary?: boolean;
  singleTopNav?: boolean;
  exploreLogsShardSplitting?: boolean;
  exploreLogsAggregatedMetrics?: boolean;
  exploreLogsLimitedTimeRange?: boolean;
  homeSetupGuide?: boolean;
  appPlatformGrpcClientAuth?: boolean;
  appSidecar?: boolean;
  groupAttributeSync?: boolean;
  alertingQueryAndExpressionsStepMode?: boolean;
  improvedExternalSessionHandling?: boolean;
  useSessionStorageForRedirection?: boolean;
  rolePickerDrawer?: boolean;
  unifiedStorageSearch?: boolean;
  pluginsSriChecks?: boolean;
  unifiedStorageBigObjectsSupport?: boolean;
}
