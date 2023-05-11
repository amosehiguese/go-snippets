package snippet

import (
	"context"
	"fmt"
	"os"
	"time"
)

var gen = func (ctx context.Context, values ...any ) <-chan any  {
	valueStream := make(chan any)

	go func() {
		defer close(valueStream)

		for v := range values  {
			select {
			case <-ctx.Done():
				return 
			case valueStream <- v:
			}
		}
	}()
	return valueStream
}

var multiply = func (ctx context.Context, valueStream <-chan any, multiplier int ) <-chan any {
	multipliedStream := make(chan any)

	go func() {
		defer close(multipliedStream)

		for v := range valueStream {
			select {
			case <-ctx.Done():
				return
			case multipliedStream <- v:
			}
		}
	}()
	return  multipliedStream
}

var newadd = func (ctx context.Context, valueStream <-chan any, additive int) <-chan any {
	time.Sleep(200 * time.Millisecond)
	addedStream := make(chan any)

	go func() {
		defer close(addedStream)

		for v := range valueStream {
			select {
			case <-ctx.Done():
				return
			case addedStream <- v:
			}
		}
	}()
	return addedStream
}

func CtxPipeMain()  {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(300 * time.Millisecond))
	defer cancel()
	// ctx, cancel :=context.WithTimeout(ctx, 300 * time.Millisecond)
	
	for v := range multiply(ctx,newadd(ctx,multiply(ctx, gen(ctx, 1,2,3,4,5), 2), 1), 2){
		fmt.Println(v)
	}
	
	if deadline, ok := ctx.Deadline(); ok  {
		if time.Until(deadline) <= 0 {
			cancel()
			fmt.Fprintf(os.Stdout, "Error: %v\n", ctx.Err())
			
		}
	}
}
