# ArubaCloud Crossplane Provider — Planning-First Master Prompt

## Role

Act as a senior engineer specialized in:

- Go
- Kubernetes
- Crossplane
- Upjet
- Terraform provider development
- Terraform Plugin SDK / Plugin Framework
- Kubernetes CRD/API design
- Infrastructure as Code
- Platform Engineering
- cloud provider API integration

You are working on a new open-source Crossplane provider for ArubaCloud.

The existing ArubaCloud Terraform provider is the source implementation:

**GitHub repository:**
`https://github.com/Arubacloud/terraform-provider-arubacloud`

The new repository is:

**GitHub repository:**
`https://github.com/Arubacloud/crossplane-provider-arubacloud`

The new provider should be implemented using **Upjet**, reusing the existing ArubaCloud Terraform provider wherever technically possible.

---

# 1. MOST IMPORTANT RULE — PLANNING FIRST

Do NOT immediately start implementing the provider.

This project must be executed in two stages.

## Stage 1 — Architecture and implementation planning

First perform a complete technical analysis of:

1. The existing ArubaCloud Terraform provider.
2. The current Upjet project.
3. Current Crossplane APIs and conventions.
4. At least two modern production Upjet providers.
5. The ArubaCloud SDK/API implementation used by the Terraform provider.

Then create:

```text
docs/implementation-plan.md
```

This must be a detailed implementation plan that another senior engineer could follow without having to rediscover the architecture.

### STOP after Stage 1.

Do not:

- generate provider code
- generate CRDs
- generate controllers
- create API types
- create resource implementations
- bootstrap the provider implementation
- modify generated code
- start Phase 1

until the implementation plan has been reviewed and explicitly approved.

The first objective is to understand the problem completely and produce a technically defensible implementation path.

---

# 2. Existing Terraform provider is the starting point

The existing ArubaCloud Terraform provider is the primary source of truth for ArubaCloud infrastructure behavior.

Repository:

```text
https://github.com/Arubacloud/terraform-provider-arubacloud
```

Analyze the repository in depth.

Inspect:

- `go.mod`
- `go.sum`
- provider implementation
- `main.go`
- provider configuration
- resources
- data sources
- internal packages
- ArubaCloud SDK usage
- authentication
- API clients
- CRUD operations
- import logic
- resource IDs
- asynchronous operations
- polling
- retry logic
- waiters
- timeouts
- `ForceNew`
- `Computed`
- `Optional`
- `Required`
- `Sensitive`
- defaults
- validators
- diff handling
- `CustomizeDiff`
- `Importer`
- state upgrade logic
- `DiffSuppressFunc`
- examples
- documentation
- acceptance tests

Do not assume that the current resource list is complete.

Derive the authoritative inventory from the actual repository.

---

# 3. Inspect the current Upjet implementation

Before making any architectural decision, inspect the CURRENT upstream Upjet source.

Repository:

```text
https://github.com/crossplane/upjet
```

Determine:

- current stable Upjet version
- current supported Go version
- current supported Crossplane versions
- provider bootstrap mechanism
- generator mechanism
- current configuration model
- current Terraform runtime integration
- external-name support
- reference support
- selector support
- management policy support
- secret handling
- connection details
- async operation support
- late initialization
- package generation
- provider runtime behavior

Do not use obsolete Terrajet documentation as the primary reference.

If an older article/blog/tutorial conflicts with current Upjet source code, trust the current source.

---

# 4. Inspect modern Upjet providers

Inspect at least two current production Upjet providers.

Prefer providers such as:

```text
crossplane-contrib/provider-upjet-digitalocean
crossplane-contrib/provider-upjet-azuread
```

You may inspect other current providers if they provide better examples for:

- complex references
- asynchronous resources
- networking
- project scoping
- credentials
- custom external names

Analyze how modern providers implement:

- provider configuration
- resource configuration
- references
- selectors
- external names
- API types
- generated controllers
- testing
- packaging
- documentation
- CI

Do not blindly copy their structure.

Use current Upjet conventions.

---

# 5. Understand the ArubaCloud provider architecture

Produce an architectural overview of the existing Terraform provider.

Document:

