// Code generated by github.com/jewzaam/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"context"
	"net/http"
)

// Permission represents a permission
type Permission struct {
	ID             string         `json:"id,omitempty"`
	ResourceID     string         `json:"_rid,omitempty"`
	Timestamp      int            `json:"_ts,omitempty"`
	Self           string         `json:"_self,omitempty"`
	ETag           string         `json:"_etag,omitempty"`
	Token          string         `json:"_token,omitempty"`
	PermissionMode PermissionMode `json:"permissionMode,omitempty"`
	Resource       string         `json:"resource,omitempty"`
}

// PermissionMode represents a permission mode
type PermissionMode string

// PermissionMode constants
const (
	PermissionModeAll  PermissionMode = "All"
	PermissionModeRead PermissionMode = "Read"
)

// Permissions represents permissions
type Permissions struct {
	Count       int           `json:"_count,omitempty"`
	ResourceID  string        `json:"_rid,omitempty"`
	Permissions []*Permission `json:"Permissions,omitempty"`
}

type permissionClient struct {
	*databaseClient
	path string
}

// PermissionClient is a permission client
type PermissionClient interface {
	Create(context.Context, *Permission) (*Permission, error)
	List() PermissionIterator
	ListAll(context.Context) (*Permissions, error)
	Get(context.Context, string) (*Permission, error)
	Delete(context.Context, *Permission) error
	Replace(context.Context, *Permission) (*Permission, error)
}

type permissionListIterator struct {
	*permissionClient
	continuation string
	done         bool
}

// PermissionIterator is a permission iterator
type PermissionIterator interface {
	Next(context.Context) (*Permissions, error)
}

// NewPermissionClient returns a new permission client
func NewPermissionClient(userc UserClient, userid string) PermissionClient {
	return &permissionClient{
		databaseClient: userc.(*userClient).databaseClient,
		path:           userc.(*userClient).path + "/users/" + userid,
	}
}

func (c *permissionClient) all(ctx context.Context, i PermissionIterator) (*Permissions, error) {
	allpermissions := &Permissions{}

	for {
		permissions, err := i.Next(ctx)
		if err != nil {
			return nil, err
		}
		if permissions == nil {
			break
		}

		allpermissions.Count += permissions.Count
		allpermissions.ResourceID = permissions.ResourceID
		allpermissions.Permissions = append(allpermissions.Permissions, permissions.Permissions...)
	}

	return allpermissions, nil
}

func (c *permissionClient) Create(ctx context.Context, newpermission *Permission) (permission *Permission, err error) {
	err = c.do(ctx, http.MethodPost, c.path+"/permissions", "permissions", c.path, http.StatusCreated, &newpermission, &permission, nil)
	return
}

func (c *permissionClient) List() PermissionIterator {
	return &permissionListIterator{permissionClient: c}
}

func (c *permissionClient) ListAll(ctx context.Context) (*Permissions, error) {
	return c.all(ctx, c.List())
}

func (c *permissionClient) Get(ctx context.Context, permissionid string) (permission *Permission, err error) {
	err = c.do(ctx, http.MethodGet, c.path+"/permissions/"+permissionid, "permissions", c.path+"/permissions/"+permissionid, http.StatusOK, nil, &permission, nil)
	return
}

func (c *permissionClient) Delete(ctx context.Context, permission *Permission) error {
	if permission.ETag == "" {
		return ErrETagRequired
	}
	headers := http.Header{}
	headers.Set("If-Match", permission.ETag)
	return c.do(ctx, http.MethodDelete, c.path+"/permissions/"+permission.ID, "permissions", c.path+"/permissions/"+permission.ID, http.StatusNoContent, nil, nil, headers)
}

func (c *permissionClient) Replace(ctx context.Context, newpermission *Permission) (permission *Permission, err error) {
	err = c.do(ctx, http.MethodPost, c.path+"/permissions/"+newpermission.ID, "permissions", c.path+"/permissions/"+newpermission.ID, http.StatusCreated, &newpermission, &permission, nil)
	return
}

func (i *permissionListIterator) Next(ctx context.Context) (permissions *Permissions, err error) {
	if i.done {
		return
	}

	headers := http.Header{}
	if i.continuation != "" {
		headers.Set("X-Ms-Continuation", i.continuation)
	}

	err = i.do(ctx, http.MethodGet, i.path+"/permissions", "permissions", i.path, http.StatusOK, nil, &permissions, headers)
	if err != nil {
		return
	}

	i.continuation = headers.Get("X-Ms-Continuation")
	i.done = i.continuation == ""

	return
}
