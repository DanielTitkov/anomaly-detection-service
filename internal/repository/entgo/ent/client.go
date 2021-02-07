// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/migrate"

	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/systemsummary"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/tasktype"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/user"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// SystemSummary is the client for interacting with the SystemSummary builders.
	SystemSummary *SystemSummaryClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
	// TaskType is the client for interacting with the TaskType builders.
	TaskType *TaskTypeClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Item = NewItemClient(c.config)
	c.SystemSummary = NewSystemSummaryClient(c.config)
	c.Task = NewTaskClient(c.config)
	c.TaskType = NewTaskTypeClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Item:          NewItemClient(cfg),
		SystemSummary: NewSystemSummaryClient(cfg),
		Task:          NewTaskClient(cfg),
		TaskType:      NewTaskTypeClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:        cfg,
		Item:          NewItemClient(cfg),
		SystemSummary: NewSystemSummaryClient(cfg),
		Task:          NewTaskClient(cfg),
		TaskType:      NewTaskTypeClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Item.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Item.Use(hooks...)
	c.SystemSummary.Use(hooks...)
	c.Task.Use(hooks...)
	c.TaskType.Use(hooks...)
	c.User.Use(hooks...)
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `item.Hooks(f(g(h())))`.
func (c *ItemClient) Use(hooks ...Hook) {
	c.hooks.Item = append(c.hooks.Item, hooks...)
}

// Create returns a create builder for Item.
func (c *ItemClient) Create() *ItemCreate {
	mutation := newItemMutation(c.config, OpCreate)
	return &ItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Item entities.
func (c *ItemClient) CreateBulk(builders ...*ItemCreate) *ItemCreateBulk {
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	mutation := newItemMutation(c.config, OpUpdate)
	return &ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItem(i))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id int) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItemID(id))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	mutation := newItemMutation(c.config, OpDelete)
	return &ItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemClient) DeleteOneID(id int) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{config: c.config}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id int) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id int) *Item {
	i, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return i
}

// QueryTask queries the task edge of a Item.
func (c *ItemClient) QueryTask(i *Item) *TaskQuery {
	query := &TaskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, item.TaskTable, item.TaskColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemClient) Hooks() []Hook {
	return c.hooks.Item
}

// SystemSummaryClient is a client for the SystemSummary schema.
type SystemSummaryClient struct {
	config
}