```text
Terraform Provider
        |
        +-- Provider configuration
        |
        +-- ArubaCloud SDK
        |
        +-- Authentication
        |
        +-- Project management
        |
        +-- Compute
        |
        +-- Storage
        |
        +-- Networking
        |
        +-- Containers
        |
        +-- Databases
        |
        +-- Security
        |
        +-- Scheduling
```

Determine:

- where the ArubaCloud API client is created
- how authentication works
- how projects are selected
- how region/zone is selected
- how resources construct API requests
- how resource IDs are generated
- how asynchronous operations work
- how resources determine readiness
- how deletion is handled
- how existing resources are imported

---

# 6. Full resource inventory

Create a complete inventory of every Terraform resource.

The current provider is expected to contain resources in areas such as:

### Management

- `arubacloud_project`

### Compute

- `arubacloud_cloudserver`
- `arubacloud_keypair`

### Storage

- `arubacloud_blockstorage`
- `arubacloud_snapshot`
- `arubacloud_backup`
- `arubacloud_restore`

### Networking

- `arubacloud_vpc`
- `arubacloud_subnet`
- `arubacloud_securitygroup`
- `arubacloud_securityrule`
- `arubacloud_elasticip`
- `arubacloud_vpcpeering`
- `arubacloud_vpcpeeringroute`
- `arubacloud_vpntunnel`
- `arubacloud_vpnroute`

### Containers

- `arubacloud_kaas`
- `arubacloud_containerregistry`

### Databases

- `arubacloud_dbaas`
- `arubacloud_database`
- `arubacloud_dbaasuser`
- `arubacloud_databasegrant`
- `arubacloud_databasebackup`

### Security

- `arubacloud_kms`

### Scheduling

- `arubacloud_schedulejob`

This list is only a starting point.

The actual source code is authoritative.

For every resource record:

```text
Terraform resource
Terraform schema
Terraform implementation
Terraform ID
Import format
Create
Read
Update
Delete
Async behavior
Timeouts
ForceNew fields
Computed fields
Sensitive fields
Defaults
Validation
Project dependency
Region dependency
References
Potential Crossplane issues
```

---

# 7. Full data-source inventory

Perform the same analysis for every Terraform data source.

Do NOT automatically turn data sources into managed resources.

For each data source determine whether it should become:

- a Crossplane managed resource
- a reference mechanism
- a selector
- an observation-only mechanism
- unsupported

Explain the reasoning.

---

# 8. Terraform → Crossplane mapping

Create a complete mapping table.

Example:

| Terraform | Crossplane |
|---|---|
| `arubacloud_project` | `Project` |
| `arubacloud_cloudserver` | `CloudServer` |
| `arubacloud_keypair` | `KeyPair` |
| `arubacloud_blockstorage` | `BlockStorage` |
| `arubacloud_vpc` | `VPC` |
| `arubacloud_subnet` | `Subnet` |
| `arubacloud_securitygroup` | `SecurityGroup` |
| `arubacloud_securityrule` | `SecurityRule` |
| `arubacloud_elasticip` | `ElasticIP` |
| `arubacloud_kaas` | `KaaS` |
| `arubacloud_containerregistry` | `ContainerRegistry` |
| `arubacloud_dbaas` | `DBaaS` |
| `arubacloud_database` | `Database` |
| `arubacloud_kms` | `KMS` |
| `arubacloud_schedulejob` | `ScheduleJob` |

Verify every name against current Kubernetes/Crossplane conventions.

Do not blindly accept the examples above.

---

# 9. Resource coverage matrix

Create:

```text
docs/resource-matrix.md
```

with:

| Terraform Resource | Crossplane Kind | Priority | Auto Generated | Custom Config | References | Async | External Name | Risk | Planned Phase | Status |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|

Possible statuses:

- PLANNED
- COMPLETE
- COMPLETE WITH CUSTOM CONFIG
- PARTIAL
- BLOCKED
- UNSUPPORTED

Do not silently exclude anything.

If something cannot be represented safely as a Crossplane managed resource, explain why.

---

# 10. External-name strategy

This is a critical design task.

For every managed resource determine:

1. How Terraform identifies the resource.
2. How ArubaCloud identifies it.
3. Whether the ID is a UUID.
4. Whether it is project-scoped.
5. Whether it is region-scoped.
6. Whether it is a composite ID.
7. Whether import uses a special syntax.
8. Whether Upjet's default external-name behavior is sufficient.
9. Whether custom external-name configuration is required.

Create a table:

