package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/dun-io/imhub/types"
	v1base "github.com/dun-io/imhub/types/v1"
	"github.com/dun-io/imhub/x/node/types"
	"github.com/dun-io/imhub/x/node/types/v1"
	protobuf "github.com/gogo/protobuf/types"
	"time"
)

// SetActiveNode stores an active node in the module's KVStore.
func (k *Keeper) SetActiveNode(ctx sdk.Context, node v1.Node) {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(node.GetAddr())
	value := k.cdc.MustMarshal(&node)

	store.Set(key, value)
}

// HasActiveNode checks if an active node exists in the module's KVStore based on the node address.
func (k *Keeper) HasActiveNode(ctx sdk.Context, addr base.NodeAddress) bool {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)

	return store.Has(key)
}

// GetActiveNode retrieves an active node from the module's KVStore based on the node address.
// If the active node exists, it returns the node and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetActiveNode(ctx sdk.Context, addr base.NodeAddress) (v v1.Node, found bool) {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// DeleteActiveNode removes an active node from the module's KVStore based on the node address.
func (k *Keeper) DeleteActiveNode(ctx sdk.Context, addr base.NodeAddress) {
	store := k.Store(ctx)
	key := types.ActiveNodeKey(addr)

	store.Delete(key)
}

// SetInactiveNode stores an inactive node in the module's KVStore.
func (k *Keeper) SetInactiveNode(ctx sdk.Context, node v1.Node) {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(node.GetAddr())
	value := k.cdc.MustMarshal(&node)

	store.Set(key, value)
}

// HasInactiveNode checks if an inactive node exists in the module's KVStore based on the node address.
func (k *Keeper) HasInactiveNode(ctx sdk.Context, addr base.NodeAddress) bool {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)

	return store.Has(key)
}

// GetInactiveNode retrieves an inactive node from the module's KVStore based on the node address.
// If the inactive node exists, it returns the node and 'found' as true; otherwise, it returns 'found' as false.
func (k *Keeper) GetInactiveNode(ctx sdk.Context, addr base.NodeAddress) (v v1.Node, found bool) {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)
	value := store.Get(key)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

// DeleteInactiveNode removes an inactive node from the module's KVStore based on the node address.
func (k *Keeper) DeleteInactiveNode(ctx sdk.Context, addr base.NodeAddress) {
	store := k.Store(ctx)
	key := types.InactiveNodeKey(addr)

	store.Delete(key)
}

// SetNode stores a node in the module's KVStore based on its status (active or inactive).
func (k *Keeper) SetNode(ctx sdk.Context, node v1.Node) {
	switch node.Status {
	case v1base.StatusActive:
		k.SetActiveNode(ctx, node)
	case v1base.StatusInactive:
		k.SetInactiveNode(ctx, node)
	default:
		panic(fmt.Errorf("failed to set the node %v", node))
	}
}

// HasNode checks if a node exists in the module's KVStore based on the node address.
func (k *Keeper) HasNode(ctx sdk.Context, addr base.NodeAddress) bool {
	return k.HasActiveNode(ctx, addr) || k.HasInactiveNode(ctx, addr)
}

// GetNode retrieves a node from the module's KVStore based on the node address.
// It checks both active and inactive nodes and returns the node if found.
func (k *Keeper) GetNode(ctx sdk.Context, addr base.NodeAddress) (node v1.Node, found bool) {
	node, found = k.GetActiveNode(ctx, addr)
	if found {
		return
	}

	node, found = k.GetInactiveNode(ctx, addr)
	if found {
		return
	}

	return node, false
}

// GetNodes retrieves all nodes stored in the module's KVStore.
//func (k *Keeper) GetNodes(ctx sdk.Context) (items []v1.Node) {
//	store := k.Store(ctx)
//	iterator := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
//
//	defer iterator.Close()
//
//	for ; iterator.Valid(); iterator.Next() {
//		var item v1.Node
//		k.cdc.MustUnmarshal(iterator.Value(), &item)
//
//		items = append(items, item)
//	}
//
//	return items
//}

func (k *Keeper) SetNodeForInactiveAt(ctx sdk.Context, at time.Time, addr base.NodeAddress) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.NodeForInactiveAtKey(at, addr)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}
