// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/api"
	"github.com/suyuan32/simple-admin-core/rpc/ent/configuration"
	"github.com/suyuan32/simple-admin-core/rpc/ent/department"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/rpc/ent/medicine"
	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"
	"github.com/suyuan32/simple-admin-core/rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/rpc/ent/position"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"
	"github.com/suyuan32/simple-admin-core/rpc/ent/token"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The APIFunc type is an adapter to allow the use of ordinary function as a Querier.
type APIFunc func(context.Context, *ent.APIQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f APIFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.APIQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.APIQuery", q)
}

// The TraverseAPI type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAPI func(context.Context, *ent.APIQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAPI) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAPI) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.APIQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.APIQuery", q)
}

// The ConfigurationFunc type is an adapter to allow the use of ordinary function as a Querier.
type ConfigurationFunc func(context.Context, *ent.ConfigurationQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ConfigurationFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ConfigurationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ConfigurationQuery", q)
}

// The TraverseConfiguration type is an adapter to allow the use of ordinary function as Traverser.
type TraverseConfiguration func(context.Context, *ent.ConfigurationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseConfiguration) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseConfiguration) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ConfigurationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ConfigurationQuery", q)
}

// The DepartmentFunc type is an adapter to allow the use of ordinary function as a Querier.
type DepartmentFunc func(context.Context, *ent.DepartmentQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DepartmentFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DepartmentQuery", q)
}

// The TraverseDepartment type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDepartment func(context.Context, *ent.DepartmentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDepartment) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDepartment) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DepartmentQuery", q)
}

// The DictionaryFunc type is an adapter to allow the use of ordinary function as a Querier.
type DictionaryFunc func(context.Context, *ent.DictionaryQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DictionaryFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DictionaryQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DictionaryQuery", q)
}

// The TraverseDictionary type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDictionary func(context.Context, *ent.DictionaryQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDictionary) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDictionary) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DictionaryQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DictionaryQuery", q)
}

// The DictionaryDetailFunc type is an adapter to allow the use of ordinary function as a Querier.
type DictionaryDetailFunc func(context.Context, *ent.DictionaryDetailQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DictionaryDetailFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DictionaryDetailQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DictionaryDetailQuery", q)
}

// The TraverseDictionaryDetail type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDictionaryDetail func(context.Context, *ent.DictionaryDetailQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDictionaryDetail) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDictionaryDetail) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DictionaryDetailQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DictionaryDetailQuery", q)
}

// The MedicineFunc type is an adapter to allow the use of ordinary function as a Querier.
type MedicineFunc func(context.Context, *ent.MedicineQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MedicineFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MedicineQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MedicineQuery", q)
}

// The TraverseMedicine type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMedicine func(context.Context, *ent.MedicineQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMedicine) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMedicine) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MedicineQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MedicineQuery", q)
}

// The MenuFunc type is an adapter to allow the use of ordinary function as a Querier.
type MenuFunc func(context.Context, *ent.MenuQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MenuFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MenuQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MenuQuery", q)
}

// The TraverseMenu type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMenu func(context.Context, *ent.MenuQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMenu) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMenu) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MenuQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MenuQuery", q)
}

// The OauthProviderFunc type is an adapter to allow the use of ordinary function as a Querier.
type OauthProviderFunc func(context.Context, *ent.OauthProviderQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f OauthProviderFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.OauthProviderQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.OauthProviderQuery", q)
}

// The TraverseOauthProvider type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOauthProvider func(context.Context, *ent.OauthProviderQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOauthProvider) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOauthProvider) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OauthProviderQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.OauthProviderQuery", q)
}

// The PositionFunc type is an adapter to allow the use of ordinary function as a Querier.
type PositionFunc func(context.Context, *ent.PositionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f PositionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.PositionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.PositionQuery", q)
}

// The TraversePosition type is an adapter to allow the use of ordinary function as Traverser.
type TraversePosition func(context.Context, *ent.PositionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePosition) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePosition) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PositionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.PositionQuery", q)
}

// The RoleFunc type is an adapter to allow the use of ordinary function as a Querier.
type RoleFunc func(context.Context, *ent.RoleQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f RoleFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.RoleQuery", q)
}

// The TraverseRole type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRole func(context.Context, *ent.RoleQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRole) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRole) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoleQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.RoleQuery", q)
}

// The TokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type TokenFunc func(context.Context, *ent.TokenQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TokenFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TokenQuery", q)
}

// The TraverseToken type is an adapter to allow the use of ordinary function as Traverser.
type TraverseToken func(context.Context, *ent.TokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseToken) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseToken) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TokenQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *ent.UserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *ent.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.APIQuery:
		return &query[*ent.APIQuery, predicate.API, api.OrderOption]{typ: ent.TypeAPI, tq: q}, nil
	case *ent.ConfigurationQuery:
		return &query[*ent.ConfigurationQuery, predicate.Configuration, configuration.OrderOption]{typ: ent.TypeConfiguration, tq: q}, nil
	case *ent.DepartmentQuery:
		return &query[*ent.DepartmentQuery, predicate.Department, department.OrderOption]{typ: ent.TypeDepartment, tq: q}, nil
	case *ent.DictionaryQuery:
		return &query[*ent.DictionaryQuery, predicate.Dictionary, dictionary.OrderOption]{typ: ent.TypeDictionary, tq: q}, nil
	case *ent.DictionaryDetailQuery:
		return &query[*ent.DictionaryDetailQuery, predicate.DictionaryDetail, dictionarydetail.OrderOption]{typ: ent.TypeDictionaryDetail, tq: q}, nil
	case *ent.MedicineQuery:
		return &query[*ent.MedicineQuery, predicate.Medicine, medicine.OrderOption]{typ: ent.TypeMedicine, tq: q}, nil
	case *ent.MenuQuery:
		return &query[*ent.MenuQuery, predicate.Menu, menu.OrderOption]{typ: ent.TypeMenu, tq: q}, nil
	case *ent.OauthProviderQuery:
		return &query[*ent.OauthProviderQuery, predicate.OauthProvider, oauthprovider.OrderOption]{typ: ent.TypeOauthProvider, tq: q}, nil
	case *ent.PositionQuery:
		return &query[*ent.PositionQuery, predicate.Position, position.OrderOption]{typ: ent.TypePosition, tq: q}, nil
	case *ent.RoleQuery:
		return &query[*ent.RoleQuery, predicate.Role, role.OrderOption]{typ: ent.TypeRole, tq: q}, nil
	case *ent.TokenQuery:
		return &query[*ent.TokenQuery, predicate.Token, token.OrderOption]{typ: ent.TypeToken, tq: q}, nil
	case *ent.UserQuery:
		return &query[*ent.UserQuery, predicate.User, user.OrderOption]{typ: ent.TypeUser, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
