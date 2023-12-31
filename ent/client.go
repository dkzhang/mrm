// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"mrm/ent/migrate"

	"mrm/ent/meeting"
	"mrm/ent/meetingdateroom"
	"mrm/ent/room"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Meeting is the client for interacting with the Meeting builders.
	Meeting *MeetingClient
	// MeetingDateRoom is the client for interacting with the MeetingDateRoom builders.
	MeetingDateRoom *MeetingDateRoomClient
	// Room is the client for interacting with the Room builders.
	Room *RoomClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Meeting = NewMeetingClient(c.config)
	c.MeetingDateRoom = NewMeetingDateRoomClient(c.config)
	c.Room = NewRoomClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Meeting:         NewMeetingClient(cfg),
		MeetingDateRoom: NewMeetingDateRoomClient(cfg),
		Room:            NewRoomClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Meeting:         NewMeetingClient(cfg),
		MeetingDateRoom: NewMeetingDateRoomClient(cfg),
		Room:            NewRoomClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Meeting.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Meeting.Use(hooks...)
	c.MeetingDateRoom.Use(hooks...)
	c.Room.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Meeting.Intercept(interceptors...)
	c.MeetingDateRoom.Intercept(interceptors...)
	c.Room.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MeetingMutation:
		return c.Meeting.mutate(ctx, m)
	case *MeetingDateRoomMutation:
		return c.MeetingDateRoom.mutate(ctx, m)
	case *RoomMutation:
		return c.Room.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MeetingClient is a client for the Meeting schema.
type MeetingClient struct {
	config
}