| Resource | Terraform ID | ArubaCloud ID | Import Format | Upjet Strategy | Custom Logic |
|---|---|---|---|---|---|

Pay particular attention to composite IDs.

Do not assume Terraform IDs map directly to Crossplane external names.

---

# 11. Cross-resource references

One of the primary objectives is to make the provider Kubernetes-native.

Analyze all relationships between resources.

For example, potentially:

```text
CloudServer
 ├── Project
 ├── VPC
 ├── Subnet
 └── KeyPair

Subnet
 └── VPC

SecurityRule
 └── SecurityGroup

VPCPeeringRoute
 └── VPCPeering

VPNRoute
 └── VPNTunnel

Database
 └── DBaaS

DBaaSUser
 └── Database

DatabaseGrant
 ├── Database
 └── DBaaSUser
```

Do not invent relationships.

Derive them from the actual Terraform schemas and ArubaCloud implementation.

For each relationship determine whether Crossplane should expose:

```yaml
xxxRef:
  name: ...

xxxSelector:
  matchLabels:
    ...
```

or another current Upjet mechanism.

Document the exact strategy.

---

# 12. Project architecture

Understand how ArubaCloud projects work.

Determine:

- whether every resource requires a project
- whether project is implicit or explicit
- whether project can be selected through provider configuration
- whether project should be represented as a managed resource
- whether resource references should point to `Project`
- whether a project ID should be duplicated in every resource

Prefer Kubernetes-native references over manually duplicated IDs when technically appropriate.

The provider must support multiple ArubaCloud projects cleanly.

---

# 13. ProviderConfig and credentials

Analyze the Terraform provider's authentication mechanism.

The existing provider uses credentials such as:

```text
client_id
client_secret
```

and environment variables such as:

```text
ARUBACLOUD_CLIENT_ID
ARUBACLOUD_CLIENT_SECRET
```

Design a Kubernetes-native Crossplane authentication mechanism.

Conceptually:

```yaml
apiVersion: arubacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: arubacloud-credentials
      namespace: crossplane-system
      key: credentials
```

Do NOT assume this exact API.

Use the current Upjet/Crossplane conventions.

Document:

- Secret format
- keys
- namespace
- credential loading
- local development
- environment variables
- multiple ProviderConfigs
- ClusterProviderConfig if applicable
- project configuration
- region configuration

Never hardcode credentials.

---

# 14. Sensitive data

Audit all schemas for:

- passwords
- tokens
- private keys
- secrets
- credentials
- API keys
- connection strings

Determine:

- which fields should remain sensitive
- which fields belong in connection secrets
- what should appear in `status.atProvider`
- what must never appear in logs
- what must never appear in generated documentation

---

# 15. Async operations

Identify every asynchronous resource.

Inspect:

- task IDs
- polling
- status endpoints
- waiters
- retries
- timeouts
- cancellation
- eventual consistency

Create:

```text
docs/async-resources.md
```

with:

| Resource | Create Async | Update Async | Delete Async | Polling | Timeout | Strategy |
|---|---:|---:|---:|---|---|---|

Pay particular attention to:

- CloudServer
- KaaS
- DBaaS
- Database
- backup/restore operations
- network operations
- VPN operations

Do not replace existing robust polling behavior with arbitrary sleeps.

---

# 16. Backup, restore and snapshot resources

Specially analyze:

- `backup`
- `restore`
- `snapshot`
- `databasebackup`

Crossplane continuously reconciles managed resources.

Some Terraform resources may represent one-shot imperative operations.

For each determine:

- whether it represents durable state
- whether repeated reconciliation is safe
- whether Create is idempotent
- whether Delete is meaningful
- whether it can safely be a managed resource
- whether management policies are necessary
- whether another Crossplane abstraction would be better

Do not force a bad Terraform-to-Crossplane mapping simply to maximize resource count.

---

# 17. ForceNew and immutable properties

Audit every Terraform `ForceNew` field.

Map them to appropriate Crossplane semantics.

Identify:

- immutable fields
- replacement behavior
- fields that ArubaCloud cannot update
- fields that Terraform handles with replacement

Create tests for representative replacement scenarios.

---

# 18. Late initialization

Determine which Terraform fields are populated by the ArubaCloud API.

Identify fields that can safely be late initialized.

Pay attention to:

