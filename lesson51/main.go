package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	// eg := errgroup.Group{}
	// eg.Go(func() error {
	// 	fmt.Println("go1")
	// 	return nil
	// })
	// eg.Go(func() error {
	// 	fmt.Println("go2")
	// 	err := errors.New("go2 err")
	// 	return err
	// })
	// err := eg.Wait()
	// if err != nil {
	// 	fmt.Println("err =", err)
	// }

	// eg, ctx := errgroup.WithContext(context.Background())
	// eg.Go(func() error {
	// 	time.Sleep(1 * time.Second)
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("go1 cancel, err = ", ctx.Err())
	// 	default:
	// 		fmt.Println("go1 run")
	// 	}
	// 	return nil
	// })
	// eg.Go(func() error {
	// 	err := errors.New("go2 err")
	// 	return err
	// })
	// err := eg.Wait()
	// if err != nil {
	// 	fmt.Println("err =", err)
	// }

	eg := errgroup.Group{}
	eg.SetLimit(2)
	eg.TryGo(func() error {
		fmt.Println("go1 run")
		return nil
	})
	eg.TryGo(func() error {
		err := errors.New("go2 err")
		return err
	})
	eg.TryGo(func() error {
		fmt.Println("go3 run")
		return nil
	})
	err := eg.Wait()
	if err != nil {
		fmt.Println("err =", err)
	}
}