// NewMeetingClient returns a client for the Meeting from the given config.
func NewMeetingClient(c config) *MeetingClient {
	return &MeetingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `meeting.Hooks(f(g(h())))`.
func (c *MeetingClient) Use(hooks ...Hook) {
	c.hooks.Meeting = append(c.hooks.Meeting, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `meeting.Intercept(f(g(h())))`.
func (c *MeetingClient) Intercept(interceptors ...Interceptor) {
	c.inters.Meeting = append(c.inters.Meeting, interceptors...)
}

// Create returns a builder for creating a Meeting entity.
func (c *MeetingClient) Create() *MeetingCreate {
	mutation := newMeetingMutation(c.config, OpCreate)
	return &MeetingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Meeting entities.
func (c *MeetingClient) CreateBulk(builders ...*MeetingCreate) *MeetingCreateBulk {
	return &MeetingCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MeetingClient) MapCreateBulk(slice any, setFunc func(*MeetingCreate, int)) *MeetingCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MeetingCreateBulk{err: fmt.Errorf("calling to MeetingClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MeetingCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MeetingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Meeting.
func (c *MeetingClient) Update() *MeetingUpdate {
	mutation := newMeetingMutation(c.config, OpUpdate)
	return &MeetingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MeetingClient) UpdateOne(m *Meeting) *MeetingUpdateOne {
	mutation := newMeetingMutation(c.config, OpUpdateOne, withMeeting(m))
	return &MeetingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MeetingClient) UpdateOneID(id int64) *MeetingUpdateOne {
	mutation := newMeetingMutation(c.config, OpUpdateOne, withMeetingID(id))
	return &MeetingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Meeting.
func (c *MeetingClient) Delete() *MeetingDelete {
	mutation := newMeetingMutation(c.config, OpDelete)
	return &MeetingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MeetingClient) DeleteOne(m *Meeting) *MeetingDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MeetingClient) DeleteOneID(id int64) *MeetingDeleteOne {
	builder := c.Delete().Where(meeting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MeetingDeleteOne{builder}
}

// Query returns a query builder for Meeting.
func (c *MeetingClient) Query() *MeetingQuery {
	return &MeetingQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMeeting},
		inters: c.Interceptors(),
	}
}

// Get returns a Meeting entity by its id.
func (c *MeetingClient) Get(ctx context.Context, id int64) (*Meeting, error) {
	return c.Query().Where(meeting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MeetingClient) GetX(ctx context.Context, id int64) *Meeting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMdrs queries the mdrs edge of a Meeting.
func (c *MeetingClient) QueryMdrs(m *Meeting) *MeetingDateRoomQuery {
	query := (&MeetingDateRoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(meeting.Table, meeting.FieldID, id),
			sqlgraph.To(meetingdateroom.Table, meetingdateroom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, meeting.MdrsTable, meeting.MdrsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MeetingClient) Hooks() []Hook {
	return c.hooks.Meeting
}

// Interceptors returns the client interceptors.
func (c *MeetingClient) Interceptors() []Interceptor {
	return c.inters.Meeting
}

func (c *MeetingClient) mutate(ctx context.Context, m *MeetingMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MeetingCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MeetingUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MeetingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MeetingDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Meeting mutation op: %q", m.Op())
	}
}

// MeetingDateRoomClient is a client for the MeetingDateRoom schema.
type MeetingDateRoomClient struct {
	config
}

// NewMeetingDateRoomClient returns a client for the MeetingDateRoom from the given config.
func NewMeetingDateRoomClient(c config) *MeetingDateRoomClient {
	return &MeetingDateRoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `meetingdateroom.Hooks(f(g(h())))`.
func (c *MeetingDateRoomClient) Use(hooks ...Hook) {
	c.hooks.MeetingDateRoom = append(c.hooks.MeetingDateRoom, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `meetingdateroom.Intercept(f(g(h())))`.
func (c *MeetingDateRoomClient) Intercept(interceptors ...Interceptor) {
	c.inters.MeetingDateRoom = append(c.inters.MeetingDateRoom, interceptors...)
}

// Create returns a builder for creating a MeetingDateRoom entity.
func (c *MeetingDateRoomClient) Create() *MeetingDateRoomCreate {
	mutation := newMeetingDateRoomMutation(c.config, OpCreate)
	return &MeetingDateRoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MeetingDateRoom entities.
func (c *MeetingDateRoomClient) CreateBulk(builders ...*MeetingDateRoomCreate) *MeetingDateRoomCreateBulk {
	return &MeetingDateRoomCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MeetingDateRoomClient) MapCreateBulk(slice any, setFunc func(*MeetingDateRoomCreate, int)) *MeetingDateRoomCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MeetingDateRoomCreateBulk{err: fmt.Errorf("calling to MeetingDateRoomClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MeetingDateRoomCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MeetingDateRoomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MeetingDateRoom.
func (c *MeetingDateRoomClient) Update() *MeetingDateRoomUpdate {
	mutation := newMeetingDateRoomMutation(c.config, OpUpdate)
	return &MeetingDateRoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MeetingDateRoomClient) UpdateOne(mdr *MeetingDateRoom) *MeetingDateRoomUpdateOne {
	mutation := newMeetingDateRoomMutation(c.config, OpUpdateOne, withMeetingDateRoom(mdr))
	return &MeetingDateRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MeetingDateRoomClient) UpdateOneID(id int) *MeetingDateRoomUpdateOne {
	mutation := newMeetingDateRoomMutation(c.config, OpUpdateOne, withMeetingDateRoomID(id))
	return &MeetingDateRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MeetingDateRoom.
func (c *MeetingDateRoomClient) Delete() *MeetingDateRoomDelete {
	mutation := newMeetingDateRoomMutation(c.config, OpDelete)
	return &MeetingDateRoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MeetingDateRoomClient) DeleteOne(mdr *MeetingDateRoom) *MeetingDateRoomDeleteOne {
	return c.DeleteOneID(mdr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MeetingDateRoomClient) DeleteOneID(id int) *MeetingDateRoomDeleteOne {
	builder := c.Delete().Where(meetingdateroom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MeetingDateRoomDeleteOne{builder}
}

// Query returns a query builder for MeetingDateRoom.
func (c *MeetingDateRoomClient) Query() *MeetingDateRoomQuery {
	return &MeetingDateRoomQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMeetingDateRoom},
		inters: c.Interceptors(),
	}
}

// Get returns a MeetingDateRoom entity by its id.
func (c *MeetingDateRoomClient) Get(ctx context.Context, id int) (*MeetingDateRoom, error) {
	return c.Query().Where(meetingdateroom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MeetingDateRoomClient) GetX(ctx context.Context, id int) *MeetingDateRoom {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMeeting queries the meeting edge of a MeetingDateRoom.
func (c *MeetingDateRoomClient) QueryMeeting(mdr *MeetingDateRoom) *MeetingQuery {
	query := (&MeetingClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := mdr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(meetingdateroom.Table, meetingdateroom.FieldID, id),
			sqlgraph.To(meeting.Table, meeting.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, meetingdateroom.MeetingTable, meetingdateroom.MeetingColumn),
		)
		fromV = sqlgraph.Neighbors(mdr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoom queries the room edge of a MeetingDateRoom.
func (c *MeetingDateRoomClient) QueryRoom(mdr *MeetingDateRoom) *RoomQuery {
	query := (&RoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := mdr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(meetingdateroom.Table, meetingdateroom.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, meetingdateroom.RoomTable, meetingdateroom.RoomColumn),
		)
		fromV = sqlgraph.Neighbors(mdr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MeetingDateRoomClient) Hooks() []Hook {
	return c.hooks.MeetingDateRoom
}

// Interceptors returns the client interceptors.
func (c *MeetingDateRoomClient) Interceptors() []Interceptor {
	return c.inters.MeetingDateRoom
}

func (c *MeetingDateRoomClient) mutate(ctx context.Context, m *MeetingDateRoomMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MeetingDateRoomCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MeetingDateRoomUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MeetingDateRoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MeetingDateRoomDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown MeetingDateRoom mutation op: %q", m.Op())
	}
}

// RoomClient is a client for the Room schema.
type RoomClient struct {
	config
}

// NewRoomClient returns a client for the Room from the given config.
func NewRoomClient(c config) *RoomClient {
	return &RoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `room.Hooks(f(g(h())))`.
func (c *RoomClient) Use(hooks ...Hook) {
	c.hooks.Room = append(c.hooks.Room, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `room.Intercept(f(g(h())))`.
func (c *RoomClient) Intercept(interceptors ...Interceptor) {
	c.inters.Room = append(c.inters.Room, interceptors...)
}

// Create returns a builder for creating a Room entity.
func (c *RoomClient) Create() *RoomCreate {
	mutation := newRoomMutation(c.config, OpCreate)
	return &RoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Room entities.
func (c *RoomClient) CreateBulk(builders ...*RoomCreate) *RoomCreateBulk {
	return &RoomCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoomClient) MapCreateBulk(slice any, setFunc func(*RoomCreate, int)) *RoomCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoomCreateBulk{err: fmt.Errorf("calling to RoomClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoomCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Room.
func (c *RoomClient) Update() *RoomUpdate {
	mutation := newRoomMutation(c.config, OpUpdate)
	return &RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoomClient) UpdateOne(r *Room) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoom(r))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoomClient) UpdateOneID(id int) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoomID(id))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Room.
func (c *RoomClient) Delete() *RoomDelete {
	mutation := newRoomMutation(c.config, OpDelete)
	return &RoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoomClient) DeleteOne(r *Room) *RoomDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoomClient) DeleteOneID(id int) *RoomDeleteOne {
	builder := c.Delete().Where(room.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoomDeleteOne{builder}
}

// Query returns a query builder for Room.
func (c *RoomClient) Query() *RoomQuery {
	return &RoomQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRoom},
		inters: c.Interceptors(),
	}
}

// Get returns a Room entity by its id.
func (c *RoomClient) Get(ctx context.Context, id int) (*Room, error) {
	return c.Query().Where(room.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoomClient) GetX(ctx context.Context, id int) *Room {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMdrs queries the mdrs edge of a Room.
func (c *RoomClient) QueryMdrs(r *Room) *MeetingDateRoomQuery {
	query := (&MeetingDateRoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(meetingdateroom.Table, meetingdateroom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, room.MdrsTable, room.MdrsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoomClient) Hooks() []Hook {
	return c.hooks.Room
}

// Interceptors returns the client interceptors.
func (c *RoomClient) Interceptors() []Interceptor {
	return c.inters.Room
}

func (c *RoomClient) mutate(ctx context.Context, m *RoomMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoomCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoomDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Room mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Meeting, MeetingDateRoom, Room []ent.Hook
	}
	inters struct {
		Meeting, MeetingDateRoom, Room []ent.Interceptor
	}
)
