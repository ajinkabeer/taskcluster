// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package scheduler

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Information about a **task-graph** as known by the scheduler, with all the state of all individual tasks.
	//
	// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#
	InspectTaskGraphResponse struct {

		// Required task metadata
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/metadata
		Metadata struct {

			// Human readable description of task-graph, **explain** what it does!
			//
			// Max length: 32768
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/metadata/properties/description
			Description string `json:"description"`

			// Human readable name of task-graph
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/metadata/properties/name
			Name string `json:"name"`

			// E-mail of person who caused this task-graph, e.g. the person who did `hg push`
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/metadata/properties/owner
			Owner string `json:"owner"`

			// Link to source of this task-graph, should specify file, revision and repository
			//
			// Max length: 4096
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/metadata/properties/source
			Source string `json:"source"`
		} `json:"metadata"`

		// List of scopes (or scope-patterns) that tasks of the task-graph is authorized to use.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/scopes
		Scopes []string `json:"scopes"`

		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/status
		Status TaskGraphStatusStructure `json:"status"`

		// Arbitrary key-value tags (only strings limited to 4k)
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tags
		Tags json.RawMessage `json:"tags"`

		// Mapping from task-labels to task information and state.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks
		Tasks []struct {

			// List of `taskId`s that requires this task to be _complete successfully_ before they can be scheduled.
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/dependents
			Dependents []string `json:"dependents"`

			// Human readable name from the task definition
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/name
			Name string `json:"name"`

			// List of required `taskId`s
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/requires
			Requires []string `json:"requires"`

			// List of `taskId`s that have yet to complete successfully, before this task can be scheduled.
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/requiresLeft
			RequiresLeft []string `json:"requiresLeft"`

			// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
			//
			// Mininum:    0
			// Maximum:    999
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/reruns
			Reruns int `json:"reruns"`

			// Number of reruns that haven't been used yet.
			//
			// Mininum:    0
			// Maximum:    999
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/rerunsLeft
			RerunsLeft int `json:"rerunsLeft"`

			// true, if the scheduler considers the task node as satisfied and hence no-longer prevents dependent tasks from running.
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/satisfied
			Satisfied bool `json:"satisfied"`

			// State of the task as considered by the scheduler
			//
			// Possible values:
			//   * "unscheduled"
			//   * "scheduled"
			//   * "completed"
			//   * "failed"
			//   * "exception"
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/state
			State string `json:"state"`

			// Unique task identifier, this is UUID encoded as [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and stripped of `=` padding.
			//
			// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
			//
			// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-response.json#/properties/tasks/items/properties/taskId
			TaskID string `json:"taskId"`
		} `json:"tasks"`
	}

	// Information about a **task** in a task-graph as known by the scheduler.
	//
	// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#
	InspectTaskGraphTaskResponse struct {

		// List of `taskId`s that requires this task to be _complete successfully_ before they can be scheduled.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/dependents
		Dependents []string `json:"dependents"`

		// Human readable name from the task definition
		//
		// Max length: 255
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/name
		Name string `json:"name"`

		// List of required `taskId`s
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/requires
		Requires []string `json:"requires"`

		// List of `taskId`s that have yet to complete successfully, before this task can be scheduled.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/requiresLeft
		RequiresLeft []string `json:"requiresLeft"`

		// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
		//
		// Mininum:    0
		// Maximum:    999
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/reruns
		Reruns int `json:"reruns"`

		// Number of reruns that haven't been used yet.
		//
		// Mininum:    0
		// Maximum:    999
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/rerunsLeft
		RerunsLeft int `json:"rerunsLeft"`

		// true, if the scheduler considers the task node as satisfied and hence no-longer prevents dependent tasks from running.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/satisfied
		Satisfied bool `json:"satisfied"`

		// State of the task as considered by the scheduler
		//
		// Possible values:
		//   * "unscheduled"
		//   * "scheduled"
		//   * "completed"
		//   * "failed"
		//   * "exception"
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/state
		State string `json:"state"`

		// Unique task identifier, this is UUID encoded as [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and stripped of `=` padding.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		//
		// See http://schemas.taskcluster.net/scheduler/v1/inspect-task-graph-task-response.json#/properties/taskId
		TaskID string `json:"taskId"`
	}

	// Definition of a task that can be scheduled
	//
	// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#
	TaskDefinitionRequest struct {

		// Creation time of task
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/created
		Created tcclient.Time `json:"created"`

		// Deadline of the task, `pending` and `running` runs are
		// resolved as **failed** if not resolved by other means
		// before the deadline. Note, deadline cannot be more than
		// 5 days into the future
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/deadline
		Deadline tcclient.Time `json:"deadline"`

		// List of dependent tasks. These must either be _completed_ or _resolved_
		// before this task is scheduled. See `requires` for semantics.
		//
		// Default:    []
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/dependencies
		Dependencies []string `json:"dependencies,omitempty"`

		// Task expiration, time at which task definition and status is deleted.
		// Notice that all artifacts for the task must have an expiration that is no
		// later than this. If this property isn't it will be set to `deadline`
		// plus one year (this default may subject to change).
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/expires
		Expires tcclient.Time `json:"expires,omitempty"`

		// Object with properties that can hold any kind of extra data that should be
		// associated with the task. This can be data for the task which doesn't
		// fit into `payload`, or it can supplementary data for use in services
		// listening for events from this task. For example this could be details to
		// display on _treeherder_, or information for indexing the task. Please, try
		// to put all related information under one property, so `extra` data keys
		// for treeherder reporting and task indexing don't conflict, hence, we have
		// reusable services. **Warning**, do not stuff large data-sets in here,
		// task definitions should not take-up multiple MiBs.
		//
		// Default:    map[]
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/extra
		Extra json.RawMessage `json:"extra,omitempty"`

		// Required task metadata
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/metadata
		Metadata struct {

			// Human readable description of the task, please **explain** what the
			// task does. A few lines of documentation is not going to hurt you.
			//
			// Max length: 32768
			//
			// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/metadata/properties/description
			Description string `json:"description"`

			// Human readable name of task, used to very briefly given an idea about
			// what the task does.
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/metadata/properties/name
			Name string `json:"name"`

			// E-mail of person who caused this task, e.g. the person who did
			// `hg push`. The person we should contact to ask why this task is here.
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/metadata/properties/owner
			Owner string `json:"owner"`

			// Link to source of this task, should specify a file, revision and
			// repository. This should be place someone can go an do a git/hg blame
			// to who came up with recipe for this task.
			//
			// Max length: 4096
			//
			// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/metadata/properties/source
			Source string `json:"source"`
		} `json:"metadata"`

		// Task-specific payload following worker-specific format. For example the
		// `docker-worker` requires keys like: `image`, `commands` and
		// `features`. Refer to the documentation of `docker-worker` for details.
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/payload
		Payload json.RawMessage `json:"payload"`

		// Priority of task, this defaults to `normal`. Additional levels may be
		// added later.
		// **Task submitter required scopes** `queue:task-priority:high` for high
		// priority tasks.
		//
		// Possible values:
		//   * "high"
		//   * "normal"
		//
		// Default:    "normal"
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/priority
		Priority string `json:"priority,omitempty"`

		// Unique identifier for a provisioner, that can supply specified
		// `workerType`
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/provisionerId
		ProvisionerID string `json:"provisionerId"`

		// The tasks relation to its dependencies. This property specifies the
		// semantics of the `task.dependencies` property.
		// If `all-completed` is given the task will be scheduled when all
		// dependencies are resolved _completed_ (successful resolution).
		// If `all-resolved` is given the task will be scheduled when all dependencies
		// have been resolved, regardless of what their resolution is.
		//
		// Possible values:
		//   * "all-completed"
		//   * "all-resolved"
		//
		// Default:    "all-completed"
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/requires
		Requires string `json:"requires,omitempty"`

		// Number of times to retry the task in case of infrastructure issues.
		// An _infrastructure issue_ is a worker node that crashes or is shutdown,
		// these events are to be expected.
		//
		// Default:    5
		// Mininum:    0
		// Maximum:    49
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/retries
		Retries int `json:"retries,omitempty"`

		// List of task specific routes, AMQP messages will be CC'ed to these routes.
		// **Task submitter required scopes** `queue:route:<route>` for
		// each route given.
		//
		// Default:    []
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/routes
		Routes []string `json:"routes,omitempty"`

		// Identifier for the scheduler that _defined_ this task, this can be an
		// identifier for a user or a service like the `"task-graph-scheduler"`.
		// **Task submitter required scopes**
		// `queue:assume:scheduler-id:<schedulerId>/<taskGroupId>`.
		// This scope is also necessary to _schedule_ a defined task, or _rerun_ a
		// task.
		//
		// Default:    "-"
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/schedulerId
		SchedulerID string `json:"schedulerId,omitempty"`

		// List of scopes (or scope-patterns) that the task is
		// authorized to use.
		//
		// Default:    []
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/scopes
		Scopes []string `json:"scopes,omitempty"`

		// Arbitrary key-value tags (only strings limited to 4k). These can be used
		// to attach informal meta-data to a task. Use this for informal tags that
		// tasks can be classified by. You can also think of strings here as
		// candidates for formal meta-data. Something like
		// `purpose: 'build' || 'test'` is a good example.
		//
		// Default:    map[]
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/tags
		Tags json.RawMessage `json:"tags,omitempty"`

		// Identifier for a group of tasks scheduled together with this task, by
		// scheduler identified by `schedulerId`. For tasks scheduled by the
		// task-graph scheduler, this is the `taskGraphId`.  Defaults to `taskId` if
		// property isn't specified.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/taskGroupId
		TaskGroupID string `json:"taskGroupId,omitempty"`

		// Unique identifier for a worker-type within a specific provisioner
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		//
		// See http://schemas.taskcluster.net/queue/v1/create-task-request.json#/properties/workerType
		WorkerType string `json:"workerType"`
	}

	// Definition of a task-graph that can be scheduled
	//
	// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#
	TaskGraphDefinition struct {

		// Required task metadata
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/metadata
		Metadata struct {

			// Human readable description of task-graph, **explain** what it does!
			//
			// Max length: 32768
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/metadata/properties/description
			Description string `json:"description"`

			// Human readable name of task-graph, give people finding this an idea
			// what this graph is about.
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/metadata/properties/name
			Name string `json:"name"`

			// E-mail of person who caused this task-graph, e.g. the person who did
			// `hg push` or whatever triggered it.
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/metadata/properties/owner
			Owner string `json:"owner"`

			// Link to source of this task-graph, should specify file, revision and
			// repository
			//
			// Max length: 4096
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/metadata/properties/source
			Source string `json:"source"`
		} `json:"metadata"`

		// List of task-graph specific routes, AMQP messages will be CC'ed to these
		// routes prefixed by `'route.'`.
		//
		// Default:    []
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/routes
		Routes []string `json:"routes,omitempty"`

		// List of scopes (or scope-patterns) that tasks of the task-graph is
		// authorized to use.
		//
		// Default:    []
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/scopes
		Scopes []string `json:"scopes,omitempty"`

		// Arbitrary key-value tags (only strings limited to 4k)
		//
		// Default:    map[]
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/tags
		Tags json.RawMessage `json:"tags,omitempty"`

		// List of nodes in the task-graph, each featuring a task definition and scheduling preferences, such as number of _reruns_ to attempt.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/tasks
		Tasks []struct {

			// List of required `taskId`s
			//
			// Default:    []
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/tasks/items/properties/requires
			Requires []string `json:"requires,omitempty"`

			// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
			//
			// Default:    0
			// Mininum:    0
			// Maximum:    100
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/tasks/items/properties/reruns
			Reruns int `json:"reruns,omitempty"`

			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/tasks/items/properties/task
			Task TaskDefinitionRequest `json:"task"`

			// Task identifier (`taskId`) for the task when submitted to the queue, also used in `requires` below. This must be formatted as a **slugid** that is a uuid encoded in url-safe base64 following [RFC 4648 sec. 5](http://tools.ietf.org/html/rfc4648#section-5)), but without `==` padding.
			//
			// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph.json#/properties/tasks/items/properties/taskId
			TaskID string `json:"taskId"`
		} `json:"tasks"`
	}

	// Definition of a task-graph that can be scheduled
	//
	// See http://schemas.taskcluster.net/scheduler/v1/extend-task-graph-request.json#
	TaskGraphDefinition1 struct {

		// List of nodes in the task-graph, each featuring a task definition and scheduling preferences, such as number of _reruns_ to attempt.
		//
		// See http://schemas.taskcluster.net/scheduler/v1/extend-task-graph-request.json#/properties/tasks
		Tasks []struct {

			// List of required `taskId`s
			//
			// Default:    []
			//
			// See http://schemas.taskcluster.net/scheduler/v1/extend-task-graph-request.json#/properties/tasks/items/properties/requires
			Requires []string `json:"requires,omitempty"`

			// Number of times to _rerun_ the task if it completed unsuccessfully. **Note**, this does not capture _retries_ due to infrastructure issues.
			//
			// Default:    0
			// Mininum:    0
			// Maximum:    100
			//
			// See http://schemas.taskcluster.net/scheduler/v1/extend-task-graph-request.json#/properties/tasks/items/properties/reruns
			Reruns int `json:"reruns,omitempty"`

			// See http://schemas.taskcluster.net/scheduler/v1/extend-task-graph-request.json#/properties/tasks/items/properties/task
			Task TaskDefinitionRequest `json:"task"`

			// Task identifier (`taskId`) for the task when submitted to the queue, also used in `requires` below. This must be formatted as a **slugid** that is a uuid encoded in url-safe base64 following [RFC 4648 sec. 5](http://tools.ietf.org/html/rfc4648#section-5)), but without `==` padding.
			//
			// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
			//
			// See http://schemas.taskcluster.net/scheduler/v1/extend-task-graph-request.json#/properties/tasks/items/properties/taskId
			TaskID string `json:"taskId"`
		} `json:"tasks"`
	}

	// Response for a request for task-graph information
	//
	// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#
	TaskGraphInfoResponse struct {

		// Required task metadata
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/metadata
		Metadata struct {

			// Human readable description of task-graph, **explain** what it does!
			//
			// Max length: 32768
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/metadata/properties/description
			Description string `json:"description"`

			// Human readable name of task-graph
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/metadata/properties/name
			Name string `json:"name"`

			// E-mail of person who caused this task-graph, e.g. the person who did `hg push`
			//
			// Max length: 255
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/metadata/properties/owner
			Owner string `json:"owner"`

			// Link to source of this task-graph, should specify file, revision and repository
			//
			// Max length: 4096
			//
			// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/metadata/properties/source
			Source string `json:"source"`
		} `json:"metadata"`

		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/status
		Status TaskGraphStatusStructure `json:"status"`

		// Arbitrary key-value tags (only strings limited to 4k)
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-info-response.json#/properties/tags
		Tags json.RawMessage `json:"tags"`
	}

	// Response containing the status structure for a task-graph
	//
	// See http://schemas.taskcluster.net/scheduler/v1/task-graph-status-response.json#
	TaskGraphStatusResponse struct {

		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-status-response.json#/properties/status
		Status TaskGraphStatusStructure `json:"status"`
	}

	// A representation of **task-graph status** as known by the scheduler, without the state of all individual tasks.
	//
	// See http://schemas.taskcluster.net/scheduler/v1/task-graph-status.json#
	TaskGraphStatusStructure struct {

		// Unique identifier for task-graph scheduler managing the given task-graph
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-status.json#/properties/schedulerId
		SchedulerID string `json:"schedulerId"`

		// Task-graph state, this enum is **frozen** new values will **not** be added.
		//
		// Possible values:
		//   * "running"
		//   * "blocked"
		//   * "finished"
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-status.json#/properties/state
		State string `json:"state"`

		// Unique task-graph identifier, this is UUID encoded as [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and stripped of `=` padding.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		//
		// See http://schemas.taskcluster.net/scheduler/v1/task-graph-status.json#/properties/taskGraphId
		TaskGraphID string `json:"taskGraphId"`
	}
)
