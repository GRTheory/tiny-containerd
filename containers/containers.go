package containers

import (
	"context"
	"time"

	"github.com/containerd/typeurl/v2"
)

// Container represents the set of data pinned by a container. Unless otherwise
// noted, the resources here are considered in use by the container.
//
// The resources specified in this object are used to create tasks from the container.
type Container struct {
	// ID uniquely identifies the container in a namespace.
	//
	// This property is required and cannot be changed after creation.
	ID string

	// Labels provide metadata extension for a container.
	//
	// These are optional and fully mutable.
	Labels map[string]string

	// Image specifies the image reference used for a container.
	//
	// This property is optional and mutable.
	Image string

	// Runtime specifies which runtime should be used when launching container
	// tasks.
	//
	// this property is required and immutable.
	Runtime RuntimeInfo

	// Spec should carry the runtime specification used to implement the
	// container.
	//
	// This field is required but mutable.
	Spec typeurl.Any

	// SnapshotKey specifies the snapshot key to use for the container's root
	// filesystem. When starting a task from this container, a caller should
	// look up the mounts from the snapshot service and include those on the
	// task create request.
	//
	// This field is not required by mutable.
	SnapshotKey string

	// Snapshotter specifies the snapshotter name used for rootfs
	//
	// This field is not required but immutable.
	Snapshotter string

	// CreateAt is the time at which the container was created.
	CreateAt time.Time

	// UpdateAt is the time at which the container was updated.
	UpdateAt time.Time

	// Extensions stores client-specified metadata
	Extensions map[string]typeurl.Any

	// SandboxID is an identifier of sandbox this container belongs to.
	//
	// This property is optional, but can't be changed after creation.
	SandboxID string
}

// RuntimeInfo holds runtime specific information.
type RuntimeInfo struct {
	Name    string
	Options typeurl.Any
}

// Store interacts with the underlying container storage
type Store interface {
	// Get a container using the id.
	//
	// Container object is returned on success. If the id is not known to the
	// store, an error will be returned.
	Get(ctx context.Context, id string) (Container, error)

	// List returns containers that match one or more of the provided filters.
	List(ctx context.Context, filters ...string) ([]Container, error)

	// Create a container in the store from the provided container.
	Create(ctx context.Context, container Container) (Container, error)

	// Update the container with the provided container object. ID must be set.
	//
	// If one or more fieldpaths are provided, only the field corresponding to
	// the fieldpaths will be mutated.
	Update(ctx context.Context, container Container, fieldpaths ...string) (Container, error)

	// Delete a container using the id.
	//
	// nil will be returned on success. If the container is not know to the
	// store, ErrNotFound will be returned.
	Delete(ctx context.Context, id string) error
}