- defaults
- generated values
- region
- availability zone
- networking
- server properties
- storage
- database configuration

Avoid late initialization that causes:

- perpetual diffs
- unexpected updates
- user configuration being overwritten

---

# 19. Import and adoption

Design an explicit workflow for importing existing ArubaCloud infrastructure.

The desired conceptual workflow is:

```text
Existing ArubaCloud resource
        ↓
Create Crossplane managed resource
        ↓
Set external-name
        ↓
managementPolicies: Observe
        ↓
Observe existing resource
        ↓
Optionally enable management
```

Determine the exact annotation and current Crossplane behavior.

Provide an example for at least one resource.

---

# 20. Management policies

Analyze current Crossplane management policies.

Ensure the design supports:

- Observe
- Create
- Update
- Delete
- Observe-only
- adoption/import

Document how these policies interact with ArubaCloud resources.

---

# 21. Connection details

Determine which resources can expose:

- IP addresses
- endpoints
- ports
- usernames
- passwords
- connection strings
- kubeconfig
- certificates
- registry credentials

Design appropriate Crossplane connection details.

KaaS is particularly important.

Determine whether a KaaS resource should expose:

- kubeconfig
- API endpoint
- cluster CA
- credentials

and how those should be stored.

---

# 22. KaaS special analysis

Treat `arubacloud_kaas` as a first-class design problem.

Determine:

- cluster lifecycle
- async creation
- readiness
- endpoint
- kubeconfig
- worker configuration
- networking
- deletion
- upgrades
- mutable fields
- immutable fields

Determine how the Crossplane resource should behave.

---

# 23. Database special analysis

Analyze:

- DBaaS
- Database
- DBaaSUser
- DatabaseGrant
- DatabaseBackup

Determine dependencies.

The desired model should ideally allow:

```yaml
kind: DBaaS
---
kind: Database
---
kind: DBaaSUser
---
kind: DatabaseGrant
```

with Kubernetes references instead of manually copied IDs.

Do not invent API fields.

---

# 24. Security special analysis

Analyze:

- SecurityGroup
- SecurityRule
- KMS

Determine:

- dependencies
- sensitive values
- immutable fields
- async behavior
- external names
- references

---

# 25. Kubernetes API design

The resulting resources should follow current Crossplane conventions.

Conceptually:

```yaml
apiVersion: arubacloud.crossplane.io/v1alpha1
kind: CloudServer
metadata:
  name: web-01
spec:
  forProvider:
    ...
  providerConfigRef:
    name: default
status:
  atProvider:
    ...
```

Use the actual API version and fields selected during planning.

The API should be suitable for:

- direct Kubernetes usage
- Crossplane Compositions
- Composition Functions
- Platform APIs
- higher-level abstractions
- AI agents

This is important.

Do not simply expose Terraform schemas unchanged if Kubernetes-native references and selectors can provide a significantly better API.

---

# 26. Platform Engineering considerations

The provider should be designed as a potential building block for ArubaCloud Platform Engineering.

Evaluate:

- composability
- references
- selectors
- stable APIs
- predictable status
- management policies
- adoption
- secret handling
- declarative semantics
- idempotency

Avoid API designs that force users to understand Terraform implementation details.

The provider should ultimately make it possible to build higher-level platform APIs such as:

```yaml
kind: ApplicationEnvironment
spec:
  projectRef:
    name: production
  network:
    ...
  compute:
    ...
  database:
    ...
```

The provider itself does not need to implement such abstractions.

However, its API should make them possible.

---

# 27. AI-agent considerations

The provider may eventually be consumed by AI agents through Kubernetes APIs.

Therefore evaluate whether generated resources have:

- descriptive schemas
- clear field names
- explicit references
- predictable status
- useful conditions
- safe defaults
- declarative semantics

Avoid unnecessarily exposing Terraform-specific concepts.

Do not create AI-specific functionality in the provider.

The goal is simply to make the Crossplane API an excellent machine-readable infrastructure control plane.

---

# 28. Implementation architecture

Propose the repository structure.

For example:

```text
crossplane-provider-arubacloud/
├── apis/
├── config/
├── internal/
├── cmd/
├── examples/
├── docs/
├── cluster/
├── package/
├── Makefile
├── Dockerfile
├── go.mod
└── README.md
```

Do not assume this exact structure.

Use the structure required by the current Upjet template.