// NewSystemSummaryClient returns a client for the SystemSummary from the given config.
func NewSystemSummaryClient(c config) *SystemSummaryClient {
	return &SystemSummaryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `systemsummary.Hooks(f(g(h())))`.
func (c *SystemSummaryClient) Use(hooks ...Hook) {
	c.hooks.SystemSummary = append(c.hooks.SystemSummary, hooks...)
}

// Create returns a create builder for SystemSummary.
func (c *SystemSummaryClient) Create() *SystemSummaryCreate {
	mutation := newSystemSummaryMutation(c.config, OpCreate)
	return &SystemSummaryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of SystemSummary entities.
func (c *SystemSummaryClient) CreateBulk(builders ...*SystemSummaryCreate) *SystemSummaryCreateBulk {
	return &SystemSummaryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SystemSummary.
func (c *SystemSummaryClient) Update() *SystemSummaryUpdate {
	mutation := newSystemSummaryMutation(c.config, OpUpdate)
	return &SystemSummaryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SystemSummaryClient) UpdateOne(ss *SystemSummary) *SystemSummaryUpdateOne {
	mutation := newSystemSummaryMutation(c.config, OpUpdateOne, withSystemSummary(ss))
	return &SystemSummaryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SystemSummaryClient) UpdateOneID(id int) *SystemSummaryUpdateOne {
	mutation := newSystemSummaryMutation(c.config, OpUpdateOne, withSystemSummaryID(id))
	return &SystemSummaryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SystemSummary.
func (c *SystemSummaryClient) Delete() *SystemSummaryDelete {
	mutation := newSystemSummaryMutation(c.config, OpDelete)
	return &SystemSummaryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SystemSummaryClient) DeleteOne(ss *SystemSummary) *SystemSummaryDeleteOne {
	return c.DeleteOneID(ss.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SystemSummaryClient) DeleteOneID(id int) *SystemSummaryDeleteOne {
	builder := c.Delete().Where(systemsummary.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SystemSummaryDeleteOne{builder}
}

// Query returns a query builder for SystemSummary.
func (c *SystemSummaryClient) Query() *SystemSummaryQuery {
	return &SystemSummaryQuery{config: c.config}
}

// Get returns a SystemSummary entity by its id.
func (c *SystemSummaryClient) Get(ctx context.Context, id int) (*SystemSummary, error) {
	return c.Query().Where(systemsummary.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SystemSummaryClient) GetX(ctx context.Context, id int) *SystemSummary {
	ss, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ss
}

// Hooks returns the client hooks.
func (c *SystemSummaryClient) Hooks() []Hook {
	return c.hooks.SystemSummary
}

// TaskClient is a client for the Task schema.
type TaskClient struct {
	config
}

// NewTaskClient returns a client for the Task from the given config.
func NewTaskClient(c config) *TaskClient {
	return &TaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `task.Hooks(f(g(h())))`.
func (c *TaskClient) Use(hooks ...Hook) {
	c.hooks.Task = append(c.hooks.Task, hooks...)
}

// Create returns a create builder for Task.
func (c *TaskClient) Create() *TaskCreate {
	mutation := newTaskMutation(c.config, OpCreate)
	return &TaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Task entities.
func (c *TaskClient) CreateBulk(builders ...*TaskCreate) *TaskCreateBulk {
	return &TaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Task.
func (c *TaskClient) Update() *TaskUpdate {
	mutation := newTaskMutation(c.config, OpUpdate)
	return &TaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskClient) UpdateOne(t *Task) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTask(t))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskClient) UpdateOneID(id int) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTaskID(id))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Task.
func (c *TaskClient) Delete() *TaskDelete {
	mutation := newTaskMutation(c.config, OpDelete)
	return &TaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TaskClient) DeleteOne(t *Task) *TaskDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TaskClient) DeleteOneID(id int) *TaskDeleteOne {
	builder := c.Delete().Where(task.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaskDeleteOne{builder}
}

// Query returns a query builder for Task.
func (c *TaskClient) Query() *TaskQuery {
	return &TaskQuery{config: c.config}
}

// Get returns a Task entity by its id.
func (c *TaskClient) Get(ctx context.Context, id int) (*Task, error) {
	return c.Query().Where(task.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskClient) GetX(ctx context.Context, id int) *Task {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryItems queries the items edge of a Task.
func (c *TaskClient) QueryItems(t *Task) *ItemQuery {
	query := &ItemQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, task.ItemsTable, task.ItemsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Task.
func (c *TaskClient) QueryUser(t *Task) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, task.UserTable, task.UserColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryType queries the type edge of a Task.
func (c *TaskClient) QueryType(t *Task) *TaskTypeQuery {
	query := &TaskTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, id),
			sqlgraph.To(tasktype.Table, tasktype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, task.TypeTable, task.TypeColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TaskClient) Hooks() []Hook {
	return c.hooks.Task
}

// TaskTypeClient is a client for the TaskType schema.
type TaskTypeClient struct {
	config
}

// NewTaskTypeClient returns a client for the TaskType from the given config.
func NewTaskTypeClient(c config) *TaskTypeClient {
	return &TaskTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tasktype.Hooks(f(g(h())))`.
func (c *TaskTypeClient) Use(hooks ...Hook) {
	c.hooks.TaskType = append(c.hooks.TaskType, hooks...)
}

// Create returns a create builder for TaskType.
func (c *TaskTypeClient) Create() *TaskTypeCreate {
	mutation := newTaskTypeMutation(c.config, OpCreate)
	return &TaskTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of TaskType entities.
func (c *TaskTypeClient) CreateBulk(builders ...*TaskTypeCreate) *TaskTypeCreateBulk {
	return &TaskTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TaskType.
func (c *TaskTypeClient) Update() *TaskTypeUpdate {
	mutation := newTaskTypeMutation(c.config, OpUpdate)
	return &TaskTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskTypeClient) UpdateOne(tt *TaskType) *TaskTypeUpdateOne {
	mutation := newTaskTypeMutation(c.config, OpUpdateOne, withTaskType(tt))
	return &TaskTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskTypeClient) UpdateOneID(id int) *TaskTypeUpdateOne {
	mutation := newTaskTypeMutation(c.config, OpUpdateOne, withTaskTypeID(id))
	return &TaskTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TaskType.
func (c *TaskTypeClient) Delete() *TaskTypeDelete {
	mutation := newTaskTypeMutation(c.config, OpDelete)
	return &TaskTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TaskTypeClient) DeleteOne(tt *TaskType) *TaskTypeDeleteOne {
	return c.DeleteOneID(tt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TaskTypeClient) DeleteOneID(id int) *TaskTypeDeleteOne {
	builder := c.Delete().Where(tasktype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaskTypeDeleteOne{builder}
}

// Query returns a query builder for TaskType.
func (c *TaskTypeClient) Query() *TaskTypeQuery {
	return &TaskTypeQuery{config: c.config}
}

// Get returns a TaskType entity by its id.
func (c *TaskTypeClient) Get(ctx context.Context, id int) (*TaskType, error) {
	return c.Query().Where(tasktype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskTypeClient) GetX(ctx context.Context, id int) *TaskType {
	tt, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return tt
}

// QueryTasks queries the tasks edge of a TaskType.
func (c *TaskTypeClient) QueryTasks(tt *TaskType) *TaskQuery {
	query := &TaskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tasktype.Table, tasktype.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tasktype.TasksTable, tasktype.TasksColumn),
		)
		fromV = sqlgraph.Neighbors(tt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TaskTypeClient) Hooks() []Hook {
	return c.hooks.TaskType
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryTasks queries the tasks edge of a User.
func (c *UserClient) QueryTasks(u *User) *TaskQuery {
	query := &TaskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TasksTable, user.TasksColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
