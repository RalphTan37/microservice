package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/RalphTan37/microservice/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client //client property
}

// generating order key function
func orderIDKey(id uint64) string {
	return fmt.Sprintf("order: %d", id)
}

// insert function
func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order) //encoding

	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderID) //generates id

	txn := r.Client.TxPipeline() //new transaction

	//marshal method returns byte array, cast to string
	//NX = not exist - client will not override any data that already exists, instead return an error
	res := txn.SetNX(ctx, key, string(data), 0)

	//checks if error occurs
	if err := res.Err(); err != nil {
		txn.Discard() //discard any potential changes
		return fmt.Errorf("failed to set: %w", err)
	}

	//order key to new set
	if err := txn.SAdd(ctx, "orders", key).Err(); err != nil {
		txn.Discard() //discard any potential changes
		return fmt.Errorf("failed to add to orders set: %w", err)
	}

	//level of guarantee, not left in a partial state
	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil //error in order to communicate this w/ caller
}

// custom Redis error
var ErrNotExist = errors.New("order does not exist")

// counterpart to insert method
func (r *RedisRepo) FindByID(ctx context.Context, id uint64) (model.Order, error) {
	key := orderIDKey(id) //generate key from id

	//handles any errors that are received
	value, err := r.Client.Get(ctx, key).Result()

	//checks if error is redis.nil error
	if errors.Is(err, redis.Nil) {
		return model.Order{}, ErrNotExist
	} else if err != nil {
		return model.Order{}, fmt.Errorf("get order: %w", err)
	}

	//convert json data received into model type
	var order model.Order
	err = json.Unmarshal([]byte(value), &order) //make changes to the original order instance

	// error case for json unmarshal
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to decode order json: %w", err)
	}

	return order, nil //returns order instance
}

// delete method
func (r *RedisRepo) DeleteByID(ctx context.Context, id uint64) error {
	key := orderIDKey(id) //generate key from id

	txn := r.Client.TxPipeline() //transaction

	err := txn.Del(ctx, key).Err() //initialize error
	if errors.Is(err, redis.Nil) { //checks error
		txn.Discard()
		return ErrNotExist //returns custom error - operation failed, expected state is still correct
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("get order: %w", err)
	}

	//remove key from the order set
	if err := txn.SRem(ctx, "orders", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to remove from orders set: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

// update method
func (r *RedisRepo) Update(ctx context.Context, order model.Order) error {
	//similar to insert + delete method but only want to update existing records
	data, err := json.Marshal(order)

	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderID)

	//set value if it already exists
	err = r.Client.SetXX(ctx, key, string(data), 0).Err()

	if errors.Is(err, redis.Nil) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("get order: %w", err)
	}

	return nil
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

// defines orders & next cursor
type FindResult struct {
	Orders []model.Order
	Cursor uint64 // caller knows where to pick up from
}

// find all method
func (r *RedisRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	res := r.Client.SScan(ctx, "orders", page.Offset, "*", int64(page.Size))

	keys, cursor, err := res.Result() //capture in a result value
	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get order ids: %w", err)
	}

	//check key size
	if len(keys) == 0 {
		return FindResult{
			Orders: []model.Order{},
		}, nil
	}

	//pass in all keys to single redis call
	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("failed to get orders: %w", err)
	}

	orders := make([]model.Order, len(xs)) //create an order slice w/ same length as resulting slice

	//iterate over each element in result array + cast to a string
	for i, x := range xs {
		x := x.(string)
		var order model.Order

		err := json.Unmarshal([]byte(x), &order)
		if err != nil {
			return FindResult{}, fmt.Errorf("failed to decode order json: %w", err)
		}

		orders[i] = order
	}
	return FindResult{
		Orders: orders,
		Cursor: cursor,
	}, nil
}