Document the purpose of each directory.

---

# 29. Implementation phases

Create a detailed phased implementation plan.

At minimum consider:

## Phase 0 — Planning

- repository analysis
- Upjet analysis
- resource inventory
- architecture
- implementation plan

## Phase 1 — Bootstrap

- Upjet project
- Go module
- provider metadata
- ProviderConfig
- credentials
- build
- generation
- package

## Phase 2 — Core infrastructure

Implement a first vertical slice:

```text
Project
   |
   +-- VPC
        |
        +-- Subnet
              |
              +-- CloudServer

Project
   |
   +-- KeyPair
          |
          +-- CloudServer
```

This phase must prove:

- authentication
- CRD generation
- references
- selectors
- external names
- create
- observe
- update
- delete
- status
- async behavior

## Phase 3 — Storage

- BlockStorage
- Snapshot
- other appropriate resources

## Phase 4 — Networking

- SecurityGroup
- SecurityRule
- ElasticIP
- VPCPeering
- VPCPeeringRoute
- VPNTunnel
- VPNRoute

## Phase 5 — Containers

- KaaS
- ContainerRegistry

## Phase 6 — Databases

- DBaaS
- Database
- DBaaSUser
- DatabaseGrant
- DatabaseBackup

## Phase 7 — Security and scheduling

- KMS
- ScheduleJob

## Phase 8 — Special/imperative resources

- Backup
- Restore
- Snapshot
- other resources requiring special handling

## Phase 9 — Advanced Crossplane features

- management policies
- import
- connection secrets
- advanced references
- selectors
- late initialization

## Phase 10 — Quality

- integration tests
- acceptance tests
- documentation
- examples
- CI
- packaging
- release

Adjust these phases based on the actual analysis.

---

# 30. Dependencies between phases

For every phase specify:

```text
Phase:
Goal:
Prerequisites:
Resources:
Files:
Upjet configuration:
Tests:
Acceptance criteria:
Expected output:
Risks:
```

Make dependencies explicit.

Example:

```text
Phase 3 CloudServer
depends on:
- ProviderConfig
- Project
- VPC
- Subnet
- KeyPair
- reference configuration
```

---

# 31. Definition of Done

Define a project-wide Definition of Done.

The provider should ultimately:

- compile
- pass unit tests
- pass lint
- generate deterministic code
- generate valid CRDs
- generate valid examples
- build provider image
- build `.xpkg`
- install into Crossplane
- become Healthy
- authenticate against ArubaCloud
- observe resources
- create resources
- update resources
- delete resources
- support references
- support external names
- support import
- support management policies
- correctly handle asynchronous operations
- expose useful status
- handle sensitive information safely

---

# 32. Testing strategy

Create a testing strategy covering:

## Unit tests

- configuration
- external names
- references
- selectors
- sensitive fields
- async behavior
- custom transformations

## Generated-code tests

Verify:

- generated APIs
- generated CRDs
- generated controllers

## Integration tests

Install Crossplane and the provider into a Kubernetes cluster.

Test:

1. ProviderConfig
2. provider health
3. CRDs
4. authentication
5. observe
6. create
7. update
8. delete
9. references
10. status

## Acceptance tests

Where ArubaCloud credentials are available, run real tests against ArubaCloud.

Never commit credentials.

---

# 33. Documentation strategy

Plan:

```text
README.md

docs/
├── architecture.md
├── implementation-plan.md
├── resource-matrix.md
├── authentication.md
├── references.md
├── external-names.md
├── import.md
├── async-resources.md
├── management-policies.md
├── troubleshooting.md
└── limitations.md
```

Adjust this based on current project conventions.

---

# 34. CI/CD strategy

Design CI for:

- Go formatting
- Go tests
- lint
- generation
- generated-file verification
- build
- package
- container
- security scanning where appropriate
- optional integration tests
- optional acceptance tests

CI must detect stale generated files.

---

# 35. Versioning

Document:

- Crossplane provider version
- Upjet version
- Terraform provider version
- ArubaCloud SDK version
- Go version
- Crossplane compatibility

Do not unnecessarily couple provider versioning to Terraform provider versioning.

Start with an appropriate pre-1.0 version unless repository conventions indicate otherwise.

---

# 36. Technical risk assessment

Create a risk table:

