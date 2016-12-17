package main

import (
	"fmt"
	"time"
)

func main() {
	//========== channel ==========
	//channel get  1
	//channel get  2
	//channel get  3
	//channel close
	channel()

	//========== channelBuffer1 ==========
	//1
	//2
	//3
	channelBuffer1()

	//========== channelBuffer2 ==========
	//0   2016-12-17 15:10:26.3476449 +0900 JST
	//1   2016-12-17 15:10:26.3868289 +0900 JST
	//2   2016-12-17 15:10:26.8871794 +0900 JST
	//3   2016-12-17 15:10:26.8871794 +0900 JST
	//4   2016-12-17 15:10:27.3877214 +0900 JST
	//5   2016-12-17 15:10:27.3877214 +0900 JST
	//6   2016-12-17 15:10:27.8886089 +0900 JST
	//7   2016-12-17 15:10:27.8886089 +0900 JST
	//8   2016-12-17 15:10:28.3888282 +0900 JST
	//9   2016-12-17 15:10:28.3888282 +0900 JST
	channelBuffer2()
}

// 通常のチャンネル実行
func channel() {
	fmt.Println("========== channel ==========")
	ch := make(chan int)

	// Error: バッファなしの場合、gorutine内でない場合エラーになる
	//	ch <- 1

	// chは引数で使いまわせる
	// 無名関数なので、別に引数で持ちまわす必要はない(今回は引数)
	go func(_ch chan int) {
		_ch <- 1
		_ch <- 2
		_ch <- 3
		close(_ch)
	}(ch)

	for {
		// closeをキャッチするとokがfalseになる
		i, ok := <-ch
		if !ok {
			fmt.Println("channel close")
			break
		}

		// closeがcloseと同時にキャッチされないことを確認するためにsleep
		time.Sleep(500 * time.Millisecond)
		fmt.Println("channel get ", i)
	}
}

// チャンネルバッファサンプル
func channelBuffer1() {
	fmt.Println("========== channelBuffer1 ==========")

	// バッファ分だけ値を突っ込めるようになる
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	// 後から取得可能
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(ch)
}

// チャンネルバッファサンプル
// 同時実行数を制御
func channelBuffer2() {
	fmt.Println("========== channelBuffer2 ==========")

	// 同時実行する数とバッファを作成
	// バッファ分しか同時に値を入れられないのでそれを利用
	parallelNum := 2
	ch := make(chan bool, parallelNum)

	for i := 0; i < 10; i++ {
		// バッファ分しか追加できないのでここで停止
		ch <- true
		fmt.Println(i, " ", time.Now())

		// なんか同時に処理したい処理
		go func(i int) {
			time.Sleep(500 * time.Millisecond)
			// 終わったことを通知
			<-ch
		}(i)
	}
}
