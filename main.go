package goroutine_example

import (
	"log"
)

type Point struct {
	X int64
	Y int64
}

func main() {
}

// 戻り値がなく、処理が終わっていればよい場合
func Goroutine1() {

	// 値の受け渡しが不要な場合のchannelには空structがよく使われます。サイズがゼロだからです。
	done := make(chan struct{}, 0)

	go func() {
		// 何か処理をする
		defer close(done)
	}()

	// chanがcloseされるまでブロックする
	<-done
}

// 戻り値が1つ存在する
func Goroutine2() {
	intChan := make(chan int, 1)

	go func() {
		// 何か処理をする
		someValue := 123456789
		intChan <- someValue
	}()

	value := <-intChan
	log.Println(value)
}

//　戻り値が2つ存在する
func Goroutine3() {
	intChan := make(chan int, 1)
	errChan := make(chan error, 1)

	go func() {
		// 何か処理をする
		/*
			if err != nil {
				errChan <- err
			} else {
				intChan <- someValue
			}
		*/
	}()

	select {
	case value := <-intChan:
		// 処理が成功した場合の処理
		log.Println(value)
	case err := <-errChan:
		// 処理が失敗した場合の処理
		log.Println(err)
	}
}

//　戻り値が3つ存在する
func Goroutine4() {
	errChan := make(chan error, 1)
	pointChan := make(chan Point, 1)

	go func() {
		// 何か処理をする
		/*
			p, err := someFunc()
			errChan <- err
			pointChan <- p
		*/
	}()

	err := <-errChan
	point := <-pointChan
	if err != nil {
		// 処理が失敗した場合の処理
	} else {
		// 処理が成功した場合の処理
		log.Println(point)
	}
}