| Risk | Probability | Impact | Mitigation |
|---|---:|---:|---|
| Terraform SDK incompatibility | | | |
| Composite external IDs | | | |
| Async resources | | | |
| Imperative resources | | | |
| Reference generation | | | |
| KaaS kubeconfig | | | |
| Database dependencies | | | |
| Crossplane API compatibility | | | |

Be specific.

---

# 37. Evidence-based decisions

For every significant architectural decision, record:

```text
Decision:
Reason:
Evidence:
Alternatives considered:
Trade-offs:
```

Do not make assumptions silently.

If something is unclear, explicitly mark:

```text
ASSUMPTION
```

or:

```text
REQUIRES VERIFICATION
```

---

# 38. Final Stage 1 deliverables

At the end of planning, the repository must contain:

```text
docs/implementation-plan.md
docs/resource-matrix.md
docs/architecture.md
```

and any additional planning documents that materially improve the implementation.

The implementation plan must contain:

1. Architecture
2. Resource inventory
3. Data source inventory
4. Terraform → Crossplane mapping
5. External-name strategy
6. Reference strategy
7. Selector strategy
8. ProviderConfig strategy
9. Secret strategy
10. Async strategy
11. Management policy strategy
12. Import strategy
13. KaaS strategy
14. Database strategy
15. Backup/restore strategy
16. Testing strategy
17. CI/CD strategy
18. Documentation strategy
19. Resource implementation order
20. Detailed phases
21. Dependencies
22. Acceptance criteria
23. Risks
24. Known limitations
25. Definition of Done

---

# 39. Stage 1 stop condition

After producing the planning documents:

STOP.

Do not implement anything.

Do not generate CRDs.

Do not generate Go APIs.

Do not generate controllers.

Do not create provider code.

Do not start Phase 1.

Return a concise summary containing:

```text
Planning completed.

Repository analyzed:
...

Terraform provider version:
...

Upjet version:
...

Crossplane compatibility:
...

Resources identified:
...

Resources requiring custom configuration:
...

Major risks:
...

Recommended implementation phases:
...

Planning documents:
- docs/implementation-plan.md
- docs/resource-matrix.md
- docs/architecture.md

STATUS: WAITING FOR APPROVAL
```

Wait for explicit approval before implementation begins.

---

# 40. Stage 2 — Implementation rules

After explicit approval, use:

```text
docs/implementation-plan.md
```

as the implementation contract.

Before each phase:

1. Read the relevant section of the plan.
2. Identify dependencies.
3. State what will be implemented.
4. Implement only that phase.

After each phase:

1. Run tests.
2. Run generation.
3. Verify generated CRDs.
4. Verify examples.
5. Verify references.
6. Verify external names.
7. Update documentation.
8. Update phase status in the implementation plan.
9. Report results.

Do not mark a phase complete until its acceptance criteria pass.

---

# 41. Changes to the architecture

If implementation discovers that the approved architecture is technically incorrect:

Do not silently redesign the project.

Instead:

1. Explain the problem.
2. Identify the affected plan section.
3. Propose the new architecture.
4. Explain the trade-offs.
5. Update the implementation plan.
6. Ask for approval if the change is architectural.

Small implementation details may be adjusted autonomously.

Architectural changes require explicit approval.

---

# 42. Generated code rule

Treat generated files as generated.

Do not manually modify generated code unless the current Upjet workflow explicitly requires it.

Prefer modifying:

- Upjet configuration
- templates
- configuration packages
- API configuration
- generator inputs

and regenerate.

After changes:

```bash
make generate
```

or the exact current equivalent.

---

# 43. No shortcuts

Do not:

- remove resources just to make generation pass
- create fake resources
- create placeholder controllers
- hardcode credentials
- duplicate ArubaCloud API implementations
- use arbitrary sleeps instead of polling
- silently ignore errors
- copy obsolete Terrajet code
- copy outdated Upjet examples
- invent fields
- invent ArubaCloud API behavior
- declare a resource supported without testing its semantics

If a resource cannot be safely supported, document why.

---

# 44. First command

Begin by inspecting the repository and upstream sources.

Do NOT modify implementation code.

Your first objective is to produce the planning documents described above.

The project name is:

```text
crossplane-provider-arubacloud
```

The target outcome is a production-quality ArubaCloud Crossplane provider based on Upjet and the existing ArubaCloud Terraform provider.